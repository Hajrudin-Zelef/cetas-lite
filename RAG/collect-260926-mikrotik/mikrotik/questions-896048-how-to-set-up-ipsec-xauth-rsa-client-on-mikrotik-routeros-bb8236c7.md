---
id: collect-260926-mikrotik/mikrotik/questions-896048-how-to-set-up-ipsec-xauth-rsa-client-on-mikrotik-routeros-bb8236c7
title: "questions-896048-how-to-set-up-ipsec-xauth-rsa-client-on-mikrotik-routeros-bb8236c7"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/questions-896048-how-to-set-up-ipsec-xauth-rsa-client-on-mikrotik-routeros-bb8236c7.md
source_anchor: ""
source_lines: [1, 17]
sha256: f683286bd7262895db9ca6e709770e68388cc5022364b02f65c90af7e4a80b47
---

# questions-896048-how-to-set-up-ipsec-xauth-rsa-client-on-mikrotik-routeros-bb8236c7

I suggest you check the MikroTik manual regarding IPsec. There's a lot of info and examples there.
Here's an XAuth example from the manual.
  Simple Mutual PSK XAuth Config
  
  Server side config:
/ip ipsec peer
add address=2.2.2.1 auth-method=pre-shared-key-xauth secret="123" passive=yes
/ip ipsec user
add name=test password=345
  
  Client side config:
/ip ipsec peer
add address=2.2.2.2 auth-method=pre-shared-key-xauth secret="123" \
  xauth-login=test xauth-password=345
  
  Note: On server side it is mandatory to set passive to yes when XAuth
  is used.
