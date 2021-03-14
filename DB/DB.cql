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
) WITH COMMENT = 'This is master entity of user. All data of sub entities must be present on this table';

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
) WITH COMMENT = 'This is master entity of user. All data of sub entities must be present on this table';

CREATE TABLE IF NOT EXISTS tickets
(
    title     VARCHAR,
    id        UUID,
    subject   VARCHAR,
    status    TINYINT,
    is_closed BOOLEAN,
    user_id   UUID,
    PRIMARY KEY ( user_id, is_closed )
) WITH CLUSTERING ORDER BY (is_closed ASC)
   AND COMMENT = 'This is master entity of tickets. All data of sub entities must be present on this table';

CREATE TABLE IF NOT EXISTS tickets_pk_id
(
    title     VARCHAR,
    id        UUID,
    subject   VARCHAR,
    status    TINYINT,
    is_closed BOOLEAN,
    user_id   UUID,
    PRIMARY KEY ( id )
) WITH COMMENT = 'This table is used for reading a single ticket. `id` column value is shared between this and master table';

CREATE TABLE IF NOT EXISTS ticket_messages
(
    message   MAP<VARCHAR,VARCHAR>,
    ticket_id UUID,
    sender    BOOLEAN,
    sent_at   TIMESTAMP,
    PRIMARY KEY ( ticket_id, sent_at )
) WITH CLUSTERING ORDER BY (sent_at DESC)
   AND COMMENT = 'This table will hold messages related to every ticket. For field sender `True` indicates
    That the sender is admin and `False` indicates that the message sender is user';

CREATE TABLE IF NOT EXISTS phone_contact_requests
(
    mobile    VARCHAR,
    is_called BOOLEAN,
    PRIMARY KEY ( mobile, is_called )
) WITH CLUSTERING ORDER BY (is_called ASC)