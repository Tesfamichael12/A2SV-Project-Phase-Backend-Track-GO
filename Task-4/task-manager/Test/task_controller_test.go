package Test

import (
	"net/http"
	"net/http/httptest"
	"task_manager/router"
	"testing"

	"github.com/stretchr/testify/assert"
)

var r = router.Router()

func TestGetTasks(t *testing.T) {

	// create test request
	req, _ := http.NewRequest("GET", "/tasks", nil)

	// record the response status-code and body
	w := httptest.NewRecorder()

	// serve the requst
	r.ServeHTTP(w, req)

	// assertions
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Task 1")
}

func TestGetTaskByID(t *testing.T) {

	req, _ := http.NewRequest("GET", "/tasks/1", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Task 1")
}

func TestDeleteTask(t *testing.T) {

	req, _ := http.NewRequest("DELETE", "/tasks/1", nil)
	w := httptest.NewRecorder()
	
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Task deleted successfully")
}
