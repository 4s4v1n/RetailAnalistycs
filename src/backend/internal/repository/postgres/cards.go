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
	cardsTable = `cards`
)

func (r *Repository) AddCard(ctx context.Context, role uint8, item entity.Card) error {
	insertQuery, _, err := goqu.Insert(cardsTable).Rows(item).ToSQL()
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

func (r *Repository) GetCard(ctx context.Context, role uint8) ([]entity.Card, error) {
	selectQuery, _, err := goqu.From(cardsTable).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("configure query: %w", err)
	}

	var cards []entity.Card
	switch role {
	case auth.RoleVisitor:
		err = r.visitor.DB.SelectContext(ctx, &cards, selectQuery)
	case auth.RoleAdmin:
		err = r.admin.DB.SelectContext(ctx, &cards, selectQuery)
	default:
		return nil, errors.New("select data: cannot exec with unknown role")
	}
	return cards, nil
}

func (r *Repository) UpdateCard(ctx context.Context, role uint8, item entity.Card) error {
	updateQuery, _, err := goqu.Update(cardsTable).Set(goqu.Record{
		"customer_id": item.CustomerId,
	}).Where(goqu.C("customer_card_id").Eq(item.CustomerCardId)).Returning("customer_card_id").ToSQL()
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

func (r *Repository) DeleteCard(ctx context.Context, role uint8, key string) error {
	deleteQuery, _, err := goqu.Delete(cardsTable).
		Where(goqu.C("customer_card_id").Eq(key)).
		Returning("customer_card_id").ToSQL()
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
