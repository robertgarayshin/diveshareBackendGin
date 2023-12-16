package handler

import (
	"diveshareBackendGin/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

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

func (h *Handler) GetAllCars(c *gin.Context) {
	cars, err := h.services.Car.GetAll()
	if err != nil {
		logrus.Print(err)
		newErrorResponse(c, http.StatusInternalServerError, "internal error")
		return
	}
	c.JSON(http.StatusOK, cars)
}

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
