package domain

type HTTPClient interface {
	Post(url string, body []byte) ([]byte, error)
}
