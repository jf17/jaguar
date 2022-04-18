package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"

	javac "github.com/jf17/jaguar/compiler"
	download "github.com/jf17/jaguar/dependency"
	jar "github.com/jf17/jaguar/packager"
)

type project struct {
	GroupId    string `xml:"groupId"`
	ArtifactId string `xml:"artifactId"`
	FileName   string `xml:"fileName"`
	Version    string `xml:"version"`
}

type Environment struct {
	JavaPath  string `xml:"java"`
	JarPath   string `xml:"jar"`
	JavacPath string `xml:"javac"`
}

type manifest struct {
	Version   string
	MainClass string
	ClassPath string
}

func clearTargetDir() {
	err := os.RemoveAll("target")
	if err != nil {
		log.Fatal(err)
	}
	err2 := os.RemoveAll("classes")
	if err2 != nil {
		log.Fatal(err2)
	}
}

func readProject() project {
	xmlFile, err := os.Open("jaguar/project.xml")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(xmlFile)
	var proj project
	xml.Unmarshal(byteValue, &proj)
	return proj
}

func writeProject(proj project) {
	file, _ := xml.MarshalIndent(proj, "", " ")

	dirPath := "jaguar"
	_, err := os.Stat(dirPath)
	if err != nil {
		os.MkdirAll(dirPath, os.ModePerm)
	}

	_ = ioutil.WriteFile("jaguar/project.xml", file, 0644)
}

func createManifestFile(man manifest) {
	file, err := os.OpenFile("target/Manifest.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

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

func getOsVersion() string {
	if runtime.GOOS == "windows" {
		return "windows"
	} else if runtime.GOOS == "linux" {
		return "linux"
	} else if runtime.GOOS == "darwin" {
		panic("Sorry. Apple MacOS not supported !!!")
		os.Exit(10)
	} else {
		panic("OS version not recognized")
		os.Exit(42)
	}
	return ""
}

func getEnvironment(osVersion string) Environment {
	var env Environment
	if osVersion == "windows" {
		javaPath, ok := os.LookupEnv("JAGUAR-JAVA")
		if !ok {
			fmt.Println("[ERROR] JAGUAR-JAVA is not present!")
			os.Exit(2)
		} else {
			if _, err := os.Stat(javaPath); err == nil {
				env.JavaPath = javaPath
			} else {
				fmt.Println("[ERROR] JAGUAR-JAVA file is not exists!")
				os.Exit(2)
			}
		}

		javacPath, okJavac := os.LookupEnv("JAGUAR-JAVAC")
		if !okJavac {
			fmt.Println("[ERROR] JAGUAR-JAVAC is not present")
			os.Exit(2)
		} else {
			if _, err := os.Stat(javacPath); err == nil {
				env.JavacPath = javacPath
			} else {
				fmt.Println("[ERROR] JAGUAR-JAVAC file is not exists!")
				os.Exit(2)
			}
		}

		jarPath, okJar := os.LookupEnv("JAGUAR-JAR")
		if !okJar {
			fmt.Println("[ERROR] JAGUAR-JAR is not present")
			os.Exit(2)
		} else {
			if _, err := os.Stat(jarPath); err == nil {
				env.JarPath = jarPath
			} else {
				fmt.Println("[ERROR] JAGUAR-JAR file is not exists!")
				os.Exit(2)
			}
		}
	} else if osVersion == "linux" {
	} else if runtime.GOOS == "darwin" {
		panic("Sorry. Apple MacOS not supported !!!")
		os.Exit(10)
	} else {
		panic("OS version not recognized")
		os.Exit(42)
	}
	return env
}

func getHelp() {
	help := `Usage: jaguar [options]
where options include:
    clean       -clean target and classes directory
    cln         -clean target and classes directory
    compile
    package
    p 
    full-package
    fp
    download
    d`
	fmt.Println(help)
}

func main() {
	osVersion := getOsVersion()
	env := getEnvironment(osVersion)

	argsWithProg := os.Args
	//fmt.Println(argsWithProg)

	var proj project

	if _, err := os.Stat("jaguar/project.xml"); err == nil {
		proj = readProject()
	} else {
		proj = project{FileName: "MyIDE",
			GroupId:    "ru.jf17.ide",
			ArtifactId: "Main",
			Version:    "1.0-SNAPSHOT",
		}
		writeProject(proj)
	}

	man := manifest{Version: "Manifest-Version: 1.0",
		MainClass: "Main-Class: " + proj.GroupId + "." + proj.ArtifactId,
		ClassPath: "",
	}

	if len(argsWithProg) == 2 {
		oneArg := argsWithProg[1]
		if oneArg == "clean" || oneArg == "cln" {
			clearTargetDir()
		} else if oneArg == "compile" {
			javac.Compile(osVersion, env.JavacPath)
		} else if oneArg == "package" || oneArg == "p" {
			createManifestFile(man)
			jar.Pack(osVersion, env.JarPath, proj.FileName+"-"+proj.Version)
		} else if oneArg == "full-package" || oneArg == "fp" {
			clearTargetDir()
			man.ClassPath = download.FromPom("", "")
			createManifestFile(man)
			javac.Compile(osVersion, env.JavacPath)
			jar.Pack(osVersion, env.JarPath, proj.FileName+"-"+proj.Version)
		} else if oneArg == "download" || oneArg == "d" {
			clearTargetDir()
			man.ClassPath = download.FromPom("", "")
		} else if oneArg == "help" || oneArg == "h" || oneArg == "-help" || oneArg == "-h" {
			getHelp()
		}
	} else {
		getHelp()
	}

	// os.Setenv("JAGUAR-JAVA", "C:\\Program Files\\Java\\jdk-16\\bin\\java.exe")
	// os.Setenv("JAGUAR-JAVAC", "C:\\Program Files\\Java\\jdk-16\\bin\\javac.exe")
	// os.Setenv("JAGUAR-JAR", "C:\\Program Files\\Java\\jdk-16\\bin\\jar.exe")

	// fmt.Println("Environment:")
	// fmt.Println(env.JavaPath)
	// fmt.Println(env.JavacPath)
	// fmt.Println(env.JarPath)

	// fmt.Println("Manifest:")
	// fmt.Println(man.Version)
	// fmt.Println(man.MainClass)
	// fmt.Println(man.ClassPath)
}
