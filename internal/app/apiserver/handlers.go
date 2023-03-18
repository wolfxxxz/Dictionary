package apiserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Wolfxxxz/Dictionary/internal/app/middleware"
	"github.com/Wolfxxxz/Dictionary/internal/app/models"
	"github.com/Wolfxxxz/Dictionary/library"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gorilla/mux"
)

type Message struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	IsError    bool   `json:"is_error"`
}

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}

// Возвращает все статьи из бд на данный момент
func (api *APIServer) GetAllWords(writer http.ResponseWriter, req *http.Request) {
	//Инициализируем хедеры
	initHeaders(writer)
	//Логируем момент начало обработки запроса
	word, err := api.store.Word().SelectAll()
	if err != nil {
		api.logger.Info(err)
		msg := Message{
			StatusCode: 501,
			Message:    "We have some troubles to accessing words in database. Try later",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	api.logger.Info("Get All Words GET /words")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(word)
}

func (api *APIServer) PostWord(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Post Word POST /words")
	var word models.Word
	err := json.NewDecoder(req.Body).Decode(&word)
	if err != nil {
		api.logger.Info("Invalid json recieved from client")
		msg := Message{
			StatusCode: 400,
			Message:    "Provided json is invalid",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	a, err := api.store.Word().Create(&word)
	if err != nil {
		api.logger.Info("Troubles while creating new word:", err)
		msg := Message{
			StatusCode: 501,
			Message:    "We have some troubles to accessing database. Try again",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(a)

}

func (api *APIServer) GetWordById(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Get Article by ID /api/v1/word/{id}")
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		api.logger.Info("Troubles while parsing {id} param:", err)
		msg := Message{
			StatusCode: 400,
			Message:    "Unapropriate id value. don't use ID as uncasting to int value.",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	word, ok, err := api.store.Word().FindWordById(id)
	if err != nil {
		api.logger.Info("Troubles while accessing database table (word) with id. err:", err)
		msg := Message{
			StatusCode: 500,
			Message:    "We have some troubles to accessing database. Try again",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	if !ok {
		api.logger.Info("Can not find word with that ID in database")
		msg := Message{
			StatusCode: 404,
			Message:    "Word with that ID does not exists in database.",
			IsError:    true,
		}

		writer.WriteHeader(404)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(word)

}

func (api *APIServer) DeleteWordById(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Delete Article by Id DELETE /api/v1/word/{id}")
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		api.logger.Info("Troubles while parsing {id} param:", err)
		msg := Message{
			StatusCode: 400,
			Message:    "Unapropriate id value. don't use ID as uncasting to int value.",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	_, ok, err := api.store.Word().FindWordById(id)
	if err != nil {
		api.logger.Info("Troubles while accessing database table (word) with id. err:", err)
		msg := Message{
			StatusCode: 500,
			Message:    "We have some troubles to accessing database. Try again",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	if !ok {
		api.logger.Info("Can not find word with that ID in database")
		msg := Message{
			StatusCode: 404,
			Message:    "Word with that ID does not exists in database.",
			IsError:    true,
		}

		writer.WriteHeader(404)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	_, err = api.store.Word().DeleteById(id)
	if err != nil {
		api.logger.Info("Troubles while deleting database elemnt from table (word) with id. err:", err)
		msg := Message{
			StatusCode: 501,
			Message:    "We have some troubles to accessing database. Try again",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.WriteHeader(202)
	msg := Message{
		StatusCode: 202,
		Message:    fmt.Sprintf("Word with ID %d successfully deleted.", id),
		IsError:    false,
	}
	json.NewEncoder(writer).Encode(msg)
}

func (api *APIServer) PostUserRegister(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Post User Register POST /api/v1/user/register")
	var user models.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		api.logger.Info("Invalid json recieved from client")
		msg := Message{
			StatusCode: 400,
			Message:    "Provided json is invalid",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	//Пытаемся найти пользователя с таким логином в бд
	_, ok, err := api.store.User().FindByLogin(user.Login)
	if err != nil {
		api.logger.Info("Troubles while accessing database table (users) with id. err:", err)
		msg := Message{
			StatusCode: 500,
			Message:    "We have some troubles to accessing database. Try again",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	//Смотрим, если такой пользователь уже есть - то никакой регистрации мы не делаем!
	if ok {
		api.logger.Info("User with that ID already exists")
		msg := Message{
			StatusCode: 400,
			Message:    "User with that login already exists in database",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	//Теперь пытаемся добавить в бд
	userAdded, err := api.store.User().Create(&user)
	if err != nil {
		api.logger.Info("Troubles while accessing database table (users) with id. err:", err)
		msg := Message{
			StatusCode: 500,
			Message:    "We have some troubles to accessing database. Try again",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	msg := Message{
		StatusCode: 201,
		Message:    fmt.Sprintf("User {login:%s} successfully registered!", userAdded.Login),
		IsError:    false,
	}
	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(msg)
}

func (api *APIServer) PostAllWords(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Post Word POST /words")

	var word []*models.Word = library.Takejson("txt/library.json")

	/*err := json.NewDecoder(req.Body).Decode(&word)
	if err != nil {
		api.logger.Info("Invalid json recieved from client")
		msg := Message{
			StatusCode: 400,
			Message:    "Provided json is invalid",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}*/
	//
	for _, v := range word {
		a, err := api.store.Word().Create(v)
		if err != nil {
			api.logger.Info("Troubles while creating new word:", err)
			msg := Message{
				StatusCode: 501,
				Message:    "This words already in db",
				IsError:    true,
			}
			writer.WriteHeader(501)
			json.NewEncoder(writer).Encode(msg)
			continue
		}
		writer.WriteHeader(201)
		json.NewEncoder(writer).Encode(a)
	}

}

// jwt
func (api *APIServer) PostToAuth(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Post to Auth /api/v1/user/auth")
	var userFromJson models.User
	err := json.NewDecoder(req.Body).Decode(&userFromJson)
	// Если не json or some troubles
	if err != nil {
		api.logger.Info("Invalid json recieved from client")
		msg := Message{
			StatusCode: 400,
			Message:    "Provided json is invalid",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	//Пытаемся найти пользователя в бд
	userInDB, ok, err := api.store.User().FindByLogin(userFromJson.Login)
	// Проблема доступа к бд
	if err != nil {
		api.logger.Info("Can not make user search in database", err)
		msg := Message{
			StatusCode: 500,
			Message:    "We have some troubles while accessing database",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	//Если нет такого пользователя в бд
	if !ok {
		api.logger.Info("User with that Login does not exists")
		msg := Message{
			StatusCode: 400,
			Message:    "User with that login does'nt exists in database. Try register first",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	// Проверяем что пароль совпадает
	if userInDB.Password != userFromJson.Password {
		api.logger.Info("Invalid credetials to auth")
		msg := Message{
			StatusCode: 404,
			Message:    "Your password is invalid",
			IsError:    true,
		}
		writer.WriteHeader(404)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	// Теперь выбиваем токен как знак успешной аутентификации
	token := jwt.New(jwt.SigningMethodHS256)                // Тот же метод подписания что и в JwtMiddleware
	claims := token.Claims.(jwt.MapClaims)                  // Дополнительные действия (в формате map) для шифрования
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix() //время после которого токен будет считатся протухшим (здесь 2 часа)
	claims["admin"] = true
	claims["name"] = userInDB.Login
	// получаем токен
	tokenString, err := token.SignedString(middleware.SecretKey)
	// Если токен выбить не удалось по каким либо причинам
	if err != nil {
		api.logger.Info("Can not claim jwt-token")
		msg := Message{
			StatusCode: 500,
			// Не указывать что проблема с получением токена иначе могут быть траблы
			Message: "We have some troubles. Try again",
			IsError: true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	// Если токен успешно выбит отдаём его клиенту
	msg := Message{
		StatusCode: 201,
		Message:    tokenString,
		IsError:    false,
	}
	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(msg)

}
