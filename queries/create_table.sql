CREATE TABLE "User"(
    user_id uuid PRIMARY KEY,
    email varchar(70),
    password varchar(30),
    username varchar(20) UNIQUE NOT NULL,
    status boolean NOT NULL
);

CREATE TABLE "Customer"(
    customer_id uuid PRIMARY KEY,
    user_id uuid REFERENCES "User"(user_id),
    customer_name varchar(20) NOT NULL,
    customer_address text,
    customer_phone varchar(12) NOT NULL
);

CREATE TABLE "Market"(
    market_id uuid PRIMARY KEY,
    user_id uuid REFERENCES "User"(user_id),
    market_name varchar(20) NOT NULL ,
    market_address text NOT NULL ,
    market_phone varchar(12) NOT NULL
);

CREATE TABLE "Products"(
    product_id uuid PRIMARY KEY,
    market_id uuid REFERENCES "Market"(market_id),
    category_id uuid,
    product_name varchar(40) NOT NULL,
    product_image bytea NOT NULL,
    keyword text,
    description text
);

CREATE TABLE "Cart"(
    cart_id uuid PRIMARY KEY,
    customer_id uuid REFERENCES "Customer"(customer_id)
);

CREATE TABLE "Order"(
    order_id uuid PRIMARY KEY,
    customer_id uuid REFERENCES "Customer"(customer_id),
    cart_id uuid REFERENCES "Cart"(cart_id),
    status text NOT NULL,
    date_order date NOT NULL,
    total_price float NOT NULL
);

CREATE TABLE "Price"(
    price_id uuid PRIMARY KEY,
    product_id uuid REFERENCES "Products"(product_id),
    price float NOT NULL,
    stock boolean NOT NULL
);

CREATE TABLE "Category"(
    category_id uuid PRIMARY KEY,
    name text NOT NULL UNIQUE
);

CREATE TABLE "Color"(
    color_id uuid PRIMARY KEY,
    name text NOT NULL UNIQUE
);

CREATE TABLE "CartItem"(
    cart_id uuid REFERENCES "Cart"(cart_id),
    price_id uuid REFERENCES "Price"(price_id),
    quantity integer NOT NULL,
    status text NOT NULL,
    PRIMARY KEY (cart_id,price_id)
);

CREATE TABLE "OrderItem"(
    order_id uuid REFERENCES "Order"(order_id),
    price_id uuid REFERENCES "Price"(price_id),
    quantity integer NOT NULL,
    price_at_purchase float NOT NULL,
    PRIMARY KEY (order_id,price_id)
);

CREATE TABLE "MarketComment"(
    mcomment_id uuid PRIMARY KEY,
    customer_id uuid REFERENCES "Customer"(customer_id),
    market_id uuid REFERENCES "Market"(market_id),
    star float NOT NULL,
    comment text NOT NULL,
    date date NOT NULL
);

CREATE TABLE "ProductComment"(
    pcomment_id uuid PRIMARY KEY,
    customer_id uuid REFERENCES "Customer"(customer_id),
    product_id uuid REFERENCES "Products"(product_id),
    star float NOT NULL,
    comment text NOT NULL,
    date date NOT NULL
);
