package main

import (
	"database/sql"
	"net/http"
	"newsfeed/platform/newsfeed"

	"github.com/go-chi/chi/v5"
	_ "github.com/mattn/go-sqlite3"
	"github.com/qkgo/yin"
)

func main() {
	db, _ := sql.Open("sqlite3", "./newsfeed.db")
	feed := newsfeed.NewFeed(db)

	// // feed.Add(newsfeed.Item{
	// // 	Content: "Hello!",
	// // })

	// items := feed.Get()

	// fmt.Println(items)

	r := chi.NewRouter()
	r.Use(yin.SimpleLogger)

	r.Get("/posts", func(w http.ResponseWriter, r *http.Request) {
		res, _ := yin.Event(w, r)
		items := feed.Get()
		res.SendJSON(items)
	})

	r.Post("/posts", func(w http.ResponseWriter, r *http.Request) {
		res, req := yin.Event(w, r)
		body := map[string]string{}
		req.BindBody(&body)
		item := newsfeed.Item{
			Content: body["content"],
		}
		feed.Add(item)
		res.SendJSON(204)
	})

	http.ListenAndServe(":3000", r)

}

// curl -i -X POST -H 'Content-Type: application/json' -d '{"content": "New item"}' http://localhost:3000/posts

// fetch('/posts', {
// 	method: 'POST',
// 	headers: {'content-type': 'application/json'},
// 	body: JSON.stringify({
// 		content: 'Hola'
// 	})
// })
