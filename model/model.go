package model

type Error struct {
	Message string `json:"message"`
}

type GetAllQuestionsRequest struct {
	TopicName  string `form:"topic_name"`
	Type       string `form:"type"`
	Name       string `form:"name"`
	Number     int64  `form:"number"`
	Difficulty string `form:"difficulty"`
}

type UpdateQuestionRequest struct {
	TopicId    string `json:"topic_id"`
	Type       string `json:"type"`
	Name       string `json:"name"`
	Number     int64  `json:"number"`
	Difficulty string `json:"difficulty"`
	InputInfo  string `json:"input_info"`
	OutputInfo string `json:"output_info"`
}
