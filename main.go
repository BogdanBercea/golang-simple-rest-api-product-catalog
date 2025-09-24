package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Product represents the structure of each product
type Product struct {
	ID             int      `json:"id"`
	Name           string   `json:"name"`
	AvailableSizes []string `json:"available_sizes"`
	Images         []string `json:"images"`
	Brand          string   `json:"brand"`
}

func main() {
	router := gin.Default()

	router.GET("/products", getProducts)
	router.GET("/products/:id", getProduct)
	router.POST("/products", createProduct)

	router.Run("localhost:9090")
}

// getProductsFromFile reads the JSON file and returns a slice of Product
func getProductsFromFile(filename string) ([]Product, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var products []Product
	if err := json.Unmarshal(bytes, &products); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return products, nil
}

func getProducts(context *gin.Context) {
	products, err := getProductsFromFile("products.json")

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Failed retrieving the products catalog"})
	} else {
		brandFileter := context.Query("brand")

		if brandFileter != "" {
			var filtered []Product

			for _, product := range products {
				if product.Brand == brandFileter {
					filtered = append(filtered, product)
				}
			}

			context.IndentedJSON(http.StatusOK, filtered)
		} else {
			context.IndentedJSON(http.StatusOK, products)
		}
	}
}

func createProduct(context *gin.Context) {
	var newProduct Product

	if err := context.BindJSON(&newProduct); err != nil {
		return
	}

	products, err := getProductsFromFile("products.json")
	if err != nil {
		return
	}

	products = append(products, newProduct)

	bytes, err := json.MarshalIndent(products, "", "  ")
	if err != nil {
		return
	}

	if err := ioutil.WriteFile("products.json", bytes, 0644); err != nil {
		return
	}

	context.IndentedJSON(http.StatusCreated, newProduct)
}

func getProductById(id int) (*Product, error) {
	products, err := getProductsFromFile("products.json")

	if err != nil {
		return nil, errors.New("Failed getting the products.")
	}

	for _, product := range products {
		if product.ID == id {
			return &product, nil
		}
	}

	return nil, fmt.Errorf("product with ID %d not found", id)
}

func getProduct(context *gin.Context) {
	stringId := context.Param("id")
	id, err := strconv.Atoi(stringId)

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "The provided id must be an integer."})
	}

	product, err := getProductById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Product not found."})
	}

	context.IndentedJSON(http.StatusOK, product)
}
