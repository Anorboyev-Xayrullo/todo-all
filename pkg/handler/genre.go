package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/zhashkevych/todo-app/model"
	"net/http"
	"strconv"
)

// creteGenre
// @Summary Create genre
// @Security ApiKeyAuth
// @Tags genre
// @Description create genre
// @ID create-genr
// @Accept  json
// @Produce  json
// @Param input body model.Genre true "genre info"
// @Success 200 {integer} model.SuccessData
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/genr/ [post]
func (h *Handler) creteGenre(c *gin.Context) {
	var input model.Genre
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Genre.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, model.SuccessData{
		Data: id,
	})
}
type GetGenreResponse struct {
	Genre []model.Genre `json:"genre"`
}
//GetGenre
// @Summary Get Genre	
// @Security ApiKeyAuth
// @Tags genre
// @Description get genre
// @ID get-genr
// @Accept  json
// @Produce  json
// @Success 200 {object} GetGenreResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/genr/ [get]
func (h *Handler) GetGenre(c *gin.Context) {
	genrelist, err := h.services.Genre.GetGenre()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, GetGenreResponse{
		Genre: genrelist,
	})
}

// GetGenreById
// @Summary Get Genre By id
// @Security ApiKeyAuth
// @Tags genre
// @Description get genre by id
// @ID get-genr-by-id
// @Accept  json
// @Produce  json
// @Param id query int true "id"
// @Success 200 {object} model.SuccessData
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/genr/id  [get]
func (h *Handler) GetGenreById(c *gin.Context) {
	id, err := strconv.Atoi(c.DefaultQuery("id", "0"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	list, err := h.services.Genre.GetGenreById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, model.SuccessData{
		Data: list,
	})
}


// DeleteGenre
// @Summary delete genre
// @Security ApiKeyAuth
// @Tags genre
// @Description delete genre
// @ID delete genr
// @Accept  json
// @Produce  json
// @Param id query int true "ID"
// @Success 200 {object} model.SuccessData
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/genr/ [delete]
func (h *Handler) DeleteGenre(c *gin.Context) {

	id, err := strconv.Atoi(c.DefaultQuery("id", "0"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.Genre.DeleteGenre(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, model.SuccessData{
		Data: "ok",
	})
}

// UpdateGenre
// @Summary update genre
// @Security ApiKeyAuth
// @Tags genre
// @Description update genre
// @ID update genr
// @Accept  json
// @Produce  json
// @Param id query int true "ID"
// @Param input body model.UpdateGenreInput true "book info"
// @Success 200 {object} model.SuccessData
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/genr/ [put]
func (h *Handler) UpdateGenre(c *gin.Context) {
	genreID, err := strconv.Atoi(c.DefaultQuery("id", "0"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input model.UpdateGenreInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.services.Genre.UpdateGenre(int(genreID), input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, model.SuccessData{
		Data: "ok",
	})
}
