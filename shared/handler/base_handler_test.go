package handler

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"net/url"
	"testing"

	"github.com/at-vudang95/go-food-market-api/infrastructure"

	"io/ioutil"
	"os"

	"net/http/httptest"

	"encoding/json"

	"github.com/at-vudang95/go-food-market-api/shared/utils"
	"github.com/go-chi/chi"

	"github.com/stretchr/testify/assert"
)

type User struct {
	Name string `form:"name" json:"name" validate:"required"`
	Pass string `form:"pass" json:"pass" validate:"required"`
}
type ErrorUser struct {
	Name int `form:"name" json:"name" validate:"required"`
	Pass int `form:"pass" json:"pass" validate:"required"`
}
type Image struct {
	FileType string `form:"file_type" validate:"omitempty,eq=image/bmp|eq=image/dib|eq=image/jpeg|eq=image/jp2|eq=image/png|eq=image/webp|eq=image/x-portable-anymap|eq=image/x-portable-bitmap|eq=image/x-portable-graymap|eq=image/x-portable-pixmap|eq=image/x-cmu-raster|eq=image/tiff|eq=image/gif"`
}

var fileContent = []string{"image/bmp", "image/dib", "image/jpeg", "image/jp2", "image/png", "image/webp", "image/x-portable-anymap", "image/x-portable-bitmap", "image/x-portable-graymap", "image/x-portable-pixmap", "image/x-cmu-raster", "image/tiff", "image/gif"}

func TestBaseHandlerNewBaseHandler(t *testing.T) {
	bh := NewBaseHTTPHandler(infrastructure.NewLogger().Log)
	assert.NotEmpty(t, bh)
}

func TestBaseHandlerParse(t *testing.T) {
	bh := NewBaseHTTPHandler(infrastructure.NewLogger().Log)

	// request param
	values := url.Values{
		"name": {"circle"},
		"pass": {"pass"},
	}
	body := bytes.NewBufferString(values.Encode())
	// request
	req, _ := http.NewRequest("POST", "/", body)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	user := User{}
	err := bh.Parse(req, &user)
	assert.Equal(t, "circle", user.Name)
	assert.Equal(t, "pass", user.Pass)
	assert.NoError(t, err)

	// r is nil.
	err = bh.Parse(nil, &user)
	assert.Error(t, err)
	// data is nil.
	err = bh.Parse(req, nil)
	assert.Error(t, err)
	// error struct.
	errorUser := ErrorUser{}
	err = bh.Parse(req, &errorUser)
	assert.Error(t, err)
}

