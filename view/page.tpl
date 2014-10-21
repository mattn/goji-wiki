<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>{{ page.Title }}</title>
	<link rel="stylesheet" href="/assets/css/style.css" media="all">
</head>
<body>
	<div id="content">
		<h1 class="title">{{ page.Title }}</h1>
		<div class="body">
		{{ page.Body | markdown }}
		</div>
		<span class="date">Updated at: {{ page.UpdatedAt | to_localdate | date: "2006/01/02 03:04:05" }}</span>
		<br />
		<a href="{{ wiki.URL }}">Page index</a>
		<a href="{{ page.URL }}/edit">Edit this page</a>
	</div>
</body>
</html>
