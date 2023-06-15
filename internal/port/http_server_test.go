package port

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/trevatk/go-pkg/logging"
)

type HTTPServerSuite struct {
	suite.Suite
	mux *chi.Mux
}

func (suite *HTTPServerSuite) SetupTest() {

	a := assert.New(suite.T())

	l, e := logging.New()
	a.NoError(e)

	server := NewHTTPServer(l)

	suite.mux = NewRouter(server)
}

func (suite *HTTPServerSuite) TestHealth() {

	assert := assert.New(suite.T())

	cases := []struct {
		code int
		body string
	}{
		{
			// success case
			code: http.StatusOK,
			body: "OK",
		},
	}

	for _, c := range cases {

		req, err := http.NewRequest(http.MethodGet, "/health", nil)
		assert.NoError(err)

		rr := httptest.NewRecorder()

		suite.mux.ServeHTTP(rr, req)

		assert.Equal(c.code, rr.Code)

		body, err := io.ReadAll(rr.Body)
		assert.NoError(err)

		assert.Equal(c.body, string(body))
	}
}

func TestHttpServerSuite(t *testing.T) {
	suite.Run(t, new(HTTPServerSuite))
}
