package book

import (
	"CRUDTest/internal/handlers"
	"database/sql"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io"
	"log"
	"net/http"
)

type handler struct {
	db *sql.DB
}

func NewHandler(db *sql.DB) handlers.Handler {
	return &handler{
		db: db,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET("/book/:uuid", h.GetBook)
	router.GET("/books", h.GetBooks)
	router.POST("/books", h.CreateBook)
	router.PUT("/book/:uuid", h.UpdateBook)
	router.DELETE("/book/:uuid", h.DeleteBook)
}

func (h *handler) GetBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	id := ps.ByName("uuid")
	var books Book

	rows, err := h.db.Query("SELECT * FROM books where id = $1", id)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		b := Book{}
		err := rows.Scan(&b.Id, &b.Name, &b.Price, &b.Time)
		if err != nil {
			log.Fatal(err)
		}

		books = b
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}

func (h *handler) GetBooks(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var books []Book

	rows, err := h.db.Query("SELECT * FROM books")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		b := Book{}
		err := rows.Scan(&b.Id, &b.Name, &b.Price, &b.Time)
		if err != nil {
			log.Fatal(err)
		}

		books = append(books, b)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(books)
}

func (h *handler) CreateBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var input BookInput
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	defer r.Body.Close()

	if err := json.Unmarshal(body, &input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
	}

	_, err = h.db.Exec("INSERT into books (name, price) VALUES ($1, $2)", input.Name, input.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) UpdateBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var input BookInput
	var id = ps.ByName("uuid")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}

	defer r.Body.Close()

	if err = json.Unmarshal(body, &input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
	}

	_, err = h.db.Exec("update books set name=$1, price=$2 where id=$3", input.Name, input.Price, id)
	w.WriteHeader(http.StatusOK)

}

func (h *handler) DeleteBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("uuid")

	_, err := h.db.Exec("delete from books where id=$1", id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
	}

	w.WriteHeader(http.StatusOK)
}
