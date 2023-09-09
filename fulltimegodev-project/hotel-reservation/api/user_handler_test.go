package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/fulltimegodev/hotel-reservation-nana/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPostUser(t *testing.T) {
	tdb := setup(t)
	defer tdb.teardown(t)

	app := fiber.New()

	userHandler := NewUserHandler(tdb.store.User)

	app.Post("/", userHandler.HandlePostUser)

	params := types.CreateUserParam{
		FirstName: "Apollo Norm",
		LastName:  "AAA-0003",
		Email:     "apollonorm@uncf.org",
		Password:  "advancedarmamentartillery",
	}

	b, _ := json.Marshal(params)
	req := httptest.NewRequest("POST", "/", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")
	resp, _ := app.Test(req)
	var user types.User
	json.NewDecoder(resp.Body).Decode(&user)

	if user.FirstName != params.FirstName {
		t.Errorf("expected first name %s but got %s", params.FirstName, user.FirstName)
	}
	if user.LastName != params.LastName {
		t.Errorf("expected last name %s but got %s", params.LastName, user.LastName)
	}
	if user.Email != params.Email {
		t.Errorf("expected email %s but got %s", params.Email, user.Email)
	}
	if len(user.ID) == 0 {
		t.Errorf("expecting a user id to be set")
	}
	if len(user.EncryptedPassword) > 0 {
		t.Errorf("Encrypted password must not shown in json")
	}
}

func TestGetUserById(t *testing.T) {
	tdb := setup(t)
	defer tdb.teardown(t)

	app := fiber.New()
	userHandler := NewUserHandler(tdb.store.User)
	app.Get("/user/:id", userHandler.HandleGetUser)

	expectedUser := types.User{
		ID:                primitive.NewObjectID(),
		FirstName:         "Apollo Norm",
		LastName:          "AAA0003",
		Email:             "apollonrom@uncf.org",
		EncryptedPassword: "advancedartilleryarmament",
	}

	insertedUser, err := tdb.store.User.InsertUser(context.TODO(), &expectedUser)
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest("GET", fmt.Sprintf("/user/%s", insertedUser.ID.Hex()), nil)
	req.Header.Add("Content-Type", "application/json")
	resp, _ := app.Test(req)

	var retrievedUser types.User
	json.NewDecoder(resp.Body).Decode(&retrievedUser)

	if retrievedUser.ID.Hex() != insertedUser.ID.Hex() {
		t.Error("The ID is not match")
	}

	if retrievedUser.FirstName != insertedUser.FirstName {
		t.Errorf("Expecting first name %s but got %s", retrievedUser.FirstName, insertedUser.FirstName)
	}

	if retrievedUser.LastName != insertedUser.LastName {
		t.Errorf("Expecting last name %s but got %s", retrievedUser.LastName, insertedUser.LastName)
	}

	if retrievedUser.Email != insertedUser.Email {
		t.Errorf("Expecting email %s but got %s", retrievedUser.Email, retrievedUser.Email)
	}
}

func TestGetUsers(t *testing.T) {
	tdb := setup(t)
	defer tdb.teardown(t)

	app := fiber.New()
	userHandler := NewUserHandler(tdb.store.User)
	app.Get("/user", userHandler.HandleGetUsers)

	expectedUser := []types.User{
		{
			ID:                primitive.NewObjectID(),
			FirstName:         "Apollo",
			LastName:          "Norm",
			Email:             "apollonorm@uncf.org",
			EncryptedPassword: "advancedartilleryarmament",
		},
		{
			ID:                primitive.NewObjectID(),
			FirstName:         "Antares",
			LastName:          "AAA-0005",
			Email:             "antares@uncf.org",
			EncryptedPassword: "advancedartilleryarmament",
		},
	}

	for _, user := range expectedUser {
		_, err := tdb.store.User.InsertUser(context.TODO(), &user)
		if err != nil {
			t.Fatal(err)
		}
	}

	req := httptest.NewRequest("GET", "/user", nil)
	req.Header.Add("Content-Type", "application/json")
	resp, _ := app.Test(req)

	var retrievedUsers []types.User
	json.NewDecoder(resp.Body).Decode(&retrievedUsers)

	if len(retrievedUsers) != len(expectedUser) {
		t.Errorf("Expecting number of users is %d but only got %d", len(expectedUser), len(retrievedUsers))
	}

	for i, expected := range expectedUser {
		if retrievedUsers[i].ID.Hex() != expected.ID.Hex() {
			t.Errorf("expecting ID %s but got %s", expected.ID.Hex(), retrievedUsers[i].ID.Hex())
		}

		if retrievedUsers[i].FirstName != expected.FirstName {
			t.Errorf("expecting first name %s but got %s", expected.FirstName, retrievedUsers[i].FirstName)
		}

		if retrievedUsers[i].LastName != expected.LastName {
			t.Errorf("expecting first name %s but got %s", expected.LastName, retrievedUsers[i].LastName)
		}

		if retrievedUsers[i].Email != expected.Email {
			t.Errorf("expecting first name %s but got %s", expected.Email, retrievedUsers[i].Email)
		}
	}
}

func TestDeleteUser(t *testing.T) {
	tdb := setup(t)
	defer tdb.teardown(t)

	app := fiber.New()
	userHandler := NewUserHandler(tdb.store.User)
	app.Delete("/user/:id", userHandler.HandleDeleteUser)

	expectedUser := types.User{
		ID:                primitive.NewObjectID(),
		FirstName:         "Medalusa",
		LastName:          "QIU",
		Email:             "medalusa@gatlantis.org",
		EncryptedPassword: "zorderalliance",
	}

	insertedUser, err := tdb.store.User.InsertUser(context.TODO(), &expectedUser)
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest("DELETE", fmt.Sprintf("/user/%s", insertedUser.ID.Hex()), nil)
	req.Header.Add("Content-Type", "application/json")
	resp, _ := app.Test(req)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expecting status code %d but got %d", http.StatusOK, resp.StatusCode)
	}

	retrievedUser, err := tdb.store.User.GetUserByID(context.TODO(), insertedUser.ID.Hex())
	if err != mongo.ErrNoDocuments {
		t.Errorf("Expecting user to be deleted, but %+v user is retrieved", retrievedUser)
	}
}

