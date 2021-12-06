package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	helpInfo := flag.Bool("h", false, "Displays this help file")
	pathToRun := flag.String("path", "c:", `Path to check. Network example("\\client.local\sharename")`)
	extToFind := flag.String("ext",`.ps1,.bat,.sh,.py`, "Extentions to search for")
	flag.Parse()

	if *helpInfo == true{
		flag.Usage()
		os.Exit(2)
	}

	if *pathToRun == `c:` {
		searchthesefolder := getalloweddir()
		for _, x := range searchthesefolder {
			walkdir(x,*extToFind)
		}
	}else{
		walkdir(*pathToRun,*extToFind)
	}
}

func checkallowed(fileString string, extensions string) {
	extlist := strings.Split(extensions,",")
	for _, y := range extlist {
		if filepath.Ext(fileString) == y {
			fmt.Println(fileString)
			break
		}
	}

}

//Recusive search though passed folder
func walkdir(allowedFolder string,extentions string) {
	errO := filepath.Walk(allowedFolder, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		checkallowed(path,extentions)

		return nil
	})
	if errO != nil {
		log.Println(errO)
	}
}

// Select all directories in in c:\ that are not system folder. Just too many false positives.
func getalloweddir() []string {
	forbiddenString := `C:\ProgramData
C:\Program Files (x86)
C:\Program Files
C:\$WinREAgent
C:\$Windows.~WS
C:\$WINDOWS.~BT
C:\Windows"	
C:\Windows10Upgrade
C:\Documents and Settings
C:\Python27`
	var searchTheseFolders []string
	forbiddenDir := strings.Split(forbiddenString, "\n")
	rootDir, err := filepath.Glob("c:\\*")
	if err != nil {
		log.Fatal(err)
	}
	for _, x := range rootDir {
		isthisadir, errz := os.Stat(x)
		if errz != nil {
			log.Fatal(errz)
		}
		if isthisadir.IsDir() == false {
			continue
		}

		match := false
		for _, y := range forbiddenDir {
			if strings.ToLower(x) == strings.ToLower(y) {
				match = true
			}
		}
		if match == false {
			searchTheseFolders = append(searchTheseFolders, x)
		}

	}

/*	fmt.Println(rootDir)
	fmt.Println(forbiddenDir)
	fmt.Println(searchTheseFolders)*/
	return searchTheseFolders

}

