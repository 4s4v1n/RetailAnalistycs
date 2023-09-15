CREATE OR REPLACE FUNCTION fnc_defining_offer_increasing_margin(cnt_group INT, max_churn_rate NUMERIC,
                                                            max_stability_index NUMERIC,max_index_sku NUMERIC,
                                                            margin_share NUMERIC)
RETURNS TABLE
(
    customer_id          INT,
    sku_name             TEXT,
    offer_discount_depth NUMERIC
)
AS
$$
WITH cte_preparing_metrics AS (
    SELECT g.customer_id, sg.group_name, g.group_churn_rate, g.group_stability_index, g.group_minimum_discount,
           max(s4.sku_retail_price - s4.sku_purchase_price) OVER (PARTITION BY g.customer_id, g.group_id, pg.sku_id),
           (count(s4.transaction_store_id) OVER (PARTITION BY pg.sku_id))::FLOAT
               / (count(s4.transaction_store_id) OVER (PARTITION BY g.group_id))::FLOAT AS share_sku_group,
           (sku_retail_price - sku_purchase_price)::FLOAT * margin_share / sku_retail_price AS offer_discount_depth,
           dense_rank() OVER (PARTITION BY g.customer_id ORDER BY g.group_id) AS ranks
    FROM groups AS g
        JOIN customers c2 ON g.customer_id = c2.customer_id
        JOIN cards c4 ON c2.customer_id = c4.customer_id
        JOIN product_grid pg ON g.group_id = pg.group_id
        JOIN stores s4 ON pg.sku_id = s4.sku_id
        JOIN transactions t2 ON c4.customer_card_id = t2.customer_card_id
        JOIN sku_group sg ON pg.group_id = sg.group_id)
SELECT DISTINCT customer_id, group_name, offer_discount_depth
FROM cte_preparing_metrics
WHERE max_churn_rate >= group_churn_rate AND max_stability_index >= group_stability_index AND cnt_group >= Ranks
    AND max_index_sku <= share_sku_group AND offer_discount_depth >= (group_minimum_discount * 1.05)
$$
LANGUAGE sql;
