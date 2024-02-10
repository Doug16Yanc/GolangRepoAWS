package main

import (
	"container/list"
	"fmt"
	"math/rand"
	"strings"
	"os"
)

type Repository struct {
	idRepo   int
	nameRepo string
	sizeRepo float64
}

func main() {
	listRepo := list.New()
	fmt.Println("Welcome to our Amazon Web Services repository management system.\n")
	interactes(listRepo)
}

func interactes(listRepo *list.List) {
	for {
	    fmt.Println("Choose an option:")
    	fmt.Println("C/c - Create repository")
    	fmt.Println("S/s - Query repository")
    	fmt.Println("U/u - Update repository")
    	fmt.Println("D/d - Delete repository")
    	fmt.Println("Q/q - Quit system")
    	var option string
    	fmt.Scanln(&option)
    
    	switch strings.ToLower(option) {
    	case "c":
    		createRepository(listRepo)
    	case "s":
    		queryRepository(listRepo)
    	case "u":
    		updateRepository(listRepo)
    	case "d":
    		deleteRepository(listRepo)
    	case "q" :
    	    printMessage("Losing you is such a pity, I will miss you, farewell!\n")
    	    os.Exit(0)
    	default:
    		fmt.Println("Option no-existent.\n")
    	}
	}
}

func generateId(listRepo *list.List) int {
	var id int
	for {
		id = rand.Intn(10000) + 1000000000
		duplicate := false
		for e := listRepo.Front(); e != nil; e = e.Next() {
			repo := e.Value.(Repository)
			if repo.idRepo == id {
				duplicate = true
				break
			}
		}
		if !duplicate {
			break
		}
	}
	return id
}

func printMessage(message string) {
	fmt.Println("***********************************")
	fmt.Println(message)
	fmt.Println("***********************************")
}

func createRepository(listRepo *list.List) {
	id := generateId(listRepo)

	fmt.Println("\nName of your repository:")
	var name string
	fmt.Scanln(&name)

	size := 20.00

	repo := Repository{
		idRepo:   id,
		nameRepo: name,
		sizeRepo: size,
	}

	listRepo.PushBack(repo)

	proofCreation(repo)
}

func proofCreation(repo Repository) {
	printMessage("\n      PROOF REPOSITORY CREATION       " +
		"\n       Id : " + fmt.Sprintf("%d", repo.idRepo) + 
		"\n       Name : " + fmt.Sprintf("%s", repo.nameRepo) +
		"\n       Size by Amazon server : " + fmt.Sprintf("%.2f KB", repo.sizeRepo),
		)
}

func searchRepository(listRepo *list.List, id int) (*Repository, error) {
	for e := listRepo.Front(); e != nil; e = e.Next() {
		repo := e.Value.(Repository)
		if repo.idRepo == id {
			return &repo, nil
		}
	}
	return nil, fmt.Errorf("Repository AWS not found")
}

func queryRepository(listRepo *list.List) {
	fmt.Println("Enter a repository ID:")
	var id int
	fmt.Scanln(&id)

	repo, err := searchRepository(listRepo, id)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Repository found: ID=%d, Name=%s\n", repo.idRepo, repo.nameRepo)
}

func updateRepository(listRepo *list.List) {
	fmt.Println("Enter a repository ID to update:")
	var id int
	fmt.Scanln(&id)

	repo, err := searchRepository(listRepo, id)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Choose an option to update:")
	fmt.Println("N/n - Name")
	fmt.Println("S/s - Size")

	var option string
	fmt.Scanln(&option)

	switch strings.ToLower(option) {
	case "n":
		fmt.Println("New name:")
		var newName string
		fmt.Scanln(&newName)

		repo.nameRepo = newName
		fmt.Printf("Repository %d updated successfully with name %s.\n", repo.idRepo, repo.nameRepo)
	case "s":
		fmt.Println("New size request:")
		var newSize float64
		fmt.Scanln(&newSize)

		repo.sizeRepo = newSize
		fmt.Printf("Repository %d updated successfully with %.2f KB.\n", repo.idRepo, repo.sizeRepo)
	default:
		fmt.Println("Option nonexistent.\n")
	}
}

func deleteRepository(listRepo *list.List) {
}

