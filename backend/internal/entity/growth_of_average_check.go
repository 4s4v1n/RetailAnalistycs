package entity

type GrowthOfAverageCheck struct {
	CustomerId           int32   `json:"customer_id" csv:"customer_id" db:"customer_id"`
	RequiredCheckMeasure float64 `json:"required_check_measure" csv:"required_check_measure" db:"required_check_measure"`
	GroupName            string  `json:"group_name" csv:"group_name" db:"group_name"`
	OfferDiscountDepth   float64 `json:"offer_discount_depth" csv:"offer_discount_depth" db:"offer_discount_depth"`
}
