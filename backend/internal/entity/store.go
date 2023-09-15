package entity

type Store struct {
	TransactionStoreId int64   `json:"transaction_store_id" csv:"transaction_store_id" db:"transaction_store_id"`
	SkuId              *int64  `json:"sku_id" csv:"sku_id" db:"sku_id"`
	SkuPurchasePrice   float64 `json:"sku_purchase_price" csv:"sku_purchase_price" db:"sku_purchase_price"`
	SkuRetailPrice     float64 `json:"sku_retail_price" csv:"sku_retail_price" db:"sku_retail_price"`
}
