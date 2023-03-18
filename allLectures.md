## Лекция 1. Работа с JSON файлами

**JSON** - формат файлов (расширение файлов), которое повсеместно используется для реализации передачи данных
между серверами на уровне API.

**JSON** == **JavaScript Object Notation** (Object - аналог map в Go только для мира JS)

**JSON** - это простейшее файловое расширение, поддерживающее элементарную структуризацию (выглядит как набор пар ключ: значение)

### Шаг 1. Создадим простой .json файл
Для этого определим файл ```users.json```.
```
{"users" : [{"name": "Vasya"}, {"name" : "Vitya"}]}
```
Обратите внимание на то, что ***ПО СТАНДАРТУ В JSON ИСПОЛЬЗУЮТСЯ ДВОИНЫЕ КАВЫЧКИ***.

### Шаг 2. Создадим чуть более сложный .json
Создадим сразу читаемым
```
{
    "users" : [
        {
            "name" : "Alex",
            "type" : "Admin",
            "age" : 32,
            "social" : {
                "vkontakte" : "https://vk.com/id=123512",
                "facebook": "https://fb.com/id=172835"
            }
        },
        {
            "name" : "Bob",
            "type" : "Regular",
            "age" : 12,
            "social" : {
                "vkontakte" : "https://vk.com/id=123561235",
                "facebook": "https://fb.com/id=19283712"
            }
        },
        {
            "name" : "Alice",
            "type" : "Regular",
            "age" : 19,
            "social" : {
                "vkontakte" : "https://vk.com/id=12123123",
                "facebook": "https://fb.com/id=172123123"
            }
        },
        {
            "name" : "George",
            "type" : "Regular",
            "age" : 42,
            "social" : {
                "vkontakte" : "https://vk.com/id=999999",
                "facebook": "https://fb.com/id=98888888"
            }
        }
    ]
}
```

***ВАЖНО*** - .json не гарантирует соблюдения упорядоченности при выдаче ключей!


### Шаг 3. Как в принципе читают из таких файлов?
* Для начала нужно создать файловый экземпляр (файловый дескриптор)
```
jsonFile, err := os.Open("users.json")
```
* Сразу же обрабатываем ошибки!
```
if err != nil {....}
```

* Не забываем файл закрывать!
```
defer jsonFile.Close()
```

* Затем нам нужно из файл-дескриптора забрать данные и куда-то их поместить!
```
json.Unmarshall(byteArr, &куда_помещаем)
```

### Шаг 4. Теперь более конкретно.

В Go существует 2 способа работы с JSON файлами:
* структуризованная сериализация/десериализация
* неструктуриозованная -//-


#### Шаг 4.1 Структуризация
***Серилазция*** - процесс конвертации объекта в последовательгость байтов. 
***Десериализация*** - процесс конвертации последовательности байтов в объект.

Идея структуризованного подхода состоит в том, что мы заранее подготавливаем набор структур, с ***ЯВНО ПРОПИСАННЫМИ ПРАВИЛАМИ*** сериализации/десериализации полей объектов.

```
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

//Struct for representation total slice
// First Level ob JSON object Parsing
type Users struct {
	Users []User `json:"users"`
}

//Internal user representation
//Second level of object JSON parsin
type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"age"`
	Social Social `json:"social"`
}

//Socail block representation
//Third level of object parsing
type Social struct {
	Vkontakte string `json:"vkontakte"`
	Facebook  string `json:"facebook"`
}

//Функция для распечатывания User
func PrintUser(u *User) {
	fmt.Printf("Name: %s\n", u.Name)
	fmt.Printf("Type: %s\n", u.Type)
	fmt.Printf("Age: %d\n", u.Age)
	fmt.Printf("Social. VK: %s and FB: %s\n", u.Social.Vkontakte, u.Social.Facebook)
}

