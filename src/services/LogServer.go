package services

import (
	"log"
	"os"
)

type webLog *log.Logger

type fileLog string

func (fl fileLog) Write(data []byte) (int, error) {
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	return file.Write(data)
}

func Run(desination string, serviceName string) {
	log.New(fileLog(desination), serviceName, log.LstdFlags)
}

func RegisterHandler() {

}
