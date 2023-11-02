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
    password TEXT NOT NULL
);

CREATE TABLE shopping_cart
(
    id          varchar(8) PRIMARY KEY,
    products_id varchar(8) REFERENCES products (id),
    user_id     varchar(8) REFERENCES customers (id)
);

INSERT INTO products(id, product_name, category, price) values('abcdefg', 'Dino', 'mainan' , 130000),
                                                              ('asasadd','Permen', 'makanan', 25000);