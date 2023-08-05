CREATE OR REPLACE FUNCTION fnc_defining_offer_conditions(first_date DATE, last_date DATE, value_transaction INT)
RETURNS TABLE
(
    customer_id                 INT,
    start_date                  TIMESTAMP,
    end_date                    TIMESTAMP,
    required_transactions_count INT
)
AS
$$
SELECT customer_id, first_date AS start_date, last_date AS end_date,
       round((last_date - first_date) / customer_frequency) + value_transaction AS required_transactions_count
FROM customers
$$
LANGUAGE sql;

CREATE OR REPLACE FUNCTION fnc_defining_offer_increasing_frequency_visits(first_date DATE, last_date DATE,
                                                                      value_transaction INT, max_churn_rate NUMERIC,
                                                                      max_discount_share NUMERIC, margin_share NUMERIC)
RETURNS TABLE
(
    customer_id                 INT,
    start_date                  TIMESTAMP,
    end_date                    TIMESTAMP,
    required_transactions_count NUMERIC,
    group_name                  TEXT,
    offer_discount_depth        NUMERIC
)
AS
$$
SELECT t1.*, t3.group_name, t2.offer_discount_depth
FROM fnc_defining_offer_conditions(first_date, last_date, value_transaction) AS t1
    JOIN fnc_defining_group(max_churn_rate, max_discount_share, margin_share) t2 ON t1.customer_id = t2.customer_id
    JOIN sku_group t3 ON t3.group_id = t2.group_id
$$
LANGUAGE sql;

