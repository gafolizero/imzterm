package temporary

import (
	"fmt"
	"os"
)

func CreateTempFile(historyCount *int) *os.File {
	tempfileName := fmt.Sprintf("/home/gafoli/fingers/imzterm/.imzterm/htry-%v.png", *historyCount)
	tempfile, err := os.Create(tempfileName)
	if err != nil {
		panic(err)
	}
	*historyCount++
	return tempfile
}

func CreateDir() {
	path := "/home/gafoli/fingers/imzterm/.imzterm"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, 0755)
		if err != nil {
			fmt.Println("Failed to create directory:", err)
			return
		}
	}
}
