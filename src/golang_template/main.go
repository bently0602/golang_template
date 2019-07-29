package main

import (
	"fmt"
	"net/http"
	"log"
	"time"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

//go:generate go run ../../scripts/include_static_assets.go ./ main

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	r := mux.NewRouter()
	s := r.PathPrefix(viper.GetString("pathPrefix")).Subrouter()

	fmt.Printf("==============================================\n")
	fmt.Printf("* golang_template *\n")
	fmt.Printf("==============================================\n\n")
	fmt.Printf("   Listening on: %s\n", viper.GetString("httpBind"))
	if viper.GetString("pathPrefix") == "" {
	fmt.Printf("        Path is: %s\n", "\\")
	} else {
	fmt.Printf("        Path is: %s\n", viper.GetString("pathPrefix"))
	}
	fmt.Printf("   Read Timeout: %d\n", viper.GetInt("HTTPReadTimeout"))
	fmt.Printf("  Write Timeout: %d\n\n", viper.GetInt("HTTPWriteTimeout"))

	// static files
	s.PathPrefix("/static/").Handler(
		http.StripPrefix(viper.GetString("pathPrefix") + "/static/",
						(CacheControlWrapper(http.FileServer(http.Dir("./static"))))))

	////////////////////////////////////
	
	// index or default
	s.HandleFunc("/", (
		func(w http.ResponseWriter, r *http.Request) {
			data := struct {
				Title string
				PathPrefix string
				URL string
				Host string
			} {
				"golang_template",
				viper.GetString("pathPrefix"),
				r.URL.Path,
				r.Host,
			}

			tmpHtml, err := RenderTemplateStaticAsset(index_html, data)

			if err != nil {
				// generic error
				// http.Error(
				// 	w, 
				// 	http.StatusText(http.StatusInternalServerError),
				// 	http.StatusInternalServerError)
				
				// error with info 
				http.Error(w, err.Error(), http.StatusInternalServerError)
			} else {
				w.Header().Add("Content-Type", "text/html")
				fmt.Fprint(w, tmpHtml)
			}
	})).Methods("GET")

	// post handler
	s.HandleFunc("/posttest", (func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		
		PARAM_A := r.FormValue("PARAM_A")
		PARAM_B := r.URL.Query().Get("PARAM_B")
		fmt.Printf("PARAM_A: %s\nPARAM_B: %s", PARAM_A, PARAM_B)

		fmt.Fprint(w, "{'success': true}")
	})).Methods("POST")

	////////////////////////////////////

	srv := &http.Server{
		Handler:      r,
		Addr:         viper.GetString("httpBind"),
		WriteTimeout: time.Duration(viper.GetInt("HTTPWriteTimeout")) * time.Second,
		ReadTimeout:  time.Duration(viper.GetInt("HTTPReadTimeout")) * time.Second,
	}

	if viper.GetBool("openBrowser") {
		go OpenBrowser("http://" + viper.GetString("httpBind") + viper.GetString("pathPrefix"))
	}

	log.Fatal(srv.ListenAndServe())
}
