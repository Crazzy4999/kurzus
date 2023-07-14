INSERT INTO users (id, first_name, last_name, email, password) VALUES (1 'Álmos', 'Fábián', 'fabian.almos77@gmail.com', '$2a$10$FGzCU3hQt14pusBK7a5/IuW6XkEdZ9ydSq3YiR2ASAVRiFk/X.Fqm');



INSERT INTO suppliers (id, type, image, name, email, password, description, delivery_time, delivery_fee, opening, closing) VALUES (1, 1, 'https://images.deliveryhero.io/image/Fd-hu/LH/i5zJ-hero.jpg?width=2500&height=625&quality=45', 'Pizza King', 'supplier.pizza.king@example.com', 'PizzaPassword123', 'Pizza King the pizza house choice for the real kings', 45, 750, '7:00', '21:00');
INSERT INTO suppliers (id, type, image, name, email, password, description, delivery_time, delivery_fee, opening, closing) VALUES (2, 2, 'https://images.deliveryhero.io/image/Fd-hu/LH/im8f-hero.jpg', 'Tesco', 'supplier.tesco@example.com', 'SupermarketPassword123', 'Tesco to get everything you need at home', 30, 500, '5:00', '23:00');
INSERT INTO suppliers (id, type, image, name, email, password, description, delivery_time, delivery_fee, opening, closing) VALUES (3, 3, 'https://images.deliveryhero.io/image/Fd-hu/LH/o4sx-hero.jpg?width=2500&height=625&quality=45', 'Starbucks', 'supplier.starbucks@example.com', 'CoffeePassword123', 'Starbucks, best coffee in the world!', 60, 1000, '9:00', '22:00');



INSERT INTO items (id, ingredient) VALUES (1, 'coca-cola');
INSERT INTO items (id, ingredient) VALUES (2, 'tomato-base');
INSERT INTO items (id, ingredient) VALUES (3, 'mozzarella');
INSERT INTO items (id, ingredient) VALUES (4, 'ham');
INSERT INTO items (id, ingredient) VALUES (5, 'corn');
INSERT INTO items (id, ingredient) VALUES (6, 'onion');
INSERT INTO items (id, ingredient) VALUES (7, 'feta cheese');
INSERT INTO items (id, ingredient) VALUES (8, 'bbq sauce');
INSERT INTO items (id, ingredient) VALUES (9, 'mushrooms');
INSERT INTO items (id, ingredient) VALUES (10, 'brown sugar');
INSERT INTO items (id, ingredient) VALUES (11, 'milk');
INSERT INTO items (id, ingredient) VALUES (12, 'pumpkin spice');
INSERT INTO items (id, ingredient) VALUES (13, 'caramel syrup');
INSERT INTO items (id, ingredient) VALUES (14, 'vanilla syrup');
INSERT INTO items (id, ingredient) VALUES (15, 'red bean');
INSERT INTO items (id, ingredient) VALUES (16, 'arabica coffee');
INSERT INTO items (id, ingredient) VALUES (17, 'ice');
INSERT INTO items (id, ingredient) VALUES (18, 'black olive');
INSERT INTO items (id, ingredient) VALUES (19, 'gyros');
INSERT INTO items (id, ingredient) VALUES (20, 'honey');
INSERT INTO items (id, ingredient) VALUES (21, 'puff pastry');
INSERT INTO items (id, ingredient) VALUES (22, 'walnut');
INSERT INTO items (id, ingredient) VALUES (23, 'pistachio');
INSERT INTO items (id, ingredient) VALUES (24, 'butter');
INSERT INTO items (id, ingredient) VALUES (25, 'whipped cream');
INSERT INTO items (id, ingredient) VALUES (26, 'flour');
INSERT INTO items (id, ingredient) VALUES (27, 'egg');
INSERT INTO items (id, ingredient) VALUES (28, 'chocolate');
INSERT INTO items (id, ingredient) VALUES (29, 'vanilla');
INSERT INTO items (id, ingredient) VALUES (30, 'white sugar');



