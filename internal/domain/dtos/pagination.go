package dtos

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mxpadidar/letsgo/internal/domain/errors"
)

type PaginateDto struct {
	Limit     int    `json:"limit"`  // Maximum number of items to return
	Offset    int    `json:"offset"` // Number of items to skip before returning results
	Order     string `json:"order"`  // Order by field
	Direction string `json:"sort"`   // Direction of sorting (ASC or DESC)
}

func NewPaginateDtoFromMap(params map[string]string) (*PaginateDto, error) {
	paginte := &PaginateDto{Limit: 10, Offset: 0, Order: "id", Direction: "ASC"}

	if param, exists := params["limit"]; exists {
		limit, err := strconv.Atoi(param)
		if err != nil {
			return nil, errors.NewErr(errors.ErrValidation, fmt.Sprintf("invalid limit value: %v", err), err)
		}
		paginte.Limit = limit
	}

	if param, exists := params["offset"]; exists {
		offset, err := strconv.Atoi(param)
		if err != nil {
			return nil, errors.NewErr(errors.ErrValidation, fmt.Sprintf("invalid offset value: %v", err), err)
		}
		paginte.Offset = offset
	}

	if param := params["sort"]; param != "" {
		sort := strings.ToUpper(param)
		if sort != "ASC" && sort != "DESC" {
			return nil, errors.NewErr(errors.ErrValidation, fmt.Sprintf("invalid sort value: %v", sort), nil)
		}
		paginte.Direction = sort
	}

	if param := params["order"]; param != "" {
		paginte.Order = param
	}

	return paginte, nil
}
