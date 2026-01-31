package main

import (
	"fmt"
	"multithreadedHandlert/internal/url"
	"multithreadedHandlert/internal/workerPool"
	"sync"
)

func main() {
	workerCount := 5  //Количество воркеров
	jobPerWorker := 3 //задачи на воркер

	totalJobs := len(url.UrlsList)
	//Размер буферв
	bufferSize := min(workerCount*jobPerWorker, totalJobs)
	fmt.Println(bufferSize)

	jobs := make(chan *workerPool.Job, bufferSize)
	results := make(chan *workerPool.Result, bufferSize)

	var wg sync.WaitGroup

	pool := workerPool.NewWorkerPool(jobs, results, &wg, workerCount)
	pool.Start()

	// Разделяем results на два канала
	statsChan := make(chan *workerPool.Result)
	printChan := make(chan *workerPool.Result)

	go func() {
		defer close(statsChan)
		defer close(printChan)

		for result := range results {
			statsChan <- result
			printChan <- result
		}
	}()

	// Собираем статистику
	statsResult := make(chan workerPool.Statistics, 1)
	go func() {
		stats := workerPool.CalculateStats(statsChan)
		statsResult <- stats
		close(statsResult)
	}()

	// Выводим результаты
	var wgPrint sync.WaitGroup
	wgPrint.Add(1)
	go func() {
		defer wgPrint.Done()
		for result := range printChan {
			result.Print()
		}
	}()

	for i, item := range url.UrlsList {
		job := &workerPool.Job{
			ID:   i + 1,
			Data: item,
			Func: func(data interface{}) error {
				return url.ProcessingSimulation()
			},
		}
		jobs <- job
	}
	close(jobs)

	// Ждем завершения
	pool.Stop()

	// Ждем завершения вывода
	wgPrint.Wait()

	// Получаем статистику
	totalStats := <-statsResult
	totalStats.PrintStats()
}
