package voucher

import (
	"airline-voucher-seat-assignment/internal/constant"
	"airline-voucher-seat-assignment/internal/entity"
	repo "airline-voucher-seat-assignment/internal/repository/voucher"
	"airline-voucher-seat-assignment/pkg/helper"
	"airline-voucher-seat-assignment/pkg/model"
	"airline-voucher-seat-assignment/pkg/model/request"
	"airline-voucher-seat-assignment/pkg/model/response"
	"airline-voucher-seat-assignment/pkg/responsehelper"
	"fmt"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type IVoucherService interface {
	GetList(pageReq request.Pagination) (vouchers []response.GetVoucher, pageRes *response.Pagination)
	Check(req request.CheckVoucher) (res *response.CheckVoucher)
	Generate(req request.GenerateVoucher) (res *response.GenerateVoucher, err error)
}

type voucherService struct {
	voucherRepo repo.IVoucherRepository
}

func New(voucherRepo repo.IVoucherRepository) IVoucherService {
	return &voucherService{
		voucherRepo,
	}
}

func (v *voucherService) GetList(pageReq request.Pagination) (vouchers []response.GetVoucher, pageRes *response.Pagination) {
	vouchers = make([]response.GetVoucher, 0)

	total := v.voucherRepo.Count()

	entities := v.voucherRepo.GetList(pageReq.Page, pageReq.Size)

	for _, en := range entities {
		voucher := response.GetVoucher{
			ID:           en.ID,
			CrewID:       en.CrewID,
			CrewName:     en.CrewName,
			FlightNumber: en.FlightNumber,
			FlightDate:   en.FlightDate,
			AircraftType: en.AircraftType,
			Seat:         strings.Split(en.Seat, ","),
			CreatedAt:    helper.FormatDate(en.CreatedAt),
		}
		vouchers = append(vouchers, voucher)
	}

	pageRes = responsehelper.ToPagination(&pageReq, total)

	return
}

func (v *voucherService) Check(req request.CheckVoucher) (res *response.CheckVoucher) {
	voucher := v.voucherRepo.GetByFlightNumberAndDate(req.FlightNumber, req.Date)
	return &response.CheckVoucher{
		Exists: voucher != nil,
	}
}

func (v *voucherService) Generate(req request.GenerateVoucher) (res *response.GenerateVoucher, err error) {
	existVoucher := v.voucherRepo.GetByFlightNumberAndDate(req.FlightNumber, req.Date)

	if existVoucher != nil {
		log.Errorf("flight number %s already generate voucher at %s", req.FlightNumber, req.Date)
		return nil, fmt.Errorf("flight number %s already generate voucher at %s", req.FlightNumber, req.Date)
	}

	aircraftType := v.findAircraftType(req.Aircraft)

	if aircraftType == nil {
		log.Errorf("aircraft type %s not found", req.Aircraft)
		return nil, fmt.Errorf("aircraft type %s not found", req.Aircraft)
	}

	seats := v.generateSeats(aircraftType)

	newVoucher := &entity.Voucher{
		CrewID:       req.ID,
		CrewName:     req.Name,
		FlightNumber: req.FlightNumber,
		FlightDate:   req.Date,
		AircraftType: req.Aircraft,
		Seat:         strings.Join(seats, ","),
	}
	err = v.voucherRepo.Create(newVoucher)

	if err != nil {
		log.Errorf("failed to create new voucher: %s", err.Error())
		return nil, fmt.Errorf("failed to create new voucher: %s", err.Error())
	}

	return &response.GenerateVoucher{
		Success: true,
		Seats:   seats,
	}, nil
}

func (v *voucherService) findAircraftType(aircraftType string) *model.AircraftType {
	for _, at := range constant.AircraftTypes {
		if at.Name == aircraftType {
			return at
		}
	}
	return nil
}

func (v *voucherService) generateSeats(aircraftType *model.AircraftType) []string {
	seats := make([]string, 0)
	for i := aircraftType.RowStart; i <= aircraftType.RowEnd; i++ {
		for _, s := range aircraftType.Seats {
			seats = append(seats, strconv.FormatInt(int64(i), 10)+s)
		}
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	r.Shuffle(len(seats), func(i, j int) {
		seats[i], seats[j] = seats[j], seats[i]
	})

	return seats[:3]
}
