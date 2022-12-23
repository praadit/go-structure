package config

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitServer() *echo.Echo {
	e := echo.New()

	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize:         1 << 10, // 1 KB
		DisableStackAll:   true,
		DisablePrintStack: false,
		LogLevel:          1,
	}))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${method}][${status}] uri=${uri} time=${latency_human}\n",
	}))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization", "X-Requested-With", "*"},
		AllowCredentials: true,
		MaxAge:           int(12 * time.Hour),
	}))

	// e.HTTPErrorHandler = exception.ErrorHandler

	// redisService := services.NewRedisService(rdb)
	// e.Use(mid.DoubleRequestMiddleware(redisService))
	// e.Use(mid.MaintenanceFilter(db))
	// e.Use(mid.TransactionHandler(db))

	// r := router.NewRouter(e, rdb, gcsClient, pddikti)
	// r.Init()

	// pmCron := router.NewCron(db, rdb, gcsClient, pddikti)
	// if config.AppConfig[config.Env] == "debug" {
	// 	mailService := services.NewMailService()
	// 	redisService := services.NewRedisService(rdb)
	// 	pddiktiService := services.NewPddiktiService(pddikti, redisService)
	// 	gcsService := services.NewGcsUploadService(gcsClient)

	// 	confUsecase := usecase.NewConfigUsecase(redisService)
	// 	mailUsecase := usecase.NewMailUsecase(mailService)
	// 	rvptReportGen := usecase.NewReportRevPtGenerator(gcsService)

	// 	cronMail := usecase.NewMailCronUsecase(db, mailService)
	// 	pddiktiProdiCronUsecase := usecase.NewPddiktiProdiCronUsecase(db, pddiktiService)
	// 	cronPkApproval := usecase.NewCronPkApprovalUsecase(db, mailUsecase)
	// 	cronReportRvpt := usecase.NewReportRvptSchedulercron(db, rvptReportGen)
	// 	cronTutupKolab := usecase.NewForceCloseCronUsecase(db, confUsecase)

	// 	e.GET("/api/cron/mail", cronMail.SendScheduledMailEndpoint, mid.SuperOnly)
	// 	e.GET("/api/cron/prodi", pddiktiProdiCronUsecase.ExecuteProdiSchedulerEndpoint, mid.SuperOnly)
	// 	e.GET("/api/cron/notify/approved", cronPkApproval.NotifyApprovedPraktisiEndpoint, mid.SuperOnly)
	// 	e.GET("/api/cron/notify/rejected", cronPkApproval.NotifyRejectedPraktisiEndpoint, mid.SuperOnly)
	// 	e.GET("/api/cron/notify/delete", cronPkApproval.DeleteRejectedColabEndpoint, mid.SuperOnly)
	// 	e.GET("/api/cron/report/rvpt", cronReportRvpt.ManualExecuteScheduler, mid.SuperOnly)
	// 	e.GET("/api/cron/lap/force-close", cronTutupKolab.ManualExecuteScheduler, mid.SuperOnly)
	// } else {
	// 	pmCron.Init()
	// }

	return e
}
