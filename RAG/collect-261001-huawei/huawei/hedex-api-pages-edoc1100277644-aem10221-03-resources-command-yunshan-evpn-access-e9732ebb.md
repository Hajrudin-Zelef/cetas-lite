---
id: collect-261001-huawei/huawei/hedex-api-pages-edoc1100277644-aem10221-03-resources-command-yunshan-evpn-access-e9732ebb
title: "hedex-api-pages-edoc1100277644-aem10221-03-resources-command-yunshan-evpn-access-e9732ebb"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/hedex-api-pages-edoc1100277644-aem10221-03-resources-command-yunshan-evpn-access-e9732ebb.md
source_anchor: ""
source_lines: [1, 9]
sha256: 2bb80262b1d0794c49eac08d8ccc8b6f9435714e58a77fa8eda552a75472edeb
---

# hedex-api-pages-edoc1100277644-aem10221-03-resources-command-yunshan-evpn-access-e9732ebb

The evpn access vll convergence separate disable command enables the coupling flag in the EVPN-accessing-VLL direction.
The undo evpn access vll convergence separate disable command enables the decoupling flag in the EVPN-accessing-VLL direction.
By default, the decoupling flag is enabled in the EVPN-accessing-VLL direction.
evpn access vll convergence separate disable
undo evpn access vll convergence separate disable
Usage Scenario
To enable the coupling flag in the EVPN-accessing-VLL direction, run the evpn access vll convergence separate disable command. This configuration allows EVPN services to be forwarded only through an LDP, TE, or LDP over TE LSP as the public network tunnel.
Precautions
If the coupling flag is enabled in the EVPN-accessing-VLL direction, only LDP, TE, or LDP over TE LSPs are supported on the public network side of EVPN accessing VLL.
