CREATE TABLE IF NOT EXISTS number_segment
(
    segment_id                     SERIAL PRIMARY KEY,
    customer_average_check_segment TEXT   NOT NULL,
    customer_frequency_segment     TEXT   NOT NULL,
    customer_churn_segment         TEXT   NOT NULL
);

BEGIN;
    COPY number_segment FROM '/tmp/data/number_segment.tsv' DELIMITER E'\t';
COMMIT;

CREATE TABLE IF NOT EXISTS margin_setting
(
    method TEXT NOT NULL ,
    param  INT  NOT NULL
);

INSERT INTO margin_setting
VALUES ('Default', 10);

CREATE OR REPLACE FUNCTION fnc_create_or_update_purchase_history() RETURNS VOID AS
$$
CREATE OR REPLACE VIEW purchase_history AS
WITH purchases AS (
    SELECT c.customer_id, t.transaction_id, transaction_datetime, group_id, sku_purchase_price, sku_amount,
           sku_sum, sku_sum_paid
    FROM transactions AS t
        JOIN cards c ON c.customer_card_id = t.customer_card_id
        JOIN personal_information pi ON c.customer_id = pi.customer_id
        JOIN checks c2 ON t.transaction_id = c2.transaction_id
        JOIN product_grid pg ON pg.sku_id = c2.sku_id
        JOIN stores s2 ON pg.sku_id = s2.sku_id AND t.transaction_store_id = s2.transaction_store_id)
SELECT DISTINCT customer_id, transaction_id, transaction_datetime, group_id,
                sum(sku_purchase_price * sku_amount)
                OVER (PARTITION BY customer_id, group_id, transaction_id, transaction_datetime) AS group_cost,
                sum(sku_sum)
                OVER (PARTITION BY customer_id, group_id, transaction_id, transaction_datetime) AS group_sum,
                sum(sku_sum_paid)
                OVER (PARTITION BY customer_id, group_id, transaction_id, transaction_datetime) AS group_sum_paid
FROM purchases;
$$
LANGUAGE sql;

SELECT fnc_create_or_update_purchase_history();

CREATE OR REPLACE FUNCTION fnc_create_or_update_periods() RETURNS VOID AS
$$
CREATE OR REPLACE VIEW periods AS
WITH сte_first_and_last_transaction_date AS (
    SELECT customer_id, group_id,
           min(transaction_datetime) AS first_group_purchase_date,
           max(transaction_datetime) AS last_group_purchase_date
    FROM purchase_history
    GROUP BY customer_id, group_id),
cte_cnt_purchase_date AS (
    SELECT h.customer_id, h.group_id, count(transaction_id) AS group_purchase
    FROM purchase_history AS h
    GROUP BY h.customer_id, h.group_id),
cte_group_frequency AS (
    SELECT p.customer_id, p.group_id, first_group_purchase_date, last_group_purchase_date, group_purchase,
           (extract(EPOCH FROM (last_group_purchase_date - first_group_purchase_date)) / 86400.0 + 1) /
           group_purchase AS group_frequency
    FROM cte_cnt_purchase_date AS p
        JOIN сte_first_and_last_transaction_date td ON p.customer_id = td.customer_id AND p.group_id = td.group_id),
cte_minimum_discount AS (
    SELECT DISTINCT f.*,
                    min(sku_discount / sku_sum) OVER (PARTITION BY ph.customer_id, ph.group_id) AS group_min_discount
    FROM checks AS c
        JOIN purchase_history ph ON c.transaction_id = ph.transaction_id
        JOIN cte_group_frequency f ON ph.group_id = f.group_id AND ph.customer_id = f.customer_id)
SELECT *
FROM cte_minimum_discount;
$$
LANGUAGE sql;

SELECT fnc_create_or_update_periods();

