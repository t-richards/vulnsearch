package views

var homeSrc = `
{{ define "content" }}
<div class="container">
  <h1>Start typing to find a product...</h1>
  <form action="/search" method="POST">
    <div class="row">
      <div class="col">
        <div class="form-group">
          <label for="vendor">Vendor</label>
          <input class="form-control" type="text" id="vendor" name="vendor" list="vendor-list" placeholder="apache">
          <datalist id="vendor-list"></datalist>
        </div>
      </div>
      <div class="col">
        <div class="form-group">
          <label for="name">Product</label>
          <input class="form-control" type="text" id="name" name="name" list="name-list" placeholder="http_server">
          <datalist id="name-list"></datalist>
        </div>
      </div>
      <div class="col">
        <div class="form-group">
          <label for="version">Version</label>
          <input class="form-control" type="text" id="version" name="version" list="version-list" placeholder="2.4">
          <datalist id="version-list"></datalist>
        </div>
      </div>
    </div>
    <div class="d-flex justify-content-end">
      <input type="submit" class="btn btn-primary" value="Show Product Details">
    </div>
  </form>
</div>
{{ end }}
`
