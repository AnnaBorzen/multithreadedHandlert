package url

import (
	"errors"
	"math/rand"
	"time"
)

var UrlsList = []string{
	"https://example.com",
	"https://example.com/status/201",
	"https://example.com/status/301",
	"https://example.com/status/404",
	"https://example.com/status/500",
	"https://google.com",
	"https://youtube.com",
	"https://stackoverflow.com",
	"https://github.com",
	"https://example.com",
	"https://example.com/status/201",
	"https://example.com/status/301",
	"https://example.com/status/404",
	"https://example.com/status/500",
	"https://google.com",
	"https://youtube.com",
	"https://stackoverflow.com",
	"https://github.com",
}

func ProcessingSimulation() error {
	sleepTime := time.Duration(rand.Intn(500)) * time.Millisecond
	time.Sleep(sleepTime)

	// 20% вероятность ошибки
	if rand.Float32() < 0.2 {
		return errors.New("random processing error")
	}

	return nil
}
