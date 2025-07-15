package voucher

import (
	"airline-voucher-seat-assignment/internal/entity"

	"gorm.io/gorm"
)

type IVoucherRepository interface {
	GetByFlightNumberAndDate(flightNumber, flightDate string) (voucher entity.Voucher)
	Count() int64
	GetList(page, size int) []entity.Voucher
	Create(voucher *entity.Voucher) error
}

type voucherRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) IVoucherRepository {
	return &voucherRepository{db}
}

func (a *voucherRepository) GetByFlightNumberAndDate(flightNumber, flightDate string) (voucher entity.Voucher) {
	a.db.Where("flight_number = ?", flightNumber).
		Where("flight_date = ?", flightDate).
		First(&voucher)
	return
}

func (a *voucherRepository) Count() (count int64) {
	a.db.Model(&entity.Voucher{}).Count(&count)
	return
}

func (a *voucherRepository) GetList(page, size int) (vouchers []entity.Voucher) {
	vouchers = make([]entity.Voucher, 0)
	a.db.Limit(size).
		Offset((page - 1) * size).
		Find(&vouchers)
	return
}

func (a *voucherRepository) Create(voucher *entity.Voucher) error {
	return a.db.Create(voucher).Error
}
