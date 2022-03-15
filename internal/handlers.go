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
	var d interface{}
	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		SendResponse(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	b, _ := json.Marshal(d)
	arg := string(b)

	currDir, _ := os.Getwd()
	p := path.Join(currDir, "py_scripts/post.py")

	cmd := exec.Command("python", p, arg)
	out, err := cmd.Output()
	if err != nil {
		log.Println(err)
		SendResponse(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	SendResponse(w, http.StatusOK, map[string]interface{}{
		"output": string(out),
	})
}

func GetData(w http.ResponseWriter, r *http.Request) {
	currDir, err := os.Getwd()
	if err != nil {
		SendResponse(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	p := path.Join(currDir, "py_scripts/get.py")
	fmt.Println(p)

	cmd := exec.Command("python", p)
	out, err := cmd.Output()
	if err != nil {
		SendResponse(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	d := data{}
	err = json.Unmarshal(out, &d)
	if err != nil {
		SendResponse(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	SendResponse(w, http.StatusOK, map[string]interface{}{
		"output": d,
	})
}
