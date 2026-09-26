---
id: collect-260926-mikrotik/mikrotik/questions-884962-how-to-disable-tcp-timestamps-on-a-mikrotik-router-e1961a01
title: "questions-884962-how-to-disable-tcp-timestamps-on-a-mikrotik-router-e1961a01"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/questions-884962-how-to-disable-tcp-timestamps-on-a-mikrotik-router-e1961a01.md
source_anchor: ""
source_lines: [1, 21]
sha256: 03cfe66e6eff640fbee7e49b964099a69ac300ae31cc5e7c8319cf20c1dbe2e0
---

# questions-884962-how-to-disable-tcp-timestamps-on-a-mikrotik-router-e1961a01

I'm using OpenVAS as vulnerability scanner and I completed a local network scanning which also involves a MikroTik router as default gateway. 
 
The router shows a weakness, better known as "TCP timestamp" vulnerability. 
In light of the above, how to disable TCP timestamp on a MikroTik router?
2 Answers 2
AFAIK you cannot disable this on MikroTik.
I'd consider this a low risk vulnerability since you can only infer the uptime of the device with it.
If I am not mistaken, by using a proper firewall on the WAN side of your network this shouldn't be an issue anyway. If the router does not reply to any TCP packets from unknown hosts then they won't be able to get any timestamps.
If you really need to disable this, you should then open a ticket at support@mikrotik.com to request this as an option for MikroTik RouterOS. They are the only ones that can implement this as RouterOS is closed source software (albeit based on the linux kernel).
I didn't actually used nor have a mikrotik router but as far as I know the OS is linux based. So I guess you could disable TCP Timestamps the same way you do it on a linux server:
Login via ssh to the mikrotik router and add the following line to the /etc/sysctl.conf file
#Disable TCP timestamps 
net.ipv4.tcp_timestamps = 0
#Enable TCP Timestamps
net.ipv4.tcp_timestamps = 1
Then run a sysctl -p to enable the settings.
Or when the system is running...
echo 0 > /proc/sys/net/ipv4/tcp_timestamps
echo 1 > /proc/sys/net/ipv4/tcp_timestamps
- 
        MikroTik has its own CLI/GUI. You don't have access to anything that MikroTik itself does not allow. That includes the filesystem.Cha0s– Cha0s2017-11-24 12:24:59 +00:00Commented Nov 24, 2017 at 12:24
