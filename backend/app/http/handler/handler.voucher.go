package handler

import (
	"airline-voucher-seat-assignment/pkg/model/request"
	"airline-voucher-seat-assignment/pkg/requesthelper"
	"airline-voucher-seat-assignment/pkg/responsehelper"

	"github.com/gin-gonic/gin"
)

// GetVoucherList
//
//	@Summary	Get all vouchers
//	@Tags		Voucher
//	@Accept		json
//	@Produce	json
//	@Param		page	query	int	false	"default is 1"
//	@Param		size	query	int	false	"default is 10"
//	@Success	200		{array}	response.GetVoucher
//	@Router		/api/v1/voucher [GET]
func (h *Handler) GetVoucherList(c *gin.Context) {
	var pagination request.Pagination

	if err := c.BindQuery(&pagination); err != nil {
		responsehelper.BadRequest(c, err.Error())
		return
	}

	requesthelper.SetDefaultPagination(&pagination)

	vouchers, pageRes := h.voucherService.GetList(pagination)

	responsehelper.SuccessPage(c, vouchers, pageRes)
}

// CheckVoucher
//
//	@Summary	Check voucher by flight number and date
//	@Tags		Voucher
//	@Accept		json
//	@Produce	json
//	@Param		body	body		request.CheckVoucher	true	"Request body"
//	@Success	200		{object}	response.CheckVoucher
//	@Router		/api/v1/voucher/check [POST]
func (h *Handler) CheckVoucher(c *gin.Context) {
	var body request.CheckVoucher

	if err := c.BindJSON(&body); err != nil {
		responsehelper.BadRequest(c, err.Error())
		return
	}

	res := h.voucherService.Check(body)

	responsehelper.Success(c, res)
}

// GenerateVoucher
//
//	@Summary	Generate voucher
//	@Tags		Voucher
//	@Accept		json
//	@Produce	json
//	@Param		body	body		request.GenerateVoucher	true	"Request body"
//	@Success	200		{object}	response.GenerateVoucher
//	@Router		/api/v1/voucher/generate [POST]
func (h *Handler) GenerateVoucher(c *gin.Context) {
	var body request.GenerateVoucher

	if err := c.BindJSON(&body); err != nil {
		responsehelper.BadRequest(c, err.Error())
		return
	}

	res, err := h.voucherService.Generate(body)

	if err != nil {
		responsehelper.BadRequest(c, err.Error())
		return
	}

	responsehelper.Success(c, res)
}
