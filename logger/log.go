package logger

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

const (
	DEBUG = iota
	INFO
	ERROR
)

type logStruct struct {
	logfile  *os.File
	size     int64
	rotation bool
	m        sync.Mutex
	filePart int
	path     string
}

type Log interface {
	LogString(int, int, string) error
	LogJson(int, []byte) error
}

func NewLogger(path string, size int64, rotation bool) Log {

	fmt.Println("Initializing the log module...")

	//Create file path and file if not present

	filePath := fmt.Sprintf("%s\\%s", path, time.Now().Format("2006-01-02"))
	log.Println("File path   : ", filePath)
	if err := os.MkdirAll(path, 666); err != nil {
		log.Println("File Path creation error ", path)
		return nil
	}
	//os.fi
	//file := os.Create()

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND, 666)

	if err != nil {
		log.Println("File creation error ", filePath)
		return nil

	}

	return &logStruct{path: path, logfile: file, size: size, rotation: rotation}

}

func (l *logStruct) LogString(level int, status int, msg string) error {

	//check if the file size exceeds to the specified size
	l.m.Lock()
	defer l.m.Unlock()

	fInfo, err := os.Stat(l.logfile.Name())

	if err != nil {
		log.Println("Error getting file stat")
		return err
	}

	if fInfo.Size()/1024 >= l.size {
		log.Println("Creating New file")
		l.filePart++

		filePath := fmt.Sprintf("%s\\%s-%d", l.path, time.Now().Format("2006-01-02"), (l.filePart))
		//log.Println("File path   : ", filePath)

		f, cerr := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND, 666)

		if cerr != nil {
			log.Println("File creation error ", filePath)
			return cerr

		}
		l.logfile = f

	}

	//Prepare the log message

	var logLevel string

	switch level {
	case 1:
		logLevel = "INFO"
		break
	case 0:
		logLevel = "DEBUG"
		break
	case 2:
		logLevel = "ERROR"
		break
	}

	logMsg := fmt.Sprintf("%s:%s:%s\n", logLevel, time.Now().Format("2006-01-02 15:04:05.000000000"), msg)

	//Write the log message

	//l.logfile.Seek(0, 2)
	_, werr := l.logfile.WriteString(logMsg)
	if werr != nil {
		log.Println("Error writing to the file")
		return werr
	}
	//l.m.Unlock()

	return nil

}

func (l *logStruct) writerouting(level int, status int, msg string) {

}

func (l *logStruct) LogJson(status int, msg []byte) error {
	return nil

}
