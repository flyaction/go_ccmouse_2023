package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"

	"imooc.com/ccmouse/learngo/errhandling/filelistingserver/filelisting"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handle appHandler) func(http.ResponseWriter, *http.Request) {

	return func(writer http.ResponseWriter, request *http.Request) {

		defer func() {
			r := recover()
			if r != nil {
				log.Printf("Panic: %v", r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		err := handle(writer, request)
		if err != nil {
			log.Printf("Error handling request: %s", err.Error())

			if userErr, ok := err.(userError); ok {
				http.Error(writer, userErr.Message(), http.StatusBadRequest)
				return
			}

			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}

}

type userError interface {
	error // Error()
	Message() string
}

func main() {
	http.HandleFunc("/", errWrapper(filelisting.HandleFileList))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
