package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Weather struct {
	TargetArea   string
	HeadlineText string
	Text         string
}

func GetWeather() string {
	fmt.Println("rt")
	url := "https://www.jma.go.jp/bosai/forecast/data/forecast/130000.json"
	data := httpRequest(url)
	weather := strToJson(data)
	response := weather.ToS()
	return response
}

func httpRequest(url string) string {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal("Get Http Error:", err)
	}
	// レスポンスボディを読み込む
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal("IO Read Error:", err)
	}
	// 読み込み終わったらレスポンスボディを閉じる
	defer response.Body.Close()
	return string(body)
}

func strToJson(data string) *Weather {
	weather := new(Weather)
	if err := json.Unmarshal([]byte(data), &weather); err != nil {
		log.Fatal("JSON Unmarshall Error", err)
	}
	fmt.Println(weather)
	return weather
}

func (w *Weather) ToS() string {
	area := fmt.Sprintf("%sの天気です。\n", w.TargetArea)
	head := fmt.Sprintf("%s\n", w.HeadlineText)
	text := fmt.Sprintf("%s\n", w.Text)
	return (area + head + text)
}
