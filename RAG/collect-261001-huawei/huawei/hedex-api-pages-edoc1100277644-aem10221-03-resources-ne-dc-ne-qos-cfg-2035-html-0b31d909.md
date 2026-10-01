---
id: collect-261001-huawei/huawei/hedex-api-pages-edoc1100277644-aem10221-03-resources-ne-dc-ne-qos-cfg-2035-html-0b31d909
title: "hedex-api-pages-edoc1100277644-aem10221-03-resources-ne-dc-ne-qos-cfg-2035-html-0b31d909"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/hedex-api-pages-edoc1100277644-aem10221-03-resources-ne-dc-ne-qos-cfg-2035-html-0b31d909.md
source_anchor: ""
source_lines: [1, 7]
sha256: b763ddc50b6fc1711be17231de7d945585fa02c6491c270149368e198d3d627d
---

# hedex-api-pages-edoc1100277644-aem10221-03-resources-ne-dc-ne-qos-cfg-2035-html-0b31d909

After QPPB is configured, you can view QPPB information.
Context
You can run the display commands in any view to check QPPB running information. 
Procedure
- Run the display qppb local-policy configuration policy-name command to check the configuration of a specific QPPB local policy.
- Run the display qppb local-policy statistics interface interface-type interface-number [ qos-local-id qos-local-id ] { inbound | outbound } command to check the statistics about a specific QPPB local policy.
- Run the display ip routing-table { ip-address [ [ mask ] | [ mask-length ] ] |  | host-name host-name } verbose command to check the QoS local ID carried in a specified route.