//1. Рассмотрим процесс десериализации (то есть когда из последовательности в объект)
func main() {
	//1. Создадим файл дескриптор
	jsonFile, err := os.Open("users.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	fmt.Println("File descriptor successfully created!")

	//2. Теперь десериализуем содержимое jsonFile в экземпляр Go
	// Инициализируем экземпляр Users
	var users Users

	// Вычитываем содержимое jsonFile в ВИДЕ ПОСЛЕДОВАТЕЛЬНОСТИ БАЙТ!
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	// Теперь задача - перенести все из byteValue в users - это и есть десериализация!
	json.Unmarshal(byteValue, &users)
	for _, u := range users.Users {
		fmt.Println("================================")
		PrintUser(&u)
	}
}

```

Идея структурированной сериализации/десериализации состоит в том, чтобы общаться с JSON объектами напрямую , через стыкову полей.

Для того, чтобы настроить стыкову полей нужно:
* Определить необходимые уровни объектности JSON (в нашем случае их 3)
* Для каждого уровня объектности подготовить свою структуру, учитывающую набор полей объекта.
```
type Person struct {
    Name string `json:"name"
}
```
* И все! Больше ничего делать не нужно. Остается только считать из файла и поместить в экземпляр!


### Шаг 4.2 Неструктуризованный подход
В этом подходе читаемость кода стремится к нулю, но его можно использовать на этапе отладки, чтобы быстро
посмотреть, что вообще в принципе находится в json.

```
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	//1. Создадим файл дескриптор
	jsonFile, err := os.Open("users.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	fmt.Println("File descriptor successfully created!")

	//2. Вычитываю набор байт из файл-дескриптора
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	var result map[string]interface{}
	json.Unmarshal(byteValue, &result)

	fmt.Println(result["users"])

}
```


### Шаг 5. Сериализация
* Только один способ - структуризованный.
```
//1. Превратим профессора в последовательность байтов
	byteArr, err := json.MarshalIndent(prof1, "", "   ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(byteArr))
	err = os.WriteFile("output.json", byteArr, 0666) //-rw-rw-rw-
	if err != nil {
		log.Fatal(err)
	}
```

Сериализация - процесс перегона в байты. Поэтому на это этапе у нас на руках будет ```[]byte``` , который в последствии будет помещен в файл ```output.json```

***ВАЖНО*** : 0664/0666 - права доступа к файлу (в нашем случае это ```rw```)


## Лекция 3. Простейший API и термины

**Задача**: создать простой REST API , который будет позволять получать информацию про пиццу.

### Шаг 1. Желаемый функционал.
Хотим собрать простей веб-сервер, который будет взаимодействовать с окружающим миром через API поддерживающий REST.

### Шаг 1.1 Идея
Хочется, чтобы наш сервер помогал клиентам узнавать следующую информацию:
* Какая пицца есть в наличии?
* Информация, про какую-то конкретную пиццу.

### Шаг 1.2 Виды запросов, поддерживаемые api
Будет существовать и поддерживаться 2 запроса:
* ```http://localhost:8080/pizzas``` - возвращает json со всеми пиццами в наличии
* ```http://localhost:8080/pizza/{id}``` - возвращает информацию про пиццу с ```id``` в случае если она имеется в наличии, или сообщает клиенту, что такой пиццы нет.


### Шаг 2. Реализация
* Для начала создадим файл ```main.go```
* Сразу инициализируем ```go mod init 3.Trash_API```

### Шаг 2.2. Базовый скелет
```
package main

import (
	"log"
	"net/http"
)

var (
	port string = "8080"
)

func main() {
	log.Println("Trying to start REST API pizza!")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

```
**ВАЖНО**: для остановки приложения используем ```Ctrl+C``` (остановка текущего процесса и ОСВОБОЖДЕНИЕ РЕСУРСОВ).

### Шаг 2.3 Маршрутизатор и исполнители
***Маршрутизатор (router)*** - это экземпляр, который имеет внутренний функционал , заключающийся в следующем:
* принимает на вход адрес запроса (по сути это строка ```http://localhost:8080/pizzas```) и вызывает исполнителя, который будет ассоциирован с этим запросом.

***Исполнитель (handler)*** - это функция/метод, котоырй вызывает маршрутизатором.

Для того, чтобы удобно работать с маршрутизатором и не писать его с нуля. Для этого установим уже готовую библиотеку:  ```github.com/gorilla/mux``` . А устаналвивается запросом ```go get -u github.com/gorilla/mux```

```
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	port string = "8080"
)

func GetAllPizzas(writer http.ResponseWriter, request *http.Request) {}
func GetPizzaById(writer http.ResponseWriter, request *http.Request) {}

func main() {
	log.Println("Trying to start REST API pizza!")
	// Инициализируем маршрутизатор
	router := mux.NewRouter()
	//1. Если на вход пришел запрос /pizzas
	router.HandleFunc("/pizzas", GetAllPizzas).Methods("GET")
	//2. Если на вход пришел запрос вида /pizza/{id}
	router.HandleFunc("/pizza/{id}", GetPizzaById).Methods("GET")
	log.Println("Router configured successfully! Let's go!")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

```

На данный момент нами написан базовый скелет функционала API (сейчас отсутствует хранилище данных и внутренняя логика), но тем не менее, сервер конфигурируется и уже запускается.

### Шаг 2.4 Создаем модель данных
В качестве хранилища выберем слайс с экземплярами пиццы.
Для этого реализуем следующий функционал:
```
var db []Pizza


//Наша модель
type Pizza struct {
	ID       int     `json:"id"`
	Diameter int     `json:"diameter"`
	Price    float64 `json:"price"`
	Title    string  `json:"title"`
}
// Вспомогательная функция для модели (модельный метод)
func FindPizzaById(id int) (Pizza, bool) {
	var pizza Pizza
	var found bool
	for _, p := range db {
		if p.ID == id {
			pizza = p
			found = true
			break
		}
	}
	return pizza, found
}
```

Определили базу данных ```db```. Определили структуру ```Pizza``` с сопоставлением полей, а также определили функцию, которая будет просматривать слайс и говорить, есть ли в нем нужная нам пицца, или ее нет.

### Шаг 2.5 Реализуем исполнителей (handlers)
Поскольку мы собираемся запускать наш сервер , как поддерживающий REST API архитектуру, нужно, чтобы в теле каждого ответа фигурировала информация про то, каким образом наш сервер общается!
```
func GetAllPizzas(writer http.ResponseWriter, request *http.Request) {
	//Прописывать хедеры .
	writer.Header().Set("Content-Type", "application/json")
	log.Println("Get infos about all pizzas in database")
	writer.WriteHeader(200)            // StatusCode для запроса
	json.NewEncoder(writer).Encode(db) // Сериализация + запись в writer
}

```

Данный обработчик возвращает сериализованный json полученный на основе ```db``` (слайса Пицц)
```
func GetAllPizzas(writer http.ResponseWriter, request *http.Request) {
	//Прописывать хедеры .
	writer.Header().Set("Content-Type", "application/json")
	log.Println("Get infos about all pizzas in database")
	writer.WriteHeader(200)            // StatusCode для запроса
	json.NewEncoder(writer).Encode(db) // Сериализация + запись в writer
}

func GetPizzaById(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	//Считаем id из строки запроса и конвертируем его в int
	vars := mux.Vars(request) // {"id" : "12"}
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("client trying to use invalid id param:", err)
		msg := ErrorMessage{Message: "do not use ID not supported int casting"}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	log.Println("Trying to send to client pizza with id #:", id)
	pizza, ok := FindPizzaById(id)
	if ok {
		writer.WriteHeader(200)
		json.NewEncoder(writer).Encode(pizza)
	} else {
		msg := ErrorMessage{Message: "pizza with that id does not exists in database"}
		writer.WriteHeader(404)
		json.NewEncoder(writer).Encode(msg)
	}
}

```


## Лекция 5. Стандартный веб-сервер

### Шаг 1. Инициализация go mod
```
go mod init github.com/vlasove/go2/5.StandardWebServer
```

### Шаг 2. Где найти стандартные шаблоны?
**Полезная ссылка**: https://github.com/golang-standards/project-layout (здесь можно найти советы по стрктурированию/пакетированию/рефакторингу любого Go-приложения)

### Шаг 3. Создадим входную точку в приложение
Стандартный шаблон входной точки :
```
cmd/<app_name>/main.go
```
Мы создадим :
```
cmd/api/main.go
```

### Шаг 3. Инициализация ядра сервера
Стандартным шаблоном диктуется следующая схема
```
internal/app/<app_name>/<app_name>.go
```
У нас будет ```internal/app/api/api.go```

### Шаг 4. Важный пункт про конфигурацию
**Правило**: в Go принято, что:
* конфигурация всегда хранится в сторонних файлах (.toml, .env) 
* в Go проектах ВСЕГДА присутствуют дефолтные настройки (исключение - бд стараются не дефолтить)

### Шаг 5. Конфигурирование API сервера
Базово, для конфигурация нужен лишь порт.
```
intrenal/app/api/config.go
```

### Шаг 6. Создадим конфиги 
```
configs/<app_name>.toml или configs/.env
```

```
//api.toml
bind_addr = ":8080"
```

### Шаг 7. Как конфиги передавать?
Хочется запускать :
```
api.exe -path configs/api.toml
```

### Шаг 8. Конфигурация http сервера
```
go get -u github.com/gorilla/mux
```

### В дз на потом
Добавить в ```makefile```
* go build -v ./cmd/api/


Хочется, чтобы была возможность запускать наше приложение как с ```.toml``` файлом, так и с ```.env```
Добавить в код необходимые блоки, для того, чтобы можно было запускать приложение следующими командами:
* Должна быть возможность запускать проект с конфигами в ```.toml```
```
api -format .env -path configs/.env
```
* Должна быть возможность запускать проект с конфигами в ```.env```
```
api -format .toml -path configs/api.toml
```
* Должна быть возможность запускать проект с дефолтными параметрами (дефолтным будем считать ```api.toml```, если его нет, то запускаем с значениями из структуры ```Config```)
```
api
```

## Лекция 6. Подключение к БД и стандартные схемы миграции


### Шаг 0. Общие соображения
* Определить модель данных (определить правило стыковки объекта в таблице вашей СУБД с объектом внутри языка) - как объект представлен в БД
* Обработчики модели (модельные методы) - как объект взаимодействует с БД
* Выделение публичных обработчиков и стыковка их с серверными запросами

***Миграция*** - это процесс изменения схемы хранения данных. (положительный - upgrade / up миграция) (отрицательная downgrade /down миграция)

***Data repository*** (репозиторий с обработчиками) - это и есть то место, где будут жить публичные обработчики (модельные методы).

### Шаг 1. Библиотеки для работы с бд
```database/sql``` 
```sqlx```
```gosql```

### Шаг 2. Инициализация хранилища
```storage/storage.go```
Цель данного моделя определить:
* инстанс хранилища
* конструктор хранилища
* публичный метод Open (установка соединения)
* публичный метод Close (закрытие соединения)


### Шаг 3. Инициализация Storage
```storage.go```
Главная проблема кроется внутри метода Open, т.к. по факту низкоуровненвый sql.Open "ленивый" (устанавливает соединение с бд только в момент осуществления первого запрос)
```config.go```
Содержит инстанс конфига и конструктор. Атрибутом конфига является лишь строка подключения вида :
```
"host=localhost port=5432 user=postgres password=postgres dbname=restapi sslmode=disable"
```

### Шаг 4. Добавление хранилища к API
Добавим новый атрибут к API
```
//Base API server instance description
type API struct {
	//UNEXPORTED FIELD!
	config *Config
	logger *logrus.Logger
	router *mux.Router
	//Добавление поля для работы с хранилищем
	storage *storage.Storage
}
```

Добавим новый конфиуратор:
```
//Пытаемся отконфигурировать наше хранилище (storage API)
func (a *API) configreStorageField() error {
	storage := storage.New(a.config.Storage)
	//Пытаемся установить соединениение, если невозможно - возвращаем ошибку!
	if err := storage.Open(); err != nil {
		return err
	}
	a.storage = storage
	return nil
}

```

### Шаг 5. Первичная миграция
Для начала установим ```scoop```
* Открываем PowerShell: ```Set-ExecutionPolicy RemoteSigned -scope CurrentUser``` и ```Invoke-Expression (New-Object System.Net.WebClient).DownloadString('https://get.scoop.sh')```

* Для линукса/мака : https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md

После установки ```scoop``` выполним: ```scoop install migrate```

### 5.1 Создание миграционного репозитория
В данном репозитории будут находится up/down пары sql миграционных запросов к бд.
```
migrate create -ext sql -dir migrations UsersCreationMigration
```

### 5.2 Создание up/down sql файлов
См. ```migrations/....up.sql``` и ```migrations/...down.sql```

### Шаг 5.3 Применить миграцию
```
migrate -path migrations -database "postgres://localhost:5432/restapi?sslmode=disable&user=postgres&password=postgres" up
```

## Лекция 7. Работа с моделью

### Шаг 0. Откат миграции
Для выполнения отката ```migrate -path migrations -database "postgres://localhost:5432/restapi?sslmode=disable&user=postgres&password=postgres" down```


### Шаг 1. Новая миграция
Заходим в файл ```migrations/.....up.sql```
```
CREATE TABLE users (
    id bigserial not null primary key,
    login varchar not null unique,
    password varchar not null
);

CREATE TABLE articles (
    id bigserial not null primary key,
    title varchar not null unique,
    author varchar not null,
    content varchar not null
);
```

Выполним команду ```migrate -path migrations -database "postgres://localhost:5432/restapi?sslmode=disable&user=postgres&password=postgres" down```

### Шаг 2. Определим модели
Для того, чтобы определить модели ```internal/app/models/``` 2 модуля:
* user.go
* article.go

```
//user.go
package models

//User model defeniton
type User struct {
	ID       int    `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

```

```
//article.go
package models

//Article model defenition
type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Content string `json:"content"`
}

```

### Шаг 3. Определение "репозиториев"
Работать с моделями будем через репозитории. Для этого инициализируем 2 файла:
* ```storage/userrepository.go```
* ```storage/articlerepository.go```

```
//articlerepository.go
package storage

//Instance of Article repository (model interface)
type ArticleRepository struct {
    storage *Storage
}

```

Аналогично для юзера.

### Шаг 4. Выделение публичного доступа к репозиторию
Хотим, чтобы наше приложение общалось с моделями через репозитории (которые будут содержать необходимый набор метод для взаимодействия с бд). Нам необходимо определить 2 метода у хранилища , которые будут предоставлять публичные репозитории:
```
//storage.go

//Instance of storage
type Storage struct {
	config *Config
	// DataBase FileDescriptor
	db *sql.DB
	//Subfield for repo interfacing (model user)
	userRepository *UserRepository
	//Subfield for repo interfaceing (model article)
	articleRepository *ArticleRepository
}

....

//Public Repo for Article
func (s *Storage) User() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		storage: s,
	}
	return nil
}

