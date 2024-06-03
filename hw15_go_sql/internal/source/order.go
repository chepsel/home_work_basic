package source

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"
)

type Order struct {
	ID          int64          `json:"id,omitempty"`
	UserID      int            `json:"userId,omitempty"`
	OrderDate   time.Time      `json:"orderDate,omitempty"`
	TotalAmount float32        `json:"totalAmount,omitempty"`
	Products    []OrderProduct `json:"products,omitempty"`
}

type OrderProduct struct {
	ProductID    int   `json:"productId,omitempty"`
	ProductCount uint8 `json:"productCount,omitempty"`
}

func (ctr *Order) MarshalJSON() ([]byte, error) {
	type dropDefaultInf Order
	result, err := json.Marshal((*dropDefaultInf)(ctr))
	return result, err
}

func (ctr *Order) UnmarshalJSON(data []byte) error {
	type dropDefaultInf Order
	err := json.Unmarshal(data, (*dropDefaultInf)(ctr))
	if err != nil {
		return err
	}
	return nil
}

func (src *Database) UnmarshalJSONOrders(data []byte) ([]Order, error) {
	var ctr []Order
	if err := json.Unmarshal(data, &ctr); err != nil {
		return nil, err
	}
	return ctr, nil
}

func (src *Database) MarshalJSONOrders(ctr []Order) ([]byte, error) {
	result, err := json.Marshal(ctr)
	return result, err
}

func (src *Database) NewOrder() *Order {
	return &Order{}
}

func (src *Database) InsertOrder(order *Order) error {
	var emptyTIme time.Time
	var queryType, query string
	funcName := "InsertOrder"

	// Проверка продукта по id и суммирование цены

	if order.OrderDate == emptyTIme {
		order.OrderDate = time.Now()
	}

	tx, err := src.DB.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})

	defer func() {
		tx.Rollback() // Откат транзакции при завершении
	}()
	if err != nil {
		queryType = fmt.Sprintf("%s.%s", funcName, "beginTx")
		return src.LogError(queryType, err)
	}

	err = src.getTotalAmmount(tx, order)
	if err != nil {
		return src.LogError(funcName, err)
	}

	// Создаем заказ в БД
	queryType = fmt.Sprintf("%s.%s", funcName, "orders.insert")
	query = `INSERT INTO orders(user_id,order_date,total_amount) VALUES($1, $2, $3) RETURNING id`

	go src.LogDBRequest(queryType, query, fmt.Sprintf("'%d','%s','%f'",
		order.UserID,
		order.OrderDate,
		order.TotalAmount,
	)) // Log

	err = tx.QueryRow(query, order.UserID, order.OrderDate, order.TotalAmount).Scan(&order.ID)
	if err != nil {
		return src.LogError(queryType, err)
	}

	go src.LogDBResult(queryType, fmt.Sprintf("new order id(%d)", order.ID)) // Log

	// Наполняем таблицу связок заказ<>продукт
	queryType = fmt.Sprintf("%s.%s", funcName, "orderproducts.insert")
	query = `INSERT INTO orderproducts(order_id, product_id, product_count) VALUES($1, $2, $3)`

	for _, product := range order.Products {
		go src.LogDBRequest(queryType, query, fmt.Sprintf("'%d','%d','%d'",
			order.ID,
			product.ProductID,
			product.ProductCount,
		)) // Log

		result, errExec := tx.Exec(query, order.ID, product.ProductID, product.ProductCount)
		if errExec != nil {
			src.LogError(queryType, errExec)
			return src.LogError(queryType, errExec)
		}

		rowsAffected, _ := result.RowsAffected()
		go src.LogDBResult(queryType, fmt.Sprintf("inserted rows(%d)", rowsAffected)) // Log
		if rowsAffected == 0 {
			src.LogError(queryType, NoRowsAffected)
			return NoRowsAffected
		}
	}

	// Успешное сохранение
	err = tx.Commit()
	if err != nil {
		queryType = fmt.Sprintf("%s.%s", funcName, "commit")
		return src.LogError(queryType, err)
	}
	return nil
}

func (src *Database) getTotalAmmount(tx *sql.Tx, order *Order) error {
	var price float32
	queryType := "getTotalAmmount.price.get"
	query := `SELECT price FROM products WHERE id=$1`

	// Проверка продукта по id и добавление цены
	for _, product := range order.Products {
		id := product.ProductID
		count := float32(product.ProductCount)

		go src.LogDBRequest(queryType, query, fmt.Sprintf("%d", id)) // Log

		err := tx.QueryRow(query, id).Scan(&price)
		if err != nil {
			return src.LogError(queryType, err)
		}

		go src.LogDBResult(queryType, fmt.Sprintf("%f", price)) // Log

		productPrice := (price * count)
		order.TotalAmount += productPrice
	}
	return nil
}

