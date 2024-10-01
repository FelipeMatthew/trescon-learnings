package services

import (
	"fmt"
	"total_count/internal/api/routes"
)

func PostImage() {
	countImageRoute := routes.NewApiRoutes().CountItens

	fmt.Println(countImageRoute)
}
