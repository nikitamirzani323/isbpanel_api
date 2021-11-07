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
type Model_keluaran struct {
	Keluaran_datekeluaran string `json:"keluaran_datekeluaran"`
	Keluaran_periode      string `json:"keluaran_periode"`
	Keluaran_nomor        string `json:"keluaran_nomor"`
}
type Model_keluaranpaitominggu struct {
	Keluaran_nomorminggu interface{} `json:"keluaran_nomor"`
}
type Model_keluaranpaitosenin struct {
	Keluaran_nomorsenin interface{} `json:"keluaran_nomor"`
}
type Model_keluaranpaitoselasa struct {
	Keluaran_nomorselasa interface{} `json:"keluaran_nomor"`
}
type Model_keluaranpaitorabu struct {
	Keluaran_nomorrabu interface{} `json:"keluaran_nomor"`
}
type Model_keluaranpaitokamis struct {
	Keluaran_nomorkamis interface{} `json:"keluaran_nomor"`
}
type Model_keluaranpaitojumat struct {
	Keluaran_nomorjumat interface{} `json:"keluaran_nomor"`
}
type Model_keluaranpaitosabtu struct {
	Keluaran_nomorsabtu interface{} `json:"keluaran_nomor"`
}

type Controller_keluaran struct {
	Pasaran_id string `json:"pasaran_id" validate:"required"`
}
