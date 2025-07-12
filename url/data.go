package url

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "postgres"
)

var db *sql.DB

func init() {
	sqlcon := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, _ = sql.Open("postgres", sqlcon)
	db.Ping()
}
func increasecount(s int) {
	var id int
	r1 := db.QueryRow("select hits from url where id=$1", s)
	r1.Scan(&id)
	id++
	db.Exec("update url set hits=$1 where id=$2", id, s)
}
func addUrl(a, b string) (UrlData, error) {
	var dt UrlData
	r1 := db.QueryRow("select * from url where original=$1", b)
	er := r1.Scan(&dt.Id, &dt.Short, &dt.Url, &dt.Create, &dt.Update, &dt.hits)
	if er != sql.ErrNoRows {
		return dt, nil
	}
	qry := `insert into url(short,original) values($1,$2)`
	_, er = db.Exec(qry, a, b)
	if er != nil {
		fmt.Println(er)
		return dt, errors.New("user already exists")
	}
	row := db.QueryRow("select * from url where id=(select max(id) from url)")
	row.Scan(&dt.Id, &dt.Short, &dt.Url, &dt.Create, &dt.Update, &dt.hits)
	return dt, nil

}

func retrieveUrl(a string) (UrlData, error) {
	var dt UrlData
	r1 := db.QueryRow("select * from url where short=$1", a)
	er := r1.Scan(&dt.Id, &dt.Short, &dt.Url, &dt.Create, &dt.Update, &dt.hits)
	if er != sql.ErrNoRows {
		increasecount(dt.Id)
		return dt, nil
	} else {
		return dt, errors.New("not found")
	}
}

func updateurl(a, b string) (UrlData, error) {
	var dt UrlData
	tim := time.Now().Format(time.RFC3339)
	r, _ := db.Exec("update url set original=$1,updated=$2,hits=$3 where short=$4", b, tim, 0, a)
	ct, _ := r.RowsAffected()
	if ct == 0 {
		return dt, errors.New("url not found")
	} else {
		r1 := db.QueryRow("select * from url where short=$1", a)
		r1.Scan(&dt.Id, &dt.Short, &dt.Url, &dt.Create, &dt.Update, &dt.hits)
		return dt, nil
	}
}

func deleteurl(a string) error {
	r, _ := db.Exec("delete from url where short=$1", a)
	ct, _ := r.RowsAffected()
	if ct == 0 {
		return errors.New("url not found")
	} else {
		return nil
	}
}
