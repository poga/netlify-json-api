# netlify-rest-api

[![Go Report Card](https://goreportcard.com/badge/github.com/poga/netlify-rest-api)](https://goreportcard.com/report/github.com/poga/netlify-rest-api)

Publish your data as static RESTful JSON API to [Netlify](https://netlify.com).

## Install

`go get github.com/poga/netlify-rest-api`

## Usage

Use our example:

```
$ git clone git@github.com:poga/netlify-rest-api.git
$ cd netlify-rest-api/example
$ netlify-rest-api users.csv http://YOUR-NETLIFY-DOMAIN out
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

