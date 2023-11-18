package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	"foxgres/internal/service"
)

func NewHandler(s service.Service) *Handler {
	return &Handler{
		service: s,
	}
}

type Handler struct {
	service service.Service
}

// Auth godoc
//
//		@Summary		authentication
//		@Description	Поиск пользователя в таблице users
//	 @Tags authentication
//		@Accept			json
//		@Produce		json
//		@Param			login header string true "Логин"
//		@Param			password header string true "Пароль"
//		@Success		200 {object} AccessLevel
//		@Failure		400	{object} ResponseErr
//		@Failure		401	{object} ResponseErr
//		@Failure		500	{object} ResponseErr
//		@Router			/auth [get]
func (h *Handler) Auth(reqCtx echo.Context) error {
	login := reqCtx.Request().Header.Get("login")
	if login == "" {
		return reqCtx.JSON(http.StatusBadRequest, ResponseErr{Message: "login is empty"})
	}
	pass := reqCtx.Request().Header.Get("password")
	if pass == "" {
		return reqCtx.JSON(http.StatusBadRequest, ResponseErr{Message: "password is empty"})
	}

	log.Printf("start auth handle with login: %v, pass: %v", login, pass)

	loginNumber, err := strconv.Atoi(login)
	if err != nil {
		log.Errorf("can't convert login: %v", err.Error())
		return reqCtx.JSON(http.StatusInternalServerError, ResponseErr{Message: "internal error"})
	}
	level, err := h.service.FindUser(reqCtx.Request().Context(), loginNumber, pass)
	if err != nil {
		log.Errorf("can't find user: %v", err.Error())
		if errors.Is(err, pgx.ErrNoRows) {
			return reqCtx.JSON(http.StatusUnauthorized, ResponseErr{Message: "user not found"})
		}
		return reqCtx.JSON(http.StatusInternalServerError, ResponseErr{Message: "internal error"})
	}

	return reqCtx.JSON(http.StatusOK, AccessLevel{AccessLevel: level})
}
