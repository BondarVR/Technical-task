package handler

import (
	"net/http"
	"technical-task/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/auth/sign-up", h.signUp)
	router.HandleFunc("/auth/sign-in", h.signIn)

	router.HandleFunc("/product/create", h.userIdentity(h.createProduct))
	router.HandleFunc("/product/get", h.userIdentity(h.getProductByID))
	router.HandleFunc("/product/update", h.userIdentity(h.updateProduct))
	router.HandleFunc("/product/delete", h.userIdentity(h.deleteProduct))

	return router
}
