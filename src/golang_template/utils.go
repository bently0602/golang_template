package main

import (
	"io"
	"io/ioutil"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"net/http"
	"html/template"
	"bytes"
	"runtime"
	"log"

	"crypto/rand"
	"crypto/subtle"
	
	"github.com/spf13/viper"
)

func OpenBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	if err != nil {
		log.Fatal(err)
	}
}

func NewUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("naomi-%x%x%x%x%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}

func ReplaceSliceTokens(x []string, key string, value string) []string {
	for n, val := range x {
		if val == key {
			x[n] = value
		}
	}
	return x
}

func GetExePath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	return exPath
}

// BasicAuth wraps a handler requiring HTTP basic auth for it using the given
// username and password and the specified realm, which shouldn't contain quotes.
//
// Most web browser display a dialog with something like:
//
//    The website says: "<realm>"
//
// Which is really stupid so you may want to set the realm to a message rather than
// an actual realm.
//
// http.HandleFunc("/", BasicAuth("admin", "123456", "Please enter your username and password for this site")
//
func BasicAuth(handler http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
		username := viper.GetString("basicAuthUsername")
		password := viper.GetString("basicAuthPassword")
		realm := viper.GetString("basicAuthRealm")
		user, pass, ok := r.BasicAuth()

        if !ok || subtle.ConstantTimeCompare([]byte(user), []byte(username)) != 1 || subtle.ConstantTimeCompare([]byte(pass), []byte(password)) != 1 {
            w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
            w.WriteHeader(401)
            w.Write([]byte("Unauthorised.\n"))
            return
        }

        handler(w, r)
    }
}

func CacheControlWrapper(h http.Handler) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")
        h.ServeHTTP(w, r)
    })
}

func ServePage(w http.ResponseWriter, page string) {
    exPath := GetExePath()
	file, err := ioutil.ReadFile(path.Join(exPath, "../static", page))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Fprint(w, string(file))
	}
}

func ReadPage(page string) string {
    exPath := GetExePath()
	file, err := ioutil.ReadFile(path.Join(exPath, "../static", page))
	if err != nil {
		return ""
	} else {
		return string(file)
	}
}

func RenderTemplate(templatePath string, data interface{}) (string, error) {
	pageSrc := ReadPage(templatePath)
	t := template.New(templatePath)
	t, _ = t.Parse(pageSrc)
	var tpl bytes.Buffer
	if err := t.Execute(&tpl, data); err != nil {
		return "", err
	} else {
		return tpl.String(), nil
	}
}

func RenderTemplateStaticAsset(src string, data interface{}) (string, error) {
	var tmplResults bytes.Buffer
	htmlTmpl, err := template.New("").Parse(src)
	if err != nil {
		return "", err
	}
	if err := htmlTmpl.Execute(&tmplResults, data); err != nil {
		return "", err
	}
	return tmplResults.String(), nil
}