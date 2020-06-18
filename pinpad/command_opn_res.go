package pinpad

type OpnResponse struct {
	Status int `json:"status"`
}

func (opn *OpnResponse) Validate() error {
	return nil
}

func (opn *OpnResponse) Parse(rawData string) error {
	return nil
}

func (opn *OpnResponse) String() string {
	return ""
}

func (opn *OpnResponse) GetStatus() int {
	return opn.Status
}
