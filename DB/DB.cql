CREATE KEYSPACE IF NOT EXISTS rabetpal WITH REPLICATION = {'class':'SimpleStrategy','replication_factor':3};

USE rabetpal;

CREATE TABLE IF NOT EXISTS users
(
    name          VARCHAR,
    lastname      VARCHAR,
    username      VARCHAR,
    mobile        VARCHAR,
    phone         VARCHAR,
    address       MAP<VARCHAR,VARCHAR>,
    bank_accounts MAP<VARCHAR,VARCHAR>,
    national_code VARCHAR,
    profile_pic   VARCHAR,
    email         VARCHAR,
    user_type     VARCHAR,
    balance       INT,
    admin         MAP<VARCHAR,VARCHAR>,
    seller        MAP<VARCHAR,VARCHAR>,
    broker        MAP<VARCHAR,VARCHAR>,
    password      VARCHAR,
    PRIMARY KEY (mobile)
) WITH COMMENT = 'This is master entity of user. All data of sub entities must be present on this table too';

CREATE TABLE IF NOT EXISTS products
(
    id           UUID,
    title        VARCHAR,
    photos       MAP<INT,VARCHAR>,
    description  TEXT,
    post         MAP<VARCHAR,VARCHAR>,
    product_type VARCHAR,
    attributes   MAP<VARCHAR,VARCHAR>,
    discount     INT,
    active       BOOLEAN,
    price        INT,
    weight       INT,
    is_suspended MAP<VARCHAR,VARCHAR>,
    views        INT,
    PRIMARY KEY ( id )
) WITH COMMENT = 'This is master entity of user. All data of sub entities must be present on this table'