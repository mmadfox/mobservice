package main

import (
	"encoding/json"
	"github.com/Shaked/gomobiledetect"
	"github.com/gorilla/mux"
	"github.com/pmylund/go-cache"
	"mobservice/props"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var c *cache.Cache

func init() {
	c = cache.New(time.Hour*48, time.Minute*30)
}

func getCache() *cache.Cache {
	return c
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/device.json", DeviceHandler)
	router.HandleFunc("/os.json", OsHandler)
	router.HandleFunc("/phone.json", PhoneHandler)
	router.HandleFunc("/tablet.json", TabletHandler)
	router.HandleFunc("/browser.json", BrowserHandler)
	http.ListenAndServe(":6062", Middleware(router))
}

func DeviceHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	userAgent, _ := url.QueryUnescape(query.Get("ua"))
	userAgent = strings.Trim(userAgent, " ")

	var device *props.Device
	var cid string = "ua" + userAgent

	if fromCache, found := getCache().Get(cid); found {
		device = fromCache.(*props.Device)
	} else {
		m := mobiledetect.NewMobileDetect(r, nil)
		m.SetUserAgent(userAgent)
		device = props.NewDevice()
		device.Configure(m)
		getCache().Set(cid, device, 0)
	}

	js, err := json.Marshal(&device)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(js)
}

func PhoneHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	term := query.Get("term")

	js, err := json.Marshal(props.FindPhone(term))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(js)
}

func TabletHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	term := query.Get("term")

	js, err := json.Marshal(props.FindTablet(term))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(js)
}

func OsHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	term := query.Get("term")

	js, err := json.Marshal(props.FindOs(term))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(js)
}

func BrowserHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	term := query.Get("term")

	js, err := json.Marshal(props.FindBrowser(term))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(js)
}

func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, r)
	})
}