//Public Repo for User
func (s *Storage) Article() *ArticleRepository {
	if s.articleRepository != nil {
		return s.articleRepository
	}
	s.articleRepository = &ArticleRepository{
		storage: s,
	}
	return nil
}

```

### Шаг 5. Что будет уметь делать UserRepo?
* Сохранять нового пользователя в бд (INSERT user'a или Create)
* Для аутентификации нужен функционал поиска пользователя по ```login```
* Выдача всех пользователей из бд
```
package storage

import (
	"fmt"
	"log"

	"github.com/vlasove/go2/7.ServerAndDB2/internal/app/models"
)

//Instance of User repository (model interface)
type UserRepository struct {
	storage *Storage
}

var (
	tableUser string = "users"
)

//Create User in db
func (ur *UserRepository) Create(u *models.User) (*models.User, error) {
	query := fmt.Sprintf("INSERT INTO %s (login, password) VALUES ($1, $2) RETURNING id", tableUser)
	if err := ur.storage.db.QueryRow(query, u.Login, u.Password).Scan(&u.ID); err != nil {
		return nil, err
	}
	return u, nil
}

//Find user by login
func (ur *UserRepository) FindByLogin(login string) (*models.User, bool, error) {
	users, err := ur.SelectAll()
	var founded bool
	if err != nil {
		return nil, founded, err
	}
	var userFinded *models.User
	for _, u := range users {
		if u.Login == login {
			userFinded = u
			founded = true
			break
		}
	}
	return userFinded, founded, nil
}

