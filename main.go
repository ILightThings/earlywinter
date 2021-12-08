package main

import (
	"bufio"
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
	extToFind := flag.String("ext", `.ps1,.bat,.sh,.py,.rdp`, "Extentions to search for")
	searchFile := flag.Bool("search", false, "Search file for strings that you would typically find with hard coded creds ")
	flag.Parse()

	if *helpInfo == true {
		flag.Usage()
		os.Exit(2)
	}

	if *pathToRun == `c:` {
		searchthesefolder := getalloweddir()
		for _, x := range searchthesefolder {
			walkdir(x, *extToFind, *searchFile)
		}
	} else {
		walkdir(*pathToRun, *extToFind, *searchFile)
	}
}

func checkallowed(fileString string, extensions string, searchbool bool) {
	extlist := strings.Split(extensions, ",")
	for _, y := range extlist {
		if filepath.Ext(fileString) == y {
			fmt.Println(fileString)
			if searchbool == true {
				searchData(fileString)
			}
			break
		}
	}

}

//Recusive search though passed folder
func walkdir(allowedFolder string, extentions string, searchbool bool) {
	errO := filepath.Walk(allowedFolder, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		checkallowed(path, extentions, searchbool)

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
C:\Windows
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

func searchData(fileToSearch string) int {
	file, err := os.Open(fileToSearch)
	if err != nil {
		println(err)
		return 1
	}
	defer file.Close()
	sensitivesStrings := `/u:
/p:
psexec
net user
net use
net account
/runas`
	stringArray := strings.Split(sensitivesStrings, "\n")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for _, x := range stringArray {
			if strings.Contains(scanner.Text(), x) {
				fmt.Printf("FOUND: %s : %s \t (%s)\n", fileToSearch,x, scanner.Text())
			}
		}
	}
	return 0
}
