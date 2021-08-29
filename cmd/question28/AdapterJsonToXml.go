package main

import "fmt"

/*
Паттерн адаптер, который реализует интерфейс старого обработчика
*/

type AdapterJsonToXml struct {
	handler *NewHandlerJSON
}

func (adp *AdapterJsonToXml) uploadXML() {
	adp.handler.uploadJSON()
	fmt.Println("Convert JSON to XML\nUpload XML file in service")
}
