package repo

import (
	"database/sql"
	"ecommerce/domain"
	"ecommerce/product"

	"github.com/jmoiron/sqlx"
)

type ProductRepo interface {
	product.ProductRepo
}

type productRepo struct {
	db *sqlx.DB
}

// constructor or condtructor function
func NewProductRepo(db *sqlx.DB) ProductRepo {
	return &productRepo{
		db: db,
	}
}

func (r *productRepo) Create(p domain.Product) (*domain.Product, error) {
	query := `
		INSERT INTO products (title, description, price, img_url)
		values ($1, $2, $3, $4)
		RETURNING id
	`
	row := r.db.QueryRow(query, p.Title, p.Description, p.Price, p.ImgUrl)
	err := row.Scan(&p.ID)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *productRepo) Get(productID int) (*domain.Product, error) {
	query := `
		SELECT id, title, description, price, img_url
		FROM products
		WHERE id = $1
	`
	var product domain.Product
	err := r.db.Get(&product, query, productID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &product, nil
}

func (r *productRepo) List() ([]*domain.Product, error) {
	query := `
		SELECT id, title, description, price, img_url
		FROM products
	`
	var products []*domain.Product
	err := r.db.Select(&products, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return products, nil
}

func (r *productRepo) Delete(productID int) error {
	query := `
		DELETE FROM products where id = $1
	`
	_, err := r.db.Exec(query, productID)
	if err != nil {
		return err
	}

	return nil
}

func (r *productRepo) Update(product domain.Product) (*domain.Product, error) {
	query := `
		UPDATE products
		SET title = $1, description = $2, price = $3, img_url = $4
		WHERE id = $5 
	`
	row := r.db.QueryRow(query, product.Title, product.Description, product.Price, product.ImgUrl, product.ID)
	err := row.Err()
	if err != nil {
		return nil, err
	}

	return &product, nil
}
