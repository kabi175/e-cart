package repository

import (
	"database/sql"
	"log"

	"github.com/kabi175/e-cart/backend/model"
	"github.com/kabi175/e-cart/backend/model/apperror"
)

type ProductRepository struct {
	DB *sql.DB
}

type ProductConfig struct {
	DB *sql.DB
}

func NewProductRepository(c *ProductConfig) model.ProductRepository {
	return &ProductRepository{
		DB: c.DB,
	}
}

func (pr *ProductRepository) Create(p *model.Product) error {
	query := `INSERT INTO products(id, name, seller_id, stock, category) VALUES( $1 , $2 , $3 , $4 , $5 )`

	result, err := pr.DB.Exec(query, p.Id, p.Name, p.SellerID, p.Stock, p.Category)
	if err != nil {
		log.Println(err)
		return apperror.NewInternal()
	}

	rows, err := result.RowsAffected()

	if err != nil {
		log.Println(err)
		return apperror.NewInternal()
	}

	if rows != 1 {
		log.Println("UserRepository.Create: rows affected not equal to 1")
		return apperror.NewInternal()
	}

	return nil
}

func (pr *ProductRepository) Delete(pId string) error {
	query := `DELETE FROM products WHERE id=$1`
	result, err := pr.DB.Exec(query, pId)
	if err != nil {
		log.Println(err)
		return apperror.NewInternal()
	}

	rows, err := result.RowsAffected()

	if err != nil {
		log.Println(err)
		return apperror.NewInternal()
	}
	if rows != 1 {
		return apperror.NewNotFound("product_id", pId)
	}
	return nil
}

func (pr *ProductRepository) FindById(pId string) (*model.Product, error) {
	p := model.Product{}
	query := `SELECT id, seller_id, name, stock, category  
						FROM products 
						WHERE id = $1`

	err := pr.DB.QueryRow(query, pId).Scan(&p.Id, &p.SellerID, &p.Name, &p.Stock, &p.Category)

	if err != nil {

		if err == sql.ErrNoRows {
			return nil, apperror.NewNotFound("product_id", pId)
		}

		return nil, apperror.NewInternal()
	}
	return &p, nil
}

func (pr *ProductRepository) FindByCategory(category string) ([]*model.Product, error) {
	var products []*model.Product
	query := `SELECT id, seller_id, name, stock, category 
						FROM products
						WHERE category = $1`
	result, err := pr.DB.Query(query, category)
	if err != nil {
		return nil, apperror.NewInternal()
	}
	for result.Next() {
		p := &model.Product{}
		result.Scan(&p.Id, &p.SellerID, &p.Name, &p.Stock, &p.Category)
		products = append(products, p)
	}
	return products, nil
}

func (pr *ProductRepository) FindBySeller(sellerId string) ([]*model.Product, error) {
	var products []*model.Product
	query := `SELECT id, seller_id, name, stock, category 
						FROM products
						WHERE seller_id = $1`
	result, err := pr.DB.Query(query, sellerId)
	if err != nil {
		return nil, apperror.NewInternal()
	}
	for result.Next() {
		p := &model.Product{}
		result.Scan(&p.Id, &p.SellerID, &p.Name, &p.Stock, &p.Category)
		products = append(products, p)
	}
	return products, nil
}

func (pr *ProductRepository) Fetch(limit, offset int) ([]*model.Product, error) {
	var products []*model.Product
	query := `SELECT id, seller_id, name, stock, category 
						FROM products 
						WHERE stock > 0 
						LIMIT $1 OFFSET $2`

	result, err := pr.DB.Query(query, limit, offset)
	if err != nil {
		return nil, apperror.NewInternal()
	}

	for result.Next() {
		p := &model.Product{}
		result.Scan(&p.Id, &p.SellerID, &p.Name, &p.Stock, &p.Category)
		products = append(products, p)
	}

	return products, nil
}
