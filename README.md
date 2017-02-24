# example-RESTapp in Microsoft Azure
Repo for an simple boilerplate RESTful web API written in Go. This doesn't really do much.

[![GoDoc](https://godoc.org/github.com/ZachtimusPrime/example-RESTapp?status.svg)](https://godoc.org/github.com/ZachtimusPrime/example-RESTapp)
[![Build Status](https://travis-ci.org/ZachtimusPrime/example-RESTapp.svg?branch=master)](https://travis-ci.org/ZachtimusPrime/example-RESTapp) 
[![Coverage Status](https://coveralls.io/repos/github/ZachtimusPrime/example-RESTapp/badge.svg?branch=master)](https://coveralls.io/github/ZachtimusPrime/example-RESTapp?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/ZachtimusPrime/example-RESTapp)](https://goreportcard.com/report/github.com/ZachtimusPrime/example-RESTapp) 

## Table of Contents ##

* [Installation](#installation)
* [Usage](#usage)


## Setup ##
This is a boilerplate API used to test cloud hosting platforms (Azure, AWS, Google Cloud). 

There are two keys to success on Azure. The first is to include the following section when defining the port that your API will listen on. This is because Azure randomly decides which port to make your API listen on when building and running your app. Look inside main.go for the full server setup. The second is that all your go files must reside in root.

```go
// Start server
	ServicePort := "8080"
	portConfig := os.Getenv("HTTP_PLATFORM_PORT")
	if portConfig != "" {
		ServicePort = portConfig
	}
  
```

Visit the following location to see the app in action:

[http://example-restapp.azurewebsites.net](http://example-restapp.azurewebsites.net)
