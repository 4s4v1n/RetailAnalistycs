package entity

type SkuGroup struct {
	GroupId   int64  `json:"group_id,omitempty" csv:"group_id" db:"group_id" goqu:"skipinsert"`
	GroupName string `json:"group_name" csv:"group_name" db:"group_name"`
}