func TestBaseHandlerParseForm(t *testing.T) {
	bh := NewBaseHTTPHandler(infrastructure.NewLogger().Log)
	// request
	req, _ := http.NewRequest("GET", "/?name=circle&pass=pass", nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	user := User{}
	bh.ParseForm(req, &user)
	assert.Equal(t, "circle", user.Name)

}

func TestBaseHandlerMultipartParse(t *testing.T) {
	bh := NewBaseHTTPHandler(infrastructure.NewLogger().Log)

	body := new(bytes.Buffer)
	w := multipart.NewWriter(body)
	w.WriteField("name", "circle")
	w.WriteField("pass", "pass")
	w.Close()
	req, _ := http.NewRequest("POST", "/", body)
	req.Header.Set("Content-Type", w.FormDataContentType())

	user := User{}
	err := bh.ParseMultipart(req, &user)
	assert.Equal(t, "circle", user.Name)
	assert.Equal(t, "pass", user.Pass)
	assert.NoError(t, err)

	// req is nil.
	err = bh.ParseMultipart(nil, &user)
	assert.Error(t, err)
	// user is nil.
	err = bh.ParseMultipart(req, nil)
	assert.Error(t, err)
}

func TestBaseHandlerSaveToFile(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		sampleFile := os.Getenv("FR_CIRCLE_API_DIR") + "/test/data/sample.txt"

		bh := NewBaseHTTPHandler(infrastructure.NewLogger().Log)

		body := new(bytes.Buffer)
		w := multipart.NewWriter(body)
		utils.MultipartFileWriter(w, "text_file", sampleFile)
		w.Close()
		req, _ := http.NewRequest("POST", "/", body)
		req.Header.Set("Content-Type", w.FormDataContentType())

		bh.SaveToFile(req, "text_file", "dst_sample.txt")

		buf, _ := ioutil.ReadFile("dst_sample.txt")

		assert.Equal(t, "sample text", string(buf))

		os.Remove("dst_sample.txt")
	})
	t.Run("failed by nothing formfile", func(t *testing.T) {
		bh := NewBaseHTTPHandler(infrastructure.NewLogger().Log)

		req, _ := http.NewRequest("POST", "/", nil)

		err := bh.SaveToFile(req, "", "dst_sample.txt")

		assert.Error(t, err)
	})
	t.Run("failed by nothing tofile", func(t *testing.T) {
		sampleFile := os.Getenv("FR_CIRCLE_API_DIR") + "/test/data/sample.txt"

		bh := NewBaseHTTPHandler(infrastructure.NewLogger().Log)

		body := new(bytes.Buffer)
		w := multipart.NewWriter(body)
		utils.MultipartFileWriter(w, "text_file", sampleFile)
		w.Close()
		req, _ := http.NewRequest("POST", "/", body)
		req.Header.Set("Content-Type", w.FormDataContentType())

		err := bh.SaveToFile(req, "text_file", "")

		assert.Error(t, err)
	})
	t.Run("failed by nil", func(t *testing.T) {
		sampleFile := os.Getenv("FR_CIRCLE_API_DIR") + "/test/data/sample.txt"

		bh := NewBaseHTTPHandler(infrastructure.NewLogger().Log)

		body := new(bytes.Buffer)
		w := multipart.NewWriter(body)
		utils.MultipartFileWriter(w, "text_file", sampleFile)
		w.Close()
		req, _ := http.NewRequest("POST", "/", body)
		req.Header.Set("Content-Type", w.FormDataContentType())

		err := bh.SaveToFile(nil, "text_file", "")

		assert.Error(t, err)
	})
}

func TestBaseHandlerGetRandomTempFileName(t *testing.T) {
	bh := NewBaseHTTPHandler(infrastructure.NewLogger().Log)
	name, err := bh.GetRandomTempFileName("prefix-", "sample.txt")
	assert.Regexp(t, "prefix-", name)
	assert.NoError(t, err)
}

func TestBaseHandlerGetRandomFileName(t *testing.T) {
	bh := NewBaseHTTPHandler(infrastructure.NewLogger().Log)
	fileName, err := bh.GetRandomFileName("file-", "example.jpg")
	assert.Regexp(t, "file-", fileName)
	assert.NoError(t, err)
}

func TestBaseHandlerResponseJson(t *testing.T) {
	user := User{
		"circle",
		"password",
	}

	bh := NewBaseHTTPHandler(infrastructure.NewLogger().Log)

	// server
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bh.ResponseJSON(w, user)
		return
	})
	// httpTest newServer
	ts := httptest.NewServer(handler)
	defer ts.Close()

	// client
	res, err := http.Get(ts.URL)
	if err != nil {
		return
	}
	defer res.Body.Close()

	buf, _ := ioutil.ReadAll(res.Body)
	resUser := User{}
	json.Unmarshal(buf, &resUser)
	assert.Equal(t, "application/json", res.Header.Get("Content-Type"))
	assert.Equal(t, user.Name, resUser.Name)
	assert.Equal(t, user.Pass, resUser.Pass)
}

func TestBaseHandlerStatusRedirect(t *testing.T) {
	bh := NewBaseHTTPHandler(infrastructure.NewLogger().Log)

	// handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bh.StatusRedirect(w, "/redirect")
		return
	})

	// router
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	// response writer
	rec := httptest.NewRecorder()
	// new request
	req, _ := http.NewRequest("GET", "/", nil)
	// request
	mux.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusFound, rec.Code)
	path := rec.Header().Get("Location")
	assert.Equal(t, "/redirect", path)
}

