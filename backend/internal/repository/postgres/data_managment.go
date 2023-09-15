package postgres

import (
	"APG6/internal/controller/auth"
	"APG6/internal/entity"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
	"io"
)

func (r *Repository) Import(ctx context.Context, role uint8, table string, body io.Reader) error {
	var tx *sqlx.Tx
	var err error

	switch role {
	case auth.RoleVisitor:
		tx, err = r.visitor.DB.Beginx()
	case auth.RoleAdmin:
		tx, err = r.admin.DB.Beginx()
	default:
		return errors.New("import data: create transaction with unknown role")
	}

	defer func() {
		_ = tx.Rollback()
	}()
	if err != nil {
		return fmt.Errorf("cannot start transaction: %w", err)
	}

	var insertQuery string

	switch table {
	case personalInformationTable:
		var data []entity.PersonalInformation
		if err = json.NewDecoder(body).Decode(&data); err != nil {
			return err
		}
		insertQuery, _, err = goqu.Insert(personalInformationTable).
			Rows(data).
			OnConflict(goqu.DoNothing()).
			ToSQL()
		if err != nil {
			return fmt.Errorf("configure insert query personal_information: %w", err)
		}
		if _, err = tx.ExecContext(ctx, insertQuery); err != nil {
			return fmt.Errorf("insert personal_information: %w", err)
		}
	case cardsTable:
		var data []entity.Card
		if err = json.NewDecoder(body).Decode(&data); err != nil {
			return err
		}
		insertQuery, _, err = goqu.Insert(cardsTable).
			Rows(data).
			OnConflict(goqu.DoNothing()).
			ToSQL()
		if err != nil {
			return fmt.Errorf("configure insert query cards: %w", err)
		}
		if _, err = tx.ExecContext(ctx, insertQuery); err != nil {
			return fmt.Errorf("insert cards: %w", err)
		}
	case skuGroupTable:
		var data []entity.SkuGroup
		if err = json.NewDecoder(body).Decode(&data); err != nil {
			return err
		}
		insertQuery, _, err = goqu.Insert(skuGroupTable).
			Rows(data).
			OnConflict(goqu.DoNothing()).
			ToSQL()
		if err != nil {
			return fmt.Errorf("configure insert query sku_group: %w", err)
		}
		if _, err = tx.ExecContext(ctx, insertQuery); err != nil {
			return fmt.Errorf("insert sku_group: %w", err)
		}
	case productGridTable:
		var data []entity.ProductGrid
		if err = json.NewDecoder(body).Decode(&data); err != nil {
			return err
		}
		insertQuery, _, err = goqu.Insert(productGridTable).
			Rows(data).
			OnConflict(goqu.DoNothing()).
			ToSQL()
		if err != nil {
			return fmt.Errorf("configure insert query product_grid: %w", err)
		}
		if _, err = tx.ExecContext(ctx, insertQuery); err != nil {
			return fmt.Errorf("insert product_grid: %w", err)
		}
	case storeTable:
		var data []entity.Store
		if err = json.NewDecoder(body).Decode(&data); err != nil {
			return err
		}
		insertQuery, _, err = goqu.Insert(storeTable).
			Rows(data).
			OnConflict(goqu.DoNothing()).
			ToSQL()
		if err != nil {
			return fmt.Errorf("configure insert query stores: %w", err)
		}
		if _, err = tx.ExecContext(ctx, insertQuery); err != nil {
			return fmt.Errorf("insert stores: %w", err)
		}
	case transactionsTable:
		var data []entity.Transaction
		if err = json.NewDecoder(body).Decode(&data); err != nil {
			return err
		}
		insertQuery, _, err = goqu.Insert(transactionsTable).
			Rows(data).
			OnConflict(goqu.DoNothing()).
			ToSQL()
		if err != nil {
			return fmt.Errorf("configure insert transactions: %w", err)
		}
		if _, err = tx.ExecContext(ctx, insertQuery); err != nil {
			return fmt.Errorf("insert transactions: %w", err)
		}
	case checksTable:
		var data []entity.Check
		if err = json.NewDecoder(body).Decode(&data); err != nil {
			return err
		}
		insertQuery, _, err = goqu.Insert(checksTable).
			Rows(data).
			OnConflict(goqu.DoNothing()).
			ToSQL()
		if err != nil {
			return fmt.Errorf("configure insert checks: %w", err)
		}
		if _, err = tx.ExecContext(ctx, insertQuery); err != nil {
			return fmt.Errorf("insert checks: %w", err)
		}
	case dateOfAnalysingFormationTable:
		var data []entity.DateOfAnalysingFormation
		if err = json.NewDecoder(body).Decode(&data); err != nil {
			return err
		}
		insertQuery, _, err = goqu.Insert(dateOfAnalysingFormationTable).
			Rows(data).
			OnConflict(goqu.DoNothing()).
			ToSQL()
		if err != nil {
			return fmt.Errorf("configure insert date_of_analysing_formation: %w", err)
		}
		if _, err = tx.ExecContext(ctx, insertQuery); err != nil {
			return fmt.Errorf("insert date_of_analysing_formation: %w", err)
		}
	default:
		return fmt.Errorf("unknown table: %s", table)
	}

	return tx.Commit()
}

func (r *Repository) Export(ctx context.Context, role uint8, table string) (any, error) {
	var data any
	var err error

	switch table {
	case personalInformationTable:
		if data, err = r.GetPersonalInformation(ctx, role); err != nil {
			return nil, err
		}
	case cardsTable:
		if data, err = r.GetCard(ctx, role); err != nil {
			return nil, err
		}
	case skuGroupTable:
		if data, err = r.GetSkuGroup(ctx, role); err != nil {
			return nil, err
		}
	case productGridTable:
		if data, err = r.GetProductGrid(ctx, role); err != nil {
			return nil, err
		}
	case storeTable:
		if data, err = r.GetStore(ctx, role); err != nil {
			return nil, err
		}
	case transactionsTable:
		if data, err = r.GetTransaction(ctx, role); err != nil {
			return nil, err
		}
	case checksTable:
		if data, err = r.GetCheck(ctx, role); err != nil {
			return nil, err
		}
	case dateOfAnalysingFormationTable:
		if data, err = r.GetDateOfAnalysingFormation(ctx, role); err != nil {
			return nil, err
		}
	case purchaseHistoryView:
		if data, err = r.GetPurchaseHistory(ctx, role); err != nil {
			return nil, err
		}
	case periodsView:
		if data, err = r.GetPeriods(ctx, role); err != nil {
			return nil, err
		}
	case groupsView:
		if data, err = r.GetGroups(ctx, role); err != nil {
			return nil, err
		}
	case customersView:
		if data, err = r.GetCustomers(ctx, role); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unknown table: %s", table)
	}

	return data, nil
}
