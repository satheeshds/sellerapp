package models

type Order struct {
	Id             int    `json:"id,omitempty"`
	User_id        int    `json:"user_id,omitempty"`
	Transaction_id int    `json:"transaction_id,omitempty"`
	Status         string `json:"status,omitempty"`
	Created_at     string `json:"created_at,omitempty"`
}

func (o *Order) CancelOrder() {
	o.Status = "cancelled"
}
