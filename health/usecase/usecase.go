package usecase

import (
	types "github.com/chiragsamtani/template-pattern/common/types"
	repository "github.com/chiragsamtani/template-pattern/health/repository"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

type HealthUsecase struct {
	healthRepo repository.HealthRepository
}

func NewHealthUsecase(healthRepo repository.HealthRepository) *HealthUsecase {
	return &HealthUsecase{
		healthRepo: healthRepo,
	}
}

func (h *HealthUsecase) CreateProduct(ctx echo.Context) error {
	var dto types.Products
	if err := ctx.Bind(&dto); err != nil {
		return err
	}
	// TODO: Change types.Products to DTO and map DTO to entity struct
	err := h.healthRepo.CreateProduct(&dto)
	if err != nil {
		log.Errorf("Internal server error on insert: %s", err)
		ctx.JSON(500, "Internal Server Error!")
	}
	ctx.JSON(201, "")
	return nil
}

func (h *HealthUsecase) GetProductDetail(ctx echo.Context, productCode string) error {
	entity, err := h.healthRepo.GetProductByProductCode(productCode)
	if err != nil {
		log.Errorf("Record not found for productCode: %s %s", err, productCode)
		return echo.ErrNotFound
	}
	// TODO: setup response schema to avoid returning sensitive metadata publicly
	ctx.JSON(200, entity)
	return nil
}
