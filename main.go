package main

import (
	"fmt"
	"log"
	"qtk-store-api/component/appctx"
	docs "qtk-store-api/docs"
	"qtk-store-api/middleware"
	"qtk-store-api/module/author/authortransport"
	"qtk-store-api/module/category/categorytransport"
	"qtk-store-api/module/importnote/importnotetransport/ginimportnote"
	"qtk-store-api/module/inventorychecknote/inventorychecknotetransport/gininventorychecknote"
	"qtk-store-api/module/product/producttransport"
	"qtk-store-api/module/supplier/suppliertransport/ginsupplier"
	"qtk-store-api/module/user/usertransport/ginuser"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type appConfig struct {
	Port string
	Env  string

	DBUsername string
	DBPassword string
	DBHost     string
	DBDatabase string

	SecretKey string
}

// @title           QTK Store API
// @version         1.0

// @contact.name   Bui Vi Quoc
// @contact.url    https://www.faceproduct.com/bviquoc/
// @contact.email  21520095@gm.uit.edu.vn

// @host      localhost:8080
// @BasePath  /v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	cfg, err := loadConfig()
	if err != nil {
		log.Fatalln("Error when loading config:", err)
	}

	db, err := connectDatabase(cfg)
	if err != nil {
		log.Fatalln("Error when connecting to database:", err)
	}
	if cfg.Env == "dev" {
		db = db.Debug()
	}

	appCtx := appctx.NewAppContext(db, cfg.SecretKey)

	r := gin.Default()
	r.Use(middleware.Recover(appCtx))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	docs.SwaggerInfo.BasePath = "/v1"
	v1 := r.Group("/v1")
	{
		v1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

		authortransport.SetupRoutes(v1, appCtx)
		categorytransport.SetupRoutes(v1, appCtx)
		producttransport.SetupRoutes(v1, appCtx)
		ginimportnote.SetupRoutes(v1, appCtx)
		gininventorychecknote.SetupRoutes(v1, appCtx)
		ginsupplier.SetupRoutes(v1, appCtx)
		ginuser.SetupRoutes(v1, appCtx)
	}

	if err := r.Run(fmt.Sprintf(":%s", cfg.Port)); err != nil {
		log.Fatalln("Error running server:", err)
	}
}

func loadConfig() (*appConfig, error) {
	env, err := godotenv.Read()
	if err != nil {
		log.Fatalln("Error when loading .env", err)
	}

	return &appConfig{
		Port:       env["PORT"],
		Env:        env["ENVIRONMENT"],
		DBUsername: env["DB_USERNAME"],
		DBPassword: env["DB_PASSWORD"],
		DBHost:     env["DB_HOST"],
		DBDatabase: env["DB_DATABASE"],
		SecretKey:  env["SECRET_KEY"],
	}, nil
}

func connectDatabase(cfg *appConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.DBUsername, cfg.DBPassword, cfg.DBHost, cfg.DBDatabase)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return db.Debug(), nil
}