func TestBaseHandlerStatusBadRequest(t *testing.T) {
	user := User{
		"circle",
		"password",
	}

	bh := NewBaseHTTPHandler(infrastructure.NewLogger().Log)

	// handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bh.StatusBadRequest(w, user)
		return
	})

	// response writer
	rec := httptest.NewRecorder()
	// new request
	req, _ := http.NewRequest("GET", "/", nil)
	// request
	handler(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))

	resUser := User{}
	json.Unmarshal(rec.Body.Bytes(), &resUser)
	assert.NotEmpty(t, user)
	assert.Equal(t, user.Name, resUser.Name)
	assert.Equal(t, user.Pass, resUser.Pass)
}

func TestBaseHandlerStatusNotFoundRequest(t *testing.T) {
	user := User{
		"circle",
		"password",
	}

	bh := NewBaseHTTPHandler(infrastructure.NewLogger().Log)

	t.Run("response json", func(t *testing.T) {
		// handler
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			bh.StatusNotFoundRequest(w, user)
			return
		})

		// response writer
		rec := httptest.NewRecorder()
		// new request
		req, _ := http.NewRequest("GET", "/", nil)
		// request
		handler(rec, req)

		assert.Equal(t, http.StatusNotFound, rec.Code)
		assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))
		resUser := User{}
		json.Unmarshal(rec.Body.Bytes(), &resUser)
		assert.Equal(t, user.Name, resUser.Name)
		assert.Equal(t, user.Pass, resUser.Pass)
	})
	t.Run("no response json.", func(t *testing.T) {
		// handler
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			bh.StatusNotFoundRequest(w, nil)
			return
		})

		// response writer
		rec := httptest.NewRecorder()
		// new request
		req, _ := http.NewRequest("GET", "/", nil)
		// request
		handler(rec, req)
		assert.Equal(t, http.StatusNotFound, rec.Code)
		assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))
		assert.Nil(t, rec.Body.Bytes())
	})
}

func TestBaseHandlerStatusServerError(t *testing.T) {
	t.Run("status server error", func(t *testing.T) {
		bh := NewBaseHTTPHandler(infrastructure.NewLogger().Log)

		// handler
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			bh.StatusServerError(w, nil)
			return
		})

		// response writer
		rec := httptest.NewRecorder()
		// new request
		req, _ := http.NewRequest("GET", "/", nil)
		// request
		handler(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})

	t.Run("status server error with body", func(t *testing.T) {
		user := User{
			"circle",
			"password",
		}

		bh := NewBaseHTTPHandler(infrastructure.NewLogger().Log)

		// handler
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			bh.StatusServerError(w, user)
			return
		})

		// response writer
		rec := httptest.NewRecorder()
		// new request
		req, _ := http.NewRequest("GET", "/", nil)
		// request
		handler(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))

		resUser := User{}
		json.Unmarshal(rec.Body.Bytes(), &resUser)
		assert.Equal(t, user.Name, resUser.Name)
		assert.Equal(t, user.Pass, resUser.Pass)
	})
}

func TestBaseHandlerGetTranslaterFunc(t *testing.T) {
	bh := NewBaseHTTPHandler(infrastructure.NewLogger().Log)

	translation := infrastructure.NewTranslation()

	// server
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		translaterFunc := bh.GetTranslaterFunc(r)
		assert.NotEmpty(t, translaterFunc)
		return
	})
	// router
	mux := chi.NewRouter()
	mux.Use(translation.Middleware.Middleware)
	mux.HandleFunc("/", handler)

	// response writer
	rec := httptest.NewRecorder()
	// new request
	req, _ := http.NewRequest("GET", "/", nil)
	// request
	mux.ServeHTTP(rec, req)
}

