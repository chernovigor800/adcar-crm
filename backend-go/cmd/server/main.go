package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"adcar-crm/backend-go/internal/rest"
	"adcar-crm/backend-go/internal/repository"
	"adcar-crm/backend-go/internal/service"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	pool, err := pgxpool.New(context.Background(), "user=... host=localhost dbname=... sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer pool.Close()

	repo := repository.NewCarRepository(pool)
	svc := service.NewCarService(repo)

	r := gin.Default()
	rest.SetupHTTPRouter(r, svc)
	r.Run(":8080")
}
