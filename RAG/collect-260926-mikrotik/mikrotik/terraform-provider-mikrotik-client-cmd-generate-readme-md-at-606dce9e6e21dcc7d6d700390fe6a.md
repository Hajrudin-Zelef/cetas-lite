---
id: collect-260926-mikrotik/mikrotik/terraform-provider-mikrotik-client-cmd-generate-readme-md-at-606dce9e6e21dcc7d6d700390fe6a
title: "terraform-provider-mikrotik-client-cmd-generate-readme-md-at-606dce9e6e21dcc7d6d700390fe6ad07b414e71"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/misc/terraform-provider-mikrotik-client-cmd-generate-readme-md-at-606dce9e6e21dcc7d6d700390fe6ad07b414e71.md
source_anchor: ""
source_lines: [1, 12]
sha256: eb091f7f80fb60368f0b4a488ac475715555e98b92b900951e5f3fb92da06dbf
---

# terraform-provider-mikrotik-client-cmd-generate-readme-md-at-606dce9e6e21dcc7d6d700390fe6ad07b414e71

This tool allows generating MikroTik resources for API client.

You can generate MikroTik resource and save it to file using remote RouterOS instance (requires valid credentials in environment)

`$ go run ./client/cmd/generate resource -query -basePath /ip/arp -outFile client/arp.go`
or using offline definition file created with inspection tool:

`$ go run ./client/cmd/generate resource -basePath /ip/arp -definitionFile arp.json -outFile client/arp.go`
If you have already generated resource in `arp.go` file, you can also generate a basic test file for it:

`$ go run ./client/cmd/generate test -sourceFile client/arp.go -outFile client/arp_test.go`
carefully review the test file, as codegen tool is still experimental.
