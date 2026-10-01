---
id: collect-261001-fortinet/fortinet/questions-311695-assign-ca-certificate-to-fortigate-https-wan-interface-f5c85cd6
title: "questions-311695-assign-ca-certificate-to-fortigate-https-wan-interface-f5c85cd6"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-fortinet/questions-311695-assign-ca-certificate-to-fortigate-https-wan-interface-f5c85cd6.md
source_anchor: ""
source_lines: [1, 21]
sha256: c936b19e3f6fa70a190adbb3ccb7b3db0b2852bc2a0e754ffb440cf5643acac1
---

# questions-311695-assign-ca-certificate-to-fortigate-https-wan-interface-f5c85cd6

Server Fault is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
3
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I have a Fortigate 80C that allows remote administration via https.
I access the URL and all fine but the one thing that really bugs me is that is brings up "untrusted connection" in chrome with the whole "click to proceed" thing.
At the moment the cert is self-signed by the Fortigate unit.
On the unit there are 5 CA signed certs for use but I cannot figure out how to assign these certs to the routers interfaces.
Does anyone know how to assign the CA signed certs to the WAN interface on port 443 so it wont ask me to confirm the cert all the time?
(I know the traffic is still encrypted but it is still nice to have)
Upload your certificates to the firewall, Fortigate certificate user guide will help you out on this. Now to use this certificate for HTTPS admin access. Use the following CLI commands:
config system global
set admin-server-cert <certname>
end
Okay this is what I did on 5.2 so that I could use a certificate signed by our internal CA:
System->Certificates->Local Certificates->Generate (this will generate the CR)
Check-click the newly created CR and Download it and process it on your favorite CA
System->Certificates->Local Certificates->Import (this will import the signed cert), set Type to 'Local Certificate if it isn't already. The FortiGate should now have the CA info filled in for what was the CR.
Follow instructions above for setting the certificate as the admin interface cert:
