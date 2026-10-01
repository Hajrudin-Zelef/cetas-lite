---
id: collect-261001-meraki/meraki/r-meraki-comments-wxlwl9-meraki-firewall-sanity-check-and-questions-9d377a1c
title: "r-meraki-comments-wxlwl9-meraki-firewall-sanity-check-and-questions-9d377a1c"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution"]
source: docs/RAG/collect-261001-meraki/r-meraki-comments-wxlwl9-meraki-firewall-sanity-check-and-questions-9d377a1c.md
source_anchor: ""
source_lines: [1, 42]
sha256: 8950f5959eb81695893d5e8410d9eb14991ed9ee1fc9a64ca702c5e555165de5
---

# r-meraki-comments-wxlwl9-meraki-firewall-sanity-check-and-questions-9d377a1c

Meraki firewall sanity check and questions. 
        
        
        
    
    
    Inherited a site with 2 Meraki MX84s in HA. The only Meraki hardware onsite is the firewalls. Slightly confused it is not seeming to work like I expected. I appear to get inter-vlan traffic through the Meraki firewalls when I want to control most vlans on an internal switch. I do want 2 separate isolated vlans to go out the meraki to internet only and never touch the internal network or each other.
I need to Re IP the site over the next month(s) from a flat 192.x.x.x to a vlaned 10.x.x.x
I want to make 10.100.1.254 the new Meraki internal IP address for all 10.x.x.x vlans. I need to keep the current 192.168.0.10 IP address until everything has moved over.
10.100.1.1 will be the internal layer 3 switch for all internal 10.x.x.x vlans
--------------------------------------------
Current vlan 1 192.168.0.0/20 MX IP 192.168.1.10
New vlan 2 guest 172.16.1.0/24 MX IP 172.16.1.254
New vlan 3 lab 172.16.2.0/24 MX IP 172.16.2.254
New vlan 10 servers 10.100.1.0/24 MX IP 10.100.1.254
New static route xxx2 10.100.2.0/24 gateway ip 10.100.2.1
New static route xxx3 10.100.3.0/24 gateway IP 10.100.3.1
New static route etc....
Now do I use Firewall rules or Group Policy to deny inter-vlan traffic?
Deny/Deny 172.16.0.0/12 to 192.168.0.0/20
Deny/Deny 172.16.0.0/12 to 10.100.0.0/20
Deny/Deny 192.168.0.0/20 to 172.16.0.0/12
Deny/Deny 10.100.0.0/20 to 172.16.0.0/12
Deny/Deny 172.16.1.0/24 to 172.16.2.0/24
Deny/Deny 172.16.2.0 to 172.16.1.0/24
Should I also deny both ways between 192.168.0.0/20 to 10.100.0.0/20 and use internal switch routing?
Will it matter that the 192.168.x.x vlans will be able to reach the Meraki at 192.168.0.10 and 10.100.1.254 ?
Creating the above rules I can still ping from 192.168.1.20 to 172.16.1.254, why?
Section des commentaires
Mixing layer 3 routing in a switch, and routing with an MX is going to get messy with Meraki. Unless you have lots of inter vlan traffic I would recommend leaving the layer 3 to the MX.
Also worth mentioning that I’m fairly certain you can ping any MX interface IP regardless of firewall rules, you just won’t reach anything else in that subnet.
Now you can create groups of objects on meraki, I create a group with 10.0.0.0/8,172.16.0.0/12 and 192.168.0.0/16. Then make a deny all from that group to that group.
This way all inter vlan traffic is blocked except what you explicitly allow. Useful with the absence of a deny all rule.
Ah, so that is what is happening. I can ping all the Meraki IP addresses but cannot access the subnets.
I don't see a place to make groups. That would be quite helpful.
There is a ton of layer 3 routing that will need to be done at the site. About 300 users with a couple of racks of IT equipment and some manufacturing and office space. The switch will be non blocking for internal 10/25gb from say the user or production vlan to server vlan. The only place I want Meraki to route traffic is to the Internet from the primary server vlan for the majority or traffic or from the isolated guest and lab vlans. I have no clue how they have been working on a single large flat vlan for so long.
I am much more a server guy than networking but wearing all the hats in my new role and had never run into a piece of equipment that seems to allow inter-vlan traffic by default, but again my networking experience has been limited.
Thank you.
Update the firmware to the latest, and you should see the ability to create object based groups. You can either do this on the fly when looking at firewall rules, or go to organisation > policy objects.
I just recently finished reconfiguring a network that has close to 300 clients on 2 /19 networks also with all Meraki Gear. I created a switch Stack for Collapsed core/dist and I am doing all L3 there the MX is only doing main internet firewalling/VPN work. and now have 10 use specific vlans. This is the right way to go, otherwise its working as a pseudo router on a stick, yes I know the MX technically has switch interfaces but you should not be using the MX as a top of rack or distribution switch because it does not participate in STP, CDP, or any other basic switch layer work.
Whether you use group policy to do so or not depends on your use case. I don't see a reason you need to use it though. It's overkill for simple firewalling.
From your description, the MX isn't doing routing and your switch is. If that is the case, idk what you expect a rule on the MX to do. If that's not the case, you should call Meraki and have them look with you.
