package controller_order

type CreateOrderPayload struct {
	ProductIDs []int64 `json:"product_ids"`
}