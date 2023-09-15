-- TODO fix error CREATE DATABASE cannot run inside a transaction block)
-- BEGIN;
--     DROP DATABASE IF EXISTS retail_analytics;
-- COMMIT;

DROP TABLE IF EXISTS personal_information CASCADE;
DROP TABLE IF EXISTS cards CASCADE;
DROP TABLE IF EXISTS sku_group CASCADE;
DROP TABLE IF EXISTS product_grid CASCADE;
DROP TABLE IF EXISTS stores CASCADE;
DROP TABLE IF EXISTS transactions CASCADE;
DROP TABLE IF EXISTS checks CASCADE;
DROP TABLE IF EXISTS date_of_analysing_formation CASCADE;

DROP INDEX IF EXISTS transaction_index CASCADE;
DROP INDEX IF EXISTS cards_index CASCADE;
DROP INDEX IF EXISTS product_grid_index CASCADE;
DROP INDEX IF EXISTS stores_index CASCADE;
DROP INDEX IF EXISTS checks_index CASCADE;