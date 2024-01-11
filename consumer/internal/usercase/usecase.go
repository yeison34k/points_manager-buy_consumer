package usecase

import "consumer/internal/domain"

type HttpClientCase struct {
	HTTPClient domain.HTTPClient
}

func (u *HttpClientCase) Post(url string, body []byte) ([]byte, error) {
	return u.HTTPClient.Post(url, body)
}
