package rest

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// ****************** Sorting ******************

type SortDirection string

func NewSortDirection(direction string) SortDirection {
	switch direction {
	case "DESC":
		return SortDescending
	case "ASC":
		return SortAscending
	}
	return SortAscending
}

const (
	SortDescending SortDirection = "DESC"
	SortAscending  SortDirection = "ASC"
)

type QuerySort struct {
	Field     string
	Direction SortDirection
}

func (querySort QuerySort) ToGormOrderString() string {
	return fmt.Sprintf("%s %s", querySort.Field, querySort.Direction)
}

// ****************** Query Params ******************

type QueryParams struct {
	Offset int
	Limit  int
	Order  string
}

func (queryParams QueryParams) BuildGormOrderString() string {
	sorts := queryParams.GetSort()
	mappedSorts := make([]string, len(sorts))
	for index, item := range sorts {
		mappedSorts[index] = item.ToGormOrderString()
	}
	return strings.Join(mappedSorts[:], ",")
}

func (queryParams QueryParams) GetSort() []QuerySort {
	if queryParams.Order == "" {
		return []QuerySort{}
	}

	orderArray := strings.Split(queryParams.Order, ",")
	var sorts []QuerySort

	for _, sort := range orderArray {
		sortValues := strings.Split(sort, " ")
		field, direction := sortValues[0], sortValues[1]

		querySort := QuerySort{
			Field:     field,
			Direction: NewSortDirection(direction),
		}

		sorts = append(sorts, querySort)
	}

	return sorts
}

func NewQueryParams() *QueryParams {
	return &QueryParams{
		Offset: 0,
		Limit:  25,
		Order:  "",
	}
}

func BuildQueryParams(query url.Values) *QueryParams {
	queryParams := NewQueryParams()

	if offset := query.Get("offset"); offset != "" {
		if parsedOffset, err := strconv.Atoi(offset); err == nil {
			queryParams.Offset = parsedOffset
		}
	}

	if limit := query.Get("limit"); limit != "" {
		if parsedLimit, err := strconv.Atoi(limit); err == nil {
			queryParams.Limit = parsedLimit
		}
	}

	if order := query.Get("order"); order != "" {
		queryParams.Order = order
	}

	return queryParams
}
