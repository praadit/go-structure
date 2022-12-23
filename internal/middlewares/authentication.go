package middleware

// import (
// 	"go-best-practice/internal/exception"
// 	"go-best-practice/internal/utilities"
// 	"strings"

// 	"github.com/dgrijalva/jwt-go"
// 	"github.com/labstack/echo/v4"
// )

// func authenticate(ctx echo.Context) (*dto.AuthCustomClaims, error) {
// 	authToken := ctx.Request().Header.Get("Authorization")
// 	if authToken == "" {
// 		return nil, &exception.UnauthorizedError{
// 			Code:    constants.ERR_REQ_AutenticationFailed,
// 			Message: "Unauthorized",
// 		}
// 	}
// 	jwtString := strings.Split(authToken, "Bearer ")[1]
// 	if jwtString == "" {
// 		return nil, &exception.UnauthorizedError{
// 			Code:    constants.ERR_REQ_AutenticationFailed,
// 			Message: "Unauthorized",
// 		}
// 	}

// 	claims := dto.AuthCustomClaims{}

// 	token, _ := jwt.ParseWithClaims(jwtString, &claims, func(token *jwt.Token) (interface{}, error) {
// 		return []byte(config.AppConfig[config.ApiSecretKey]), nil
// 	})

// 	if !token.Valid {
// 		return nil, &exception.UnauthorizedError{
// 			Code:    constants.ERR_REQ_AutenticationFailed,
// 			Message: "Unauthorized",
// 		}
// 	}

// 	return &claims, nil

// }

// func AdminOnly(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(ctx echo.Context) error {
// 		claims, err := authenticate(ctx)
// 		if err != nil {
// 			panic(err)
// 		}
// 		if claims.Role != constants.UserRole_Admin {
// 			panic(&exception.ForbiddendError{
// 				Code:    constants.ERR_REQ_AccessForbidden,
// 				Message: "Forbidden",
// 			})
// 		}
// 		ctx.Set(constants.CONTEXT_VAR_UserClaims, claims)
// 		return next(ctx)
// 	}
// }

// func PraktisiOnly(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(ctx echo.Context) error {
// 		claims, err := authenticate(ctx)
// 		if err != nil {
// 			panic(err)
// 		}
// 		if claims.Role != constants.UserRole_Praktisi {
// 			panic(&exception.ForbiddendError{
// 				Code:    constants.ERR_REQ_AccessForbidden,
// 				Message: "Forbidden",
// 			})
// 		}
// 		ctx.Set(constants.CONTEXT_VAR_UserClaims, claims)
// 		return next(ctx)
// 	}
// }

// func ReviewerOnly(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(ctx echo.Context) error {
// 		claims, err := authenticate(ctx)
// 		if err != nil {
// 			panic(err)
// 		}
// 		if claims.Role != constants.UserRole_Reviewer_Praktisi {
// 			panic(&exception.ForbiddendError{
// 				Code:    constants.ERR_REQ_AccessForbidden,
// 				Message: "Forbidden",
// 			})
// 		}
// 		ctx.Set(constants.CONTEXT_VAR_UserClaims, claims)
// 		return next(ctx)
// 	}
// }

// func ReviewerPTOnly(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(ctx echo.Context) error {
// 		claims, err := authenticate(ctx)
// 		if err != nil {
// 			panic(err)
// 		}
// 		if claims.Role != constants.UserRole_Reviewer_PT {
// 			panic(&exception.ForbiddendError{
// 				Code:    constants.ERR_REQ_AccessForbidden,
// 				Message: "Forbidden",
// 			})
// 		}
// 		ctx.Set(constants.CONTEXT_VAR_UserClaims, claims)
// 		return next(ctx)
// 	}
// }

// func PTOnly(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(ctx echo.Context) error {
// 		claims, err := authenticate(ctx)
// 		if err != nil {
// 			panic(err)
// 		}

// 		if claims.Role != constants.UserRole_PT {
// 			panic(&exception.ForbiddendError{
// 				Code:    constants.ERR_REQ_AccessForbidden,
// 				Message: "Forbidden",
// 			})
// 		}
// 		ctx.Set(constants.CONTEXT_VAR_UserClaims, claims)
// 		return next(ctx)
// 	}
// }

// func DosenOnly(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(ctx echo.Context) error {
// 		claims, err := authenticate(ctx)
// 		if err != nil {
// 			panic(err)
// 		}

// 		if claims.Role != constants.UserRole_Dosen {
// 			panic(&exception.ForbiddendError{
// 				Code:    constants.ERR_REQ_AccessForbidden,
// 				Message: "Forbidden",
// 			})
// 		}
// 		ctx.Set(constants.CONTEXT_VAR_UserClaims, claims)
// 		return next(ctx)
// 	}
// }

// func AllUser(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(ctx echo.Context) error {
// 		claims, err := authenticate(ctx)
// 		if err != nil {
// 			panic(err)
// 		}
// 		ctx.Set(constants.CONTEXT_VAR_UserClaims, claims)
// 		return next(ctx)
// 	}
// }

// func SuperOnly(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(ctx echo.Context) error {
// 		claims, err := authenticate(ctx)
// 		if err != nil {
// 			panic(err)
// 		}

// 		if claims.Role != constants.UserRole_SuperAdmin {
// 			panic(&exception.ForbiddendError{
// 				Code:    constants.ERR_REQ_AccessForbidden,
// 				Message: "Forbidden",
// 			})
// 		}
// 		ctx.Set(constants.CONTEXT_VAR_UserClaims, claims)
// 		return next(ctx)
// 	}
// }

// func ViewerOnly(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(ctx echo.Context) error {
// 		claims, err := authenticate(ctx)
// 		if err != nil {
// 			panic(err)
// 		}

// 		if claims.Role != constants.UserRole_Reporter {
// 			panic(&exception.ForbiddendError{
// 				Code:    constants.ERR_REQ_AccessForbidden,
// 				Message: "Forbidden",
// 			})
// 		}
// 		ctx.Set(constants.CONTEXT_VAR_UserClaims, claims)
// 		return next(ctx)
// 	}
// }

// func MultipleRole(allowedRole []constants.UserRoleConst) echo.MiddlewareFunc {
// 	return func(next echo.HandlerFunc) echo.HandlerFunc {
// 		return func(ctx echo.Context) error {
// 			claims, err := authenticate(ctx)
// 			if err != nil {
// 				panic(err)
// 			}

// 			if !utilities.Contains(allowedRole, claims.Role) {
// 				panic(&exception.ForbiddendError{
// 					Code:    constants.ERR_REQ_AccessForbidden,
// 					Message: "Forbidden",
// 				})
// 			}
// 			ctx.Set(constants.CONTEXT_VAR_UserClaims, claims)
// 			return next(ctx)
// 		}
// 	}
// }
