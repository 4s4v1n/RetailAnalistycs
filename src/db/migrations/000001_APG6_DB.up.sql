-- TODO fix error CREATE DATABASE cannot run inside a transaction block
-- BEGIN;
--     CREATE DATABASE retail_analytics;
-- COMMIT;

CREATE TABLE IF NOT EXISTS personal_information
(
    customer_id            SERIAL PRIMARY KEY,
    customer_name          TEXT NOT NULL CHECK (customer_name ~ '[А-Я][а-я -]+'),
    customer_surname       TEXT NOT NULL CHECK (customer_surname ~ '[А-Я][а-я -]+'),
    customer_primary_email TEXT NOT NULL,
    customer_primary_phone TEXT NOT NULL CHECK (customer_primary_phone ~ '([+]7)[0-9]{10}')
);

CREATE TABLE IF NOT EXISTS cards
(
    customer_card_id SERIAL PRIMARY KEY,
    customer_id      BIGINT NOT NULL,

    FOREIGN KEY (customer_id) REFERENCES personal_information (customer_id)
);

CREATE TABLE IF NOT EXISTS sku_group
(
    group_id   SERIAL PRIMARY KEY,
    group_name TEXT NOT NULL CHECK (group_name ~ '\D+')
);

CREATE TABLE IF NOT EXISTS product_grid
(
    sku_id   SERIAL PRIMARY KEY,
    sku_name TEXT   NOT NULL CHECK (sku_name ~ '\D+'),
    group_id BIGINT NOT NULL,

    FOREIGN KEY (group_id) REFERENCES sku_group (group_id)
);

CREATE TABLE IF NOT EXISTS stores
(
    transaction_store_id BIGINT  NOT NULL,
    sku_id               BIGINT  NOT NULL,
    sku_purchase_price   NUMERIC NOT NULL CHECK (sku_purchase_price >= 0),
    sku_retail_price     NUMERIC NOT NULL CHECK (sku_retail_price >= 0),

    UNIQUE (transaction_store_id, sku_id),

    FOREIGN KEY (sku_id) REFERENCES product_grid (sku_id)
);

CREATE TABLE IF NOT EXISTS transactions
(
    transaction_id       SERIAL    PRIMARY KEY,
    customer_card_id     BIGINT    NOT NULL,
    transaction_sum      NUMERIC   NOT NULL CHECK (transaction_sum >= 0),
    transaction_datetime TIMESTAMP NOT NULL,
    transaction_store_id BIGINT    NOT NULL,

    FOREIGN KEY (customer_card_id) REFERENCES cards (customer_card_id)
);

CREATE TABLE IF NOT EXISTS checks
(
    transaction_id BIGINT  NOT NULL,
    sku_id         BIGINT  NOT NULL,
    sku_amount     NUMERIC NOT NULL CHECK (sku_amount >= 0),
    sku_sum        NUMERIC NOT NULL CHECK (sku_sum >= 0),
    sku_sum_paid   NUMERIC NOT NULL CHECK (sku_sum_paid >= 0),
    sku_discount   NUMERIC NOT NULL CHECK (sku_discount >= 0),

    UNIQUE (transaction_id, sku_id),

    FOREIGN KEY (transaction_id) REFERENCES transactions (transaction_id),
    FOREIGN KEY (sku_id) REFERENCES product_grid (sku_id)
);

CREATE TABLE IF NOT EXISTS date_of_analysing_formation
(
    date TIMESTAMP NOT NULL,

    UNIQUE (date)
);

CREATE INDEX IF NOT EXISTS transaction_index ON transactions (transaction_id, customer_card_id);
CREATE INDEX IF NOT EXISTS cards_index ON cards (customer_card_id, customer_id);
CREATE INDEX IF NOT EXISTS product_grid_index ON product_grid (sku_id, group_id);
CREATE INDEX IF NOT EXISTS stores_index ON stores (transaction_store_id, sku_id);
CREATE INDEX IF NOT EXISTS checks_index ON checks (transaction_id, sku_id);

BEGIN;
    SET session_replication_role = REPLICA;
    SET datestyle = dmy;

    COPY personal_information FROM '/tmp/data/mini/personal_data.tsv' DELIMITER E'\t';
    SELECT setval('personal_information_customer_id_seq', (SELECT max(customer_id)
                                                           FROM personal_information));

    COPY cards FROM '/tmp/data/mini/cards.tsv' DELIMITER E'\t';
    SELECT setval('cards_customer_card_id_seq', (SELECT max(customer_card_id)
                                                 FROM cards));

    COPY sku_group FROM '/tmp/data/mini/groups_sku.tsv' DELIMITER E'\t';
    SELECT setval('sku_group_group_id_seq', (SELECT max(group_id)
                                             FROM sku_group));

    COPY product_grid FROM '/tmp/data/mini/sku.tsv' DELIMITER E'\t';
    SELECT setval('product_grid_sku_id_seq', (SELECT max(sku_id)
                                              FROM product_grid));

    COPY stores FROM '/tmp/data/mini/stores.tsv' DELIMITER E'\t';

    COPY transactions FROM '/tmp/data/mini/transactions.tsv' DELIMITER E'\t';
    SELECT setval('transactions_transaction_id_seq', (SELECT max(transaction_id)
                                                      FROM transactions));

    COPY checks FROM '/tmp/data/mini/checks.tsv' DELIMITER E'\t';

    COPY date_of_analysing_formation FROM '/tmp/data/date_of_analysis_formation.tsv' DELIMITER E'\t';

    SET session_replication_role = origin;
COMMIT;