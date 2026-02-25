package repo

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(productID int) (*Product, error)
	List() ([]*Product, error)
	Delete(productID int) error
	Update(product Product) (*Product, error)
}

type productRepo struct {
	productList []*Product
}

// constructor or condtructor function
func NewProductRepo() ProductRepo {
	repo := &productRepo{}
	generateInitialProducts(repo)
	return repo
}

func (r *productRepo) Create(p Product) (*Product, error) {
	p.ID = len(r.productList) + 1
	r.productList = append(r.productList, &p)
	return &p, nil
}

func (r *productRepo) Get(productID int) (*Product, error) {
	for _, p := range r.productList {
		if p.ID == productID {
			return p, nil
		}
	}
	return nil, nil
}

func (r *productRepo) List() ([]*Product, error) {
	return r.productList, nil
}

func (r *productRepo) Delete(productID int) error {
	var temp []*Product
	for _, p := range r.productList {
		if p.ID != productID {
			temp = append(temp, p)
		}
	}
	r.productList = temp
	return nil
}

func (r *productRepo) Update(product Product) (*Product, error) {
	for i, p := range r.productList {
		if p.ID == product.ID {
			r.productList[i] = &product
			return &product, nil
		}
	}
	return nil, nil
}

func generateInitialProducts(r *productRepo) {
	p1 := &Product{
		ID:          1,
		Title:       "Product 1",
		Price:       19.99,
		Description: "This is the first product.",
		ImgUrl:      "http://example.com/product1.jpg",
	}
	p2 := &Product{
		ID:          2,
		Title:       "Product 2",
		Price:       29.99,
		Description: "This is the second product.",
		ImgUrl:      "http://example.com/product2.jpg",
	}
	p3 := &Product{
		ID:          3,
		Title:       "Product 3",
		Price:       39.99,
		Description: "This is the third product.",
		ImgUrl:      "http://example.com/product3.jpg",
	}
	p4 := &Product{
		ID:          4,
		Title:       "Product 4",
		Price:       49.99,
		Description: "This is the fourth product.",
		ImgUrl:      "http://example.com/product4.jpg",
	}
	p5 := &Product{
		ID:          5,
		Title:       "Product 5",
		Price:       59.99,
		Description: "This is the fifth product.",
		ImgUrl:      "http://example.com/product5.jpg",
	}
	r.productList = append(r.productList, p1, p2, p3, p4, p5)
}
