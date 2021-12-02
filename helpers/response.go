package helpers

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Record  interface{} `json:"record"`
	Time    string      `json:"time"`
}
type ResponseKeluaran struct {
	Status       int         `json:"status"`
	Message      string      `json:"message"`
	Record       interface{} `json:"record"`
	Paito_minggu interface{} `json:"paito_minggu"`
	Paito_senin  interface{} `json:"paito_senin"`
	Paito_selasa interface{} `json:"paito_selasa"`
	Paito_rabu   interface{} `json:"paito_rabu"`
	Paito_kamis  interface{} `json:"paito_kamis"`
	Paito_jumat  interface{} `json:"paito_jumat"`
	Paito_sabtu  interface{} `json:"paito_sabtu"`
	Time         string      `json:"time"`
}
type Responsemobilemovie struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Slider  interface{} `json:"slider"`
	Genre   interface{} `json:"genre"`
	Time    string      `json:"time"`
}
type ResponseAdmin struct {
	Status   int         `json:"status"`
	Message  string      `json:"message"`
	Record   interface{} `json:"record"`
	Listrule interface{} `json:"listruleadmin"`
	Time     string      `json:"time"`
}
type ErrorResponse struct {
	Field string
	Tag   string
}

func ErrorCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
}
