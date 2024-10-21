package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
	"image"
	_ "image/gif"  // Import image formats to support GIF, JPEG, and PNG
	_ "image/jpeg" // We use underscore imports to register the image formats
	_ "image/png"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

type ApiError struct {
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

func HandleFile(c *gin.Context) (string, string) {
	path, err := filepath.Abs("./")

	file, _ := c.FormFile("file")
	filename := filepath.Base(file.Filename)
	var basePath = filepath.Join(path, "assets")

	if _, err = os.Stat(basePath); os.IsNotExist(err) {
		var dirMod uint64
		if dirMod, err = strconv.ParseUint("0775", 8, 32); err == nil {
			err = os.Mkdir(basePath, os.FileMode(dirMod))
		}
	}
	dst := filepath.Join(basePath, filename)

	// Upload the file to specific dst.
	c.SaveUploadedFile(file, dst)

	return dst, filename
}

func GetFileMimeType(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Read the first 512 bytes of the file
	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use http.DetectContentType to determine the file's MIME type
	mimeType := http.DetectContentType(buffer)
	return mimeType, nil
}

func GetFileSize(filename string) (int64, error) {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return 0, err
	}

	// Use the Size() method to get the file size in bytes
	return fileInfo.Size(), nil
}

func IsImageMimeType(mimeType string) bool {
	return mimeType == "image/jpeg" || mimeType == "image/png" || mimeType == "image/gif"
}

func GetImageDimensions(filename string) (int, int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()

	// Decode the image to get the dimensions
	img, _, err := image.Decode(file)
	if err != nil {
		return 0, 0, err
	}

	// Get the width and height of the image
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	return width, height, nil
}

func HandleError(err error) (int, ApiError) {
	var response ApiError

	var code int

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		code = http.StatusNotFound
	case errors.Is(err, gorm.ErrInvalidTransaction):
		code = http.StatusInternalServerError
	case errors.Is(err, gorm.ErrNotImplemented):
		code = http.StatusNotImplemented
	case errors.Is(err, gorm.ErrMissingWhereClause):
		code = http.StatusBadRequest
	case errors.Is(err, gorm.ErrUnsupportedRelation):
		code = http.StatusBadRequest
	case errors.Is(err, gorm.ErrPrimaryKeyRequired):
		code = http.StatusBadRequest
	case errors.Is(err, gorm.ErrModelValueRequired):
		code = http.StatusBadRequest
	case errors.Is(err, gorm.ErrModelAccessibleFieldsRequired):
		code = http.StatusBadRequest
	case errors.Is(err, gorm.ErrSubQueryRequired):
		code = http.StatusBadRequest
	case errors.Is(err, gorm.ErrInvalidData):
		code = http.StatusBadRequest
	case errors.Is(err, gorm.ErrUnsupportedDriver):
		code = http.StatusInternalServerError
	case errors.Is(err, gorm.ErrRegistered):
		code = http.StatusConflict
	case errors.Is(err, gorm.ErrInvalidField):
		code = http.StatusBadRequest
	case errors.Is(err, gorm.ErrEmptySlice):
		code = http.StatusBadRequest
	case errors.Is(err, gorm.ErrDryRunModeUnsupported):
		code = http.StatusNotImplemented
	case errors.Is(err, gorm.ErrInvalidDB):
		code = http.StatusInternalServerError
	case errors.Is(err, gorm.ErrInvalidValue):
		code = http.StatusBadRequest
	case errors.Is(err, gorm.ErrInvalidValueOfLength):
		code = http.StatusBadRequest
	case errors.Is(err, gorm.ErrPreloadNotAllowed):
		code = http.StatusBadRequest
	case errors.Is(err, gorm.ErrDuplicatedKey):
		code = http.StatusConflict
	case errors.Is(err, gorm.ErrForeignKeyViolated):
		code = http.StatusConflict
	case errors.Is(err, gorm.ErrCheckConstraintViolated):
		code = http.StatusConflict
	default:
		code = http.StatusInternalServerError
	}
	if pgErr, ok := err.(*pgconn.PgError); ok {
		switch pgErr.Code {
		case "23505": // unique_violation
			code = http.StatusConflict
		case "23503": // foreign_key_violation
			code = http.StatusConflict
		case "23502": // not_null_violation
			code = http.StatusBadRequest
		case "23514": // check_violation
			code = http.StatusConflict
		case "42703": // undefined_column
			code = http.StatusBadRequest
		case "42883": // undefined_function
			code = http.StatusBadRequest
		case "42601": // syntax_error
			code = http.StatusBadRequest
		case "23508": // exclusion_violation
			code = http.StatusConflict
		case "22P02": // invalid_text_representation (e.g. invalid input syntax for type)
			code = http.StatusBadRequest
		case "22007": // invalid_datetime_format
			code = http.StatusBadRequest
		case "42P01": // undefined_table
			code = http.StatusNotFound
		case "42P07": // duplicate_table
			code = http.StatusConflict
		case "40P01": // deadlock_detected
			code = http.StatusInternalServerError
		default:
			code = http.StatusInternalServerError
		}
	}

	if pgErr, ok := err.(*pgconn.PgError); ok {
		response.Message = pgErr.Message
		response.Detail = pgErr.Detail
	}

	return code, response
}

func Home(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Ok")
}
