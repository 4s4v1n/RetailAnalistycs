package entity

type Card struct {
	CustomerCardId int64  `json:"customer_card_id,omitempty" csv:"customer_card_id" db:"customer_card_id" goqu:"skipinsert"`
	CustomerId     *int64 `json:"customer_id" csv:"customer_id" db:"customer_id"`
}
