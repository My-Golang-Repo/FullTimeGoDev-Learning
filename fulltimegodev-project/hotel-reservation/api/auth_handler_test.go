package api

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/fulltimegodev/hotel-reservation-nana/db"
	"github.com/fulltimegodev/hotel-reservation-nana/types"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

var (
	ctx = context.Background()
)

func insertUserTest(t *testing.T, store db.UserStore) *types.User {

	user, err := types.NewUserFromParams(types.CreateUserParam{
		FirstName: "Antares",
		LastName:  "AAA-0005",
		Email:     "antares@uncf.org",
		Password:  "supersecurepassword",
	})

	if err != nil {
		log.Fatal(err)
	}

	_, err = store.InsertUser(ctx, user)
	if err != nil {
		log.Fatal(err)
	}

	return user
}

func TestAuthenticateSuccess(t *testing.T) {
	tdb := setup(t)
	defer tdb.teardown(t)
	insertedUser := insertUserTest(t, tdb)

	app := fiber.New()
	authHandler := NewAuthHandler(tdb)
	app.Post("/auth", authHandler.HandleAuthenticate)

	params := AuthParams{
		Email:    "antares@uncf.org",
		Password: "supersecurepassword",
	}

	b, _ := json.Marshal(params)

	req := httptest.NewRequest("POST", "/auth", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")
	resp, _ := app.Test(req)

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expecting status code 200 but got %d", resp.StatusCode)
	}

	var authResp AuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&authResp); err != nil {
		t.Fatal(err)
	}

	if authResp.Token == "" {
		t.Fatal("expected jwt token to be present in the auth response")
	}

	insertedUser.EncryptedPassword = ""
	if !reflect.DeepEqual(authResp.User, insertedUser) {
		t.Fatalf("expecting user %+v but got %+v", authResp.User, insertedUser)
	}
}

func TestAuthenticateWithWrongPassword(t *testing.T) {
	tdb := setup(t)
	defer tdb.teardown(t)

	app := fiber.New()
	authHandler := NewAuthHandler(tdb)
	app.Post("/auth", authHandler.HandleAuthenticate)

	params := AuthParams{
		Email:    "antares@uncf.org",
		Password: "supersecurepasswordNotCorrect",
	}

	b, _ := json.Marshal(params)

	req := httptest.NewRequest("POST", "/auth", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")
	resp, _ := app.Test(req)

	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expecting status code 400 but got %d", resp.StatusCode)
	}

	var genResp genericResp
	if err := json.NewDecoder(resp.Body).Decode(&genResp); err != nil {
		t.Fatal(err)
	}

	if genResp.Type != "error" {
		t.Fatalf("expected generic response type to be error but got %s", genResp.Type)
	}

	if genResp.Message != "invalid credentials" {
		t.Fatalf("expected the message response is <invalid credentials> but got %s", genResp.Message)
	}
}
