---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-78-2
title: "Overview"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-78.md
source_anchor: ""
source_lines: [134, 146]
sha256: 4bb3ac5cbe39bbebce36e03a662ad18495f002cecdf110962feb0e86e5bb3e9c
---

# Overview

SNMP is limited to *ftp,reboot,**read,write,test,romon* script policies. If the script has greater policies than  *ftp,reboot,**read,write,test,romon* - then the script will not be executed, make sure your scripts do not exceed the mentioned policies.

## Running scripts with GET

It is possible to run **/system scripts** via SNMP GET request of the script OID (since 6.37). For this to work SNMP community with write permission is required. OIDs for scripts can be retrieved via the SNMPWALK command as the table is dynamic.

Add script:

Get the script OID table

To run the script use table 18

SNMP is limited to *ftp,reboot,**read,write,test,romon* script policies. If the script has greater policies than  *ftp,reboot,**read,write,test,romon* - then the script will not be executed, make sure your scripts do not exceed the mentioned policies.
