
-- User table
INSERT INTO "User" (user_id, email, password, username, status) VALUES
('bde21a3b-0f8c-4d5d-9c56-2a3e5bb7d532', 'john.doe@example.com', 'pass123', 'johndoe', TRUE),
('a5e4b57d-6fa9-49e2-ae0d-1f4d22b7365c', 'jane.smith@example.com', 'mypassword', 'janesmith', TRUE);

-- Customer table
INSERT INTO "Customer" (customer_id, user_id, customer_name, customer_address, customer_phone) VALUES
('c6c7ae41-d841-4e87-bb3e-90c5c507e710', 'bde21a3b-0f8c-4d5d-9c56-2a3e5bb7d532', 'John Doe', '123 Elm Street', '5551234567'),
('e3a12c13-4bdf-4506-8e47-9d6f45f20ea9', 'a5e4b57d-6fa9-49e2-ae0d-1f4d22b7365c', 'Jane Smith', '456 Oak Avenue', '5559876543');

-- Market table
INSERT INTO "Market" (market_id, user_id, market_name, market_address, market_phone) VALUES
('d1a86b0c-6e49-4fd4-8f2e-78fa96b4e8c8', 'bde21a3b-0f8c-4d5d-9c56-2a3e5bb7d532', 'Johns Market', '789 Pine Road', '5551122334');

-- Category table
INSERT INTO "Category" (category_id, name) VALUES
('f7c1f28a-9e93-43c9-bf1d-3ad43b28e97c', 'Electronics'),
('a99b9f0c-2ebf-4634-8c7f-2f9c52db4732', 'Home Appliances');

-- Products table
INSERT INTO "Products" (product_id, market_id, category_id, product_name, product_image, keyword, description) VALUES
('a6712b99-24f3-4b4b-9dcf-8a2e2a5e1a47', 'd1a86b0c-6e49-4fd4-8f2e-78fa96b4e8c8', 'f7c1f28a-9e93-43c9-bf1d-3ad43b28e97c', 'Smartphone', '\xFFD8FFE000104A464946000101', 'phone, mobile', 'High-end smartphone with great features'),
('b342d3f6-4ae5-4d79-a14b-7a389fde8a9c', 'd1a86b0c-6e49-4fd4-8f2e-78fa96b4e8c8', 'a99b9f0c-2ebf-4634-8c7f-2f9c52db4732', 'Vacuum Cleaner', '\xFFD8FFE000104A464946000102', 'cleaning, home appliance', 'Efficient vacuum cleaner for home use'),
('c456d7f3-9e8f-4d7b-a6e2-5a2c3f6e7b89', 'd1a86b0c-6e49-4fd4-8f2e-78fa96b4e8c8', 'f7c1f28a-9e93-43c9-bf1d-3ad43b28e97c', 'Laptop', '\xFFD8FFE000104A464946000103', 'computer, electronics', 'Powerful laptop for professional use'),
('d879e3b2-8e9d-4c3b-b4a6-6e5f2b4d3c78', 'd1a86b0c-6e49-4fd4-8f2e-78fa96b4e8c8', 'a99b9f0c-2ebf-4634-8c7f-2f9c52db4732', 'Air Purifier', '\xFFD8FFE000104A464946000104', 'air, purifier, appliance', 'Compact air purifier for better indoor air quality'),
('e123f4c7-4b5d-47a6-a3e2-7e5d4c3b2a9f', 'd1a86b0c-6e49-4fd4-8f2e-78fa96b4e8c8', 'f7c1f28a-9e93-43c9-bf1d-3ad43b28e97c', 'Tablet', '\xFFD8FFE000104A464946000105', 'tablet, mobile, electronics', 'Portable tablet for entertainment and productivity');

-- Price table
INSERT INTO "Price" (price_id, product_id, price, stock) VALUES
('8d7628b9-5f3b-42c3-9e87-5f9d52bb6f84', 'a6712b99-24f3-4b4b-9dcf-8a2e2a5e1a47', 799.99, TRUE),
('af7b7cc0-7a8b-4ea6-8a5f-4e2e8a6c7b35', 'b342d3f6-4ae5-4d79-a14b-7a389fde8a9c', 199.99, TRUE),
('9e5f2a4b-3c7d-4f5a-b6e3-7e8d5c9f4a8b', 'c456d7f3-9e8f-4d7b-a6e2-5a2c3f6e7b89', 1299.99, TRUE),
('7d6e8c3b-2f5a-47a6-b4c5-6e3d7f9b2a5f', 'd879e3b2-8e9d-4c3b-b4a6-6e5f2b4d3c78', 149.99, TRUE),
('5e7d4c3a-2f6b-4a8c-b5e7-3d9f4a2b6e5f', 'e123f4c7-4b5d-47a6-a3e2-7e5d4c3b2a9f', 499.99, TRUE);

