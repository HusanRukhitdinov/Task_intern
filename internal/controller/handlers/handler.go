package handlers

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"intern/pkg/logger"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var (
	defaultPage  = "1"
	defaultLimit = "10"
)

var (
	errorInnNotFound = errors.New("inn not found")

	errorWrongGtinLen = errors.New("wrong gtin length")
	wrongID           = errors.New("invalid input syntax for type uuid: string (SQLSTATE 22P02)")

	errorBadRequest        = errors.New("bad request")
	errorInternal          = errors.New("internal error")
	errorWrongCredentials  = errors.New("wrong credentials")
	errorAccessDenied      = errors.New("access denied")
	ErrPermissionDenied    = errors.New("permission denied")
	emailErr               = errors.New("email is not valid")
	timeErr                = errors.New("time parse error")
	errNotFoundID          = errors.New("no rows in result set")
	idError                = errors.New("SQLSTATE 22P02")
	referenceIDError       = errors.New("SQLSTATE 23503")
	notFound               = "Not found"
	wrongPassword          = "Wrong credentials"
	emailResponse          = "Your email is not Strong "
	timeResponse           = "Time syntax Input"
	errAdminUserNameInsert = errors.New("ERROR: new row for relation \"answers\\\" violates check constraint \"check_solution_required\" (SQLSTATE 23514)")
)

func (h *Handler) handleError(c *gin.Context, err error) bool {
	if err == nil {
		return false
	}

	h.log.Error("ERROR:-~~~~~~~~~~~~~``", logger.Error(err))
	fmt.Println("---------", err.Error())
	switch err.Error() {
	case errorWrongGtinLen.Error():
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	case errAdminUserNameInsert.Error():
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bu username da admin bor"})

		c.JSON(http.StatusBadRequest, gin.H{"error": "Seller Bir xill nomdagi product yaratish mumkin emas"})
	case emailErr.Error():
		c.JSON(http.StatusBadRequest, gin.H{"error": emailResponse})
	case wrongID.Error():
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("ID wrong is that input")})
	case errorBadRequest.Error():
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	case errorWrongCredentials.Error():
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	case bcrypt.ErrMismatchedHashAndPassword.Error():
		c.JSON(http.StatusUnauthorized, gin.H{"error": wrongPassword})
	case errorAccessDenied.Error():
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
	case sql.ErrNoRows.Error():
		c.JSON(http.StatusNotFound, gin.H{"error": notFound})
	case errorInnNotFound.Error():
		c.JSON(http.StatusNotFound, gin.H{"error": errorInnNotFound.Error()})
	case errorInternal.Error():
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	return true
}

func (h *Handler) makeContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Hour)
	return ctx, cancel
}

func (h *Handler) compareHashedPassword(hashedPassword []byte, password []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	if err != nil {
		return err
	}

	return nil
}

func (h *Handler) hashingPassword(password []byte) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}
