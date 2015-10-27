package config

import (
	"fmt"
	"os"
	"testing"
)

var initCasesContext = `
process = --args=/bin/bash --cwd=/bin --terminal=true;--args=/bin/bash;--cwd=/bin
args  = --args=xxxx
capablitis    = --cwd=xxx
mount= --mountpoint=bbbb
`

func TestConfig(t *testing.T) {
	f, err := os.Create("cases.conf")
	if err != nil {
		t.Fatal(err)
	}

	_, err = f.WriteString(initCasesContext)
	if err != nil {
		f.Close()
		t.Fatal(err)
	}
	f.Close()
	// defer os.Remove("cases.conf")

	if data := GetConfig("process"); len(data) != 3 {
		t.Fatal("Get process err", data)
	} else if data[0] != "--args=/bin/bash --cwd=/bin --terminal=true" {
		t.Fatal("Get first params of process err")
	}

	for _, c := range CaseArray {
		fmt.Println(c)
	}

	fmt.Println(ConfigLen)
}
