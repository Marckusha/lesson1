package main

import (
	"lesson1/pkg/question28/adapter"
	"lesson1/pkg/question28/handleXml"
	"lesson1/pkg/question28/service"
)

func main() {
	service := service.Service{}

	xmlHandle := handleXml.OldHandlerXML{}
	service.UpdateService(&xmlHandle)
	adapter := adapter.AdapterJsonToXml{}
	service.UpdateService(&adapter)
}
