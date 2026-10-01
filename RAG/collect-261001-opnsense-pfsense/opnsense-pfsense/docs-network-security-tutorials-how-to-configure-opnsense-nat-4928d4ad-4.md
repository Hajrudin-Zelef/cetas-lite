---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/docs-network-security-tutorials-how-to-configure-opnsense-nat-4928d4ad-4
title: "docs-network-security-tutorials-how-to-configure-opnsense-nat-4928d4ad"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/docs-network-security-tutorials-how-to-configure-opnsense-nat-4928d4ad.md
source_anchor: ""
source_lines: [242, 346]
sha256: a0af823d4246c04d75f12b5d07ddc712eec5a67d168e82c6f0c283c0a065ac0f
---

# docs-network-security-tutorials-how-to-configure-opnsense-nat-4928d4ad

1. 
Click the clone icon to copy the port forwarding rule for the HTTPS(443) service created above.
2. 
Change the Destination Port Range option to `HTTP` .**Figure 9.***Port forwarding rule configuration for HTTP in OPNsense-1*
3. 
Set the Redirect Target Port to `HTTP` .
4. 
Change the Description field to `Allow HTTP access to Webserver_10.10.10.13` .
5. 
Verify that the Filter rule association option is set to `Add associated filter rule`
6. 
Leave other options as they are.
7. 
Click `Save` button at the bottom of the page.**Figure 10.***Port forwarding rule configuration for HTTP in OPNsense-2*

#### Port Forwarding For HTTP Service of WebServer2 on Custom External Port(81)

To create a port forwarding rule for the HTTP service of the WebServer2 on custom port(81), you may clone the port forwarding rule for the HTTP(80) service created above and change the related settings by following the step given below.

**Figure 11.** *Port forwarding rules list in OPNsense*

1. 
Click the clone icon to copy the port forwarding rule for the HTTP(80) service created above.
2. 
Change the Destination Port Range option to `other` and enter`81` to the related field.**Figure 12.***Port forwarding rule configuration for HTTP(81) in OPNsense-1*
3. 
Set the Redirect Target IP to `10.10.10.14`
4. 
Set the Redirect Target Port to `HTTP` .
5. 
Change the Description field to `Allow HTTP access to Webserver_10.10.10.14` .
6. 
Verify that the Filter rule association option is set to `Add associated filter rule`
7. 
Leave other options as they are.
8. 
Click `Save` button at the bottom of the page.**Figure 13.***Port forwarding rule configuration for HTTP(81) in OPNsense-2*

#### Port Forwarding For HTTPS Service of WebServer2 on Custom External Port (8443)

To create a port forwarding rule for the HTTPS service of the WebServer2 on a custom external port(8443), you may clone the port forwarding rule for the HTTP(81) service created above and change the related settings by following the step given below.

**Figure 14.** *Port forwarding rules list in OPNsense*

1. 
Click the clone icon to copy the port forwarding rule for the HTTP(81) service created above.
2. 
Change the Destination Port Range option to `8443` .**Figure 15.***Port forwarding rule configuration for HTTP(8443) in OPNsense-1*
3. 
Set the Redirect Target Port to `HTTPS` .
4. 
Change the Description field to `Allow HTTPS access to Webserver_10.10.10.14` .
5. 
Verify that the Filter rule association option is set to `Add associated filter rule`
6. 
Leave other options as they are.
7. 
Click `Save` button at the bottom of the page.**Figure 16.***Port forwarding rule configuration for HTTP(8443) in OPNsense-2*Now, you have completed the port forwarding configurations of both web servers. Your port forwarding rules list should look like this. **Figure 17.***Port forwarding rules list for web servers in OPNsense*
8. 
Click `Apply Changes` at the upper right of the page to activate the settings.Since we have selected the `Add associated filter rule` option, the related firewall rules are created on the WAN interface automatically. To view the automatically added associated rules, navigate to the`Firewall` →`Rules` →`WAN` . Firewall rules list on WAN interfaces should look like this:**Figure 18.***WAN firewall rules for web server port forwarding in OPNsense*Although internal users should access the web servers by connecting to the private IP address (local IP) of the servers, they may try to connect to a local server by using the public IP addresses. To allow local users to access the public IP addresses of these servers, you must allow the NAT reflection. For NAT reflection, first you should enable the NAT reflection by checking on the `Reflection for port forwards` option on the`Firewall` →`Settings` →`Advanced` page.**Figure 19.***Enabling Reflection for port forwards*Then, you should select the interface where the local users are, such as LAN, as well as the WAN interface during the port forwarding rule configuration. **Figure 20.***NAT reflection*Ensure that NAT reflection is enabled in the port forwarding rule configuration. **Figure 21.***NAT reflection is enabled in port forwarding rule*

### How to Configure Port Forwarding For SSH and RDP Services on Custom Ports

Assume that a web administrator needs remote(SSH & RDP) access to the web servers from his home. He is using a static public IP address at home. Since management services such as SSH and RDP are critical and pose a high security risk, it is recommended that they are not accessible from the entire Internet. As a result, you will create a port forwarding rule to allow the web administrator's IP address to connect to the web servers. Also, because the default ports are already in use for accessing other servers, you must enable SSH and RDP services on custom ports.

| Server Name | External IP | External Port | Local IP | Local Port | Client IP | 
|---|---|---|---|---|---|
| WebServer1 | Public Internet IP | 2222 | 10.10.10.13 | 22 | 1.1.1.1 | 
| WebServer2 | Public Internet IP | 5555 | 10.10.10.14 | 3389 | 1.1.1.1 | 

**Figure 22.** *Port Forwarding topology for SSH and RDP services*

After completing the port forwarding configurations in your OPNsense firewall, port 2222 requests coming from web administrator IP address(1.1.1.1) to your WAN IP will be redirected to the WebServer1(10.10.10.13), while port 5555 requests coming from web administrator IP address(1.1.1.1) to your WAN IP will be redirected to the WebServer2(10.10.10.14).

#### Port Forwarding For SSH Service of WebServer1 on Custom External Port(2222)

To create a port forwarding rule for the SSH service of the WebServer1 on custom port(2222), you may clone the port forwarding rule for the HTTP(80) service created above and change the related settings by following the step given below.

1. 
Click the clone icon to copy the port forwarding rule for the HTTP(80) service created above.
2. 
Click the `Advanced` button in the`Source` option. This will displays the details of the Source option.
3. 
Select `Single Host or Network` from the Source dropdown menu.
4. 
Enter the Web Administrator's static public IP address, such as 1.1.1.1/32.
5. 
Leave Source Port Range as `any` .**Figure 23.***Port forwarding rule configuration for SSH(2222) in OPNsense-1*
6. 
Change the Destination Port Range option to `2222` .
7. 
Set the Redirect Target Port to `SSH` .
8. 
Change the Description field to `Allow SSH access to Webserver_10.10.10.13` .**Figure 24.***Port forwarding rule configuration for SSH(2222) in OPNsense-2*
9. 
Verify that the Filter rule association option is set to `Add associated filter rule`
10. 
Leave other options as they are.
11. 
Click `Save` button at the bottom of the page.**Figure 25.***Port forwarding rule configuration for SSH(2222) in OPNsense-3*

#### Port Forwarding For RDP Service of WebServer2 on Custom External Port(5555)

To create a port forwarding rule for the RDP service of the WebServer2 on custom port(5555), you may clone the port forwarding rule for the SSH(2222) service created above and change the related settings by following the step given below.

