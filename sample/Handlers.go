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


var sentences []Sentence = []Sentence{
	Sentence{Content: "Hello World!"},
	Sentence{Content: "World Wide Web."},
	Sentence{Content: "Three free throws."},
	Sentence{Content: "The blue bluebird blinks."},
	Sentence{Content: "Red leather yellow leather."},
	Sentence{Content: "Four fine fresh fish for you."},
	Sentence{Content: "Kitty caught the kitten in the kitchen."},
	Sentence{Content: "How can a clam cram in a clean cream can?"},
	Sentence{Content: "Can you can a can as a canner can can a can?"},
	Sentence{Content: "I thought I thought of thinking of thanking you."},
	Sentence{Content: "To put a pipe in byte mode, type PIPE_TYPE_BYTE."},
	Sentence{Content: "I scream, you scream, we all scream for ice cream!"},
	Sentence{Content: "If you want to buy, buy, if you don't want to buy, bye, bye!"},
	Sentence{Content: "One-one was a race horse. Two-two was one too. One-one won one race. Two-two won one too."},
	Sentence{Content: "Whether the weather is warm, whether the weather is hot, we have to put up with the weather, whether we like it or not."},
}

func GetSentence(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s\t%s\t%s", r.Method, r.RequestURI, time.Since(time.Now()))
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	sentenceId := vars["level"]
	id, err:= strconv.Atoi(sentenceId)
	if err != nil || len(sentences) < id || id < 1 {
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
	if err != nil || len(sentences) < id || id < 1 {
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


func FindSentence(id int) Sentence {
	return sentences[id - 1]
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