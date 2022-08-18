package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)



var db *sql.DB

type Album struct {
	ID		int64
	Title   string
	Artist  string
	Price   float64
}

func main(){
	// handling database connection
	config := mysql.Config{
        User:   "root",
        Passwd: "cyber",
		Net : 	 "tcp" ,
		Addr:    "127.0.0.1:3306",
        DBName:  "recordings",
	}
    // Get a database handle.
	var err error 
	db , err = sql.Open("mysql" , config.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Connected!")

    albums, err := albumsByArtist("john cortrane")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Albums found: %v\n", albums)

	alb , err := findById(2);
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Single album found: %v\n", alb)

	albID, err := addAlbum(Album{
		Title:  "The Modern Sound of Betty Carter",
		Artist: "Betty Carter",
		Price:  49.99,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of added album: %v\n", albID)

	stmtId , err := preparedStatementAlbumByID(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Prepared statemnt query found : %v\n", stmtId)
}

func albumsByArtist(name string) ([]Album , error){
	var albums []Album

	rows , err := db.Query("SELECT * FROM album WHERE artist = ?" , name);
	if err != nil {
		return nil , fmt.Errorf("albumsByArtist %q: %v", name, err)
	}

	defer rows.Close()

	//loop over document
	for rows.Next(){
		var alb Album
		if err := rows.Scan(&alb.ID , &alb.Title , &alb.Artist , &alb.Price); err != nil {
			return nil , fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums , alb)
	}

	//final check for error than might have happended in the quary
	if err:= rows.Err() ; err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return albums , nil
}

func findById(id int64) (Album , error){
	var alb Album
	row := db.QueryRow("SELECT * FROM album WHERE id = ?" , id);
	
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price) ; err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumsById %d: no such album", id)
		}
		return alb, fmt.Errorf("albumsById %d: %v", id, err)
	}

	return alb , nil ;
}

func addAlbum(alb Album) (int64 , error){
	result , err := db.Exec("INSERT INTO album (title , artist , price) VALUES (? , ? , ?)" , alb.Title  , alb.Artist , alb.Price);
	if err != nil {
		return 0 , fmt.Errorf("addAlbum: %v", err)
	}
	id , err := result.LastInsertId()
	if err !=nil {
		return 0 , fmt.Errorf("addAlbum: %v", err)
	}
	return id , nil
}

func preparedStatementAlbumByID(id int64) (Album , error) {
	stmt , err := db.Prepare("SELECT * FROM album WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	
	var album Album ;
	err = stmt.QueryRow(id).Scan(&album.ID, &album.Title, &album.Artist, &album.Price);

	stmt.Close();

	if err != nil { 
		if err == sql.ErrNoRows {
			return album , err
		}
		return album , err ;
	}

	return album , err ;
}