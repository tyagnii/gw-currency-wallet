CREATE TABLE wallets
(
    id         serial PRIMARY KEY,
    uuid       varchar NOT NULL UNIQUE,
    balanceRUB float DEFAULT 0,
    balanceUSD float DEFAULT 0,
    balanceEUR float DEFAULT 0
);

CREATE TABLE users
(
    id serial PRIMARY KEY,
    username varchar NOT NULl UNIQUE ,
    password varchar NOT NULL,
    email varchar NOT NULL UNIQUE
);

ALTER TABLE users
    ADD FOREIGN KEY (wallet_id) REFERENCES wallets(id);