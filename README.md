# chrome-browserback-demo
A demo site for browser back in chrome.

URL: https://chrome-browserback-demo.herokuapp.com/index

## How to run in local environment

```
$ PORT=3000 go run main.go
```

## How to deploy using herku

```
# Create a heroku app
$ heroku login
$ heroku create <app name>

# Deploy
$ git push heroku main
```
