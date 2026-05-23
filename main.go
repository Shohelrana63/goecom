package main
import (
	"encoding/json"
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About Us")
}

type Product struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Price float64 `json:"price"`
	Description string `json:"description"`
	ImgUrl string `json:"imageUrl"`
}

var productList []Product

func getProductsHandler(w http.ResponseWriter, r *http.Request) {
   w.Header().Set("Access-Control-Allow-Origin", "*") // Allow CORS for all origins
   w.Header().Set("Content-Type", "application/json") // Set content type to JSON

	if r.Method != "GET"{
		http.Error(w, "Please use GET request", 400)
		return
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(productList)
}

func createProductHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // Allow CORS for all origins
	w.Header().Set("Access-Control-Allow-Methods", "POST") // Allow only POST method
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type") // Allow Content-Type header
	w.Header().Set("Content-Type", "application/json") // Set content type to JSON

	if r.Method != "POST"{
		http.Error(w, "Please use POST request", 400)
		return
	}

	var newProduct Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		http.Error(w, "Please send valid JSON in  request body", 400)
		return
	}
	newProduct.ID = len(productList) + 1 // Assign a new ID based on the current length of the product list
	productList = append(productList, newProduct)

	encoder := json.NewEncoder(w)
	encoder.Encode(newProduct) // Return the newly created product as a response
}

func main () {
	mux := http.NewServeMux() //router

	mux.HandleFunc("/hello", helloHandler) //route
	mux.HandleFunc("/about", aboutHandler) //route
    mux.HandleFunc("/getproducts", getProductsHandler) //route
	mux.HandleFunc("/createproduct", createProductHandler) //route

	fmt.Println("Server is running on http://localhost:8080")

	err := http.ListenAndServe(":8080", mux) //Failed to start server

	if err != nil {
		fmt.Printf("Error starting server: %v", err)
	}
}

func init() {
	prd1 := Product{
		ID: 1,
		Title: "Product 1",
		Price: 10.99,
		Description: "This is product 1",
		ImgUrl: "https://example.com/product1.jpg",
	}
	prd2 := Product{
		ID: 2,
		Title: "Product 2",
		Price: 20.99,
		Description: "This is product 2",
		ImgUrl: "https://example.com/product2.jpg",
	}
	prd3 := Product{
		ID: 3,
		Title: "Product 3",
		Price: 30.99,
		Description: "This is product 3",
		ImgUrl: "https://example.com/product3.jpg",
	}
	prd4 := Product{
		ID: 4,
		Title: "Product 4",
		Price: 40.99,
		Description: "This is product 4",
		ImgUrl: "https://example.com/product4.jpg",
	}

	productList = append(productList, prd1, prd2, prd3, prd4)
}