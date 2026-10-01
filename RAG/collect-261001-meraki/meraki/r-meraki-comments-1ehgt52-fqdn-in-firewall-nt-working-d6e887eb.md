---
id: collect-261001-meraki/meraki/r-meraki-comments-1ehgt52-fqdn-in-firewall-nt-working-d6e887eb
title: "r-meraki-comments-1ehgt52-fqdn-in-firewall-nt-working-d6e887eb"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/r-meraki-comments-1ehgt52-fqdn-in-firewall-nt-working-d6e887eb.md
source_anchor: ""
source_lines: [1, 32]
sha256: 8586dff48af89e71bd128021ed2334eeb9c41c42f47ad9b8da1dec8ac07df119
---

# r-meraki-comments-1ehgt52-fqdn-in-firewall-nt-working-d6e887eb

FQDN in firewall nt working
Hi everyone,
      im trying to block some specifc site on my mx from my iot-wifi.
My client gets his ip and uses the meraki as gateway and also as dns.
on the firewall-rules i blocked heise.de but i can ping and visit the site everytime i try.
In my Understanding meraki should snoop the dns replies and block the ip. But it does not work.
When i use specific ip-address rules everything gets blocked to this ip.
    
Is there something wrong in my concept?
Section des commentaires
Where are your DNS servers? (Does your DNS request go through the firewall?) If it doesn't then I don't believe the DNS snooping can work.
Client has 172.26.8.5. dns and gateway is the 172.26.8.1(mx). Im using a galaxy s23 ultra without private dns. I took captures but do not see any dns requests to the mx. Even when i use google dns(8.8.8.8) its not going to the mx in the capture.
I do not know if android is using something else than udp 53. But i blocked tcp 53 and 853
That's strange. I would setup a packet capture to capture all traffic from that IP and see what its doing.
Are you maybe using dns over tls/https?
I dont think so. I ran multiple packet captures and can not see my phones dns requests to the configured dns servers in the trace. I think my phone is doing something fishy and not using the dns servers correctly... Need to test it with a laptop.
Probably doing dns over https… you Can block is using content filtering and then you should be able to do dns filtering…
Don’t use the Layer 3 firewall rules for that. Use the Layer 7 to block to the host domain further down to the bottom of the same page
So layer 7 http hostname and then the domain? But shouldnt it work with layer 3 also?
Just trying to think of a workaround, but it should work for layer 3 as well
What firmware version are you on? I use a lot of FQDNs on my MXs. After upgrading to 18.211.2 I started having a ton of random issues. Worked with support and ran tons of packet captures, and eventually was advised to roll back to a prior firmware version. From what I understand it's a known issue, but I still haven't seen it in the patch notes.
L3 fqdn blocks don't work if the client already has DNS for that name. The mx uses DNS snooping to make those rules work. https://documentation.meraki.com/MX/Firewall_and_Traffic_Shaping/MX_Firewall_Settings#FQDN_Support
I'd follow the below steps to effectively block specific sites on your Meraki MX from your IoT-WiFi:
Layer 7 Blocking:
Instead of relying solely on Layer 3 firewall rules, you can also block the domain at the Layer 7 level. This will help inspect and block traffic at the application layer.
Content Filtering:
Utilize Meraki's content filtering feature under Security & SD-WAN > Content Filtering. This feature allows you to block access to specific websites more effectively.
Firewall Rule Order:
Ensure that no other firewall rules are overriding the block rule for heise.de. Sometimes, other rules might take precedence and affect the blocking.
Firmware Stability:
If you have recently updated your firmware, check if it is stable or still in beta state. Beta firmware might have issues that can impact functionality.
On a different note, if you have the capacity to leverage more, I'd recommend using Cisco Umbrella. I have seen it do a good job of blocking sites without much hassle.
