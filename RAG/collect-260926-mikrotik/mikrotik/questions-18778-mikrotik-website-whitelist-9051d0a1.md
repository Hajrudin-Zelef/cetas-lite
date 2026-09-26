---
id: collect-260926-mikrotik/mikrotik/questions-18778-mikrotik-website-whitelist-9051d0a1
title: "questions-18778-mikrotik-website-whitelist-9051d0a1"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/misc/questions-18778-mikrotik-website-whitelist-9051d0a1.md
source_anchor: ""
source_lines: [1, 10]
sha256: eb037bc8aa0d0a43e6b8bde2c6c15d6c2b24f674284025e832f4ca4d7f96385b
---

# questions-18778-mikrotik-website-whitelist-9051d0a1

On the mikrotik wiki there are lot of examples on how to use the http proxy feature.
Anyway:
On the IP/Proxy menu you can set the options of the embedded http proxy: in the "/ip proxy access" part you can define you access policies in terms of http host names or ip addresses.
Then, you have two options to make it working as a trasparent proxy.
A fully trasparent proxy means that the users do not have to configure it in their browser.
http://wiki.mikrotik.com/wiki/How_to_make_transparent_web_proxy
BUT, doing so, you cannot filter HTTPS protocol.
If you want to filter also https using your proxy, I can advise to use a "proxy auto-discover" mechanism, like WPAD.
http://findproxyforurl.com/wpad-introduction/
Is this enough to start?
