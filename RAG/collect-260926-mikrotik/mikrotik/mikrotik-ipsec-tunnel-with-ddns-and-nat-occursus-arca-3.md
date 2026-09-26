---
id: collect-260926-mikrotik/mikrotik/mikrotik-ipsec-tunnel-with-ddns-and-nat-occursus-arca-3
title: "mikrotik-ipsec-tunnel-with-ddns-and-nat-occursus-arca"
domain: mikrotik
role: reference
task: reference
actors: []
dates: ["2017-02-28", "2017-03-01", "2017-05-27", "2017-07-25", "2017-10-17", "2017-11-13", "2017-12-01", "2017-12-04", "2017-12-08", "2017-12-12", "2017-12-15", "2018-06-19", "2018-06-20", "2019-04-07"]
keywords: []
source: docs/RAG/lot-mikrotik/forum/ipsec/mikrotik-ipsec-tunnel-with-ddns-and-nat-occursus-arca.md
source_anchor: ""
source_lines: [236, 364]
sha256: e182afb57e5c1a632b696dec963c56a8ec24d39daddf9a873f8b2dcf283a85ec
---

# mikrotik-ipsec-tunnel-with-ddns-and-nat-occursus-arca

Regarding your second question: IPSec tunnel in config is configured to connect to remote site in case there is traffic going from local private network to remote private network ( in example above: if traffic will go from host in 10.10.10.0/24 network to network host in 10.10.20.0/24 network, then tunnel will be established from Mikrotik router 1 to Mikrotik router 2 and traffic sent through this tunnel – this works also vice versa ). So if you will have additional Mikrotik box in one of the sites with the same config as the other Mikrotik on the same site, it should work let’s say as its backup ( i.e. if there will be traffic going through additional box from local to remote private network, tunnel will be established to remote site ).

## DQ · February 28, 2017 at 10:34

Thanks Pressoft! I was not luck, your post was not working well on my side. I have almost same environment as shown in your topology. After I configured bother MT routers the IPSec tunnel was not up. in the logging file I could find that the phase one packet was sent out in both sides the routers could not receive response packets. What could be the potential reason?

I also have some questions which confused to me. Hope to get your answers.

1. What is the purpose of IP 127.99.99.99/32? is it a real IP or just an sample IP?

2. Is the IP an IP of a public stun server?

3. By using script and DDNS the MT router can know the public IP of peer, but how does the up lever NAT router know the received IPSec negotiation packets need to be forwarded to its direct connected MT router if there is no NAT translation table established before? vice verse.

Thanks a lot!

## Pessoft · March 1, 2017 at 00:16

Hi DQ,

At first make sure that 500/UDP and 4500/UDP traffic is being forwarded from gateways of your MTs to MT routers.

1. 127.99.99.99/32 is just a temporary placeholder IP, it will get replaced by IP of remote peer by the script

2. There is no STUN server used in the configuration, each MT knows about the IP of remote peer from the DDNS name of remote host.

3. There actually needs to be forwarding configured before ( in the article it is in the top – situation description point 4: Each MikroTik router has IPSec protocol, NAT-Traversal (4500/UDP) and IPSec IKE (500/UDP) traffic forwarded from its gateway (ISP Router) ). This is usually done on routers using configuration called port forwarding or DMZ host.

## Anton · May 27, 2017 at 22:49

This was the most useful guilde I have ever encountered! Works like a charm and it updates FAST!

## Anton · May 27, 2017 at 22:49

guide*

## Eduardo · July 25, 2017 at 15:36

Hello!.

I just connect two nets using this tutorial without problems. This is a great guide.

Thanks a lot!.

## Stephane · October 17, 2017 at 12:37

Thanks a lot.

Great job, great sharing .. please do share for us beginner Mikrotik admin, some of your useful tips.

## Dablah · November 13, 2017 at 23:57

Solo puedo decir, muchas gracias!

## hardoverflow · December 1, 2017 at 06:31

Thanks for your great guide! One question about routing. On every site i have a few vlans configured. Is it possible to route them over the ipsec tunnel ?

