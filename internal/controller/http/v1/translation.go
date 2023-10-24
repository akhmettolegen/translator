package v1

import (
	"encoding/json"
	"errors"
	"github.com/akhmettolegen/translator/internal/entity"
	"github.com/akhmettolegen/translator/internal/usecase"
	"github.com/akhmettolegen/translator/pkg/logger"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
)

type TranslationRoutes struct {
	translation usecase.Translation
	logger      logger.Interface
}

func NewTranslationRoutes(
	t usecase.Translation,
	l logger.Interface) *TranslationRoutes {
	return &TranslationRoutes{
		translation: t,
		logger:      l,
	}
}

func (rs *TranslationRoutes) Routes() chi.Router {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Post("/", rs.doTranslate)
		r.Get("/history", rs.history)
	})

	return r
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
// @Failure     500 {object} Response
// @Router      /translation/history [get]
func (rs *TranslationRoutes) history(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	translations, err := rs.translation.History(ctx)
	if err != nil {
		rs.logger.Error(err, "http - v1 - history")
		_ = render.Render(w, r, errResponse(http.StatusInternalServerError, "internal error"))

		return
	}

	render.JSON(w, r, historyResponse{translations})
}

type doTranslateRequest struct {
	Source      string `json:"source"      validate:"required"  example:"auto"`
	Destination string `json:"destination" validate:"required"  example:"en"`
	Original    string `json:"original"    validate:"required"  example:"текст для перевода"`
}

func (r *doTranslateRequest) validate() error {
	if r.Source == "" {
		return errors.New("source is empty")
	}
	if r.Original == "" {
		return errors.New("original is empty")
	}
	if r.Destination == "" {
		return errors.New("destination is empty")
	}

	return nil
}

// @Summary     Translate
// @Description Translate a text
// @ID          do-translate
// @Tags  	    translation
// @Accept      json
// @Produce     json
// @Param       request body doTranslateRequest true "Set up translation"
// @Success     200 {object} entity.Translation
// @Failure     400 {object} Response
// @Failure     500 {object} Response
// @Router      /translation [post]
func (rs *TranslationRoutes) doTranslate(w http.ResponseWriter, r *http.Request) {
	var request doTranslateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		rs.logger.Error(err, "http - v1 - doTranslate")
		_ = render.Render(w, r, errResponse(http.StatusBadRequest, "invalid request body"))

		return
	}

	if err := request.validate(); err != nil {
		rs.logger.Error(err, "http - v1 - doTranslate")
		_ = render.Render(w, r, errResponse(http.StatusBadRequest, "invalid request body"))

		return
	}

	translation, err := rs.translation.Translate(
		r.Context(),
		entity.Translation{
			Source:      request.Source,
			Destination: request.Destination,
			Original:    request.Original,
		},
	)
	if err != nil {
		rs.logger.Error(err, "http - v1 - doTranslate")
		_ = render.Render(w, r, errResponse(http.StatusInternalServerError, "internal error"))

		return
	}

	render.JSON(w, r, translation)
}
