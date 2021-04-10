package argsmap

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func GetCommandLineArgMap(fileName string, args []string) (map[string]string, error) {
	var argHelpMap = make(map[string]oneArg)

	b := loadFile(fileName)
	jserr := json.Unmarshal(b, &argHelpMap)
	if jserr != nil {
		fmt.Printf("解析文件%s出错！\n", fileName)
	}

	var userInputArgMap = make(map[string]string)

	if len(os.Args) > 1 {
		for i := 1; i < len(args); i++ {
			flag := strings.TrimSpace(args[i])
			if usage, ok := argHelpMap[flag]; ok {
				if usage.MustHaveValue {
					v, err := getFlagValueFromArgs(usage, i, args)
					if err != nil {
						showError(usage.ArgValueErrorMsg + ",User Input:'" + v + "', Expect for:" + usage.ValueExpect)
					} else {
						log.Println("argsmap","Binding success:", flag, v)
						userInputArgMap[flag] = v
						i++
					}
				} else {
					log.Println("argsmap","Bidding success:", flag)
					userInputArgMap[flag] = "1"
				}
			} else {
				log.Println("argsmap:", "Unknown param:" + flag)
				return userInputArgMap, errors.New("Unknown param:" + flag)
			}
		}
	}
	log.Println("argsmap","------------Command line configuration------------------")
	log.Println("argsmap",userInputArgMap)
	return userInputArgMap, nil
}

type oneArg struct {
	ArgFlag      string `json:"flag"`
	ValuePattern string `json:"pattern"`
	/*是否必须有值，比如-h显示帮助就不需要，而-p 8888指定端口则必须指定，当必须指定的时候该值为true*/
	ArgValue         string `json:"value"`
	ValueExpect      string `json:"expect"`
	ArgUsage         string `json:"usage"`
	ArgValueErrorMsg string `json:"err"`
	MustHaveValue    bool   `json:"must_have_value"`
}

/**
为struct添加默认值，该方法会被自动调用
*/
func (o *oneArg) UnmarshalJSON(b []byte) error {
	type xOneArg oneArg
	/*是否必须有值，比如-h显示帮助就不需要，而-p 8888指定端口则必须指定，当必须指定的时候该值为true, 不指定则为false*/
	xo := &xOneArg{MustHaveValue: true}
	if err := json.Unmarshal(b, xo); err != nil {
		return err
	}
	*o = oneArg(*xo)
	return nil
}

func loadFile(filePath string) []byte {
	body, err := ioutil.ReadFile(filePath)
	if err != nil {
		showError("Failed to load file:" + filePath)
		return []byte{}
	}
	return body
}
func showError(msg string) {
	log.Println("argsmap",msg)
}


func getFlagValueFromArgs(usage oneArg, i int, args []string) (string, error) {
	if i >= (len(args) - 1) {
		return "", errors.New(usage.ArgValueErrorMsg)
	} else {
		return args[i+1], nil
	}
}
