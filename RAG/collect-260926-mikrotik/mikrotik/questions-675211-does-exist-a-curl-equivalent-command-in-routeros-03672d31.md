---
id: collect-260926-mikrotik/mikrotik/questions-675211-does-exist-a-curl-equivalent-command-in-routeros-03672d31
title: "questions-675211-does-exist-a-curl-equivalent-command-in-routeros-03672d31"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/questions-675211-does-exist-a-curl-equivalent-command-in-routeros-03672d31.md
source_anchor: ""
source_lines: [1, 16]
sha256: f7788bf5fb9c1d1c2c42a9bb84d48117abc069da8cb0440e4fd70f587e738d93
---

# questions-675211-does-exist-a-curl-equivalent-command-in-routeros-03672d31

Just starting with routeros and after searching the official docs just fetch appears
I need to do a POST request to REST API but I cannot find any curl like command. Is it possible to install curl somehow? 
Unfortunately there is no way to install curl on a standard Mikrotik installation.
The only way to make http requests from a vanilla mikrotik installation is using fetch as you already found.
But it does not support POST requests, only GET.
There are a couple solutions to this problem depending on your situation.
If you are using 5.x version on an x86 machine then there is an ISO out there that will install a debian underneath Mikrotik so you can then install any debian package you need on your Mikrotik using ssh and apt-get.
curl will not be able to use mikrotik's scripting language or any other info directly from mikrotik though.
Since the ISO not only patches mikrotik to be able to install debian packages, but also cracks the mikrotik licensing, for obvious reasons I cannot tell you any more details about it.
The other (legal) method if you are using Routerboard hardware (with a few exceptions) and you have enough RAM (at least 64MB), would be to create a metarouter (a VM essentially) and use an openwrt image on it.
Then on openwrt you can install any package you like and do what you need.
Again, you won't be able to access any mikrotik's internals with this method.
Routeros manual page about fetch (https://wiki.mikrotik.com/wiki/Manual:Tools/Fetch) says it does support POST:
Sending information to a remote host
It is possible to use HTTP POST request to send information to a remote server, that is prepared to accept it. In the following example, we send geographic coordinates to a PHP page:
/tool fetch http-method=post http-header-field="Content-Type: application/json" http-data="{\"lat\":\"56.12\",\"lon\":\"25.12\"}" url="http://testserver.lv/index.php"
