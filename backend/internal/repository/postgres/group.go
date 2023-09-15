package postgres

import (
	"APG6/internal/controller/auth"
	"APG6/internal/entity"
	"context"
	"fmt"

	"github.com/doug-martin/goqu/v9"
)

const (
	groupsView = `groups`
)

func (r *Repository) GetGroups(ctx context.Context, role uint8) ([]entity.Group, error) {
	selectQuery, _, err := goqu.From(groupsView).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("configure query: %w", err)
	}

	var groups []entity.Group
	switch role {
	case auth.RoleVisitor:
		err = r.visitor.DB.SelectContext(ctx, &groups, selectQuery)
	case auth.RoleAdmin:
		err = r.admin.DB.SelectContext(ctx, &groups, selectQuery)
	default:
		return nil, fmt.Errorf("select data: cannot exec with unknown role")
	}
	if err != nil {
		return nil, fmt.Errorf("select data: %w", err)
	}
	return groups, nil
}