-- Cart table
INSERT INTO "Cart" (cart_id, customer_id) VALUES
('c81c43d7-7985-4e14-b6eb-6a7e5cde3a20', 'c6c7ae41-d841-4e87-bb3e-90c5c507e710');

-- CartItem table
INSERT INTO "CartItem" (cart_id, price_id, quantity, status) VALUES
('c81c43d7-7985-4e14-b6eb-6a7e5cde3a20', '8d7628b9-5f3b-42c3-9e87-5f9d52bb6f84', 1, 'In Cart');

-- Order table
INSERT INTO "Order" (order_id, customer_id, cart_id, status, date_order, total_price) VALUES
('a52d4c62-786d-4528-90c6-8e5e25d7a8f2', 'c6c7ae41-d841-4e87-bb3e-90c5c507e710', 'c81c43d7-7985-4e14-b6eb-6a7e5cde3a20', 'Processing', '2025-01-10', 799.99);

-- MarketComment table
INSERT INTO "MarketComment" (mcomment_id, customer_id, market_id, star, comment, date) VALUES
('af35b8d2-5e3a-453b-b6e5-9e7b5d2a7e35', 'c6c7ae41-d841-4e87-bb3e-90c5c507e710', 'd1a86b0c-6e49-4fd4-8f2e-78fa96b4e8c8', 4.5, 'Great market with excellent products!', '2025-01-11');

-- ProductComment table
INSERT INTO "ProductComment" (pcomment_id, customer_id, product_id, star, comment, date) VALUES
('bf5e6c77-7d89-453b-8b6e-9c7a5b28f75e', 'c6c7ae41-d841-4e87-bb3e-90c5c507e710', 'a6712b99-24f3-4b4b-9dcf-8a2e2a5e1a47', 5.0, 'Amazing phone with superb performance!', '2025-01-12');

-- Query 0: Retrieve details of a customer's cart and the total price of items in it
SELECT c.customer_name, ci.quantity, p.product_name, pr.price, (ci.quantity * pr.price) AS total_item_price
FROM "Cart" ca
JOIN "Customer" c ON ca.customer_id = c.customer_id
JOIN "CartItem" ci ON ca.cart_id = ci.cart_id
JOIN "Price" pr ON ci.price_id = pr.price_id
JOIN "Products" p ON pr.product_id = p.product_id
WHERE c.customer_name = 'John Doe';


-- Query 1: Retrieve all products with their market and category details
SELECT
    p.product_name,
    p.description,
    m.market_name,
    c.name AS category_name,
    pr.price,
    pr.stock
FROM "Products" p
JOIN "Market" m ON p.market_id = m.market_id
JOIN "Category" c ON p.category_id = c.category_id
JOIN "Price" pr ON p.product_id = pr.product_id;

-- Query 2: Retrieve all orders for a specific customer
SELECT
    o.order_id,
    o.date_order,
    o.status,
    o.total_price,
    c.customer_name
FROM "Order" o
JOIN "Customer" c ON o.customer_id = c.customer_id
WHERE c.customer_name = 'John Doe';

-- Query 3: Fetch all comments for a specific product
SELECT
    pc.star,
    pc.comment,
    pc.date,
    cu.customer_name
FROM "ProductComment" pc
JOIN "Customer" cu ON pc.customer_id = cu.customer_id
WHERE pc.product_id = 'a6712b99-24f3-4b4b-9dcf-8a2e2a5e1a47';

-- Query 4: Retrieve all items in a specific customer's cart
SELECT
    p.product_name,
    pr.price,
    ci.quantity,
    (ci.quantity * pr.price) AS total_item_price
FROM "Cart" ca
JOIN "CartItem" ci ON ca.cart_id = ci.cart_id
JOIN "Price" pr ON ci.price_id = pr.price_id
JOIN "Products" p ON pr.product_id = p.product_id
WHERE ca.customer_id = 'c6c7ae41-d841-4e87-bb3e-90c5c507e710';

-- Query 5: Fetch all products with their comments and average rating
SELECT
    p.product_name,
    p.description,
    AVG(pc.star) AS average_rating,
    COUNT(pc.pcomment_id) AS total_reviews
FROM "Products" p
LEFT JOIN "ProductComment" pc ON p.product_id = pc.product_id
GROUP BY p.product_name, p.description;

-- Query 6: Fetch all markets and their ratings
SELECT
    m.market_name,
    m.market_address,
    AVG(mc.star) AS average_rating,
    COUNT(mc.mcomment_id) AS total_reviews
FROM "Market" m
LEFT JOIN "MarketComment" mc ON m.market_id = mc.market_id
GROUP BY m.market_name, m.market_address;

-- Query 7: Fetch the most recent orders
SELECT
    o.order_id,
    o.date_order,
    o.total_price,
    c.customer_name
FROM "Order" o
JOIN "Customer" c ON o.customer_id = c.customer_id
ORDER BY o.date_order DESC
LIMIT 10;
