package product

import "fmt"

type Product struct {
	ID          int
	Name        string
	Price       int
	Description string
	Category    string
}

type ProductSlice []Product

var Products ProductSlice = []Product{
	{044, "Divina Commedia", 50, "Poem by Dante Alighieri", "Books"},
	{192, "iPhone 1", 540, "A phone!", "Phones"},
}

// Appends a Product into a slice
func (p Product) Save(products *ProductSlice) {

	//should check errors
	*products = append(*products, p)
}

func (products ProductSlice) GetAll() {

	for _, p := range products {
		fmt.Printf("Product %d: %s - $%d %s %s \n", p.ID, p.Name, p.Price, p.Description, p.Category)
	}

}
func getById(productId int) (Product, string) {

	for _, p := range Products {

		if p.ID == productId {
			return p, ""
		}

	}
	return Product{}, "Error, did not found a product"
}
