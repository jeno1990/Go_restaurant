package product

import (
	"basic_go_backend/types"
	"database/sql"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetProductByID(id int) (*types.Product, error) {
	return nil, nil
}
func (s *Store) GetProductsByID(ids []int) ([]types.Product, error) {
	return nil, nil
}
func (s *Store) GetProducts() ([]*types.Product, error) {
	query := `SELECT * FROM products`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	products := make([]*types.Product, 0)
	for rows.Next() {
		product, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (s *Store) CreateProduct(payload types.CreateProductPayload) error {
	query := `INSERT INTO products (name, description,image , price,quantity) VALUES (?, ?, ? ,?, ?)`
	_, err := s.db.Exec(query, payload.Name, payload.Description, payload.Image, payload.Price, payload.Quantity)
	if err != nil {
		return err
	}
	return nil
}
func (s *Store) UpdateProduct(p types.Product) error {
	query := `UPDATE products SET name = ?, description = ?, image = ?, price = ?, quantity = ? WHERE id = ?`
	_, err := s.db.Exec(query, p.Name, p.Description, p.Image, p.Price, p.Quantity, p.ID)
	if err != nil {
		return err
	}
	return nil
}

func scanRowsIntoProduct(rows *sql.Rows) (*types.Product, error) {
	product := new(types.Product)

	err := rows.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Image,
		&product.Price,
		&product.Quantity,
		&product.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return product, nil
}
