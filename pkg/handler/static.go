package handler

import (
	"diveshareBackendGin/pkg/service"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"os"
	"strconv"
)

func getWd() string {
	wd, _ := os.Getwd()
	return wd
}

type Auto struct {
	Image template.URL
	Model string
	New   string
	Price string
	Id    template.URL
}

func (h *Handler) staticMain(c *gin.Context) {
	wd := getWd()
	tmpl, err := template.ParseFiles(wd + "/internal/static/index.html")
	if err != nil {
		c.Writer.WriteHeader(http.StatusNotFound)
		return
	}
	style, err := os.ReadFile(wd + "/internal/static/main.css")
	script, err := os.ReadFile(wd + "/internal/static/scripts/main.js")
	if err != nil {
		// handle error
		return
	}
	tmplData := struct {
		Style  template.CSS
		Script template.JS
	}{Style: template.CSS(style), Script: template.JS(script)}
	err = tmpl.Execute(c.Writer, tmplData)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *Handler) staticRegistration(c *gin.Context) {

}

func initCatalog(service *service.Service) []Auto {
	Autos := make([]Auto, 0)
	cars, err := service.Car.GetAll()
	if err != nil {
		return nil
	}

	for _, car := range cars {
		tmp := Auto{}
		tmp.Model = car.Brand + " " + car.Model
		tmp.Price = car.Price
		tmp.Image = template.URL(car.Image)
		tmp.New = car.Category
		tmp.Id = template.URL(strconv.Itoa(car.Id))
		Autos = append(Autos, tmp)
	}

	return Autos
}

