package api

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/fulltimegodev/hotel-reservation-nana/db"
	"github.com/fulltimegodev/hotel-reservation-nana/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http/httptest"
	"testing"
)

const dbnametest = "hotel-reservation-test"
const testmongouri = "mongodb://localhost:27017"

type testdb struct {
	db.UserStore
}

func (tdb *testdb) teardown(t *testing.T) {
	if err := tdb.UserStore.Drop(context.TODO()); err != nil {
		t.Fatal(err)
	}
}

func setup(t *testing.T) *testdb {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(testmongouri))
	if err != nil {
		log.Fatal(err)
	}

	return &testdb{
		UserStore: db.NewMongoUserStore(client, dbnametest),
	}
}

func TestPostUser(t *testing.T) {
	tdb := setup(t)
	defer tdb.teardown(t)

	app := fiber.New()

	userHandler := NewUserHandler(tdb)

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
