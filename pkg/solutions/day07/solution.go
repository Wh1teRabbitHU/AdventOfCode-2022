package day07

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"utils"

	"golang.org/x/exp/maps"
)

const inputPath string = "../inputs/day07.txt"
const totalDiskCapacity int = 70000000
const neededSpace int = 30000000

type FsFolder struct {
	name    string
	parent  *FsFolder
	folders map[string]*FsFolder
	files   map[string]*FsFile
}

func (inst *FsFolder) size() int {
	sizeSum := 0

	for _, file := range inst.files {
		sizeSum += file.size
	}

	for _, folder := range inst.folders {
		sizeSum += folder.size()
	}

	return sizeSum
}

func (inst *FsFolder) subFolders(recursive bool) []*FsFolder {
	folders := make([]*FsFolder, 0)

	for _, f := range inst.folders {
		folders = append(folders, f)

		if recursive {
			folders = append(folders, f.subFolders(true)...)
		}
	}

	return folders
}

type FsFile struct {
	name string
	size int
}

func createFolder(name string) FsFolder {
	return FsFolder{name: name, files: make(map[string]*FsFile, 0), folders: make(map[string]*FsFolder, 0)}
}

func createAndAddFolder(parentFolder *FsFolder, name string) *FsFolder {
	folder := createFolder(name)
	folder.parent = parentFolder

	parentFolder.folders[name] = &folder

	return &folder
}

func arrayContains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func interpretCommand(fileScanner *bufio.Scanner, rootFolder *FsFolder, currentFolder *FsFolder, cmd string) (*FsFolder, string, bool) {
	if !strings.HasPrefix(cmd, "$") {
		panic("Not a command!")
	}

	if strings.Contains(cmd, "$ cd") {
		target := cmd[5:]

		if target == ".." {
			currentFolder = currentFolder.parent
		} else if target == "/" {
			currentFolder = rootFolder
		} else if arrayContains(maps.Keys(currentFolder.folders), target) {
			currentFolder = currentFolder.folders[target]
		} else {
			currentFolder = createAndAddFolder(currentFolder, target)
		}
	} else if strings.Contains(cmd, "$ ls") {
		for fileScanner.Scan() {
			element := fileScanner.Text()

			if strings.HasPrefix(element, "$") {
				return currentFolder, element, true
			}

			parts := strings.Split(element, " ")
			itemName := parts[1]

			if parts[0] == "dir" { // Folder
				createAndAddFolder(currentFolder, itemName)
			} else { // File
				fileSize, err := strconv.Atoi(parts[0])

				utils.HandleError(err)

				currentFolder.files[itemName] = &FsFile{name: itemName, size: fileSize}
			}
		}
	}

	return currentFolder, "", false
}

func smallFolderSizeSum(rootFolder *FsFolder) int {
	sizeSum := 0

	for _, f := range rootFolder.subFolders(true) {
		folderSize := f.size()

		if folderSize < 100000 {
			sizeSum += folderSize
		}
	}

	return sizeSum
}

func findDeletableFolderSize(rootFolder *FsFolder) int {
	minDeletableSize := rootFolder.size() - (totalDiskCapacity - neededSpace)
	size := math.MaxInt64

	for _, f := range rootFolder.subFolders(true) {
		folderSize := f.size()

		if folderSize > minDeletableSize && folderSize < size {
			size = folderSize
		}
	}

	return size
}

func Solve() {
	inputFile, err := os.Open(inputPath)

	utils.HandleError(err)
	defer inputFile.Close()

	fileScanner := bufio.NewScanner(inputFile)
	fileScanner.Split(bufio.ScanLines)

	rootFolder := createFolder("/")
	currentFolder := &rootFolder

	for fileScanner.Scan() {
		cmd := fileScanner.Text()
		interpret := true

		// In case it reads into the next command, it should try to re-interpret without scanning
		// This only happens after ls, because the end of the output is only known by finding the next command
		for interpret {
			currentFolder, cmd, interpret = interpretCommand(fileScanner, &rootFolder, currentFolder, cmd)
		}
	}

	folderSizeSum := smallFolderSizeSum(&rootFolder)
	deletableFolderSize := findDeletableFolderSize(&rootFolder)

	fmt.Print("Day 07 - Solution 01: ")
	fmt.Println(folderSizeSum)
	fmt.Print("Day 07 - Solution 02: ")
	fmt.Println(deletableFolderSize)
}
