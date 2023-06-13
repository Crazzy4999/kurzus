SELECT * FROM addresses;

SELECT name, email FROM users;

SELECT status FROM order_status;

SELECT DISTINCT house_number FROM addresses;



SELECT * FROM items
WHERE price > 3000;

SELECT * FROM order_status
WHERE status LIKE '%o%';

SELECT * FROM order_status
WHERE status IN ('created', 'done');

SELECT * FROM addresses
WHERE house_number BETWEEN 20 AND 50;

SELECT * FROM addresses
WHERE floor_number IS NULL;

SELECT * FROM items
WHERE name ILIKE 's%';

SELECT * FROM suppliers
WHERE supplier_type = 'restaurant';



--returns all suppliers from Budapest
SELECT * FROM suppliers
WHERE address_id IN (
	SELECT id FROM addresses
	WHERE city_name = 'Budapest'
);

--returns all items supplier number 1 is selling
SELECT * FROM items
WHERE id IN (
	SELECT item_id FROM suppliers_items
	WHERE supplier_id = 1
);

--returns all users from Budapest
SELECT * FROM users
WHERE address_id IN (
	SELECT id FROM addresses
	WHERE city_name = 'Budapest'
);



SELECT * FROM addresses
ORDER BY house_number;



SELECT name FROM users
UNION
SELECT name FROM delivery_drivers;



SELECT *
FROM users u
INNER JOIN orders o
ON u.id = o.user_id;

SELECT *
FROM delivery_drivers dd
INNER JOIN orders o
ON dd.id = o.driver_id;



SELECT *
FROM addresses a
LEFT JOIN orders o
ON a.id = o.address_id;

SELECT *
FROM order_status os
RIGHT JOIN orders o
ON os.id = o.order_status_id;



SELECT MAX(delivery_fee) from orders;



SELECT LOWER(name) from users;



--Users may share one address, example: members of a family
SELECT address_id FROM users
GROUP BY address_id;

--Selects addresses that are used with more than 1 user
SELECT COUNT(id), address_id FROM users
GROUP BY address_id
HAVING 1 < COUNT(id);
