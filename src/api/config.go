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

func CVGet(CV string) string {
	prompt := `现在有一些简历，我使用LayoutXLM算法对它们进行了关键信息提取(Key Instruction Extraction, KIE)，我将给出KIE处理后的信息，您需要将它们转化成标准的、人类可读的json简历格式。您只需要给出经过格式化的json格式即可，千万不要加入其他的文字。以下是您需要参照的json格式：
    {
		"name": "",
    "age": 0,
    "birthday": "",
    "tel": "",
    "email": "",
    "degree": "",
    "working_years": 0,
    "education": [
        {
            "begin": 0,
            "end": 0,
            "school": "",
            "degree": "",
            "major": "",
            "lessons": ""
        }
    ],
    "work_experience": [
        {
            "begin": 0,
            "end": 0,
            "company_or_organization": "",
            "position": ""
        }
    ],
    "school_experience": [
        {
            "company_or_organization": "",
            "position": ""
        }
    ],
    "internship_experience": [
        {
            "company_or_organization": "",
            "position": ""
        }
    ],
    "project_experience": [
        {
            "name":"",
            "desc":""
        }
    ],
    "award":[
        {
            "name":"",
            "level":""
        }
    ],
    "skill":[
        {
            "name":""
        }
    ],"self_desc":""}
""}`
	request := &Request{
		Model: model,
		Messages: []*Message{
			{
				Role:    "system",
				Content: prompt,
			},
			{
				Role:    "system",
				Content: CV,
			},
		},
	}

	response, err := gptApi.Chat(request)
	if err != nil {
		panic(err)
	}

	return response.Choices[0].Message.Content
}
