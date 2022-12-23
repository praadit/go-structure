package middleware

// import (
// 	"fmt"
// 	"go-best-practice/internal/constants"
// 	"go-best-practice/internal/exception"
// 	"strings"
// 	"time"

// 	"github.com/labstack/echo/v4"
// )

// func DoubleRequestMiddleware(cache services.RedisServiceInterface) echo.MiddlewareFunc {
// 	return func(next echo.HandlerFunc) echo.HandlerFunc {
// 		return echo.HandlerFunc(func(c echo.Context) error {
// 			if c.Request().Method == "GET" {
// 				return next(c)
// 			}
// 			header := c.Request().Header["Authorization"]
// 			if header == nil {
// 				return next(c)
// 			}

// 			if header[0] == "" {
// 				return next(c)
// 			}

// 			token := strings.Split(header[0], "Bearer ")[1]
// 			if token == "" {
// 				return next(c)
// 			}

// 			endpoint := c.Request().URL.String()

// 			cacheName := fmt.Sprintf("request:%s:%s", token, endpoint)
// 			exist := cache.GetCacheValue(cacheName)
// 			if exist != "" {
// 				panic(&exception.BadRequestError{
// 					Code:    constants.ERR_REQ_DoubleRequest,
// 					Message: "Mohon menunggu beberapa saat sebelum melakukan aksi yang sama.",
// 				})
// 			}

// 			cache.SetCacheString(cacheName, "exist", 2*time.Second)
// 			return next(c)
// 		})
// 	}
// }
