package handler

import (
	"bookshelf-api-pemula-dicoding/model"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

var books = []model.Book{}

func PostBookHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	book := model.BookRequest{}

	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&book); err != nil {
		panic(err)
	}

	if book.Name == "" {
		writer.Header().Add("Content-Type", "application/json; charset=utf-8")
		writer.WriteHeader(400)
		encoder := json.NewEncoder(writer)
		err := encoder.Encode(model.ResponseFailJSON{
			Status:  "fail",
			Message: "Gagal menambahkan buku. Mohon isi nama buku",
		})

		if err != nil {
			panic(err)
		}
		return
	}

	if book.PageCount < book.ReadPage {
		writer.Header().Add("Content-Type", "application/json; charset=utf-8")
		writer.WriteHeader(400)
		encoder := json.NewEncoder(writer)
		err := encoder.Encode(model.ResponseFailJSON{
			Status:  "fail",
			Message: "Gagal menambahkan buku. readPage tidak boleh lebih besar dari pageCount",
		})

		if err != nil {
			panic(err)
		}
		return
	}

	id := strconv.Itoa(rand.Int() % 1000)
	finished := false
	insertedAt := time.Now().Format(time.RFC3339)
	updatedAt := insertedAt

	if book.ReadPage == book.PageCount {
		finished = true
	}

	bookMap := model.Book{
		Id:         id,
		Name:       book.Name,
		Year:       book.Year,
		Author:     book.Author,
		Summary:    book.Summary,
		Publisher:  book.Publisher,
		PageCount:  book.PageCount,
		ReadPage:   book.ReadPage,
		Finished:   book.Reading,
		Reading:    finished,
		InsertedAt: insertedAt,
		UpdatedAt:  updatedAt,
	}

	books = append(books, bookMap)

	bookId := model.PostBookResponse{Id: id}

	writer.Header().Add("Content-Type", "application/json; charset=utf-8")
	writer.WriteHeader(200)
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(model.ResponseJSON{
		Status:  "success",
		Message: "Buku berhasil ditambahkan",
		Data:    bookId,
	})

	if err != nil {
		panic(err)
	}
}

func GetBooksHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var book = []model.BooksResponse{}

	for i := 0; i < len(books); i++ {
		book = append(book, model.BooksResponse{
			Id:        books[i].Id,
			Name:      books[i].Name,
			Publisher: books[i].Publisher,
		})
	}

	writer.Header().Add("Content-Type", "application/json; charset=utf-8")
	writer.WriteHeader(200)
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(model.ResponseJSON{
		Status: "success",
		Data:   model.GetBooksResponse{book},
	})

	if err != nil {
		panic(err)
	}
}

func GetBookByIdHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var book = model.Book{}

	id := params.ByName("id")

	found := false

	for i := 0; i < len(books); i++ {
		if books[i].Id == id {
			book.Id = books[i].Id
			book.Name = books[i].Name
			book.Year = books[i].Year
			book.Author = books[i].Author
			book.Summary = books[i].Summary
			book.Publisher = books[i].Publisher
			book.PageCount = books[i].PageCount
			book.ReadPage = books[i].ReadPage
			book.Finished = books[i].Finished
			book.Reading = books[i].Reading
			book.InsertedAt = books[i].InsertedAt
			book.UpdatedAt = books[i].UpdatedAt

			found = true
		}
	}

	if found {
		writer.Header().Add("Content-Type", "application/json; charset=utf-8")
		writer.WriteHeader(200)
		encoder := json.NewEncoder(writer)
		err := encoder.Encode(model.ResponseJSON{
			Status: "success",
			Data:   model.GetBookByIdResponse{book},
		})

		if err != nil {
			panic(err)
		}
		return
	} else {
		writer.Header().Add("Content-Type", "application/json; charset=utf-8")
		writer.WriteHeader(404)
		encoder := json.NewEncoder(writer)
		err := encoder.Encode(model.ResponseFailJSON{
			Status:  "fail",
			Message: "Buku tidak ditemukan",
		})

		if err != nil {
			panic(err)
		}
		return
	}
}

func PutBookByIdHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	book := model.BookRequest{}

	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&book); err != nil {
		panic(err)
	}

	if book.Name == "" {
		writer.Header().Add("Content-Type", "application/json; charset=utf-8")
		writer.WriteHeader(400)
		encoder := json.NewEncoder(writer)
		err := encoder.Encode(model.ResponseFailJSON{
			Status:  "fail",
			Message: "Gagal memperbarui buku. Mohon isi nama buku",
		})

		if err != nil {
			panic(err)
		}
		return
	}

	if book.PageCount < book.ReadPage {
		writer.Header().Add("Content-Type", "application/json; charset=utf-8")
		writer.WriteHeader(400)
		encoder := json.NewEncoder(writer)
		err := encoder.Encode(model.ResponseFailJSON{
			Status:  "fail",
			Message: "Gagal memperbarui buku. readPage tidak boleh lebih besar dari pageCount",
		})

		if err != nil {
			panic(err)
		}
		return
	}

	id := params.ByName("id")

	found := false
	var index int

	for i := 0; i < len(books); i++ {
		if books[i].Id == id {
			index = i
			found = true
		}
	}

	if found {
		finished := false
		updatedAt := time.Now().Format(time.RFC3339)

		if book.ReadPage == book.PageCount {
			finished = true
		}

		books[index].Name = book.Name
		books[index].Year = book.Year
		books[index].Author = book.Author
		books[index].Summary = book.Summary
		books[index].Publisher = book.Publisher
		books[index].PageCount = book.PageCount
		books[index].ReadPage = book.ReadPage
		books[index].Finished = finished
		books[index].Reading = book.Reading
		books[index].UpdatedAt = updatedAt

		writer.Header().Add("Content-Type", "application/json; charset=utf-8")
		writer.WriteHeader(200)
		encoder := json.NewEncoder(writer)
		err := encoder.Encode(model.ResponseJSON{
			Status:  "success",
			Message: "Buku berhasil diperbarui",
		})

		if err != nil {
			panic(err)
		}
		return
	} else {
		writer.Header().Add("Content-Type", "application/json; charset=utf-8")
		writer.WriteHeader(404)
		encoder := json.NewEncoder(writer)
		err := encoder.Encode(model.ResponseFailJSON{
			Status:  "fail",
			Message: "Gagal memperbarui buku. Id tidak ditemukan",
		})

		if err != nil {
			panic(err)
		}
		return
	}
}

func DeleteBookByIdHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id := params.ByName("id")

	found := false
	var index int

	for i := 0; i < len(books); i++ {
		if books[i].Id == id {
			index = i
			found = true
		}
	}

	if found {
		books = append(books[:index], books[index+1:]...)

		writer.Header().Add("Content-Type", "application/json; charset=utf-8")
		writer.WriteHeader(200)
		encoder := json.NewEncoder(writer)
		err := encoder.Encode(model.ResponseJSON{
			Status:  "success",
			Message: "Buku berhasil dihapus",
		})

		if err != nil {
			panic(err)
		}
		return
	} else {
		writer.Header().Add("Content-Type", "application/json; charset=utf-8")
		writer.WriteHeader(404)
		encoder := json.NewEncoder(writer)
		err := encoder.Encode(model.ResponseFailJSON{
			Status:  "fail",
			Message: "Buku gagal dihapus. Id tidak ditemukan",
		})

		if err != nil {
			panic(err)
		}
		return
	}
}
