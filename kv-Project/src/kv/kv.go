package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {

    store := &Store{}
    // Ausgabe aller Werte
    if len(os.Args) == 1 {
        store.PrintAttributes(true)
    }
    isValid, operation := validateInput()
    if (isValid) {
        switch operation {
            case "set":

                kv := prepareDataToSet()
                if (kv == nil) {
                   fmt.Printf("Ungültige Eingabe") 
                }
                user := newUser(kv)
                saveUser(user)

            case "get":

                store.PrintAttributes(false)
                
        }
    } else {
        fmt.Printf("Ungültige Eingabe")
    } 

}

func prepareDataToSet()(map[string]string) {
    kv := map[string]string{}

    for i := range os.Args {
        if (i == 0) {
            continue
        }
        kvTmp := strings.SplitN(os.Args[i], "=", 2)
        if (kvTmp[0] != "name" && kvTmp[0] != "vorname" && kvTmp[0] != "alter") {
            return nil
        } else {
            kv[kvTmp[0]] = kvTmp[1] 
        }
    }
    return kv            
}



func validateInput() (bool, string){
    var operation string 
    for i := range os.Args {
        if (i == 0) {
            continue
        }
        if strings.Contains(os.Args[i], "=") {
            if ( operation == "get") {
                return false, ""
            }
            operation = "set"
        } else {
            if ( operation == "set") {
                return false, ""
            }
            operation = "get"
        }
    }
    
    return true, operation
}



