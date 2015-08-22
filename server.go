package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "bytes"
)

const (
    key string = "please get one from https://console.developers.google.com"
    url string = "https://www.google.com/speech-api/v2/recognize?output=json&lang=en-us&key=" + key
)

func translate(file, _type string) string {

    stream, err := ioutil.ReadFile(file)
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(stream))
    if err != nil {
        panic(err)
    }

    //set audio type
    req.Header.Set("Content-Type", _type)


    //submit the request
    client := &http.Client{}

    resp, err := client.Do(req)
    
    if err != nil {
    	panic(err)
    }

    defer resp.Body.Close()

    var status int = resp.StatusCode
    var body, _ = ioutil.ReadAll(resp.Body)
    fmt.Println(status)

    return string(body)
}

func handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var res = translate("/home/alexander/Desktop/audio.wav", "audio/l16; rate=16000")
    fmt.Fprintln(w, res)
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":9999", nil)
}
