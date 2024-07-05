// Web-server for encode-decode base64 and decode jwt
// Веб-сервер для кодирования-декодирования base64 и декодирования jwt

package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"

	//"path/filepath"
	"sync"
)

// var (
// 	crtFile = filepath.Join(".", "cert", "server.crt")
// 	keyFile = filepath.Join(".", "cert", "server.key")
// )

func main() {
	log.SetPrefix("Event main: ")
	log.SetFlags(log.Lshortfile)

	LogInfo("web-server listening on :8060")

	// Мультиплексор запросов. Router of http-requests.
	mux := http.NewServeMux()
	mux.HandleFunc("/encode/", encoder)
	mux.HandleFunc("/decode/", decoder)
	mux.HandleFunc("/jwtdecode/", jwtdecoder)
	log.Fatal(http.ListenAndServe(":8060", mux))
	//log.Fatal(http.ListenAndServeTLS(":8060", crtFile, keyFile, mux))
}

// Хэндл енкодера. Handle of encoder
func encoder(w http.ResponseWriter, r *http.Request) {

	fmt.Println("\nRequest of client:")

	// Параметры http-запроса. Parameters of headers
	fmt.Fprintf(w, "Method = %s\nURL = %s\nProto = %s\n", r.Method, r.URL, r.Proto)
	fmt.Printf("Method = %s\nURL = %s\nProto = %s\n", r.Method, r.URL, r.Proto)

	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		fmt.Printf("Header[%q] = %q\n", k, v)
	}

	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Printf("Host = %q\n", r.Host)

	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	fmt.Printf("RemoteAddr = %q\n", r.RemoteAddr)

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Printf("Form[%q] = %q\n", k, v)
	}

	var mu sync.Mutex
	chg := make(chan string, 1)

	mu.Lock()
	defer mu.Unlock()
	//pars := r.URL.Path[8:]
	pars := r.URL.Query().Get("secret")
	if pars == "" {
		log.Println("Empty secret")
		fmt.Fprintf(w, "\nEmpty secret %v", nil)
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		encoded := base64.StdEncoding.EncodeToString([]byte(pars))
		log.Println("Encoded base64: ", encoded)
		chg <- encoded
	}()

	fmt.Fprintf(w, "\nEncoded base64: %v", <-chg)

	go func() {
		wg.Wait()
		for range <-chg {
			//nil
		}
		close(chg)
	}()

}

// Хэндл декодера. Handle of decoder
func decoder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request of client:")

	// Параметры http-запроса. Parameters of headers
	fmt.Fprintf(w, "Method = %s\nURL = %s\nProto = %s\n", r.Method, r.URL, r.Proto)
	fmt.Printf("Method = %s\nURL = %s\nProto = %s\n", r.Method, r.URL, r.Proto)

	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		fmt.Printf("Header[%q] = %q\n", k, v)
	}

	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Printf("Host = %q\n", r.Host)

	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	fmt.Printf("RemoteAddr = %q\n", r.RemoteAddr)

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Printf("Form[%q] = %q\n", k, v)
	}

	var mu sync.Mutex
	chg := make(chan string, 1)
	chr := make(chan string, 1)

	mu.Lock()
	defer mu.Unlock()
	//pars := r.URL.Path[8:]
	pars := r.URL.Query().Get("secret")
	if pars == "" {
		log.Println("Empty secret")
		fmt.Fprintf(w, "\nEmpty secret %v", nil)
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		wg.Done()

		decoded, err := base64.StdEncoding.DecodeString(pars)
		if err != nil {
			log.Println("Decoded error: ", err)
			chr <- err.Error()
		}
		// decoded, err := base64.URLEncoding.DecodeString(pars)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		log.Println("Decoded base64: ", string(decoded))
		chg <- string(decoded)
	}()

	//fmt.Fprintf(w, "\nDecoded error: %v", <-chr)
	fmt.Fprintf(w, "\nDecoded base64: %v", <-chg)

	go func() {
		wg.Wait()
		for range <-chg {
			//nil
		}
		for range <-chr {
			//nil
		}
		close(chg)
		close(chr)
	}()
}

// Хэндл jwt-декодера. Handle of jwt-decoder
func jwtdecoder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request of client: ")

	// Параметры http-запроса. Parameters of headers
	fmt.Fprintf(w, "Method = %s\nURL = %s\nProto = %s\n", r.Method, r.URL, r.Proto)
	fmt.Printf("Method = %s\nURL = %s\nProto = %s\n", r.Method, r.URL, r.Proto)

	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		fmt.Printf("Header[%q] = %q\n", k, v)
	}

	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Printf("Host = %q\n", r.Host)

	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	fmt.Printf("RemoteAddr = %q\n", r.RemoteAddr)

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Printf("Form[%q] = %q\n", k, v)
	}

	var mu sync.Mutex
	chg := make(chan string, 2)
	chr := make(chan string, 1)

	mu.Lock()
	defer mu.Unlock()
	pars := r.URL.Query().Get("jwt")
	if pars == "" {
		log.Println("Empty jwt")
		fmt.Fprintf(w, "\nEmpty jwt %v", nil)
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		wg.Done()

		jwt := []byte(pars)
		d := new(TokenDecoder)

		header, body, err := d.Decode(jwt)
		if err != nil {
			log.Fatalf("Error: %v\n", err)
			chr <- err.Error()
		} else {
			log.Printf("Header: %v, \nBody: %v", string(header), string(body))
		}
		chg <- string(header)
		chg <- string(body)
	}()

	fmt.Fprintf(w, "\nDecoded JWT: %v, %v", <-chg, <-chg)

	go func() {
		wg.Wait()
		for range <-chg {
			//nil
		}
		for range <-chr {
			//nil
		}
		close(chg)
		close(chr)
	}()
}

// Logger
var logger = log.Default()

func LogInfo(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	logger.Printf("[Info]: %s\n", msg)
}
