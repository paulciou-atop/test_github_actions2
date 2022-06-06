package db

import (
	"database/sql"
	"fmt"
	"github.com/go-redis/redis"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

type ConntDbInfo struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
}

func init() {

	db := Connet2Postgre(ConntDbInfo{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		User:     "postgres",
		Password: "postgres",
		Dbname:   "postgres",
	})
	db.Query("DROP TABLE IF EXISTS albums")

	db.Query(`
	CREATE TABLE album (
		id         SERIAL PRIMARY KEY,
		title      VARCHAR(128) NOT NULL,
		artist     VARCHAR(255) NOT NULL,
		price      DECIMAL(5,2) NOT NULL
	  )
	`)

	//albums, err := albumsByArtist("John Coltrane")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("Albums found: %v\n", albums)
	//
	//// Hard-code ID 2 here to test the query.
	//alb, err := albumByID(2)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("Album found: %v\n", alb)

}

func Connet2Postgre(dbInfo ConntDbInfo) *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbInfo.Host, dbInfo.Port, dbInfo.User, dbInfo.Password, dbInfo.Dbname)

	client, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	//defer db.Close()
	pingErr := client.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	return client
}

func QueryAlbumsByArtist(name string, client *sql.DB) ([]Album, error) {
	// An albums slice to hold data from returned rows.
	var albums []Album

	rows, err := client.Query("SELECT * FROM album WHERE artist = $1", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist scan %q: %v", name, err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist  rows error %q: %v", name, err)
	}
	return albums, nil
}

func InsertPostgre(alb Album, client *sql.DB) (int64, error) {
	var id int64
	err := client.QueryRow("INSERT INTO album (title, artist, price) VALUES ($1, $2, $3) RETURNING id", alb.Title, alb.Artist, alb.Price).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}

	return id, nil
}

// redis

func ConnectRedis(dbInfo ConntDbInfo) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     dbInfo.Host + ":" + dbInfo.Port,
		Password: "",
		DB:       0,
	})
	if err := client.Ping().Err(); err != nil {
		log.Fatal(err)
	}
	return client
}

func InsertRedis(key string, value string, client *redis.Client) {
	err := client.Set(key, value, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("insert finished with : ", key, value)
}

func QueryRedis(key string, client *redis.Client) string {
	response, err := client.Get(key).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("query result: ", response)
	return response
}
