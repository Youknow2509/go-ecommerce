package main

import (
    "github.com/Bot-SomeOne/go-ecommerce/internal/routers"
)

func main() {
	r := routers.NewRouter()

    r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
