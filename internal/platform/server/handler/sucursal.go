package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wilgustavo/turnos/internal/domain/sucursal"
)

type crearSucursalRequest struct {
	ID        string `json:"id" binding:"required"`
	Nombre    string `json:"nombre" binding:"required"`
	Direccion string `json:"direccion" binding:"required"`
	Localidad string `json:"localidad" binding:"required"`
}

// CrearSucursalHandler retorna un handler para crear sucursal
func CrearSucursalHandler(crearService sucursal.CrearSucursalService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req crearSucursalRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		err := crearService(req.ID, req.Nombre, req.Direccion, req.Localidad)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
		}
		ctx.Status(http.StatusCreated)
	}
}
