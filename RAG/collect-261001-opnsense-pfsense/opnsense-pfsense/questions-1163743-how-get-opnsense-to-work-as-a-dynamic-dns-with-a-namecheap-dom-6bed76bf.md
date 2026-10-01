---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/questions-1163743-how-get-opnsense-to-work-as-a-dynamic-dns-with-a-namecheap-dom-6bed76bf
title: "How get OpnSense to work as a dynamic dns with a NameCheap domain?"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-opnsense-pfsense/questions-1163743-how-get-opnsense-to-work-as-a-dynamic-dns-with-a-namecheap-dom-6bed76bf.md
source_anchor: ""
source_lines: [1, 18]
sha256: d33f9251b413d2ed84cacbcc1e0750dce23345ec5be6b40b5904fa9b8bbeec47
---

# How get OpnSense to work as a dynamic dns with a NameCheap domain?

*Score : -1 | Source : https://serverfault.com/questions/1163743/how-get-opnsense-to-work-as-a-dynamic-dns-with-a-namecheap-domain*

I have a domain from NameCheap and I'm running the latest OpnSense on my router, but trying to set up the DynDns gives me an error in my log file:
   <ResponseString>Validation error; not found; domain name(s)</ResponseString> 
   <ResponseNumber>316153</ResponseNumber>  
   <Description>Domain name not found</Description>
I've seen other people having this problem, but nothing comprehensive as a guide.

---

### Reponse — score 1

There are quite a few steps to getting NameCheap domain working as a dyndns with OpnSense:
Let's say you want to use the domain example.com
@.example.comServices>DynDns tab.
service use "NameCheap"username use the name of your domain (e.g. "example.com")Hostname(s) use "@" your domain name so No Records updated. A record not Found;Interface which means get the IP to name the domain to from the IP on your ethernet/wireless cardInterface to monitor which is in the case of having multiple ethernet cards. If "WAN" is an option, you probably want that one.
