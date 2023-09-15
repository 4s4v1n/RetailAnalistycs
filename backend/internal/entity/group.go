package entity

type Group struct {
	CustomerId           int64   `json:"customer_id" csv:"customer_id" db:"customer_id"`
	GroupId              int64   `json:"group_id" csv:"group_id" db:"group_id"`
	GroupAffinityIndex   float64 `json:"group_affinity_index" csv:"group_affinity_index" db:"group_affinity_index"`
	GroupChurnRate       float64 `json:"group_churn_rate" csv:"group_churn_rate" db:"group_churn_rate"`
	GroupStabilityIndex  float64 `json:"group_stability_index" csv:"group_stability_index" db:"group_stability_index"`
	GroupMargin          float64 `json:"group_margin" csv:"group_margin" db:"group_margin"`
	GroupDiscountShare   float64 `json:"group_discount_share" csv:"group_discount_share" db:"group_discount_share"`
	GroupMinimumDiscount float64 `json:"group_minimum_discount" csv:"group_minimum_discount" db:"group_minimum_discount"`
	GroupAverageDiscount float64 `json:"group_average_discount" csv:"group_average_discount" db:"group_average_discount"`
}
