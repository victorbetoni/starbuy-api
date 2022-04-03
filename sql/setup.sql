CREATE TABLE IF NOT EXISTS users (
    username        VARCHAR(20) PRIMARY KEY NOT NULL,
    email           VARCHAR(32)             NOT NULL,
    name            VARCHAR(64)             NOT NULL,
    gender          INT                     NOT NULL,
    registration    DATE                    NOT NULL,
    birthdate       DATE                    NOT NULL,
    seller          BOOLEAN                 NOT NULL,
    profile_picture VARCHAR(512)            NOT NULL,
    city            VARCHAR(60)             NOT NULL
);

CREATE TABLE IF NOT EXISTS phone_number (
    holder       VARCHAR(20) NOT NULL,
    phone_number VARCHAR(15) NOT NULL,

    FOREIGN KEY (holder) REFERENCES users (username),
    PRIMARY KEY (holder, phone_number)
);

CREATE TABLE IF NOT EXISTS product_categories (
    identifier INTEGER      PRIMARY KEY NOT NULL,
    name       VARCHAR(64)              NOT NULL
);

INSERT INTO product_categories VALUES (1, 'Tecnologia e Informática')  ON CONFLICT (identifier) DO NOTHING;
INSERT INTO product_categories VALUES (2, 'Moda e Acessórios') ON CONFLICT (identifier) DO NOTHING;
INSERT INTO product_categories VALUES (3, 'Casa e Decoração') ON CONFLICT (identifier) DO NOTHING;
INSERT INTO product_categories VALUES (4, 'Livros e Revistas') ON CONFLICT (identifier) DO NOTHING;
INSERT INTO product_categories VALUES (5, 'Papelaria') ON CONFLICT (identifier) DO NOTHING;
INSERT INTO product_categories VALUES (6, 'Brinquedos e Jogos') ON CONFLICT (identifier) DO NOTHING;
INSERT INTO product_categories VALUES (7, 'Música e Instrumentos') ON CONFLICT (identifier) DO NOTHING;

CREATE TABLE IF NOT EXISTS products (
    identifier  VARCHAR(64)    UNIQUE  NOT NULL,
    title       VARCHAR(256)           NOT NULL,
    seller      VARCHAR(20)            NOT NULL,
    price       DECIMAL(10, 2)         NOT NULL,
    stock       INTEGER                NOT NULL,
    category    INTEGER                NOT NULL,
    description VARCHAR(512),

    FOREIGN KEY (seller) REFERENCES users (username),
    PRIMARY KEY (identifier)
);

CREATE TABLE IF NOT EXISTS reviews (
    identifier VARCHAR(64)  NOT NULL PRIMARY KEY,
    product    VARCHAR(64)  NOT NULL,
    user       VARCHAR(20)  NOT NULL,
    msg        VARCHAR(512) NOT NULL,
    rate       INT          NOT NULL,

    FOREIGN KEY (user)    REFERENCES users    (username),
    FOREIGN KEY (product) REFERENCES products (identifier)
)

CREATE TABLE IF NOT EXISTS juridic_id (
    holder     VARCHAR(20) NOT NULL PRIMARY KEY,
    identifier VARCHAR(20) NOT NULL,
    type       INT         NOT NULL,

    FOREIGN KEY (holder) REFERENCES users (username)
);

CREATE TABLE IF NOT EXISTS addresses (
    identifier VARCHAR(64) NOT NULL PRIMARY KEY,
    holder     VARCHAR(20) NOT NULL,
    cep        CHAR(8)     NOT NULL,
    number     INTEGER     NOT NULL,
    complement VARCHAR(60),

    FOREIGN KEY (holder) REFERENCES users (username),
    PRIMARY KEY (holder, cep, number)
);

CREATE TABLE IF NOT EXISTS tags (
    product VARCHAR(64) NOT NULL,
    tag     VARCHAR(15) NOT NULL,

    FOREIGN KEY (product) REFERENCES products (identifier),
    PRIMARY KEY (product, tag)
);

CREATE TABLE IF NOT EXISTS login (
    username VARCHAR(20) PRIMARY KEY NOT NULL,
    password VARCHAR(512)            NOT NULL,

    FOREIGN KEY (username) REFERENCES users (username)
);

CREATE TABLE IF NOT EXISTS shopping_cart (
    holder   VARCHAR(20) NOT NULL,
    product  VARCHAR(64) NOT NULL,
    quantity INTEGER     NOT NULL,

    FOREIGN KEY (holder) REFERENCES users (username),
    FOREIGN KEY (product) REFERENCES products (identifier),
    PRIMARY KEY (holder, product)
);

CREATE TABLE IF NOT EXISTS purchase_log (
    identifier    VARCHAR(64)    NOT NULL,
    holder        VARCHAR(20)    NOT NULL,
    seller        VARCHAR(20)    NOT NULL,
    product       VARCHAR(64)    NOT NULL,
    quantity      INTEGER        NOT NULL,
    price         DECIMAL(10, 2) NOT NULL,

    FOREIGN KEY (holder) REFERENCES users (username)
);

CREATE TABLE IF NOT EXISTS purchase_update (
    purchase    VARCHAR(64) NOT NULL,
    status      INTEGER     NOT NULL,
    update_date DATE        NOT NULL
);

CREATE TABLE IF NOT EXISTS product_images (
    product VARCHAR(64)  NOT NULL,
    url     VARCHAR(512) NOT NULL,

    FOREIGN KEY (product) REFERENCES products (identifier)
);