package entity

import "time"

type PurchaseHistory struct {
	CustomerId          int64     `json:"customer_id" csv:"customer_id" db:"customer_id"`
	TransactionId       int32     `json:"transaction_id" csv:"transaction_id" db:"transaction_id"`
	TransactionDatetime time.Time `json:"transaction_datetime" csv:"transaction_datetime" db:"transaction_datetime"`
	GroupId             int64     `json:"group_id" csv:"group_id" db:"group_id"`
	GroupCost           float64   `json:"group_cost" csv:"group_cost" db:"group_cost"`
	GroupSum            float64   `json:"group_sum" csv:"group_sum" db:"group_sum"`
	GroupSumPaid        float64   `json:"group_sum_paid" csv:"group_sum_paid" db:"group_sum_paid"`
}
