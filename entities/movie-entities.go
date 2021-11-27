package entities

type Model_moviecategory struct {
	Movie_idcategory string      `json:"movie_idcategory"`
	Movie_category   string      `json:"movie_category"`
	Movie_list       interface{} `json:"movie_list"`
}
type Model_movie struct {
	Movie_id        int         `json:"movie_id"`
	Movie_type      string      `json:"movie_type"`
	Movie_title     string      `json:"movie_title"`
	Movie_label     string      `json:"movie_label"`
	Movie_thumbnail string      `json:"movie_thumbnail"`
	Movie_video     interface{} `json:"movie_video"`
}
type Model_movievideo struct {
	Movie_src string `json:"movie_src"`
}

type Model_movieseason struct {
	Season_id    int    `json:"season_id"`
	Season_title string `json:"season_title"`
}
type Model_movieepisode struct {
	Episode_id    int    `json:"episode_id"`
	Episode_title string `json:"episode_title"`
	Episode_src   string `json:"episode_src"`
}
type Controller_clientmovie struct {
	Client_hostname string `json:"client_hostname" validate:"required"`
}
type Controller_season struct {
	Client_hostname string `json:"client_hostname" validate:"required"`
	Movie_id        int    `json:"movie_id" validate:"required"`
}
type Controller_episode struct {
	Client_hostname string `json:"client_hostname" validate:"required"`
	Season_id       int    `json:"season_id" validate:"required"`
}
