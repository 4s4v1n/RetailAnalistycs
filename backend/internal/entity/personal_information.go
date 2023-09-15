package entity

type PersonalInformation struct {
	CustomerId           int64  `json:"customer_id,omitempty" csv:"customer_id" db:"customer_id" goqu:"skipinsert"`
	CustomerName         string `json:"customer_name" csv:"customer_name" db:"customer_name"`
	CustomerSurname      string `json:"customer_surname" csv:"customer_surname" db:"customer_surname"`
	CustomerPrimaryEmail string `json:"customer_primary_email" csv:"customer_primary_email" db:"customer_primary_email"`
	CustomerPrimaryPhone string `json:"customer_primary_phone" csv:"customer_primary_phone" db:"customer_primary_phone"`
}
