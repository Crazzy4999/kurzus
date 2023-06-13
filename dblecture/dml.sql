INSERT INTO addresses (country, region, city_name, street_name, house_number, floor_number, postal_code)
VALUES ('HU', 'Pest', 'Budapest', 'Váci utca', 20, null, 1155);

INSERT INTO addresses (country, region, city_name, street_name, house_number, floor_number, postal_code)
VALUES ('HU', 'Pest', 'Budapest', 'Váci utca', 43, null, 1155);

INSERT INTO users (address_id, name, email, password)
VALUES (1, 'Kis Jani', 'kis.janesz@gmail.com', 'passwordhashwithsalt');

INSERT INTO users (address_id, name, email, password)
VALUES (1, 'Nagy Jani', 'nagy.jano@gmail.com', 'passwordhashwithsalt');

INSERT INTO suppliers (address_id, name, supplier_type,  email, password)
VALUES (2, 'Pizza King', 'restaurant', 'pizzaking@gmail.com', 'passwordhashwithsalt');

INSERT INTO items (supplier_id, name, price)
VALUES (1, 'Songoku', 3000);

UPDATE items
SET price = 3200
WHERE id = 1;

INSERT INTO items (supplier_id, name, price)
VALUES (1, 'Négysajtos', 3500);

DELETE FROM items WHERE id = 2;

INSERT INTO delivery_drivers (name, email, password)
VALUES ('Nagy Miklós', 'nagy.miki@gmail.com', 'passwordhashwithsalt');

INSERT INTO order_status (status)
VALUES ('created');

INSERT INTO order_status (status)
VALUES ('in_progress');

INSERT INTO order_status (status)
VALUES ('done');

INSERT INTO order_status (status)
VALUES ('failed');

UPDATE order_status
SET status = 'delivered'
WHERE id = 3;

DELETE FROM order_status WHERE status = 'failed';

INSERT INTO orders (user_id, supplier_id, driver_id, address_id, order_status_id, comment_to_driver, delivery_fee, estimated_delivery_time, delivery_time)
VALUES (1, 1, 1, 1, 1, 'Leave the food at the door please.', 499, '2023-06-12 14:30:00+00', null);

INSERT INTO suppliers_items (supplier_id, item_id)
VALUES (1, 1);

INSERT INTO orders_items (order_id, item_id, quantity)
VALUES (1, 1, 4);
