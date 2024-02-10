package main

import (
	"fmt"
	"strings"
	"math/rand"
	"container/list"
)

type Repository struct{
	idRepo int
	nameRepo string
	sizeRepo float64
}

func main(){
    listRepo := list.New()
    fmt.Println("Welcome to our Amazon Web Services repository management system.\n")
    interactes(listRepo)
}

func interactes(listRepo *list.List){
    fmt.Println("   Choose an option:   \n")
    fmt.Println("   C/c - Create repository\n" +
                "   S/s - Search repository\n" +                
                "   U/u - Update repository\n" +
                "   D/d - Delete repository\n")
    var option string
    fmt.Scanln(&option)
    
    switch strings.ToLower(option){
        case "c" :
            createRepo(listRepo)
        case "s" :
            searchRepository(listRepo)
        case "u" :
            updateRepository(listRepo)
        case "d" :
            deleteRepository(listRepo)
        default:
            fmt.Println("Option no-existent.\n")
    }
}

func generateId(listRepo *list.List) int {
    var enter int
    var id int
    for {
        enter = rand.Intn(10000) + 1000000000
        for e := listRepo.Front(); e != nil; e = e.Next() {
    		if e.Value.(int) == id {
    			return 1
    			break
    		}
        }
    }
    listRepo.PushBack(enter)
    return enter
}

func printMessage(message string) string{
    fmt.Println("***********************************\n")
    fmt.Println(message)
    fmt.Println("***********************************\n")
    return message
}

func createRepo(listRepo *list.List){
    id := generateId(listRepo)
    
    fmt.Println("Name of your repository:")
    var name string
    fmt.Scanln(&name)
    
    size := 20.00
    
    repo := Repository {
        idRepo : id,
        nameRepo : name,
        sizeRepo : size,
    }
    
    listRepo.PushBack(repo)
    
    proofCreation(listRepo);
}

func proofCreation(listRepo *list.List){
    
}

func searchRepository(listRepo *list.List){
    
}

func updateRepository(listRepo *list.List){
    
}

func deleteRepository(listRepo *list.List){
    
}
