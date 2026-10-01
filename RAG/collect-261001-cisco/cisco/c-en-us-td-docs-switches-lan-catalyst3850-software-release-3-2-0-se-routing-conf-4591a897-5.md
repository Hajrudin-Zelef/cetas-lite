---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897-5
title: "c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897.md
source_anchor: ""
source_lines: [237, 296]
sha256: dd58fe4f7bb43ce2382d7b33480c6629f1e4ac6289ae3b5144265495c598ffa5
---

# c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897

| Step 7 | copy running-config 				  startup-config Example: Device# copy running-config startup-config    | (Optional) Saves your entries in the configuration file. | 
Enabling Proxy ARP
By default, the Device uses proxy ARP to help hosts learn MAC addresses of hosts on other networks or subnets.
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example: Device> enable   | Enables privileged EXEC mode. Enter your password if prompted.  | 
| Step 2 | configure  				terminal Example:  Device# configure terminal   | Enters the global configuration mode. | 
| Step 3 | interface  				interface-id Example:  Device(config)# interface gigabitethernet 1/0/2  | Enters interface configuration mode, and specifies the Layer 3 interface to configure. | 
| Step 4 | ip 				  proxy-arp Example:  Device(config-if)# ip proxy-arp  | Enables proxy ARP on the interface. | 
| Step 5 | end Example:  Device(config)# end   | Returns to privileged EXEC mode. | 
| Step 6 | show ip interface [interface-id] Example:  Device# show ip interface gigabitethernet 1/0/2  | Verifies the configuration on the interface or all interfaces. | 
| Step 7 | copy running-config 				  startup-config Example: Device# copy running-config startup-config    | (Optional) Saves your entries in the configuration file. | 
Routing Assistance When IP Routing is Disabled
These mechanisms allow the Device to learn about routes to other networks when it does not have IP routing enabled:
Proxy ARP
Proxy ARP is enabled by default. To enable it after it has been disabled, see the “Enabling Proxy ARP” section. Proxy ARP works as long as other routers support it.
Default Gateway
Another method for locating routes is to define a default router or default gateway. All non-local packets are sent to this router, which either routes them appropriately or sends an IP Control Message Protocol (ICMP) redirect message back, defining which local router the host should use. The Device caches the redirect messages and forwards each packet as efficiently as possible. A limitation of this method is that there is no means of detecting when the default router has gone down or is unavailable.
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example: Device> enable   | Enables privileged EXEC mode. Enter your password if prompted.  | 
| Step 2 | configure  				terminal Example:  Device# configure terminal   | Enters the global configuration mode. | 
| Step 3 | ip 				  default-gateway  				ip-address Example:  Device(config)# ip default gateway 10.1.5.1  | Sets up a default gateway (router). | 
| Step 4 | end Example:  Device(config)# end   | Returns to privileged EXEC mode. | 
| Step 5 | show ip 				  redirects Example:  Device# show ip redirects  | Displays the address of the default gateway router to verify the setting. | 
| Step 6 | copy running-config 				  startup-config Example: Device# copy running-config startup-config    | (Optional) Saves your entries in the configuration file. | 
ICMP Router Discovery Protocol (IRDP)
The only required task for IRDP routing on an interface is to enable IRDP processing on that interface. When enabled, the default parameters apply.
You can optionally change any of these parameters. If you change the maxadvertinterval value, the holdtime and minadvertinterval values also change, so it is important to first change the maxadvertinterval value, before manually changing either the holdtime or minadvertinterval values.
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example: Device> enable   | Enables privileged EXEC mode. Enter your password if prompted.  | 
| Step 2 | configure  				terminal Example:  Device# configure terminal   | Enters the global configuration mode. | 
| Step 3 | interface  				interface-id Example:  Device(config)# interface gigabitethernet 1/0/1  | Enters interface configuration mode, and specifies the Layer 3 interface to configure. | 
| Step 4 | ip 				  irdp Example:  Device(config-if)# ip irdp  | Enables IRDP processing on the interface. | 
| Step 5 | ip irdp 				  multicast Example:  Device(config-if)# ip irdp multicast  | (Optional) Sends IRDP advertisements to the multicast address (224.0.0.1) instead of IP broadcasts. | 
| Step 6 | ip irdp 				  holdtime  				seconds Example:  Device(config-if)# ip irdp holdtime 1000  | (Optional) Sets the IRDP period for which advertisements are valid. The default is three times the maxadvertinterval value. It must be greater than maxadvertinterval and cannot be greater than 9000 seconds. If you change the maxadvertinterval value, this value also changes. | 
| Step 7 | ip irdp 				  maxadvertinterval  				seconds Example:  Device(config-if)# ip irdp maxadvertinterval 650  | (Optional) Sets the IRDP maximum interval between advertisements. The default is 600 seconds. | 
| Step 8 | ip irdp 				  minadvertinterval  				seconds Example:  Device(config-if)# ip irdp minadvertinterval 500  | (Optional) Sets the IRDP minimum interval between advertisements. The default is 0.75 times the maxadvertinterval. If you change the maxadvertinterval, this value changes to the new default (0.75 of maxadvertinterval). | 
| Step 9 | ip irdp 				  preference  				number Example:  Device(config-if)# ip irdp preference 2  | (Optional) Sets a device IRDP preference level. The allowed range is –231 to 231. The default is 0. A higher value increases the router preference level. | 
| Step 10 | ip irdp address  				address [number] Example:  Device(config-if)# ip irdp address 10.1.10.10  | (Optional) Specifies an IRDP address and preference to proxy-advertise. | 
| Step 11 | end Example:  Device(config)# end   | Returns to privileged EXEC mode. | 
| Step 12 | show ip 				  irdp Example:  Device# show ip irdp  | Verifies settings by displaying IRDP values. | 
| Step 13 | copy running-config 				  startup-config Example: Device# copy running-config startup-config    | (Optional) Saves your entries in the configuration file. | 
| Note | This command allows for compatibility with Sun Microsystems Solaris, which requires IRDP packets to be sent out as multicasts. Many implementations cannot receive these multicasts; ensure end-host ability before using this command. | 
Configuring Broadcast Packet Handling
Perform the tasks in these sections to enable these schemes:
- Enabling Directed Broadcast-to-Physical Broadcast Translation
- Forwarding UDP Broadcast Packets and Protocols
- Establishing an IP Broadcast Address
- Flooding IP Broadcasts
Enabling Directed Broadcast-to-Physical Broadcast Translation
By default, IP directed broadcasts are dropped; they are not forwarded. Dropping IP-directed broadcasts makes routers less susceptible to denial-of-service attacks.
You can enable forwarding of IP-directed broadcasts on an interface where the broadcast becomes a physical (MAC-layer) broadcast. Only those protocols configured by using the ip forward-protocol global configuration command are forwarded.
You can specify an access list to control which broadcasts are forwarded. When an access list is specified, only those IP packets permitted by the access list are eligible to be translated from directed broadcasts to physical broadcasts. For more information on access lists, see the “Information about Network Security with ACLs" section in the Security Configuration Guide.
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example: Device> enable   | Enables privileged EXEC mode. Enter your password if prompted.  | 
| Step 2 | configure  				terminal Example:  Device# configure terminal   | Enters the global configuration mode. | 
| Step 3 | interface  				interface-id Example:  Device(config)# interface gigabitethernet 1/0/2  | Enters interface configuration mode, and specifies the interface to configure. | 
