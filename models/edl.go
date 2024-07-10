package models

type Remarks struct {
	Idreservation string `json:"idreservation"`
	Remark        string `json:"remark"`
	Status        bool   `json:"status"`
	Final         bool   `json:"final"`
}

type RemarksDTO struct {
	Edl []Remarks `json:"edl"`
}
