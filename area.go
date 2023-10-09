package calculate

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// APIClient представляет клиент для работы с API.
type APIClient struct {
	BaseURL string
	APIKey  string
}

// NewAPIClient создает новый экземпляр APIClient.
func NewAPIClient(baseURL, apiKey string) *APIClient {
	return &APIClient{
		BaseURL: baseURL,
		APIKey:  apiKey,
	}
}

// SetAPIKey устанавливает API-ключ для клиента.
func (c *APIClient) SetAPIKey(apiKey string) {
	c.APIKey = apiKey
}

// GetUser делает запрос к API для получения информации о пользователе по идентификатору.
func (c *APIClient) GetUser(userID int) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/user/%d", c.BaseURL, userID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Добавляем заголовок X-MPBX-API-AUTH-TOKEN с вашим API-ключом
	req.Header.Add("X-MPBX-API-AUTH-TOKEN", c.APIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status code %d", resp.StatusCode)
	}

	var user map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}

	return user, nil
}
