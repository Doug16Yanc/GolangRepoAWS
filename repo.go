package main

import (
	"fmt"
	"strings"
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

func createRepo(listRepo *list.List){
    
}

func searchRepository(listRepo *list.List){
    
}

func updateRepository(listRepo *list.List){
    
}

func deleteRepository(listRepo *list.List){
    
}

