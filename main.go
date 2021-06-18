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
 /*   fmt.Println("Hello Arnaud")
   a := 4
   b := 12

    a += b
    fmt.Println("a += b  = ", a)


    var x int = 50
    var y float32 = 30.5

    fmt.Printf("x + y = ", float32(x )+ y)

    var todo1 api.Todo
    SetTodo(todo1, "Salut les copains 1")
    println("titre", todo1.Titre)

    var todo2 api.Todo
    SetTodoPtr(&todo2, "Salut les copains 2")
    println("titre", todo2.Titre)

    //api.Todos

   scanner := bufio.NewScanner(os.Stdin) // création du scanner capturant une entrée utilisateur
    fmt.Print("Entrez quelque chose : ")
    scanner.Scan()                      // lancement du scanner
    entreeUtilisateur := scanner.Text() // stockage du résultat du scanner dans une variable
    fmt.Println("Resultat de la saisie utilisateur : " + entreeUtilisateur)


    fmt.Print("Entrez un nombre entier : ")
    scanner.Scan()

    nbr, err := strconv.Atoi(scanner.Text()) // conversion du type string en int

    if err!= nil {
        fmt.Printf("Erreur retournée %s res : %d\n", err.Error(), (nbr + 6))
    } else {
        fmt.Printf("res : %d\n", (nbr + 6))
    }
*/

    api.Create("Apprentissage", "Programmation", 161220)
    api.Create("Cuisine", "Cookies", 171220)
    api.Create("Golang", "Cours", 270121)
    api.Create("Javascript", "Vue", 280121)
    api.Create("TypeScript", "Typer les éléments", 230121)
    api.Create("Immobilier", "FI", 230121)
    api.Create("Neuf", "FI9", 230121)
    api.Create("Sport", "Faire un footing", 241221)
    api.Create("Sieste", "Se reposer", 301121)

    for _, valeur := range api.List() {
        println(" ", valeur.Id, " ", valeur.Titre, " ", valeur.Description, " ", valeur.DueDate)
    }


    http.HandleFunc("/hello", serverweb.Accueil)
    http.HandleFunc("/create", serverweb.Create)
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
    CheckError(err)

    api.DataBasePtr = db

    // close database
    defer db.Close()

    // check db
    err = db.Ping()
    CheckError(err)

    fmt.Println("Connected!")
    RecordTest(db)

    erreur := http.ListenAndServe(":8090", nil)
    println(erreur.Error())
}

func CheckError(err error) {
    if err != nil {
      //  panic(err)
    }
}

func RecordTest(myDb *sql.DB) {
    sqlStatement := `INSERT INTO "user" ("email", "password") VALUES ($1, $2) RETURNING id`
    id := 0
    err := myDb.QueryRow(sqlStatement, "test@test.com", "arnaud").Scan(&id)
    println("id :", id)
   // _, err := myDb.Exec(sqlStatement)
    if err != nil {
        panic(err)
    }

}





/*func SetTodo (todo api.Todo, titre string) {
    todo.Titre = titre
    println("titre", todo.Titre)
}

func SetTodoPtr (todo *api.Todo, titre string) {
todo.Titre = titre
println("titre", todo.Titre)
}*/



