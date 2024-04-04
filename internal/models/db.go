package models

import (
	"database/sql"
	"fmt"
	"log"
	"url-shorter/internal/config"

	_ "github.com/lib/pq"
)

var db *sql.DB

type Postgres struct {
	db *sql.DB
}

func (p *Postgres) Init() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Dbname)
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		db.Close()
		log.Fatalln(err)
	}
	p.db = db
}

func (p *Postgres) InsertLink(link string, shortenedLink string) {
	query := fmt.Sprintf("INSERT INTO Links(link, shortened_link) VALUES('%v', '%v')", link, shortenedLink)
	_, err := p.db.Exec(query)
	if err != nil {
		p.db.Close()
		log.Fatalln(err)
	}
}

func (p *Postgres) GetLink(shortenedLink string) (string, bool) {
	var link string
	query := fmt.Sprintf("SELECT link FROM Links WHERE shortened_link = '%v'", shortenedLink)
	row := p.db.QueryRow(query)
	err := row.Scan(&link)
	if err != nil {
		log.Println(err)
		return "", false
	}
	return link, true

}

func (p *Postgres) GetShortenedLink(link string) (string, bool) {
	var shortenedLink string
	query := fmt.Sprintf("SELECT shortened_link FROM Links WHERE link = '%v'", link)
	row := p.db.QueryRow(query)
	err := row.Scan(&shortenedLink)
	if err != nil {
		log.Println(err)
		return "", false
	}
	return shortenedLink, true
}
