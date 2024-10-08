package main

import (
	"fmt"
	"net/http"
	"log"
	"flag"
	"database/sql"


	_ "github.com/go-sql-driver/mysql"

	"urls/urlshort"
)
type Dburls struct {
	ID int
	PATH string
	URL string
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	fmt.Printf("connected to DB with %s\n", dsn)
	return db, nil
}

func getDBRows(db *sql.DB) ([]*Dburls, error) {
	stmt := `SELECT ID, PATH, URL FROM urls`
	rows, err := db.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	dbdata := []*Dburls{}
	for rows.Next(){
		s := &Dburls{}
		err := rows.Scan(&s.ID, &s.PATH, &s.URL)
		if err != nil {
			return nil, err
		}
		dbdata = append(dbdata, s)
	}
	if err = rows.Err(); err != nil{
		return nil, err
	}
	return dbdata, nil

}

func buildPathMap(sqldata []*Dburls) map[string]string {
	pathsToUrls := make(map[string]string)
	for _, entry := range sqldata {
		pathsToUrls[entry.PATH] = entry.URL
	}
	return pathsToUrls
}


func main() {
	db_conn := "web:pass@/urlshort?parseTime=true"
	db, err := openDB(db_conn)

	defer db.Close()

	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}

	yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
- path: /anek
  url: https://www.anekdot.ru/
`
	json := `[{"path": "/urlshort", "url": "https://github.com/gophercises/urlshort"}, {"path": "/urlshort-final", "url": "https://github.com/gophercises/urlshort/tree/solution"}, {"path": "/anek", "url": "https://www.anekdot.ru/"}]`

	useYaml := flag.Bool("isyaml", false, "use yaml handler")
	useJson := flag.Bool("isjson", false, "use json handler")
	useSql := flag.Bool("issql", false, "use sql handler")
	flag.Parse()

	mux := defaultMux()
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	var aHandler http.Handler
	sqldata, err := getDBRows(db)

	switch {
	case *useYaml:
		fmt.Println("using yaml source...")
		aHandler, err = urlshort.YAMLHandler([]byte(yaml), mapHandler)
	case *useJson:
		fmt.Println("using json source...")
		aHandler, err = urlshort.JSONHandler([]byte(json), mapHandler)
	case *useSql:
		fmt.Printf("Type of sqldata: %T\n", sqldata)
		
		 fmt.Println("using sql source...")
		 parsedData := buildPathMap(sqldata)
		 fmt.Printf("Type of parsedData: %T\n", parsedData)
		 fmt.Println("parsedData: ", parsedData)
		 for path, url := range parsedData {
			fmt.Printf("Path: %s, URL: %s\n", path, url)
		}
		aHandler =urlshort.MapHandler(parsedData, mapHandler)
	default:
		fmt.Println("Chief everything was justlost, using default")
	}

	if err != nil {
		panic(err)
	}

	fmt.Println("start listening on port 8080...")
	err = http.ListenAndServe(":8080", aHandler)
	log.Fatal(err)

}
