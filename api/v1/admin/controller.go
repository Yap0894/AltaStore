package admin

import (
	"AltaStore/api/common"
	"AltaStore/api/middleware"
	"AltaStore/api/v1/admin/request"
	"AltaStore/api/v1/admin/response"
	"AltaStore/business/admin"

	uuid "github.com/google/uuid"
	echo "github.com/labstack/echo/v4"
)

//Controller Get item API controller
type Controller struct {
	service admin.Service
}

//NewController Construct item API controller
func NewController(service admin.Service) *Controller {
	return &Controller{
		service,
	}
}

// InsertAdmin Create new Admin handler
func (controller *Controller) InsertAdmin(c echo.Context) error {
	insertAdminRequest := new(request.InsertAdminRequest)

	if err := c.Bind(insertAdminRequest); err != nil {
		return c.JSON(common.BadRequestResponse())
	}

	admin := *insertAdminRequest.ToUpsertAdminSpec()

	err := controller.service.InsertAdmin(admin)
	if err != nil {
		return c.JSON(common.NewBusinessErrorResponse(err))
	}

	return c.JSON(common.SuccessResponseWithoutData())
}

//GetItemByID Get item by ID echo handler
func (controller *Controller) FindAdminByID(c echo.Context) error {
	id, _ := uuid.Parse(c.Param("id"))

	admin, err := controller.service.FindAdminByID(id.String())
	if err != nil {
		return c.JSON(common.NewBusinessErrorResponse(err))
	}

	response := response.NewGetAdminResponse(*admin)

	return c.JSON(common.SuccessResponseWithData(response))
}

// UpdateAdmin update existing Admin handler
func (controller *Controller) UpdateAdmin(c echo.Context) error {
	id, _ := uuid.Parse(c.Param("id"))

	adminId, err := middleware.ExtractToken(c)
	if err != nil {
		return c.JSON(common.BadRequestResponse())
	}

	updateAdminRequest := new(request.UpdateAdminRequest)

	if err = c.Bind(updateAdminRequest); err != nil {
		return c.JSON(common.BadRequestResponse())
	}
	admin := *updateAdminRequest.ToUpsertAdminSpec()

	err = controller.service.UpdateAdmin(id.String(), admin, adminId)
	if err != nil {
		return c.JSON(common.NewBusinessErrorResponse(err))
	}

	return c.JSON(common.SuccessResponseWithoutData())
}

// UpdateAdminPassword update existing Admin handler
func (controller *Controller) UpdateAdminPassword(c echo.Context) error {
	id, _ := uuid.Parse(c.Param("id"))

	updateAdminPasswordRequest := new(request.UpdateAdminPasswordRequest)

	adminId, err := middleware.ExtractToken(c)
	if err != nil {
		return c.JSON(common.BadRequestResponse())
	}

	if err = c.Bind(updateAdminPasswordRequest); err != nil {
		return c.JSON(common.BadRequestResponse())
	}
	admin := *updateAdminPasswordRequest.ToUpsertAdminSpec()

	err = controller.service.UpdateAdminPassword(id.String(), admin.NewPassword, admin.OldPassword, adminId)
	if err != nil {
		return c.JSON(common.NewBusinessErrorResponse(err))
	}

	return c.JSON(common.SuccessResponseWithoutData())
}

// DeleteAdmin delete existing Admin handler
func (controller *Controller) DeleteAdmin(c echo.Context) error {
	id, _ := uuid.Parse(c.Param("id"))

	adminId, err := middleware.ExtractToken(c)
	if err != nil {
		return c.JSON(common.BadRequestResponse())
	}

	err = controller.service.DeleteAdmin(id.String(), adminId)
	if err != nil {
		return c.JSON(common.NewBusinessErrorResponse(err))
	}

	return c.JSON(common.SuccessResponseWithoutData())
}
