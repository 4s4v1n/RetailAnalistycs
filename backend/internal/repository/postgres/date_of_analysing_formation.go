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
	dateOfAnalysingFormationTable = `date_of_analysing_formation`
)

func (r *Repository) GetDateOfAnalysingFormation(ctx context.Context,
	role uint8) ([]entity.DateOfAnalysingFormation, error) {
	selectQuery, _, err := goqu.From(dateOfAnalysingFormationTable).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("configure query: %w", err)
	}

	var dates []entity.DateOfAnalysingFormation
	switch role {
	case auth.RoleVisitor:
		err = r.visitor.DB.SelectContext(ctx, &dates, selectQuery)
	case auth.RoleAdmin:
		err = r.admin.DB.SelectContext(ctx, &dates, selectQuery)
	default:
		return nil, errors.New("select data: cannot exec with unknown role")
	}
	return dates, nil
}

func (r *Repository) UpdateDateOfAnalysingFormation(ctx context.Context, role uint8,
	item entity.UpdateFormation) error {
	updateQuery, _, err := goqu.Update(dateOfAnalysingFormationTable).Set(goqu.Record{
		"date": item.NewDate,
	}).Where(goqu.C("date").Eq(item.OldDate)).Returning("date").ToSQL()
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
