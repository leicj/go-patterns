package main

import "fmt"
import prototype "p7/prototype"

func main() {
	file1 := &prototype.File{Name: "File1"}
	file2 := &prototype.File{Name: "File2"}
	file3 := &prototype.File{Name: "File3"}

	folder1 := &prototype.Folder{
		Children: []prototype.Inode{file1},
		Name: "Folder1",
	}

	folder2 := &prototype.Folder{
		Children: []prototype.Inode{folder1, file2, file3},
		Name: "Folder2",
	}
	fmt.Println("\nPrinting hierarchy for Folder2")
    folder2.Print("  ")

    cloneFolder := folder2.Clone()
    fmt.Println("\nPrinting hierarchy for clone Folder")
    cloneFolder.Print("  ")
}