package views

var homeSrc = `
{{ define "content" }}
<div class="container mt-5">
  <h1>Start typing to find a product...</h1>
  <form id="search" action="/search" method="POST">
    <div class="row">
      <div class="col">
        <div class="form-floating">
          <input class="form-control" type="text" id="vendor" name="vendor" list="vendor-list" placeholder="apache">
          <label for="vendor">Vendor</label>
          <datalist id="vendor-list"></datalist>
        </div>
      </div>
      <div class="col">
        <div class="form-floating">
          <input class="form-control" type="text" id="name" name="name" list="name-list" placeholder="http_server">
          <label for="name">Product</label>
          <datalist id="name-list"></datalist>
        </div>
      </div>
      <div class="col">
        <div class="form-floating">
          <input class="form-control" type="text" id="version" name="version" list="version-list" placeholder="2.4">
          <label for="version">Version</label>
          <datalist id="version-list"></datalist>
        </div>
      </div>
    </div>
    <div class="mt-3">
      <button type="submit" class="btn btn-primary">Show Product Details</button>
    </div>
  </form>
</div>
{{ end }}
`