//Select all users in db
func (ur *UserRepository) SelectAll() ([]*models.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s", tableUser)
	rows, err := ur.storage.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	//Подготовим, куда будем читать
	users := make([]*models.User, 0)
	for rows.Next() {
		u := models.User{}
		err := rows.Scan(&u.ID, &u.Login, &u.Password)
		if err != nil {
			log.Println(err)
			continue
		}
		users = append(users, &u)
	}
	return users, nil
}

```

### Шаг 6. Что нужно от ArticleRepo?
* Уметь доавлять статью в бд
* Уметь удалять по id
* Получать все статьи
* Получать статью по id
* Обновлять (дома)
```
articlerepository.go
package storage

import (
	"fmt"
	"log"

	"github.com/vlasove/go2/7.ServerAndDB2/internal/app/models"
)

//Instance of Article repository (model interface)
type ArticleRepository struct {
	storage *Storage
}

var (
	tableArticle string = "articles"
)

//Добавить статью в бд
func (ar *ArticleRepository) Create(a *models.Article) (*models.Article, error) {
	query := fmt.Sprintf("INSERT INTO %s (title, author, content) VALUES ($1, $2, $3) RETURNING id", tableArticle)
	if err := ar.storage.db.QueryRow(query, a.Title, a.Author, a.Content).Scan(&a.ID); err != nil {
		return nil, err
	}

	return a, nil

}

