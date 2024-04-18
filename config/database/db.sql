CREATE DATABASE zxsd_sales_db;
\c zxsd_sales_db;

CREATE TABLE users(
    id VARCHAR(64) PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    nama VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE tokens(
    id VARCHAR(64) PRIMARY KEY,
    user_id VARCHAR(64) NOT NULL,
    token VARCHAR(255) NOT NULL,
    expired_at DATE NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TYPE jenis_transaksi AS ENUM ('barang', 'jasa');

CREATE TABLE sales(
    id VARCHAR(64) PRIMARY KEY,
    user_id VARCHAR(64) NOT NULL,
    tanggal_transaksi DATE NOT NULL,
    jenis jenis_transaksi NOT NULL,
    nominal FLOAT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);