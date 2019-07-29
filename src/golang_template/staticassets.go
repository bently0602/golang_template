package main 

const (
index_html = `<!DOCTYPE html>
<html>
<head>
	<meta charset="UTF-8"> 
	<title>{{.Title}}</title>
	<link href="https://fonts.googleapis.com/css?family=Open+Sans" rel="stylesheet">
	<style>
	html, body {
		width: 100%;
		margin: 0;
		padding: 0;
	}
	body {
		font-family: 'Open Sans', sans-serif;
		background: #fafafa;
	}
	.nyancat {
		margin: 0 auto;
		width: 400px;
		display: block;
		padding: 2rem;
	}
	</style>
</head>
<body>
	<marquee><h1>{{.Title}}</h1></marquee>
	<img class="nyancat" src="{{.PathPrefix}}/static/nyancat.gif" />
</body>
</html>
`
)
