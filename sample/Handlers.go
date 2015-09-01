package main

import (
	"encoding/json"
	"log"
	"time"
	"net/http"
	"strconv"
	"io"
	"os"
	"io/ioutil"
	"bytes"
	"strings"
	
	"github.com/gorilla/mux"
)

const (
	filePath string = "/path/to/file"
	key string = "please get one from https://console.developers.google.com"
    url string = "https://www.google.com/speech-api/v2/recognize?output=json&lang=en-us&key=" + key
)

func GetSentences() []Sentence {
	file, err := ioutil.ReadFile("./sample/sentences.json")
	if err != nil {
		panic(err)
	}
	var jsonType jsonObject
	json.Unmarshal(file, &jsonType)
	return jsonType.Object
}

func FindSentence(level int) Sentence {
	return GetSentences()[level - 1]
}


func GetSentence(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s\t%s\t%s", r.Method, r.RequestURI, time.Since(time.Now()))
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	sentenceId := vars["level"]
	id, err:= strconv.Atoi(sentenceId)
	if err != nil || len(GetSentences()) < id || id < 1 {
		if err := json.NewEncoder(w).Encode(Error{StatusCode:400, Message:"Level not found"}); err != nil {
			panic(err)
		}
		return
	}
	Sentence := FindSentence(id)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(Sentence); err != nil {
		panic(err)
	}
}

func PostSentence(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s\t%s\t%s", r.Method, r.RequestURI, time.Since(time.Now()))
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	sentenceId := vars["level"]
	id, err:= strconv.Atoi(sentenceId)
	if err != nil || len(GetSentences()) < id || id < 1 {
		if err := json.NewEncoder(w).Encode(Error{StatusCode:400, Message:"Level not found"}); err != nil {
			panic(err)
		}
		return
	}
	file, _, err := r.FormFile("file")
	if err != nil {
		if err := json.NewEncoder(w).Encode(Error{StatusCode:400, Message:"File not found"}); err != nil {
			panic(err)
		}
		return
	}
	defer file.Close()

	out, err := os.Create(filePath)
	if err != nil {
		if err := json.NewEncoder(w).Encode(Error{StatusCode:500, Message:"Unable to create the file for writing"}); err != nil {
			panic(err)
		}
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		if err := json.NewEncoder(w).Encode(Error{StatusCode:500, Message:"Unable to create the file for writing"}); err != nil {
			panic(err)
		}
	}
	
	var edited string = Translate(filePath)
	edited = Parse(edited)


	var original string = FindSentence(id).Content
	var success bool = strings.EqualFold(Edit(original), Edit(edited))
	result := Result{Right:success, Sentence:edited}

	if err := json.NewEncoder(w).Encode(result); err != nil {
		panic(err)
	}
}

func Translate(file string) string {
	stream, err := ioutil.ReadFile(file)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(stream))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "audio/l16; rate=44100;")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var body, _ = ioutil.ReadAll(resp.Body)
	return string(body)
}




func Edit(change string) string {
	change = strings.Replace(change, "?", "", -1)
	change = strings.Replace(change, "!", "", -1)
	change = strings.Replace(change, ",", "", -1)
	change = strings.Replace(change, ".", "", -1)
	change = strings.Replace(change, "_", " ", -1)
	change = strings.Replace(change, "-", " ", -1)
	return strings.ToLower(change)
}

func Parse(change string) string {
	var start int = strings.Index(change, ":\"") + 2
	change = change[start:]
	var end int = strings.Index(change, "\"")
	return change[:end]
}  