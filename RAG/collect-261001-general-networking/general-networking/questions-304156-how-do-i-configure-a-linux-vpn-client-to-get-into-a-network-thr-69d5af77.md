---
id: collect-261001-general-networking/general-networking/questions-304156-how-do-i-configure-a-linux-vpn-client-to-get-into-a-network-thr-69d5af77
title: "questions-304156-how-do-i-configure-a-linux-vpn-client-to-get-into-a-network-thr-69d5af77"
domain: general-networking
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-general-networking/questions-304156-how-do-i-configure-a-linux-vpn-client-to-get-into-a-network-thr-69d5af77.md
source_anchor: ""
source_lines: [1, 14]
sha256: 851bb2302fa34e777427d355f200f36e6f6a899e94a7bdf9e2ec909f6410e5d2
---

# questions-304156-how-do-i-configure-a-linux-vpn-client-to-get-into-a-network-thr-69d5af77

Server Fault is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
2
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
In order to connect to my job's VPN, I have been given by the network admin:
a username
a password
a PSK
I run Ubuntu at home. I know Fortigate's VPN should be a vanilla IPSec, so OpenSwan should do the trick. Still, I can't get it to work.
I have tried a program called "Forticlient" for Linux that I found through Google but it doesn't have the appropriate fields for the 3 items listed above.
I suggest you check out Openfortivpn. I had to resort to that, as our implementation of fortigate VPN doesn't have a functioning linux client. OpenfortiVPN works great for me:
The trusted cert is the certificate offered by the VPN gateway, and will be displayed if you try to connect. Then you edit the config to add the certificate. Several certificates can be specifie on separate lines.
