package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jf17/jaguar/dependency"
)

type environment struct {
	JavaPath  string `xml:"java"`
	JarPath   string `xml:"jar"`
	JavacPath string `xml:"javac"`
}

type manifest struct {
	Version   string
	MainClass string
	ClassPath string
}

func readEnvironment() environment {
	xmlFile, err := os.Open("jaguar/environment.xml")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(xmlFile)
	var env environment
	xml.Unmarshal(byteValue, &env)
	return env
}

func writeEnvironment(env environment) {
	file, _ := xml.MarshalIndent(env, "", " ")

	dirPath := "jaguar"
	_, err := os.Stat(dirPath)
	if err != nil {
		os.MkdirAll(dirPath, os.ModePerm)
	}

	_ = ioutil.WriteFile("jaguar/environment.xml", file, 0644)
}

func main() {
	man := manifest{Version: "Manifest-Version: 1.0",
		MainClass: "",
		ClassPath: "",
	}
	var env environment

	if _, err := os.Stat("jaguar/environment.xml"); err == nil {
		env = readEnvironment()
	} else {
		env = environment{JavaPath: "C:\\Program Files\\Java\\jdk-16\\bin\\java.exe",
			JavacPath: "C:\\Program Files\\Java\\jdk-16\\bin\\javac.exe",
			JarPath:   "C:\\Program Files\\Java\\jdk-16\\bin\\jar.exe",
		}
		writeEnvironment(env)
	}

	fmt.Println(env.JavacPath)
	fmt.Println(env.JarPath)
	fmt.Println(env.JavaPath)

	fmt.Println(man.Version)
	fmt.Println(man.MainClass)
	fmt.Println(man.ClassPath)

	download.FromPom("", "")

}
