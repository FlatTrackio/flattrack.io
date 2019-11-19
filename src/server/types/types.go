package types

type JSONMessageResponse struct {
	Message string `json:"message"`
	Version string `json:"version"`
}

type EmailStore struct {
	Emails []string
}
