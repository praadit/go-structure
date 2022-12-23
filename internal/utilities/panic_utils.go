package utilities

import (
	"go-best-practice/internal/constants"
	"go-best-practice/internal/exception"

	"github.com/jackc/pgconn"
)

func PanicIfError(err error, logMsg interface{}) {
	if err != nil {
		ExceptionError(err)
		HandlerPgError(err)
		panic(err)
	}
}

func ExceptionError(err error) {
	if _, ok := err.(*exception.MaintenanceError); ok {
		panic(err)
	}
	if _, ok := err.(*exception.UnauthorizedError); ok {
		panic(err)
	}
	if _, ok := err.(*exception.ForbiddendError); ok {
		panic(err)
	}
	if _, ok := err.(*exception.NotFoundError); ok {
		panic(err)
	}
	if _, ok := err.(*exception.ValidationError); ok {
		panic(err)
	}
	if _, ok := err.(*exception.BadRequestError); ok {
		panic(err)
	}
}

func HandlerPgError(err error) {
	if pgerr, ok := err.(*pgconn.PgError); ok {
		switch pgerr.Code {
		case "42703":
			throwPanicPgError(constants.ERR_REQ_RequestNotValid, "PGERR (42703) : Field tidak ditemukan")
			break
		case "22001":
			throwPanicPgError(constants.ERR_REQ_RequestNotValid, "PGERR (22001) : Value tidak valid")
			break
		case "08P01":
			throwPanicPgError(constants.ERR_REQ_RequestNotValid, "PGERR (08P01) : Format data tidak valid")
			break
		case "23503":
			throwPanicPgError(constants.ERR_REQ_RequestNotValid, "PGERR (23503) : Relasi/Foreign Key tidak Valid")
			break
		case "42601":
			throwPanicPgError(constants.ERR_REQ_RequestNotValid, "PGERR (42601) : Query sysntax error")
			break
		default:
			panic(err)
		}
	}
}

func throwPanicPgError(code string, mesage string) {
	panic(&exception.InternalError{
		Code:            code,
		InternalMessage: mesage,
	})
}
