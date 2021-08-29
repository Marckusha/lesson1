package main

func main() {
	service := &Service{}

	xmlHandle := OldHandlerXML{}
	service.updateService(&xmlHandle)
	adapter := AdapterJsonToXml{}
	service.updateService(&adapter)
}
