package db_test

import (
	"fmt"
	"learn_git_actions/db"
	"os"

	"log"
	"testing"
)

func TestInsert_PostgreSql(t *testing.T) {
	client := db.Connet2Postgre(db.ConntDbInfo{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		User:     "postgres",
		Password: "postgres",
		Dbname:   "postgres",
	})
	albID, err := db.InsertPostgre(db.Album{
		Title:  "11133",
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
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
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

//redis

func TestConnectRedis(t *testing.T) {
	_ = db.ConnectRedis(db.ConntDbInfo{
		Host: "redis",
		Port: "6379",
	})
}

func TestInsertRedis(t *testing.T) {
	client := db.ConnectRedis(db.ConntDbInfo{
		Host: "redis",
		Port: "6379",
	})
	db.InsertRedis("abc", "123", client)
}

func TestQueryRedis(t *testing.T) {
	client := db.ConnectRedis(db.ConntDbInfo{
		Host: "redis",
		Port: "6379",
	})
	db.QueryRedis("abc", client)
}
