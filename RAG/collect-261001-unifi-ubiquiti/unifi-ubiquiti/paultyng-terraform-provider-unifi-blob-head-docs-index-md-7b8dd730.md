---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/paultyng-terraform-provider-unifi-blob-head-docs-index-md-7b8dd730
title: "paultyng-terraform-provider-unifi-blob-head-docs-index-md-7b8dd730"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/paultyng-terraform-provider-unifi-blob-head-docs-index-md-7b8dd730.md
source_anchor: ""
source_lines: [1, 21]
sha256: 6f6ff2b2ecba5889a842c7167306dc5b546efff1804ace303073fe7faec3eb53
---

# paultyng-terraform-provider-unifi-blob-head-docs-index-md-7b8dd730

| layout |  | 
|---|---|
| page_title | Provider: Unifi | 
| description | The Unifi provider provides resources to interact with a Unifi controller API. | 
The Unifi provider provides resources to interact with a Unifi controller API.
It is not recommended to use your own account for management of your controller. A user specific to Terraform is recommended. You can create a Limited Admin with Local Access Only and provide that information for authentication. Two-factor authentication is not supported in the provider.
provider "unifi" {
  username = var.username # optionally use UNIFI_USERNAME env var
  password = var.password # optionally use UNIFI_PASSWORD env var
  api_url  = var.api_url  # optionally use UNIFI_API env var
  # you may need to allow insecure TLS communications unless you have configured
  # certificates for your controller
  allow_insecure = var.insecure # optionally use UNIFI_INSECURE env var
  # if you are not configuring the default site, you can change the site
  # site = "foo" or optionally use UNIFI_SITE env var
}
- allow_insecure (Boolean) Skip verification of TLS certificates of API requests. You may need to set this totrue if you are using your local API without setting up a signed certificate. Can be specified with theUNIFI_INSECURE environment variable.
- api_url (String) URL of the controller API. Can be specified with theUNIFI_API environment variable. You should NOT supply the path (/api ), the SDK will discover the appropriate paths. This is to support UDM Pro style API paths as well as more standard controller paths.
- password (String) Password for the user accessing the API. Can be specified with theUNIFI_PASSWORD environment variable.
- site (String) The site in the Unifi controller this provider will manage. Can be specified with theUNIFI_SITE environment variable. Default:default
- username (String) Local user name for the Unifi controller API. Can be specified with theUNIFI_USERNAME environment variable.
