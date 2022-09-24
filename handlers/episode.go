package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	episodesdto "testdumpflix/dto/episode"
	dto "testdumpflix/dto/result"
	"testdumpflix/models"
	"testdumpflix/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type handlerForEpisode struct {
	EpisodeRepository repositories.EpisodeRepository
}

func HandlerEpisode(EpisodeRepository repositories.EpisodeRepository) *handlerForEpisode {
	return &handlerForEpisode{EpisodeRepository}
}

func (h *handlerForEpisode) FindEpisode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	episode, err := h.EpisodeRepository.FindEpisode()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: episode}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerForEpisode) GetEpisode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	episode, err := h.EpisodeRepository.GetEpisode(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseEpisode(episode)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerForEpisode) CreateEpisode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dataContex := r.Context().Value("dataFile")
	filename := dataContex.(string)

	idfilm, _ := strconv.Atoi(r.FormValue("film_id"))

	request := episodesdto.CreateEpisodeRequest{
		TitleEpisode:     r.FormValue("title_episode"),
		ThumbnailEpisode: filename,
		LinkFilm:         r.FormValue("linkfilm"),
		FilmID:           idfilm,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	episode := models.Episode{
		TitleEpisode:     request.TitleEpisode,
		LinkFilm:         request.LinkFilm,
		ThumbnailEpisode: request.ThumbnailEpisode,
		FilmID:           request.FilmID,
	}

	data, err := h.EpisodeRepository.CreateEpisode(episode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseEpisode(data)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerForEpisode) UpdateEpisode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(episodesdto.UpdateEpisodeRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	episode, err := h.EpisodeRepository.GetEpisode(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.TitleEpisode != "" {
		episode.TitleEpisode = request.TitleEpisode
	}

	if request.ThumbnailEpisode != "" {
		episode.ThumbnailEpisode = request.ThumbnailEpisode
	}

	if request.LinkFilm != "" {
		episode.LinkFilm = request.LinkFilm
	}

	data, err := h.EpisodeRepository.UpdateEpisode(episode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseEpisode(data)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerForEpisode) DeleteEpisode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	episode, err := h.EpisodeRepository.GetEpisode(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.EpisodeRepository.DeleteEpisode(episode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseEpisode(data)}
	json.NewEncoder(w).Encode(response)
}

func convertResponseEpisode(u models.Episode) episodesdto.EpisodeResponse {
	return episodesdto.EpisodeResponse{
		Title:         u.TitleEpisode,
		ThumbnailFilm: u.ThumbnailEpisode,
		LinkFilm:      u.LinkFilm,
	}
}
