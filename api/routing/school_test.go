package routing

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/mouzkolit/GOCli/database"
	"github.com/stretchr/testify/assert"
)

var testDB *database.DB

func TestMain(m *testing.M) {
	// Set up the test database
	var err error
	testDB, err = database.NewDB()
	if err != nil {
		panic("Failed to open database: " + err.Error())
	}
	defer testDB.Close()

	testDB.InitSchema()

	// Run the tests
	code := m.Run()

	// Exit with the test result code
	os.Exit(code)
}

func TestCreateSchool(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("successful school creation", func(t *testing.T) {

		r := gin.Default()
		CreateSchool(r, testDB)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/school?name=Test School&place=Test Place&web=http://test.com", nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)

		// Verify the school was created in the database
		school, err := testDB.GetSchool(0)
		assert.NoError(t, err)
		assert.Equal(t, "Test School", school.Name)
		assert.Equal(t, "Test Place", school.Location)
		assert.Equal(t, "http://test.com", school.SchoolWeb)

	})

	t.Run("failed school creation", func(t *testing.T) {

		// Simulate a database error by closing the connection

		r := gin.Default()
		CreateSchool(r, testDB)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/school?name=Test School&place=Test Place&web=http://test.com", nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, 500, w.Code)
	})
}
