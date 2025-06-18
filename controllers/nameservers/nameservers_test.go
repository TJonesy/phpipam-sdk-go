package nameservers

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/tjonesy/phpipam-sdk-go/phpipam"
	"github.com/tjonesy/phpipam-sdk-go/phpipam/session"
)

var testCreateNameserverInput = Nameserver{
	Name:        "foolan",
	NameSrv1:    "8.8.8.8",
	Description: "google",
	Permissions: "1",
}

const testCreateNameserverOutputExpected = `Nameserver created`
const testCreateNameserverOutputJSON = `
{
  "code": 201,
  "success": true,
  "data": "Nameserver created"
}
`

var testGetNameserverByIDOutputExpected = Nameserver{
	ID:          3,
	Name:        "foolan",
	NameSrv1:    "8.8.8.8",
	Description: "google",
	Permissions: "1",
}

const testGetNameserverByIDOutputJSON = `
{
  "code": 200,
  "success": true,
  "data": {
    "id": 3,
    "nameSrv1": "8.8.8.8",
    "name": "foolan",
    "description": "google",
    "permissions": "1",
    "editDate": null,
    "links": [
      {
        "rel": "self",
        "href": "/api/test/nameservers/3/",
        "methods": [
          "GET",
          "POST",
          "DELETE",
          "PATCH"
        ]
      }
	]
  }
}
`

var testUpdateNameserverInput = Nameserver{
	ID:   3,
	Name: "bazlan",
}

const testUpdateNameserverOutputExpected = `Nameserver updated`
const testUpdateNameserverOutputJSON = `
{
  "code": 200,
  "success": true,
  "data": "Nameserver updated"
}
`

const testDeleteNameserverOutputExpected = `Nameserver deleted`
const testDeleteNameserverOutputJSON = `
{
  "code": 200,
  "success": true,
  "data": "Nameserver deleted"
}
`

func newHTTPTestServer(f func(w http.ResponseWriter, r *http.Request)) *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(f))
	return ts
}

func httpOKTestServer(output string) *httptest.Server {
	return newHTTPTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		http.Error(w, output, http.StatusOK)
	})
}

func httpCreatedTestServer(output string) *httptest.Server {
	return newHTTPTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		http.Error(w, output, http.StatusCreated)
	})
}

func fullSessionConfig() *session.Session {
	return &session.Session{
		Config: phpipam.Config{
			AppID:    "0123456789abcdefgh",
			Password: "changeit",
			Username: "nobody",
		},
		Token: session.Token{
			String: "foobarbazboop",
		},
	}
}

func TestCreateNameserver(t *testing.T) {
	ts := httpCreatedTestServer(testCreateNameserverOutputJSON)
	defer ts.Close()
	sess := fullSessionConfig()
	sess.Config.Endpoint = ts.URL
	client := NewController(sess)

	in := testCreateNameserverInput
	expected := testCreateNameserverOutputExpected
	actual, err := client.CreateNameserver(in)
	if err != nil {
		t.Fatalf("Bad: %s", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected %#v, got %#v", expected, actual)
	}
}

func TestGetNameserverByID(t *testing.T) {
	ts := httpOKTestServer(testGetNameserverByIDOutputJSON)
	defer ts.Close()
	sess := fullSessionConfig()
	sess.Config.Endpoint = ts.URL
	client := NewController(sess)

	expected := testGetNameserverByIDOutputExpected
	actual, err := client.GetNameserverByID(3)
	if err != nil {
		t.Fatalf("Bad: %s", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected %#v, got %#v", expected, actual)
	}
}

func TestUpdateNameserver(t *testing.T) {
	ts := httpOKTestServer(testUpdateNameserverOutputJSON)
	defer ts.Close()
	sess := fullSessionConfig()
	sess.Config.Endpoint = ts.URL
	client := NewController(sess)

	in := testUpdateNameserverInput
	expected := testUpdateNameserverOutputExpected
	actual, err := client.UpdateNameserver(in)
	if err != nil {
		t.Fatalf("Bad: %s", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected %#v, got %#v", expected, actual)
	}
}

func TestDeleteNameserver(t *testing.T) {
	ts := httpOKTestServer(testDeleteNameserverOutputJSON)
	defer ts.Close()
	sess := fullSessionConfig()
	sess.Config.Endpoint = ts.URL
	client := NewController(sess)

	expected := testDeleteNameserverOutputExpected
	actual, err := client.DeleteNameserver(3)
	if err != nil {
		t.Fatalf("Bad: %s", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected %#v, got %#v", expected, actual)
	}
}
