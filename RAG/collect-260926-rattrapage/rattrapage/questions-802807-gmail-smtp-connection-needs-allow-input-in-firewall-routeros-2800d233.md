---
id: collect-260926-rattrapage/rattrapage/questions-802807-gmail-smtp-connection-needs-allow-input-in-firewall-routeros-2800d233
title: "questions-802807-gmail-smtp-connection-needs-allow-input-in-firewall-routeros-2800d233"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-rattrapage/ai-llm/questions-802807-gmail-smtp-connection-needs-allow-input-in-firewall-routeros-2800d233.md
source_anchor: ""
source_lines: [1, 15]
sha256: f6067df62e16fb7f5e57951a3149cf0a3eba9e368716d46f342227ac6d42959e
---

# questions-802807-gmail-smtp-connection-needs-allow-input-in-firewall-routeros-2800d233

I am trying to set-up Mikrotik (RouterOS v6.24) for sending emails.
I have google account and I am using SMTP connection. I know that I need to allow outbound connection for it to work. But when I try to send an email, the inbound filter will block connection and sending an email will fail.
16:37:04 firewall,info input: in:ether1-WAN out:(none), src-mac 00:13:60:16:4f:c6, proto TCP (SYN,ACK), 74.125.128.108:587->x.x.x.x:5462
When I disable the input filter. The email will be send correctly.
Why do I need to allow input connection for sending an email?
RouterOS settings:
   address: 74.125.128.108
       port: 587
  start-tls: yes
       from: xxxxx@gmail.com
       user: xxxxx
   password: xxxxxxxx
last-status: failed
Command used:
send to=xxxxx@xxxx.com from="xxxx@gmail.com" subject="test email" body="test body"
