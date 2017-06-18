# netlify-rest-api

[![Go Report Card](https://goreportcard.com/badge/github.com/poga/netlify-rest-api)](https://goreportcard.com/report/github.com/poga/netlify-rest-api)

Publish your data as static RESTful JSON API to [Netlify](https://netlify.com).

## Install

`go get github.com/poga/netlify-rest-api`

## Usage

```
$ netlify-rest-api -h
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
$ git clone git@github.com:poga/netlify-rest-api.git
$ cd netlify-rest-api/example
$ netlify-rest-api -file users.csv -host http://YOUR-NETLIFY-DOMAIN -out out
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

