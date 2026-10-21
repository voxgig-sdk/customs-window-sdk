# API Definition
 
API definition (OpenAPI etc) files live here.


## Changes from the vendor's published definition

`customs-window-openapi.json` is https://api.customswindow.com/schema/?format=json
with two changes:

1. `servers` is injected (`https://api.customswindow.com`). The published
   definition declares none.
2. `components.securitySchemes["Secret Key"].name` is changed from `"Basic Auth"`
   to `"Authorization"` (2026-10-21). `Basic Auth` is not a valid HTTP header name:
   fetch throws on it, and the server never reads it. Probed on 2026-10-21, the
   live API answers `Authorization: Basic <secret key>` with "Incorrect/invalid
   secret key" and every other placement with "Authentication credentials were
   not provided". The `Basic` prefix is set in `model/project.aontu`
   (`main.kit.config.auth`), because an apiKey scheme cannot express one.
