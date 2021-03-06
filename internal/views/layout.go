package views

var layoutSrc = `<!doctype html>
<html lang="en">

<head>
  <meta charset="utf-8">
  <title>Vulnsearch</title>
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <link rel="stylesheet" href="/vendor/bootstrap-v5.0.0-beta3.min.css" />
</head>

<body>
  <div class="mt-5">
    {{ block "content" . }}{{ end }}
  </div>
  <script type="module" src="/application.js"></script>
</body>

</html>`
