package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"test-majoo-new/auth"
	"test-majoo-new/handler"
	"test-majoo-new/helper"
	user "test-majoo-new/modules/User"
	"test-majoo-new/modules/area"
	"test-majoo-new/modules/merchant"
	"test-majoo-new/modules/outlet"
	"test-majoo-new/modules/report"
	"test-majoo-new/modules/transaction"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	// GlobalCache *bigcache.BigCache
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("sad .env file found")
	}
}
func main() {
	time.LoadLocation("Asia/Jakarta")

	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file. Make sure .env file is exists!")
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// koneksi mengunkana sql biasa
	dbManual, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer dbManual.Close()
	// end koneksi

	db.Debug().AutoMigrate(
		&user.User{},
		&merchant.Merchant{},
		&outlet.Outlet{},
		&transaction.Transaction{},
		&area.Area{},
	)

	// repository
	userRepository := user.NewRepository(db)
	merchantRepository := merchant.NewRepository(db, dbManual)
	outletRepository := outlet.NewRepository(db, dbManual)

	transactionRepository := transaction.NewRepository(db, dbManual)
	reportRepository := report.NewRepository(db, dbManual)

	areaRepository := area.NewRepository(db, dbManual)

	// service
	userService := user.NewService(userRepository)
	merchantService := merchant.NewService(merchantRepository)
	outletService := outlet.NewService(outletRepository, merchantRepository)

	transactionService := transaction.NewService(transactionRepository)
	reportService := report.NewService(reportRepository)

	areaService := area.NewService(areaRepository)

	// auth
	authService := auth.NewService()

	// handler
	userHandler := handler.NewUserHandler(userService, authService)
	merchantHandler := handler.NewMerchantHandler(merchantService)
	outletHandler := handler.NewOutletHandler(outletService)

	transactionHandler := handler.NewTransactionHandler(transactionService)
	fmt.Println(transactionHandler)
	reportHandler := handler.NewReportHandler(reportService)

	areaHandler := handler.NewAreaHandler(areaService)

	// route
	router := gin.Default()
	// cors
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	router.Use(cors.New(corsConfig)) // CORS configuraion

	api := router.Group("/api/v1")

	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)

	api.GET("/users", authMiddleware(authService, userService), userHandler.GetUserAll)
	api.GET("/users/fetch", authMiddleware(authService, userService), userHandler.FetchUser)
	api.POST("/users/update/:uuid", authMiddleware(authService, userService), userHandler.UpdateUser)
	api.POST("/users/delete/:uuid", authMiddleware(authService, userService), userHandler.DeleteUser)

	// merchant
	api.GET("/merchants", authMiddleware(authService, userService), merchantHandler.GetMerchants)
	api.GET("/merchant", authMiddleware(authService, userService), merchantHandler.GetMerchantByUserID)
	api.POST("/merchants", authMiddleware(authService, userService), merchantHandler.CreateMerchant)

	// outlet
	api.GET("/outlets", authMiddleware(authService, userService), outletHandler.GetOutlets)
	// api.GET("/outlet", authMiddleware(authService, userService), outletHandler.GetOutletByMerchantID)
	api.POST("/outlet", authMiddleware(authService, userService), outletHandler.CreateOutlet)

	// transaction
	api.POST("/transactions", authMiddleware(authService, userService), transactionHandler.CreateTransaction)

	// report
	api.POST("/report/merchant", authMiddleware(authService, userService), reportHandler.ReportMerchant)
	api.POST("/report/merchantByid", authMiddleware(authService, userService), reportHandler.ReportMerchantByid)
	api.POST("/report/outlet", authMiddleware(authService, userService), reportHandler.ReportOutletByid)

	api.POST("/area", authMiddleware(authService, userService), areaHandler.CreateArea)

	router.Run()

}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := uint64(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}

}
