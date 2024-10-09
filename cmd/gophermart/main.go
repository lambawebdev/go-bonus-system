package main

import (
	"database/sql"
	"embed"
	"net/http"

	"github.com/go-chi/chi"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/lambawebdev/go-bonus-system/internal/gophemart/config"
	"github.com/lambawebdev/go-bonus-system/internal/gophemart/handlers"
	"github.com/lambawebdev/go-bonus-system/internal/gophemart/middleware"
	"github.com/lambawebdev/go-bonus-system/internal/gophemart/repositories"
	orderservice "github.com/lambawebdev/go-bonus-system/internal/gophemart/services/orderService"
	"github.com/pressly/goose/v3"
)

//go:embed database/migrations/*.sql
var embedMigrations embed.FS

func main() {
	config.ParseFlags()

	db, err := ConnectToDB(config.GetDsn())

	if err != nil {
		panic(err)
	}
	defer db.Close()

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.Up(db, "database/migrations"); err != nil {
		panic(err)
	}

	err = RunServer(db)

	if err != nil {
		panic(err)
	}
}

func ConnectToDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func RunServer(db *sql.DB) error {
	r := chi.NewRouter()

	userRepo := repositories.NewUserRepository(db)
	orderRepo := repositories.NewOrderRepository(db)
	transRepo := repositories.NewTransactionRepository(db)
	withdrawalRepo := repositories.NewWithdrawalRepository(db)

	orderService := orderservice.NewOrderService(orderRepo, transRepo)

	regHandler := handlers.NewRegistrationHandler(userRepo)
	authHandler := handlers.NewAuthenticationHandler(userRepo)
	orderHandler := handlers.NewOrderHandler(orderRepo)
	balanceHandler := handlers.NewBalanceHandler(orderRepo, transRepo, withdrawalRepo)

	r.Post("/api/user/register", regHandler.Register)
	r.Post("/api/user/login", authHandler.Authenticate)

	r.Get("/api/user/orders", middleware.AuthMiddleware(orderHandler.GetOrders))
	r.Post("/api/user/orders", middleware.AuthMiddleware(orderHandler.CreateOrder))

	r.Get("/api/user/withdrawals", middleware.AuthMiddleware(balanceHandler.GetWithdrawals))
	r.Post("/api/user/balance/withdraw", middleware.AuthMiddleware(balanceHandler.Withdraw))
	r.Get("/api/user/balance", middleware.AuthMiddleware(balanceHandler.GetBalance))

	go orderService.RunUpdateOrdersStatuses()

	err := http.ListenAndServe(config.GetHost(), r)

	return err
}