//Удалять статью по id
func (ar *ArticleRepository) DeleteById(id int) (*models.Article, error) {
	article, ok, err := ar.FindArticleById(id)
	if err != nil {
		return nil, err
	}
	if ok {
		query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", tableArticle)
		_, err := ar.storage.db.Exec(query, id)
		if err != nil {
			return nil, err
		}
	}
	return article, nil
}

//Получаем статью по id
func (ar *ArticleRepository) FindArticleById(id int) (*models.Article, bool, error) {
	articles, err := ar.SelectAll()
	var founded bool
	if err != nil {
		return nil, founded, err
	}
	var articleFinded *models.Article
	for _, a := range articles {
		if a.ID == id {
			articleFinded = a
			founded = true
			break
		}
	}
	return articleFinded, founded, nil
}

//Получим все статьи в бд
func (ar *ArticleRepository) SelectAll() ([]*models.Article, error) {
	query := fmt.Sprintf("SELECT * FROM %s", tableArticle)
	rows, err := ar.storage.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	//Подготовим, куда будем читать
	articles := make([]*models.Article, 0)
	for rows.Next() {
		a := models.Article{}
		err := rows.Scan(&a.ID, &a.Title, &a.Author, &a.Content)
		if err != nil {
			log.Println(err)
			continue
		}
		articles = append(articles, &a)
	}
	return articles, nil
}

