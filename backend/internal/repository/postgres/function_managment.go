package postgres

import (
	"APG6/internal/controller/auth"
	"APG6/internal/entity"
	"context"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"time"
)

const (
	fncGrowthOfAverageCheck                   = `fnc_growth_of_average_check`
	fncDefiningOfferIncreasingFrequencyVisits = `fnc_defining_offer_increasing_frequency_visits`
	fncDefiningOfferIncreasingMargin          = `fnc_defining_offer_increasing_margin`
)

func (r *Repository) GrowthOfAverageCheck(ctx context.Context, role uint8, method string, first time.Time,
	last time.Time, number int32, coefficient float64, maxChurnRate float64, maxDiscountShare float64,
	marginShare float64) ([]entity.GrowthOfAverageCheck, error) {
	selectQuery, _, err := goqu.From(goqu.Func(fncGrowthOfAverageCheck, method, first, last, number, coefficient,
		maxChurnRate, maxDiscountShare, marginShare)).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("configure query: %w", err)
	}

	var growth []entity.GrowthOfAverageCheck
	switch role {
	case auth.RoleVisitor:
		err = r.visitor.DB.SelectContext(ctx, &growth, selectQuery)
	case auth.RoleAdmin:
		err = r.admin.DB.SelectContext(ctx, &growth, selectQuery)
	default:
		return nil, fmt.Errorf("select data: cannot exec with unknown role")
	}
	if err != nil {
		return nil, fmt.Errorf("select data: %w", err)
	}
	return growth, nil
}

func (r *Repository) DefiningOfferIncreasingFrequencyVisits(ctx context.Context, role uint8, first time.Time,
	last time.Time, valueTransaction int32, maxChurnRate float64, maxDiscountShare float64,
	marginShare float64) ([]entity.DefiningOfferIncreasingFrequencyVisits, error) {
	selectQuery, _, err := goqu.From(goqu.Func(fncDefiningOfferIncreasingFrequencyVisits, first, last,
		valueTransaction, maxChurnRate, maxDiscountShare, marginShare)).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("configure query: %w", err)
	}

	var increase []entity.DefiningOfferIncreasingFrequencyVisits
	switch role {
	case auth.RoleVisitor:
		err = r.visitor.DB.SelectContext(ctx, &increase, selectQuery)
	case auth.RoleAdmin:
		err = r.admin.DB.SelectContext(ctx, &increase, selectQuery)
	default:
		return nil, fmt.Errorf("select data: cannot exec with unknown role")
	}
	if err != nil {
		return nil, fmt.Errorf("select data: %w", err)
	}
	return increase, nil
}

func (r *Repository) DefiningOfferIncreasingMargin(ctx context.Context, role uint8, countGroup int32,
	maxChurnRate float64, maxStabilityIndex float64, maxIndexSku float64,
	marginShare float64) ([]entity.DefiningOfferIncreasingMargin, error) {
	selectQuery, _, err := goqu.From(goqu.Func(fncDefiningOfferIncreasingMargin, countGroup, maxChurnRate,
		maxStabilityIndex, maxIndexSku, marginShare)).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("configure query: %w", err)
	}

	var increase []entity.DefiningOfferIncreasingMargin
	switch role {
	case auth.RoleVisitor:
		err = r.visitor.DB.SelectContext(ctx, &increase, selectQuery)
	case auth.RoleAdmin:
		err = r.admin.DB.SelectContext(ctx, &increase, selectQuery)
	default:
		return nil, fmt.Errorf("select data: cannot exec with unknown role")
	}
	if err != nil {
		return nil, fmt.Errorf("select data: %w", err)
	}
	return increase, nil
}
