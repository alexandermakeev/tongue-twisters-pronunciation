package main

type jsonObject struct {
	Object []Sentence
}

type Sentence struct {
	Content string `json:"content"`
}