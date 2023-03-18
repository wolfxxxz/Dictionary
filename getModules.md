### Стандарты составления пакетов и схема
``` https://github.com/golang-standards/project-layout``` 
Dictionary
    cmd
       apiserver
                main.go
    configs
          .env
          api.toml
    internal
        app
           apiserver
           middleware
           models
    migrations
           up.sql
           down.sql
    store

### Create and 
Создать C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Dictionary

### Makefile and command
* go mod init 
   Доступ будет не только локально
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Dictionary> 
 ```go mod init github.com/Wolfxxxz/Dictionary```

* Go build 
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Dictionary>go build -v ./cmd/api/
 ```go build -v ./cmd/api/```

* run 
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Dictionary>./api
 ```./api```

* ```Makefile``` 
Установить ```gnuwin32.sourceforge.net/packages/make.htm```
--Прописать в настройках Windows в puth путь до make.exe
build:
	go build -v ./cmd/api/
run:
	./api

### Библиотеки
-u обновить версию пакета
* 1 Библиотека для работы с маршрутизатором:  ```github.com/gorilla/mux``` . 
```go get -u github.com/gorilla/mux```

* 2 Библиотека для работы с .env файлами:  ```github.com/joho/godotenv```
```go get -u github.com/joho/godotenv```
   
* 3 Пакет для работы с api.toml файлами ```github.com/BurntSushi/toml```
```go get -u github.com/BurntSushi/toml```

* 4 Стандартная библиотека flag
flag.StringVar()
flag.Parse()
 ```api -path configs/api.toml``` or ``api.exe -path configs/api.toml``
 ```api -help```

* 5 Библиотека для отправки логов на разные серверы ```github.com/sirupsen/logrus```
```go get -u github.com/sirupsen/logrus```
--Добавить api logger logrus
api.toml - Можно менять порт и уровень логирования ошибок

* 6 Библиотека pgsql ```github.com/lib/pq```
 ```go get -u github.com/lib/pq```
--Бывает тупит
    _"github.com/lib/pq"

* 7 Библиотеки для создания и проверки токенов ```github.com/auth0/go-jwt-middleware``` ```github.com/form3tech-oss/jwt-go```
 ```go get -u github.com/auth0/go-jwt-middleware```
 ```go get -u github.com/form3tech-oss/jwt-go```

### PgSQL
* 1 Библиотека pgsql ```github.com/lib/pq```
 ```go get -u github.com/lib/pq```
--Бывает тупит
    _"github.com/lib/pq"

* 2 PgAdmin
порт sql 5433 
 ```host=localhost dbname=restapi port=5433 user=postgres password=1 sslmode=disable```

### Migration sql up down
 * 1 Instal Scoop
```scoop.sh```
 PowerShell, run 
