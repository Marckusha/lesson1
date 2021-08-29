package handleJson

import "fmt"

/*
Новый обработчик json файлов и его реализация
*/

type HandlerJSON interface {
	UploadJSON()
}

type NewHandlerJSON struct{}

func (json *NewHandlerJSON) UploadJSON() {
	fmt.Println("Upload json file")
}
