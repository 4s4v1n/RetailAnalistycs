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
	storeTable = `stores`
)

func (r *Repository) AddStore(ctx context.Context, role uint8, item entity.Store) error {
	insertQuery, _, err := goqu.Insert(storeTable).Rows(item).ToSQL()
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

func (r *Repository) GetStore(ctx context.Context, role uint8) ([]entity.Store, error) {
	selectQuery, _, err := goqu.From(storeTable).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("configure query: %w", err)
	}

	var stores []entity.Store
	switch role {
	case auth.RoleVisitor:
		err = r.visitor.DB.SelectContext(ctx, &stores, selectQuery)
	case auth.RoleAdmin:
		err = r.admin.DB.SelectContext(ctx, &stores, selectQuery)
	default:
		return nil, errors.New("select data: cannot exec with unknown role")
	}
	return stores, nil
}

func (r *Repository) UpdateStore(ctx context.Context, role uint8, item entity.Store) error {
	updateQuery, _, err := goqu.Update(storeTable).Set(goqu.Record{
		"sku_purchase_price": item.SkuPurchasePrice,
		"sku_retail_price":   item.SkuRetailPrice,
	}).Where(goqu.C("transaction_store_id").Eq(item.TransactionStoreId),
		goqu.C("sku_id").Eq(item.SkuId)).
		Returning("transaction_store_id").ToSQL()
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

func (r *Repository) DeleteStore(ctx context.Context, role uint8, key1 string, key2 string) error {
	deleteQuery, _, err := goqu.Delete(storeTable).
		Where(goqu.C("transaction_store_id").Eq(key1),
			goqu.C("sku_id").Eq(key2)).
		Returning("transaction_store_id").ToSQL()
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
