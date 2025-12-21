package model

import "github.com/google/uuid"

type BaseQueryParams struct {
	Id        *uuid.UUID
	Offset    *int
	Limit     *int
	Sort      *string
	SortOrder *string
}
