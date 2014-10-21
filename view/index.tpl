<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>Wiki Pages</title>
	<link rel="stylesheet" href="/assets/css/style.css" media="all">
</head>
<body>
	<div id="content">
		<h1 class="title">Wiki Pages</h1>
		<ul>
		{% for page in pages %}
		<li><a href="{{ page.URL }}">{{ page.Title }}</a> <span class="date">{{ page.UpdatedAt | date: "2006/01/02 03:04:05" }}</span></li>
		{% endfor %}</ul>
	</div>
</body>
</html>
