package filmsdto

type FilmResponse struct {
	ID            int                   `json:"id" gorm:"primary_key:auto_increment"`
	Title         string                `json:"title" form:"title" gorm:"type: varchar(255)"`
	ThumbnailFilm string                `json:"thumbnailfilm" form:"thumbnailfilm" gorm:"type: varchar(255)"`
	Description   string                `json:"desc" gorm:"type:text" form:"desc"`
	Year          int                   `json:"year" form:"year" gorm:"type: int"`
	Category      CategoryFilmResponse  `json:"category"`
	Episode       []EpisodeFilmResponse `json:"episode"`
}

type CategoryFilmResponse struct {
	Name string `json:"name"`
}

type EpisodeFilmResponse struct {
	ID               int    `json:"id"`
	TitleEpisode     string `json:"title_episode"`
	ThumbnailEpisode string `json:"thumbnailepisode"`
	LinkFilm         string `json:"linkfilm"`
	FilmID           int    `json:"film_id"`
}
