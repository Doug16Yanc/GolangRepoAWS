package main

import (
	"fmt"
	"strings"
)

type Repository struct{
	idRepo int
	nameRepo string
	sizeRepo float64
}

func main(){
    fmt.Println("Welcome to our Amazon Web Services repository management system.\n")
    interactes()
}

func interactes(){
    fmt.Println("Choose an option:\n")
    fmt.Println("   C/c - Create repository\n" +
                "   U/u - Update repository\n" +
                "   S/s - Search repository\n"
                "   D/d - Delete repository\n")
    var option string
    fmt.Scanln(&option)
    
    switch strings.ToLower(option){
        case "c" :
        
        case "u" :
        
        case "s" :
        
        case "d" :
        
        default:
            fmt.Println("Option no-existent.\n")
    }
}