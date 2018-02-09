package test

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"

	"github.com/bxcodec/faker"
	"github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
	"github.com/rafaeljusto/redigomock"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

// NewSQLMock is mock db
func NewSQLMock() (sqlmock.Sqlmock, *gorm.DB) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("can't create sqlmock: %s", err)
	}

	gormDB, gerr := gorm.Open("postgres", db)
	if gerr != nil {
		log.Fatalf("can't open gorm connection: %s", err)
	}
	gormDB.LogMode(true)
	gormDB.SingularTable(true)

	return mock, gormDB.Set("gorm:update_column", true)
}

// NewCacheMock returns new cacheHandler.
// repository: https://github.com/garyburd/redigo/redis
func NewCacheMock() *redigomock.Conn {
	return redigomock.NewConn()
}

// StartMockServer func
func StartMockServer(path string) *httptest.Server {
	// visenze mock server
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		buf, _ := ioutil.ReadFile(path)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(buf)
	})
	rs := httptest.NewServer(handler)
	return rs
}

//GetFilePath get file path
func GetFilePath(folder string, name string) string {
	dir := os.Getenv("FR_CIRCLE_API_DIR")
	filePath := dir + "/test/" + folder + "/" + name
	return filePath
}

// FixedFullRe return fix full query
func FixedFullRe(s string) string {
	return fmt.Sprintf("^%s$", regexp.QuoteMeta(s))
}

// NewLogger func
func NewLogger() *logrus.Logger {
	var logger logrus.Logger
	if err := faker.FakeData(&logger); err != nil {
		_ = fmt.Errorf("error occured by fakeData for logger: %s", err)
	}
	return &logger
}
