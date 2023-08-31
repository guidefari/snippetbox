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
