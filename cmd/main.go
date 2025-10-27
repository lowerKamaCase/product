package main

import (
	"fmt"
	"net/http"

	"github.com/lowerKamaCase/product/configs"
	"github.com/lowerKamaCase/product/pkg/db"
	"github.com/lowerKamaCase/product/pkg/product"
)

const PORT = 8081

func main() {
	config := configs.LoadConfig()
	database := db.NewDb(config)
	mux := http.NewServeMux()

	productRepository := product.NewProductRepository(database)

	product.NewProductHandler(mux, product.ProductHandlerDeps{
		ProductRepository: productRepository,
	})

	Addr := fmt.Sprintf(":%d", PORT)
	server := http.Server{
		Addr:    Addr,
		Handler: mux,
	}

	fmt.Println("Server started at port: ", PORT)

	err := server.ListenAndServe()

	if err != nil {
		panic(err.Error())
	}
}
