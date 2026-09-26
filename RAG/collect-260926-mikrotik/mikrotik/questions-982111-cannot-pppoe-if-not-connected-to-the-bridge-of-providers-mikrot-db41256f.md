---
id: collect-260926-mikrotik/mikrotik/questions-982111-cannot-pppoe-if-not-connected-to-the-bridge-of-providers-mikrot-db41256f
title: "questions-982111-cannot-pppoe-if-not-connected-to-the-bridge-of-providers-mikrot-db41256f"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/lot-mikrotik/forum/vlan-bridge/questions-982111-cannot-pppoe-if-not-connected-to-the-bridge-of-providers-mikrot-db41256f.md
source_anchor: ""
source_lines: [1, 3]
sha256: 2047e7dc6719aa23f202fc9ce09b015227d40f993deda6522f904f98220394bc
---

# questions-982111-cannot-pppoe-if-not-connected-to-the-bridge-of-providers-mikrot-db41256f

my italian provider open fiber just installed in my rack a new optical fiber GPON connectivity.
They passed into my office's walls an optical cable that ends into a zte zxhn f601 converter sc to rj45. Then they connected the zte to the wan interface of a basic mikrotik router board hEx lite.
From this point they gived to me pppoe authentication parameters, pppoe ip interface and a network of 8 static ips that I requested into the contract. To make a successful pppoe connection, I need to connect the second port of the hEx lite mikrotik to the wan interface of another router(in my case is a mikrotik too but this one is an enterprise version) and this is working perfectly. The point is that I want to find a way to remove this hEx lite mikrotik from my environment. Considering that I tried to connect directly with pppoe parameters from my enterprise mikrotik connected to the zte without success, I'm trying to understand what they have done on their hex mikrotik side to make it works. During the Installation I asked the technician if this hex mikrotik was something like a pppoe server from my side or if is taking a specific ip only for the monitoring. He told me that wan port and lan port are bridged and there is no pppoe server, they simply create a separate connection only for monitoring. So at this point I'm sure that I'm connecting directly to the outside with my enterprise mikrotik but for some reason if I'm not connected to this bridge the connections fails(there is no authentication error, simply it trying to connect without success) so I'm asking you were do you think the lock is, and what I can do with my mikrotik to overcame this. Many thanks
