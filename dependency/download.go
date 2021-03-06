package download

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func writeStringToFile(filepath, s string) error {
	fo, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer fo.Close()

	_, err = io.Copy(fo, strings.NewReader(s))
	if err != nil {
		return err
	}

	return nil
}

type pomDependenciesStr struct {
	Locations xml.Name           `xml:"project"`
	Depen     []pomDependencyStr `xml:"dependencies>dependency"`
}

type pomDependencyStr struct {
	GroupId    string `xml:"groupId"`
	ArtifactId string `xml:"artifactId"`
	Version    string `xml:"version"`
	Scope      string `xml:"scope"`
}

func downloadFile(fileName string, url string) error {

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	oneFolder := "target"
	twoFolder := "dependency"
	//threeFolder := "lib"

	//dirPath := filepath.Join(pwd, oneFolder, twoFolder, threeFolder)
	dirPath := filepath.Join(pwd, oneFolder, twoFolder)

	fullPath := filepath.Join(dirPath, fileName)

	_, err = os.Stat(dirPath)
	if err != nil {
		//fmt.Println("create dir ...")
		os.MkdirAll(dirPath, os.ModePerm)
	}

	// Create the file
	out, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	size, err := io.Copy(out, resp.Body)
	//_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	if (size / 1024) > 0 {
		if (size / 1048576) > 0 {
			result := size / 1048576

			fmt.Printf("%s with %v Mbytes downloaded \n", fileName, result)
		} else {
			result := size / 1024
			fmt.Printf("%s with %v Kbytes downloaded \n", fileName, result)
		}
	} else {
		fmt.Printf("%s with %v bytes downloaded \n", fileName, size)
	}

	return nil
}

func FromPom(resourceUrl string, pomFilepath string) string {
	var fileSTR string = "Class-Path: "

	v := pomDependenciesStr{}

	if pomFilepath == "" {
		pomFilepath = "pom.xml"
	}

	raw_data, err := ioutil.ReadFile(pomFilepath)

	if err != nil {
		fmt.Printf("error: %v \n", err)
		os.Exit(1)
	}

	err = xml.Unmarshal(raw_data, &v)
	if err != nil {
		fmt.Printf("error: %v \n", err)
		os.Exit(1)
	}

	if resourceUrl == "" {
		resourceUrl = "https://repo1.maven.org/maven2/"
	}

	fmt.Println("Start downloading files:")

	for i := range v.Depen {

		if v.Depen[i].Scope == "" || v.Depen[i].Scope == "compile" {
			art := v.Depen[i].ArtifactId
			ver := v.Depen[i].Version

			fileName := art + "-" + ver + ".jar"

			newGroupID := strings.Replace(v.Depen[i].GroupId, ".", "/", -1)

			fileUrl := resourceUrl + newGroupID + "/" + art + "/" + ver + "/" + fileName

			fileSTR = fileSTR + "dependency/" + fileName + " "

			//	fmt.Println(fileUrl, fileName)

			err := downloadFile(fileName, fileUrl)
			if err != nil {
				panic(err)
			}

		}

	}

	/*
		fileName := "listLib.txt"
		oneFolder := "jar"
		fullPath := filepath.Join(oneFolder, fileName)
		writeStringToFile(fullPath, fileSTR)

	*/

	//fmt.Println("........ Done !")

	return fileSTR
}
