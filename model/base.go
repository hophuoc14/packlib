package model

type BaseQueryParams struct {
	Id        *uint
	Offset    *int
	Limit     *int
	Sort      *string
	SortOrder *string
}
