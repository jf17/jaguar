package jar

import (
	"fmt"
	//"os"
	"os/exec"
)

func Pack(osVersion string, jarPath string, fileName string) {
	if osVersion == "linux" {
		packLinux()
	} else if osVersion == "windows" {
		packWindows(jarPath, fileName)
	}
}

func packWindows(jarPath string, fileName string) {
	name := "build/" + fileName + ".jar"
	cmd := exec.Command(jarPath, "cvfm", name, "Manifest.txt", "ru")
	cmd.Dir = "jar"

	//cmd.Stdout = os.Stdout
	//cmd.Stderr = os.Stderr
	err := cmd.Run()

	if err != nil {
		fmt.Println(err)
	}
}

func packLinux() {
	fmt.Println("START run package script")
	cmd := exec.Command("/bin/sh", "create-FAT-jar.sh")

	//cmd.Stdout = os.Stdout
	//cmd.Stderr = os.Stderr
	err := cmd.Run()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("END package script")
}
