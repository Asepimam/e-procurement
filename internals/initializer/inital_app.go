package initializer

import (
	"database/sql"
	"e-procurement/internals/delivery/routers"
	"e-procurement/internals/repositories"
	"e-procurement/internals/usecases"
	"e-procurement/pkg/auth"
	"e-procurement/pkg/connections"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

type App struct {
	Router http.Handler
	DB     *sql.DB
}


func InitializeApp() (*App, error) {
	dbConfig := connections.DBConfig{
		Driver:          "postgres",
		DSN: 		   "postgres://postgres:1234@localhost:5432/dropify_store?sslmode=disable" ,
		MaxOpenConns:    10,
		MaxIdleConns:    5,
		ConnMaxLifetime: 5 * time.Minute,
	}

	db, err := connections.ConnectDB(dbConfig)
	if err != nil {
		return nil, err
	}

	// initialize jwt
	jwtSecret := "secreate"
	JWT:=auth.NewJWT(jwtSecret)
	// if os.Getenv("SECRET_KE")

	// intial repositories
	categoryRepo := repositories.NewCategoryRepository(db)
	userRepo := repositories.NewUserRepository(db)
	productRepo := repositories.NewProductUseCase(db)
	// intial usecases
	authUseCase := usecases.NewAuthUseCase(userRepo)
	productUsecase := usecases.NewProductUsecase(productRepo)
	categoryUsecase := usecases.NewCategoryUsecase(categoryRepo)
	// inital routers
	r := routers.Router{
		User:   "user",
		Auth: *authUseCase,
		Vendor: "vendor",
		Product: *productUsecase,
		Category: *categoryUsecase,
		JWT: JWT,
	}
	routers := routers.NewRouter(&r)
	app := &App{
		Router: routers,
		DB:     db,
	}
	return app, nil

}