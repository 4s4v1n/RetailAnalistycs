package entity

type Check struct {
	TransactionId *int64  `json:"transaction_id" csv:"transaction_id" db:"transaction_id"`
	SkuId         *int64  `json:"sku_id" csv:"sku_id" db:"sku_id"`
	SkuAmount     float64 `json:"sku_amount" csv:"sku_amount" db:"sku_amount"`
	SkuSum        float64 `json:"sku_sum" csv:"sku_sum" db:"sku_sum"`
	SkuSumPaid    float64 `json:"sku_sum_paid" csv:"sku_sum_paid" db:"sku_sum_paid"`
	SkuDiscount   float64 `json:"sku_discount" csv:"sku_discount" db:"sku_discount"`
}
