package postgres

import (
	"APG6/internal/controller/auth"
	"APG6/internal/entity"
	"context"
	"fmt"

	"github.com/doug-martin/goqu/v9"
)

const (
	purchaseHistoryView = `purchase_history`
)

func (r *Repository) GetPurchaseHistory(ctx context.Context, role uint8) ([]entity.PurchaseHistory, error) {
	selectQuery, _, err := goqu.From(purchaseHistoryView).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("configure query: %w", err)
	}

	var history []entity.PurchaseHistory
	switch role {
	case auth.RoleVisitor:
		err = r.visitor.DB.SelectContext(ctx, &history, selectQuery)
	case auth.RoleAdmin:
		err = r.admin.DB.SelectContext(ctx, &history, selectQuery)
	default:
		return nil, fmt.Errorf("select data: cannot exec with unknown role")
	}
	if err != nil {
		return nil, fmt.Errorf("select data: %w", err)
	}
	return history, nil
}