func TestBaseHandlerGetFileHeaderContentType(t *testing.T) {
	t.Run("get file header content type success", func(t *testing.T) {
		imageFile := os.Getenv("FR_CIRCLE_API_DIR") + "/test/image/W_ULD_171201_01.jpg"

		bh := NewBaseHTTPHandler(infrastructure.NewLogger().Log)

		body := new(bytes.Buffer)
		w := multipart.NewWriter(body)
		utils.MultipartFileWriter(w, "image_file", imageFile)
		w.Close()
		req, _ := http.NewRequest("POST", "/", body)
		req.Header.Set("Content-Type", w.FormDataContentType())

		f, _, _ := req.FormFile("image_file")
		defer f.Close()
		fileContentType, _ := bh.GetFileHeaderContentType(f)
		assert.Equal(t, "image/jpeg", fileContentType)
	})
	t.Run("arguments is nil", func(t *testing.T) {
		bh := NewBaseHTTPHandler(infrastructure.NewLogger().Log)
		f, err := bh.GetFileHeaderContentType(nil)
		assert.Error(t, err)
		assert.Empty(t, f)
	})
}
func TestBaseHandlerValidatorFunc(t *testing.T) {
	t.Run("validate success", func(t *testing.T) {
		user := User{
			Name: "username",
			Pass: "pass",
		}
		// base handler
		bh := NewBaseHTTPHandler(infrastructure.NewLogger().Log)
		// response write
		w := httptest.NewRecorder()

		err := bh.Validate(user, w)
		assert.NoError(t, err)
	})
	t.Run("validate faild", func(t *testing.T) {
		user := User{
			Pass: "password",
		}
		// base handler
		bh := NewBaseHTTPHandler(infrastructure.NewLogger().Log)
		// response writer
		w := httptest.NewRecorder()

		err := bh.Validate(user, w)
		assert.Error(t, err)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
	t.Run("validate faild with file text/css", func(t *testing.T) {
		image := Image{
			FileType: "text/css",
		}
		bh := NewBaseHTTPHandler(infrastructure.NewLogger().Log)
		w := httptest.NewRecorder()

		err := bh.Validate(image, w)
		assert.Error(t, err)
		assert.Equal(t, http.StatusUnsupportedMediaType, w.Code)
	})
	for _, v := range fileContent {
		t.Run("validate success with image valid "+v, func(t *testing.T) {
			image := Image{
				FileType: v,
			}
			bh := NewBaseHTTPHandler(infrastructure.NewLogger().Log)
			w := httptest.NewRecorder()
			err := bh.Validate(image, w)
			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, w.Code)
		})
	}
}

func TestBaseHandlerStatusUnsupportedMediaTypeRequest(t *testing.T) {
	image := Image{
		FileType: "text/css",
	}
	bh := NewBaseHTTPHandler(infrastructure.NewLogger().Log)

	t.Run("response json", func(t *testing.T) {
		// handler
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			bh.StatusUnsupportedMediaTypeRequest(w, image)
			return
		})

		// response writer
		rec := httptest.NewRecorder()
		// new request
		req, _ := http.NewRequest("GET", "/", nil)
		// request
		handler(rec, req)

		assert.Equal(t, http.StatusUnsupportedMediaType, rec.Code)
		assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))
		resImage := Image{}
		json.Unmarshal(rec.Body.Bytes(), &resImage)
		assert.Equal(t, image.FileType, resImage.FileType)
	})
	t.Run("no response json.", func(t *testing.T) {
		// handler
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			bh.StatusUnsupportedMediaTypeRequest(w, nil)
			return
		})

		// response writer
		rec := httptest.NewRecorder()
		// new request
		req, _ := http.NewRequest("GET", "/", nil)
		// request
		handler(rec, req)
		assert.Equal(t, http.StatusUnsupportedMediaType, rec.Code)
		assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))
		assert.Nil(t, rec.Body.Bytes())
	})
}
