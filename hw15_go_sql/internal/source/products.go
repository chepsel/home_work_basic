package source

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	ID    int     `json:"id,omitempty"`
	Name  string  `json:"name,omitempty"`
	Price float32 `price:"price,omitempty"`
}

func (ctr *Product) MarshalJSON() ([]byte, error) {
	type dropDefaultInf Product
	result, err := json.Marshal((*dropDefaultInf)(ctr))
	return result, err
}

func (ctr *Product) UnmarshalJSON(data []byte) error {
	type dropDefaultInf Product
	err := json.Unmarshal(data, (*dropDefaultInf)(ctr))
	if err != nil {
		return err
	}
	return nil
}

func (src *Database) UnmarshalJSONProducts(data []byte) ([]Product, error) {
	var ctr []Product
	if err := json.Unmarshal(data, &ctr); err != nil {
		return nil, err
	}
	return ctr, nil
}

func (src *Database) MarshalJSONProducts(ctr []Product) ([]byte, error) {
	result, err := json.Marshal(ctr)
	return result, err
}

func (src *Database) NewProduct() *Product {
	return &Product{}
}

func (src *Database) InsertProduct(product *Product) error {
	queryType := "InsertProduct"
	query := `INSERT INTO products(name, price) VALUES($1, $2) RETURNING id`

	go src.LogDBRequest(queryType, query, fmt.Sprintf("'%s','%f'", product.Name, product.Price)) // Log

	err := src.DB.QueryRow(query, product.Name, product.Price).Scan(&product.ID)
	if err != nil {
		return src.LogError(queryType, err)
	}
	go src.LogDBResult(queryType, fmt.Sprintf("%d", product.ID)) // Log

	if product.ID == 0 {
		return src.LogError(queryType, MissingID)
	}
	return nil
}

func (src *Database) DeleteProduct(id int64) error {
	queryType := "DeleteProduct"
	query := `DELETE FROM products WHERE id=$1`

	go src.LogDBRequest(queryType, query, fmt.Sprintf("'%d'", id)) // Log

	result, err := src.DB.Exec(query, id)
	if err != nil {
		return src.LogError(queryType, err)
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		src.LogError(queryType, NoRowsAffected)
		return NoRowsAffected
	}

	go src.LogDBResult(queryType, fmt.Sprintf("deleted rows(%d)", rowsAffected)) // Log

	return nil
}

func (src *Database) UpdateProduct(product *Product) error {
	queryType := "UpdateProduct"
	query := product.GetUpdateQuery()

	go src.LogDBRequest(queryType, query, fmt.Sprintf("'%s','%f' - where id='%d'",
		product.Name,
		product.Price,
		product.ID,
	)) // Log

	result, err := src.DB.Exec(query, product.ID)
	if err != nil {
		return src.LogError(queryType, err)
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		src.LogError(queryType, NoRowsAffected)
		return NoRowsAffected
	}

	go src.LogDBResult(queryType, fmt.Sprintf("updated rows(%d)", rowsAffected)) // Log

	return nil
}

func (ctr *Product) GetUpdateQuery() string {
	var query string
	switch {
	case len(ctr.Name) > 0 && ctr.Price > 0:
		query = fmt.Sprintf(`UPDATE products SET NAME='%s', PRICE=%f where ID=$1;`, ctr.Name, ctr.Price)
	case len(ctr.Name) > 0:
		query = fmt.Sprintf(`UPDATE products SET NAME='%s' where ID=$1;`, ctr.Name)
	case ctr.Price > 0:
		query = fmt.Sprintf(`UPDATE products SET PRICE=%f where ID=$1;`, ctr.Price)
	}
	return query
}

func (src *Database) GetProduct(userID int64) (Product, error) {
	queryType := "GetProduct"
	var product Product

	query := `SELECT id, name, price FROM products WHERE id=$1`

	go src.LogDBRequest(queryType, query, fmt.Sprintf("%d", userID)) // Log

	err := src.DB.QueryRow(query, userID).Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		return product, src.LogError(queryType, err)
	}
	go src.LogDBResult(queryType, fmt.Sprintf("'%v','%v','%v'",
		product.ID,
		product.Name,
		product.Price,
	)) // Log

	return product, nil
}

func (src *Database) GetProductsList() ([]Product, error) {
	queryType := "GetProductsList"
	var products []Product

	query := `SELECT id, name, price FROM products`

	go src.LogDBRequest(queryType, query, "") // Log

	rows, err := src.DB.Query(query)
	if err != nil {
		return nil, src.LogError(queryType, err)
	}
	defer rows.Close()
	for rows.Next() {
		var product Product
		if err = rows.Scan(&product.ID, &product.Name, &product.Price); err != nil {
			return nil, src.LogError(queryType, err)
		}
		products = append(products, product)
	}
	if err = rows.Err(); err != nil {
		return nil, src.LogError(queryType, err)
	}

	go src.LogDBResult(queryType, fmt.Sprintf("rows found(%d)", len(products))) // Log

	return products, nil
}
