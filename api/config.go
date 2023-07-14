package api

import (
	"log"

	"gopkg.in/ini.v1"
)

var (
	// GPT configs
	gptApi *Api
	model  Model
)

func Init() {

	var file, err = ini.Load("config.ini")
	if err != nil {
		log.Fatalln("Failed to load config.ini, err:", err.Error())
	}
	loadGPT(file)
}

func loadGPT(file *ini.File) {
	var gptSection = file.Section("gpt")
	// GPTKey = gptSection.Key("GPTKey").String()
	gptApi = NewApi(gptSection.Key("GPTKey").String())
	model = Model(gptSection.Key("Model").String())
}

func GptGet(role, prompt string) string {
	request := &Request{
		Model: model,
		Messages: []*Message{
			{
				Role:    role,
				Content: prompt,
			},
		},
	}

	response, err := gptApi.Chat(request)
	if err != nil {
		panic(err)
	}

	return response.Choices[0].Message.Content
}
