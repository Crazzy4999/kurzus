CREATE TABLE addresses (
	id SERIAL PRIMARY KEY,
	country VARCHAR(256) NOT NULL,
	region VARCHAR(256) NOT NULL,
	city_name VARCHAR(256) NOT NULL,
	street_name VARCHAR(256) NOT NULL,
	house_number INTEGER NOT NULL,
	floor_number INTEGER,
	postal_code INTEGER NOT NULL
);

CREATE TABLE users (
	id SERIAL PRIMARY KEY,
	address_id INTEGER NOT NULL,
	name VARCHAR(256) NOT NULL,
	email VARCHAR(256) NOT NULL,
	password VARCHAR(64) NOT NULL,
	created_at TIMESTAMPTZ DEFAULT NOW(),
	updated_at TIMESTAMPTZ DEFAULT NULL,
	CONSTRAINT fk_address_id FOREIGN KEY(address_id) REFERENCES addresses(id)
);

CREATE TABLE suppliers (
	id SERIAL PRIMARY KEY,
	address_id INTEGER NOT NULL,
	name VARCHAR(256) NOT NULL,
	supplier_type VARCHAR(64) NOT NULL,
	email VARCHAR(256) NOT NULL,
	password VARCHAR(64) NOT NULL,
	CONSTRAINT fk_address_id FOREIGN KEY(address_id) REFERENCES addresses(id)
);

CREATE TABLE items (
	id SERIAL PRIMARY KEY,
	supplier_id INTEGER NOT NULL,
	name VARCHAR(256) NOT NULL,
	price DECIMAL NOT NULL,
	CONSTRAINT fk_supplier_id FOREIGN KEY(supplier_id) REFERENCES suppliers(id)
);

CREATE TABLE delivery_drivers (
	id SERIAL PRIMARY KEY,
	name VARCHAR(256) NOT NULL,
	email VARCHAR(256) NOT NULL,
	password VARCHAR(64) NOT NULL,
	created_at TIMESTAMPTZ DEFAULT NOW(),
	updated_at TIMESTAMPTZ DEFAULT NULL
);

CREATE TABLE order_status (
	id SERIAL PRIMARY KEY,
	status VARCHAR(64) NOT NULL
);

CREATE TABLE orders (
	id SERIAL PRIMARY KEY,
	user_id INTEGER NOT NULL,
	supplier_id INTEGER NOT NULL,
	driver_id INTEGER NOT NULL,
	address_id INTEGER NOT NULL,
	order_status_id INTEGER NOT NULL,
	comment_to_driver VARCHAR(512),
	delivery_fee DECIMAL,
	estimated_delivery_time TIMESTAMPTZ NOT NULL,
	delivery_time TIMESTAMPTZ,
	CONSTRAINT fk_user_id FOREIGN KEY(user_id) REFERENCES users(id),
	CONSTRAINT fk_supplier_id FOREIGN KEY(supplier_id) REFERENCES suppliers(id),
	CONSTRAINT fk_driver_id FOREIGN KEY(driver_id) REFERENCES delivery_drivers(id),
	CONSTRAINT fk_address_id FOREIGN KEY(address_id) REFERENCES addresses(id),
	CONSTRAINT fk_order_status_id FOREIGN KEY(order_status_id) REFERENCES order_status(id)
);

CREATE TABLE suppliers_items (
	id SERIAL PRIMARY KEY,
	supplier_id INTEGER NOT NULL,
	item_id INTEGER NOT NULL
);

ALTER TABLE suppliers_items
ADD CONSTRAINT fk_suppliers_id FOREIGN KEY(supplier_id) REFERENCES suppliers(id);

ALTER TABLE suppliers_items
ADD CONSTRAINT fk_item_id FOREIGN KEY(item_id) REFERENCES items(id);

CREATE TABLE orders_items (
	id SERIAL PRIMARY KEY,
	order_id INTEGER NOT NULL,
	item_id INTEGER NOT NULL,
	quantity INTEGER DEFAULT 1
);

ALTER TABLE orders_items
ADD CONSTRAINT fk_order_id FOREIGN KEY(order_id) REFERENCES orders(id);

ALTER TABLE orders_items
ADD CONSTRAINT fk_item_id FOREIGN KEY(item_id) REFERENCES items(id);
