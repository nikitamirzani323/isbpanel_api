package entities

type Model_moviecategory struct {
	Movie_idcategory string      `json:"movie_idcategory"`
	Movie_category   string      `json:"movie_category"`
	Movie_list       interface{} `json:"movie_list"`
}
type Model_movie struct {
	Movie_id        int    `json:"movie_id"`
	Movie_type      string `json:"movie_type"`
	Movie_title     string `json:"movie_title"`
	Movie_label     string `json:"movie_label"`
	Movie_thumbnail string `json:"movie_thumbnail"`
	Movie_video     string `json:"movie_video"`
}
