package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//var db *sql.DB

type Album struct {
	ID     int64   `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float32 `json:"price"`
}

// albumsByArtist queries for albums that have the specified artist name.
//
//	func albumsByArtist(name string) ([]Album, error) {
//		// An albums slice to hold data from returned rows.
//		var albums []Album
//
//		rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
//		if err != nil {
//			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
//		}
//		defer rows.Close()
//		// Loop through rows, using Scan to assign column data to struct fields.
//		for rows.Next() {
//			var alb Album
//			if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
//				return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
//			}
//			albums = append(albums, alb)
//		}
//		if err := rows.Err(); err != nil {
//			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
//		}
//		return albums, nil
//	}
func main() {

	dsn := "root:#Mdsunny75@tcp(127.0.0.1:3306)/record?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	var db *gorm.DB
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("got error:")
		fmt.Println(err)
		return
	}
	p := &Album{
		ID:     12,
		Title:  "a",
		Artist: "b",
		Price:  123,
	}
	_ = p
	err = db.Debug().Migrator().DropTable(&Album{})
	if err != nil {
		panic(fmt.Sprintf("Failed to drop table: %v", err))
	}
	//if err := db.AutoMigrate(&Album{}); err != nil {
	//	fmt.Println("got error:")
	//	fmt.Println(err)
	//	return
	//}
	//
	//db.Create(&p)
	////if err != nil {
	////	fmt.Println("got error:")
	////	fmt.Println(err)
	////	return
	////}
	//
	////x := db.Take(&Album{})
	//var x []Album
	//result := db.Find(&x)
	//if result.Error != nil {
	//	panic("Failed to retrieve data from the database")
	//}
	//fmt.Println("hello")
	//fmt.Println(x)

	//// Capture connection properties.
	//// Capture connection properties.
	////cfg := mysql.Config{
	////	User:   os.Getenv("DBUSER"),
	////	Passwd: os.Getenv("DBPASS"),
	////	Net:    "tcp",
	////	Addr:   "127.0.0.1:3306",
	////	DBName: "recordings",
	////}
	////cfg := mysql.Config{
	////	User:                 "root",
	////	Passwd:               "#Mdsunny75",
	////	Net:                  "tcp",
	////	Addr:                 "127.0.0.1:3306",
	////	DBName:               "recordings",
	////	AllowNativePasswords: true,
	////}
	////_ = cfg
	////a := cfg.FormatDSN()
	//b := "root:#Mdsunny75@tcp(127.0.0.1:3306)/recordings"
	////_ = a
	//_ = b
	//
	//fmt.Println("a: ", b)
	//fmt.Println("b: ", b)
	//
	////db, err = sql.Open("mysql", cfg.FormatDSN())
	//var err error
	////db, err = sql.Open("mysql", b)
	//
	//db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:4000)/test"), &gorm.Config{})
	//if err != nil {
	//	panic("failed to connect database")
	//}
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	////pingErr := db.
	////if pingErr != nil {
	////	log.Fatal(pingErr)
	////}
	//fmt.Println("Connected!")
	//
	//albums, err := albumsByArtist("John Coltrane")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//_ = albums
	//fmt.Printf("Albums found:\n")
	//for _, i := range albums {
	//	fmt.Println(i)
	//}
}