```

### Шаг 7. Описание маршрутизатора для данного проекта.
Зайдем в ```api```
```
//Пытаемся отконфигурировать маршрутизатор (а конкретнее поле router API)
func (a *API) configreRouterField() {
	a.router.HandleFunc(prefix+"/articles", a.GetAllArticles).Methods("GET")
	a.router.HandleFunc(prefix+"/articles/{id}", a.GetArticleById).Methods("GET")
	a.router.HandleFunc(prefix+"/articles/{id}", a.DeleteArticleById).Methods("DELETE")
	a.router.HandleFunc(prefix+"/articles", a.PostArticle).Methods("POST")
	a.router.HandleFunc(prefix+"/user/register", a.PostUserRegister).Methods("POST")

}
```

Создадим файл ```internal/app/api/handlers.go```
```
```



## Лекция 8. Реализация обработчиков

Из-за того, что пока у ```users``` всего один обработчик, будет держать все handlers в одном месте :
```
internal/app/api/handlers.go
```

Внутри определим 2 сущности:
```
package api

import "net/http"

//Вспомогательная структура для формирования сообщений
type Message struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	IsError    bool   `json:"is_error"`
}

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}

```

### Шаг 1. Реализация обработчика GetAllArticles
```
//Возвращает все статьи из бд на данный момент
func (api *API) GetAllArticles(writer http.ResponseWriter, req *http.Request) {
	//Инициализируем хедеры
	initHeaders(writer)
	//Логируем момент начало обработки запроса
	api.logger.Info("Get All Artiles GET /api/v1/articles")
	//Пытаемся что-то получить от бд
	articles, err := api.storage.Article().SelectAll()
	if err != nil {
		//Что делаем, если была ошибка на этапе подключения?
		api.logger.Info("Error while Articles.SelectAll : ", err)
		msg := Message{
			StatusCode: 501,
			Message:    "We have some troubles to accessing database. Try again later",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(articles)
}
```

### Шаг 2. Реализация PostArticle
```
func (api *APIServer) PostArticle(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Post Article POST /articles")
	var article models.Article
	err := json.NewDecoder(req.Body).Decode(&article)
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

	a, err := api.store.Article().Create(&article)
	if err != nil {
		api.logger.Info("Troubles while creating new article:", err)
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
```
###  Update Article by ID
```
func (api *APIServer) PutArticleById(writer http.ResponseWriter, req *http.Request) {
	//now its get
	initHeaders(writer)
	api.logger.Info("PUT Article by ID /api/v1/articles/{id}")
	// Если ID не инт
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
	// Проверка на наличие в дб
	_, ok, err := api.store.Article().FindArticleById(id)
	if err != nil {
		api.logger.Info("Troubles while accessing database table (articles) with id. err:", err)
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
		api.logger.Info("Can not find article with that ID in database")
		msg := Message{
			StatusCode: 404,
			Message:    "Article with that ID does not exists in database.",
			IsError:    true,
		}

		writer.WriteHeader(404)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	//Достать из постман json
	var article models.Article
	error := json.NewDecoder(req.Body).Decode(&article)
	if error != nil {
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
	article.ID = id
	//fmt.Println(article.ID)
	_, errar := api.store.Article().Update(&article)
	if errar != nil {
		api.logger.Info("Troubles while updating new article:", errar)
		msg := Message{
			StatusCode: 501,
			Message:    "We have some troubles to accessing database. Try again",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(article)

}
```


## Лекция 10. Простейши механизм аутентификации

На данный момент у нас реализован API , с одной проблмой - кто угодно может получить доступ к элемента в БД через публичные запросы, и например, удалить все что там имеется.

***Идея*** : сделать так, чтобы пользователь, который собирается использовать наш API не был анонимным, а мог зарегестрироваться и пройти базовую аутентификацию.

### Шаг 0. Термины
***Аутентификация*** - процесс узнавания свой/чужой. (Подразуемвает под собой сопоставление данных стороннего пользователя с данными, которые уже имеются в бд.) 
***Авторизация*** - процесс выдачи прав доступа различного уровня.


### Шаг 1. Простейшая логика при аутентификации
* К нам пришел какой-то пользователь
* Пользователь должен пройти регистрацию 
* Пользователь переходит на ресурс аутентификации и получает какой-либо аутентификационный ключ
* Далее пользователь с этим ключом может ходить по всем ресурсам нашего api.

### Шаг 2. Аутентификация с помощью JWT токена
***JWT** - ```JsonWebToken``` - символьная строка с закодированным ключом.

### Шаг 3. Немного про то, где будут выполняться действия по работе с JWT
***Middleware*** - часть ПО (архитектурная часть), которая напрямую не взаимодействует ни с клиентом, ни с сервером, а осуществляет какие-либо команды или запросы во-время клиент-серврного общения.
Например:
* Пользователь вызывает ```POST /api/v1/article +.json``` 
* Auth Middleware - проверяет, может ли данный клиент данный запрос вообще выполнять или у него не ххватает прав? (Мы не знаем кто это)
* Сервер должен принять данные и обработать запрос (добавить в бд инфу про статью)

### Шаг 4. Реализация
Добавим 2 зависмости в проект:
* ```go get -u github.com/auth0/go-jwt-middleware```
* ```go get -u github.com/form3tech-oss/jwt-go```

В следующей директории :```internal/app/middleware/middleware.go```

```
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

```

### Шаг 5. Как пользователю получить этот токен?
Нам необходимо реализовать ресурс ```/auth``` или ```api/v1/user/auth```.
```
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

```

### шАГ 6. Реализация PostToAuth
```
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
```

### Шаг 7. Проверим, что токен выбивается
Для этого идем в postman

### Шаг 8. Завернем необходимые хендлеры в JWT-REQUIRED-декоратор
Для того, чтобы обозначит факт необходимости использования JWT токена перед выполнением какого-либо запроса - заверните его в декоратор ```middleware.JwtMiddleware```
```
//Теперь требует наличия JWT
	s.router.Handle(prefix+"/articles"+"/{id}", middleware.JwtMiddleware.Handler(
		http.HandlerFunc(s.GetArticleById),
	)).Methods("GET")
	//
```

### Шаг 9. В postman
На вкладке ```Headers``` у данного запроса доавбляем пару параметров
```Authorization``` и ```Bearer <your_token_form_auth>```

## Лекция 11. Про тесты

***Проблема*** - тестируем все вручную. Это не хорошо. 

### Шаг 0. Термины
***TFD*** - это концпеция, подразумевающая написание модульных тестов еще ДО начала
реализации кода проекта (Test First Development)


### Шаг 1. Простейший пример с факториалом.
* Что оно должно уметь? Уметь вычислять факториал
* Как? ```func factorial(num int) int {}```
* Как проверить, что оно работает правильно?
```
0! = 1
1! = 1
2! = 2
3! = 6
5! = 120
6! = 720
....
Входной параметр меньше 10.
```
* В результате предыдущего пункта имеем ряд ограничений (граничные условия)

* Всегда с самого начала продумывайте граничные условия!!!

* Создаем файл с тестами
```main_test.go```
```
package main

import "testing"

type TestCase struct {
	InputData int // то, что будет подаваться на вход
	Answer    int // то, что вернет тестируемая функция
	Expected  int //то, что ожидаем получить
}

//Тестовый сценарий
var Cases []TestCase = []TestCase{
	{
		InputData: 0,
		Expected:  1,
	},
	{
		InputData: 1,
		Expected:  1,
	},
	{
		InputData: 3,
		Expected:  6,
	},
	{
		InputData: 5,
		Expected:  120,
	},
}

func TestFactorial(t *testing.T) {
	for id, test := range Cases {
		if test.Answer = factorial(test.InputData); test.Answer != test.Expected {
			t.Errorf("test case %d failed: result %v expected %v", id, test.Answer, test.Expected)
		}
	}
}

```

### Шаг 2. Реализация factoial()
```main.go```
```
func factorial(num int) int {
	if num <= 1 {
		return 1
	}
	ans := 1
	for i := 1; i <= num; i++ {
		ans *= i
	}
	return ans
}

```

### Шаг 3. Запуск.
```go test -v``` (не забывайте про go mod)

### Шаг 4. Теперь попробуем создать простейший http тест
* Что оно должно уметь? Должно уметь вычислять факториал через типичный http запрос
* Как? ```POST /factorial?num=7``` - > 7!
* Как проверить, что оно работает правильно?
```
POST /factorial?num=0 => 1
POST /factorial?num=1 => 1
POST /factorial?num=2 => 2
POST /factorial?num=3 => 6
POST /factorial?num=5 => 120
...
Ограничимся тем, что факториал 10! - самое большое значение, которое в принципе будем вычислять

```

### Шаг 5. Простейший http тест
* Откроем postman и создать тестовый запрос: ```http://localhost:8080/factorial?num=6```
* Теперь добавим тесты на уровне приложения
```
func TestHandleFactorial(t *testing.T) {
	handler := http.HandlerFunc(HandlerFactorial)
	for _, test := range HttpCases {
		//Подтест (суб-тест)
		t.Run(test.Name, func(t *testing.T) {
			recorder := httptest.NewRecorder() // Куда писать ответ
			handlerData := fmt.Sprintf("/factorial?num=%d", test.Numeric)
			request, err := http.NewRequest("GET", handlerData, nil) //Какой будет запрос
			// data := io.Reader([]byte(`{"num" : 5}`))
			// request, err := http.Post("http://localhost:8080/factorial?num=5", "application/json", data)
			if err != nil {
				t.Error(err)
			}
			handler.ServeHTTP(recorder, request) // Выполняем запрос и ответ записываем в recorder
			if string(recorder.Body.Bytes()) != string(test.Expected) {
				t.Errorf("test %s failed: input: %v! result: %v expected %v",
					test.Name,
					test.Numeric,
					string(recorder.Body.Bytes()),
					string(test.Expected),
				)
			}

		}) // Под-тестовый раннер
	}

}
```

* Теперь реализация:
```
package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

var (
	dataMap map[int]int
)

func init() {
	dataMap = make(map[int]int)
}

func main() {
	http.HandleFunc("/factorial", HandlerFactorial)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func factorial(num int) int {
	if num <= 1 {
		return 1
	}
	if result, ok := dataMap[num]; ok {
		return result
	}
	ans := 1
	for i := 1; i <= num; i++ {
		ans *= i
	}
	dataMap[num] = ans
	return ans
}

func HandlerFactorial(writer http.ResponseWriter, request *http.Request) {
	//http://localhost:8080/factorial?num=10
	num := request.FormValue("num")
	n, err := strconv.Atoi(num)
	if err != nil {
		http.Error(writer, err.Error(), 404) // msg := .....
		return
	}
	io.WriteString(writer, strconv.Itoa(factorial(n)))
}

```

### Шаг 6. Покрытие
Замер покрытия делаем через ```go test -cover``` (в будущем посмотрите на ```gotool coverage```).
По-хорошему нужно покрыть 70-85% вашего кода. 

## Заключительная лекция. Пример использования фреймворков

Попробуем реализовать простейший REST api с использованием популярной связки ```Gin + Gorm``` (```Beego```, ```echo```....).

### Шаг 0. Отвращение к фреймворкам
Никогда их не используйте. Это фу. - это миф. Перед выюором фреймворка для начала подумайте, а для чего он вам нужен? 

### Шаг 1. Зависимости:
```
go get -u github.com/jinzhu/gorm
go get -u github.com/gin-gonic/gin
go get -u github.com/lib/pq 
```

### Шаг 2. Все
Готово:)

### Шаг 3. Хитрый вопрос
У вас 2 горутины:
* main
* и какая-то запущенная из main
```
func another(){
    panic()
}

func main(){
    go another()
    ....
}
```
**another** в ходе своей работы провоцирует panic()
Вопрос - возможно ли ее обработать из main? Намек - посмотрие в сторону каналов!