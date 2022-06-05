package db_test

import (
	"fmt"
	"learn_git_actions/db"

	"log"
	"testing"
)

func TestInsert_PostgreSql(t *testing.T) {
	client := db.Connet2Postgre(db.ConntDbInfo{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "postgres",
		Dbname:   "postgres",
	})
	albID, err := db.InsertPostgre(db.Album{
		Title:  "116",
		Artist: "aaa",
		Price:  49.99,
	}, client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of added album: %v\n", albID)

}

func TestQueryAlbumsByArtist(t *testing.T) {
	client := db.Connet2Postgre(db.ConntDbInfo{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "postgres",
		Dbname:   "postgres",
	})
	albums, err := db.QueryAlbumsByArtist("aaa", client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", albums)

}
