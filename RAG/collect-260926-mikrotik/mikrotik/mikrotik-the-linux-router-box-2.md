---
id: collect-260926-mikrotik/mikrotik/mikrotik-the-linux-router-box-2
title: "Mikrotik - The Linux Router Box"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["agent"]
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/mikrotik-the-linux-router-box.md
source_anchor: ""
source_lines: [168, 178]
sha256: 13f0e5fbded114ca8c7a338197518bb2d528a6bb199fcaf24701755a119038c0
---

# Mikrotik - The Linux Router Box

I'm monitoring the device via SNMP, which provides some basic data but not much. The RouterOS also provides an API, allowing to query or manipulate the operation and configuration of the device. I found a library (librouteros) and a check_mk monitoring agent for mikrotik, which uses that API to provide more details for the router(fan speed, temperature, vrrp interface state) operation.

## Quirks and Problems I had

- I wanted to load the export-File and just got an error message which I did not understand. Through googling I found the issue(wrong extension of the file)
- Reset and load export-file did not work, because of the missing delay. I had to google to get the reason for this.
- The File upload tool ("/tool fetch" works for uploads and downloads) is a bit clunky. scp from outside is definitely the better choice. Maybe Winbox makes that easier...
- When writing a script, there was a restriction, that certain elements of the RouterOS can not manipulate Global Variables(security restriction). I found it by googling around and as of now workarounds are available, but it was a bit hard to get it running nevertheless.
- The Mikrotik Forum Administration seems quite dead. I registered my account and as security measures all posts of new users are moderated. Even after 2 weeks the posts weren't approved.
- When using the Mikrotik-internal editor for a script and the cursor goes beyond the end of the screen, the editor freezes the session and you have to open a new session.
- Files greater than 4K size (e. g. a full export) can not be edited with the Mikrotik-internal editor.
