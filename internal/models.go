package internal

type ZoneResponse struct {
	Payload Payload `json:"Payload"`
	Error   string  `json:"error"`
}
type Payload struct {
	Tags map[string][]string `json:"Tags"`
}

type LoginBody struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	BasicAuthUserName string
	BasicAuthPassword string
}
