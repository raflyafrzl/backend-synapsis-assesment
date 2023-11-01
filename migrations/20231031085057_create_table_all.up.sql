CREATE TABLE products
(
    id           varchar(8) PRIMARY KEY,
    product_name varchar(45)      NOT NULL,
    category     varchar(30)      NOT NULL,
    price        DOUBLE PRECISION NOT NULL
);

CREATE TABLE customers
(
    id       varchar(8) PRIMARY KEY,
    username varchar (15) NOT NULL,
    password varchar(30) NOT NULL
);

CREATE TABLE shopping_cart
(
    id          varchar(8) PRIMARY KEY,
    products_id varchar(8) REFERENCES products (id),
    user_id     varchar(8) REFERENCES customers (id)
);