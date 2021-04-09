package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type Page struct {
	Title string
	Body  []byte
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int, message string) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, message)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[len("/"):]
	fmt.Println(path)
	if len(path) == 0 {
		http.Redirect(w, r, "index.html", http.StatusFound)
		return
	}

	p, err := loadPage(path)
	if err != nil {
		//http.Redirect(w, r, "index.html", http.StatusFound)
		errorHandler(w, r, http.StatusNotFound, "页面找不到了:"+path)
		return
	}
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//}

	fileType, err := getFileType(path)
	switch fileType {
	case "html":
		w.Header().Set("Content-Type", "text/html")
		break
	case "css":
		w.Header().Set("Content-Type", "text/css")
		break
	case "js":
		w.Header().Set("Content-Type", "text/javascript")
		break
	}
	w.Write(p.Body)
}

func getFileType(path string) (string, error) {
	li := strings.LastIndex(path, ".")
	if li > len(path) {
		li = len(path)
	} else {
		li = li + 1
	}
	fileType := path[li:]

	if len(fileType) > 0 {
		return fileType, nil
	} else {
		return "", errors.New("未找到类型" + path)
	}
}

func loadPage(title string) (*Page, error) {
	filename := title
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func showHelp() {
	fmt.Println(
		"用法: ./gohttpserver.exe {[参数]} {[值]}\n" +
			"参数:\n" +
			"    -p: 端口号\n" +
			"    -h: 显示帮助\n" +
			"举例: 指定端口为9999" +
			"./test2.exe -p 9999")
}


func main() {
	http.HandleFunc("/", viewHandler)
	port := "8080"
	argsLen := len(os.Args)
	if argsLen >= 2 {
		firstArgs := strings.TrimSpace(os.Args[1])
		switch firstArgs {
		case "-h":
			showHelp()
			return
		case "-p":
			if argsLen >= 3 {
				secondArgs := strings.TrimSpace(os.Args[2])
				bb, _ := regexp.MatchString("\\d+", secondArgs)
				if bb {
					port = secondArgs
				} else {
					return
				}
			} else {
				fmt.Println("请输入端口号")
				return
			}
			break
		default:
			fmt.Println("参数有误")
			showHelp()
			return
		}
	}
	fmt.Printf("Listening: http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
