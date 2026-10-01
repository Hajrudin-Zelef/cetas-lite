---
id: collect-261001-general-networking/general-networking/troubleshooting-tip-avoid-errors-when-adding-an-interface-to-an-sd-wan-community
title: "troubleshooting-tip-avoid-errors-when-adding-an-interface-to-an-sd-wan-community"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/troubleshooting-tip-avoid-errors-when-adding-an-interface-to-an-sd-wan-community.md
source_anchor: ""
source_lines: [1, 29]
sha256: 4b76637a0432b69e969f51098396c9fb97aeab8a7f3993650ab7aee5e592886c
---

# troubleshooting-tip-avoid-errors-when-adding-an-interface-to-an-sd-wan-community

**Description**


This article describes how to avoid common errors when adding an interface to an SD-WAN.


**Scope**


FortiGate.

**Solution**


Before configuring FortiGate interfaces as SD-WAN members, it is necessary to remove or redirect existing configuration references to those interfaces in routes and security policies. 

This includes the default Internet access policy that’s included with many FortiGate models. 

Note that after removing the routes and security policies, traffic cannot reach the WAN ports through the FortiGate. 

Redirecting the routes and policies to reference other interfaces prevents the need to create them again later.


For example:


1) WAN2 is the physical interface to add the SD-WAN member into, but WAN2 has a reference in the static route and policies.

**Network options -> Interfaces**, then scroll left to the specific interface. Select the appropriate reference number in the reference column.
