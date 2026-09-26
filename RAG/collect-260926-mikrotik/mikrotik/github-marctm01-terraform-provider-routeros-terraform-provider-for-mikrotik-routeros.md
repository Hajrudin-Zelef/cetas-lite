---
id: collect-260926-mikrotik/mikrotik/github-marctm01-terraform-provider-routeros-terraform-provider-for-mikrotik-routeros
title: "github-marctm01-terraform-provider-routeros-terraform-provider-for-mikrotik-routeros"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/github-marctm01-terraform-provider-routeros-terraform-provider-for-mikrotik-routeros.md
source_anchor: ""
source_lines: [1, 98]
sha256: 518ac73282ee6d961faa6a19ae90bc84b418ddb0113a115ed9002abac3c18805
---

# github-marctm01-terraform-provider-routeros-terraform-provider-for-mikrotik-routeros

**Note**: In release 1.43, the resource schemas have been changed:

- `routeros_routing_bgp_connection`
- `routeros_ipv6_neighbor_discovery`
- `routeros_interface_wireguard_peer`

For the first two to work correctly, you must remove the resource state (`terraform state rm <name>`) and import it again (`terraform import [options] <name> <id>`).

This provider allows you to configure Mikrotik routers using old API or REST API, using or not using TLS. Compatibility testing is only performed within ROS version 7.x.

From version 1.0.0, the provider has been rewritten by vaerh, and their fork has now been merged. This version drastically improves adding new endpoints to the provider, enabling significantly easier development. vaerh has been added as a maintainer to this project.

*We are not affiliated in any way with Mikrotik or the development of RouterOS*

To get started with the provider, you first need to enable the REST API on your router. You can follow the Mikrotik documentation on this, but the gist is to create an SSL cert (in `/system/certificates`) and enable the `web-ssl` service (in `/ip/services`) which uses that certificate. After that, include the following in your Terraform manifests:

```
terraform {
  required_providers {
    routeros = {
      source = "terraform-routeros/routeros"
    }
  }
}
provider "routeros" {
  hosturl  = "(http|https|api|apis)://my.router.local[:port]"
  username = "my_username"
  password = "my_super_secret_password"
}
```
For more in-depth documentation about each of the resources and datasources, please read the documentation on Hashicorp's Provider registry

- go 1.24.2 and ROS 7.12, 7.15, 7.16 (stable)

For a detailed changelog, please see the changelog.md.

This version of the module greatly simplifies the process of adding new resources. You are welcome!

You can build the provider locally to test fixes by following these intructions:

- Build and copy the provider where Terraform reads it

```
go build *.go && \
mkdir -p ~/.terraform.d/plugins/terraform.local/local/routeros/1.0.0/$(uname -s | tr '[:upper:]' '[:lower:]')_$(uname -m) && \
mv main ~/.terraform.d/plugins/terraform.local/local/routeros/1.0.0/$(uname -s | tr '[:upper:]' '[:lower:]')_$(uname -m)/terraform-provider-routeros_v1.0.0
```
- Change provider from

```
required_providers {
  routeros = {
    source  = "terraform-routeros/routeros"
    version = "1.85.1"
  }
}
```
to

```
required_providers {
  routeros = {
    source  = "terraform.local/local/routeros"
    version = "1.0.0"
  }
}
```
- Clean your providers, init and apply
- Alternatively, you can edit/create ~/.terraformrc add add a provider installation block like:

```
provider_installation {
  dev_overrides {
     "terraform-routeros/routeros" = "/path/to/your/git/clone"
  }
  direct {
  }
}
```
and then build the provider using

```
go build -o terraform-provider-routeros *.go
```
in order for Terraform to find it.

Sometimes RouterOS might introduce a breaking change on a property. You can easilfy contribute to the provider by following these intructions:

- Edit `routeros/mikrotik_resource_drift.yaml` . Add the resource used as well as the old property name and the new one
- Perform the generator. It should edit file `routeros/mikrotik_resource_drift.go` .

```
cd routeros/
go run ../tools/drift/main.go
```
- Submit your changes!

Here is a example of pull request.
