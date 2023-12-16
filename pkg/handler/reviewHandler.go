package handler

import (
	"diveshareBackendGin/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func (h *Handler) NewReview(c *gin.Context) {
	var input model.Review
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Review.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) GetAllReviews(c *gin.Context) {
	cars, err := h.services.Review.GetAll()
	if err != nil {
		logrus.Print(err)
		newErrorResponse(c, http.StatusInternalServerError, "internal error")
		return
	}
	c.JSON(http.StatusOK, cars)
}

func (h *Handler) GetReviewById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Print(err)
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	car, err := h.services.Review.GetById(id)
	if err != nil {
		logrus.Print(err)
		newErrorResponse(c, http.StatusInternalServerError, "internal error")
		return
	}
	c.JSON(http.StatusOK, car)
}

func (h *Handler) EditReview(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Print(err)
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	var input model.UpdateReviewInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.services.Review.Update(id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) DeleteReview(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Print(err)
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	if err := h.services.Review.Delete(id); err != nil {
		logrus.Print(err)
		newErrorResponse(c, http.StatusInternalServerError, "internal error")
		return
	}
	c.JSON(http.StatusOK, statusResponse{"ok"})
}
