---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/docs-network-security-tutorials-how-to-configure-opnsense-nat-4928d4ad-5
title: "docs-network-security-tutorials-how-to-configure-opnsense-nat-4928d4ad"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/docs-network-security-tutorials-how-to-configure-opnsense-nat-4928d4ad.md
source_anchor: ""
source_lines: [347, 460]
sha256: 92a592fbf8b9ea52e49a806a4a87d9f9dce738a490289b5d0e6af27b2ff02585
---

# docs-network-security-tutorials-how-to-configure-opnsense-nat-4928d4ad

1. 
Click the clone icon to copy the port forwarding rule for the SSH(2222) service created above in the port forward rules list.
2. 
Change the Destination Port Range option to `5555` .
3. 
Set the Redirect Target IP to WebServer2 local IP address, such as 10.10.10.14.
4. 
Set the Redirect Target Port to `MS RDP` .
5. 
Change the Description field to `Allow RDP access to Webserver_10.10.10.14` .**Figure 26.***Port forwarding rule configuration for MS-RDP(5555) in OPNsense-1*
6. 
Verify that the Filter rule association option is set to `Add associated filter rule`
7. 
Leave other options as they are.
8. 
Click `Save` button at the bottom of the page.**Figure 27.***Port forwarding rule configuration for SSH(2222) in OPNsense-3*Now, you have completed the port forwarding rule configurations of both management services. Your port forwarding rules list should look like this. **Figure 28.***Port forwarding rules list for web servers in OPNsense*
9. 
Click `Apply Changes` at the upper right of the page to activate the settings.Since we have selected the `Add associated filter rule` option, the related firewall rules are created on the WAN interface automatically. To view the automatically added associated rules, navigate to the`Firewall` →`Rules` →`WAN` . Firewall rules list on WAN interfaces should look like this:**Figure 29.***WAN firewall rules for SSH and RDP access in OPNsense*

### Outbound NAT For Accessing a Remote Service Via Specific External IP Address

Assume that one of your application servers (WebServer1 with the IP address 10.10.10.13) needs to connect to a MySQL database on another company network via the Internet. However, in accordance with the agreements between your company and the other company, you must ensure that the remote MySQL DB server(public IP address: 3.3.3.3) is only accessible by WebServer1 and that no other devices in your LAN can access the remote DB.

To accomplish this, firstly you need a second public IP address which will be used for providing WebServer1 access to the remote MySQL database. Because, your first public IP address is being used for Internet access of the local users and servers. We will use the `2.2.2.2` as our second IP address and WebServer1 will connect to the remote MySQL database with this external IP address.

| Packet Type | Source IP Before NAT | Destination IP Before NAT | Source IP After NAT | Destination IP After NAT | 
|---|---|---|---|---|
| MySQL Request | 10.10.10.13 | 3.3.3.3 | 2.2.2.2 | 3.3.3.3 | 
| MySQL Reply | 3.3.3.3 | 2.2.2.2 | 3.3.3.3 | 10.10.10.13 | 

**Figure 30.** *Outbound NAT/SNAT topology for accessing remote Database server*

You may follow the next steps given below:

