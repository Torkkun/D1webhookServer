package domain

type TestResponse struct {
	Message string `json:"message"`
}

type TestRequest struct {
	Test string `json:"test"`
}
