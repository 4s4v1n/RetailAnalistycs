package entity

type DefiningOfferIncreasingMargin struct {
	CustomerId         int32   `json:"customer_id" csv:"customer_id" db:"customer_id"`
	SkuName            string  `json:"sku_name" csv:"sku_name" db:"sku_name"`
	OfferDiscountDepth float64 `json:"offer_discount_depth" csv:"offer_discount_depth" db:"offer_discount_depth"`
}
