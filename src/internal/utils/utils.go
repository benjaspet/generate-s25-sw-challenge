package utils

import (
	"encoding/json"
	"io"
	"os"

	"github.com/benjaspet/generate-s25-software-challenge/src/internal/types"
)

func GetMyProducts() ([]types.Product, error) {
	
	file, _ := os.Open("products.json")
	defer file.Close()

	data, err := io.ReadAll(file)

	if err != nil {
		return nil, err
	}

	var products []types.Product
	json.Unmarshal(data, &products)

	return products, nil
}

func ReverseProducts(products []types.Product) []types.Product {
	reversedProducts := make([]types.Product, len(products))
	for i, j := 0, len(products)-1; i < len(products); i, j = i+1, j-1 {
		reversedProducts[i] = products[j]
	}
	return reversedProducts
}

func ContainsAll(s []string, e []string) bool {
	for _, v := range e {
		for _, v2 := range s {
			if v == v2 {
				return true
			}
		}
	}
	return true
}

func FilterProductsByCategories(products []types.Product, categories []string) []types.Product {
	filteredProducts := []types.Product{}
	seen := make(map[string]bool)

	for _, product := range products {
		for _, category := range categories {
			for _, productCategory := range product.Categories {
				if category == string(productCategory) {
					if !seen[product.ID] {
						filteredProducts = append(filteredProducts, product)
						seen[product.ID] = true
					}
					break
				}
			}
		}
	}

	return filteredProducts
}


func FilterProductsByPriceRange(products []types.Product, min int, max int) []types.Product {
	filteredProducts := []types.Product{}
	for _, product := range products {
		if product.Price >= min && product.Price <= max {
			filteredProducts = append(filteredProducts, product)
		}
	}
	return filteredProducts
}

func FilterProductsByStarRange(products []types.Product, min int, max int) []types.Product {
	filteredProducts := []types.Product{}
	for _, product := range products {
		if product.Stars >= min && product.Stars <= max {
			filteredProducts = append(filteredProducts, product)
		}
	}
	return filteredProducts
}

func SortProductsByKey(products []types.Product, key string) []types.Product {
	comparator := getComparator(key)
	if comparator == nil {
		return products
	}

	for i := 0; i < len(products); i++ {
		for j := i + 1; j < len(products); j++ {
			if comparator(products[i], products[j]) {
				products[i], products[j] = products[j], products[i]
			}
		}
	}

	return products
}

func getComparator(key string) func(a, b types.Product) bool {
	switch key {
	case "name":
		return func(a, b types.Product) bool { return a.Name > b.Name }
	case "price":
		return func(a, b types.Product) bool { return a.Price > b.Price }
	case "stars":
		return func(a, b types.Product) bool { return a.Stars > b.Stars }
	default:
		return nil
	}
}

func ApplyOffsetAndLimit(products []types.Product, offset, limit int) []types.Product {
	
	if offset > len(products) {
		return []types.Product{}
	}
	products = products[offset:]

	if limit > 0 && limit < len(products) {
		products = products[:limit]
	}

	return products
}