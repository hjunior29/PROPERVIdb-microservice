package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hjunior29/PROPERVIdb-microservice/routes"
	"github.com/stretchr/testify/assert"
)

func TestStatus(t *testing.T) {

	loadEnv()

	r := routes.HandleRequest()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "running")

}

func ServeHTTP(w *httptest.ResponseRecorder, req *http.Request) {
	panic("unimplemented")
}

// func TestAddProperties(t *testing.T) {

// 	loadEnv()

// 	r := routes.HandleRequest()

// 	testProperties := []models.Properties{
// 		{
// 			Address: "Goiânia, GO",
// 			Price:   250000,
// 		},
// 	}

// 	payload, _ := json.Marshal(testProperties)
// 	fmt.Println(string(payload))
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("POST", "/add-properties", bytes.NewBuffer(payload))
// 	// req.Header.Set("Content-Type", "application/json")
// 	r.ServeHTTP(w, req)

// 	assert.Equal(t, 200, w.Code)
// 	assert.Contains(t, w.Body.String(), "Properties added successfully")

// }