CREATE OR REPLACE FUNCTION fnc_create_or_update_groups() RETURNS VOID AS
$$
CREATE OR REPLACE VIEW groups AS
WITH cte_group AS (
    SELECT DISTINCT ph.customer_id, ph.group_id, ph.group_sum_paid, ph.transaction_datetime, group_cost, row_number()
                    OVER (PARTITION BY ph.customer_id, ph.group_id ORDER BY transaction_datetime DESC) AS row_day,
                    p.group_purchase / (count(ph.transaction_id) OVER (PARTITION BY p.customer_id, p.group_id) + (
                        SELECT count(*)
                        FROM purchase_history AS ph1
                        WHERE ph1.customer_id = ph.customer_id
                            AND ph1.group_id != ph.group_id AND ph1.transaction_datetime
                            BETWEEN p.first_group_purchase_date
                            AND p.last_group_purchase_date))::FLOAT AS group_affinity_index,
                                extract(EPOCH FROM (SELECT *
                                                    FROM date_of_analysing_formation) - (max(transaction_datetime)
                                                        OVER (PARTITION BY p.customer_id, p.group_id))) / 86400.0 /
                    group_frequency AS group_churn_rate,
                    abs(extract(EPOCH FROM transaction_datetime - lag(transaction_datetime, 1)
                        OVER (PARTITION BY p.customer_id, p.group_id ORDER BY transaction_datetime))::FLOAT /
                            86400.0 - group_frequency) / group_frequency AS group_stability_index,
                    count(c.transaction_id)
                        FILTER (WHERE c.sku_discount > 0) OVER (PARTITION BY p.customer_id, p.group_id)::FLOAT /
                    group_purchase AS group_discount_share,
                    coalesce((SELECT min(sku_discount / sku_sum)
                              FROM checks c1
                                  JOIN purchase_history ph2 ON ph2.transaction_id = c1.transaction_id
                              WHERE (sku_discount / sku_sum) > 0
                                  AND ph2.customer_id = ph.customer_id
                                  AND ph2.group_id = ph.group_id),0)                                                                                                            AS Group_Minimum_Discount,
                    avg(group_sum_paid) OVER (PARTITION BY p.customer_id, p.group_id) /
                    avg(group_sum) OVER (PARTITION BY p.customer_id, p.group_id) AS group_average_discount
    FROM purchase_history AS ph
        JOIN periods p ON ph.customer_id = p.customer_id and ph.group_id = p.group_id
        JOIN checks c ON ph.transaction_id = c.transaction_id),
cte_group_margin AS (
    SELECT customer_id, group_id,
    CASE
        WHEN (SELECT method
              FROM margin_setting) = 'transaction'
            THEN (SELECT sum(group_sum_paid - group_cost)
                FILTER (WHERE row_day <= (SELECT param
                                          FROM margin_setting))
                OVER (PARTITION BY customer_id, group_id ORDER BY transaction_datetime DESC))
        WHEN (SELECT method
              FROM margin_setting) = 'period'
            THEN (SELECT sum(group_sum_paid - group_cost)
                FILTER (WHERE (SELECT p.transaction_datetime::DATE - (SELECT param
                                                                      FROM margin_setting)
                               FROM purchase_history AS p
                               ORDER BY transaction_datetime DESC
                               LIMIT 1) <= transaction_datetime::DATE)
                OVER (PARTITION BY customer_id, group_id ORDER BY transaction_datetime DESC))
            ELSE (sum(group_sum_paid - group_cost) OVER (PARTITION BY customer_id, group_id))
            END AS group_margin
    FROM cte_group)
SELECT DISTINCT g.customer_id, g.group_id, group_affinity_index, group_churn_rate,
                coalesce(avg(group_stability_index), 0) AS group_stability_index, group_margin,
                group_discount_share, group_minimum_discount, group_average_discount
FROM cte_group AS g
    JOIN cte_group_margin gm ON g.group_id = gm.group_id AND g.customer_id = gm.customer_id
WHERE group_minimum_discount IS NOT NULL
GROUP BY g.customer_id, g.group_id, group_affinity_index, group_churn_rate, group_margin,
    group_discount_share, group_minimum_discount, group_average_discount;
$$
LANGUAGE sql;

