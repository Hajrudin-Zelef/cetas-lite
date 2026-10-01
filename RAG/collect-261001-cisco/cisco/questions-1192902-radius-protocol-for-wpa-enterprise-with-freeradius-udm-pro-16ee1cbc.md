---
id: collect-261001-cisco/cisco/questions-1192902-radius-protocol-for-wpa-enterprise-with-freeradius-udm-pro-16ee1cbc
title: "questions-1192902-radius-protocol-for-wpa-enterprise-with-freeradius-udm-pro-16ee1cbc"
domain: cisco
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["compute"]
source: docs/RAG/collect-261001-cisco/questions-1192902-radius-protocol-for-wpa-enterprise-with-freeradius-udm-pro-16ee1cbc.md
source_anchor: ""
source_lines: [1, 29]
sha256: 793ba4f71897b6a87256727c1629fcf25c1169771d2e13a27cc20833e3369bc4
---

# questions-1192902-radius-protocol-for-wpa-enterprise-with-freeradius-udm-pro-16ee1cbc

I'm setting up WPA-Enterprise using a UniFi Dream Machine Pro and FreeRADIUS. My goal is to have the server receive cleartext user credentials (username/password). These credentials then get compared to argon2 hashes in my DB (that's why I need cleartext passwords).
I've tried EAP-TTLS with PAP inside, but my iOS devices fail to connect. The FreeRADIUS logs show:
(6) server inner-tunnel {
(6)   Restoring &session-state
(6)     &session-state:Framed-MTU = 1002
(6)   # Executing section authorize from file /etc/freeradius/3.0/sites-enabled/inner-tunnel
(6)     authorize {
(6) files: users: Matched entry bob at line 2
(6)       [files] = ok
(6) pap: No User-Password attribute in the request.  Cannot do PAP
(6)       [pap] = noop
(6) eap: Peer sent EAP Response (code 2) ID 1 length 6
(6) eap: No EAP Start, assuming it's an on-going EAP conversation
(6)       [eap] = updated
(6)     } # authorize = updated
(6)   Found Auth-Type = eap
(6)   # Executing group from file /etc/freeradius/3.0/sites-enabled/inner-tunnel
(6)     authenticate {
(6) eap: Removing EAP session with state 0x086488d308659df9
(6) eap: Previous EAP request found for state 0x086488d308659df9, released from the list
(6) eap: Peer sent packet with method EAP NAK (3)
(6) eap: Peer NAK'd asking for unsupported EAP type MSCHAPv2 (26), skipping...
(6) eap: ERROR: No mutually acceptable types found
It seems like iOS devices refuses to do PAP inside EAP-TTLS in my setup. In this test, I was just using the users file before verifying that it works and hooking it up to a real database.
My questions:
- What RADIUS / EAP methods are most widely supported on all devices and allow retrieval of cleartext credentials?
- Is there a recommended configuration for FreeRADIUS + UDM Pro for this use case?
- Is this approach even possible or should I just go with MSCHAPv2 and pre-compute nt hashes every time a user updates their password? I don't like how insecure those hashes are.
Any guidance or example configurations would be highly appreciated.
