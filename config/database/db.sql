CREATE DATABASE zxsd_sales_db;
\c zxsd_sales_db;

CREATE TABLE users(
    id VARCHAR(32) PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    nama VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE purchase_types(
    id INTEGER PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE sales(
    id VARCHAR(32) PRIMARY KEY,
    user_id VARCHAR(32) NOT NULL,
    tanggal_transaksi DATE NOT NULL,
    jenis_pembelian INTEGER NOT NULL,
    nominal INTEGER NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (jenis_pembelian) REFERENCES purchase_types(id)
);