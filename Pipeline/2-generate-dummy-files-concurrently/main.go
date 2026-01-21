package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"sync"
	"time"
)

const totalFile = 3000
const contentLength = 5000

var tempPath = filepath.Join(os.Getenv("TEMP"), "chapter-A.60-worker-pool")

type FileInfo struct {
	Index       int
	FileName    string //file-<index>.txt
	WorkerIndex int    //penanda worker mana yg melakukan pekerjaan
	Err         error  // default nya kosong, nanti diisi sama objek error kalo ada error
}

func main() {
	log.Println("start")
	start := time.Now()

	generateFiles()

	duration := time.Since(start)
	log.Println("Done in", duration.Seconds())
}

func randomString(length int) string {
	randomizer := rand.New(rand.NewSource(time.Now().Unix()))
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, length)

	for i := range b {
		b[i] = letters[randomizer.Intn(len(letters))]
	}

	return string(b)
}

func generateFiles() {
	os.RemoveAll(tempPath)
	os.MkdirAll(tempPath, os.ModePerm)

	// pipeline 1: job distribution
	// bertugas men-dispatch goroutine untuk distribusi jobs
	chanFileIndex := generateFileIndexes()

	// pipeline 2: the main logic (creating files)
	// bertugas men-dispatch gorountine untuk start worker yg masing-masing worker punya tugas utama yaitu buat files
	createFilesWorker := 100
	chanFileResult := createFiles(chanFileIndex, createFilesWorker)

	// track and print output
	// tracking channel dari fan in nilai balik fungsi pipeline ke-2
	counterTotal := 0
	counterSuccess := 0
	for fileResult := range chanFileResult {
		if fileResult.Err != nil {
			log.Printf("error creating file %s. stack trace: %s", fileResult.FileName, fileResult.Err)
		} else {
			counterSuccess++
		}
		counterTotal++
	}
	log.Printf("%d/%d of total files created", counterSuccess, counterTotal)
}

func generateFileIndexes() <-chan FileInfo {
	chanOut := make(chan FileInfo)

	go func() {
		for i := 0; i < totalFile; i++ {
			chanOut <- FileInfo{
				Index:    i,
				FileName: fmt.Sprintf("file-%d.txt", i),
			}
		}
		close(chanOut)
	}()
	return chanOut
}

// bertugas untuk terima channel output dari pipeline sebelumnya (Fan-out)
// dan menjalankan beberapa worker yg memproses channel output dari pipeline sebelumnya
// lalu output dari masing-masing worker ini (channel) - langsung di merge jadi satu channel aja
func createFiles(chanIn <-chan FileInfo, numberOfWorkers int) <-chan FileInfo {
	chanOut := make(chan FileInfo)

	// wait group to control the workers
	wg := new(sync.WaitGroup)

	//allocate N of workers
	wg.Add(numberOfWorkers)

	go func() {
		for workerIndex := 0; workerIndex < numberOfWorkers; workerIndex++ {
			go func(workerIndex int) {
				// listen to `chanIn` channel for incoming jobs
				for job := range chanIn {

					// do the jobs
					filePath := filepath.Join(tempPath, job.FileName)
					content := randomString(contentLength)
					err := os.WriteFile(filePath, []byte(content), os.ModePerm)
					log.Println("worker", workerIndex, "working on", job.FileName, "file generation")

					// construct the job's result, and send it to `chanOut`
					chanOut <- FileInfo{
						FileName:    job.FileName,
						WorkerIndex: workerIndex,
						Err:         err,
					}
				}
				// if `chanIn` is closed, and the remaining jobs are finished
				// only then we mark the worker as complate
				wg.Done()
			}(workerIndex)
		}
	}()

	// wait until `chanIn` closed and then all workers are done
	// because right after that - we need to close the `chanOut` channel
	go func() {
		wg.Wait()
		close(chanOut)
	}()

	return chanOut
}
