package main

type Error struct {
	StatusCode int `json:"code"`
	Message string	`json:"message"`
}