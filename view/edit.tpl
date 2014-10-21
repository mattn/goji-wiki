<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>{{ page.Title }}</title>
	<link rel="stylesheet" href="/assets/css/style.css" media="all">
</head>
<body>
	<div id="content">
		<form action="{{ page.URL }}" method="POST">
		<h1 class="title">{{ page.Title }}</h1>
		<textarea id="body" name="body" cols="80" rows="20">{{ page.Body }}</textarea><br />
		<input type="submit" value="Save">
		<input type="button" value="Cancel" onclick="location.href='{{ wiki.URL }}'">
		</form>
	</div>
</body>
</html>
