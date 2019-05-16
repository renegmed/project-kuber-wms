package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	//"strconv"
	//"fmt"
	"github.com/gin-gonic/gin"
)

type Product struct {
	Sku      string
	Name     string
	Quantity int
}

func main() {

	engine := gin.Default()

	engine.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	engine.LoadHTMLGlob("./templates/*.html")

	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Warehousing System",
		})
	})

	engine.GET("/products", func(c *gin.Context) {

		prods, err := readProducts("http://127.0.0.1:8080/api/products")

		if err != nil {
			c.String(http.StatusBadRequest, "URL data error")
			return
		}
		var products []Product

		//fmt.Printf("Data received:\n %s ", prods)

		err = json.Unmarshal(prods, &products)
		if err != nil {
			c.String(http.StatusInternalServerError, "Products unmarshal error")
			return
		}

		fmt.Printf("Products received:\n %v ", products)

		params := map[string]interface{}{
			"products": products,
		}

		c.HTML(http.StatusOK, "index.html", params)

	})

	engine.Run(port())
}

func readProducts(url string) ([]byte, error) {

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8484"
	}
	return ":" + port
}
