---
id: collect-261001-automatisation-infra/automatisation-infra/gns3vault-ccna-r-s-labs-list-in-order
title: "gns3vault-ccna-r-s-labs-list-in-order"
domain: automatisation-infra
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-automatisation-infra/gns3vault-ccna-r-s-labs-list-in-order.md
source_anchor: ""
source_lines: [1, 78]
sha256: 19c24b3462ee5fea44a11788f270eefc8ffc135afd3697d69b2b700402bf299f
---

# gns3vault-ccna-r-s-labs-list-in-order

Here is a list of all CCNA R&S labs on GNS3Vault and in correct order.

**Wireshark**

- Perform a packet capture and look at the layers of the OSI model.
- Perform a packet capture of a DHCP client that receives an IP address from a DHCP server.
- Perform a packet capture of a TCP 3 way handshake.

**ARP**

- Check the ARP table of your computer. What do you see here?
- Capture an ARP request and ARP reply in wireshark.

**Network Management**

- Configure SSH on your router or switch so that you can log in remotely with putty.
- (optional): Take a look at CCP (Cisco Configuration Professional) and use it to connect to your router.

**Switching:**

- Take a look at a switch MAC address table. What dynamic entries do you see?
- Configure port-security
- Connect two computers to a switch and put them in VLAN 10.
- Configure a trunk between two switches.
- Test the different switchport modes:
  - access
  - trunk
  - dynamic auto
  - dynamic desirable
- Use VTP to exchange VLAN information between 2-3 switches.

**Etherchannel**

**Spanning-Tree**

**Subnetting**

- I don’t have a specific subnetting lab. You need to get enough practice with subnetting. To force yourself to become familiar with subnetting, don’t use subnet mask 255.255.255.0 (/24) in some of your labs but use other subnet masks like /25, /26, /27, /27, etc.

**General Routing**

**Gateway Redundancy**

**Routing Protocols**

**Security**

**NAT/PAT**

**PPP**

**IPv6**

**Network Monitoring**

I don’t have a lab for this but it’s a good exercise to download LibreNMS, see if you can monitor one of your routers or switches with it.

Same thing for syslog, download Kiwi Syslog and send syslog information from your router or it.

**CDP**

**Misc**

- see if you can do a password recovery on your router.
- see if you can upgrade the IOS on your switch or router.

 
              
                  
                  
                

            
           
          
            
            
              I took down the website and migrated the labs to the GitHub repository:
