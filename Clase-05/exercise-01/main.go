package main

import "exercise-01/product"

func main() {

	salt := product.Product{99, "Salt", 1, "Just salt", "Food"}
	salt.Save(&product.Products)
	product.Products.GetAll()

}
