package batch

type Error struct {
	Code    string
	Message string
}

type Response struct {
	Response interface{} `json:"Response"`
}
