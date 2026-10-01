---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/paultyng-terraform-provider-unifi-blob-head-docs-resources-site-md-c032fbdf
title: "import using the API/UI ID"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/paultyng-terraform-provider-unifi-blob-head-docs-resources-site-md-c032fbdf.md
source_anchor: ""
source_lines: [1, 16]
sha256: c6a63a7cf9857ce080bb8d58324b3f579df9e070d6df7818cc78b2c0261f0373
---

# import using the API/UI ID

| page_title | unifi_site Resource - terraform-provider-unifi | 
|---|---|
| subcategory |  | 
| description | unifi_site manages Unifi sites | 
unifi_site manages Unifi sites
resource "unifi_site" "mysite" {
  description = "mysite"
}
- description (String) The description of the site.
- id (String) The ID of the site.
- name (String) The name of the site.
Import is supported using the following syntax:
# import using the API/UI ID
terraform import unifi_site.mysite 5fe6261995fe130013456a36
# import using the name (short ID)
terraform import unifi_site.mysite vq98kwez
