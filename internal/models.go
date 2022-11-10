package internal

type ZoneResponse struct {
	Payload Payload `json:"Payload"`
	Error   string  `json:"error"`
}
type Payload struct {
	Tags map[string][]string `json:"Tags"`
}
