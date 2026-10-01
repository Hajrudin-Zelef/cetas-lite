---
id: collect-261001-cisco/cisco/c-en-us-td-docs-ios-xml-ios-sec-conn-dmvpn-configuration-15-s-sec-conn-dmvpn-15-f8753561-4
title: "c-en-us-td-docs-ios-xml-ios-sec-conn-dmvpn-configuration-15-s-sec-conn-dmvpn-15--f8753561"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-ios-xml-ios-sec-conn-dmvpn-configuration-15-s-sec-conn-dmvpn-15--f8753561.md
source_anchor: ""
source_lines: [595, 793]
sha256: eedf298c02e95ede3f3af53fd078a9f2d82f8c609cd012bb6a91068bf3a83b6a
---

# c-en-us-td-docs-ios-xml-ios-sec-conn-dmvpn-configuration-15-s-sec-conn-dmvpn-15--f8753561

17. end
DETAILED STEPS
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example: Device> enable | Enables privileged EXEC mode. | 
| Step 2 | configure  			   			   				terminal Example: Device# configure terminal | Enters global configuration mode. | 
| Step 3 | interface  				tunnel 				 			   			   				number Example: Device(config)# interface tunnel 5  | Configures a tunnel interface and enters interface configuration mode. | 
| Step 4 | ipv6  				address  			  {ipv6-address  			 /  			 prefix-length \| 			  			 prefix-name  			   			 sub-bits  			 /  			 prefix-length Example: Device(config-if) ipv6 address 2001:DB8:1:1::72/64 | Configures an IPv6 address based on an IPv6 general prefix and enables IPv6 processing on an interface. | 
| Step 5 | ipv6  				address  			   			   				ipv6-address  			   			   				/  			   			   				prefix-length  			   			   				link-local Example: Device(config-if)# ipv6 address fe80::2001 link-local | Configures an IPv6 link-local address for an interface and enables IPv6 processing on the interface. | 
| Step 6 | ipv6  				mtu  			   			   				bytes Example: Device(config-if)# ipv6 mtu 1400  | Sets the MTU size of IPv6 packets sent on an interface. | 
| Step 7 | ipv6  				nhrp 				 				authentication  			   			   				string Example: Device(config-if)# ipv6 nhrp authentication examplexx | Configures the authentication string for an interface using the NHRP. | 
| Step 8 | ipv6 				 				nhrp 				 				map  			   			   				ipv6-address  				nbma-address Example: Device(config-if)# ipv6 nhrp map 2001:DB8:3333:4::5 10.1.1.1 | Statically configures the IPv6-to-NBMA address mapping of IPv6 destinations connected to an NBMA network. | 
| Step 9 | ipv6 				 				nhrp 				 				map  				multicast  			   			   				ipv4-nbma-address Example: Device(config-if)# ipv6 nhrp map multicast 10.11.11.99 | Maps destination IPv6 addresses to IPv4 NBMA addresses. | 
| Step 10 | ipv6 				 				nhrp 				 				nhs  			   			   				ipv6-  			   			   				nhs-address Example: Device(config-if)# ipv6 nhrp nhs 2001:0DB8:3333:4::5 2001:0DB8::/64  | Specifies the address of one or more IPv6 NHRP servers. | 
| Step 11 | ipv6 				 				nhrp 				 				network-id  			   			   				network-id Example: Device(config-if)# ipv6 nhrp network-id 99 | Enables the NHRP on an interface. | 
| Step 12 | tunnel  				source  			   			   				ip-address  			  \|  			 ipv6-address 			 \|  			 interface-type  				interface-number Example: Device(config-if)# tunnel source ethernet 0 | Sets the source address for a tunnel interface. | 
| Step 13 | Do one of the 			 following:  Example: Device(config-if)# tunnel mode gre multipoint Example: Device(config-if)# tunnel destination 10.1.1.1 | Sets the encapsulation mode to mGRE for the tunnel interface. or Specifies the destination for a tunnel interface. | 
| Step 14 | Do one of the 			 following: Example: Router(config-if)# tunnel protection ipsec profile vpnprof  Example: Router(config-if)#  tunnel protection psk test1  | Associates a tunnel interface with an IPsec profile.  or Simplifies the tunnel protection configuration for pre-shared key (PSK) by creating a default IPsec profile. | 
| Step 15 | bandwidth  			  {interzone \|  			 total \|  			 session} 			 {default \|  			 zone  			 zone-name}  			 bandwidth-size Example: Device(config-if)# bandwidth total 1200 | Sets the current bandwidth value for an interface to higher-level protocols.  | 
| Step 16 | ipv6 				 				nhrp 				 				holdtime  			   			   				seconds Example: Device(config-if)# ipv6 nhrp holdtime 3600 | Changes the number of seconds that NHRP NBMA addresses are advertised as valid in authoritative NHRP responses. | 
| Step 17 | end Example: Device(config-if)# end | Exits interface configuration mode and returns to privileged EXEC mode. | 
| Note | The NHRP authentication string must be set to the same value on all hubs and spokes that are in the same DMVPN network. | 
| Note | Only IPv4 NBMA addresses are supported, not ATM or Ethernet addresses. | 
| Note | Effective with Cisco IOS XE Denali 16.3 ipv6 nhrp network-id is enabled by default. | 
Verifying DMVPN for IPv6 Configuration
1.   
       
			  
				enable 
			  
		  
2.   
       
			  
				show 
				dmvpn 
			  [ipv4 [vrf 
			 vrf-name] | 
			 ipv6 [vrf 
			 vrf-name]] [debug-condition | [interface 
			 tunnel 
			 number | 
			 peer {nbma 
			 ip-address | 
			 network 
			 network-mask | 
			 tunnel 
			 ip-address}] [static] [detail]] 
		  
3.   
       
			  
				show 
				ipv6 
				nhrp 
			  [dynamic [ipv6-address] | 
			 incomplete | 
			 static] [address | 
			 interface ] [brief | 
			 detail] [purge] 
		  
4.   
       
			  
				show 
			  
			  
				ipv6 
			  
			  
				 
				 
				nhrp 
				multicast 
			 [ipv4-address | 
			 interface | 
			 ipv6-address] 
		  
5.   
       
			  
				show 
				ip 
				nhrp 
				multicast 
			  [nbma-address | 
			 interface] 
		  
6.   
       
			  
				show 
				ipv6 
				nhrp 
				summary 
			  
		  
7.   
       
			  
				show 
				ipv6 
				nhrp 
				traffic 
			  [ 
			 interfacetunnel 
			 number 
		  
8.   
       
			  
				show 
				ip 
				nhrp 
				shortcut 
			  
		  
9.   
       
			  
				show 
				ip 
				route 
			  
		  
10.   
       
			  
				show 
				ipv6 
				route 
			  
		  
11.   
       
			  
				show 
				nhrp 
				debug-condition 
			  
		  
DETAILED STEPS
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example: Device> enable  | Enables privileged EXEC mode. | 
| Step 2 | show  				dmvpn  			  [ipv4 [vrf  			 vrf-name] \|  			 ipv6 [vrf  			 vrf-name]] [debug-condition \| [interface  			 tunnel  			 number \|  			 peer {nbma  			 ip-address \|  			 network  			 network-mask \|  			 tunnel  			 ip-address}] [static] [detail]] Example: Device# show dmvpn 2001:0db8:1:1::72/64 | Displays DMVPN-specific session information. | 
| Step 3 | show  				ipv6  				nhrp  			  [dynamic [ipv6-address] \|  			 incomplete \|  			 static] [address \|  			 interface ] [brief \|  			 detail] [purge] Example: Device# show ipv6 nhrp | Displays NHRP mapping information. | 
| Step 4 | show  			   			   				ipv6  			   			   				  				  				nhrp  				multicast  			 [ipv4-address \|  			 interface \|  			 ipv6-address] Example: Device# show ipv6 nhrp multicast | Displays NHRP multicast mapping information. | 
| Step 5 | show  				ip  				nhrp  				multicast  			  [nbma-address \|  			 interface] Example: Device# show ip nhrp multicast | Displays NHRP multicast mapping information. | 
| Step 6 | show  				ipv6  				nhrp  				summary Example: Device# show ipv6 nhrp summary | Displays NHRP mapping summary information. | 
| Step 7 | show  				ipv6  				nhrp  				traffic  			  [  			 interfacetunnel  			 number Example: Device# show ipv6 nhrp traffic  | Displays NHRP traffic statistics information. | 
| Step 8 | show  				ip  				nhrp  				shortcut Example: Device# show ip nhrp shortcut | Displays NHRP shortcut information. | 
| Step 9 | show  				ip  				route Example: Device# show ip route | Displays the current state of the IPv4 routing table. | 
| Step 10 | show  				ipv6  				route Example: Device# show ipv6 route | Displays the current contents of the IPv6 routing table. | 
| Step 11 | show  				nhrp  				debug-condition Example: Device# show nhrp debug-condition | Displays the NHRP conditional debugging information. | 
Monitoring and Maintaining DMVPN for IPv6 Configuration and Operation
1.   
       
			  
				enable 
			  
		  
2.   
       
			  
				clear 
				dmvpn 
				session 
			  [interface 
				tunnel 
			 number | 
			 peer {ipv4-address | 
			 fqdn-string | 
			 ipv6-address} | 
			 vrf 
			 vrf-name] [static] 
		  
3.   
       
			  
				clear 
				ipv6 
				nhrp 
			  [ipv6-address | 
			 counters 
		  
4.   
       
			  
				debug 
				dmvpn 
			  {all | 
			 error | 
			 detail | 
			 packet} {all | 
			 debug-type} 
		  
5.   
       
