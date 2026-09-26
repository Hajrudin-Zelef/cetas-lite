---
id: collect-260926-mikrotik/mikrotik/questions-1142081-scripting-with-variables-in-routeros-2ce6e502
title: "questions-1142081-scripting-with-variables-in-routeros-2ce6e502"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/questions-1142081-scripting-with-variables-in-routeros-2ce6e502.md
source_anchor: ""
source_lines: [1, 13]
sha256: 351f331f18461af31c83409f584546f92fcecc9235a58bc030a355c235a6d505
---

# questions-1142081-scripting-with-variables-in-routeros-2ce6e502

I seem to be at a complete loss as to how Global variables work in RouterOS. My goal is to create a script which will make calls to other infrastructure components when DHCP leases are changed.
For each of the numerous DHCP servers, I intend to make a call to a system script which contains the actual logic.
DHCP Server Configuration:
:global leaseActIP
:log info ("Server Script Got: $leaseActIP")
/system script run LeaseChange
And in the System script LeaseChange
:global leaseActIP
:log info ("System Script Got: $leaseActIP")
However log output is:
Server Script Got: 10.3.111.145
System Script Got:
Question: How can I pass the variables available in the DHCP server script over to the System script?
