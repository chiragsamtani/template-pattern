package usecase

import (
	types "github.com/chiragsamtani/template-pattern/common/types"
	repository "github.com/chiragsamtani/template-pattern/travel/repository"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

type TravelUsecase struct {
	travelRepo repository.TravelRepository
}

func NewTravelUsecase(travelRepo repository.TravelRepository) *TravelUsecase {
	return &TravelUsecase{
		travelRepo: travelRepo,
	}
}

func (h *TravelUsecase) CreateProduct(ctx echo.Context) error {
	var dto types.Products
	if err := ctx.Bind(&dto); err != nil {
		return err
	}
	// TODO: Change types.Products to DTO and map DTO to entity struct
	err := h.travelRepo.CreateProduct(&dto)
	if err != nil {
		log.Errorf("Internal server error on insert: %s", err)
		ctx.JSON(500, "Internal Server Error!")
	}
	ctx.JSON(201, "")
	return nil
}

func (h *TravelUsecase) GetProductDetail(ctx echo.Context, productCode string) error {
	entity, err := h.travelRepo.GetProductByProductCode(productCode)
	if err != nil {
		log.Errorf("Record not found for productCode: %s %s", err, productCode)
		return echo.ErrNotFound
	}
	// TODO: setup response schema to avoid returning sensitive metadata publicly
	ctx.JSON(200, entity)
	return nil
}