1. 
Define an alias, such as `RemoteCompany_DB` . For more information about creating an alias, please refer to How to Configure OPNsense Firewall article.
2. 
To create a Virtual IP address for your second public IP address, navigate to the `Interfaces` →`Virtual IPs` →`Settings` .
3. 
Click the `+` icon to add Virtual IP address.**Figure 31.***Adding Virtual IP address in OPNsense*
4. 
Select `IP Alias` as Mode.
5. 
Select `WAN` as Interface.
6. 
Set Address to your second public IP address which is used for accessing the database server by your WebServer1, such as `2.2.2.2/32`
7. 
Enter `WAN VIP_2.2.2.2` in the Description field.
8. 
Leave other options as default.
9. 
Click `Save` .**Figure 32.***Setting Virtual IP address configuration in OPNsense*
10. 
Click `Apply Changes` to activate the VIPs settings.**Figure 33.***Virtual IP address settings in OPNsense*
11. 
Navigate to the `Firewall` →`NAT` →`Outbound` to define Outbound NAT.
12. 
Select `Hybrid outbound NAT rule generation` option.
13. 
Click `Save` button.**Figure 34.***Setting Outbound NAT mode in OPNsense*
14. 
Click `+` icon to add a manual Outbound NAT rule.
15. 
Set Interface to `WAN` .
16. 
Set TCP/IP Version to `IPv4` .
17. 
Set Protocol `TCP` .
18. 
Set Source add to `Single Host or Network`
19. 
Enter the WebServer1 IP address such as `10.10.10.13/32` .
20. 
Set Source Port to `any` .**Figure 35.***Defining Outbound NAT rule in OPNsense -1*
21. 
Select Destination Address as `RemoteCompany_DB` .
22. 
Select Destination Port as `MySQL` .
23. 
Select `2.2.2.2 (WAN IP_2.2.2.2)` for Translation / target
24. 
Enable Logging. **Figure 36.***Defining Outbound NAT rule in OPNsense -2*
25. 
Enter `Remote MySQL DB access` in Description field.
26. 
Click `Save`
27. 
Click `Apply Changes` to activate the Outbound NAT rule.

Your Outbound NAT rules list should look something like this:

**Figure 37.** *Manual Outbound NAT rules in OPNsense*

When WebServer1 tries to connect to a remote database server, you should see that it connects the DB using `2.2.2.2` IP address in your firewall logs. To view the firewall logs navigate to `Firewall` → `Log Files` → `Live View`. Your logs look like this.

**Figure 38.** *Firewall Live Log View in OPNsense*

## What is the Importance of the OPNSense NAT Configuration?

The configuration of NAT (Network Address Translation) in OPNsense is of the utmost importance, as it regulates the communication between your internal network and external networks (such as the Internet). Secure, reliable, and efficient network communication is guaranteed by the proper configuration of NAT.

The primary explanations for the significance of NAT configuration in OPNsense are given below.

- **Enabling Internet Access for Internal Networks** :  NAT enables multiple devices within your internal network to share a single public IP address that is assigned by your Internet Service Provider (ISP). Each device would necessitate its own public IP address, which is both impractical and expensive in the absence of NAT.
- **Service Accessibility and Port Forwarding** : NAT allows for the establishment of Port Forwarding protocols, which enables external users to securely access services hosted internally (e.g., web servers, VPN servers, gaming servers). The proper configuration of NAT guarantees that only authorized traffic reaches your internal servers.
- **Privacy and Security** : NAT inherently offers a layer of security by concealing internal IP addresses from external networks. Your network is rendered less susceptible to external threats by the fact that attackers are unable to directly observe or access your internal devices.
- **Streamlined Network Administration** : NAT facilitates network management by reducing the intricacy of IP addressing schemes within your internal network. It simplifies the process of administering and maintaining your network infrastructure.
- **Failover and Load Balancing** : OPNsense's NAT feature can be configured to distribute network traffic across multiple WAN connections, thereby enhancing reliability and redundancy. It enables automatic failover, which guarantees uninterrupted internet connectivity in the event of a single connection failure.
- **Policy Enforcement and Compliance** : The NAT configuration enables the enforcement of network policies, the restriction of network access, and the fulfillment of regulatory or organizational requirements.  Security and compliance can be improved by establishing policies that authorize or prohibit particular traffic.
- **Network Traffic Management** : NAT is responsible for the efficient management and routing of network traffic, ensuring that data messages reach their intended destinations. It enables the configuration of specific behaviors for inbound and outbound traffic, thereby enhancing the reliability and performance of the network.
- **Monitoring and Logging** : OPNsense's NAT feature offers network traffic visibility through archiving and monitoring capabilities.  You can effectively troubleshoot network issues, identify potential security hazards, and monitor and audit network activity.

## How does Port Forwarding Work in OPNSense?

