package main

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
    "net/http"
    "test/api"
    "test/serverweb"
)

func main() {
    http.HandleFunc("/hello", serverweb.Accueil)
    http.HandleFunc("/create", serverweb.Create)
    http.HandleFunc("/createbypost", serverweb.CreateByPost)
    http.HandleFunc("/list", serverweb.ListTodo)
    http.HandleFunc("/get", serverweb.GetTodo)
    http.HandleFunc("/del", serverweb.Delete)
    http.HandleFunc("/update", serverweb.Update)
    http.HandleFunc("/health", serverweb.Health)


    const (
        host     = "172.17.0.1"
        port     = 5432
        user     = "user1"
        password = "motdepasse"
        dbname   = "dbtodo"
    )

    psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
    println(psqlconn)

    // open database
    db, err := sql.Open("postgres", psqlconn)
    if err != nil {
        fmt.Println("err"+ err.Error())
    }

    api.DataBasePtr = db

    // close database
    defer db.Close()

    erreur := http.ListenAndServe(":8090", nil)
    println(erreur.Error())
}


