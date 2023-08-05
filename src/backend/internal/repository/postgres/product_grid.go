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
	productGridTable = `product_grid`
)

func (r *Repository) AddProductGrid(ctx context.Context, role uint8, item entity.ProductGrid) error {
	insertQuery, _, err := goqu.Insert(productGridTable).Rows(item).ToSQL()
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

func (r *Repository) GetProductGrid(ctx context.Context, role uint8) ([]entity.ProductGrid, error) {
	selectQuery, _, err := goqu.From(productGridTable).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("configure query: %w", err)
	}

	var grids []entity.ProductGrid
	switch role {
	case auth.RoleVisitor:
		err = r.visitor.DB.SelectContext(ctx, &grids, selectQuery)
	case auth.RoleAdmin:
		err = r.admin.DB.SelectContext(ctx, &grids, selectQuery)
	default:
		return nil, errors.New("select data: cannot exec with unknown role")
	}
	return grids, nil
}

func (r *Repository) UpdateProductGrid(ctx context.Context, role uint8, item entity.ProductGrid) error {
	updateQuery, _, err := goqu.Update(productGridTable).Set(goqu.Record{
		"sku_name": item.SkuName,
		"group_id": item.GroupId,
	}).Where(goqu.C("sku_id").Eq(item.SkuId)).Returning("sku_id").ToSQL()
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

func (r *Repository) DeleteProductGrid(ctx context.Context, role uint8, key string) error {
	deleteQuery, _, err := goqu.Delete(productGridTable).
		Where(goqu.C("sku_id").Eq(key)).
		Returning("sku_id").ToSQL()
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
