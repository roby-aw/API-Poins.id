package utils

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Pagination struct {
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Sort  string `json:"sort"`
}

func GeneratePagination(c echo.Context) Pagination {
	limit := 1000
	page := 1
	sort := "created_at desc"
	query := c.Request().URL.Query()
	fmt.Println(query)
	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
			break
		case "page":
			page, _ = strconv.Atoi(queryValue)
			break
		case "sort":
			sort = queryValue
			break

		}
	}
	fmt.Println(limit, " = ", page, " = ", sort)
	return Pagination{
		Limit: limit,
		Page:  page,
		Sort:  sort,
	}
}
