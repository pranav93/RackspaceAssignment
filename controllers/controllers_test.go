package controllers_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/pranav93/RackspaceAssignment/scripts"
	"github.com/pranav93/RackspaceAssignment/setup"
)

func init() {
	scripts.CreateRulesDB()
	scripts.CreateProductsDB()
	scripts.CreateCartsDB()
}

func TestCartCreateAPI(t *testing.T) {
	// The setupServer method, that we previously refactored
	// is injected into a test server
	ts := httptest.NewServer(setup.Server())
	// Shut down the server and block until all requests have gone through
	defer ts.Close()

	// Make a request to our server with the {base url}/ping
	payload := strings.NewReader(`
	{
		"cartItems": {
			"CH1": 1,
			"AP1": 3
		}
	}"
	`)

	resp, err := http.Post(fmt.Sprintf("%s/cart/save/", ts.URL), "application/json", payload)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

	val, ok := resp.Header["Content-Type"]

	// Assert that the "content-type" header is actually set
	if !ok {
		t.Fatalf("Expected Content-Type header to be set")
	}

	// Assert that it was set as expected
	if val[0] != "application/json; charset=utf-8" {
		t.Fatalf("Expected \"application/json; charset=utf-8\", got %s", val[0])
	}

	type getIDStruct struct {
		Data struct {
			Cart struct {
				ID string `json:"ID"`
			} `json:"cart"`
		} `json:"data"`
	}

	dec := json.NewDecoder(resp.Body)

	for dec.More() {
		var response getIDStruct
		err := dec.Decode(&response)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Cart id is %s\n", response.Data.Cart.ID)
	}
}

func TestCartDeleteAPI(t *testing.T) {
	// The setupServer method, that we previously refactored
	// is injected into a test server
	ts := httptest.NewServer(setup.Server())
	// Shut down the server and block until all requests have gone through
	defer ts.Close()

	// Make a request to our server with the {base url}/ping
	payload := strings.NewReader(`
	{
		"cartItems": {
			"CH1": 1,
			"AP1": 3
		}
	}"
	`)

	resp, err := http.Post(fmt.Sprintf("%s/cart/save/", ts.URL), "application/json", payload)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	type getIDStruct struct {
		Data struct {
			Cart struct {
				ID string `json:"ID"`
			} `json:"cart"`
		} `json:"data"`
	}

	dec := json.NewDecoder(resp.Body)

	var cartID string
	for dec.More() {
		var response getIDStruct
		err := dec.Decode(&response)
		if err != nil {
			t.Fatal(err)
		}
		log.Printf("Cart id is %s\n", response.Data.Cart.ID)
		cartID = response.Data.Cart.ID
	}

	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/cart/%s/", ts.URL, cartID), nil)
	if err != nil {
		t.Fatal(err)
	}

	// Fetch Request
	resp, err = client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	type GetStatusStruct struct {
		Data struct {
			Deleted bool `json:"deleted"`
		} `json:"data"`
	}
	// Read Response Body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	dec = json.NewDecoder(strings.NewReader(string(respBody)))

	for dec.More() {
		var response GetStatusStruct
		err := dec.Decode(&response)
		if err != nil {
			t.Fatal(err)
		}
		if response.Data.Deleted != true {
			t.Fatalf("Deleted should be true\n")
		}
	}
}

