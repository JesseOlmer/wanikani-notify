package main

import (
	"fmt"
	"os"
	"path"
	"time"

	"./waniclient"
	"github.com/TheCreeper/go-notify"

	"github.com/kardianos/osext"
)

var clientAPIKey string
var client *waniclient.WaniClient
var notificationID uint32

func checkQueue() {
	queue := client.GetStudyQueue()

	if queue != nil && queue.Reviews > 0 {

		folderPath, _ := osext.ExecutableFolder()

		ntf := notify.NewNotification("WaniKani", fmt.Sprintf("<b>%d</b> reviews available!", queue.Reviews))
		ntf.AppName = "WaniKani-notify"
		ntf.ReplacesID = notificationID
		ntf.AppIcon = path.Join(folderPath, "wanikani.ico")
		ntf.Actions = append(ntf.Actions, "default", "default")
		ntf.Hints = make(map[string]interface{})
		notificationID, _ = ntf.Show()
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

	timer := time.NewTicker(time.Second * 10)
	go func() {
		for range timer.C {
			checkQueue()
		}
	}()

	for {
		time.Sleep(time.Second)
	}
}
