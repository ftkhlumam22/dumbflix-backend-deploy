package routes

import (
	"testdumpflix/handlers"
	"testdumpflix/pkg/middleware"
	"testdumpflix/pkg/mysql"
	"testdumpflix/repositories"

	"github.com/gorilla/mux"
)

func EpisodeRoutes(r *mux.Router) {
	episodeRepository := repositories.RepositoryEpisode(mysql.DB)
	h := handlers.HandlerEpisode(episodeRepository)

	r.HandleFunc("/episodes", h.FindEpisode).Methods("GET")
	r.HandleFunc("/episode", middleware.Auth(middleware.IsAdmin(middleware.UploadFile(h.CreateEpisode)))).Methods("POST")
	r.HandleFunc("/episode/{id}", h.GetEpisode).Methods("GET")
	r.HandleFunc("/episode/{id}", h.DeleteEpisode).Methods("DELETE")
}
