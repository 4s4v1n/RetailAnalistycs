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
	personalInformationTable = `personal_information`
)

func (r *Repository) AddPersonalInformation(ctx context.Context, role uint8, item entity.PersonalInformation) error {
	insertQuery, _, err := goqu.Insert(personalInformationTable).Rows(item).ToSQL()
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

func (r *Repository) GetPersonalInformation(ctx context.Context, role uint8) ([]entity.PersonalInformation, error) {
	selectQuery, _, err := goqu.From(personalInformationTable).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("configure query: %w", err)
	}

	var personalInformation []entity.PersonalInformation
	switch role {
	case auth.RoleVisitor:
		err = r.visitor.DB.SelectContext(ctx, &personalInformation, selectQuery)
	case auth.RoleAdmin:
		err = r.admin.DB.SelectContext(ctx, &personalInformation, selectQuery)
	default:
		return nil, errors.New("select data: cannot exec with unknown role")
	}
	return personalInformation, nil
}

func (r *Repository) UpdatePersonalInformation(ctx context.Context, role uint8, item entity.PersonalInformation) error {
	updateQuery, _, err := goqu.Update(personalInformationTable).Set(goqu.Record{
		"customer_name":          item.CustomerName,
		"customer_surname":       item.CustomerSurname,
		"customer_primary_email": item.CustomerPrimaryEmail,
		"customer_primary_phone": item.CustomerPrimaryPhone,
	}).Where(goqu.C("customer_id").Eq(item.CustomerId)).Returning("customer_id").ToSQL()
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

func (r *Repository) DeletePersonalInformation(ctx context.Context, role uint8, key string) error {
	deleteQuery, _, err := goqu.Delete(personalInformationTable).
		Where(goqu.C("customer_id").Eq(key)).
		Returning("customer_id").ToSQL()
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
