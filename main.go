package main

import (
	"net/http"
	"os"

	"github.com/at-vudang95/go-food-market-api/infrastructure"

	"github.com/at-vudang95/go-food-market-api/router"
	"github.com/garyburd/redigo/redis"
	"github.com/go-chi/chi"
)

func main() {
	mux := chi.NewRouter()
	// sql new.
	sqlHandler := infrastructure.NewSQL()
	// s3 new.
	// s3Handler := infrastructure.NewS3()
	// cache new.
	// cacheHandler := infrastructure.NewCache()
	// logger new.
	loggerHandler := infrastructure.NewLogger()
	// translation new.
	translationHandler := infrastructure.NewTranslation()

	r := &router.Router{
		Mux:        mux,
		SQLHandler: sqlHandler,
		// S3Handler:  s3Handler,
		// CacheHandler:       cacheHandler,
		LoggerHandler:      loggerHandler,
		TranslationHandler: translationHandler,
	}

	r.InitializeRouter()
	r.SetupHandler()

	// after process
	defer closeLogger(r.LoggerHandler.Logfile)
	// defer closeRedis(r.CacheHandler.Conn)

	_ = http.ListenAndServe(":8081", mux)
}

// after process
func closeLogger(logfile *os.File) {
	// close file.
	if logfile != nil {
		_ = logfile.Close()
	}
}
func closeRedis(conn *redis.Conn) {
	// close redis connection.
	if conn != nil {
		_ = (*conn).Close()
	}
}
