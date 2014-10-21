<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>Wiki Pages</title>
	<link rel="stylesheet" href="/assets/css/style.css" media="all">
<script>
function editNewPage() {
  var title = prompt('Title');
  if (title) location.href = '{{ wiki.URL }}wiki/' + encodeURIComponent(title) + '/edit';
  return 0;
}
</script>
</head>
<body>
	<div id="content">
		<h1 class="title">Wiki Pages</h1>
		<ul>
		{% for page in pages %}
		<li><a href="{{ page.URL }}">{{ page.Title }}</a> <span class="date">{{ page.UpdatedAt | to_localdate | date: "2006/01/02 03:04:05" }}</span></li>
		{% endfor %}</ul>
		<a href="#" onclick="return editNewPage()">Edit new page</a>
	</div>
</body>
</html>
