# Fall15CMPE273Lab2

A simple “Hello World” REST API in Go to understand how handler and Http router work, and how to parse the request and send the response.

GET /hello/xxxx
        Response:
                Hello, xxxx!

Finally, open this URL in a web browser: http://localhost:8080/hello/foo
Now, implement POST that takes JSON request and returns JSON response.
POST /hello

Request:
{
   "name": "foo"
}
Response:
{
   "greeting" : "Hello, foo!"
}
