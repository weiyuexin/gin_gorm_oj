package main

import "gin_gorm_oj/router"

func main() {
	r := router.Router()

	r.Run(":9090")
}
