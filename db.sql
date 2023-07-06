CREATE TABLE addresses (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    is_active BOOLEAN NOT NULL,
    city VARCHAR(64) NOT NULL,
    street VARCHAR(64) NOT NULL,
    house_number VARCHAR(64) NOT NULL,
    zip_code VARCHAR(64) NOT NULL,
    floor_number VARCHAR(64),
    apartment VARCHAR(64),
    CONSTRAINT fk_user_id FOREIGN KEY (user_id)
    REFERENCES users(id)
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(64) NOT NULL,
    last_name VARCHAR(64) NOT NULL,
    email VARCHAR(64) NOT NULL,
    password VARCHAR(64) NOT NULL
);

CREATE TABLE drivers (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(64) NOT NULL,
    last_name VARCHAR(64) NOT NULL,
    email VARCHAR(64) NOT NULL,
    password VARCHAR(64) NOT NULL
);

CREATE TABLE supplier_types (
    id SERIAL PRIMARY KEY,
    type VARCHAR(32) NOT NULL
);

CREATE TABLE suppliers (
    id SERIAL PRIMARY KEY,
    address_id INTEGER,
    type INTEGER NOT NULL,
    image VARCHAR(1024) NOT NULL,
    name VARCHAR(64) NOT NULL,
    last_name VARCHAR(64) NOT NULL,
    email VARCHAR(64) NOT NULL,
    password VARCHAR(64) NOT NULL,
    opening VARCHAR(5) NOT NULL,
    closing VARCHAR(5) NOT NULL,
    CONSTRAINT fk_address_id FOREIGN KEY (address_id)
    REFERENCES addresses(id),
    CONSTRAINT fk_type FOREIGN KEY (type)
    REFERENCES supplier_types(id)
);

CREATE TABLE users_addresses (
    user_id INTEGER NOT NULL,
    address_id INTEGER NOT NULL,
    CONSTRAINT fk_user_id FOREIGN KEY (user_id)
    REFERENCES users(id),
    CONSTRAINT fk_address_id FOREIGN KEY (address_id)
    REFERENCES addresses(id)
);

CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(128) NOT NULL
);

CREATE TABLE menus (
    id SERIAL PRIMARY KEY,
    image VARCHAR(1024) NOT NULL,
    supplier_id INTEGER NOT NULL,
    categorie_id INTEGER NOT NULL,
    price DECIMAL DEFAULT 1.0,
    CONSTRAINT fk_supplier_id FOREIGN KEY (supplier_id)
    REFERENCES suppliers(id),
    CONSTRAINT fk_categorie_id FOREIGN KEY (categorie_id)
    REFERENCES categories(id)
);


CREATE TABLE items (
    id SERIAL PRIMARY KEY,
    ingredient VARCHAR(128) NOT NULL
);

CREATE TABLE items_menus (
    item_id INTEGER NOT NULL,
    menu_id INTEGER NOT NULL,
    CONSTRAINT fk_item_id FOREIGN KEY (item_id)
    REFERENCES items(id),
    CONSTRAINT fk_menu_id FOREIGN KEY (menu_id)
    REFERENCES menus(id)
);

CREATE TABLE order_status (
    id SERIAL PRIMARY KEY,
    status VARCHAR(32) NOT NULL
);

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    supplier_id INTEGER NOT NULL,
    driver_id INTEGER NOT NULL,
    status_id INTEGER NOT NULL,
    note VARCHAR(512),
    delivery_fee DECIMAL NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    delivered_at TIMESTAMPTZ DEFAULT '0001-01-01 00:00:00+00',
    CONSTRAINT fk_user_id FOREIGN KEY (user_id)
    REFERENCES users(id),
    CONSTRAINT fk_supplier_id FOREIGN KEY (supplier_id)
    REFERENCES suppliers(id),
    CONSTRAINT fk_driver_id FOREIGN KEY (driver_id)
    REFERENCES drivers(id),
    CONSTRAINT fk_status_id FOREIGN KEY (status_id)
    REFERENCES order_status(id)
);

CREATE TABLE orders_menus (
    order_id INTEGER NOT NULL,
    menu_id INTEGER NOT NULL,
    quantity INTEGER NOT NULL DEFAULT 1,
    CONSTRAINT fk_order_id FOREIGN KEY (order_id)
    REFERENCES orders(id),
    CONSTRAINT fk_menu_id FOREIGN KEY (menu_id)
    REFERENCES menus(id)
);