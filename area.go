package calculate

import (
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
)

func GetRectangleArea(width, height float32) float32 {
	return width * height
}

func GetCircleArea(radius float32) float32 {
	return math.Pi * radius * radius
}

// APIClient представляет клиент для работы с API.
type APIClient struct {
	BaseURL string
}

// NewAPIClient создает новый экземпляр APIClient.
func NewAPIClient(baseURL string) *APIClient {
	return &APIClient{
		BaseURL: baseURL,
	}
}

func (c *APIClient) Get(endpoint string, authToken string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", c.BaseURL, endpoint)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-MPBX-API-AUTH-TOKEN", authToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status code %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
