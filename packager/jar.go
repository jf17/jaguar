package jar

import (
	"fmt"
	//"os"
	"os/exec"
)

func Pack(jarPath string, fileName string) {
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
