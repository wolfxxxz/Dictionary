package store

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// Instance of store
type Store struct {
	config         *Config
	db             *sql.DB
	userRepository *UserRepository
	wordRepository *WordRepository
}

// Constructor for store
func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

// Open store connection method
func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		return err
	}
	//Проверим, что все ок. Реально соединение тут не создается. Соединение только при первом вызове
	//db.Ping() // Пустой SELECT *
	if err := db.Ping(); err != nil {
		return err
	}
	s.db = db
	log.Println("Connection to db successfully")
	return nil
}

// Close store method connection
func (s *Store) Close() {
	s.db.Close()
}

// Public for UserRepo
func (s *Store) User() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		store: s,
	}
	return s.userRepository
}

// Public for ArticleRepo
// Public for ArticleRepo
func (s *Store) Word() *WordRepository {
	if s.wordRepository != nil {
		return s.wordRepository
	}
	s.wordRepository = &WordRepository{
		store: s,
	}
	return s.wordRepository
}
