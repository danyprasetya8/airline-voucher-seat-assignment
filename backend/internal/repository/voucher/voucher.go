package voucher

import (
	"airline-voucher-seat-assignment/internal/entity"

	"gorm.io/gorm"
)

type IVoucherRepository interface {
	GetByFlightNumberAndDate(flightNumber, flightDate string) *entity.Voucher
	Count() int64
	GetList(page, size int) []entity.Voucher
	Create(voucher *entity.Voucher) error
}

type VoucherRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) IVoucherRepository {
	return &VoucherRepository{db}
}

func (a *VoucherRepository) GetByFlightNumberAndDate(flightNumber, flightDate string) (voucher *entity.Voucher) {
	a.db.Where("flightNumber = ?", flightNumber).
		Where("flightDate = ?", flightDate).
		First(voucher)
	return
}

func (a *VoucherRepository) Count() (count int64) {
	a.db.Model(&entity.Voucher{}).Count(&count)
	return
}

func (a *VoucherRepository) GetList(page, size int) (vouchers []entity.Voucher) {
	vouchers = make([]entity.Voucher, 0)
	a.db.Limit(size).
		Offset((page - 1) * size).
		Find(&vouchers)
	return
}

func (a *VoucherRepository) Create(voucher *entity.Voucher) error {
	return a.db.Create(voucher).Error
}
