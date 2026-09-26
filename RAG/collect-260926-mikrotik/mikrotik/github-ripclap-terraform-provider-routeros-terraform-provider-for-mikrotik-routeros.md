---
id: collect-260926-mikrotik/mikrotik/github-ripclap-terraform-provider-routeros-terraform-provider-for-mikrotik-routeros
title: "github-ripclap-terraform-provider-routeros-terraform-provider-for-mikrotik-routeros"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["attribution"]
source: docs/RAG/lot-mikrotik/RouterOS/github-ripclap-terraform-provider-routeros-terraform-provider-for-mikrotik-routeros.md
source_anchor: ""
source_lines: [1, 71]
sha256: 9fc115ca02863447664adac91d9eb8dd590c57c0f66c4670fa189869b7e5fc05
---

# github-ripclap-terraform-provider-routeros-terraform-provider-for-mikrotik-routeros

Manage MikroTik RouterOS 7 configuration with OpenTofu, over the REST API or the binary API, with or without TLS.

**This is a fork** of
terraform-routeros/terraform-provider-routeros,
licensed under MPL-2.0 and not affiliated with or endorsed by the upstream
project. See FORK.md for what differs and NOTICE for
attribution.


It registers **415 resources** covering 388 RouterOS menus, against 254
upstream, so a whole device can be held in state and drift detected across all
of it rather than a subset. COVERAGE.md lists what it adds.

Every list menu also has a **data source** of the same name — 308 in total —
returning the menu's entries under `entries`, narrowed by an optional `filter`:

```
data "routeros_interface_bridge_port" "ports" {
  filter = { bridge = "bridge" }
}
output "member_interfaces" {
  value = data.routeros_interface_bridge_port.ports.entries[*].interface
}
```
Enable the REST API on the router first: create a certificate under
`/certificate` and enable the `www-ssl` service under `/ip/service` using it.
MikroTik's documentation
covers this.

```
terraform {
  required_providers {
    routeros = {
      source  = "ripclap/routeros"
      version = "~> 2.0"
    }
  }
}
provider "routeros" {
  hosturl  = "https://my.router.local"
  username = "my_username"
  password = "my_super_secret_password"
}
```
`hosturl` accepts `https://` and `http://` for the REST API, and `apis://`
(8729) or `api://` (8728) for the binary API. The binary API is faster and is
the one that completes a read against a very large routing table; see
FORK.md.

Published to the OpenTofu Registry, where the resource and data source reference lives. It is not published to the HashiCorp Terraform Registry.

Tested against RouterOS 7.23.x, built with Go 1.25. Compatibility is only claimed within RouterOS 7.

- **2.0.2 is withdrawn and will not install.** Use 2.0.3 or later.
- Upstream 1.43 changed the schemas of `routeros_routing_bgp_connection` ,`routeros_ipv6_neighbor_discovery` and`routeros_interface_wireguard_peer` .
For the first two, remove the resource from state and import it again.

Each release ships a GPG-signed checksum file, a CycloneDX SBOM per archive, and a build provenance attestation:

```
gh attestation verify terraform-provider-routeros_2.0.3_linux_amd64.zip \
  --repo ripclap/terraform-provider-routeros
```
| FORK.md | What differs from upstream, and transport behaviour | 
| COVERAGE.md | Resources this fork adds | 
| CHANGELOG.md | Release history | 
| CONTRIBUTING.md | Building, testing and adding resources | 
| RELEASING.md | How a release is cut, and why a published version is never rebuilt | 

*Not affiliated with MikroTik. MikroTik and RouterOS are trademarks of
Mikrotikls SIA.*
