package render

import (
	"bytes"
	"log"
	"myGoWebApp/pkg/config"
	"myGoWebApp/pkg/models"
	"net/http"
	"path/filepath"
	"text/template"
)

var app *config.AppConfig

// NewTemplates sets config for template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(data *models.TemplateData) *models.TemplateData {

	return data
}

// Renders templates using html template
func RenderTemplate(w http.ResponseWriter, tmpl string, data *models.TemplateData) {

	var tc map[string]*template.Template
	if app.UseCache {
		// get template cache from config

		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// get template from cache
	t, okay := tc[tmpl]
	if !okay {
		log.Fatal("Couldn't load template cache")
	}

	buffer := new(bytes.Buffer)

	data = AddDefaultData(data)

	_ = t.Execute(buffer, data)

	// render
	_, err := buffer.WriteTo(w)
	if err != nil {
		log.Println(err)
	}

	/*
		parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
		err := parsedTemplate.Execute(w, nil)
		if err != nil {
			fmt.Println("error parseing template:", err)
			return
		}
	*/
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	// get the tmpl files from templates dir
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return cache, err
	}

	//range through all tmpl files
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return cache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return cache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return cache, err
			}
		}

		cache[name] = ts
	}

	return cache, nil
}

//First version of cache below, more simple way of doing things
/*
var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	//check if template is already stored in the cache

	_, inMap := tc[t]

	if !inMap {
		//Create the template  and add to cache
		log.Println("adding template to cache")
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}

	} else {
		//Its already in the cache
		log.Println("Pulling template from cache")
	}

	tmpl = tc[t]

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}

}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t), "./templates/base.layout.tmpl",
	}

	//parse the template
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	//add to cache
	tc[t] = tmpl

	return nil
}
*/
