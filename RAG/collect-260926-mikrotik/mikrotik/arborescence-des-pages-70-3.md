---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-70-3
title: "Overview"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-70.md
source_anchor: ""
source_lines: [147, 168]
sha256: 5852f55cb0f850cd61646a6122b663d8ea083fab57c08e9e8e2b18a14881edd0
---

# Overview

Specifying MRRU means enabling MP (Multilink PPP) over a single link. This protocol is used to split big packets into smaller ones.  Their MRRU is hardcoded to 1614. This setting is useful to overcome PathMTU discovery failures. The MP setting should be enabled on both peers.

The default *keepalive-timeout* value of 10s is OK in most cases. If you set it to 0, the router will not disconnect clients until they explicitly log out or the router is restarted. To resolve this problem, the one-session-per-host property can be used.

# Quick Example

## PPPoE Client

To configure MikroTik RouterOS to be a PPPoE client, just add a PPPoE-client with the following parameters as in the example:

## PPPoE Server

To configure MikroTik RouterOS to be an Access Concentrator (PPPoE Server):

- add an IP address pool for the clients from 10.0.0.2-10.0.0.5;
- add PPP profile;
- add PPP secret (username/password);
- add the PPPoE server itself;

Notes

Its not recommended to use large amount of pppoe-clients on one device.
