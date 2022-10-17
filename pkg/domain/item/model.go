package item

type Item struct {
	ID          uint64 `json:"id" gorm:"column:id;primaryKey"`
	Code        string `json:"code" gorm:"column:code"`
	Description string `json:"description" gorm:"column:description"`
	Quantity    string `json:"quantity" gorm:"column:quantity"`
	OrderID     string `json:"order_id" gorm:"column:order_id"`
}
