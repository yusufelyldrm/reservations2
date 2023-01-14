package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/justinas/nosurf"
	"github.com/yusufelyldrm/bookings/internal/config"
	"github.com/yusufelyldrm/bookings/internal/models"
)

// var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplates set the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

// AddDefaultData adds default data to the templateData struct and returns it back to the caller
// function (RenderTemplate) as a pointer to the
// struct (td) so that it can be used in the template file (base.layout.html) as
// {{.CSRFToken}} and {{.Flash}} etc. (see base.layout.html) and {{.Data}}
func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {

	td.CSRFToken = nosurf.Token(r)

	return td

}

// RenderTemplate renders template using html/template
func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) {

	var tc map[string]*template.Template
	if app.UseCache {
		//get the template cache from the app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	//get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("could not get template from template cache")
	}

	//create a buffer to write to
	buf := new(bytes.Buffer)

	//add default data to template
	td = AddDefaultData(td, r)

	//execute the template
	err := t.Execute(buf, td)
	if err != nil {
		log.Println(err)
	}

	//render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}

	/*	ParsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.html")
		err = ParsedTemplate.Execute(w, nil)
		if err != nil {
			fmt.Println("error parsing template", err)
			return
		}*/

}

// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	//get all the files named *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return myCache, err
	}

	//range through all files ending with *.page.tmpl
	for _, page := range pages {

		//get the file name
		name := filepath.Base(page)
		//ParseFiles returns a *Template associated with the name
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		//get all the files named *.layout.tmpl from ./templates
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		//range through all files ending with *.layout.tmpl
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		//add the template to the cache
		myCache[name] = ts
	}
	//return the cache
	return myCache, nil
}

/*
//This is the simple one

var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	//check to see if we already have the template in our cache
	_, inMap := tc[t]
	if !inMap {
		log.Println("creating template and adding to cache ")
		//need to create the template
		err = createTemplateCache(t)

		if err != nil {
			log.Println(err)
		}

	} else {
		//we have the template in cache
		log.Println("using cached template")
	}

	tmpl = tc[t]

	err = tmpl.Execute(w, nil)

	if err != nil {
		log.Println(err)
	}

}
func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.html",
	}

	//parse the template
	tmpl, err := template.ParseFiles(templates...)

	if err != nil {
		return err
	}

	//add template to cache(map)
	tc[t] = tmpl

	return nil

}
*/
