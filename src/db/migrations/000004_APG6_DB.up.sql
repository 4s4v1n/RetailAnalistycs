CREATE OR REPLACE FUNCTION sorted_groups()
RETURNS TABLE
(
    customer_id            INT,
    group_id               INT,
    group_churn_rate       NUMERIC,
    group_discount_share   NUMERIC,
    group_minimum_discount NUMERIC,
    av_margin              NUMERIC
)
AS
$$
WITH cte_row_groups AS (
    SELECT *, row_number() OVER (PARTITION BY customer_id ORDER BY group_affinity_index DESC) AS number_id,
           avg(group_margin) OVER (PARTITION BY customer_id, group_id) AS av_margin
    FROM groups)
SELECT customer_id, group_id, group_churn_rate, group_discount_share, group_minimum_discount, av_margin
FROM cte_row_groups
$$
LANGUAGE sql;

CREATE OR REPLACE FUNCTION fnc_defining_group(max_churn_rate NUMERIC, max_discount_share NUMERIC,margin_share NUMERIC)
RETURNS TABLE
(
    customer_id          INT,
    group_id             INT,
    offer_discount_depth NUMERIC
)
AS
$$
DECLARE
    id        INT  := -1;
    is_check  BOOL := TRUE;
    value     RECORD;
    group_cur CURSOR FOR
        (SELECT *
         FROM sorted_groups());
BEGIN
FOR value IN group_cur
    LOOP
        IF is_check != TRUE AND id = value.customer_id THEN
            CONTINUE;
        END IF;
        IF value.group_churn_rate <= max_churn_rate AND value.group_discount_share <= max_discount_share THEN
            IF (value.av_margin * margin_share) >= (value.group_minimum_discount * 1.05) THEN
                customer_id = value.customer_id;
                group_id = value.group_id;
                offer_discount_depth = value.group_minimum_discount * 1.05;
                is_check = FALSE;
                id = customer_id;
                RETURN NEXT;
            ELSE
                is_check = TRUE;
            END IF;
        ELSE
            is_check = TRUE;
        END IF;
    END LOOP;
END;
$$
LANGUAGE plpgsql;

CREATE TYPE calc_method AS ENUM ('PERIOD', 'QUANTITY');

CREATE OR REPLACE FUNCTION fnc_first_transaction_date() RETURNS DATE AS
$$
SELECT transaction_datetime::TIMESTAMP::DATE
FROM transactions
ORDER BY transaction_datetime
LIMIT 1
$$
LANGUAGE sql;

CREATE OR REPLACE FUNCTION fnc_last_transaction_date() RETURNS DATE AS
$$
SELECT transaction_datetime::TIMESTAMP::DATE
FROM transactions
ORDER BY transaction_datetime DESC
LIMIT 1
$$ LANGUAGE sql;

CREATE OR REPLACE FUNCTION fnc_transactions_count() RETURNS INT AS
$$
SELECT count(transaction_id)
FROM transactions
$$
LANGUAGE sql;

CREATE OR REPLACE FUNCTION fnc_get_transaction_date(out_of_end INT) RETURNS DATE AS
$$
SELECT transaction_datetime::TIMESTAMP::DATE
FROM (SELECT *
      FROM transactions
      ORDER BY transaction_datetime DESC
      LIMIT $1) AS tmp
ORDER BY transaction_datetime
LIMIT 1
$$
LANGUAGE sql;

CREATE
OR REPLACE FUNCTION fnc_get_required_check_measure(first DATE, last DATE, coefficient NUMERIC)
RETURNS TABLE
(
    customer_id            INT,
    required_check_measure NUMERIC
)
AS
$$
SELECT customer_id, round(customer_average_check * coefficient, 2) AS required_check_measure
FROM (SELECT customer_id, avg(transaction_sum) AS customer_average_check
      FROM cards c
          JOIN transactions t on c.customer_card_id = t.customer_card_id
      WHERE transaction_datetime::TIMESTAMP::DATE >= first AND transaction_datetime::TIMESTAMP::DATE <= last
      GROUP BY customer_id) AS sad
$$
LANGUAGE sql;

CREATE OR REPLACE FUNCTION fnc_growth_of_average_check(method calc_method, first DATE, last DATE,
                                                       num INT, coefficient NUMERIC, max_churn_rate NUMERIC,
                                                       max_discount_share NUMERIC, margin_share NUMERIC)
RETURNS TABLE
(
    customer_id            INT,
    required_check_measure NUMERIC,
    group_name             TEXT,
    offer_discount_depth   NUMERIC
)
AS
$$
BEGIN
CASE method
WHEN 'PERIOD'
    THEN first = greatest(least(first, last), fnc_first_transaction_date());
         last = least(greatest(first, last), fnc_last_transaction_date());
WHEN 'QUANTITY'
    THEN num = least(num, fnc_transactions_count());
         first = fnc_get_transaction_date(num);
         last = fnc_get_transaction_date(1);
END CASE;

RETURN QUERY SELECT a.customer_id, b.required_check_measure, sku_group.group_name, a.offer_discount_depth
             FROM fnc_get_required_check_measure(first, last, coefficient) b
                 JOIN fnc_defining_group(max_churn_rate, max_discount_share, margin_share) a
                     ON a.customer_id = b.customer_id
                 JOIN sku_group on sku_group.group_id = a.group_id;

END;
$$
LANGUAGE plpgsql;