INSERT INTO items_menus (item_id, menu_id) VALUES (1, 1);
INSERT INTO items_menus (item_id, menu_id) VALUES (11, 2);
INSERT INTO items_menus (item_id, menu_id) VALUES (15, 3);
INSERT INTO items_menus (item_id, menu_id) VALUES (2, 4);
INSERT INTO items_menus (item_id, menu_id) VALUES (3, 4);
INSERT INTO items_menus (item_id, menu_id) VALUES (4, 4);
INSERT INTO items_menus (item_id, menu_id) VALUES (5, 4);
INSERT INTO items_menus (item_id, menu_id) VALUES (9, 4);
INSERT INTO items_menus (item_id, menu_id) VALUES (2, 5);
INSERT INTO items_menus (item_id, menu_id) VALUES (19, 5);
INSERT INTO items_menus (item_id, menu_id) VALUES (7, 5);
INSERT INTO items_menus (item_id, menu_id) VALUES (6, 5);
INSERT INTO items_menus (item_id, menu_id) VALUES (18, 5);
INSERT INTO items_menus (item_id, menu_id) VALUES (3, 5);
INSERT INTO items_menus (item_id, menu_id) VALUES (1, 6);
INSERT INTO items_menus (item_id, menu_id) VALUES (20, 7);
INSERT INTO items_menus (item_id, menu_id) VALUES (21, 7);
INSERT INTO items_menus (item_id, menu_id) VALUES (22, 7);
INSERT INTO items_menus (item_id, menu_id) VALUES (23, 7);
INSERT INTO items_menus (item_id, menu_id) VALUES (24, 7);
INSERT INTO items_menus (item_id, menu_id) VALUES (14, 7);
INSERT INTO items_menus (item_id, menu_id) VALUES (11, 8);
INSERT INTO items_menus (item_id, menu_id) VALUES (13, 8);
INSERT INTO items_menus (item_id, menu_id) VALUES (16, 8);
INSERT INTO items_menus (item_id, menu_id) VALUES (17, 8);
INSERT INTO items_menus (item_id, menu_id) VALUES (25, 8);
INSERT INTO items_menus (item_id, menu_id) VALUES (11, 9);
INSERT INTO items_menus (item_id, menu_id) VALUES (24, 9);
INSERT INTO items_menus (item_id, menu_id) VALUES (26, 9);
INSERT INTO items_menus (item_id, menu_id) VALUES (27, 9);
INSERT INTO items_menus (item_id, menu_id) VALUES (28, 9);
INSERT INTO items_menus (item_id, menu_id) VALUES (29, 9);
INSERT INTO items_menus (item_id, menu_id) VALUES (30, 9);



INSERT INTO menus (id, image, supplier_id, category_id, price, name) VALUES (1, 'https://images.deliveryhero.io/image/darsktores-cz/Products/5449000130389.jpg?height=480&dpi=1', 6, 1,  599, 'Coca-Cola 1.75l');
INSERT INTO menus (id, image, supplier_id, category_id, price, name) VALUES (2, 'https://images.deliveryhero.io/image/fd-hu/Products/DMART/Sole-Mizo/5998200138607.jpg?height=480&dpi=1', 6, 1,  700, 'Mizo milk 2.8% 1l');
INSERT INTO menus (id, image, supplier_id, category_id, price, name) VALUES (3, 'https://images.deliveryhero.io/image/fd-hu/Products/DMART/Bonduelle/3083680009508.jpg?height=480&dpi=1', 6, 5,  850, 'Red Bean');
INSERT INTO menus (id, image, supplier_id, category_id, price, name) VALUES (4, 'https://images.deliveryhero.io/image/fd-hu/Products/518403.png?width=1200', 4, 2, 2100, 'Songoku Pizza 24cm');
INSERT INTO menus (id, image, supplier_id, category_id, price, name) VALUES (5, 'https://images.deliveryhero.io/image/fd-hu/Products/518420.png?width=1200%22', 4, 2, 3200, 'Diran 32cm');
INSERT INTO menus (id, image, supplier_id, category_id, price, name) VALUES (6, 'https://images.deliveryhero.io/image/fd-hu/Products/520623.png?width=1200', 4, 1,  299, 'Coca-Cola 0,33l');
INSERT INTO menus (id, image, supplier_id, category_id, price, name) VALUES (7, 'https://images.deliveryhero.io/image/fd-hu/Products/520546.png?width=1200', 4, 4,  500, 'Baklava');
INSERT INTO menus (id, image, supplier_id, category_id, price, name) VALUES (8, 'https://images.deliveryhero.io/image/fd-hu/Products/233884.jpg?width=1200', 5, 3, 1790, 'Caramel Frappuccino (Tall)');
INSERT INTO menus (id, image, supplier_id, category_id, price, name) VALUES (9, 'https://images.deliveryhero.io/image/fd-hu/Products/234111.png?width=1200', 5, 4, 1390, 'Chocolate cake');



INSERT INTO categories (id, name) VALUES (1, 'Drink');
INSERT INTO categories (id, name) VALUES (2, 'Pizza');
INSERT INTO categories (id, name) VALUES (3, 'Coffee');
INSERT INTO categories (id, name) VALUES (4, 'Dessert');
INSERT INTO categories (id, name) VALUES (5, 'Vegan');



INSERT INTO supplier_types (id, type) VALUES (1, 'restaurant');
INSERT INTO supplier_types (id, type) VALUES (2, 'supermarket');
INSERT INTO supplier_types (id, type) VALUES (3, 'coffee_shop');



INSERT INTO order_status (id, status) VALUES (1, 'Created');
INSERT INTO order_status (id, status) VALUES (2, 'Delivering');
INSERT INTO order_status (id, status) VALUES (3, 'Done');



INSERT INTO drivers (id, first_name, last_name, email, password, is_delivering) VALUES (1, 'Driver', 'Driver', 'driver@driver.com', 'DriverPassword123', false);
INSERT INTO drivers (id, first_name, last_name, email, password, is_delivering) VALUES (2, 'Driver1', 'Driver', 'driver1@driver.com', 'DriverPassword123', false);
INSERT INTO drivers (id, first_name, last_name, email, password, is_delivering) VALUES (3, 'Driver2', 'Driver', 'driver2@driver.com', 'DriverPassword123', false);