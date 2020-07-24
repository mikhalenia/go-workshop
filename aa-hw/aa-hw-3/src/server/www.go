package server

import (
	"auto"
	"errors"
	"fmt"
	"logs"
	"mime"
	"net/http"
	"order"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
	"user"
)

var (
	autos          auto.Catalog
	orders         order.Catalog
	users          user.Catalog
	PATH_SEPARATOR string
	baseDir        string
)

func InitServer() {

	if runtime.GOOS == "windows" {
		PATH_SEPARATOR = "\\"
	} else {
		PATH_SEPARATOR = "/"
	}

	var err error
	baseDir, err = getBaseDir()
	if err != nil {
		logs.Logs(err.Error())
	}

	autos = auto.CreateCatalog()
	orders = order.CreateCatalog()
	users = user.CreateCatalog()
}

func StartWebServer() {
	httpMux := http.NewServeMux()
	//
	httpMux.Handle("/", http.HandlerFunc(requestHandler))
	http.ListenAndServe(":80", httpMux)
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	requestpath := r.URL.RequestURI()
	file := r.URL.Path

	if file == requestpath { /// File Request

		if file == "/" || file == "" { // default page
			file = "/index.html"
			bytes, contentType := readFile(file)
			w.Header().Set("Content-Type", contentType)
			w.Header().Set("Server", "goserver/0.1")
			w.WriteHeader(200)
			w.Write(bytes)
		} else { // Loading other html`s
			bytes, contentType := readFile(file)
			w.Header().Set("Content-Type", contentType)
			w.Header().Set("Server", "goserver/0.1")
			w.WriteHeader(200)
			w.Write(bytes)
		}
	} else { /// JSON Request
		htmlVrbls := getHTMLVariable(requestpath)

		for vr, vl := range htmlVrbls {
			switch vr {
			// Examples:
			// CMD: curl "http://localhost/?createorder=0&auto=287"
			// Browser: http://localhost/?createorder=0&auto=23
			case "createorder":
				autoId := strToInt(htmlVrbls["auto"])
				logs.Logs(fmt.Sprintf("Create order for auto id: %d", autoId))

				// Id, userId, autoId, Sum, TimeFrom, TimeTo, Status
				order := order.Order{orders.CreateId(), 1, autoId, 20, false}
				orders.SetOrder(order)
			// Examples:
			// CMD: curl "http://localhost/?paidorder=34type=none"
			// Browser: http://localhost/?paidorder=881&type=erip
			case "paidorder":
				id := strToInt(vl)
				o := orders.Table[id]
				_ = o.Payment(htmlVrbls["type"])
			}
		}
	}
}

func getHTMLVariable(query string) map[string]string {

	if i := strings.Index(query, "?"); i != -1 {
		query = query[i+1:]
	}

	//var mapVar map[string]string
	mapVar := map[string]string{}

	arr := strings.Split(query, "&")
	if len(arr) > 1 {
		for _, data := range arr {

			if i := strings.Index(data, "="); i != -1 {
				vr, vl := data[:i], data[i+1:]
				mapVar[vr] = vl
			}
		}
	} else {
		if i := strings.Index(query, "="); i != -1 {
			vr, vl := query[:i], query[i+1:]
			mapVar[vr] = vl
		}
	}
	return mapVar
}

func strToInt(str string) int {
	i, err := strconv.Atoi(str) // str to int
	if err != nil {
		return 0
	}
	return i
}

func intToStr(i int) string {
	str := strconv.Itoa(i)
	return str
}

func readFile(file string) ([]uint8, string) { // readFile (file string) bytes []uint8, contentType string

	reqFile, err := os.Open(baseDir + file)
	if err != nil {
		reqFile, err = os.Open(baseDir + file + ".html")
		if err != nil {
			file = PATH_SEPARATOR + "404.html"
			reqFile, err = os.Open(baseDir + file)
		}
	}
	fi, err := reqFile.Stat()
	contentType := mime.TypeByExtension(path.Ext(file))
	var bytes = make([]uint8, fi.Size())
	reqFile.Read(bytes)
	return bytes, contentType
}

func getBaseDir() (string, error) {
	str, err := getRootDir()
	str = str + PATH_SEPARATOR + "www"
	return str, err
}

func getRootDir() (string, error) {
	var dir string

	tdir, err := os.Getwd()
	if err != nil {
		return "", errors.New("Can't get root directory.")
	}
	arrayS := strings.Split(tdir, PATH_SEPARATOR)
	dir = strings.Join(arrayS[:len(arrayS)-2], PATH_SEPARATOR)
	if _, err := os.Stat(dir + PATH_SEPARATOR + "www"); os.IsNotExist(err) {
		dir = tdir
	}
	return dir, nil
}