func TestCartCheckoutAPI(t *testing.T) {
	// The setupServer method, that we previously refactored
	// is injected into a test server
	ts := httptest.NewServer(setup.Server())
	// Shut down the server and block until all requests have gone through
	defer ts.Close()

	// Make a request to our server with the {base url}/ping
	payload := strings.NewReader(`
	{
		"cartItems": {
			"CH1": 1,
			"AP1": 3
		}
	}"
	`)

	resp, err := http.Post(fmt.Sprintf("%s/cart/save/", ts.URL), "application/json", payload)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	type getIDStruct struct {
		Data struct {
			Cart struct {
				ID string `json:"ID"`
			} `json:"cart"`
		} `json:"data"`
	}

	dec := json.NewDecoder(resp.Body)

	var cartID string
	for dec.More() {
		var response getIDStruct
		err := dec.Decode(&response)
		if err != nil {
			t.Fatal(err)
		}
		log.Printf("Cart id is %s\n", response.Data.Cart.ID)
		cartID = response.Data.Cart.ID
	}

	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/cart/checkout/%s/", ts.URL, cartID), nil)
	if err != nil {
		t.Fatal(err)
	}

	// Fetch Request
	resp, err = client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	// Read Response Body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	dec = json.NewDecoder(strings.NewReader(string(respBody)))

	for dec.More() {
		var response getIDStruct
		err := dec.Decode(&response)
		if err != nil {
			t.Fatal(err)
		}
		if response.Data.Cart.ID != cartID {
			t.Fatalf("Expected cart id %s, but got %s\n", cartID, response.Data.Cart.ID)
		}
	}
}

func TestCartGetAPI(t *testing.T) {
	// The setupServer method, that we previously refactored
	// is injected into a test server
	ts := httptest.NewServer(setup.Server())
	// Shut down the server and block until all requests have gone through
	defer ts.Close()

	// Make a request to our server with the {base url}/ping
	payload := strings.NewReader(`
	{
		"cartItems": {
			"CH1": 1,
			"AP1": 3
		}
	}"
	`)

	resp, err := http.Post(fmt.Sprintf("%s/cart/save/", ts.URL), "application/json", payload)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	type getIDStruct struct {
		Data struct {
			Cart struct {
				ID string `json:"ID"`
			} `json:"cart"`
		} `json:"data"`
	}

	dec := json.NewDecoder(resp.Body)

	var cartID string
	for dec.More() {
		var response getIDStruct
		err := dec.Decode(&response)
		if err != nil {
			t.Fatal(err)
		}
		log.Printf("Cart id is %s\n", response.Data.Cart.ID)
		cartID = response.Data.Cart.ID
	}

	resp, err = http.Get(fmt.Sprintf("%s/cart/%s/", ts.URL, cartID))

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

	// Read Response Body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf(string(respBody))
	dec = json.NewDecoder(strings.NewReader(string(respBody)))

	for dec.More() {
		var response getIDStruct
		err := dec.Decode(&response)
		if err != nil {
			t.Fatal(err)
		}
		if response.Data.Cart.ID != cartID {
			t.Fatalf("Expected cart id %s, but got %s\n", cartID, response.Data.Cart.ID)
		}
	}
}

func TestCartUpdateAPI(t *testing.T) {
	// The setupServer method, that we previously refactored
	// is injected into a test server
	ts := httptest.NewServer(setup.Server())
	// Shut down the server and block until all requests have gone through
	defer ts.Close()

	// Make a request to our server with the {base url}/ping
	payload := strings.NewReader(`
	{
		"cartItems": {
			"CH1": 1,
			"AP1": 3
		}
	}"
	`)

	resp, err := http.Post(fmt.Sprintf("%s/cart/save/", ts.URL), "application/json", payload)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	type getIDStruct struct {
		Data struct {
			Cart struct {
				ID string `json:"ID"`
			} `json:"cart"`
		} `json:"data"`
	}

	dec := json.NewDecoder(resp.Body)

	var cartID string
	for dec.More() {
		var response getIDStruct
		err := dec.Decode(&response)
		if err != nil {
			t.Fatal(err)
		}
		log.Printf("Cart id is %s\n", response.Data.Cart.ID)
		cartID = response.Data.Cart.ID
	}

	payload = strings.NewReader(
		`{
			"cartItems": {
				"add": ["AP1"]
				}
			}"
		}`)

	client := &http.Client{}
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/cart/save/%s/", ts.URL, cartID), payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	defer res.Body.Close()
	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
}
