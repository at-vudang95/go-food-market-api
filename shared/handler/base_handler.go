package handler

import (
	"crypto/rand"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"encoding/hex"
	"path/filepath"

	"io"

	"github.com/at-vudang95/go-food-market-api/infrastructure"
	"github.com/at-vudang95/go-food-market-api/shared/utils"
	"github.com/at-vudang95/go-food-market-api/shared/validator"
	"github.com/go-playground/form"

	"github.com/nicksnyder/go-i18n/i18n/bundle"
	"github.com/sirupsen/logrus"
	v "gopkg.in/go-playground/validator.v9"
)

const (
	// FilePerm is file permission.
	FilePerm = 0666
)

// BaseHTTPHandler base handler struct.
type BaseHTTPHandler struct {
	Logger *logrus.Logger
}

// Parse parse form struct.
// https://github.com/go-playground/form
func (h *BaseHTTPHandler) Parse(r *http.Request, i interface{}) error {
	if r == nil {
		return utils.ErrorsNew("first argument is nil")
	}
	if i == nil {
		return utils.ErrorsNew("second argument is nil")
	}
	// mapping post to struct.
	err := r.ParseForm()
	if err != nil {
		return utils.ErrorsWrap(err, "can't ParseForm error")
	}
	decoder := form.NewDecoder()
	err = decoder.Decode(&i, r.PostForm)
	if err != nil {
		return utils.ErrorsWrap(err, "can't decoder.Decode error")
	}
	return nil
}

// ParseMultipart parse form struct.
// https://github.com/go-playground/form
func (h *BaseHTTPHandler) ParseMultipart(r *http.Request, i interface{}) error {
	if r == nil {
		return utils.ErrorsNew("first argument is nil")
	}
	if i == nil {
		return utils.ErrorsNew("second argument is nil")
	}
	maxMemory := infrastructure.GetConfigInt64("multipart.maxmemory")
	err := r.ParseMultipartForm(maxMemory)
	if err != nil {
		return utils.ErrorsWrap(err, "can't ParseMultipartForm error")
	}
	// mapping post to struct.
	decoder := form.NewDecoder()
	err = decoder.Decode(&i, r.PostForm)
	if err != nil {
		return utils.ErrorsWrap(err, "can't decoder.Decode error")
	}
	return nil
}

// ParseForm  form struct.
// https://github.com/go-playground/form
func (h *BaseHTTPHandler) ParseForm(r *http.Request, i interface{}) error {
	// mapping post to struct.
	err := r.ParseForm()
	if err != nil {
		return utils.ErrorsWrap(err, "can't ParseForm error")
	}
	decoder := form.NewDecoder()
	err = decoder.Decode(&i, r.Form)
	if err != nil {
		return utils.ErrorsWrap(err, "can't decoder.Decode error")
	}
	return nil
}

// SaveToFile saves uploaded file to new path.
// it only operates the first one of mutil-upload form file field.
func (h *BaseHTTPHandler) SaveToFile(r *http.Request, fromfile, tofile string) (err error) {
	if r == nil {
		return utils.ErrorsNew("first argument is nil")
	}
	file, _, err := r.FormFile(fromfile)
	if err != nil {
		return utils.ErrorsWrap(err, "FormFile error")
	}
	defer func() {
		if ferr := file.Close(); err == nil && ferr != nil {
			err = utils.ErrorsWrap(ferr, "can't close file")
		}
	}()

	f, err := os.OpenFile(tofile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, FilePerm)
	if err != nil {
		return utils.ErrorsWrap(err, "can't open file")
	}
	defer func() {
		if ferr := f.Close(); err == nil && ferr != nil {
			err = utils.ErrorsWrap(ferr, "can't close file")
		}
	}()
	_, err = io.Copy(f, file)
	if err != nil {
		return utils.ErrorsWrap(err, "can't copy file")
	}
	return nil
}

// GetRandomTempFileName get temp random filename.
func (h *BaseHTTPHandler) GetRandomTempFileName(prefix, filename string) (string, error) {
	randBytes := make([]byte, 16)
	_, err := rand.Read(randBytes)
	if err != nil {
		return "", utils.ErrorsWrap(err, "can't create rand")
	}
	extension := filepath.Ext(filename)
	return filepath.Join(os.TempDir(), prefix+hex.EncodeToString(randBytes)+extension), nil
}

