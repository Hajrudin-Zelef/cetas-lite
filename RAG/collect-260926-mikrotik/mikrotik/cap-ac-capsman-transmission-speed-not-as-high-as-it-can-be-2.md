---
id: collect-260926-mikrotik/mikrotik/cap-ac-capsman-transmission-speed-not-as-high-as-it-can-be-2
title: "nov/16/2022 11:00:52 by RouterOS 7.6"
domain: mikrotik
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["intel"]
source: docs/RAG/lot-mikrotik/forum/wifi-capsman/cap-ac-capsman-transmission-speed-not-as-high-as-it-can-be.md
source_anchor: ""
source_lines: [141, 297]
sha256: 7b9bd8733e3b12cd3f3063ab34811d3036d2aea0659da7a5012159f8dd4c4185
---

# nov/16/2022 11:00:52 by RouterOS 7.6

PS:see this for the Intel rate adaptation algorithm (RAA) : https://remy.grunblatt.org/pdfs/Simulation_and_Performance_Evaluation_of_the_Intel_Rate_Adaptation_Algorithm.pdf

             
            
           
          
            
            
              OK,

thanks for the comprehensive answer  I think it was the last answer that most described what I was thinking of.

However, as far as CAPsMAN is concerned - local forwarding does not work for me, the moment I turn off “Discovery Interfaces” I get the error “CAP did not find suitable CAPsMAN” over and over again.

MGMT connection is via VLAN (it is added to the interface, it gets an address from DHCP), on CAP I have a bridge done (wlan1+wlan2+ether1), CAP sees HEX, firewall on HEX allows UDP/5246,5247, unfortunately it does not work. The moment I turn on “Discovery Interfaces” - everything returns to normal (of course, then I have to add a bridge in the datapath configuration on the capsman side). Maybe continuing with the topic of speed optimization you will have some idea what I am doing wrong.

             
            
           
          
            
            
              When toggling between local- and capsman-forwarding you don’t change discovery interface … that one is necessary for provisioning. You only change setting in datapath part on capsman. Also make sure bridge, mentioned in capsman datapath setup exists on clients and is configured with necessary VLANs.

             
            
           
          
            
            
              Then maybe I misunderstood something, because reading various information (including MUMs) everywhere using the Local forwarding option (in datapath on CAPSMAN), on CAP discovery interfaces was empty, and capsman address was given.

An interesting fact is that if I leave discovery interfaces set, master interfaces still have CAPsMAN forwarding, while slave interfaces display local forwarding (regardless of how it is set in datapath). For testing at this point on both sides, I disabled vlan filtering on the bridges.

             
            
           
          
            
            
              Re. discovery: CAP device can either auto-discover CAPsMAN if it’s present in same L2 network and it uses discovery interface for that. Or it can communicate via IP to CAPsMAN in any routed network but needs CAPsMAN’s address (the usual routing rules apply).

Re. local/remote forwarding: check CAPsMAN config that all applicable datapaths have local-forwarding enabled explicitly, I don’t remember which is default setting.

I don’t think you can see actual setting on CAP under wireless. You should see all wireless interfaces as bridge ports on CAP client though - use *print* command, export won’t show that info.

             
            
           
          
            
            
              OK, in my case it can be L2 option (but I don’t understand why IP connection doesn’t work - maybe this will be next lesson for me )

And about local forwarding - that was my missclick in Configurations tab (i missed that in one configuration “Local forwarding” option was not hidden, and not checked, so it overrides my datapath config) - now all works brilliant. But VLAN Filtering is not enabled yet.

             
            
           
          
            
            
              
Make sure you enable safe mode before enabling VLAN-filtering. If things go wrong, they go really bad on L2.

             
            
           
          
            
            
              There are several options for placing access points on the network.

Schematically, I displayed them in this figure.

In option #3, the “local forwarding” mode didn’t work for me anywhere.

An example of option #3 is a building with several floors.

             
            
           
          
            
            
              You can seriously have a cap connect to caps-man across layer 3. I had the wap at my in-laws controlled by the router at my house.

I did this over direct port forwarding.

Over VPN

And over xvlan

Point being you have to make sure you allow the traffic from the caps to the caps-manager.

Once that’s done… Caps can connect and get a config with nothing more than caps-man enabled after reset (if you are in the same IP scope). If you need to hit a server elsewhere… You tell the device to go find the caps-manager.

             
            
           
          
            
            
              With local forwarding it’s vital that whole L2 infrastructure and settings on all involved devices can support the needs. In case #3 by @BrateloSlava it could happen that not all switches are configured with all necessary VLANs. CAPsMAN forwarding largely skips the proper LAN infrastructure setup problems as it tends to work if CAP can connect CAPsMAN. It also allows some more fancy stuff (like filtering traffic between wireless clients in a single point - CAPsMAN bridge) that is otherwise hard to achieve.

But as explained, performance wise it can be a hog …

And no, multiple floor physical layout doesn’t explain why local forwarding doesn’t work …

             
            
           
          
            
            
              
Yup, as I said, usual routing rules apply (and firewall as well, I thought that’s obvious).

             
            
           
          
            
            
              @**zett93** Show me, please, full text configuration of your AP. Without serial numbers and etc. With bridge and VLANs info.

And a question along the way. When you enable “local forwarding” mode for access points, do wireless clients not get IP addresses from DHCP server?

@**mkx** Most likely I do not understand or do not know something. Consider scheme #3. Imagine a simple network with no VLAN, the “primary” router is a DHCP and CAPsMAN server. The router has only one bridge configured. Switches are connected through the LAN ports of the router. Some switches can be cascaded. Access points are connected to the switch ports. All switches are also configured with a single bridge. We get one “big” L2 network.

Activate the “local forwarding” mode. Wireless clients receive IP addresses and start moving between access points. As soon as the client connects to an access point that is connected to a different switch, not the same one as the previous one, the connection is lost. This does not happen when using “capsman forwarding”.

If necessary, I will make a simple text configuration of the devices of such a network. To clarify configuration errors.

*Why am I describing such a “simple” network scheme (several switches and no VLAN) - it’s easier for me to assemble it from the equipment that I have at home.*

             
            
           
          
            
            
              
This sounds to me as trouble with switches not updating their ARP tables properly. Or some funky configuration somewhere in L2 network, can’t tell without some advanced troubleshooting. And I’ve seen this bug happen even with switches by most renowned brand. I can imagine even some “security features” on switches to cause such behaviour.

It’s not a bug of local-forwarding but in such case capsman forwarding is a (welcome) workaround. If this was a general problem with wireless clients roaming, then roaming in wireless network without CAPsMAN would not work at all … while in reality it does (and until CAPsMAN starts to speak 802.11 r/k/v it works equally well).

             
            
           
          
            
            
              @**mkx**

I assembled a “small and simple network” at home according to scheme #3. Before that, everything worked for me according to scheme #1. The access points were directly connected to the router, and the rest of the devices worked through the CRS326 switch.

