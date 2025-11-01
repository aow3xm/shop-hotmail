package model

type GetStockResponse struct {
	Sum int `json:"sum"`
}

type ApiResponse[T any] struct {
	Success bool    `json:"success"`
	Message *string `json:"message"`
	Data    T       `json:"data"`
}

type ApiGetStock struct {
	Total   int `json:"total"`
	Outlook int `json:"outlook"`
	Hotmail int `json:"hotmail"`
}

type ApiPurchase struct {
	Username     string `json:"username"`
	Login        string `json:"login"`
	Password     string `json:"password"`
	RefreshToken string `json:"refresh_token"`
	ServiceID    string `json:"service_id"`
}

type GetStockRequest struct {
	KioskID string `query:"key"`
}

type PurchaseRequest struct {
	KioskID  string `query:"key"`
	OrderID  string `query:"order_id"`
	Quantity string `query:"quantity"`
}

type ProductItem struct {
	Product string `json:"product"`
}

type PurchaseResponse []ProductItem
