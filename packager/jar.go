package jar

import (
	"fmt"
	"os"
	"os/exec"
)

func Pack(jarPath string) {
	cmd := exec.Command(jarPath, "cvfm", "build/MyIDE.jar", "Manifest.txt", "ru")
	cmd.Dir = "jar"

	//cmd.Stdout = os.Stdout
	//cmd.Stderr = os.Stderr
	err := cmd.Run()

	if err != nil {
		fmt.Println(err)
	}
}
