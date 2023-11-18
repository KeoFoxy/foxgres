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
