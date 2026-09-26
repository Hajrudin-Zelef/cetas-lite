---
id: collect-260926-mikrotik/mikrotik/questions-559820-mikrotik-ipsec-client-fortigate-received-esp-packet-with-unknow-22a3c23e
title: "questions-559820-mikrotik-ipsec-client-fortigate-received-esp-packet-with-unknow-22a3c23e"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["incident"]
source: docs/RAG/lot-mikrotik/forum/ipsec/questions-559820-mikrotik-ipsec-client-fortigate-received-esp-packet-with-unknow-22a3c23e.md
source_anchor: ""
source_lines: [1, 53]
sha256: a5e95716b85a54b79014017fafd5c6b9c28d31fd9601a1538311ca21928ebc41
---

# questions-559820-mikrotik-ipsec-client-fortigate-received-esp-packet-with-unknow-22a3c23e

We have a client with 6 sites using IPsec. Every now and again, possibly once a week, sometimes once a month, data just stops flowing from the remote Fortigate VPN server to the local MikroTik IPsec VPN client.
In order to demonstrate the symptoms of the problem I have attached a diagram. On the diagram Installed SAs tab you will notice a source IP address x.x.186.50 trying to communicate with x.x.7.3 but 0 current bytes. x.x.186.50 is the client's remote Fortigate IPsec server, and x.x.7.73 is a MikroTik based IPsec endpoint. It appears data from the remote side to us is not always flowing.
Phase 1 and 2 are always established but traffic always refuses to flow from the remote side to us.
We tried various things over time, such as rebooting, setting clocks, dabbling with configuration, rechecking and rechecking configuration but it appears the problem is entirely random. And sometimes random things fixes it. At one stage I had a theory that if the tunnel is initiated from their side it works, but fiddling with "Send Initial Contact" has not made any difference.
We've had many chats to the client about this but they have many more international IPsec VPNs and only our MikroTik configuration is failing.
Fortigate log:
http://kb.fortinet.com/kb/microsites/microsite.do?cmd=displayKC&externalId=11654
Looking at Fortigate's knowledgebase it appears SPIs don't agree and DPD would make a difference. But I have tried every single combination of DPD on this side without avail. I would like to enable DPD on the other side but I cannot due to change control and also because the client is saying it's working on all the other sites exactly configuration the same. EDIT DPD was enabled
Local VPN client diagram showing no traffic flow:
I have included a log file showing continuous loops of "received a valid R-U-THERE, ACK sent" MikroTik log file:
echo: ipsec,debug,packet 84 bytes from x.x.7.183[500] to x.x.186.50[500]
echo: ipsec,debug,packet sockname x.x.7.183[500]
echo: ipsec,debug,packet send packet from x.x.7.183[500]
echo: ipsec,debug,packet send packet to x.x.186.50[500]
echo: ipsec,debug,packet src4 x.x.7.183[500]
echo: ipsec,debug,packet dst4 x.x.186.50[500]
echo: ipsec,debug,packet 1 times of 84 bytes message will be sent to x.x.186.50[500]
echo: ipsec,debug,packet 62dcfc38 78ca950b 119e7a34 83711b25 08100501 bc29fe11 00000054 fa115faf
echo: ipsec,debug,packet cd5023fe f8e261f5 ef8c0231 038144a1 b859c80b 456c8e1a 075f6be3 53ec3979
echo: ipsec,debug,packet 6526e5a0 7bdb1c58 e5714988 471da760 2e644cf8
echo: ipsec,debug,packet sendto Information notify.
echo: ipsec,debug,packet received a valid R-U-THERE, ACK sent
I've received various suggesions from IPsec experts and MikroTik themselves implying that the problem is at the remote side. However the situation is greatly compounded that 5 other sites are working and that the client's firewall is under change control. The setup also always worked for many years, so they claim it cannot be a configuration error on their side. This suggestion seems plausible but I cannot implement due to change control. I may only change the client side:
Make sure the IPSec responder has both passive=yes and send-initial-contact=no set.
This did not work.
EDIT 9 Dec 2013
I am pasting additional screenshots with the Fortigate configuration and what we believe are the Quick Mode selectors on the Mikrotik side.
Let me re-iterate that I don't think it's a configuration problem. I speculate it's a timing problem whereby side A or side B tries to send information too aggressively making the negotiation of the information (e.g. the SPI) out of sync.
EDIT 11 Dec 2013
Sadly I have to give up on this issue. Happily everything is working. Why it's working is still a mystery, but to further illustrate what we did I post another image inline.
We fixed it by:
- Turning off PPPoE at client.
- Installing completely new router (Router B) and tested at Border. It worked at Border.
- Switching off new router B at border. AND THEN, WITHOUT MAKING A SINGLE CHANGE, the client's end-point Router A started working. So just adding a duplicate router at the border and taking this router offline again made the original router work.
So add this fix to the list of things we've done:
- Reboot. That worked once.
- Create new tunnel with new IP. That worked once but only once. After changing IP back client endpoint came live again.
- Change time servers.
- Fiddle with every possible setting.
- Wait. Once, after a day, it just came right. This time, even after days, nothing came right.
So I postulate that there is an incompatibility on either Fortigate or MikroTik side which only happens at very random situations. The only things we haven't been able to try is upgrade firmware on Fortigate. Maybe there is hidden corrupt configuration value or timing issue invisible to configurer.
I further speculate that the issue is caused by timing issues causing SPI mismatch. And my guess is the Fortigate doesn't want to "forget" about the old SPI, as if DPD is not working. It just happens randomly and from what I can tell only when endpoint A is Fortigate and endpoint B is MikroTik. The constant aggressive attempts at trying to re-establish the connection "holds" on to old SPI values.
I'll add to this post when it happens again.
EDIT 12 Dec 2013
As expected it happened again. As you may recall we have 6 MikroTik client IPsec end-point routers configured exactly the same connecting to one Fortigate server. The latest incident was again to a random router, not the one I posted here about originally. Considering the last fix where we installed this duplicate router, I took this shortcut:
- Disable Router A, the router that does not want to receive packets from Fortigate any more.
- Copy Router A's IPsec configuration to a temporary router closer to the border of our network.
- Immediately disable the newly created configuration.
- Re-enable Router A.
- Automagically it just starts working.
Looking at @mbrownnyc comment I believe that we are having an issue with Fortigate not forgetting stale SPIs even though DPD is on. I will investigate our client's firmware and post it.
Here is a new diagram, much like the last, but just showing my "fix":
autokey keep alive. Also, setup a ping from the remote site to a host at the destination site. What about yourquick mode selectors(and whatever MikroTek calls them)? What about your fortigate debug logs (diag debug app ike -1?R-U-THEREis a function ofDPD(which functions on phase 1, it seems like phase 1 is establishing (okay on the Aggressive versus main mode), but phase 2 might be failing. I'd say, what about PFS, but I already said verify each setting is exactly the same, particularly what Fortinet calls Quick Mode Selectors. It doesn't seem you have confirmed that you have verified every single setting. Can you post what they gave you (less IPs, shared key, etc), appending to your original post?the Fortigate doesn't want to "forget" about the old SPI,YES YES YES! I have had this happen to me. I'm on v4 MR3 patch 11.
