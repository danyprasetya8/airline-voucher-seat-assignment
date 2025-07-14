package handler

import (
	"airline-voucher-seat-assignment/internal/service/voucher"
)

type Handler struct {
	voucherService voucher.IVoucherService
}

func New(
	voucherService voucher.IVoucherService,
) *Handler {
	return &Handler{
		voucherService,
	}
}