SELECT fnc_create_or_update_groups();

CREATE OR REPLACE FUNCTION fnc_create_or_update_customers() RETURNS VOID AS
$$
CREATE OR REPLACE VIEW customers AS
WITH cte_avg_check AS (
    SELECT customer_id, transaction_datetime, transaction_store_id,
           count(transaction_id) OVER (PARTITION BY customer_id) AS cnt,
           avg(transaction_sum) OVER (PARTITION BY customer_id) AS customer_average_check,
           extract(EPOCH FROM (MAX(transaction_datetime) OVER (PARTITION BY customer_id)
                - min(transaction_datetime) OVER (PARTITION BY customer_id))) / 86400.0 /
                 count(transaction_id) OVER (PARTITION BY customer_id) AS customer_frequency,
           extract(EPOCH FROM ((SELECT *
                                FROM date_of_analysing_formation)
                - max(transaction_datetime) OVER (PARTITION BY customer_id))) / 86400.0 AS customer_inactive_period,
           count(transaction_id) OVER (PARTITION BY customer_id, transaction_store_id)::FLOAT /
                 count(transaction_id) OVER (PARTITION BY customer_id)::FLOAT AS share_transaction,
           lead(transaction_store_id, 1) OVER (PARTITION BY customer_id ORDER BY transaction_datetime DESC) AS lead_1,
           lead(transaction_store_id, 2) OVER (PARTITION BY customer_id ORDER BY transaction_datetime DESC) AS lead_2,
           row_number() OVER (PARTITION BY customer_id ORDER BY transaction_datetime DESC) AS rank
    FROM cards AS c2
        JOIN transactions t ON c2.customer_card_id = t.customer_card_id),
cte_rank_customers AS (
    SELECT c.*,
        CASE
        WHEN percent_rank() OVER (ORDER BY customer_average_check DESC) < 0.1
            THEN 'High'
        WHEN percent_rank() OVER (ORDER BY customer_average_check DESC) < 0.35
            THEN 'Medium'
        ELSE 'Low'
        END AS customer_average_check_segment,
        CASE
        WHEN percent_rank() OVER (ORDER BY customer_frequency) < 0.10
            THEN 'Often'
        WHEN PERCENT_RANK() OVER (ORDER BY customer_frequency) < 0.35
            THEN 'Occasionally'
        ELSE 'Rarely'
        END customer_frequency_segment, customer_inactive_period / customer_frequency AS customer_churn_rate,
        CASE
        WHEN (customer_inactive_period / customer_frequency) BETWEEN 0 AND 2
            THEN 'Low'
        WHEN (customer_inactive_period / customer_frequency) BETWEEN 2 AND 5
            THEN 'Medium'
        ELSE 'High'
        END AS customer_churn_segment,
        max(transaction_store_id)
            OVER (PARTITION BY customer_id ORDER BY transaction_datetime DESC, share_transaction DESC) AS max_share
    FROM cte_avg_check AS c
    WHERE RANK = 1),
cte_number_customers AS (
    SELECT DISTINCT c.customer_id, c.customer_average_check, c.customer_average_check_segment, c.customer_frequency,
                    c.customer_frequency_segment, c.customer_inactive_period, c.customer_churn_rate,
                    customer_churn_segment,
                    (SELECT segment_id
                     FROM number_segment AS s
                     WHERE s.customer_average_check_segment = c.Customer_Average_Check_Segment
                        AND s.customer_churn_segment = c.Customer_Churn_Segment
                        AND s.customer_frequency_segment = c.Customer_Frequency_Segment) AS Customer_Segment,
                    CASE
                    WHEN (transaction_store_id + lead_1 + lead_2) / 3 = transaction_store_id
                        THEN transaction_store_id
                    ELSE max_share
                    END AS Customer_Primary_Store
    FROM cte_rank_customers AS c
    ORDER BY 1)
SELECT *
FROM cte_number_customers;
$$
LANGUAGE sql;

SELECT fnc_create_or_update_customers();
