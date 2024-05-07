/*Users*/
INSERT INTO application.users ("name",email,"password") VALUES ('Teddy', 'teddy@bear.pro', 'qwerty'),('Ignat', 'ignat@rules.pro', '2wsx'),('removeme', 'killme@please.org', '1234');
UPDATE application.users SET email='ignat@rules.com' WHERE "name"='Ignat';
DELETE FROM application.users WHERE "name"='removeme';
/*Products*/
INSERT INTO application.Products ("name",price) VALUES ('Beer', 23),('Bear', 700.2),('Сhair', 30),('Waffle', 2),('killme',2);
UPDATE application.Products SET price=6 WHERE "name"='Waffle';
DELETE FROM application.Products WHERE "name"='killme';
/*Orders - create\delete*/
--- Insert order transaction
	START TRANSACTION;
	DO
	$$
	DECLARE
	_price INT;
	_order_id INT;
	_user_id INT;
	_product_id INT;
	_product_count INT;
	begin
	/*init vars*/
	SELECT 1 INTO _product_id;
	SELECT 1 INTO _user_id;
	SELECT 2 INTO _product_count;
	/*get actual prices vars*/
	SELECT price INTO _price FROM application.products WHERE id=_product_id;
	IF NOT FOUND 
	THEN
		RAISE EXCEPTION 'NO product FOUND';
	else
		/*create order*/
		INSERT INTO application.orders(user_id,order_date,total_amount) VALUES(_user_id,now(),(_price*_product_count)) RETURNING id INTO _order_id;
	END IF;
	IF _order_id is null 
	THEN 
		RAISE EXCEPTION 'NO order FOUND';
	else
		/*create order-product bunch*/
		INSERT INTO application.orderproducts(order_id, product_id, product_count) VALUES(_order_id,_product_id,_product_count);
	END IF;
	END;
	$$;
	COMMIT;
--- Delete order transaction
	START TRANSACTION;
	DO
	$$
	DECLARE
	_order_id INT;
	begin
	/*init vars*/
	SELECT 2 INTO _order_id;
	/*get actual prices vars*/
	delete from application.orderproducts where order_id=_order_id;
	delete from application.orders where id=_order_id;
	END;
	$$;
	COMMIT;
/*Select*/
SELECT email FROM application.users WHERE id>1
SELECT id, "name", price FROM application.products WHERE lower("name")=lower('waffle');
/*Select user order list*/
SELECT o.id FROM application.orders o
INNER JOIN application.users u on o.user_id=u.id
WHERE u.id=1
/*User stat*/
SELECT SUM(o.total_amount) AS order_sum , AVG(p.price) AS avg_price, u."name"  FROM application.orders o 
INNER JOIN application.orderproducts op ON op.order_id=o.id
INNER JOIN application.products p ON op.product_id=p.id
INNER JOIN application.users u ON o.user_id=u.id
GROUP BY u."name"