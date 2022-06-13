package main

import (
	"net/http"

	"github.com/kodepro/inventoryservice/product"
)

/* var productList []Product

func init() {
	productsJSON := `[
		{
			"productId": 1,
			"manufacturer": "ABC ltd",
			"sku": "ABC123",
			"upc": "123456789000",
			"pricePerUnit": "456.65",
			"quantityOnHand": 5432,
			"productName": "Sticky Note"
		},
		{
			"productId": 2,
			"manufacturer": "XYZ inc",
			"sku": "QWE987",
			"upc": "987654321000",
			"pricePerUnit": "333.52",
			"quantityOnHand": 78965,
			"productName": "Leg Warmers"
		},
		{
			"productId": 3,
			"manufacturer": "QWERTY corp",
			"sku": "IJK456",
			"upc": "321654987000",
			"pricePerUnit": "489.15",
			"quantityOnHand": 45633,
			"productName": "Lamp shade"
		}
	]`

	err := json.Unmarshal([]byte(productsJSON), &productList)

	if err != nil {
		log.Fatal(err)
	}
}

func getNextID() int {
	maxID := -1

	for _, prod := range productList {
		if maxID < prod.ProductID {
			maxID = prod.ProductID
		}
	}

	return maxID + 1
}

func findProductByID(productID int) (*Product, int) {
	for i, prod := range productList {
		if prod.ProductID == productID {
			return &prod, i
		}
	}

	return nil, 0
} */

const apiBasePath = "/api"

func main() {
	//http.Handle("/foo", &fooHandler{Message: "foo called"})
	//http.HandleFunc("/bar", barHandler)

	//http.HandleFunc("/products", productsHandler)
	//http.HandleFunc("/products/", productHandler)

	//Convert to HTTP Handler
	/* productListHandler := http.HandlerFunc(productsHandler)
	productItemHandler := http.HandlerFunc(productHandler)

	http.Handle("/products", middlewareHandler(productListHandler))
	http.Handle("/products/", middlewareHandler(productItemHandler)) */

	product.SetupRoutes(apiBasePath)

	http.ListenAndServe(":5000", nil)
}
