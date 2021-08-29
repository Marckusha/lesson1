package main

import "fmt"

/*
Интерфейс обработчика xml файлов и его релизация
*/

type HandlerXML interface {
	uploadXML()
}

type OldHandlerXML struct{}

func (xml *OldHandlerXML) uploadXML() {
	fmt.Println("Upload xml file in service")
}
