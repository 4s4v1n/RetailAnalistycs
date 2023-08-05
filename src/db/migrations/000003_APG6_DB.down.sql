ALTER VIEW IF EXISTS purchase_history OWNER TO postgres;
ALTER VIEW IF EXISTS periods          OWNER TO postgres;
ALTER VIEW IF EXISTS groups           OWNER TO postgres;
ALTER VIEW IF EXISTS customers        OWNER TO postgres;

REVOKE CONNECT ON DATABASE retail_analytics FROM PUBLIC, visitor;
REVOKE USAGE ON SCHEMA public FROM PUBLIC, visitor;
REVOKE SELECT ON TABLE cards,
                       checks,
                       date_of_analysing_formation,
                       personal_information,
                       product_grid,
                       sku_group,
                       stores,
                       transactions,
                       number_segment,
                       margin_setting
FROM PUBLIC, visitor;
REVOKE SELECT ON TABLE purchase_history,
                       periods,
                       groups,
                       customers
FROM PUBLIC, visitor;
DROP ROLE IF EXISTS visitor;

REVOKE ALL ON DATABASE retail_analytics FROM PUBLIC, admin;
REVOKE USAGE ON SCHEMA public FROM PUBLIC, admin;
REVOKE USAGE ON ALL SEQUENCES IN SCHEMA public FROM admin;
REVOKE ALL ON TABLE cards,
                    checks,
                    date_of_analysing_formation,
                    personal_information,
                    product_grid,
                    sku_group,
                    stores,
                    transactions,
                    number_segment,
                    margin_setting
FROM admin;
REVOKE ALL ON TABLE purchase_history,
                    periods,
                    groups,
                    customers
FROM admin;
DROP ROLE IF EXISTS admin;
