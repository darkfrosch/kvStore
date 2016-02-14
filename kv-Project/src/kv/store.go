package main

import (
	"bufio"
    "fmt"
    "os"
    "strings"
)

type Store struct {
	Operation string
}

func (kv Store)PrintAttributes(all bool)() {

    if(all) {
        file, err := os.Open("resources/User")
        if err != nil {
            fmt.Printf("error: %v", err)
            return
        }
        defer file.Close()
        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
          fmt.Printf("%v \n", scanner.Text())
        }    
    } else {
        for i := range os.Args {
            if (i == 0) {
                continue
            }
            if (os.Args[i] != "name" && os.Args[i] != "vorname" && os.Args[i] != "alter") {
                fmt.Printf("Ungültige Eingabe")
            } else {
                file, err := os.Open("resources/User")
                if err != nil {
                    fmt.Printf("error: %v", err)
                    return
                }
                defer file.Close()
                scanner := bufio.NewScanner(file)
                for scanner.Scan() {

                    kvTmp := strings.SplitN(scanner.Text(), "=", 2)     

                    if (kvTmp[0] == os.Args[i]) {
                        fmt.Printf("%v \n", scanner.Text()) 
                    }    
                }       
            }
        }
    }       
} 

func saveUser(user User)(){
    file, err := os.Create("resources/User")
    if err != nil {
        fmt.Printf("error: %v", err)
        return
    }
    defer file.Close()
    fmt.Fprintf(file, "name=%v \nvorname=%v \nalter=%v", user.Name, user.Given, user.Age)
}