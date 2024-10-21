package routes

import (
	"net/http"

	"github.com/Dannyfsp/DaFinBotMS/utils"
)

func LoadRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /healthz", func(res http.ResponseWriter, req *http.Request) {
		utils.JSONResponse(res, 200, struct {
			Data string `json:"data"`
		}{
			Data: "Welcome to Fin BOT",
		})
	})

	mux.HandleFunc("/{$}", func(res http.ResponseWriter, req *http.Request) {
		utils.JSONResponse(res, 404, struct {
			Error string `json:"error"`
		}{
			Error: "Page Not Found",
		})
	})
}