// GetRandomFileName get temp random filename.
func (h *BaseHTTPHandler) GetRandomFileName(prefix, filename string) (string, error) {
	randBytes := make([]byte, 16)
	_, err := rand.Read(randBytes)
	if err != nil {
		return "", utils.ErrorsWrap(err, "can't create rand")
	}
	extension := filepath.Ext(filename)
	curentTime := time.Now().Unix()
	name := hex.EncodeToString(randBytes) + strconv.FormatInt(curentTime, 10)
	return prefix + name + extension, nil
}

// ResponseJSON responses status code 200 and json.
func (h *BaseHTTPHandler) ResponseJSON(w http.ResponseWriter, data interface{}) {
	// status code 200
	utils.ResponseJSON(w, http.StatusOK, data)
}

// StatusRedirect responses status code 302 and redirect.
func (h *BaseHTTPHandler) StatusRedirect(w http.ResponseWriter, url string) {
	w.Header().Set("Location", url)
	// status code 302
	w.WriteHeader(http.StatusFound)
}

// StatusBadRequest responses status code 400 and json.
func (h *BaseHTTPHandler) StatusBadRequest(w http.ResponseWriter, data interface{}) {
	// status code 400
	utils.ResponseJSON(w, http.StatusBadRequest, data)
}

// StatusNotFoundRequest responses status code 404 and json.
func (h *BaseHTTPHandler) StatusNotFoundRequest(w http.ResponseWriter, data interface{}) {
	// status code 404
	utils.ResponseJSON(w, http.StatusNotFound, data)
}

// StatusUnsupportedMediaTypeRequest responses status code 415 and json.
func (h *BaseHTTPHandler) StatusUnsupportedMediaTypeRequest(w http.ResponseWriter, data interface{}) {
	// status code 415
	utils.ResponseJSON(w, http.StatusUnsupportedMediaType, data)
}

// StatusServerError responses 500.
func (h *BaseHTTPHandler) StatusServerError(w http.ResponseWriter, data interface{}) {
	// status code 500
	utils.ResponseJSON(w, http.StatusInternalServerError, data)
}

// GetTranslaterFunc returns i18n.TranslateFunc.
// This function is necessary for translation.
func (h *BaseHTTPHandler) GetTranslaterFunc(r *http.Request) bundle.TranslateFunc {
	return r.Context().Value("i18nTfunc").(bundle.TranslateFunc)
}

// GetFileHeaderContentType return the file content-type by read 128 first bytes
func (h *BaseHTTPHandler) GetFileHeaderContentType(file multipart.File) (string, error) {
	if file == nil {
		return "", utils.ErrorsNew("file variable is nil")
	}
	fileHeader := make([]byte, 128)
	_, err := file.Read(fileHeader)
	if err == io.EOF {
		fileHeader, err = ioutil.ReadAll(file)
		if err != nil {
			return "", utils.ErrorsNew("can't read file header")
		}
	}
	return strings.ToLower(utils.DetectFileContentType(fileHeader)), err
}

// Validate func validate request of user
func (h *BaseHTTPHandler) Validate(req interface{}, w http.ResponseWriter) error {
	// validator.v9 new and add custom validators.
	mValidator := validator.New()
	err := mValidator.Struct(req)
	if err != nil {
		if _, ok := err.(*v.InvalidValidationError); ok {
			h.Logger.Error("Invalid validation error.")
			common := utils.CommonResponse{Message: "Internal server error response", Errors: []string{}}
			h.StatusServerError(w, common)
			return err

		}
		var errors []string
		for _, errV := range err.(v.ValidationErrors) {
			if errV.StructField() == "FileType" {
				common := utils.CommonResponse{Message: "Unsupported media type error response.", Errors: []string{"Unsupported media type"}}
				h.StatusUnsupportedMediaTypeRequest(w, common)
				return err
			}
			errors = append(errors, errV.Field()+" is error.")
		}
		common := utils.CommonResponse{Message: "validation error.", Errors: errors}
		h.StatusBadRequest(w, common)
	}
	return err
}

// NewBaseHTTPHandler returns BaseHTTPHandler instance.
func NewBaseHTTPHandler(logger *logrus.Logger) *BaseHTTPHandler {
	return &BaseHTTPHandler{Logger: logger}
}
