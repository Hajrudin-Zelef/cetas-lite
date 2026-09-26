---
id: collect-260926-mikrotik/mikrotik/mikrotik-ipsec-tunnel-with-ddns-and-nat-occursus-arca-4
title: "mikrotik-ipsec-tunnel-with-ddns-and-nat-occursus-arca"
domain: mikrotik
role: reference
task: reference
actors: []
dates: ["2018-06-20", "2018-08-07", "2018-08-09", "2018-10-26", "2018-10-30", "2018-11-09", "2019-02-25", "2019-02-26", "2019-03-02", "2019-03-13", "2019-03-23", "2019-03-26", "2019-03-28", "2019-04-01", "2019-06-15", "2019-06-16", "2019-06-26", "2019-06-27", "2019-06-29", "2019-07-13"]
keywords: []
source: docs/RAG/lot-mikrotik/forum/ipsec/mikrotik-ipsec-tunnel-with-ddns-and-nat-occursus-arca.md
source_anchor: ""
source_lines: [365, 509]
sha256: 5c2d18e1a9bcb0fecb087c526d4e7a6ba82d144bdd3f741d6e80feeb07794145
---

# mikrotik-ipsec-tunnel-with-ddns-and-nat-occursus-arca

`/ip ipsec policy print`. Is private network behind Mikrotik set as source address and remote network behind Mikrotik as destination address? Is SA destination address set to remote public IP? This applies also vice-verse on the other Mikrotik. Can you see in the logs what policy is missing and verify the policy configuration accordingly? It might be also beneficial to check whether Mikrotik Router OS version is up-to-date on both sides.
## Mathew · June 20, 2018 at 23:46

Everything working. Thanks

## Charles · August 7, 2018 at 01:35

Great guide. Thanks.

Everything is working once the tunnel is created except that, whilst I can ping addresses apart from each router, I can’t ping one router from the other and vice versa.

Any ideas why that might be?

## Pessoft · August 9, 2018 at 21:16

Hi,

At first I suggest to check firewall configuration, whether it allows ICMP ping from ( and to ) the other router. Next thing is routing: in the section “Netwatch and Route” of this article is a route created, which makes the routers reachable between each other. Also Torch tool could help identify whether ICMP pings leave router via a correct interface.

## Lucky · October 26, 2018 at 18:11

Hi,

Can you explain exactly what part of the update script can be ommited if DNS names are used in peer configurations ?

## Pessoft · October 30, 2018 at 08:48

Hi,

ipsec-peer-update script updates 2 values: IP address of remote peer and SA destination address. If DNS names are used in remote peer configuration, parts of script related to update of remote peer can be omitted. So the command to create the script might look like this:

`/system script add name="ipsec-peer-update-vpn01" policy=read,write source=":local peerid \"vpn01\"\ \n:local peerhost \"0123456789.sn.mynetname.net\"\ \n:local peerip [:resolve \$peerhost]\ \n:local policyuid\ \n:set policyuid [/ip ipsec policy find comment=\"\$peerid\" and sa-dst-address!=\"\$peerip\"]\ \n:if (\$policyuid != \"\") do={\ \n /ip ipsec policy set \$policyuid sa-dst-address=\"\$peerip\"\ \n :log info \"Script ipsec-peer-update updated policy '\$peerid' with address '\$peerip'\"\ \n}"`
## Miguel · November 9, 2018 at 05:50

Thank you!!!! it’s a great guide. Everything working!

## Hannan · February 25, 2019 at 15:15

Nice guide. Really insightful and easy to understand. Thumbs up!!! Will try it with PureVPN…

## Mouhamed · February 26, 2019 at 08:51

Looks like RouterOS v6.43.12 makes the peer add command invalid.

It can no longer understand:

dh-group=modp4096 enc-algorithm=aes-256,aes-128 exchange-mode=ike2 hash-algorithm=sha512

Is there a way to get this to work considering the update?

Thanks!

## Pessoft · March 2, 2019 at 19:32

Hi Mouhamed,

Thanks for the info on peer command. I’ve reviewed the change log, tested the updated configuration on the latest stable RouterOS version 6.44 and updated the article. The arguments that are no longer understood by peer command in RouterOS 6.43.12 have been moved to the new profile object and are understood by the

`/ip ipsec profile`command. RouterOS 6.44 moved also arguments`auth-method`and`secret`to the new identity object and arguments are understood by the`/ip ipsec identity`command.
## Mike · March 13, 2019 at 10:56

