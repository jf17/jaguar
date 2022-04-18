package javac

import (
	"fmt"
	"os"

	//"os"
	"os/exec"
)

func Compile(osVersion string, javacPath string) {
	if osVersion == "linux" {
		compileLinux()
	} else if osVersion == "windows" {
		compileWindows(javacPath)
	}
}

func compileWindows(javacPath string) {
	cmd := exec.Command(javacPath, "-d", "jar", "-cp", "jar/build/lib/rsyntaxtextarea-3.0.4.jar;jar/build/lib/commons-io-2.6.jar;jar/build/lib/autocomplete-2.6.1.jar", "src/main/java/ru/jf17/ide/*.java")

	//cmd.Stdout = os.Stdout
	//cmd.Stderr = os.Stderr
	err := cmd.Run()

	if err != nil {
		fmt.Println(err)
	}
}

func compileLinux() {
	fmt.Println("START run Compile script")
	cmd := exec.Command("/bin/sh", "build.sh")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("END Compile script")
}
