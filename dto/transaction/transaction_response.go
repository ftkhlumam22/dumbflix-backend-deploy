package transactionsdto

import "time"

type TransactionResponse struct {
	ID        int       `json:"id"`
	StartDate time.Time `json:"statusdate"`
	DueDate   time.Time `json:"duedate"`
	Status    string    `json:"status"`
	Price     int       `json:"price"`
	UserID    int       `json:"user_id"`
}
