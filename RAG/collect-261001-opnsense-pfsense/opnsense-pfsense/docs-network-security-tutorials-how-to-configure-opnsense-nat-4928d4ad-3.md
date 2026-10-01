---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/docs-network-security-tutorials-how-to-configure-opnsense-nat-4928d4ad-3
title: "docs-network-security-tutorials-how-to-configure-opnsense-nat-4928d4ad"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/docs-network-security-tutorials-how-to-configure-opnsense-nat-4928d4ad.md
source_anchor: ""
source_lines: [138, 241]
sha256: 939b6605c9e052214a3e826dbf7a68aa43896679473ba881882bade8bad77d21
---

# docs-network-security-tutorials-how-to-configure-opnsense-nat-4928d4ad

| Option | Description | 
|---|---|
| Disabled | Check this option to disable the rule without removing it. | 
| Do not NAT | Enabling this option will disable NAT for traffic matching this rule and stop processing Outbound NAT rules.Hint: this option is rarely used; don't use it unless you're sure you know what you're doing. | 
| Interface | Which interface the rule should apply to. The majority of the time, this will be WAN. | 
| TCP/IP version | IPv4 or IPv6. | 
| Protocol | In typical scenarios, this will be TCP | 
| Source | The source network to match | 
| Source / Invert | Invert match in `Source` field. | 
| Source port range | When applicable, the source port on which we should match. This is almost always random and almost never equals the destination port range (and should almost always be 'any'). | 
| Destination / Invert | Invert match in `Destination` field. | 
| Destination | Enter the destination network for the outbound NAT mapping. | 
| Destination port range | Service port the traffic is using. | 
| Translation / target | Packets matching this rule will be mapped to the IP address given here.If you want this rule to apply to another IP address rather than the IP address of the interface chosen above, select it here (you will need to define Virtual IP addresses on the interface first). | 
| Log | Put packets matching this rule in the logs. Use this sparingly to avoid overflowing the logs. | 
| Pool Options | This option is explained in the previous section. The default is to use Round robin. Only Round Robin types are compatible with Host Aliases. Subnets of any type can be used. **Round Robin:** Iterates over the translation addresses.** Random:** Chooses an address at random from the translation address pool.**Source Hash:** Determines the translation address by hashing the source address, ensuring that the redirection address is always the same for a given source.**Bitmask:** Uses the subnet mask while keeping the last portion the same; 172.16.10.50 → x.x.x.50.**Sticky Address:** When using the Random or Round Robin pool types, the Sticky Address option ensures that a specific source address is always mapped to the same translation address. | 
| Translation / port | Which port to use on the target | 
| Static-port | Prevents pf(4) from modifying the source port on TCP and UDP packets. | 
| Set local tag | Set a tag that other NAT rules and filters can check for. | 
| Match local tag | Check for a tag set by another rule | 
| No XMLRPC sync | Prevent this rule from being synced to a backup host. (Checking this on the backup host has no effect.) | 
| Description | A description to easily find the rule in the overview. | 

## Configure NPTv6 for IPv6 Networks

Network Prefix Translation, abbreviated as NPTv6, is used to convert IPv6 addresses. A prevalent use for this is to convert global ("WAN") IP addresses to local ones. In this context, it resembles NAT; however, NPTv6 only allows one-to-one address mapping, in contrast to NAT, which often converts one external IP to several internal addresses.

NPTv6 routes may be found under **Firewall** > **NAT** > **NPTv6**.  New regulations may be included by selecting **Add** in the top right corner.  A brief summary of the fields for NPTv6 NAT is given below.

| Field | Description | 
|---|---|
| **Disabled** | Deactivates this rule without necessitating its removal. | 
| **Interface** | To which interface should this rule be applied? This will often be a WAN interface. | 
| **Internal IPv6 Prefix** | The internal IPv6 prefix used inside the local area network(s). This will substitute the prefix of the destination address in incoming packets. The chosen prefix size will be applied to the external prefix. | 
| **External IPv6 Prefix** | The external IPv6 prefix. This will substitute the prefix of the source address in outgoing packets. | 
| **Category** | The category to which this rule pertains may serve as a filter in the overview. | 
| **Description** | A description that enables the reader to quickly ascertain the purpose of this rule in the overview. | 

**Figure 4.** *Configuring NPTv6 for IPv6 Networks in OPNsense*

## Real-World Examples for NAT Configurations in OPNsense

In this section we will give some real world scenarios for NAT configuration on OPNsense firewall:

- Port Forwarding for Web Servers
- Port Forwarding for SSH and RDP Services on Custom Ports
- Outbound NAT for Accessing a Remote Service via External IP

### How to Configure Port Forwarding For Web Services

Businesses that provide a service to their customers via the Internet must make their applications or web servers accessible from the Internet. Assume your company has two separate web servers in the DMZ network and one public IP address. Both the HTTP and HTTPS ports on these web servers should be accessible from anywhere in the world using the same public IP address. To accomplish this, you may define the port forwarding rules in your OPNsense. You may configure your rules in such a way that while requests coming to 80 and 443 ports are redirected to the first web server, the second web server is accessible via 81 and 8443 ports. For this configuration, you may follow the next steps below.

| Server Name | External IP | External Port | Local IP | Local Port | 
|---|---|---|---|---|
| WebServer1 | Public Internet IP | 80 | 10.10.10.13 | 80 | 
| WebServer1 | Public Internet IP | 443 | 10.10.10.13 | 443 | 
| WebServer2 | Public Internet IP | 81 | 10.10.10.14 | 80 | 
| WebServer2 | Public Internet IP | 8443 | 10.10.10.14 | 443 | 

**Figure 5.** *Port Forwarding topology for web services*

After completing the port forwarding configurations on your OPNsense firewall, HTTP(80) and HTTPS(443) requests for your WAN IP will be redirected to the WebServer1(10.10.10.13), while port 81 and port 8443 requests for your WAN IP will be redirected to the WebServer2(10.10.10.14).

#### Port Forwarding For HTTPS(443) Service of WebServer1

You may follow the instructions below to add a port forwarding rule for HTTPS service of WebServer1.

1. 
Navigate to `Firewall` →`NAT` →`Port Forward` in your OPNsense Web UI.
2. 
Click the `+` button in the upper right corner. This will open the port forwarding configuration window.**Figure 5.***Port forwarding rule configuration for HTTPS in OPNsense-1*
3. 
Set the Interface to `WAN` .
4. 
Set the TCP/IP Version to `IPv4` .
5. 
Set the Protocol to `TCP` .
6. 
Set the Destination to `WAN Address` .
7. 
Set the Destination Port Range to `HTTPS` .
8. 
Select `Single Host or Network` from the Redirect Target IP dropdown menu. Then, set the field to the private IP address of the WebServer1, such as`10.10.10.13` .
9. 
Set the Redirect Target Port to `HTTPS` .**Figure 6.***Port forwarding rule configuration for HTTPS in OPNsense-2*
10. 
You may enable logging by clicking the check box in the Log option.
11. 
Fill in the Description field, such as `Allow HTTPS access to Webserver_10.10.10.13` .
12. 
Select `Add associated filter rule` from the Filter rule association option.
13. 
Leave other options as default.
14. 
Click `Save` button at the bottom of the page.

**Figure 7.** *Port forwarding rule configuration for HTTPS in OPNsense-3*

#### Port Forwarding For HTTP(80) Service of WebServer1

To create a port forwarding rule for the HTTP(80) service of the WebServer1, you may clone the port forwarding rule for the HTTPS(443) service created above and change the related settings by following the step given below.

**Figure 8.** *Port forwarding rules list in OPNsense*

