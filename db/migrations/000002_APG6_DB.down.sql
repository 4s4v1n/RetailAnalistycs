DROP TABLE IF EXISTS number_segment CASCADE;
DROP TABLE IF EXISTS margin_setting CASCADE;

DROP VIEW IF EXISTS purchase_history CASCADE;
DROP VIEW IF EXISTS periods          CASCADE;
DROP VIEW IF EXISTS groups           CASCADE;
DROP VIEW IF EXISTS customers        CASCADE;

DROP FUNCTION IF EXISTS fnc_create_or_update_purchase_history CASCADE;
DROP FUNCTION IF EXISTS fnc_create_or_update_periods          CASCADE;
DROP FUNCTION IF EXISTS fnc_create_or_update_groups           CASCADE;
DROP FUNCTION IF EXISTS fnc_create_or_update_customers        CASCADE;
