package main

import (
	"encoding/json"
	"net/http"

	. "github.com/stevelacy/daz"
)

var head = H("head",
	H("meta", Attr{"charset": "utf-8"}),
	H("meta", Attr{
		"name":    "viewport",
		"content": "width=device-width, initial-scale=1.0",
	}),
	H("title", "expensive relay"),
)

func handleWebpage(w http.ResponseWriter, r *http.Request) {
	items := make([]HTML, 0, 200)
	iter := b.db.NewIter(nil)
	for iter.First(); iter.Valid(); iter.Next() {
		pubkey := string(iter.Key())
		var entity Entity
		if err := json.Unmarshal(iter.Value(), &entity); err != nil {
			continue
		}
		items = append(items, H("tr",
			H("td",
				H("code",
					pubkey),
			),
			H("td",
				H("a", Attr{
					"href": entity.URL}, entity.URL),
			),
		))
	}

	body := H("body",
		H("h1", "rsslay"),
		H("p", "rsslay turns RSS or Atom feeds into ",
			H("a", Attr{
				"href": "https://github.com/fiatjaf/nostr"}, "Nostr"),
			" profiles.",
		),
		H("h2", "How to use"),
		H("ol",
			H("li", "Get the blog URL or RSS or Atom feed URL and paste below;"),
			H("li", "Click the button to get its corresponding public key"),
			H("li", "Add the relay ",
				H("code", "wss://"+b.Domain),
				" to your Nostr client",
			),
			H("li", "Follow the feed's public key from your Nostr client."),
		),
		H("form", Attr{
			"action": "/create",
			"method": "GET",
			"class":  "my-4",
		},
			H("label",
				H("input", Attr{
					"name":        "url",
					"type":        "url",
					"placeholder": "https://.../feed",
				}),
			),
			H("button", "Get Public Key"),
		),

		H("h2", "Some of the existing feeds"),
		H("table", items),
		H("h2", "Source Code"),
		H("p", "You can find it at ",
			H("a", Attr{"href": "https://github.com/fiatjaf/rsslay"},
				"https://github.com/fiatjaf/rsslay"),
		),
	)

	w.Header().Set("content-type", "text/html")
	w.Write([]byte(
		H("html",
			head,
			body,
		)()))
}

func handleLnurlRegister(w http.ResponseWriter, r *http.Request) {}

func handleInvoicePaid(w http.ResponseWriter, r *http.Request) {}
