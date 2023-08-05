package postgres

import (
	"APG6/internal/controller/auth"
	"APG6/internal/entity"
	"context"
	"errors"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
)

const (
	transactionsTable = `transactions`
)

func (r *Repository) AddTransaction(ctx context.Context, role uint8, item entity.Transaction) error {
	insertQuery, _, err := goqu.Insert(transactionsTable).Rows(item).ToSQL()
	if err != nil {
		return fmt.Errorf("configure query: %w", err)
	}

	switch role {
	case auth.RoleVisitor:
		_, err = r.visitor.DB.ExecContext(ctx, insertQuery)
	case auth.RoleAdmin:
		_, err = r.admin.DB.ExecContext(ctx, insertQuery)
	default:
		return errors.New("insert data: cannot exec with unknown role")
	}
	if err != nil {
		return fmt.Errorf("insert data: %w", err)
	}
	return nil
}

func (r *Repository) GetTransaction(ctx context.Context, role uint8) ([]entity.Transaction, error) {
	selectQuery, _, err := goqu.From(transactionsTable).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("configure query: %w", err)
	}

	var transactions []entity.Transaction
	switch role {
	case auth.RoleVisitor:
		err = r.visitor.DB.SelectContext(ctx, &transactions, selectQuery)
	case auth.RoleAdmin:
		err = r.admin.DB.SelectContext(ctx, &transactions, selectQuery)
	default:
		return nil, errors.New("select data: cannot exec with unknown role")
	}
	return transactions, nil
}

func (r *Repository) UpdateTransaction(ctx context.Context, role uint8, item entity.Transaction) error {
	updateQuery, _, err := goqu.Update(transactionsTable).Set(goqu.Record{
		"customer_card_id":     item.CustomerCardId,
		"transaction_sum":      item.TransactionSum,
		"transaction_datetime": item.TransactionDatetime,
		"transaction_store_id": item.TransactionStoreId,
	}).Where(goqu.C("transaction_id").Eq(item.TransactionId)).Returning("transaction_id").ToSQL()
	if err != nil {
		return fmt.Errorf("configure query: %w", err)
	}

	var id string
	var row *sqlx.Row
	switch role {
	case auth.RoleVisitor:
		row = r.visitor.DB.QueryRowxContext(ctx, updateQuery)
	case auth.RoleAdmin:
		row = r.admin.DB.QueryRowxContext(ctx, updateQuery)
	default:
		return errors.New("update data: cannot exec with unknown role")
	}
	if err = row.Scan(&id); err != nil {
		return fmt.Errorf("update data: %w", err)
	}
	return nil
}

func (r *Repository) DeleteTransaction(ctx context.Context, role uint8, key string) error {
	deleteQuery, _, err := goqu.Delete(transactionsTable).
		Where(goqu.C("transaction_id").Eq(key)).
		Returning("transaction_id").ToSQL()
	if err != nil {
		return fmt.Errorf("configure query: %w", err)
	}

	var id string
	var row *sqlx.Row
	switch role {
	case auth.RoleVisitor:
		row = r.visitor.DB.QueryRowxContext(ctx, deleteQuery)
	case auth.RoleAdmin:
		row = r.admin.DB.QueryRowxContext(ctx, deleteQuery)
	default:
		return errors.New("update data: cannot exec with unknown role")
	}
	if err = row.Scan(&id); err != nil {
		return fmt.Errorf("delete data: %w", err)
	}
	return nil
}
