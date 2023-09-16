package gfShutdown

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

// Will be added to the list as a single line JSON: {"id":"0"},\n
type ListJson struct {
	Id string
}

// Iterates a counter that fills the id os ListJson.Id and calls SavingList deferred.
// Listens for a notification on shutdownCh prior to each iteration.
func FillingListAndSaving(list *string, shutdownCh chan os.Signal, wg *sync.WaitGroup) {
	path := "./output/output.txt"
	var i int
	var id1 ListJson
	defer wg.Done()
	defer SavingList(list, path)

	for {
		select {
		case <-shutdownCh:
			return
		default:
		}
		// Simulating error handling
		if i == 100 {
			log.Println("Some artificial error.")
			return
		}

		id1.Id = strconv.Itoa(i)
		id2, err := json.Marshal(id1)
		if err != nil {
			log.Println(err)
		}
		*list += string(id2) + ",\n"
		i++
	}
}

// Saves the produced list to the path specified.
func SavingList(list *string, path string) {
	fmt.Println("Saving File...")
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Error opening the file: ", err)
		return
	}
	defer f.Close()
	if _, err := f.WriteString(*list + fmt.Sprintf("Saved at: %v\n", time.Now())); err != nil {
		log.Println("Error writing to the file: ", err)
		return
	}
	fmt.Println("File saved successfully!")
}
