package helper

import (
	"airline-voucher-seat-assignment/internal/constant"
	"strings"
	"time"
)

func IsBlank(str string) bool {
	return strings.TrimSpace(str) == ""
}

func FormatDate(dt time.Time) string {
	loc, _ := time.LoadLocation(constant.AsiaJakarta)
	return dt.In(loc).Format(constant.DateLayout)
}
