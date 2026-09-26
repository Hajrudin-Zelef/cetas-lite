---
id: collect-260926-mikrotik/mikrotik/operators-security-toolkit-how-to-secure-mikrotik-devices-c19a13f8-2
title: "operators-security-toolkit-how-to-secure-mikrotik-devices-c19a13f8"
domain: mikrotik
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["research"]
source: docs/RAG/lot-mikrotik/forum/misc/operators-security-toolkit-how-to-secure-mikrotik-devices-c19a13f8.md
source_anchor: ""
source_lines: [69, 129]
sha256: 9797e20b634e30e60102b00e3429bba1d83732383b6de0e9cb2fe32f6b4494c0
---

# operators-security-toolkit-how-to-secure-mikrotik-devices-c19a13f8

By default, the mac server runs on all interfaces. That means anyone directly connected to the Mikrotik device can connect or telnet to the device. This allows miscreants inside the network to connect back to the device.
The Mikrotik MAC Connectivity Access provides details to disable MAC Connectivity by default on all entries and then add a local interface to disallow MAC connectivity from the WAN port.
https://help.mikrotik.com/docs/display/ROS/First+Time+Configuration#FirstTimeConfiguration-ProtectingtheRouter
Step 9 – Turn off Neighbor Discovery Protocol
Mikrotik provides an up-to-date neighbor discovery service that includes the MNDP (MikroTik Neighbor Discovery Protocol), CDP (Cisco Discovery Protocol), and LLDP (Link Layer Discovery Protocol) in the Layer2 broadcast domain. It can be used to map out your network.
Miscreants, malware, and APT (advanced persistent threat actors) already inside your network can use Network Discover Protocol to map out your network.
The Best Common Practice (BCP) is to turn off Network Discovery by default. Mikrotik turns network discover ON by default. So it needs to be manually turned off:
To disable neighbor discovery on all interfaces:
/ip neighbor discovery-settings set discover-interface-list=none
Step 10 – Limit WHO on the Internet Can Access Your Mikrotik!
Limit which services can access the router on the Internet-facing “public interfaces” (see illustration). This is done with Mikrotik’s Firewall Service.
There are core access list/firewall filter principles for how you protect your router, network, services, and your organization. Protecting Routers, Switches, and Network Devices is a good video tutorial on YouTube that walks through these principles. In this example, IP connectivity on the public interface must be limited by the Mikrotik firewall feature. In this example, we will accept only ICMP(ping/traceroute), IP Winbox, and ssh access.
/ip firewall filter
add chain=input connection-state=established,related action=accept comment="accept established,related";
add chain=input connection-state=invalid action=drop;
add chain=input in-interface=ether1 protocol=icmp action=accept comment="allow ICMP";
add chain=input in-interface=ether1 protocol=tcp port=8291 action=accept comment="allow Winbox";
add chain=input in-interface=ether1 protocol=tcp port=22 action=accept comment="allow SSH";
add chain=input in-interface=ether1 action=drop comment="block everything else";
Mikrotik Firewall References
The is a mix of guidelines and references. Here are some to explore when building your firewall access rules to protect your router, network, and organization.
- Mikrortik’s Orginal Security Recommendations – Manual: Securing Your Router – this has someone good examples that did not get carried over to Mikrotik’s new Confluence-based wiki. Look in the “firewall” section.
- Mikrotik’s Documentation: First Time Configuration – this has a section on firewalls.
- Unimus Guide: Validating the security of your MikroTik routers network-wide
Step 11 – Don’t let your Mikrotik be used as a “DDoS Proxy!”
Mikrotik’s proxy, socks, UPnP, and other services are getting turned ON “accidentally,” turning the Mikrotik into a powerful DDoS Weapon. These features, combined with other Mikrotik capabilities, turn the device into a “bot” that becomes a member of the threat actor’s BOTNET.
DO NOT LET MISCREANTS USE YOUR MIKROTIK DEVICE FOR CRIMINAL ACTIVITY!
Mikrotik disables these by default. Check if they are turned on and manually turn them off – to be sure:
- MikroTik caching proxy
/ip proxy set enabled=no
- MikroTik socks proxy
/ip socks set enabled=no
- MikroTik UPNP service
ip upnp set enabled=no
- MikroTik dynamic name service or IP cloud
/ip cloud set ddns-enabled=no update-time=no
Step 12 – Start Cleaning Up Your “Compromised” Router
At this point, it is time to see if your Mikrotik device has been compromised. As you can see in this Mikrotik blog post – MĒRIS BOTNET – Mikrotik devices are sought after by threat actors and miscreants. As mentioned in the beginning, we are making the assumption your device has been compromised.
- Check the Mikrotik Scheduler. The threat actor will install a rule that executes the script. Those scripts are evolving and creative. Check every scheduled script (under System → Scheduler). If you don’t know the script, download a copy, then delete it. Look for scripts with the fetch () method.
- Check Files on the Mikrotik Router – Delete Unknown Files. Check the files on your Mikrotik. If you do not recognize it, download a copy, then delete the file.
- Check the SOCKS Proxy. We covered this earlier, but it is worth checking again. An unknown SOCKS proxy server enabled on your router indicates that the router has been compromised.. You’ll find the setting under IP → SOCKS; if you do not use it, disable it;
- Check for L2TP Clients. Threat Actors will use L2TP to control the router as part of a BOTNET. Look for any L2TP clients called lvpn, (or any other L2TP client unfamiliar to you). Delete these clients.
- Check all the Firewall Rules! A threat actor will “open a hole” in the firewall to allow them to get remote access. For example, a firewall rule that allows remote access through port 5678 was used as part of the early versions of the MĒRIS botnet. Remove this rule, then check each firewall rule.
There is more. Unimus created a guide with more detailed scripts to help recover a compromised Mikrotik device (see Validating the security of your MikroTik routers network-wide).
Step 13 – Prevent Trickbot from using Your Router
Trickbot is a modular trojan that’s been around since 2016 and is often used by cybercriminals to deliver ransomware or other malware. The miscreants using Trickbot have found they can use Mikrotik Routers as “proxies” to hide their Trickbot Command and Control.
The Microsoft Defender for IoT Research Team (Section 52) published details of how Trickbot exploits Mikrotik in Uncovering Trickbot’s use of IoT devices in command-and-control infrastructure. They also have an open-source tool to help check Mikrotik devices (see routeros-scanner on Github).
The Trickbot “miscreant threat actors” are using the same approach we are working to prevent in the previous defensive steps:
- Trickbot crews are using default MikroTik passwords where the device has been plugged in before being securely configured.
- Trickbot crews scan for Mikrotik devices, then launch brute force “password guessing” attacks. Microsoft has seen attackers use some unique passwords that probably were harvested from other MikroTik devices.
- Trickbot crews are exploiting CVE-2018-14847 on devices with RouterOS versions older than 6.42. This vulnerability gives the attacker the ability to read arbitrary files like user.dat, which contain passwords.
Microsoft’s Section 52 Team we able to explore Mikrotik confirmation commands to unravel the Trickbot crew’s source and intent. For example, we observed attackers issuing the following commands:
/ip firewall nat add chain=dstnat proto=tcp dst-port=449   to-port=80 action=dst-nat to-addresses=dst-address= 
From the command, we can understand the following:
- A new rule, similar to iptables, is created.
- The rule redirects traffic from the device to a server
- The redirected traffic is received from port 449 and redirected to port 80
The command looks like a legitimate network address translation (NAT) command that allows the NAT router to perform IP address rewriting. Here the Trickbot miscreants have compromised the Mikrotik router and configured the NAT for malicious activity. Trickbot’s Command and Control (C2C) is known for using ports 443 and 449.
How can you check if TrickBot’s Firewall Rule is installed?
Run the following command to detect if the NAT rule was applied to the device (completed by the tool as well):
/ip firewall nat print
