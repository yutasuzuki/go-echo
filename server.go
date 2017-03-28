package main

import (
	"net/http"
	"github.com/labstack/echo"
	"io"
	"text/template"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
	"fmt"
	"github.com/labstack/echo/middleware"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	t := &Template{
		templates: template.Must(template.ParseGlob("template/*.html")),
	}
	e.Renderer = t
	e.GET("/", RootHandler)
	e.GET("/signup", SignUpHandler)
	e.POST("/signup", PostSignUpHandler)
	e.GET("/signin", SignInHandler)
	e.POST("/signin", PostSignInHandler)
	e.Logger.Fatal(e.Start(":1323"))
}

func RootHandler(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "username"
	cookie.Value = "jon"
	cookie.Expires = time.Now().Add(1 * time.Minute)
	c.SetCookie(cookie)
	cookie.Name = "aaa"
	cookie.Value = "hoge"
	cookie.Expires = time.Now().Add(1 * time.Minute)
	c.SetCookie(cookie)
	return c.Render(http.StatusOK, "index", "!!")
}

func SignInHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "signin", "!! Sign in!!")
}

func SignUpHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "signup", "!! Sign in!!")
}

type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type Store struct {
	gorm.Model
	Name string `json:"name" form:"name" query:"name"`
	Mail string `json:"mail" form:"mail" query:"mail"`
	Password string `json:"password" form:"password" query:"password"`
}

func PostSignUpHandler(c echo.Context) error {
	post := new(Store)
	if err := c.Bind(post); err != nil {
		return err
	}
	db, err := gorm.Open("mysql", "root:@/heaup_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	store := Store{Name: post.Name, Mail: post.Mail, Password: post.Password}
	db.Create(&store)
	res := db.First(&store)
	defer db.Close()
	return c.JSON(http.StatusCreated, res)
}

func PostSignInHandler(c echo.Context) error {
	post := new(Store)
	if err := c.Bind(post); err != nil {
		return err
	}
	db, err := gorm.Open("mysql", "root:@/heaup_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	store := Store{Mail: post.Mail, Password: post.Password}
	res := db.Where("mail = ? AND password = ?", post.Mail, post.Password).Find(&store)
	fmt.Println(res.Error)
	if res.Error == nil {

	}
	defer db.Close()
	return c.JSON(http.StatusCreated, res)
}