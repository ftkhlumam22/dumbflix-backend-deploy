package filmsdto

type CreateFilmRequest struct {
	Title            string `json:"title" form:"title" gorm:"type: varchar(255)"`
	ThumbnailFilm    string `json:"thumbnailfilm" form:"thumbnailfilm" gorm:"type: varchar(255)"`
	Description      string `json:"description" gorm:"type:text" form:"desc"`
	Year             int    `json:"year" form:"year" gorm:"type: int"`
	Category         string `json:"category"`
	TitleEpisode     string `json:"title_episode" form:"title_episode" gorm:"type: varchar(255)"`
	ThumbnailEpisode string `json:"thumbnailepisode" form:"thumbnailepisode" gorm:"type: varchar(255)"`
	LinkFilm         string `json:"linkfilm" gorm:"type:text" form:"linkfilm"`
}

type UpdateFilmRequest struct {
	Title         string `json:"title" form:"name" gorm:"type: varchar(255)"`
	ThumbnailFilm string `json:"thumbnailFilm" form:"image" gorm:"type: varchar(255)"`
	Description   string `json:"description" gorm:"type:text" form:"desc"`
	Year          int    `json:"year" form:"price" gorm:"type: int"`
	CategoryID    int    `json:"category"`
}
