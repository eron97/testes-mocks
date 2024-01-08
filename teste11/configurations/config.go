package configurations

type RestErr struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Error   string `json:"error"`
}
