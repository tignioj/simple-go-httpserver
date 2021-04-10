package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"tignioj.io/argsmap"
)

type FileObj struct {
	Title   string
	Content []byte
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

	p, err := loadWebFile(path)
	if err != nil {
		//http.Redirect(w, r, "index.html", http.StatusFound)
		errorHandler(w, r, http.StatusNotFound, "页面找不到了:"+path)
		return
	}

	fileType, err := getFileType(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		fileType = strings.ToLower(fileType)
	}
	/*写入content-Type*/
	ctm := serverConf.ContentType
	/*如果找到了相应的类型*/
	if ct, ok := ctm[fileType]; ok {
		w.Header().Set("Content-Type", ct)
	}
	w.Write(p)
}

func setContentType(w http.ResponseWriter, contentType string) {
	w.Header().Set("Content-Type", contentType)
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

func loadFile(filePath string) (*FileObj, error) {
	body, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return &FileObj{Title: filePath, Content: body}, nil
}

/**
加载页面文件
*/
func loadWebFile(filePath string) ([]byte, error) {
	path := serverConf.Root + "/" + filePath
	file, err := loadFile(path)
	if err != nil {
		return nil, err
	}
	return file.Content, nil
}

/*配置初始化*/
var serverConf = &ServerConfig{
	Port: "8080",
	Root: "./",
	ContentType: map[string]string{
		"html": "text/html",
		"css":  "text/css",
		"woff": "font/woff2",
		"js":   "text/javascript",
		"svg":  "image/svg+xml",
		"ico":  "image/x-icon",
	},
}

type ServerConfig struct {
	Port        string            `json:"port"`
	Root        string            `json:"root"`
	ContentType map[string]string `json:"content_type"`
	Header      map[string]string `json:"header"`
}

var DefaultServerConfigPath = "server-config.json"

func main() {
	//argMap, argerr := argsmap.GetCommandLineArgMap("help.json", os.Args)
	o, argerr := argsmap.NewCommandLineObj("help.json", os.Args)
	if argerr != nil {
		log.Fatal("server", argerr)
	}
	argMap:= o.GetCommandLineMap


	if argerr != nil {
		log.Fatal("server", argerr)
	}
	/*检测默认配置文件*/

	/*先加载配置文件*/
	if configPath, ok := argMap["-c"]; ok {
		initConfig(configPath)
		/*检查默认配置文件*/
	} else {
		initConfig(DefaultServerConfigPath)
	}

	/*命令行覆盖配置文件*/
	for k, v := range argMap {
		switch k {
		case "-h":
			o.ShowHelp()
			return
		case "-p":
			serverConf.Port = v
			break
		case "-r":
			serverConf.Root = v
			break
		case "-c":
			break
		default:
			fmt.Println("未知参数:" + k)
			o.ShowHelp()
			return
		}
	}

	if serverConf.Root == "" {
		serverConf.Root = "./"
	}


	fmt.Println("Apply config", serverConf)
	fmt.Printf("Listening: http://localhost:%s\n", serverConf.Port)
	http.HandleFunc("/", viewHandler)
	log.Fatal("server", http.ListenAndServe(":"+serverConf.Port, nil))
}

func initConfig(configPath string) {
	if configPath != "" {
		config, err := loadFile(configPath)
		if err != nil {
			showError("File not found:" + configPath)
			return
		}
		err = json.Unmarshal(config.Content, &serverConf)
		if err != nil {
			showError("Config file failed to load.")
			return
		}
		fmt.Println("Using config file:", config.Title)
	}
}

func showError(msg string) {
	log.Println("server:", "Error:", msg)
}
