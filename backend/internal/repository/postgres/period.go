package postgres

import (
	"APG6/internal/controller/auth"
	"APG6/internal/entity"
	"context"
	"fmt"

	"github.com/doug-martin/goqu/v9"
)

const (
	periodsView = `periods`
)

func (r *Repository) GetPeriods(ctx context.Context, role uint8) ([]entity.Period, error) {
	selectQuery, _, err := goqu.From(periodsView).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("configure query: %w", err)
	}

	var periods []entity.Period
	switch role {
	case auth.RoleVisitor:
		err = r.visitor.DB.SelectContext(ctx, &periods, selectQuery)
	case auth.RoleAdmin:
		err = r.admin.DB.SelectContext(ctx, &periods, selectQuery)
	default:
		return nil, fmt.Errorf("select data: cannot exec with unknown role")
	}
	if err != nil {
		return nil, fmt.Errorf("select data: %w", err)
	}
	return periods, nil
}
