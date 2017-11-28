package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"

	"github.com/howeyc/fsnotify"
)

func main() {
	logPath := flag.String("logpath", "./", "log abstruct path")
	for {

	}
	flag.Parse()
	fmt.Printf("you input path is:%s", *logPath)
}

func RunWatch(logFile string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()
	done := make(chan bool)

	go func(logFile string) {
		for {
			select {
			case <-watcher.Event:
				ShowError(logFile)
			case err := <-watcher.Error:
				log.Println("error:", err)
			}
		}
	}(logFile)
	err = watcher.Watch(logFile)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

func ShowError(logFile string) {
	f, err := ioutil.ReadFile(logFile)
	if err != nil {
		fmt.Println(err)
	}
	fs := string(f)
	pattern := `\[\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}([\+-]\d{4})?\].*`
	rep := regexp.MustCompile(pattern)
	p := rep.FindAllStringSubmatch(fs, -1)
	if 0 != len(p) {
		fmt.Println(p[len(p)-1])
	}
}
