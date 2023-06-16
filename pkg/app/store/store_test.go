package store

import (
	"database/sql"
	"reflect"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"github.com/anindita02/books-api-server/pkg/app/models"
)

func TestCreateBook(t *testing.T) {
	// Create a new in-memory SQLite database for testing
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("failed to open database connection: %v", err)
	}
	defer db.Close()

	// Create the books table in the test database
	createTableQuery := `
        CREATE TABLE books (
            id INTEGER PRIMARY KEY,
            name TEXT,
            author TEXT,
            year INTEGER,
            language TEXT
        );
    `
	_, err = db.Exec(createTableQuery)
	if err != nil {
		t.Fatalf("failed to create books table: %v", err)
	}

	// Initialize the BookStore with the test database connection
	bookStore := &BookStore{db: db}

	// Define test cases
	testCases := []struct {
		name    string
		book    *models.Book
		wantErr bool
	}{
		{
			name: "Create new book",
			book: &models.Book{
				Name:     "New Book",
				Author:   "New Author",
				Year:     2022,
				Language: "English",
			},
			wantErr: false,
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := bookStore.CreateBook(tc.book)
			if (err != nil) != tc.wantErr {
				t.Errorf("unexpected error: %v", err)
			}

			// Retrieve the created book from the database
			createdBook, err := bookStore.GetBookByID(tc.book.ID)
			if err != nil {
				t.Fatalf("failed to retrieve created book: %v", err)
			}

			// Compare the created book with the expected book
			if !reflect.DeepEqual(createdBook, tc.book) {
				t.Errorf("unexpected created book:\ngot:  %+v\nwant: %+v", createdBook, tc.book)
			}
		})
	}
}

func TestUpdateBook(t *testing.T) {
	// Create a new in-memory SQLite database for testing
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("failed to open database connection: %v", err)
	}
	defer db.Close()

	// Create the books table in the test database
	createTableQuery := `
        CREATE TABLE books (
            id INTEGER PRIMARY KEY,
            name TEXT,
            author TEXT,
            year INTEGER,
            language TEXT
        );
    `
	_, err = db.Exec(createTableQuery)
	if err != nil {
		t.Fatalf("failed to create books table: %v", err)
	}

	// Initialize the BookStore with the test database connection
	bookStore := &BookStore{db: db}

	// Insert the initial book record
	initialBook := &models.Book{
		ID:       1,
		Name:     "Initial Book",
		Author:   "Initial Author",
		Year:     2021,
		Language: "English",
	}
	err = bookStore.CreateBook(initialBook)
	if err != nil {
		t.Fatalf("failed to insert initial book: %v", err)
	}
	// Define test cases
	testCases := []struct {
		name    string
		book    *models.Book
		wantErr bool
	}{
		{
			name: "Update existing book",
			book: &models.Book{
				ID:       1,
				Name:     "New Book",
				Author:   "New Author",
				Year:     2022,
				Language: "English",
			},
			wantErr: false,
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := bookStore.UpdateBook(tc.book)
			if (err != nil) != tc.wantErr {
				t.Errorf("unexpected error: %v", err)
			}

			// Retrieve the updated book from the database
			updatedBook := &models.Book{}
			query := "SELECT id, name, author, year, language FROM books WHERE id = ?"
			err = db.QueryRow(query, tc.book.ID).Scan(&updatedBook.ID, &updatedBook.Name, &updatedBook.Author, &updatedBook.Year, &updatedBook.Language)
			if err != nil {
				if err == sql.ErrNoRows {
					t.Fatalf("book not found in the database")
				}
				t.Fatalf("failed to retrieve updated book: %v", err)
			}

			// Compare the updated book with the expected book
			if !reflect.DeepEqual(updatedBook, tc.book) {
				t.Errorf("unexpected updated book:\ngot:  %+v\nwant: %+v", updatedBook, tc.book)
			}
		})
	}
}

