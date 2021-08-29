package service

import "lesson1/pkg/question28/handleXml"

/*
сервис, который принимает только обработчик XML файлов со старым интерфейсом
*/

type Service struct{}

func (s *Service) UpdateService(x handleXml.HandlerXML) {
	x.UploadXML()
}
