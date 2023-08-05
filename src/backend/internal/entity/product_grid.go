package entity

type ProductGrid struct {
	SkuId   int64  `json:"sku_id,omitempty" csv:"sku_id" db:"sku_id" goqu:"skipinsert"`
	SkuName string `json:"sku_name" csv:"sku_name" db:"sku_name"`
	GroupId *int64 `json:"group_id" csv:"group_id" db:"group_id"`
}
