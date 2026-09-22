package main

import (
	"github.com/habibmrizki/week10/internal/router"
)

func main() {
	r := router.InitRouter()

	r.Run(":3000")

}
