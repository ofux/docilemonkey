# Docile Monkey

Docile Monkey is an extremely simple HTTP server that responds what you want it to respond.

It can respond with the HTTP status you want, with the body you want and in the delay you want.

It may be useful for testing.
For example, you may use it to easily check how your application reacts when it receives unexpected HTTP responses (code 500, whatever)
or when the server takes a lot of time to respond.
It can also help to test circuit breakers.

## Query params

- `s` will be the HTTP status code of the response (ex: 200, 404, 500, etc.)
- `t` time to wait before responding (ex: 10s, 200ms, etc.)
- `b` will be the body of the response
- `bb` if parameter 'b' has no value and if bb=1 then the body of the request (if any) will be sent back in the response

## Usage

### SAAS

Docile Monkey is available here for convenience: https://europe-west1-miscellaneous-ofux-stuff.cloudfunctions.net/docilemonkey

Examples: 
```sh
curl -v 'https://europe-west1-miscellaneous-ofux-stuff.cloudfunctions.net/docilemonkey?s=418&t=20ms'

curl -v -X POST \
  'https://europe-west1-miscellaneous-ofux-stuff.cloudfunctions.net/docilemonkey?s=418&t=20ms&bb=1' \
  -d '{ "foo": "bar" }'
```

**Do not call it from your automated tests.**

**Do not abuse or I will stop providing this endpoint.**

If you plan to do many calls to this server, get it and deploy it on your own infrastructure (see below).

### As a standalone server

Download Docile Monkey binary and launch it.

Parameters:
- **port**: port on which the server will listen for requests. Default value: **8080**

Example:

```
$ ./docilemonkey -port=8080
```

### In your tests (for Go projects)

You can use it in your test through the httptest.Server

```go
func TestSomething(t *testing.T) {

    // create a new docilemonkey test server. This is equivalent to httptest.NewServer(http.HandlerFunc(docilemonkey.Handler))
	ts := docilemonkey.NewTestServer()
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Error(err)
	}

    //...
}
```

