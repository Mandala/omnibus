// Copyright (c) 2017 Fadhli Dzil Ikram. All rights reserved.
// This source code is brought to you under MIT license that can be found
// on the LICENSE file.

package muxer

import (
	"html/template"
	"net/http"
)

var tmplNotFound *template.Template

func init() {
	tmplNotFound = template.Must(template.New("").Parse(`<!DOCTYPE html>
<html lang="en">
<head>
  <title>Not Found</title>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <style>
    body {
      font-family: Georgia, serif;
    }
    .resource {
      font-family: "Lucida Console", Monaco, monospace;
    }
  </style>
</head>
<body>
  <h1>Page Not Found</h1>
  <p>
    Resource <span class="resource">{{.}}</span> was not found on the server.
    Please double check the entered URL and try again.
  </p>
</body>
</html>
`))
}

// notFoundHandler acts as default not found page for the router
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	tmplNotFound.Execute(w, r.URL.Path)
}
