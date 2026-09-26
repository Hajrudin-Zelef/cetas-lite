---
id: collect-260926-mikrotik/mikrotik/runzeroinc-runzero-custom-integrations-blob-head-mikrotik-routeros-readme-md-be99a23d-4
title: "runzeroinc-runzero-custom-integrations-blob-head-mikrotik-routeros-readme-md-be99a23d"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/runzeroinc-runzero-custom-integrations-blob-head-mikrotik-routeros-readme-md-be99a23d.md
source_anchor: ""
source_lines: [272, 280]
sha256: 61a73e84efc8b32db31eac7aa9667d5ee0da1ac14e4c18776857b3d3071da34f
---

# runzeroinc-runzero-custom-integrations-blob-head-mikrotik-routeros-readme-md-be99a23d

- The router's own services. /ip/service lists which management services are enabled and on which ports — a genuinely useful exposure signal for the router asset, and a naturalService list.
- CAPsMAN-managed APs. /caps-man/remote-cap enumerates the access points a CAPsMAN controller manages, each a real device with a MAC, model, and version. Today those APs appear only if they also show up as discovery neighbors.
- RouterOS 6 via the binary API. Out of scope here, and it would need a hand-rolled protocol over socket.tcp . Worth revisiting only if v6 demand is real; the version floor is stated plainly instead.
- REST API — the base path, the JSON-string encoding rule, the .proplist /.query forms, the error envelope, and the verbatim/system/resource example.
- RouterOS documentation home — per-menu property tables for IP/ARP, IP/DHCP Server, IP/Neighbor discovery, Wireless, WifiWave2, CAPsMAN, System/Resource, and System/RouterBOARD.
- User management and policies — the policy list, including rest-api as distinct fromapi andweb .
- IP services — www andwww-ssl , and the certificate requirement.
- Reference clients worth reading for real payload shapes: socialwifi/RouterOS-api ,librouteros , and thecommunity.routeros Ansible collection.
- tikoci/restraml — per-version dumps of RouterOS's own property namespace, harvested over REST from real routers on both x86 and arm64. This is the source of the field names used here and is more current than the documentation pages.
