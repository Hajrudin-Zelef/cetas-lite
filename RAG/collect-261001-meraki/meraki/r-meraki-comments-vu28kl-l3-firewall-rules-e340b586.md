---
id: collect-261001-meraki/meraki/r-meraki-comments-vu28kl-l3-firewall-rules-e340b586
title: "r-meraki-comments-vu28kl-l3-firewall-rules-e340b586"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/r-meraki-comments-vu28kl-l3-firewall-rules-e340b586.md
source_anchor: ""
source_lines: [1, 30]
sha256: c0b7c5323e2e7a3e6ff9c73d8bfbf6cd0144e2c87ed04e4c7cabb660ca89439b
---

# r-meraki-comments-vu28kl-l3-firewall-rules-e340b586

L3 Firewall rules 
        
    Hey folks!
      I have started working with a Meraki MX100 and I need to configure L3 firewall rules...
In the past I have worked with Fortinet, Pfsence, Ubiquity, Sonicwall etc etc
And with the Meraki firewall it seems that even the most basic configuration is sometimes very hard to achieve. For exemple, I simply want to isolate a DMZ vlan but still allow some trafic to pass to the other vlan (Lan-General):
    
      vLan A is Lan-General (192.168.1.1)
vLan B is DMZ (192.168.2.1)
    
      First I wan to block the DMZ from accessing the Lan-General:
Rule 1: Deny, proto=all, from=192.168.2.0/24, to=192.168.1.0/24, ports=all
    
      Then, I wan to allow a server in the DMZ to communicate with another server on the Lan-General (lets say a syslog server):
Rule 2: Allow, proto=udp, from=192.168.2.200, to=192.168.1.100, ports=514
    
Well somebody help me understand the logic becuse as soon as the first rule is effective, the second rule has no effect, no mather the order of the rules. To circumvent that I have to create mutiple Deny rules that do not use port 514.. so one rule to deny from port 1-513/TCP, another to deny 515-99999/TCP, another to deny ALL UDP. another to deny all ICMP...
      Really?
Well thanks for your feedback folks!
    
Section des commentaires
Here's some clues: Problems arise when using multiple CIDR, Objects or Groups (of objects).
This is not making any sense. I will now contact Meraki for support. Thanks for your help.
Firewall rules work from a top to down order. So your 2nd rule would be the most specific rule so it should be at the top and your 1st rule should be under it.
Thanks all for your help but as mentionned, I did reverse both rules and I wasen't getting the expected results ... Coult it be that I am not being patient enough for the config to refresh in the unit itself? I usually wait at least 2-3 minutes after each config before attempting to verify...
Anyway thanks again for your input, I'll keep researching and testing.
if you’ve already tried this then something odd is going on. Do you have a group policy configured somewhere between those networks?
Switch the order and do the second rule first. You can't put a big deny statement first and then expect a more specific allow statement after it to be hit. Pretty much every firewall policy works top down until it matches on a rule. Then it stops.
Check out my comments on this thread. Good way of starting with a dent all!
link
