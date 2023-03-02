package errors

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Error struct {
	Message gin.H
	Code    int
}

func ErrorHandler(err error) *Error {
	error := &Error{}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		error.Message = gin.H{"error": "Record does not exist"}
		error.Code = http.StatusNotFound
	} else {
		error.Message = gin.H{"error": err.Error()}
		error.Code = http.StatusInternalServerError
	}

	return error
}
