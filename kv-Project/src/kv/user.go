package main

type User struct {
    Name  string
    Given string
    Age   string
}

func newUser(kv map[string]string)(User) {
    user := User{}
    user.Name = kv["name"]
    user.Given = kv["vorname"]
    user.Age = kv["alter"]

    return user
}