package entity

import (
	"APG6/internal/entity/utils"
	"encoding/json"
	"time"
)

type Transaction struct {
	TransactionId       int64     `json:"transaction_id,omitempty" csv:"transaction_id" db:"transaction_id" goqu:"skipinsert"`
	CustomerCardId      *int64    `json:"customer_card_id" csv:"customer_card_id" db:"customer_card_id"`
	TransactionSum      float64   `json:"transaction_sum" csv:"transaction_sum" db:"transaction_sum"`
	TransactionDatetime time.Time `json:"transaction_datetime,time.RFC3339" csv:"transaction_datetime" db:"transaction_datetime"`
	TransactionStoreId  int64     `json:"transaction_store_id" csv:"transaction_store_id" db:"transaction_store_id"`
}

func (tr *Transaction) UnmarshalJSON(b []byte) error {
	type transactionAlias Transaction
	alias := &struct {
		*transactionAlias
		TransactionDatetime string `json:"transaction_datetime"`
	}{
		transactionAlias: (*transactionAlias)(tr),
	}

	if err := json.Unmarshal(b, &alias); err != nil {
		return err
	}
	t, err := utils.ParseDatetime(alias.TransactionDatetime)
	if err != nil {
		return err
	}

	tr.TransactionDatetime = t
	return nil
}