func TestUpdateUser(t *testing.T) {
	tdb := setup(t)
	defer tdb.teardown(t)

	app := fiber.New()
	userHandler := NewUserHandler(tdb.store.User)
	app.Put("/user/:id", userHandler.HandlePutUser)

	initialUser := types.User{
		ID:                primitive.NewObjectID(),
		FirstName:         "Andromeda",
		LastName:          "AAA-0001-2202",
		Email:             "andromeda@uncf.org",
		EncryptedPassword: "advancedartilleryarmament",
	}

	_, err := tdb.store.User.InsertUser(context.TODO(), &initialUser)
	if err != nil {
		t.Fatal(err)
	}

	params := types.UpdateUserParam{
		FirstName: "Andromeda Kai",
		LastName:  "ZZZ-0001-YF-2203",
	}

	b, _ := json.Marshal(params)
	req := httptest.NewRequest("PUT", fmt.Sprintf("/user/%s", initialUser.ID.Hex()), bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")
	resp, _ := app.Test(req)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expecting %d status code but got %d", http.StatusOK, resp.StatusCode)
	}

	updatedUser, err := tdb.store.User.GetUserByID(context.TODO(), initialUser.ID.Hex())
	if err != nil {
		t.Fatal(err)
	}

	if updatedUser.FirstName != params.FirstName {
		t.Errorf("Expecting %s but got %s", updatedUser.FirstName, params.FirstName)
	}

	if updatedUser.LastName != params.LastName {
		t.Errorf("Expecting %s but got %s", updatedUser.LastName, params.LastName)
	}
}
