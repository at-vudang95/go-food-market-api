package test

import (
	"bytes"
	"mime/multipart"
	"os"
	"testing"

	"github.com/at-vudang95/go-food-market-api/infrastructure"

	"net/http"
	"net/http/httptest"

	"github.com/at-vudang95/go-food-market-api/shared/utils"
	"github.com/garyburd/redigo/redis"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	mMiddleware "github.com/at-vudang95/go-food-market-api/shared/middleware"
)

// APITest struct.
type APITest struct {
	T       *testing.T
	TRouter Router
}

// Router is application struct hold Mux and db connection
type Router struct {
	Mux                *chi.Mux
	SQLHandler         *infrastructure.SQL
	S3Handler          *infrastructure.S3
	CacheHandler       *infrastructure.Cache
	LoggerHandler      *infrastructure.Logger
	TranslationHandler *infrastructure.Translation
	SearchAPIHandler   infrastructure.SearchAPI
}

// InitializeRouter initializes Mux and middleware
func (at *APITest) InitializeRouter() {
	at.TRouter.Mux.Use(middleware.RequestID)
	at.TRouter.Mux.Use(middleware.RealIP)
	// Custom middleware(Translation)
	at.TRouter.Mux.Use(at.TRouter.TranslationHandler.Middleware.Middleware)
	// Custom middleware(Logger)
	at.TRouter.Mux.Use(mMiddleware.Logger(at.TRouter.LoggerHandler))
}

// ExecuteRequest requests ServeHTTP.
func (at *APITest) ExecuteRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	at.TRouter.Mux.ServeHTTP(rr, req)

	return rr
}

// CloseLogger close Logger.
func (at *APITest) CloseLogger(logfile *os.File) {
	// close file.
	if logfile != nil {
		_ = logfile.Close()
	}
}

// CloseRedis close redis connection.
func (at *APITest) CloseRedis(conn *redis.Conn) {
	// close redis connection.
	if conn != nil {
		_ = (*conn).Close()
	}
}

type a struct {
	conn redis.Conn
}

// NewTestMain is TestMain function initialize.
func NewTestMain(m *testing.M, fnSetupHandler func(*Router)) (*chi.Mux, *APITest) {
	mux := chi.NewRouter()
	apiHandler := infrastructure.NewSearchAPI()
	return mux, NewAPITest(nil, apiHandler, fnSetupHandler)
}

// NewAPITest initialize the apiTest
func NewAPITest(t *testing.T, apiHandler infrastructure.SearchAPI, fnSetupHandler func(*Router)) *APITest {
	mux := chi.NewRouter()
	// sql new.
	sqlHandler := infrastructure.NewSQL()

	// s3 new.
	s3Handler := infrastructure.NewS3()
	// cache new.
	// cacheHandler := infrastructure.NewCache()
	// conn := NewCacheMock()
	cacheHandler := &infrastructure.Cache{Conn: nil}

	// logger new.
	loggerHandler := infrastructure.NewLogger()
	// translation new.
	translationHandler := infrastructure.NewTranslation()

	TRouter := &Router{Mux: mux, SQLHandler: sqlHandler, S3Handler: s3Handler, CacheHandler: cacheHandler, LoggerHandler: loggerHandler, TranslationHandler: translationHandler, SearchAPIHandler: apiHandler}
	apiTest := &APITest{TRouter: *TRouter, T: t}

	apiTest.InitializeRouter()
	fnSetupHandler(TRouter)

	// after process
	//http.ListenAndServe(":8080", mux)
	return apiTest
}

//SendRequestWithBody send request with body
func (at *APITest) SendRequestWithBody(fields map[string][]string, method, url, token string) *httptest.ResponseRecorder {
	body, writer := at.CreateRequestBody(fields)
	return at.CallRequest(body, writer, method, url, token, writer.FormDataContentType())
}

// SendRequestWithContentType send request with ContentType
func (at *APITest) SendRequestWithContentType(fields map[string][]string, method, url, token, contentType string) *httptest.ResponseRecorder {
	body, writer := at.CreateRequestBody(fields)
	return at.CallRequest(body, writer, method, url, token, contentType)

}

//CallRequest call request
func (at *APITest) CallRequest(body *bytes.Buffer, writer *multipart.Writer, method, url, token, contentType string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, url, body)
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", contentType)
	resp := at.ExecuteRequest(req)
	return resp
}

//CreateRequestBody create request body
func (at *APITest) CreateRequestBody(fields map[string][]string) (body *bytes.Buffer, writer *multipart.Writer) {
	body = &bytes.Buffer{}
	writer = multipart.NewWriter(body)
	if value, ok := fields["image_file_path"]; ok {
		_ = utils.MultipartFileWriter(writer, "image_file", value[0])
		delete(fields, "image_file_path")
	}
	for key, values := range fields {
		for _, val := range values {
			_ = writer.WriteField(key, val)
		}
	}
	_ = writer.Close()
	return body, writer
}
