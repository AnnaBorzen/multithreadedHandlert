package workerPool

import (
	"fmt"
	"strings"
	"time"
)

type Statistics struct {
	TotalRequests   int     // общее количество запросов
	ErrorCount      int     // количество ошибок
	AverageRespTime float64 // среднее время ответа
}

func CalculateStats(input chan *Result) Statistics {
	stats := Statistics{}

	var totalRespTime float64
	count := 0

	for result := range input {
		count++
		stats.TotalRequests++

		if result.Status == "failed" {
			stats.ErrorCount++
		}

		if result.Duration > 0 {
			totalRespTime += float64(result.Duration)
			stats.AverageRespTime = totalRespTime / float64(count)
		}
	}

	return stats
}

func (stats Statistics) PrintStats() {
	fmt.Println("СТАТИСТИКА")
	fmt.Println(strings.Repeat("=", 50))
	fmt.Printf("Всего запросов: %d\n", stats.TotalRequests)
	fmt.Printf("Количество ошибок: %d\n", stats.ErrorCount)
	durationMs := time.Duration(stats.AverageRespTime).Seconds() * 1000
	fmt.Printf("Среднее время ответа: %.2f ms\n", durationMs)
}
