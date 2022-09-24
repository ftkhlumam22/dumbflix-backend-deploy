package episodesdto

type CreateEpisodeRequest struct {
	ID               int    `json:"id" gorm:"primary_key:auto_increment"`
	TitleEpisode     string `json:"title_episode" form:"title_episode" gorm:"type: varchar(255)"`
	ThumbnailEpisode string `json:"thumbnailepisode" form:"thumbnailepisode" gorm:"type: varchar(255)"`
	LinkFilm         string `json:"linkfilm" gorm:"type:text" form:"linkfilm"`
	FilmID           int    `json:"filmid" gorm:"type:int"`
}

type UpdateEpisodeRequest struct {
	ID               int    `json:"id" gorm:"primary_key:auto_increment"`
	TitleEpisode     string `json:"title_episode" form:"title_episode" gorm:"type: varchar(255)"`
	ThumbnailEpisode string `json:"thumbnailepisode" form:"thumbnailepisode" gorm:"type: varchar(255)"`
	LinkFilm         string `json:"linkfilm" gorm:"type:text" form:"linkfilm"`
}
