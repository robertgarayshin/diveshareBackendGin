package handler

import (
	"diveshareBackendGin/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// NewCar
// @Summary New Car
// @Tags cars
// @Description create car
// @ID create-car
// @Accept  json
// @Produce  json
// @Param input body model.Car true "car info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /car/new [post]
func (h *Handler) NewCar(c *gin.Context) {
	var input model.Car
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Car.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// GetAllCars
//
// @Summary Get All Cars
// @Tags cars
// @Description get all cars
// @ID get-cars
// @Accept  json
// @Produce  json
// @Success 200 {array} model.Car
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /car/ [get]
func (h *Handler) GetAllCars(c *gin.Context) {
	cars, err := h.services.Car.GetAll()
	if err != nil {
		logrus.Print(err)
		newErrorResponse(c, http.StatusInternalServerError, "internal error")
		return
	}
	c.JSON(http.StatusOK, cars)
}

// GetCarById
//
// @Summary Get car by id
// @Tags cars
// @Description get car by id
// @ID get-car-by-id
// @Accept  json
// @Produce  json
// @Param 	id path int true "id"
// @Success 200 {object} model.Car
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /car/{id} [get]
func (h *Handler) GetCarById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Print(err)
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	car, err := h.services.Car.GetById(id)
	if err != nil {
		logrus.Print(err)
		newErrorResponse(c, http.StatusInternalServerError, "internal error")
		return
	}
	c.JSON(http.StatusOK, car)
}

// EditCar
//
// @Summary Edit car
// @Tags cars
// @Description edit car
// @ID edit-car
// @Accept  json
// @Produce  json
// @Param 	id path int true "id"
// @Param input body model.UpdateCarInput true "car info"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /car/{id} [put]
func (h *Handler) EditCar(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Print(err)
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	var input model.UpdateCarInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.services.Car.Update(id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{"ok"})
}

// DeleteCar
//
// @Summary Delete Car
// @Tags cars
// @Description delete car by id
// @ID delete-car-by-id
// @Accept  json
// @Produce  json
// @Param 	id path int true "id"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /car/{id} [delete]
func (h *Handler) DeleteCar(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Print(err)
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	if err := h.services.Car.Delete(id); err != nil {
		logrus.Print(err)
		newErrorResponse(c, http.StatusInternalServerError, "internal error")
		return
	}
	c.JSON(http.StatusOK, statusResponse{"ok"})
}
