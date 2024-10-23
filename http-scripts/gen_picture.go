package http_scripts

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GenImageReqSt struct {
	Model             string `json:"model"`
	Prompt            string `json:"prompt"`
	NegativePrompt    string `json:"negative_prompt"`
	ImageSize         string `json:"image_size"`
	BatchSize         int    `json:"batch_size"`
	Seed              int64  `json:"seed"`
	NumInferenceSteps int    `json:"num_inference_steps"`
	GuidanceScale     int    `json:"guidance_scale"`
}

type GenImageRespSt struct {
	Images  []ImageMsgSt `json:"images"`
	Timings TimingSt     `json:"timings"`
	Seed    int          `json:"seed"`
}

type ImageMsgSt struct {
	Url string `json:"url"`
}

type TimingSt struct {
	Inference float64 `json:"inference"`
}

func GenImage(model, url, prompt, imageSize string, numInferenceSteps int) string {
	payload := GenImageReqSt{
		Model:     model,
		Prompt:    prompt,
		ImageSize: imageSize,
	}
	if len(imageSize) > 0 {
		payload.ImageSize = imageSize
	}
	if numInferenceSteps > 0 {
		payload.NumInferenceSteps = numInferenceSteps
	}

	reqBody, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	resp, err := SendHttpRequest(url, http.MethodPost, reqBody)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	response := GenImageRespSt{}
	err = json.Unmarshal(resp, &response)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	fmt.Println("prompt: ", prompt, " res_image_url: ", response.Images[0].Url, " time_cost: ", response.Timings.Inference)

	return response.Images[0].Url
}
