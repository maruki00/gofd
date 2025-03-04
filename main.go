package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"path"
)

func getRandomData() []byte {
	nRundom := rand.Intn(8096) + 4048
	data := make([]byte, nRundom)
	for i := range nRundom {
		data[i] = byte(rand.Intn(255) + 32)
	}
	return data
}
func distroytFile(filePath string, writeTimes int) {
	for range writeTimes {
		file, err := os.OpenFile(filePath, os.O_RDWR, 0777)
		if err != nil {
			continue
		}
		file.Write(getRandomData())
	}
	os.Remove(filePath)
}

func distroytDir(basePath string, writeTimes int) {
	dir, err := os.ReadDir(basePath)
	if err != nil {
		fmt.Println("Error : ", err.Error())
		return
	}
	counter := 0
	for _, file := range dir {
		if file.IsDir() {
			run(path.Join(basePath, file.Name()), writeTimes)
			os.Remove(path.Join(basePath, file.Name()))
			continue
		}
		distroytFile(path.Join(basePath, file.Name()), writeTimes)
		counter++
	}
}

func usage() {
	fmt.Printf("usage : ./%s -path=[the path of the file] \n", os.Args[0])
}
func main() {
	funcs :=map[string]func {
		"file": distroy,
		"dir": run
	}
	var basePath string
	var tp string
	var timeToWrite int
	flag.StringVar(&basePath, "path", "", "specify the path of files")
	flag.StringVar(&tp, "type", "file", "specify the path of files")
	flag.IntVar(&timeToWrite, "times", 1000, "specify how many times you want to write on file")
	flag.Parse()
	if basePath == "" {
		usage()
		return
	}
	fmt.Println("Process Started")
	if tp == "file" {
		distroy(basePath, timeToWrite)
	} 
	run(basePath, timeToWrite)
	fmt.Println("Finished")

}
