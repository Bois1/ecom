package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Bois1/ecomm/types"
	"github.com/gorilla/mux"
)

func TestUserServiceHandlers(t *testing.T) {

	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	t.Run("should fail if the user payload is invalid", func(t *testing.T) {
		payload := types.RegiserUserPayload{
			FirstName: "test",
			LastName:  "sisi",
			Email:     "",
			Password:  "asdaque",
		}
		marshalled, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("unexpected status code: expected %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})

}

type mockUserStore struct{}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {

	return nil, nil
}

func (m *mockUserStore) GetUserById(id int) (*types.User, error) {

	return nil, nil
}

func (m *mockUserStore) CreateUser(user *types.User) error {
	return nil
}