func (h *Handler) staticCatalog(c *gin.Context) {
	Cars := initCatalog(h.services)

	wd := getWd()
	tmpl, err := template.ParseFiles(wd + "/internal/static/catalog.html")
	if err != nil {
		c.Writer.WriteHeader(http.StatusNotFound)
		return
	}

	style, err := os.ReadFile(wd + "/internal/static/catalog.css")
	if err != nil {
		// handle error
		return
	}

	tmplData := struct {
		Style template.CSS
		Autos []Auto
	}{Style: template.CSS(style), Autos: Cars}
	err = tmpl.Execute(c.Writer, tmplData)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *Handler) staticAbout(c *gin.Context) {
	wd := getWd()
	tmpl, err := template.ParseFiles(wd + "/internal/static/about.html")
	if err != nil {
		c.Writer.WriteHeader(http.StatusNotFound)
		return
	}
	style, err := os.ReadFile(wd + "/internal/static/about.css")
	if err != nil {
		// handle error
		return
	}
	tmplData := struct {
		Style template.CSS
	}{Style: template.CSS(style)}
	err = tmpl.Execute(c.Writer, tmplData)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *Handler) staticBlog(c *gin.Context) {
	wd := getWd()
	tmpl, err := template.ParseFiles(wd + "/internal/static/blog.html")
	if err != nil {
		c.Writer.WriteHeader(http.StatusNotFound)
		return
	}
	style, err := os.ReadFile(wd + "/internal/static/blog.css")
	if err != nil {
		// handle error
		return
	}
	tmplData := struct {
		Style template.CSS
	}{Style: template.CSS(style)}
	err = tmpl.Execute(c.Writer, tmplData)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *Handler) staticCar(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}

	wd := getWd()
	tmpl, err := template.ParseFiles(wd + "/internal/static/car.html")
	if err != nil {
		c.Writer.WriteHeader(http.StatusNotFound)
		return
	}

	car, err := h.services.Car.GetById(id)
	if err != nil {
		return
	}

	carById := Auto{
		Image: template.URL(car.Image),
		Model: car.Brand + " " + car.Model,
		New:   car.Category,
		Price: car.Price,
		Id:    template.URL(strconv.Itoa(car.Id)),
	}

	style, err := os.ReadFile(wd + "/internal/static/car.css")
	if err != nil {
		// handle error
		return
	}

	tmplData := struct {
		Style template.CSS
		Car   Auto
	}{Style: template.CSS(style), Car: carById}

	err = tmpl.Execute(c.Writer, tmplData)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *Handler) staticCareer(c *gin.Context) {
	wd := getWd()
	tmpl, err := template.ParseFiles(wd + "/internal/static/career.html")
	if err != nil {
		c.Writer.WriteHeader(http.StatusNotFound)
		return
	}
	style, err := os.ReadFile(wd + "/internal/static/career.css")
	if err != nil {
		// handle error
		return
	}
	tmplData := struct {
		Style template.CSS
	}{Style: template.CSS(style)}
	err = tmpl.Execute(c.Writer, tmplData)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *Handler) staticContact(c *gin.Context) {
	wd := getWd()
	tmpl, err := template.ParseFiles(wd + "/internal/static/contact.html")
	if err != nil {
		c.Writer.WriteHeader(http.StatusNotFound)
		return
	}
	style, err := os.ReadFile(wd + "/internal/static/contact.css")
	if err != nil {
		// handle error
		return
	}
	tmplData := struct {
		Style template.CSS
	}{Style: template.CSS(style)}
	err = tmpl.Execute(c.Writer, tmplData)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *Handler) staticDescription(c *gin.Context) {
	wd := getWd()
	tmpl, err := template.ParseFiles(wd + "/internal/static/description.html")
	if err != nil {
		c.Writer.WriteHeader(http.StatusNotFound)
		return
	}
	style, err := os.ReadFile(wd + "/internal/static/description.css")
	if err != nil {
		// handle error
		return
	}
	tmplData := struct {
		Style template.CSS
	}{Style: template.CSS(style)}
	err = tmpl.Execute(c.Writer, tmplData)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *Handler) staticFaq(c *gin.Context) {
	wd := getWd()
	tmpl, err := template.ParseFiles(wd + "/internal/static/faq.html")
	if err != nil {
		c.Writer.WriteHeader(http.StatusNotFound)
		return
	}
	style, err := os.ReadFile(wd + "/internal/static/faq.css")
	if err != nil {
		// handle error
		return
	}
	tmplData := struct {
		Style template.CSS
	}{Style: template.CSS(style)}
	err = tmpl.Execute(c.Writer, tmplData)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *Handler) staticSafety(c *gin.Context) {
	wd := getWd()
	tmpl, err := template.ParseFiles(wd + "/internal/static/safety.html")
	if err != nil {
		c.Writer.WriteHeader(http.StatusNotFound)
		return
	}
	style, err := os.ReadFile(wd + "/internal/static/safety.css")
	if err != nil {
		// handle error
		return
	}
	tmplData := struct {
		Style template.CSS
	}{Style: template.CSS(style)}
	err = tmpl.Execute(c.Writer, tmplData)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *Handler) staticShare(c *gin.Context) {
	wd := getWd()
	tmpl, err := template.ParseFiles(wd + "/internal/static/share.html")
	if err != nil {
		c.Writer.WriteHeader(http.StatusNotFound)
		return
	}
	style, err := os.ReadFile(wd + "/internal/static/share.css")
	if err != nil {
		// handle error
		return
	}
	tmplData := struct {
		Style template.CSS
	}{Style: template.CSS(style)}
	err = tmpl.Execute(c.Writer, tmplData)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *Handler) staticUpgrade(c *gin.Context) {
	wd := getWd()
	tmpl, err := template.ParseFiles(wd + "/internal/static/upgrade.html")
	if err != nil {
		c.Writer.WriteHeader(http.StatusNotFound)
		return
	}
	style, err := os.ReadFile(wd + "/internal/static/upgrade.css")
	if err != nil {
		// handle error
		return
	}
	tmplData := struct {
		Style template.CSS
	}{Style: template.CSS(style)}
	err = tmpl.Execute(c.Writer, tmplData)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *Handler) staticNew(c *gin.Context) {
	wd := getWd()
	tmpl, err := template.ParseFiles(wd + "/internal/static/new.html")
	if err != nil {
		c.Writer.WriteHeader(http.StatusNotFound)
		return
	}
	err = tmpl.Execute(c.Writer, nil)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
	}
}
