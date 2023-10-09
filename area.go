package calculate

import (
	"encoding/json"
	"fmt"
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
	BaseURL  string
	APIToken string
}

// NewAPIClient создает новый экземпляр APIClient с указанным URL и токеном.
func NewAPIClient(baseURL, apiToken string) *APIClient {
	return &APIClient{
		BaseURL:  baseURL,
		APIToken: apiToken,
	}
}

// GetAllUsers выполняет запрос к API и выводит всех пользователей в формате JSON в консоль.
func (c *APIClient) GetAllUsers() error {
	url := fmt.Sprintf("%s/users", c.BaseURL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("X-MPBX-API-AUTH-TOKEN", c.BaseURL) //  заголовок X-MPBX-API-AUTH-TOKEN

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API request failed with status code %d", resp.StatusCode)
	}

	var users []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		return err
	}

	// Вывод пользователей в формате JSON в консоль
	jsonUsers, err := json.Marshal(users)
	if err != nil {
		return err
	}
	fmt.Println(string(jsonUsers))

	return nil
}
