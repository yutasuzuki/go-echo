package main

import (
	"net/http"
	"github.com/labstack/echo"
	"io"
	"text/template"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"time"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	t := &Template{
		templates: template.Must(template.ParseGlob("template/*.html")),
	}
	e.Renderer = t
	e.GET("/", RootHandler)
	e.GET("/signup", SignUpHandler)
	e.GET("/signin", SignInHandler)
	e.POST("/create-store", CreateStoreHandler)
	e.Logger.Fatal(e.Start(":1323"))
}

func RootHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "index", "!!")
}

func SignUpHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "index", "! Sign up!!")
}

func SignInHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "index", "!! Sign in!!")
}

type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type Store struct {
	gorm.Model
	Name string `json:"name"`
	Mail string `json:"mail"`
	Password string `json:"password"`
}

func CreateStoreHandler(c echo.Context) error {
	post := new(Store)
	println(post.Name)
	db, err := gorm.Open("mysql", "root:@/heaup_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	store := Store{Name: c.FormValue("name"), Mail: c.FormValue("mail"), Password: c.FormValue("password")}
	db.Create(&store)
	// var store Store
	res := db.First(&store)
	defer db.Close()
	// id := c.Param("name")
	return c.JSON(http.StatusCreated, res)
}