## Pessoft · December 8, 2017 at 22:36

Hi and thank you. Regarding VLAN routing, I didn’t test such configuration, but generally VLANs are working on OSI layer 2 and are terminated on routers when IP routing occurs. So VLANs on 2 sites are usually different logical networks. It might be worth a try to create a tunnel interface (GRE) on top of the IPSec and bridge the VLANs with tunnel. But as mentioned, I didn’t test such setup on Mikrotik platform, so it’s just a guess 😉 Another possibility (as an example) is to make sure that network in the VLAN 100 on the site A is configured correctly in terms of routing and firewall rules, that it can communicate with network in the VLAN 100 on the site B – and vise versa.

## Lars · April 7, 2019 at 18:07

In terms of VLAN routing over IPsec with GRE and with dynamic protocol

ls like RIPv2 you find a very good howto here:

https://administrator.de/wissen/cisco-mikrotik-vpn-standort-vernetzung-dynamischem-routing-398932.html

Unfortunately in German but the WinBox screenshots are self explaining. With “Pessofts” above guide to dynamic IPsec peer addresses it works like a charme !!

Also the added IPsec related links at the end are worth to read !

## zack12821 · December 4, 2017 at 11:45

Thanks Pressoft! I but not luck, your post was not working well on my side. I have same environment as shown in your topology is about one week that i trying to make it work but way , os 6.40.5

not packet send on ipfirewall, ipsec nagociat fail due to time up

## Pessoft · December 8, 2017 at 23:13

Hi! I tested the setup also on the 6.40.5 and it works well. Your description points out that IPSec communication is not flowing between the two routers. Here are the hints to check:

– verify that routers between Mikrotiks and Internet forward port 4500/UDP to Mikrotik device

– verify that firewall on Mikrotik accepts 4500/UDP

– if router on one site is more susceptible to block communication or difficult to configure, set on the opposing site Mikrotik’s IPSec peer configuration to be “passive” and also disable there “Send Initial Contact”

– try to change the IPSec peer exchange mode on both sites to IKEv2

## felixput · December 12, 2017 at 14:41

Hi Pessoft,

Thanks a lot for your guide, it’s really helpful. I got the IPsec connection established and I can reach both routers from both sides.

Unfortunately, I couldn’t ping any devices that other than the router from another side. When I check the connection, there is a ping request however it never got replied.

Can you help me on this one? Thanks

## Pessoft · December 15, 2017 at 00:22

Hi Felix,

Try to test between hosts on both sides. Using the Torch tool you should see incoming icmp packets on gateway interface. This should help you identify on which side is the issue or whether it is general. Also make sure that you have NAT firewall rules set and in the order prior to your masquerading or other src/dst nat rules, which could influence routing of the packets.

## Mathew · June 19, 2018 at 23:42

Hi Pessoft,

Thanks for the guide. I didnt find anything with this type of topology in mind. I tried to do everything here but cant establish IPsec IKEv2 connection between routers (ping). I dont have dynamic public ips so i used static public IP for SA and peers on both sides. Unfortunately at Remote Peers i have local address ether1-gateway address and remote address correct static public address and also no installed SAs. Can you please help me? This is same on both Mikrotik routers. Thanks in advance

## Pessoft · June 20, 2018 at 00:05

Hi Mathew,

Based on your input, it seems that addresses in Remote Peers are looking good. Check in which state your remote peer is using

`/ip ipsec remote-peers print value-list`. If it is established, then peer connection is fine. Otherwise have a look at logs what they say and also verify network between Mikrotiks: Are both local and remote ports in remote peer view using 4500/UDP? Are ISP routers forwarding 4500/UDP traffic to Mikrotik routers? Can you see packet counts increasing in Mikrotik firewall rules allowing 4500/UDP? Are there visible attempts to establish SA? This could give some information about what can be blocking the connection.
## Mathew · June 20, 2018 at 17:58

i have 4500 and it is increasing count and also peer connection is established. Log shows after 2 SA messages no policy found/generated

## Pessoft · June 20, 2018 at 22:01

It seems that peer connection works well, so the next step should be Policy review using

