package services

import (
	"errors"
	"testing"

	"github.com/anindita02/books-api-server/mocks/pkg/app/store"
	"github.com/anindita02/books-api-server/pkg/app/models"
	"gotest.tools/assert"
)

func TestCreateBook(t *testing.T) {
	// Create a new mock store
	store := &store.Store{}

	// Create the service instance using the mock store
	service := &BookService{store: store}

	// Define test cases
	testCases := []struct {
		name        string
		book        *models.Book
		storeErr    error
		expectedErr error
	}{
		{
			name: "Successful creation",
			book: &models.Book{
				ID:       1,
				Name:     "Test Book",
				Author:   "Test Author",
				Year:     2021,
				Language: "English",
			},
			storeErr:    nil,
			expectedErr: nil,
		},
		{
			name: "Failed creation",
			book: &models.Book{
				ID:       2,
				Name:     "Another Book",
				Author:   "Another Author",
				Year:     2022,
				Language: "English",
			},
			storeErr:    errors.New("failed to create book"),
			expectedErr: errors.New("failed to create book"),
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Configure the mock store's behavior
			store.On("CreateBook", tc.book).Return(tc.storeErr)

			// Call the service method
			err := service.CreateBook(tc.book)

			// Assert the result
			assert.Error(t, err, "expected an error")
			assert.Equal(t, tc.expectedErr.Error(), err.Error())

			// Verify the interactions with the mock store
			store.AssertCalled(t, "CreateBook", tc.book)
			store.AssertExpectations(t)
		})
	}
}
