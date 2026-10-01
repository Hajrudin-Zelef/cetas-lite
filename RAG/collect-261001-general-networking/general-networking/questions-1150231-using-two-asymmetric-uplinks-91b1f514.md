---
id: collect-261001-general-networking/general-networking/questions-1150231-using-two-asymmetric-uplinks-91b1f514
title: "Using two asymmetric uplinks"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/questions-1150231-using-two-asymmetric-uplinks-91b1f514.md
source_anchor: ""
source_lines: [1, 45]
sha256: 858b4c0d43e43d602c476d84836e2bf6a244d37c57ebed9475a07924a220115b
---

# Using two asymmetric uplinks

*Score : 0 | Source : https://serverfault.com/questions/1150231/using-two-asymmetric-uplinks*

My current network topology is:
Fact Firewall (I tested with OPNsense and IPFire, but I'm open to further suggestions) and the whole "ORANGE/DMZ" stuff are actually VMs/containers in a physical LXD server should not impact (but I might be wrong, of course).
My current setup (using IPFire ATM) is fully functional, BUT it uses only ISP_B
for communication.
I can easily switch to use ISP_A and it works fine just the same (with floating IP
limitation, of course).
My attempts to use both failed.
Target would be:
ISP_A for all outgoing traffic originated from GREEN/LAN and ORANGE/DMZ.webserver on ssh/git
protocol) to userver on ISP_B for all traffic
in case of My attempts failed miserably essentially because if I set the default gateway to
ISP_A packets coming from ISP_B get NATted, then correctly sent to Firewall and redirected to webserver (so far, so good), but replies are sent from Firewall to ISP_A which did never see the original packet, so it can't divine where to send it
and simply drops it.
I am missing something, but I couldn't understand what.
Note 1: both modem/routers are set in "DMZ mode" (i.e.: full port redirect) to Firewall (192.168.1.9).
Note 2: Analysis above was done under IPFire, but OPNsense behavior is similar.
Note 3: I would like to understand what should be done in such circumstances, not just "how to fix", if possible.
I am ready to provide any required info, including full configs, if deemed useful.
Here follows the log of everything I did from OPNsense installation to date.
The last part (Port redirection) does not seem to work properly, even if rule is triggered and logs (see below).
cinderella).https://192.168.7.9)root password.[Save] and [Apply changes][+]DMZ netdmzAllow DMZ to access Internetconfig-OPNsense.condarelli.it-01.xml10BackupDSL backup connection (with fixed IP)WAN192.168.1.2100config-OPNsense.condarelli.it-02.xmlBackup - 192.168.1.2TCP/UDPDMZ addressDNS to Allow access to DNS servertracerouteconfig-OPNsense.condarelli.it-03.xmlwebff14ffWEBportsPort(s)80 81 443Web standard portsHost(s)dmz web192.168.9.8Nginx web server and reverse proxyWAN addressweb dmzAllow external access to Nginx Proxy Manager
If I try to access my nginx server using a DNS entry pointing to ISP_B external address I can see in logs:
Note: 79.34.242.11 is the (floating) IP currently assigned to ISP_A modem.
Sorry if it was not clear but I have two DMZ rules:
The first allows access to DNS server on firewall and the second redirects everything else to ISP_B.
I am unsure why the first is needed, but without it DNS resolution does not work for hosts sitting on DMZ.
Hosts connected to DMZ really use ISP_B/Backup/192.168.1.2 to reach the internet; that bit works. What does not work is getting back the answer after the request packet was (apparently correctly) redirected to DMZ/webserver.
The port forwarding rule looks like:
and apparently works as logs above seem to indicate.
I am unable to understand why response packets seem to get lost (and how can I trace them to understand what went south).
Unfortunately I'm very far from being a net (and OPNsense) guru.
Specific instructions on how to overcome the problem (or further diagnose it) would be very welcome.

---

### Reponse — score 1

What you want is basically policy based routing.
In OpnSense (and pfsense) you do this in a firewall rule:
Here the gateway will be used for packets matching that rule. You can add additional gateways in addition to default under System - Gateways.
You will have to define firewall rules for the different traffic types, and assign them to correct gateways. The firewall rules can match on all the usual stuff - source IP, source network and so forth to send traffic according to your wishes.
