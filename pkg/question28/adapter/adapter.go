package adapter

import (
	"fmt"
	"lesson1/pkg/question28/handleJson"
)

/*
Паттерн адаптер, который реализует интерфейс старого обработчика
*/

type AdapterJsonToXml struct {
	handler *handleJson.NewHandlerJSON
}

func (adp *AdapterJsonToXml) UploadXML() {
	adp.handler.UploadJSON()
	fmt.Println("Convert JSON to XML\nUpload XML file in service")
}
