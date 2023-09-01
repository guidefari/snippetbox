- Following the book [Let's Go](https://lets-go.alexedwards.net/)
- `chmod +x ./run.sh` if it refuses to execute

## TIL - Block action

> This acts like the {{template}} action, except it allows you to specify some default content if the template being invoked doesn’t exist in the current template set.
> In the context of a web application, this is useful when you want to provide some default content (such as a sidebar) which individual pages can override on a case-by-case basis if they need to.

```html
{{define "base"}}
<h1>An example template</h1>
{{block "sidebar" .}}
<p>My default sidebar content</p>
{{end}} {{end}}
```

## TIL - Additional info about Go's file server

- It sanitizes all request paths by running them through `path.Clean`
  - this helps stop directory traversal attacks
- [Range requests](https://benramsey.com/blog/2008/05/206-partial-content-and-range-requests)
  are fully supported. This is great for resumable downloads!

## TIL - All incoming http requests are handled concurrently

Yes, blazingly fast, but be aware of race conditions when accessing
shared resources

## TIL - Command line flags

`addrVariable := flag.String("addr", ":4000", "HTTP network address")`

- This takes the flag `addr` from the command line
- Sets a default value of `":4000"`
- Gives it some helper text

## Do a deep dive on function signatures

```go
// Change the signature of the home handler so it is defined as a method against // *application.
func (app *application) home(w http.ResponseWriter, r *http.Request) {

// as seen in handler.go
```

- my question is especially centered around `func (app *application)`
- learning module resolution and package management will probably be helpful tooo
  - that \*application exists in `main.go`
