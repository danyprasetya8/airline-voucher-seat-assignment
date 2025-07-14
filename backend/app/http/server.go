package http

import (
	"airline-voucher-seat-assignment/app/http/handler"

	_ "airline-voucher-seat-assignment/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

type server struct {
	handler *handler.Handler
}

func NewServer(handler *handler.Handler) *server {
	return &server{handler}
}

func (s *server) Run() {
	r := gin.Default()

	api := r.Group("/api")
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	s.v1Route(api)

	r.Run()
}

func (s *server) v1Route(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	s.voucherRoute(v1)
}

func (s *server) voucherRoute(v1Group *gin.RouterGroup) {
	voucher := v1Group.Group("/voucher")
	voucher.GET("/", s.handler.GetVoucherList)
	voucher.POST("/check", s.handler.CheckVoucher)
	voucher.POST("/generate", s.handler.GenerateVoucher)
}
