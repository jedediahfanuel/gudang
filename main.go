package main

import (
	"log"
	"net/http"

	"fmt"

	"gudang/database"
	"gudang/models"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	db := database.GetDBConnection()
	database.Migrate(db)

	user := models.User{
		Name:     "Jedi",
		Password: "jedi1234",
		Level:    1,
	}
	db.Create(&user)

	wh := models.Warehouse{
		Location: "Bandung Selatan",
	}
	db.Create(&wh)
	wh2 := models.Warehouse{
		Location: "Jakarta",
	}
	db.Create(&wh2)

	p1 := models.Product{
		Name: "BF50ml",
	}
	db.Create(&p1)
	p2 := models.Product{
		Name: "BF300ml",
	}
	db.Create(&p2)

	tr18 := models.History{
		Product:   p1,
		Warehouse: wh,
		Event:     1000,
	}
	db.Create(&tr18)

	tr18 = models.History{
		Product:   p1,
		Warehouse: wh,
		Event:     -100,
	}
	db.Create(&tr18)

	tr18 = models.History{
		Product:   p2,
		Warehouse: wh2,
		Event:     2000,
	}
	db.Create(&tr18)

	tr18 = models.History{
		Product:   p2,
		Warehouse: wh2,
		Event:     -500,
	}
	db.Create(&tr18)

	tr18 = models.History{
		Product:   p2,
		Warehouse: wh,
		Event:     200,
	}
	db.Create(&tr18)

	tr18 = models.History{
		Product:   p2,
		Warehouse: wh,
		Event:     -100,
	}
	db.Create(&tr18)

	tr18 = models.History{
		Product:   p1,
		Warehouse: wh2,
		Event:     500,
	}
	db.Create(&tr18)

	// r := mux.NewRouter()
	// routes.RegisterRoutes(db, r)

	// http.Handle("/", r)
	fmt.Println("Connected to port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
