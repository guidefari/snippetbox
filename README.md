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
