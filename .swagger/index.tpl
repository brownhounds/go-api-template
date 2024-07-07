<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <meta name="description" content="SwaggerIU" />
    <title>SwaggerUI</title>
    <link
      rel="stylesheet"
      href="https://unpkg.com/swagger-ui-dist@5.17.14/swagger-ui.css"
    />
  </head>
  <body>
    <div id="swagger-ui"></div>
    <script
      src="https://unpkg.com/swagger-ui-dist@5.17.14/swagger-ui-bundle.js"
      crossorigin
    ></script>
    <script
      src="https://unpkg.com/swagger-ui-dist@5.17.14/swagger-ui-standalone-preset.js"
      crossorigin
    ></script>
    <script>
      window.onload = () => {
        window.ui = SwaggerUIBundle({
          spec: JSON.parse({{JSON_SCHEMA}}),
          dom_id: "#swagger-ui",
        });
      };
    </script>
