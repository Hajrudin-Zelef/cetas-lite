---
id: collect-261001-general-networking/general-networking/t5-switching-switching-best-practices-loop-guard-root-guard-etc-m-p-958650-highl-38959b3b
title: "t5-switching-switching-best-practices-loop-guard-root-guard-etc-m-p-958650-highl-38959b3b"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-switching-switching-best-practices-loop-guard-root-guard-etc-m-p-958650-highl-38959b3b.md
source_anchor: ""
source_lines: [1, 176]
sha256: bdb480101bdf96ffa470fbcd23cdba6bbd26d77f15f0ef30c122b1a375079f38
---

# t5-switching-switching-best-practices-loop-guard-root-guard-etc-m-p-958650-highl-38959b3b

Switching : best practices, loop guard, root guard etc
						
					
					
				
			
		
	
			
	
	
	
	
	
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-29-2008 12:26 AM - edited 03-05-2019 09:27 PM
Hi,
We are having two 6509E core switches along with 22 3560g edge switches. The edge switches have dual uplinks, one to each core. We have enabled rapid spanning tree on all switches. I am somewhat confused about on which edge switch ports "loop guard" and "root guard" "bpduguard" should be enabled. We dont want a situation where someone connects another switch and that becomes the root.The core switch vlans are on HSRP with even vlans active on core2 and odd vlans active on core1.
Would someone please explain about the best solutions/practices that they have come across?
Thanks in advance.
Sonu
- Labels:
- 
						
							
		
			Other Switching
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-29-2008 01:16 AM
hi Sonu,
The loop guard feature is enabled on a per-port basis.if BPDUs are not received on the trunk port for only one particular VLAN, only that VLAN is blocked (moved to loop-inconsistent STP state).Loop guard must be enabled on the non-designated ports (more precisely, on root and alternate ports) for all possible combinations of active topologies.
BPDU guard and root guard are similar, but their impact is different. BPDU guard disables the port upon BPDU reception if PortFast is enabled on the port. The disablement effectively denies devices behind such ports from participation in STP.
Root guard allows the device to participate in STP as long as the device does not try to become the root. If root guard blocks the port, subsequent recovery is automatic. Recovery occurs as soon as the offending device ceases to send superior BPDUs.
The root guard is used on designated ports, and it does not allow the port to become non-designated. The loop guard works on non-designated ports and does not allow the port to become designated through the expiration of max_age. The root guard cannot be enabled on the same port as the loop guard. When the loop guard is configured on the port, it disables the root guard configured on the same port.
Enable
loopguard: nondesignated ports (root/blocking)
rootguard: designated ports
bpduguard: portfast enabled ports.
HTH,
regards,
shri
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-01-2008 09:38 PM
Hi shri,
Thanks for the reply. We are using rapid spanning tree throughout. Portfast command is used make a port to be in the forwarding state faster in places a802.1d is running.But since we use rstp there is no need for "spanning-tree portfast" and hence bpduguard. Is this correct?
Also on access ports where users are connected do i need to enable roots guard to prevent an STP election in case someone connects a switch to the access port? ( Assume rstp is running)
Thanks in advance.
Sonu
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-23-2008 10:30 AM
Even when running Rapid Spanning Tree, it's still a good practice to run portfast on ports you know will connect to end devices, since they are marked as "Edge" ports and easier to identify. Just make sure to run BPDUguard or BPDUFilter as a protection method. More info:
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-27-2008 04:12 PM
Hi, if you do not run portfast on ports you will connect to end devices (hosts, routers, etc), every time a topology change occurs in the access vlan of these ports, they will transition to blocking and then from blocking to forwarding state in 30 sec. That is because the ports are not "edge" ports.
Regards,
Max
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-27-2008 04:33 PM
Sonu:
Read this link. I think it will answer most of your questions.
http://www.cisco.com/warp/public/473/185.pdf
HTH
Victor
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-27-2008 04:33 PM
Sonu:
Read this link. I think it will answer most of your questions.
http://www.cisco.com/warp/public/473/185.pdf
HTH
Victor
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-27-2008 04:33 PM
Sonu:
Read this link. I think it will answer most of your questions.
http://www.cisco.com/warp/public/473/185.pdf
HTH
Victor
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-01-2008 11:39 AM
Sonu,
I would like to clarify your question on use of portfast on edge ports: here is what cisco says on that " Unlike PortFast, an edge port that receives a BPDU immediately loses edge port status and becomes a normal spanning tree port. At this point, there is a user-configured value and an operational value for the edge port state. The Cisco implementation maintains that the PortFast keyword be used for edge port configuration. This makes the transition to RSTP simpler."
You are correct - there is no need for port-fast config on edge port.
- as per root guard on edge port, I would rather enable bpdu-guard, this will make sure no one can plug in a switch to the network. rootguard will still allow a switch to be connected as it kicks in only when it detects another switch with lower bridge priority than current root.
here is a good reference for STP features use:
-serg
[use the rating]
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-01-2008 01:02 PM
In simpler terms, since access switches have edge ports - enable bpduguard on them.
Core switches will have trunk links to access switches, so enable rootguard on core switch ports so that new switch doesnt take over as root with superior bpdu.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-01-2008 06:04 PM
Dear Serguei, I'm not agree with you. With RSTP (or Rapid PVST+ or MST), if you don't enable portfast on ports connected to end stations (or routers), when there is a TC the agreement-proposal handshake does not work fine on those ports. The switch running RSTP sends a BPDU with the proposal bit set but there will be no answer from the router/end station, so the switch thinks there is a legacy switch running STP connected to that port and the transition from blocking to forwarding takes 30 seconds. That's why you need to configure these ports as edge-ports (with the "portfast" command).
Here is a fragment of the "3560 configuration guide":
"Rapid Convergence
The RSTP provides for rapid recovery of connectivity following the failure of a switch, a switch port, or a LAN. It provides rapid convergence for edge ports, new root ports, and ports connected through point-to-point links as follows:
â€¢Edge ports: If you configure a port as an edge port on an RSTP switch by using the spanning-tree portfast interface configuration command, the edge port immediately transitions to the forwarding state. An edge port is the same as a Port Fast-enabled port, and you should enable it only on ports that connect to a single end station.
Regards,
Max
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-01-2008 07:18 PM
Max, looks like you are right here. Cisco is using fortfast command in very confusing way to define edge ports. it's not very clear from documentation.
thanks for clarification.
-serg
