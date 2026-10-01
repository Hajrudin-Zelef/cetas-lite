---
id: collect-261001-huawei/huawei/hedex-api-pages-edoc1100277644-aem10221-04-resources-vrp-dc-vrp-seamless-mpls-cf-a05efc52
title: "hedex-api-pages-edoc1100277644-aem10221-04-resources-vrp-dc-vrp-seamless-mpls-cf-a05efc52"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/hedex-api-pages-edoc1100277644-aem10221-04-resources-vrp-dc-vrp-seamless-mpls-cf-a05efc52.md
source_anchor: ""
source_lines: [1, 7]
sha256: dbb560a4648842e36a8c530b7d86bb8f7a5b1870e31847f8d11da7bbeebc2b1b
---

# hedex-api-pages-edoc1100277644-aem10221-04-resources-vrp-dc-vrp-seamless-mpls-cf-a05efc52

After configuring inter-AS seamless MPLS, you can check established LSPs and the connectivity of BGP LSPs between a CSG and an MASG.
Prerequisites
Inter-AS seamless MPLS has been configured.
Procedure
- Run the display ip routing-table command on a CSG or an MASG to check the routes to the peer end.
- Run the display mpls lsp command to check LSP information.
- Run the ping lsp [ -a source-ip | -c count | -exp exp-value | -h ttl-value | -m interval | -r reply-mode | -s packet-size | -t time-out | -v ] * bgp destination-iphost mask-length [ ip-address ] command on a CSG or MASG to check the BGP LSP connectivity.
