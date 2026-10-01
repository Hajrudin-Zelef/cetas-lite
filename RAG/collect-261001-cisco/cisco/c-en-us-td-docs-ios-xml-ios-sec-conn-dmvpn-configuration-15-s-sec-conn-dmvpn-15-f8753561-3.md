---
id: collect-261001-cisco/cisco/c-en-us-td-docs-ios-xml-ios-sec-conn-dmvpn-configuration-15-s-sec-conn-dmvpn-15-f8753561-3
title: "c-en-us-td-docs-ios-xml-ios-sec-conn-dmvpn-configuration-15-s-sec-conn-dmvpn-15--f8753561"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-ios-xml-ios-sec-conn-dmvpn-configuration-15-s-sec-conn-dmvpn-15--f8753561.md
source_anchor: ""
source_lines: [325, 594]
sha256: ac9258fcb28fdd2580aed81ca2343e14c8302d71eccaff4bc53eaab4729bb276
---

# c-en-us-td-docs-ios-xml-ios-sec-conn-dmvpn-configuration-15-s-sec-conn-dmvpn-15--f8753561

15. end
DETAILED STEPS
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example: Device> enable | Enables privileged EXEC mode. | 
| Step 2 | configure  			   			   				terminal Example: Device# configure terminal | Enters global configuration mode. | 
| Step 3 | interface  				tunnel 				 			   			   				number Example: Device(config)# interface tunnel 5  | Configures a tunnel interface and enters interface configuration mode. | 
| Step 4 | ipv6  				address  			  {ipv6-address  			 /  			 prefix-length \| 			  			 prefix-name  			   			 sub-bits  			 /  			 prefix-length Example: Device(config-if)# ipv6 address 2001:DB8:1:1::72/64 | Configures an IPv6 address based on an IPv6 general prefix and enables IPv6 processing on an interface. | 
| Step 5 | ipv6  				address  			   			   				ipv6-address  			   			   				/  			   			   				prefix-length  			   			   				link-local Example: Device(config-if)# ipv6 address fe80::2001 link-local | Configures an IPv6 link-local address for an interface and enables IPv6 processing on the interface. | 
| Step 6 | ipv6  				mtu  			   			   				bytes Example: Device(config-if)# ipv6 mtu 1400  | Sets the maximum transmission unit (MTU) size of IPv6 packets sent on an interface. | 
| Step 7 | ipv6  				nhrp  				authentication  			   			   				string Example: Device(config-if)# ipv6 nhrp authentication examplexx | Configures the authentication string for an interface using the NHRP. | 
| Step 8 | ipv6 				 				nhrp 				 				map  				multicast  				dynamic Example: Device(config-if)# ipv6 nhrp map multicast dynamic  | Allows NHRP to automatically add routers to the multicast NHRP mappings. | 
| Step 9 | ipv6 				 				nhrp 				 				network-id  			   			   				network-id Example: Device(config-if)# ipv6 nhrp network-id 99 | Enables the NHRP on an interface. Effective with Cisco IOS XE Denali 16.3 ipv6 nhrp network-id is enabled by default. | 
| Step 10 | tunnel  				source  			   			   				ip-address  			  \|  			 ipv6-address 			 \|  			 interface-type  				interface-number Example: Device(config-if)# tunnel source ethernet 0 | Sets the source address for a tunnel interface. | 
| Step 11 | tunnel  				mode 				 			  {aurp \|  			 cayman \|  			 dvmrp \|  			 eon \|  			 gre\|  			 gre  				multipoint[ipv6] \|  			 gre  				ipv6 \|  			 ipip  			 decapsulate-any] \|  			 ipsec  				ipv4 \|  			 iptalk \|  			 ipv6\|  			 ipsec  			 ipv6 \|  			 mpls \|  			 nos \|  			 rbscp Example: Device(config-if)# tunnel mode gre multipoint | Sets the encapsulation mode to mGRE for the tunnel interface. | 
| Step 12 | Do one of the 			 following: Example: Router(config-if)# tunnel protection ipsec profile vpnprof  Example: Router(config-if)#  tunnel protection psk test1  | Associates a tunnel interface with an IPsec profile.  or Simplifies the tunnel protection configuration for pre-shared key (PSK) by creating a default IPsec profile. | 
| Step 13 | bandwidth  			  {kbps \|  			 inherit 			 [kbps] \|  			 receive 			 [kbps]} Example: Device(config-if)# bandwidth 1200 | Sets the current bandwidth value for an interface to higher-level protocols. | 
| Step 14 | ipv6 				 				nhrp 				 				holdtime  			   			   				seconds Example: Device(config-if)# ipv6 nhrp holdtime 3600 | Changes the number of seconds that NHRP NBMA addresses are advertised as valid in authoritative NHRP responses. | 
| Step 15 | end Example: Device(config-if)# end | Exits interface configuration mode and returns to privileged EXEC mode. | 
| Note | The NHRP authentication string must be set to the same value on all hubs and spokes that are in the same DMVPN network. | 
| Note | Effective with Cisco IOS XE Denali 16.3 ipv6 nhrp map multicast dynamic is enabled by default. | 
Configuring the NHRP Redirect and Shortcut Features on the Hub
1.   
       
			  
				enable
				
			  
		  
2.   
       
			  
				configure 
			  
			  
				terminal 
			  
		  
3.   
       
			  
				interface 
				tunnel
				
			  
			  
				number
				
			  
		  
4.   
       
			  
				ipv6 
				address 
			  {ipv6-address 
			 / 
			 prefix-length |
			 
			 prefix-name 
			  
			 sub-bits 
			 / 
			 prefix-length 
		  
5. Do one of the following:
6.   
       
			  
				ipv6 
				nhrp 
				shortcut 
			  
		  
7. end
DETAILED STEPS
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example: Device> enable | Enables privileged EXEC mode. | 
| Step 2 | configure  			   			   				terminal Example: Device# configure terminal | Enters global configuration mode. | 
| Step 3 | interface  				tunnel 				 			   			   				number Example: Device(config)# interface tunnel 5  | Configures a tunnel interface and enters interface configuration mode. | 
| Step 4 | ipv6  				address  			  {ipv6-address  			 /  			 prefix-length \| 			  			 prefix-name  			   			 sub-bits  			 /  			 prefix-length Example: Device(config-if)# ipv6 address 2001:DB8:1:1::72/64 | Configures an IPv6 address based on an IPv6 general prefix and enables IPv6 processing on an interface. | 
| Step 5 | Do one of the 			 following: Example: Device(config-if)# ipv6 nhrp redirect Example: Device(config-if)# ipv6 nhrp redirect interest | Enables NHRP redirect. or Enables the user to specify an ACL. | 
| Step 6 | ipv6  				nhrp  				shortcut Example: Device(config-if)# ipv6 nhrp shortcut | Enables NHRP shortcut switching. | 
| Step 7 | end Example: Device(config-if)# end | Exits interface configuration mode and returns to privileged EXEC mode. | 
| Note | You must configure the ipv6 nhrp redirect command on a hub. | 
| Note | Effective with Cisco IOS XE Denali 16.3 ipv6 nhrp shortcut is enabled by default. | 
Configuring the Spoke for IPv6 over DMVPN
Perform this task to configure the spoke for IPv6 over DMVPN.
1.   
       
			  
				enable
				
			  
		  
2.   
       
			  
				configure 
			  
			  
				terminal 
			  
		  
3.   
       
			  
				interface 
				tunnel
				
			  
			  
				number
				
			  
		  
4.   
       
			  
				ipv6 
				address 
			  {ipv6-address 
			 / 
			 prefix-length |
			 
			 prefix-name 
			  
			 sub-bits 
			 / 
			 prefix-length 
		  
5.   
       
			  
				ipv6 
				address 
			  
			  
				ipv6-address 
			  
			  
				/ 
			  
			  
				prefix-length 
			  
			  
				link-local 
			  
		  
6.   
       
			  
				ipv6 
				mtu 
			  
			  
				bytes 
			  
		  
7.   
       
			  
				ipv6 
				nhrp
				
				authentication 
			  
			  
				string 
			  
		  
8.   
       
			  
				ipv6
				
				nhrp
				
				map 
			  
			  
				ipv6-address 
				nbma-address 
			  
		  
9.   
       
			  
				ipv6
				
				nhrp
				
				map 
				multicast 
			  
			  
				ipv4-nbma-address 
			  
		  
10.   
       
			  
				ipv6
				
				nhrp
				
				nhs 
			  
			  
				ipv6- 
			  
			  
				nhs-address 
			  
		  
11.   
       
			  
				ipv6
				
				nhrp
				
				network-id 
			  
			  
				network-id 
			  
		  
12.   
       
			  
				tunnel 
				source 
			  
			  
				ip-address 
			  | 
			 ipv6-address
			 | 
			 interface-type 
				interface-number 
		  
14. Do one of the following:
15.   
       
			  
				bandwidth 
			  {interzone | 
			 total | 
			 session}
			 {default | 
			 zone 
			 zone-name} 
			 bandwidth-size 
		  
16.   
       
			  
				ipv6
				
				nhrp
				
				holdtime 
			  
			  
				seconds 
			  
		  
