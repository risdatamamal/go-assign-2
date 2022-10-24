package item

type Item struct {
	ItemID      uint64 `json:"id" gorm:"column:id;primaryKey"`
	ItemCode    string `json:"itemCode" gorm:"column:itemCode"`
	Description string `json:"description" gorm:"column:description"`
	Quantity    string `json:"quantity" gorm:"column:quantity"`
}
