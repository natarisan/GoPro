package app

import(
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"GOP/service"
	"GOP/domain"
	"os"
	"fmt"
	"github.com/natarisan/gop-libs/logger"
	"time"
	"github.com/jmoiron/sqlx"
)

func sanityCheck() {
	envProps := []string{
		"SERVER_ADDRESS",
		"SERVER_PORT",
		"DB_USER",
		"DB_PASSWD",
		"DB_ADDR",
		"DB_PORT",
		"DB_NAME",
	}
	for _, k := range envProps {
		if os.Getenv(k) == "" {
			logger.Error(fmt.Sprintf("Environment variable %s not defined. Terminating application...", k))
		}
	}
}

func Start(){
	sanityCheck()
	router := mux.NewRouter()
	dbClient := getDbClient()
	customerRepositoryDb := domain.NewCustomerRepositoryDb(dbClient)
	accountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)
	ch := CustomerHandlers{service.NewCustomerService(customerRepositoryDb)}
	ah := AccountHandler{service.NewAccountService(accountRepositoryDb)}
	
    router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet).Name("GetAllCustomers")
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet).Name("GetCustomer")
	router.HandleFunc("/get-images/{customer_id:[0-9]+}", ch.getImages).Methods(http.MethodGet).Name("GetImages")
	router.HandleFunc("/post-image", ch.postImage).Methods(http.MethodPost).Name("PostImage")
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.newAccount).Methods(http.MethodPost).Name("NewAccount")
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", ah.MakeTransaction).Methods(http.MethodPost).Name("NewTransaction")
	
	am := AuthMiddleware{domain.NewAuthRepository()}
	router.Use(am.authorizationHandler())

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
    log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s",address,port), router))
}

func getDbClient() *sqlx.DB{
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWD")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",dbUser,dbPasswd,dbAddr,dbPort,dbName)
	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}