package repository

import (
	"database/sql"
	"fmt"
	"go-api/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{connection: connection}
}

func (repository *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "select id, product_name, price FROM product"

	rows, err := repository.connection.Query(query)

	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}

	var productList []model.Product
	var productObj model.Product

	for rows.Next() {
		err = rows.Scan(&productObj.Id, &productObj.Name, &productObj.Price)

		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}
		productList = append(productList, productObj)
	}
	rows.Close()
	return productList, nil
}

func (repository *ProductRepository) CreateProduct(product model.Product) (int, error) {
	var id int

	query, err := repository.connection.Prepare("insert into product" +
		"(product_name, price)" +
		" VALUES ($1,$2) returning id")

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(product.Name, product.Price).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()
	return id, nil
}

func (repository *ProductRepository) GetProductById(id int) (*model.Product, error) {
	query, err := repository.connection.Prepare("select * from product where id = $1")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var product model.Product
	err = query.QueryRow(id).Scan(&product.Id, &product.Name, &product.Price)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	query.Close()
	return &product, nil
}