1--```Set-ExecutionPolicy RemoteSigned -Scope CurrentUser```
2--(does nt work)```Invoke-Expression (New-Object System.Net.WebClient).DownloadString(`https://get.scoop.sh`)```
scoop.sh```
3--PS C:\Users\Mvmir> irm get.scoop.sh | iex                                                                              
                     >> # You can use proxies if you have network trouble in accessing GitHub, e.g.                                         
                     >> irm get.scoop.sh -Proxy 'http://<ip:port>' | iex
* 2 cmd ```scoop install migrate```
  C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Dictionary>scoop install migrate
 Для линукса/мака : https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md

* 3 Создание миграционного репозитория
 ```migrate create -ext sql -dir migrations UsersCreationMigration```
migrate - миграция
        create - создать
               -ext sql - расширение
                        -dir migrations - где хранить миграционные скрипты (директория)
                                        UsersCreationMigration - коментарий
cmd
C:\Users\Mvmir\go\src\github.com\Mvmir\go2\ServerAndDB>migrate create -ext sql -dir migrations UsersCreationMigration
Команда создаёт папку с двумя файлами ...up.sql (update) & ...down.sql (откат)

* 4 Заполняем up/down sql файлов
 См. ```migrations/....up.sql``` и ```migrations/...down.sql```

* 5 Применить миграцию
* migrate -path migrations -database "postgres://localhost:5433/words?sslmode=disable&user=postgres&password=1" up
* migrate -path migrations -database "postgres://localhost:5433/words?sslmode=disable&user=postgres&password=1" down
cmd run
migrate -path migrations -database "postgres://localhost:5433/words?sslmode=disable&user=postgres&password=1" up
migrate - Миграция
        -path migrations 
                          -database "postgres://localhost:5433/restapi?sslmode=disable&
                                   user=postgres&
                                        password=1" 
                                           up  - (Update)
down - запускает второй файл с командой DROP TABLE users

### Middleware Аутентификация с помощью JWT токена
***JWT** - ```JsonWebToken``` - символьная строка с закодированным ключом.

* Немного про то, где будут выполняться действия по работе с JWT
```Middleware``` - часть ПО (архитектурная часть), которая напрямую не взаимодействует ни с клиентом, ни с сервером, а осуществляет какие-либо команды или запросы во-время клиент-серврного общения.
Например:
 Пользователь вызывает ```POST /api/v1/article +.json``` 
 Auth Middleware - проверяет, может ли данный клиент данный запрос вообще выполнять или у него не ххватает прав? (Мы не знаем кто это)
 Сервер должен принять данные и обработать запрос (добавить в бд инфу про статью)

* Реализация
Добавим 2 зависмости в проект:
 ```go get -u github.com/auth0/go-jwt-middleware```
 ```go get -u github.com/form3tech-oss/jwt-go```

* В следующей директории :```internal/app/middleware/middleware.go```
package middleware
import (
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
)
var (
	SecretKey      []byte      = []byte("UltraRestApiSectryKey9000")
	emptyValidFunc jwt.Keyfunc = func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	}
)
var JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: emptyValidFunc,
	SigningMethod:       jwt.SigningMethodHS256,
})
* Как пользователю получить этот токен?
* Нам необходимо реализовать ресурс ```/auth``` или ```api/v1/user/auth```.
//func for configure Router
func (s *APIServer) configureRouter() {
	s.router.HandleFunc(prefix+"/articles", s.GetAllArticles).Methods("GET")
	s.router.HandleFunc(prefix+"/articles"+"/{id}", s.GetArticleById).Methods("GET")
	s.router.HandleFunc(prefix+"/articles"+"/{id}", s.DeleteArticleById).Methods("DELETE")
	s.router.HandleFunc(prefix+"/articles", s.PostArticle).Methods("POST")
	s.router.HandleFunc(prefix+"/user/register", s.PostUserRegister).Methods("POST")
	//new pair for auth
	s.router.HandleFunc(prefix+"/user/auth", s.PostToAuth).Methods("POST")
}
* Реализация PostToAuth
func (api *APIServer) PostToAuth(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Post to Auth POST /api/v1/user/auth")
	var userFromJson models.User
	err := json.NewDecoder(req.Body).Decode(&userFromJson)
	//Обрабатываем случай, если json - вовсе не json или в нем какие-либо пробелмы
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
	//Необходимо попытаться обнаружить пользователя с таким login в бд
	userInDB, ok, err := api.store.User().FindByLogin(userFromJson.Login)
	// Проблема доступа к бд
	if err != nil {
		api.logger.Info("Can not make user search in database:", err)
		msg := Message{
			StatusCode: 500,
			Message:    "We have some troubles while accessing database",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	//Если подключение удалось , но пользователя с таким логином нет
	if !ok {
		api.logger.Info("User with that login does not exists")
		msg := Message{
			StatusCode: 400,
			Message:    "User with that login does not exists in database. Try register first",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	//Если пользователь с таким логином ест ьв бд - проверим, что у него пароль совпадает с фактическим
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
	//Теперь выбиваем токен как знак успешной аутентифкации
	token := jwt.New(jwt.SigningMethodHS256)             // Тот же метод подписания токена, что и в JwtMiddleware.go
	claims := token.Claims.(jwt.MapClaims)               // Дополнительные действия (в формате мапы) для шифрования
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix() //Время жизни токена
	claims["admin"] = true
	claims["name"] = userInDB.Login
	tokenString, err := token.SignedString(middleware.SecretKey)
	//В случае, если токен выбить не удалось!
	if err != nil {
		api.logger.Info("Can not claim jwt-token")
		msg := Message{
			StatusCode: 500,
			Message:    "We have some troubles. Try again",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	//В случае, если токен успешно выбит - отдаем его клиенту
	msg := Message{
		StatusCode: 201,
		Message:    tokenString,
		IsError:    false,
	}
	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(msg)
}

* Проверим, что токен выбивается postman
Для этого идем в postman

* Завернем необходимые хендлеры в JWT-REQUIRED-декоратор
Для того, чтобы обозначит факт необходимости использования JWT токена перед выполнением какого-либо запроса - заверните его в декоратор ```middleware.JwtMiddleware```
//Теперь требует наличия JWT
	s.router.Handle(prefix+"/articles"+"/{id}", middleware.JwtMiddleware.Handler(
		http.HandlerFunc(s.GetArticleById),
	)).Methods("GET")
	//

* В postman
На вкладке ```Headers``` у данного запроса доавбляем пару параметров
```Authorization``` и ```Bearer <your_token_form_auth>```