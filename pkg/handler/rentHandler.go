package handler

import (
	"diveshareBackendGin/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// NewRent
// @Summary New Rent
// @Tags rents
// @Description create rent
// @ID create-rent
// @Accept  json
// @Produce  json
// @Param input body model.Rent true "rent info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /rent/new [post]
func (h *Handler) NewRent(c *gin.Context) {
	var input model.Rent
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Rent.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// GetAllRents
//
// @Summary Get All Rents
// @Tags rents
// @Description get all rents
// @ID get-rents
// @Accept  json
// @Produce  json
// @Success 200 {array} model.Rent
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /rent/ [get]
func (h *Handler) GetAllRents(c *gin.Context) {
	rents, err := h.services.Rent.GetAll()
	if err != nil {
		logrus.Print(err)
		newErrorResponse(c, http.StatusInternalServerError, "internal error")
		return
	}
	c.JSON(http.StatusOK, rents)
}

// GetRentById
//
// @Summary Get rent by id
// @Tags rents
// @Description get rent by id
// @ID get-rent-by-id
// @Accept  json
// @Produce  json
// @Param 	id path int true "id"
// @Success 200 {object} model.Rent
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /rent/{id} [get]
func (h *Handler) GetRentById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Print(err)
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	car, err := h.services.Rent.GetById(id)
	if err != nil {
		logrus.Print(err)
		newErrorResponse(c, http.StatusInternalServerError, "internal error")
		return
	}
	c.JSON(http.StatusOK, car)
}

// EditRent
//
// @Summary Edit rent
// @Tags rents
// @Description edit rent
// @ID edit-rent
// @Accept  json
// @Produce  json
// @Param 	id path int true "id"
// @Param input body model.UpdateRentInput true "rent info"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /rent/{id} [put]
func (h *Handler) EditRent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Print(err)
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	var input model.UpdateRentInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.services.Rent.Update(id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{"ok"})
}

// DeleteRent
//
// @Summary Delete rent
// @Tags rents
// @Description delete rent by id
// @ID delete-rent-by-id
// @Accept  json
// @Produce  json
// @Param 	id path int true "id"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /rent/{id} [delete]
func (h *Handler) DeleteRent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Print(err)
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	if err := h.services.Rent.Delete(id); err != nil {
		logrus.Print(err)
		newErrorResponse(c, http.StatusInternalServerError, "internal error")
		return
	}
	c.JSON(http.StatusOK, statusResponse{"ok"})
}
