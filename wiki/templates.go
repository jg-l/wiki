package wiki

const showTpl = `<!DOCTYPE html>
<html>
	<head>
	<!-- required in all modes -->
	<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/hack/0.8.0/hack.css">
	<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/hack/0.8.0/dark.css">




	</head>
	<body class="hack dark" style="padding: 0.5em;">
		<form action="{{.Path}}" method="GET">
    	<button class="btn btn-info btn-ghost" type="submit">Edit</button>
		</form>
		<div class="container" style="padding: 0em 1em; border: 1px solid #ccc;">
			{{if .Created.IsZero }}

			{{ else }}
				<p>Created on: {{ .Created.Month }} {{ .Created.Day }} {{ .Created.Year }} {{ .Created.Hour }}:{{ .Created.Minute }}</p>
			{{ end }}
			{{if .Modified.IsZero }}

			{{ else }}
				<p>Last modified: {{ .Modified.Month }} {{ .Modified.Day }} {{ .Modified.Year }} {{ .Modified.Hour }}:{{ .Modified.Minute }}</p>
			{{ end }}
			{{.Text}} 
		</div>
	</body>
</html>`

const editTpl = `<!DOCTYPE html>
<html>
	<head>
	<!-- required in all modes -->
	<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/hack/0.8.0/hack.css">
	<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/hack/0.8.0/dark.css">
	</head>
	<body class="hack dark" style="padding: 0.5em;">
		<div class="container">
			<form action="{{.Path}}" method="POST">
				<input class="btn btn-info" type="submit" value="Update" /><br />

				<textarea name="text" class="form-control" style="width: 100%; height: 40em;">{{.Text}}</textarea>
			</form>
		</div>
	</body>
</html>`

const emptyPageString = `# Empty page
So this is an empty page
`
