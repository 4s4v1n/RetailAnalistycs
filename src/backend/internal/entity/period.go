package entity

import "time"

type Period struct {
	CustomerId         int64     `json:"customer_id" csv:"customer_id" db:"customer_id"`
	GroupId            int64     `json:"group_id" csv:"group_id" db:"group_id"`
	FirstGroupPurchase time.Time `json:"first_group_purchase_date" csv:"first_group_purchase_date" db:"first_group_purchase_date"`
	LastGroupPurchase  time.Time `json:"last_group_purchase_date" csv:"last_group_purchase_date" db:"last_group_purchase_date"`
	GroupPurchase      int64     `json:"group_purchase" csv:"group_purchase" db:"group_purchase"`
	GroupFrequency     float64   `json:"group_frequency" csv:"group_frequency" db:"group_frequency"`
	GroupMinDiscount   float64   `json:"group_min_discount" csv:"group_min_discount" db:"group_min_discount"`
}
