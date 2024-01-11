package domain

type Buy struct {
	ID          string  `json:"id"`
	BuyID       string  `json:"buyId"`
	User        string  `json:"user"`
	ProductName string  `json:"productName"`
	Price       float64 `json:"price"`
	Points      int     `json:"points"`
}

type Response struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}
