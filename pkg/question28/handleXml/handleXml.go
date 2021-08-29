package handleXml

import "fmt"

/*
Интерфейс обработчика xml файлов и его релизация
*/

type HandlerXML interface {
	UploadXML()
}

type OldHandlerXML struct{}

func (xml *OldHandlerXML) UploadXML() {
	fmt.Println("Upload xml file in service")
}
