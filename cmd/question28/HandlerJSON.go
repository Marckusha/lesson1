package main

import "fmt"

/*
Новый обработчик json файлов и его реализация
*/

type HandlerJSON interface {
	uploadJSON()
}

type NewHandlerJSON struct{}

func (json *NewHandlerJSON) uploadJSON() {
	fmt.Println("Upload json file")
}
