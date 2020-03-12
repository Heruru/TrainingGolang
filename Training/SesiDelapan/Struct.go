package SesiDelapan

type StsDisplay struct {
	Status      Status `json:"status"`
	StatusWater string
	StatusWind  string
}

type Sts struct {
	Status Status `json:"status"`
}

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}
