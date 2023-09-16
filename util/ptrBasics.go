package gfShutdown

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

type ListJson struct {
	Id string
}

func FillingListAndSaving(list *string, shutdownCh chan struct{}) {
	path := "./pointer/ptrBasics/ptrBasicsOutput/output.txt"
	var i int
	var id1 ListJson
	defer SavingList(list, path)

	for {
		select {
		case <-shutdownCh:
			return
		default:
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

func SavingList(list *string, path string) {
	fmt.Println("Initializing File Saving...")
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Error opening the file: ", err)
	}
	defer f.Close()
	if _, err := f.WriteString(*list); err != nil {
		log.Println(err)
	}
	fmt.Println("File saved successfully!")
}
