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

func NewHandler(s service.Service) *handler {
	return &handler{
		service: s,
	}
}

type handler struct {
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
func (h *handler) Auth(reqCtx echo.Context) error {
	login := reqCtx.Request().Header.Get("login")
	if login == "" {
		return reqCtx.String(http.StatusBadRequest, "login is empty")
	}
	pass := reqCtx.Request().Header.Get("password")
	if pass == "" {
		return reqCtx.String(http.StatusBadRequest, "password is empty")
	}

	log.Printf("start auth handle with login: %v, pass: %v", login, pass)

	loginNumber, err := strconv.Atoi(login)
	if err != nil {
		log.Errorf("can't convert login: %v", err.Error())
		return reqCtx.String(http.StatusInternalServerError, "internal error")
	}
	level, err := h.service.FindUser(reqCtx.Request().Context(), loginNumber, pass)
	if err != nil {
		log.Errorf("can't find user: %v", err.Error())
		if errors.Is(err, pgx.ErrNoRows) {
			return reqCtx.String(http.StatusUnauthorized, "user not found")
		}
		return reqCtx.String(http.StatusInternalServerError, "internal error")
	}

	return reqCtx.JSON(http.StatusOK, AccessLevel{AccessLevel: level})
}
