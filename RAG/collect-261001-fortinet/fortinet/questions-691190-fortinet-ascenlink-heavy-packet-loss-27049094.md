---
id: collect-261001-fortinet/fortinet/questions-691190-fortinet-ascenlink-heavy-packet-loss-27049094
title: "questions-691190-fortinet-ascenlink-heavy-packet-loss-27049094"
domain: fortinet
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/questions-691190-fortinet-ascenlink-heavy-packet-loss-27049094.md
source_anchor: ""
source_lines: [1, 12]
sha256: 1e854826bf1bf175fc8cb0b28358050775a2c89e3bdc4c4d0e8e184bf76239ab
---

# questions-691190-fortinet-ascenlink-heavy-packet-loss-27049094

I'm working on a company with some Fortinet appliances. One of them is a Load Balancer named AscenLink. The specific model is an AscenLink AL-700.
There's a insane problem: a huge packet loss on the device, here's an excerpt or the mtr Unix software output:
Packets               Pings
 Host                                                                      Loss%   Snt   Last   Avg  Best  Wrst StDev
 1. 172.16.214.1                                                            0.0%   106    0.7   1.0   0.7   2.0   0.0
 2. 189.3.123.133                                                          52.4%   106    1.3   1.3   1.0   2.3   0.0
 3. 189.3.123.129                                                           0.9%   106    1.6   1.8   1.5   2.7   0.0
 4. embratel-s2-0-0-1-2-2-2-0-gacc01.vta.embratel.net.br                   25.7%   106  1045. 933.2 576.4 1259. 164.4
The architecture is simple. The first hop is the firewall/NAT on a Fortigate device, the second one with the 52.4% of packet loss is the Ascenlink Load Balancer and the next hop is the Cisco 2900 router from the ISP.
I've double checked and certified all the cables connecting the devices and they are good. Checked the diagnostics and statistics from the Ascenlink device and it claims to be good, looked for troubleshooting options on the manual and nothing was really conclusive. So I'm running out of options and I can't even use Google, since appears that no one uses this product.
I was wondering if someone here on ServerFault uses this Load Balancer from Fortinet and if there's some similar cases on the community.
Thanks in advance,
