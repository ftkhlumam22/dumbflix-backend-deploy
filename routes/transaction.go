package routes

import (
	"testdumpflix/handlers"
	"testdumpflix/pkg/middleware"
	"testdumpflix/pkg/mysql"
	"testdumpflix/repositories"

	"github.com/gorilla/mux"
)

func TransactionRoutes(r *mux.Router) {
	transactionRepository := repositories.RepositoryforTransaction(mysql.DB)
	h := handlers.HandlerTransaction(transactionRepository)

	r.HandleFunc("/transactions", h.FindTransactions).Methods("GET")
	r.HandleFunc("/transaction/{id}", h.GetTransaction).Methods("GET")
	r.HandleFunc("/transaction", middleware.Auth(h.CreateTransaction)).Methods("POST")
	r.HandleFunc("/notification", h.Notification).Methods("POST")
	r.HandleFunc("/transaction/{id}", middleware.Auth(middleware.IsAdmin(h.DeleteTransaction))).Methods("DELETE")

}
