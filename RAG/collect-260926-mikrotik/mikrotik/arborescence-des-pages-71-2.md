---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-71-2
title: "Overview"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-71.md
source_anchor: ""
source_lines: [96, 115]
sha256: 8473da860753c43df3f659b7dd0dd433fcb638a3b03927e33f5660478deb5c9e
---

# Overview

- **certificate is not yet valid** - notBefore certificate date is after the current time;
- **certificate has expired** - certificate expiry date is before the current time;
- **cinvalid certificate purpose** - the supplied certificate cannot be used for the specified purpose;
- **cself signed certificate in a chain** - the certificate chain could be built up using the untrusted certificates but the root could not be found locally;
- **cunable to get issuer certificate locally** - CA certificate is not imported locally;
- **cserver's IP address does not match certificate** - server address verification is enabled, but the address provided in certificate does not match the server's address;

# Quick Example

## SSTP Client

In the following configuration example, e will create a simple SSTP client without using a certificate:

## SSTP Server

We will configure PPP secret for a particular user, afterwards simply enable an SSTP server:

In P2P setups network address will be same with other endpoint local address.

As with any other ppp tunnel, SSTP also supports BCP which allows it to bridge SSTP tunnel with a local interface. For example in setups where routers are connected to Internet through ether1, workstations and laptops are connected to ether2. Both local networks are routed through SSTP client, and they are not in the same broadcast domain BCP is used.
