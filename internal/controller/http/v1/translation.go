package v1

import (
	"encoding/json"
	"github.com/akhmettolegen/translator/internal/entity"
	"github.com/akhmettolegen/translator/internal/usecase"
	resp "github.com/akhmettolegen/translator/pkg/api/response"
	"github.com/akhmettolegen/translator/pkg/logger"
	"github.com/go-chi/render"
	"net/http"
)

type translationRoutes struct {
	t usecase.Translation
	l logger.Interface
}

func newTranslationRoutes(handler *gin.RouterGroup, t usecase.Translation, l logger.Interface) {
	r := &translationRoutes{t, l}

	h := handler.Group("/translation")
	{
		h.GET("/history", r.history)
		h.POST("/do-translate", r.doTranslate)
	}
}

type historyResponse struct {
	History []entity.Translation `json:"history"`
}

// @Summary     Show history
// @Description Show all translation history
// @ID          history
// @Tags  	    translation
// @Accept      json
// @Produce     json
// @Success     200 {object} historyResponse
// @Failure     500 {object} response
// @Router      /translation/history [get]
func (c *translationRoutes) history(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	translations, err := c.t.History(ctx)
	if err != nil {
		c.l.Error(err, "http - v1 - history")
		errorResponse(ctx, http.StatusInternalServerError, "database problems")
		render.JSON(w, r, resp.Error("internal error"))

		return
	}

	render.JSON(w, r, historyResponse{translations})
}

type doTranslateRequest struct {
	Source      string `json:"source"       binding:"required"  example:"auto"`
	Destination string `json:"destination"  binding:"required"  example:"en"`
	Original    string `json:"original"     binding:"required"  example:"текст для перевода"`
}

// @Summary     Translate
// @Description Translate a text
// @ID          do-translate
// @Tags  	    translation
// @Accept      json
// @Produce     json
// @Param       request body doTranslateRequest true "Set up translation"
// @Success     200 {object} entity.Translation
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /translation/do-translate [post]
func (c *translationRoutes) doTranslate(w http.ResponseWriter, r *http.Request) {
	const op = "handlers.url.save.New"
	var request doTranslateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		c.l.Error(err, op)
		render.JSON(w, r, resp.Error("invalid request body"))

		return
	}

	translation, err := c.t.Translate(
		r.Context(),
		entity.Translation{
			Source:      request.Source,
			Destination: request.Destination,
			Original:    request.Original,
		},
	)
	if err != nil {
		c.l.Error(err, op)
		render.JSON(w, r, resp.Error("translation service problems"))

		return
	}

	render.JSON(w, r, translation)
}
