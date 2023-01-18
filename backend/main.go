package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Lubpui/Musian/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ListSong struct {
	Title       string `json:"title"`
	Album       string `json:"album"`
	Produced    string `json:"produced"`
	ReleaseDate string `json:"releaseDate"`
	Link        string `json:"link"`
}

func getByID(c echo.Context) error {
	listSong := ListSong{}
	id := c.Param("id")

	row := db.QueryRow(`SELECT title, album, produced, releaseDate, link 
	FROM Yoasobi_Song WHERE id=?`, id)

	err := row.Scan(&listSong.Title, &listSong.Album, &listSong.Produced, &listSong.ReleaseDate, &listSong.Link)
	switch err {
	case nil:
		return c.JSON(http.StatusOK, listSong)
	case sql.ErrNoRows:
		return c.JSON(http.StatusNotFound, map[string]string{"message": "not found"})
	default:
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
}

var db *sql.DB

func main() {
	db = config.Connect_Database()

	e := echo.New()
	port := "8080"

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
	}))

	e.GET("/get/:id", getByID)

	log.Println("Starting Server... Port:", port)
	log.Fatal(e.Start(":" + port))
}
