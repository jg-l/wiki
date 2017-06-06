package wiki

const showTpl = `<!DOCTYPE html>
<html>
	<head>

	<!-- Google Fonts -->
	<link rel="stylesheet" href="//fonts.googleapis.com/css?family=Roboto:300,300italic,700,700italic">

	<!-- CSS Reset -->
	<link rel="stylesheet" href="//cdn.rawgit.com/necolas/normalize.css/master/normalize.css">

	<!-- Milligram CSS minified -->
	<link rel="stylesheet" href="//cdn.rawgit.com/milligram/milligram/master/dist/milligram.min.css">

	<style>

	#footer {
  position: fixed;
  right: 0;
  bottom: 0;
  left: 0;
  padding: 1rem;
  background-color: #efefef;
  text-align: center;
			  }

	</style>


	<title>{{.Name }} - Wiki</title>

	</head>
	<body class="hack dark" style="padding: 0.5em;">
		<div class="container" style="padding: 0em 1em;">
		<form action="{{.Path}}" method="GET">
    	<button class="btn btn-info btn-ghost" type="submit">Edit</button>
		</form>
			{{.Text}} 
			<div id="footer">
			{{if .Created.IsZero }}

			{{ else }}
				<span class="float-left">Created on: {{ .Created.Format "2006 Jan _2, 3:04:05 PM" }} </span>
			{{ end }}
			{{if .Modified.IsZero }}

			{{ else }}
				<span class="float-right">Last modified: {{ .Modified.Format "2006 Jan _2, 3:04:05 PM" }} </span>

			{{ end }}</div>
		</div>
	</body>
</html>`

const editTpl = `<!DOCTYPE html>
<html>
	<head>
	<!-- Google Fonts -->
	<link rel="stylesheet" href="//fonts.googleapis.com/css?family=Roboto:300,300italic,700,700italic">

	<!-- CSS Reset -->
	<link rel="stylesheet" href="//cdn.rawgit.com/necolas/normalize.css/master/normalize.css">

	<!-- Milligram CSS minified -->
	<link rel="stylesheet" href="//cdn.rawgit.com/milligram/milligram/master/dist/milligram.min.css">
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
