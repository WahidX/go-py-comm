package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
)

type data struct {
	Name  string  `json:"name"`
	Id    int     `json:"id"`
	Score float32 `json:"score"`
}

func PostData(w http.ResponseWriter, r *http.Request) {

	log.Println("Sent data to py script")
}

func GetData(w http.ResponseWriter, r *http.Request) {
	currDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	p := path.Join(currDir, "py_scripts/get.py")
	fmt.Println(p)

	cmd := exec.Command("python", p)
	out, err := cmd.Output()
	if err != nil {
		log.Println(err)
		return
	}

	d := data{}
	err = json.Unmarshal(out, &d)
	if err != nil {
		log.Println("invalid JSON out from Python script")
		return
	}

	SendResponse(w, http.StatusOK, map[string]interface{}{
		"output": d,
	})
}
