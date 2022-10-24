package order

import model "github.com/risdatamamal/go-assign-2/pkg/domain/item"

type Order struct {
	OrderID      uint64       `json:"id" gorm:"column:id;primaryKey"`
	CustomerName string       `json:"customerName" gorm:"column:customerName"`
	OrderedAt    string       `json:"orderedAt" gorm:"column:orderedAt"`
	Items        []model.Item `json:"items" gorm:"column:items"`
}
