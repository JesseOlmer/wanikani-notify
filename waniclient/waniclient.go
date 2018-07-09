package waniclient

import (
	"encoding/json"
	"net/http"
)

type response struct {
	User       UserInformation `json:"user_information"`
	StudyQueue StudyQueue      `json:"requested_information"`
}

type UserInformation struct {
	Username string
	Level    int
}

type StudyQueue struct {
	Lessons         int `json:"lessons_available"`
	Reviews         int `json:"reviews_available"`
	ReviewsNextHour int `json:"reviews_available_next_hour"`
	ReviewsNextDay  int `json:"reviews_available_next_day"`
}

type WaniClient struct {
	apiKey string
}

func (c *WaniClient) GetStudyQueue() *StudyQueue {
	r, err := http.Get("https://www.wanikani.com/api/user/" + c.apiKey + "/study-queue")
	if err != nil {
		return nil
	}
	defer r.Body.Close()

	var response response
	json.NewDecoder(r.Body).Decode(&response)
	return &response.StudyQueue
}

func NewClient(APIKey string) *WaniClient {
	var client WaniClient
	client.apiKey = APIKey
	return &client
}
