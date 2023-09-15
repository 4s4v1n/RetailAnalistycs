package entity

import "time"

type DefiningOfferIncreasingFrequencyVisits struct {
	CustomerId                int32     `json:"customer_id" csv:"customer_id" db:"customer_id"`
	StartDate                 time.Time `json:"start_date,time.RFC3339" csv:"start_date" db:"start_date"`
	EndDate                   time.Time `json:"end_date,time.RFC3339" csv:"end_date" db:"end_date"`
	RequiredTransactionsCount float64   `json:"required_transactions_count" csv:"required_transactions_count" db:"required_transactions_count"`
	GroupName                 string    `json:"group_name" csv:"group_name" db:"group_name"`
	OfferDiscountDepth        float64   `json:"offer_discount_depth" csv:"offer_discount_depth" db:"offer_discount_depth"`
}
