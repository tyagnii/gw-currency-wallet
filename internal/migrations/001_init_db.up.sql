CREATE TABLE users (
                       id serial PRIMARY KEY,
                       name varchar NOT NULL,
                       email varchar NOT NULL,
                       FOREIGN KEY (wallet_id) REFERENCES wallets(id)
);

CREATE TABLE wallets (
                         id serial PRIMARY KEY,
                         uuid varchar(30) NOT NULL,
                         balanceRUB float,
                         balanceUSD float,
                         balanceEUR float
)