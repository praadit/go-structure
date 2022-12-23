package middleware

// import (
// 	"errors"
// 	"go-best-practice/internal/exception"

// 	"github.com/labstack/echo/v4"
// 	"gorm.io/gorm"
// )

// func MaintenanceFilter(db *gorm.DB) echo.MiddlewareFunc {
// 	return func(next echo.HandlerFunc) echo.HandlerFunc {
// 		return echo.HandlerFunc(func(c echo.Context) error {
// 			conf := entity.SystemConfig{}
// 			err := db.Model(&entity.SystemConfig{}).Where("config", constants.SysConf_MaintenanceState).First(&conf).Error
// 			if err != nil {
// 				if errors.Is(err, gorm.ErrRecordNotFound) {
// 					next(c)
// 					return nil
// 				}
// 				c.Logger().Info("Getting error while reading a maintenance state. Skipping maintenance filter")
// 				next(c)
// 				return nil
// 			}

// 			if conf.Value == "true" {
// 				panic(&exception.MaintenanceError{
// 					Code:    constants.ERR_SYSTEM_UnderMaintenance,
// 					Message: "Under Maintenance",
// 				})
// 			}

// 			next(c)
// 			return nil
// 		})
// 	}
// }
