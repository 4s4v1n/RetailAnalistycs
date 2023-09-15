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
	checksTable = `checks`
)

func (r *Repository) AddCheck(ctx context.Context, role uint8, item entity.Check) error {
	insertQuery, _, err := goqu.Insert(checksTable).Rows(item).ToSQL()
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

func (r *Repository) GetCheck(ctx context.Context, role uint8) ([]entity.Check, error) {
	selectQuery, _, err := goqu.From(checksTable).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("configure query: %w", err)
	}

	var checks []entity.Check
	switch role {
	case auth.RoleVisitor:
		err = r.visitor.DB.SelectContext(ctx, &checks, selectQuery)
	case auth.RoleAdmin:
		err = r.admin.DB.SelectContext(ctx, &checks, selectQuery)
	default:
		return nil, errors.New("select data: cannot exec with unknown role")
	}
	return checks, nil
}

func (r *Repository) UpdateCheck(ctx context.Context, role uint8, item entity.Check) error {
	updateQuery, _, err := goqu.Update(checksTable).Set(goqu.Record{
		"sku_amount":   item.SkuAmount,
		"sku_sum":      item.SkuSum,
		"sku_sum_paid": item.SkuSumPaid,
		"sku_discount": item.SkuDiscount,
	}).Where(goqu.C("transaction_id").Eq(item.TransactionId),
		goqu.C("sku_id").Eq(item.SkuId)).Returning("transaction_id").ToSQL()
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

func (r *Repository) DeleteCheck(ctx context.Context, role uint8, key1 string, key2 string) error {
	deleteQuery, _, err := goqu.Delete(checksTable).
		Where(goqu.C("transaction_id").Eq(key1),
			goqu.C("sku_id").Eq(key2)).
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
