package main

import (
	"fmt"
	"os"
	"path"
	"time"

	"./waniclient"
	"github.com/0xAX/notificator"
	"github.com/kardianos/osext"
)

var clientAPIKey string
var client *waniclient.WaniClient

func checkQueue() {
	queue := client.GetStudyQueue()

	if queue != nil && queue.Reviews > 0 {
		notify := notificator.New(notificator.Options{
			AppName: "WaniKani",
		})

		folderPath, _ := osext.ExecutableFolder()

		notify.Push("WaniKani",
			fmt.Sprintf("%d reviews available!", queue.Reviews),
			path.Join(folderPath, "wanikani.ico"),
			notificator.UR_NORMAL)
	}
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		panic("Incorrect number of arguments.")
	}
	clientAPIKey = args[0]
	client = waniclient.NewClient(clientAPIKey)

	checkQueue()

	timer := time.NewTicker(time.Minute * 10)
	go func() {
		for range timer.C {
			checkQueue()
		}
	}()

	for {
		time.Sleep(time.Second)
	}
}
