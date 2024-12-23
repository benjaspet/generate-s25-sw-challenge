package types

type Product struct {
	ID         string            `json:"id"`
	Name       string            `json:"name"`
	Categories []ProductCategory `json:"categories"`
	Stars      int               `json:"stars"`
	Price      int               `json:"price"`
}

type ProductCategory string

const (
	Electronics    ProductCategory = "electronics"
	Apparel        ProductCategory = "apparel"
	HomeGoods      ProductCategory = "home goods"
	Sports         ProductCategory = "sports"
	Beauty         ProductCategory = "beauty"
	Grocery        ProductCategory = "grocery"
	OfficeSupplies ProductCategory = "office supplies"
	Outdoor        ProductCategory = "outdoor"
	Toys           ProductCategory = "toys"
	Health         ProductCategory = "health"
	Automotive     ProductCategory = "automotive"
	Luxury         ProductCategory = "luxury"
	Books          ProductCategory = "books"
)

type SortableField string

const (
	Name  SortableField = "name"
	Price SortableField = "price"
	Stars SortableField = "stars"
)

type SortOrder string

const (
	Ascending  SortOrder = "asc"
	Descending SortOrder = "desc"
)
