---
id: collect-261001-fortinet/fortinet/questions-20890-fortigate-30d-ipsec-vpn-could-not-locate-phase1-configuration-111bf081
title: "questions-20890-fortigate-30d-ipsec-vpn-could-not-locate-phase1-configuration-111bf081"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: ["2015-08-24", "2015-09-08"]
keywords: []
source: docs/RAG/collect-261001-fortinet/questions-20890-fortigate-30d-ipsec-vpn-could-not-locate-phase1-configuration-111bf081.md
source_anchor: ""
source_lines: [1, 28]
sha256: 170ef484fc4aab4493aaccabeb14b4ff65c6fa1855e22b884ac999d4ebcb509e
---

# questions-20890-fortigate-30d-ipsec-vpn-could-not-locate-phase1-configuration-111bf081

Hello network engineers,
I have an IPSEC VPN tunnel between two offices, the HQ is a fortigate 200B(os:v5.0,build0292 (GA Patch 9)) and the branch is fortigate 30D(os:5.2.3).
Both are now on static IPs. The branch side has an PPPoE though. auto-reconnect is also enabled on the branch side.
The issue we're facing since day 1 is the tunnel works fine for the day but the next morning is down and does not come up on its own at all, until some minor change is made to the phase1 configuration on the branch side.
I have a case locked with the fortigate TAC since then, it's over 2 months of submitting logs to them and no solution has been provided at all. Submitting logs to them is now a daily thing.
Running debugging during the time of the issue on the branch 30D the initial out put is 2015-08-24 21:44:34 ike 0:mandhana: could not locate phase1 configuration. 
I could find only this one similar case on their forums, my branch side is already on 5.2.3 so the solution of updating to 5.2 which worked there won't work for me.
Really appreciate any help towards what could be causing this.
UPDATE 4th NOV 2015:
The wan interface has an static IP address but its through PPPoE, so i guess there's some sort of ip address renewal still happening even though the ip address which the wan interface gets is still the same and is static.
When the ip is renewed the tunnel goes down and doesn't come up on its own for some reason.
The TAC has now declared this as a bug and tells me to wait till FortiOS 5.4.0 is released.
I am of the opinion though that this is something to do with hardware itself. If this was a bug we would have more of such cases here and elsewhere but i couldn't find any.
Here's the fortinet TAC's logging info:
Hi,
` -Here is what the issue started occurring today morning. 2015-09-08 18:29:41 is PST time(your device has PST time zone), not IST time zone. So actually issue started exactly at 7:30AM today IST time. ---/--- 2015-09-08 18:29:41 ike 0: IP 59.x.x.x (59) is down ----<<<< probably the pppoe interface went down as the 'IP is down' in the debug logs. 2015-09-08 18:29:41 ike 0:mandhana: deleting 2015-09-08 18:29:41 ike 0:mandhana: flushing 2015-09-08 18:29:41 ike 0:mandhana:79: send IPsec SA delete, spi c21221f8 ---<<<<
-The below debug logs confirms that default route was no longer available in FGT as per 'len=92: error 101:Network is unreachable'
-So it's highly possible that default route was removed from FGT routing table.
---/---
2015-09-08 18:29:41 ike 0:mandhana:79: could not send IKE Packet(ISAKMP SA DELETE-NOTIFY):59.x.y.z:500->14.x.z.c:500, len=92: error 101:Network is unreachable
---/---
Above continues till next 1 min.
-Next min, we see the following error
2015-09-08 18:30:16 ike 0:mandhana: could not locate phase1 configuration.
` I suspect, there is an issue with pppoe related -(routes, ip address, policy) when this issue occurs. I will login to fortigate tomorrow at 9:00 am, and figure it out that. Please don't make any changes.
The other option(i think its best) I advised you earlier is that configure the modem in L3 mode, do the nat port forwarding to Fortigate using the public IP. doing this way, avoids the pppoe config in fortigate.
You have port forward (https, http, VPN traffic) from modem to FGT.
Thanks.
