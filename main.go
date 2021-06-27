package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/jf17/jaguar/dependency"
	"github.com/jf17/jaguar/packager"
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
	xmlFile, err := os.Open("jaguar/tmp/environment.xml")
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

	dirPath := "jaguar/tmp"
	_, err := os.Stat(dirPath)
	if err != nil {
		os.MkdirAll(dirPath, os.ModePerm)
	}

	_ = ioutil.WriteFile("jaguar/tmp/environment.xml", file, 0644)
}

func createManifestFile(man manifest) {
	file, err := os.OpenFile("jar/Manifest.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	datawriter := bufio.NewWriter(file)

	_, _ = datawriter.WriteString(man.Version + "\n")
	_, _ = datawriter.WriteString(man.MainClass + "\n")
	_, _ = datawriter.WriteString(man.ClassPath + "\n")

	datawriter.Flush()
	file.Close()
}

func main() {
	man := manifest{Version: "Manifest-Version: 1.0",
		MainClass: "Main-Class: ru.jf17.ide.Main",
		ClassPath: "",
	}
	var env environment

	if _, err := os.Stat("jaguar/tmp/environment.xml"); err == nil {
		env = readEnvironment()
	} else {
		env = environment{JavaPath: "C:\\Program Files\\Java\\jdk-16\\bin\\java.exe",
			JavacPath: "C:\\Program Files\\Java\\jdk-16\\bin\\javac.exe",
			JarPath:   "C:\\Program Files\\Java\\jdk-16\\bin\\jar.exe",
		}
		writeEnvironment(env)
	}

	man.ClassPath = download.FromPom("", "")

	fmt.Println("Environment:")
	fmt.Println(env.JavacPath)
	fmt.Println(env.JarPath)
	fmt.Println(env.JavaPath)

	fmt.Println("Manifest:")
	fmt.Println(man.Version)
	fmt.Println(man.MainClass)
	fmt.Println(man.ClassPath)

	createManifestFile(man)
	jar.Pack(env.JarPath)

}
