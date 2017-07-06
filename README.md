# netlify-json-api

[![Go Report Card](https://goreportcard.com/badge/github.com/poga/netlify-json-api)](https://goreportcard.com/report/github.com/poga/netlify-json-api)

Publish your data as a RESTful JSON API on [Netlify](https://netlify.com).

## Install

`go get github.com/poga/netlify-json-api`

## Usage

```
$ netlify-json-api -h
  -file string
    	data file
  -host string
    	host domain
  -id string
    	ID column names (seperate by ",") (default "id")
  -out string
    	output directory
  -perPage int
    	items per page (default 10)
  -type string
    	resource type name
```

Use our example:

```
$ git clone git@github.com:poga/netlify-json-api.git
$ cd netlify-json-api/example
$ netlify-json-api -file users.csv -host http://YOUR-NETLIFY-DOMAIN -out out
```

Then, deploy `out` to netlify.

Now you have a RESTful JSON API! Try these URLs:

* `GET http://YOUR-NETLIFY-DOMAIN/users.json`
* `GET http://YOUR-NETLIFY-DOMAIN/users.json?page=1`
* `GET http://YOUR-NETLIFY-DOMAIN/users/1.json`


## Todos

- [ ] Support other source data. sql? json?
- [ ] Auto deploy to netlify
- [ ] `POST`, `PUT`, and `DELETE` with proxy

## License

The MIT License

