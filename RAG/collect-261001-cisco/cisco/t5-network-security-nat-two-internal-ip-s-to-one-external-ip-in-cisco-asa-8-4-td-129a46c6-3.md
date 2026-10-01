---
id: collect-261001-cisco/cisco/t5-network-security-nat-two-internal-ip-s-to-one-external-ip-in-cisco-asa-8-4-td-129a46c6-3
title: "t5-network-security-nat-two-internal-ip-s-to-one-external-ip-in-cisco-asa-8-4-td-129a46c6"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/t5-network-security-nat-two-internal-ip-s-to-one-external-ip-in-cisco-asa-8-4-td-129a46c6.md
source_anchor: ""
source_lines: [212, 304]
sha256: 4d43b8fb8b02be37b6566e2c1266e43640343de862ffd1025b90805821b09de9
---

# t5-network-security-nat-two-internal-ip-s-to-one-external-ip-in-cisco-asa-8-4-td-129a46c6

network-object host 192.168.0.3
network-object host 192.168.0.6
object network NAT-IP
 host 
nat (inside,outside) 1 after-auto source dynamic SOURCE-ADDRESSES NAT-IP
object network SMTP-SERVER
host 192.168.0.3
nat (inside,outside) static NAT-IP service tcp 25 25
NAT Configuration When Using ASA "outside" Interface Public IP Address
object-group network SOURCE-ADDRESSES
network-object host 192.168.0.3
network-object host 192.168.0.6
nat (inside,outside) 1 after-auto source dynamic SOURCE-ADDRESSES interface
object network SMTP-SERVER
host 192.168.0.3
nat (inside,outside) static interface service tcp 25 25
Naturally you will have to make sure that you open the TCP/25 port on the ACLs on the ASA
Also possible existing configurations can affect if this configuration works or not. But it can be confirmed either with testing traffic OR using the "packet-tracer" command on the ASA
For example to test the incoming SMTP traffic
packet-tracer input outside tcp 1.2.3.4 12345 
To test the outgoing traffic from the hosts
packet-tracer input inside tcp 192.168.0.3 12345 1.2.3.4 
packet-tracer input inside tcp 192.168.0.6 12345 1.2.3.4 
Hopefully this helps
Remember to mark the question as answered if it was. Or ask more if needed.
- Jouni
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-09-2013 09:19 AM
Thank you for the reply and my apologies for not elaborating more. Yes that is what I was trying to do. I did what you suggested but there is an issue with that configuration because we have a PAT setup so if I do after auto then the NAT rule comes after that PAT for all the outbound internet traffic and it does not work.
If I remove after auto and just use 1 then it puts it all the way on the top and then the incoming mail does not reach us. So here is what I did:
object-group network Email_InOut
network-object host 192.168.0.3
network-object host 192.168.0.6
exit
object network obj-1.1.1.1
host 1.1.1.1
exit
nat (inside,outside) after-auto 1 source dynamic Email_InOut obj-1.1.1.1 (did not use this)
object network Inbound_Email
host 192.168.0.3
nat (inside,outside) static 1.1.1.1 service tcp 25 25
exit
object network IronPort
host 192.168.0.6
 nat (inside,outside) static 1.1.1.1
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-09-2013 09:43 AM
Hi,
Yeah, since we dont see the full NAT configuration we wont know how the current existing configuration affects what we are trying to achieve.
The reason why I personally suggest configuring Network Object NAT for the Static PAT / Port Forward AND Twice NAT/Manual NAT type of configuration for the Dynamic PAT is how I personally order the NAT rules in my configurations
- Static PAT and Static NAT always as Network Object NAT (Section 2)
- Default Dynamic PAT/NAT always as Twice NAT / Manual NAT (Usually Section 3 in special cases Section 1)
- Special NAT setups like NAT0 and Policy NAT/PAT type configurations as Twice NAT / Manual NAT (Section 1)
Naturally Static PAT and Static NAT can be done in the Section 1 also but I prefer keeping strict roles for every Section and so far it has worked for me.
I wrote a NAT 8.3+ Document which pretty much states the way I configure and section the different type of configurations. Have a look if you want. Will probably add a lot more information to it later
https://supportforums.cisco.com/docs/DOC-31116
To me it seems the NAT configuration ordering/sectioning is causing the problems why the suggest configurations dont work. The existing configurations is set up so that it overrides the configurations suggested.
Glad to hear you got it working though
- Jouni
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-09-2013 10:00 AM
Yes you are right the existing NAT configuration was causing the issue so I had to modify it a bit but it seems to be working. Thank you for the document link, good read. I have saved it
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-09-2013 09:42 AM
Hello Mohammad,
Exactly, the configuration I sent will do it
Regards
Senior Network Security and Core Specialist
CCIE #42930, 2xCCNP, JNCIP-SEC
