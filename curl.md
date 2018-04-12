# Using Curl For Ad Hoc Testing Of RESTful Microservices

There are plenty of tools available for testing RESTful microservices today. Most of them, e.g. SoapUI, are comprehensive solutions and best fit for creating testing suites. Using such a tool to check a single faulty endpoint would be an overkill.
So what should you choose for ad hoc testing instead? There are simplified GUI tools, e.g. Postman, and many developers are happy with them. But if you are after ultimate performance and love command line, there is a better option - curl. In this post I'll show how to check RESTful endpoints using curl with a lot of examples.

## curl
Many developers use curl just to fetch remote files and even do not suspect what a powerful tool they run. curl can transfer data from or to a server, using FILE, FTP, FTPS, GOPHER, HTTP, HTTPS, IMAP, IMAPS, LDAP, LDAPS, POP3, POP3S, RTMP, RTSP, SCP, SFTP, SMB, SMBS, SMTP, SMTPS, TELNET and TFTP protocols. It has lots of options for almost every edge case.
For testing RESTful endpoints, however, you don't need all that power. Basically, HTTP and HTTPS is all you need. I'll use HTTP to simplify examples, but feel free to experiment with HTTPS if you feel like.
Mac and Linux users should already have curl installed out of the box. Windows users can downloaded it from here.

## Simple service in Go
The best way to learn something new is to play and experiment with it. curl is not an exception, so I've rustled up a simple service in Go for curl experiments:

```golang
package main

import (  
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
)

type RequestSummary struct {  
    URL     string
    Method  string
    Headers http.Header
    Params  url.Values
    Auth    *url.Userinfo
    Body    string
}

func main() {  
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        bytes, err := ioutil.ReadAll(r.Body)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        rs := RequestSummary{
            URL:     r.URL.RequestURI(),
            Method:  r.Method,
            Headers: r.Header,
            Params:  r.URL.Query(),
            Auth:    r.URL.User,
            Body:    string(bytes),
        }

        resp, err := json.MarshalIndent(&rs, "", "\t")
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        w.Write(resp)
        w.Write([]byte("\n"))
    })

    http.ListenAndServe(":8080", nil)
    fmt.Println("Exiting...")
}
```

The service is running on localhost:8080 and returns the most interesting pieces of request in json format, so that you can see what is received. Below is an example of what service returns for a GET request:

```
$ curl -X GET http://localhost:8080/something
{
    "URL": "/something",
    "Method": "GET",
    "Headers": {
        "Accept": [
            "*/*"
        ],
        "User-Agent": [
            "curl/7.43.0"
        ]
    },
    "Params": {},
    "Auth": null,
    "Body": ""
}
```

## Checking RESTful endpoints
Before jumping in with both feet, let's first get familiar with some curl options commonly used with all HTTP methods:


-X or --request specifies HTTP method (defaulting to GET), e.g -X POST.

-H or --header sets request header. E.g. -H "Content-Type:application/json" sets content type to application/json.

## GET
Most often you'll need to call get with some parameters:

```
$ curl -X GET "http://localhost:8080/something?param1=value1&param2=value2"
{
    "URL": "/something?param1=value1\u0026param2=value2",
    "Method": "GET",
    "Headers": {
        "Accept": [
            "*/*"
        ],
        "User-Agent": [
            "curl/7.43.0"
        ]
    },
    "Params": {
        "param1": [
            "value1"
        ],
        "param2": [
            "value2"
        ]
    },
    "Auth": null,
    "Body": ""
}
```

Note that url with parameters needs to be escaped to prevent shell from interpreting &.
If you need to urlencode your data, use -G and --data-urlencode instead:

```
$ curl -G --data-urlencode "message=hello world" http://localhost:8080/something
{
    "URL": "/something?message=hello%20world",
    "Method": "GET",
    "Headers": {
        "Accept": [
            "*/*"
        ],
        "User-Agent": [
            "curl/7.43.0"
        ]
    },
    "Params": {
        "message": [
            "hello world"
        ]
    },
    "Auth": null,
    "Body": ""
}
```

--data-urlencode, as it follows from its name, urlencodes your parameter. It's a POST method's option, so you have to add -G option, which append all data specified with -d, --data, --data-binary or --data-urlencode to the URL with a ? separator.

## POST and PUT
For POST and PUT you often need to pass some data in request body. With -d or --data you can pass data directly in command line or in a file using @filename syntax. In both cases you need to set Content-Type to application/json.
If data is small, it could be convenient to pass it directly in command line:

```
$ curl -X POST -H "Content-Type:application/json" -d '{"first": "John", "last": "Dow"}' http://localhost:8080/something
{
    "URL": "/something",
    "Method": "POST",
    "Headers": {
        "Accept": [
            "*/*"
        ],
        "Content-Length": [
            "32"
        ],
        "Content-Type": [
            "application/json"
        ],
        "User-Agent": [
            "curl/7.43.0"
        ]
    },
    "Params": {},
    "Auth": null,
    "Body": "{\"first\": \"John\", \"last\": \"Dow\"}"
}
```

In most cases however, it's preferred to serve data from a file:

```
$ cat data.json
{
        "first": "John",
        "last": "Dow"
}
```
```
$ curl -X PUT -H "Content-Type:application/json" -d @data.json http://localhost:8080/something
{
    "URL": "/something",
    "Method": "PUT",
    "Headers": {
        "Accept": [
            "*/*"
        ],
        "Content-Length": [
            "33"
        ],
        "Content-Type": [
            "application/json"
        ],
        "User-Agent": [
            "curl/7.43.0"
        ]
    },
    "Params": {},
    "Auth": null,
    "Body": "{\t\"first\": \"John\",\t\"last\": \"Dow\"}"
}
```

Note that --data strips out carriage returns and newlines when reading from a file. Use --data-binary if you need your data to be passed intact:

```
$ curl -X PUT -H "Content-Type:application/json" --data-binary @data.json http://localhost:8080/something
{
    "URL": "/something",
    "Method": "PUT",
    "Headers": {
        "Accept": [
            "*/*"
        ],
        "Content-Length": [
            "37"
        ],
        "Content-Type": [
            "application/json"
        ],
        "User-Agent": [
            "curl/7.43.0"
        ]
    },
    "Params": {},
    "Auth": null,
    "Body": "{\n\t\"first\": \"John\",\n\t\"last\": \"Dow\"\n}\n"
}
```
## DELETE
Delete is less often used in practice:
```
$ curl -X DELETE http://localhost:8080/collection/id
{
    "URL": "/collection/id",
    "Method": "DELETE",
    "Headers": {
        "Accept": [
            "*/*"
        ],
        "User-Agent": [
            "curl/7.43.0"
        ]
    },
    "Params": {},
    "Auth": null,
    "Body": ""
}
```

Basically, this is enough for ad hoc testing restful microservices. In the next post I'll write on more advanced
curl options, which help to smooth off the rough edges in day to day usage.
