package main

import (
	"sort"
	"strconv"
	"strings"
	"sync"
)

func ExecutePipeline(jobs ...job) {
	wg := &sync.WaitGroup{}
	in := make(chan interface{})
	for _, currentJob := range jobs {
		out := make(chan interface{})
		wg.Add(1)
		go func(in, out chan interface{}, j job) {
			defer wg.Done()
			j(in, out)
			close(out)
		}(in, out, currentJob)
		in = out
	}
	wg.Wait()
}

func SingleHash(in, out chan interface{}) {
	var mu *sync.Mutex = &sync.Mutex{}
	wgAll := &sync.WaitGroup{}
	for data := range in {
		wgAll.Add(1)
		go func(data string) {
			defer wgAll.Done()
			crc32Chan := make(chan string, 2)
			defer close(crc32Chan)
			wg := &sync.WaitGroup{}
			wg.Add(1)
			go func(data string, wg *sync.WaitGroup) {
				defer wg.Done()
				crc32Chan <- DataSignerCrc32(data)
			}(data, wg)

			wg.Add(1)
			go func(data string, crc32Chan chan string, wg *sync.WaitGroup, mu *sync.Mutex) {
				defer wg.Done()
				var md5Result string
				mu.Lock()
				md5Result = DataSignerMd5(data)
				mu.Unlock()
				crc32Result := DataSignerCrc32(md5Result)
				crc32Chan <- crc32Result
			}(data, crc32Chan, wg, mu)

			wg.Wait()
			result := <-crc32Chan + "~"
			result += <-crc32Chan
			out <- result

		}(strconv.Itoa(data.(int)))
	}
	wgAll.Wait()
}

func MultiHash(in, out chan interface{}) {
	WgAll := &sync.WaitGroup{}
	for data := range in {
		WgAll.Add(1)
		go func(data string) {
			defer WgAll.Done()
			wg := &sync.WaitGroup{}
			results := make([]string, 6)
			for i := 0; i < 6; i++ {
				wg.Add(1)
				go func(results []string, i int) {
					defer wg.Done()
					results[i] = DataSignerCrc32(strconv.Itoa(i) + data)
				}(results, i)
			}
			wg.Wait()
			out <- strings.Join(results, "")
		}(data.(string))
	}
	WgAll.Wait()
}

func CombineResults(in, out chan interface{}) {
	result := []string{}
	for data := range in {
		result = append(result, data.(string))
	}
	sort.Strings(result)
	out <- strings.Join(result, "_")
}
