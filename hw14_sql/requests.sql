/*Users*/
INSERT INTO application.users ("name",email,"password") VALUES ('Teddy', 'teddy@bear.pro', 'qwerty'),('Ignat', 'ignat@rules.pro', '2wsx'),('removeme', 'killme@please.org', '1234');
UPDATE application.users SET email='ignat@rules.com' WHERE "name"='Ignat';
DELETE FROM application.users WHERE "name"='removeme';
/*Products*/
INSERT INTO application.Products ("name",price) VALUES ('Beer', 23),('Bear', 700.2),('Сhair', 30),('Waffle', 2),('killme',2);
UPDATE application.Products SET price=6 WHERE "name"='Waffle';
DELETE FROM application.Products WHERE "name"='killme';
/*UsersOrders*/
INSERT INTO application.orders (user_id,order_date,total_amount) VALUES (1, now(),30),(2, to_timestamp('2022-05-17 10:10:10.022.000001', 'YYYY-MM-DD HH:MI:SS.MS.US'),700.2),(1, to_timestamp('2023-05-17 10:10:10.022.000001', 'YYYY-MM-DD HH:MI:SS.MS.US'),730.2),(2, now(),50);
INSERT INTO application.orderproducts (order_id,product_id) VALUES (1,3),(2,2),(3,2),(3,3);
DELETE FROM application.orders WHERE id=4;
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