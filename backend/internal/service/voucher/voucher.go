package voucher

import (
	repo "airline-voucher-seat-assignment/internal/repository/voucher"
	"airline-voucher-seat-assignment/pkg/model/request"
	"airline-voucher-seat-assignment/pkg/model/response"
)

type IVoucherService interface {
	GetList(pageReq request.Pagination) (vouchers []response.GetVoucher, pageRes response.Pagination)
	Check(req request.CheckVoucher) (res response.CheckVoucher, err error)
	Generate(req request.GenerateVoucher) (res response.GenerateVoucher, err error)
}

type voucherService struct {
	voucherRepo repo.IVoucherRepository
}

func New(voucherRepo repo.IVoucherRepository) IVoucherService {
	return &voucherService{
		voucherRepo,
	}
}

func (v *voucherService) GetList(pageReq request.Pagination) (vouchers []response.GetVoucher, pageRes response.Pagination) {
	return
}

func (v *voucherService) Check(req request.CheckVoucher) (res response.CheckVoucher, err error) {
	return
}

func (v *voucherService) Generate(req request.GenerateVoucher) (res response.GenerateVoucher, err error) {
	return
}
