package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/peraltafederico/go-rules/internal/middleware"
)

type RequestBody struct {
	Trace   bool    `json:"trace"`
	Context Context `json:"context"`
}

type Context struct {
	Company Company `json:"company"`
}

type Company struct {
	Turnover int    `json:"turnover"`
	Type     string `json:"type"`
	Country  string `json:"country"`
}

func InitRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		var reqBody RequestBody

		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		url := fmt.Sprintf("https://%s.gorules.io/api/projects/%s/evaluate/%s", os.Getenv("DOMAIN"), os.Getenv("PROJECT_ID"), os.Getenv("DOCUMENT_PATH"))

		jsonData, err := json.Marshal(reqBody)

		if err != nil {
			log.Fatal(err)
			return
		}

		req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))

		req.Header.Add("accept", "application/json")
		req.Header.Add("content-type", "application/json")
		req.Header.Add("X-Access-Token", os.Getenv("ACCESS_TOKEN"))

		res, err := http.DefaultClient.Do(req)

		if err != nil {
			log.Fatal(err)
			return
		}

		defer res.Body.Close()
		body, _ := io.ReadAll(res.Body)

		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	})

	return r
}

// DlqH36z5aHa5bbEZSpGYNGMv
