package apputils

type StatusCuaca struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

type BodyCuaca struct {
	Status StatusCuaca
}

const (
	IndexHTML = "index.html"
	JSONFile  = "cuaca.json"
	SleepTime = 15
)
