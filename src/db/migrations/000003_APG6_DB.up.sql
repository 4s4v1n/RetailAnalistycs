CREATE ROLE visitor WITH LOGIN PASSWORD 'visitor';
GRANT CONNECT ON DATABASE retail_analytics TO visitor;
GRANT USAGE ON SCHEMA public TO visitor;
GRANT SELECT ON TABLE cards,
                      checks,
                      date_of_analysing_formation,
                      personal_information,
                      product_grid,
                      sku_group,
                      stores,
                      transactions,
                      number_segment,
                      margin_setting
TO visitor;
GRANT SELECT ON TABLE purchase_history,
                      periods,
                      groups,
                      customers
TO visitor;

CREATE ROLE admin WITH LOGIN PASSWORD 'admin';
GRANT ALL ON DATABASE retail_analytics TO admin;
GRANT USAGE ON SCHEMA public TO admin;
GRANT USAGE ON ALL SEQUENCES IN SCHEMA public TO admin;
GRANT ALL ON TABLE cards,
                   checks,
                   date_of_analysing_formation,
                   personal_information,
                   product_grid,
                   sku_group,
                   stores,
                   transactions,
                   number_segment,
                   margin_setting
TO admin;
GRANT ALL ON TABLE purchase_history,
                   periods,
                   groups,
                   customers
TO admin;

ALTER VIEW IF EXISTS purchase_history OWNER TO admin;
ALTER VIEW IF EXISTS periods          OWNER TO admin;
ALTER VIEW IF EXISTS groups           OWNER TO admin;
ALTER VIEW IF EXISTS customers        OWNER TO admin;