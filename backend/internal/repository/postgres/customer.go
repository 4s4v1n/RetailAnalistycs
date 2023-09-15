package postgres

import (
	"APG6/internal/controller/auth"
	"APG6/internal/entity"
	"context"
	"fmt"

	"github.com/doug-martin/goqu/v9"
)

const (
	customersView = `customers`
)

func (r *Repository) GetCustomers(ctx context.Context, role uint8) ([]entity.Customer, error) {
	selectQuery, _, err := goqu.From(customersView).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("configure query: %w", err)
	}

	var customers []entity.Customer
	switch role {
	case auth.RoleVisitor:
		err = r.visitor.DB.SelectContext(ctx, &customers, selectQuery)
	case auth.RoleAdmin:
		err = r.admin.DB.SelectContext(ctx, &customers, selectQuery)
	default:
		return nil, fmt.Errorf("select data: cannot exec with unknown role")
	}
	if err != nil {
		return nil, fmt.Errorf("select data: %w", err)
	}
	return customers, nil
}
