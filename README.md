# Docile Monkey

Docile Monkey is an extremely simple HTTP server that responds what you want it to respond.

It can respond with the HTTP status you want, with the body you want and in the delay you want.

It may be useful for testing.
For example, you may use it to easily check how your application reacts when it receives unexpected HTTP responses (code 500, whatever)
or when the server takes a lot of time to respond.
It can also help to test circuit breakers.

## Usage

### SAAS

Docile Monkey is available here for convenience: 

**Do not call it from your automated tests.**

**Do not abuse or I will stop providing this server.**

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

