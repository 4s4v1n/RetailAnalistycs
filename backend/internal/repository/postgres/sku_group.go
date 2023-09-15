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
	skuGroupTable = `sku_group`
)

func (r *Repository) AddSkuGroup(ctx context.Context, role uint8, item entity.SkuGroup) error {
	insertQuery, _, err := goqu.Insert(skuGroupTable).Rows(item).ToSQL()
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

func (r *Repository) GetSkuGroup(ctx context.Context, role uint8) ([]entity.SkuGroup, error) {
	selectQuery, _, err := goqu.From(skuGroupTable).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("configure query: %w", err)
	}

	var groups []entity.SkuGroup
	switch role {
	case auth.RoleVisitor:
		err = r.visitor.DB.SelectContext(ctx, &groups, selectQuery)
	case auth.RoleAdmin:
		err = r.admin.DB.SelectContext(ctx, &groups, selectQuery)
	default:
		return nil, errors.New("select data: cannot exec with unknown role")
	}
	return groups, nil
}

func (r *Repository) UpdateSkuGroup(ctx context.Context, role uint8, item entity.SkuGroup) error {
	updateQuery, _, err := goqu.Update(skuGroupTable).Set(goqu.Record{
		"group_name": item.GroupName,
	}).Where(goqu.C("group_id").Eq(item.GroupId)).Returning("group_id").ToSQL()
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

func (r *Repository) DeleteSkuGroup(ctx context.Context, role uint8, key string) error {
	deleteQuery, _, err := goqu.Delete(skuGroupTable).
		Where(goqu.C("group_id").Eq(key)).
		Returning("group_id").ToSQL()
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
