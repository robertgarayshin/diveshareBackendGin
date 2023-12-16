package handler

import (
	"diveshareBackendGin/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// NewReview
// @Summary New Review
// @Tags reviews
// @Description create review
// @ID create-review
// @Accept  json
// @Produce  json
// @Param input body model.Review true "rent info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /review/new [post]
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

// GetAllReviews
//
// @Summary Get All Reviews
// @Tags reviews
// @Description get all reviews
// @ID get-reviews
// @Accept  json
// @Produce  json
// @Success 200 {array} model.Review
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /review/ [get]
func (h *Handler) GetAllReviews(c *gin.Context) {
	cars, err := h.services.Review.GetAll()
	if err != nil {
		logrus.Print(err)
		newErrorResponse(c, http.StatusInternalServerError, "internal error")
		return
	}
	c.JSON(http.StatusOK, cars)
}

// GetReviewById
//
// @Summary Get review by id
// @Tags reviews
// @Description get review by id
// @ID get-review-by-id
// @Accept  json
// @Produce  json
// @Param 	id path int true "id"
// @Success 200 {object} model.Review
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /review/{id} [get]
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

// EditReview
//
// @Summary Edit review
// @Tags reviews
// @Description edit review
// @ID edit-rent
// @Accept  json
// @Produce  json
// @Param 	id path int true "id"
// @Param input body model.UpdateReviewInput true "rent info"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /review/{id} [put]
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

// DeleteReview
//
// @Summary Delete review
// @Tags reviews
// @Description delete review by id
// @ID delete-review-by-id
// @Accept  json
// @Produce  json
// @Param 	id path int true "id"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /review/{id} [delete]
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
