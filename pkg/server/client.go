package server

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

var (
	emptyRequestError = errors.New("empty request")
)

type client struct {
	BaseURL string
	*http.Client
}

func NewClient(baseUrl string) *client {
	c := &http.Client{}

	return &client{
		BaseURL: baseUrl,
		Client:  c,
	}
}

func (c *client) V1MethodPost(dto *RequestDTO) (*ResponseDTO, error) {
	if dto == nil {
		return nil, emptyRequestError
	}

	payloadBuf := new(bytes.Buffer)
	err := json.NewEncoder(payloadBuf).Encode(dto)
	if err != nil {
		return nil, fmt.Errorf("failed encode data V1Method from API Server: %v", err)
	}

	res, err := c.Post(c.BaseURL, "application/json", payloadBuf)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed ReadAll body response V1Method from API Server: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusNoContent {
			return nil, nil
		}

		return nil, fmt.Errorf(`not success status V1Method from API Server.
			Actual status: %d\nBody response: %s`, res.StatusCode, string(data))
	}

	result := &ResponseDTO{}

	if err = json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("failed decode response V1Method from API Server: %v", err)
	}

	return result, err
}
