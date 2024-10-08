package model

import "api/genproto/question"

type Error struct {
	Message string `json:"message"`
}

type GetAllQuestionsRequest struct {
	TopicId    string `form:"topic_id"`
	Type       string `form:"type"`
	Name       string `form:"name"`
	Number     int64  `form:"number"`
	Difficulty string `form:"difficulty"`
	Language   string `form:"language"`
}

type UpdateQuestionRequest struct {
	TopicId     string `json:"topic_id"`
	Type        string `json:"type"`
	Name        string `json:"name"`
	Number      int64  `json:"number"`
	Difficulty  string `json:"difficulty"`
	InputInfo   string `json:"input_info"`
	OutputInfo  string `json:"output_info"`
	Language    string `json:"language"`
	TimeLimit   int64  `json:"time_limit"`
	MemoryLimit int64  `json:"memory_limit"`
	Description string `json:"description"`
	Constrains  string `json:"constrains"`
	Image       string `json:"image"`
}

type CreateQuestionRequest struct {
	TopicID       string        `json:"topic_id"`
	Type          string        `json:"type"`
	Name          string        `json:"name"`
	Number        int64         `json:"number"`
	Difficulty    string        `json:"difficulty"`
	Description   string        `json:"description"`
	Image         string        `json:"image"`
	Constrains    string        `json:"constrains"`
	InputInfo     string        `json:"input_info"`
	OutputInfo    string        `json:"output_info"`
	Language      string        `json:"language"`
	InputsOutputs []InputOutput `json:"inputs_outputs"`
}

type InputOutput struct {
	Input  string `json:"input"`
	Output string `json:"output"`
}

type GetQuestionInputWithOutputsResponse struct {
	Input  *question.GetQuestionInputResponse  `json:"input"`
	Output *question.GetQuestionOutputResponse `json:"output"`
}

type GetQuestionInputWithOutput struct {
	Input  *question.GetQuestionInputResponse  `json:"input"`
	Output *question.GetQuestionOutputResponse `json:"output"` // Output maydoni
}

type GetAllQuestionInputsWithOutputsByQuestionIdResponse struct {
	InputsWithOutputs []GetQuestionInputWithOutput `json:"inputs_with_outputs"`
}
