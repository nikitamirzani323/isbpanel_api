package entities

type Model_pasaran struct {
	Pasaran_id            string `json:"pasaran_id"`
	Pasaran_name          string `json:"pasaran_name"`
	Pasaran_url           string `json:"pasaran_url"`
	Pasaran_diundi        string `json:"pasaran_diundi"`
	Pasaran_jamjadwal     string `json:"pasaran_jamjadwal"`
	Pasaran_datekeluaran  string `json:"pasaran_datekeluaran"`
	Pasaran_keluaran      string `json:"pasaran_keluaran"`
	Pasaran_dateprediksi  string `json:"pasaran_dateprediksi"`
	Pasaran_bbfsprediksi  string `json:"pasaran_bbfsprediksi"`
	Pasaran_nomorprediksi string `json:"pasaran_nomorprediksi"`
}
