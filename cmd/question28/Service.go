package main

/*
сервис, который принимает только обработчик XML файлов со старым интерфейсом
*/

type Service struct{}

func (s *Service) updateService(x HandlerXML) {
	x.uploadXML()
}