func (src *Database) DeleteOrder(id int64) error {
	var queryType, query string
	funcName := "DeleteOrder"

	tx, err := src.DB.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		queryType = fmt.Sprintf("%s.%s", funcName, "beginTx")
		return src.LogError(queryType, err)
	}
	defer func() {
		tx.Rollback() // Откат транзакции при завершении
	}()

	queryType = fmt.Sprintf("%s.%s", funcName, "orderproducts.delete")
	query = `DELETE FROM orderproducts WHERE order_id=$1`

	go src.LogDBRequest(queryType, query, fmt.Sprintf("%d", id)) // Log

	result, err := tx.Exec(query, id)
	if err != nil {
		return src.LogError(queryType, err)
	}

	rowsAffected, _ := result.RowsAffected()
	go src.LogDBResult(queryType, fmt.Sprintf("deleted rows(%d)", rowsAffected)) // Log
	if rowsAffected == 0 {
		src.LogError(queryType, NoRowsAffected)
		return NoRowsAffected
	}

	queryType = fmt.Sprintf("%s.%s", funcName, "order.delete")
	query = `DELETE FROM orders WHERE id=$1`

	go src.LogDBRequest(queryType, query, fmt.Sprintf("%d", id)) // Log

	result, err = tx.Exec(query, id)
	if err != nil {
		src.LogError(queryType, err)
	}

	rowsAffected, _ = result.RowsAffected()
	go src.LogDBResult(queryType, fmt.Sprintf("deleted rows(%d)", rowsAffected)) // Log
	if rowsAffected == 0 {
		src.LogError(queryType, NoRowsAffected)
		return NoRowsAffected
	}
	err = tx.Commit()
	if err != nil {
		queryType = fmt.Sprintf("%s.%s", funcName, "commit")
		return src.LogError(queryType, err)
	}
	return nil
}

func (src *Database) GetUserOrders(userID int64) ([]Order, error) {
	var queryType, query string
	funcName := "GetUserOrders"
	var orders []Order

	query = `SELECT id, user_id, order_date, total_amount FROM orders WHERE user_id=$1`
	queryType = fmt.Sprintf("%s.%s", funcName, "order.select")

	go src.LogDBRequest(queryType, query, fmt.Sprintf("%d", userID)) // Log

	rows, err := src.DB.Query(query, userID)
	if err != nil {
		return nil, src.LogError(queryType, err)
	}
	defer rows.Close()

	for rows.Next() {
		var tmpOrder Order
		err = rows.Scan(&tmpOrder.ID, &tmpOrder.UserID, &tmpOrder.OrderDate, &tmpOrder.TotalAmount)
		if err != nil {
			return nil, src.LogError(queryType, err)
		}
		tmpOrder.Products, err = src.getOrderProducts(tmpOrder.ID)
		if err != nil {
			return nil, src.LogError(queryType, err)
		}

		orders = append(orders, tmpOrder)
	}
	if err = rows.Err(); err != nil {
		return nil, src.LogError(queryType, err)
	}

	go src.LogDBResult(queryType, fmt.Sprintf("rows found(%d)", len(orders))) // Log

	return orders, nil
}

func (src *Database) GetOrder(orderID int64) (Order, error) {
	var queryType, query string
	funcName := "GetOrder"
	var order Order

	query = `SELECT id, user_id, order_date, total_amount FROM orders WHERE id=$1`
	queryType = fmt.Sprintf("%s.%s", funcName, "order.select")

	go src.LogDBRequest(queryType, query, fmt.Sprintf("%d", orderID)) // Log

	err := src.DB.QueryRow(query, orderID).Scan(&order.ID, &order.UserID, &order.OrderDate, &order.TotalAmount)
	if err != nil {
		return order, src.LogError(queryType, err)
	}
	order.Products, err = src.getOrderProducts(order.ID)
	if err != nil {
		return order, src.LogError(queryType, err)
	}

	go src.LogDBResult(queryType, fmt.Sprintf("%v,%v,%v,%v,%v",
		order.ID,
		order.UserID,
		order.OrderDate,
		order.TotalAmount,
		order.Products,
	)) // Log
	return order, nil
}

func (src *Database) getOrderProducts(orderID int64) ([]OrderProduct, error) {
	var orderProduct []OrderProduct
	queryType := "getTotalAmmount.orderproducts.get"
	query := `SELECT product_id, product_count FROM orderproducts WHERE order_id=$1`

	go src.LogDBRequest(queryType, query, fmt.Sprintf("%d", orderID)) // Log

	rows, err := src.DB.Query(query, orderID)
	if err != nil {
		return nil, src.LogError(queryType, err)
	}
	defer rows.Close()

	for rows.Next() {
		var tmpOrderProduct OrderProduct
		err = rows.Scan(&tmpOrderProduct.ProductID, &tmpOrderProduct.ProductCount)
		if err != nil {
			return nil, src.LogError(queryType, err)
		}
		orderProduct = append(orderProduct, tmpOrderProduct)
	}

	// Проверка продукта по id и добавление цены
	if err = rows.Err(); err != nil {
		return nil, src.LogError(queryType, err)
	}

	go src.LogDBResult(queryType, fmt.Sprintf("rows found(%d)", len(orderProduct))) // Log

	return orderProduct, nil
}
