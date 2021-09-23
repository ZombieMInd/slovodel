package server

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/ZombieMInd/slovodel/internal/store"
	"github.com/ZombieMInd/slovodel/internal/store/redisstore"
	"github.com/ZombieMInd/slovodel/internal/store/sqlstore"
	"github.com/go-redis/redis"
	"github.com/gorilla/handlers"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
)

func Start(conf *Config) error {
	store, err := initStore(conf)
	if err != nil {
		log.Fatal(err)
	}

	srv := NewServer(store)
	srv.configLogger(conf)
	srv.InitServices(conf)
	initRouter(srv)

	return http.ListenAndServe(conf.BindAddr, srv)
}

func InitConfig(conf *Config) error {
	err := envconfig.Process("API", conf)
	if err != nil {
		return fmt.Errorf("error while parsing env config: %s", err)
	}
	return nil
}

func initRouter(s *server) {
	s.router.Use(s.setRequestID)
	s.router.Use(s.logRequest)
	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
}

func newDB(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func initStore(conf *Config) (store.Store, error) {
	if conf.StoreMode == "postgres" {
		db, err := newDB(conf.DBURL)
		if err != nil {
			return nil, err
		}
		return sqlstore.New(db), nil
	} else if conf.StoreMode == "redis" {
		client := redis.NewClient(&redis.Options{
			Addr:     conf.Redis.Host,
			Password: conf.Redis.Password,
			DB:       conf.Redis.DB,
		})

		_, err := client.Ping().Result()
		if err != nil {
			return nil, err
		}

		return redisstore.New(client), nil
	}
	return nil, errors.New("unknown store mode")
}