Hi, will this Script also working if i use only one LAN port (ethernet2) to “inject” VPN between 2 Locations?

## Pessoft · March 13, 2019 at 21:57

Hi Mike,

Purpose of the script is to resolve the dynamic DNS hostname to an IP address and then update IPSec VPN configuration in a case new IP address is resolved. Script will work, if router is able to resolve the dynamic DNS hostname – probably in your case, as you mentioned, via the one LAN port.

## Kyaw Swa Hein · March 23, 2019 at 07:55

I got this message when ip cloud print “Router is behind a NAT. Remote connection might not work.”

Will this script work for router behind a NAT?

Thank in advance.

## Pessoft · March 26, 2019 at 20:40

Hi,

Yes, the point of the configuration is that both routers are behind NAT. IP Cloud sets the public IP of ISP’s router to a dynamic DNS entry – instead of the public IP of Mikrotik router, because Mikrotik is not connected directly to the public network. On remote side, script takes the IP of dynamic DNS entry and updates the IPSec configuration accordingly. Important point is that IPSec NAT-T traffic is forwarded by ISP router to the Mikrotik router.

## Kyaw Swa Hein · March 28, 2019 at 18:15

Thank you sir, I tried it and i still can’t ping to other client, but route show it reachable…..Can you please guide me where to fix by mail?

## Pessoft · April 1, 2019 at 20:19

Hi,

It might be caused by the firewall configuration. I sent you an email.

## Silvio · June 15, 2019 at 17:50

First thanks to the tutorial, works very well even through dynamic ip addresses.

The question is it possible to connect 3 (or more) mikrotik with VPN (all dynamic IP addresses). I was trying, but it will not work.

## Pessoft · June 16, 2019 at 23:21

Hi,

You can connect more Mikrotik’s together via VPNs. If you have 3 Mikrotiks (for example A, B, C) and you want each site to communicate with any other one, just add VPN between each site (i.e. A-B, B-C, A-C). You need to create every step of the guide again for a new VPN (also scripts). Make sure that you use separate configurations for each VPN, so they do not mix up. Alternative is also to have one Mikrotik configured so it routes between VPNs and then connect all other Mikrotiks there. This routing Mikrotik will then forward traffic between tunnels and will work as a central node. In this case you will just need each site to be connected to this central node (i.e. A-B, A-C).

## Silvio · June 26, 2019 at 00:31

Greeting,

the settings should be B-A-C, router A should work the tunnel according to B and C. Router A already has a tunnel set 10.10.20.0/24 – 10.10.10.0/24 – call it vpn1.

Question: Which ip addresses should be for another tunnel – vpn2? 10.10.30.0/24 – 10.10.40.0/24?

I got caught in Phase 2, I think that the Mikrotik communion is doing, assimilating SA Dst. address for each tunnel. Tunnel A connects almost immediately, but another tunnel B stops at Phase 2

## Silvio · June 27, 2019 at 16:04

Update,

all OK, firewall on ISP ADSL blocked everything possible. 🙁

If someone needs it: since I have more tunnels, on central MT, I used similar addresses 10.10.x.x. Therefore i used firewalls rules and nat 10.10.0.0/16. Each policies must have a special proposal, and peer a special profile.

Thanks.

## Pessoft · June 29, 2019 at 00:23

Hi Silvio,

Thank you for sharing your experience and I’m glad it works for you.

## Manuel · July 13, 2019 at 17:11

Hi Pessoft, i used your example and evrything works great, instead of use the 127.99.99.99 dummy address y use the actual public ip. Problem was when the ip changes. it won’t connect anymore.

How can i check if there is a part of the script not working or something

Thanks for your work!

## Pessoft · July 13, 2019 at 21:06

Hi Manuel,

127.99.99.99 is a dummy address, which is replaced automatically by the script with IP of remote Mikrotik’s dynamic DNS hostname ( IP cloud ). Watchdog enables schedule for the script whenever remote Mikrotik router is not reachable via VPN. You can check if script works, by setting again the dummy IP and then running the script manually. If IP is not replaced with remote Mikrotik’s IP by the script, then something with IP cloud configuration is not working ( you can check it further by trying to ping the remotehost name to see if it resolves to correct IP ) or incorrect hostname is set in peerhost variable of the script. You can also check if scheduler is working in the logs, where there should be the log message “changed scheduled script settings” whenever script schedule has been enabled and disabled ( when remote Mikrotik was not reachable and reachable again ).

