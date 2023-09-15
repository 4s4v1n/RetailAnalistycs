package repository

import (
	"APG6/internal/entity"
	"context"
	"io"
	"time"
)

type Repository interface {
	TableManager
	ViewManager
	DataManager
	FunctionManager
}

type TableManager interface {
	AddPersonalInformation(ctx context.Context, role uint8, item entity.PersonalInformation) error
	GetPersonalInformation(ctx context.Context, role uint8) ([]entity.PersonalInformation, error)
	UpdatePersonalInformation(ctx context.Context, role uint8, item entity.PersonalInformation) error
	DeletePersonalInformation(ctx context.Context, role uint8, key string) error

	AddCard(ctx context.Context, role uint8, item entity.Card) error
	GetCard(ctx context.Context, role uint8) ([]entity.Card, error)
	UpdateCard(ctx context.Context, role uint8, item entity.Card) error
	DeleteCard(ctx context.Context, role uint8, key string) error

	AddSkuGroup(ctx context.Context, role uint8, item entity.SkuGroup) error
	GetSkuGroup(ctx context.Context, role uint8) ([]entity.SkuGroup, error)
	UpdateSkuGroup(ctx context.Context, role uint8, item entity.SkuGroup) error
	DeleteSkuGroup(ctx context.Context, role uint8, key string) error

	AddProductGrid(ctx context.Context, role uint8, item entity.ProductGrid) error
	GetProductGrid(ctx context.Context, role uint8) ([]entity.ProductGrid, error)
	UpdateProductGrid(ctx context.Context, role uint8, item entity.ProductGrid) error
	DeleteProductGrid(ctx context.Context, role uint8, key string) error

	AddStore(ctx context.Context, role uint8, item entity.Store) error
	GetStore(ctx context.Context, role uint8) ([]entity.Store, error)
	UpdateStore(ctx context.Context, role uint8, item entity.Store) error
	DeleteStore(ctx context.Context, role uint8, key1 string, key2 string) error

	AddTransaction(ctx context.Context, role uint8, item entity.Transaction) error
	GetTransaction(ctx context.Context, role uint8) ([]entity.Transaction, error)
	UpdateTransaction(ctx context.Context, role uint8, item entity.Transaction) error
	DeleteTransaction(ctx context.Context, role uint8, key string) error

	AddCheck(ctx context.Context, role uint8, item entity.Check) error
	GetCheck(ctx context.Context, role uint8) ([]entity.Check, error)
	UpdateCheck(ctx context.Context, role uint8, item entity.Check) error
	DeleteCheck(ctx context.Context, role uint8, key1 string, key2 string) error

	GetDateOfAnalysingFormation(ctx context.Context, role uint8) ([]entity.DateOfAnalysingFormation, error)
	UpdateDateOfAnalysingFormation(ctx context.Context, role uint8, item entity.UpdateFormation) error
}

type ViewManager interface {
	GetPurchaseHistory(ctx context.Context, role uint8) ([]entity.PurchaseHistory, error)

	GetPeriods(ctx context.Context, role uint8) ([]entity.Period, error)

	GetGroups(ctx context.Context, role uint8) ([]entity.Group, error)

	GetCustomers(ctx context.Context, role uint8) ([]entity.Customer, error)
}

type DataManager interface {
	Import(ctx context.Context, role uint8, table string, body io.Reader) error
	Export(ctx context.Context, role uint8, table string) (any, error)
}

type FunctionManager interface {
	GrowthOfAverageCheck(ctx context.Context, role uint8, method string, first time.Time, last time.Time,
		number int32, coefficient float64, maxChurnRate float64, maxDiscountShare float64,
		marginShare float64) ([]entity.GrowthOfAverageCheck, error)
	DefiningOfferIncreasingFrequencyVisits(ctx context.Context, role uint8, first time.Time, last time.Time,
		valueTransaction int32, maxChurnRate float64, maxDiscountShare float64,
		marginShare float64) ([]entity.DefiningOfferIncreasingFrequencyVisits, error)
	DefiningOfferIncreasingMargin(ctx context.Context, role uint8, countGroup int32, maxChurnRate float64,
		maxStabilityIndex float64, maxIndexSku float64,
		marginShare float64) ([]entity.DefiningOfferIncreasingMargin, error)
}
