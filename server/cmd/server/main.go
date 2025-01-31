package main

import (
	"log"
	"portal/loggerhelper"
	"portal/migrations"
	"portal/models"
	"portal/routing"

	"github.com/gin-gonic/gin"
	_ "github.com/go-pg/pg/v10/orm"
	"github.com/oiime/logrusbun"
	"github.com/pro-assistance-dev/sprob/config"
	helperPack "github.com/pro-assistance-dev/sprob/helper"
	"github.com/sirupsen/logrus"
)

func main() {
	conf, err := config.LoadConfig(".env")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	helper := helperPack.NewHelper(*conf)
	logger := loggerhelper.NewLogger()

	router := gin.New()
	router.Use(gin.Recovery())
	//
	router.Use(loggerhelper.LoggingMiddleware(logger))
	routing.Init(router, helper)
	helper.DB.DB.AddQueryHook(logrusbun.NewQueryHook(logrusbun.QueryHookOptions{Logger: logger, ErrorLevel: logrus.ErrorLevel, QueryLevel: logrus.DebugLevel}))
	helper.DB.DB.RegisterModel((*models.EventToSponsor)(nil))
	helper.DB.DB.RegisterModel((*models.EventToPlace)(nil))
	helper.DB.DB.RegisterModel((*models.EventToPerfom)(nil))
	helper.DB.DB.RegisterModel((*models.PerfomToSpeaker)(nil))
	helper.DB.DB.RegisterModel((*models.SessionToSpeaker)(nil))
	helper.DB.DB.RegisterModel((*models.ScheduleToPerfom)(nil))
	helper.Run(migrations.Init(), routing.Init)
}
