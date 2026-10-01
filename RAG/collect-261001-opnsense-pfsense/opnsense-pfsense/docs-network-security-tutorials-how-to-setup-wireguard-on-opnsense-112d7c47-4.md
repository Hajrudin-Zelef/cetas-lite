---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/docs-network-security-tutorials-how-to-setup-wireguard-on-opnsense-112d7c47-4
title: "docs-network-security-tutorials-how-to-setup-wireguard-on-opnsense-112d7c47"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/docs-network-security-tutorials-how-to-setup-wireguard-on-opnsense-112d7c47.md
source_anchor: ""
source_lines: [317, 433]
sha256: fdf8e36a9a72f2c76810a58e6143956ca06e2931e5270a754662ac16ad7c9b1d
---

# docs-network-security-tutorials-how-to-setup-wireguard-on-opnsense-112d7c47

Figure 20. Adding VPN client endpoints to the WireGuard server configuration
Step 8. Enabling WireGuard Server on OPNsense
Once you have added the necessary peers (VPN clients) to the WireGuard configuration, the next step is to enable the WireGuard service so that the server can start listening for client connections.
Follow the steps below to enable the WireGuard server.
- 
Navigate to the VPN → WireGuard → Interfaces tab on the OPNsense Web UI.
- 
Check the Enable WireGuard checkbox to activate the service.
- 
Click Apply to save and apply the changes.
This step starts the WireGuard service on your OPNsense firewall and prepares it to establish secure VPN tunnels with authorized clients.
If you’ve made changes to the instance or peer configurations, it’s a good idea to disable and re-enable the WireGuard service from this screen to ensure the changes take effect.
Step 9. Creating WireGuard Interface on OPNsense
Although not strictly required, it is recommended to create a dedicated WireGuard interface on OPNsense for better control and flexibility. This step allows you to apply specific firewall rules, simplifies alias management, and ensures proper NAT behavior for tunnel traffic.
If you only need to access the local network (LAN) via WireGuard and not public IP addresses (e.g., internet access), you may skip this step.
Although creating a WireGuard interface is not strictly required for a road warrior setup, it is strongly recommended due to several practical advantages:
- It automatically creates an alias for the tunnel subnet(s), simplifying the application of firewall rules. Otherwise, you'd have to define them manually.
- It adds an automatic outbound NAT rule for IPv4, enabling access to public IP addresses outside the local network without manual configuration.
- It allows each WireGuard instance (e.g., wg0, wg1, etc.) to have its own firewall rule set, ensuring clean separation and easier management.
To create and configure a WireGuard interface, follow the steps below.
- 
Navigate to Interfaces → Assignments on the OPNsense Web UI.
- 
Under Assign a new interface, find your WireGuard device in the dropdown, for example: wg0 (WireGuard - MyWireGuard).
- 
Click the Add button.
- 
The interface will be added to the list above with a default name like opt1. Figure 21. Creating WireGuard Interface on OPNsense
- 
Click on the newly added interface name (e.g., [WireGuardVPN]) to open its configuration page.
- 
Check the Enable interface box.
- 
Provide a Description, such as MyWireGuardVPN.
- 
In the Lock section, enable the checkbox Prevent interface removal.
- 
Leave other settings as default.
- 
Click Save.
- 
Click Apply changes.
Figure 22. Enabling WireGuard Interface on OPNsense
You do not need to manually assign an IP address to the WireGuard interface. Once the WireGuard service is restarted, the IP address specified during the initial WireGuard configuration will be automatically applied to the interface. To ensure this, restart the WireGuard service with the following steps.
- Go to VPN → WireGuard → Instances.
- Uncheck the Enable WireGuard checkbox and click Apply.
- Then check the Enable WireGuard checkbox again and click Apply.
This ensures that the interface receives the correct tunnel address and is ready for use before configuring firewall rules or testing the VPN tunnel.
Step 10. Creating Firewall Rules
To ensure secure and functional communication between WireGuard VPN clients and the OPNsense network, you need to define two firewall rules that are shared below.
- 
A rule on the WAN interface to allow incoming connections to the WireGuard VPN server from remote clients.
- 
A rule on the WireGuard interface (e.g., MyWireGuard) to allow VPN clients to access internal network resources (e.g., LAN devices or the Internet), depending on your specific needs.
10.1. Allowing VPN clients to access the OPNsense WireGuard Server
You have installed and configured a WireGuard VPN server to provide secure remote access to your internal network.
To make your WireGuard server accessible from the Internet, you must define a firewall rule on the WAN interface that allows incoming connections to the WireGuard port (e.g., 51820/UDP).
Once this firewall rule is defined, your VPN clients will be able to connect to the WireGuard server and access internal or external resources through the OPNsense gateway.
For detailed guidance on how to define this rule, you can refer to the How to Configure OPNsense Firewall Rules article by Sunny Valley Networks.
10.2. Allowing VPN clients to access the internal networks
If you want your WireGuard VPN clients to access internal network resources (e.g., LAN devices, printers, file servers), you need to define a second firewall rule on the WireGuard interface you previously created (e.g., MyWireGuard).
This rule allows traffic from the connected VPN clients to reach the rest of the internal network or any destination, depending on how broadly you configure it.
To define the rule follow these steps below.
- 
Navigate to Firewall → Rules, and select the WireGuard interface (e.g., MyWireGuard) from the tabs. Figure 23. Creating Firewall Rule on the WireGuardVPN Interface
- 
Click the ➕ (Add) button to create a new rule with the following settings. Option Value Action Pass Interface WireGuardVPN Direction in TCP/IP Version IPv4 Protocol any Source WireGuardVPN net Source Port any Destination any Destination Port any Description Allow WireGuard clients access to all destinations
- 
Click Save at the bottom of the page.
- 
Click Apply Changes to activate the rule.
Figure 24. Firewall Rule Settings on OPNsense
Step 11. Verifying the WireGuard Setup on OPNsense
After completing both the WireGuard server and client configurations, you can verify the VPN setup using the following steps.
11.1. Activating WireGuard on Windows Client:
To establish a VPN connection from your Windows PC follow these steps.
- 
Open the WireGuard application.
- 
Click on the tunnel named MyWireGuard and press the Activate button.
- 
Once connected, the tunnel Status will change to Active.
Figure 25. Activating the WireGuard tunnel on Windows client
11.2. Activating WireGuard on Android Client:
To connect your Android device to the VPN:
- 
Open the WireGuard app.
- 
Toggle the MyWireGuard tunnel to ON.
11.3. Monitoring VPN Clients on OPNsense:
To view active connections navigate to VPN → WireGuard → Status on OPNsense. You should see each connected client with the following details.
- Public key of the peer
- Time since last handshake
- Amount of data transferred
Figure 26. List of connected WireGuard clients in OPNsense
11.4. Performing a Ping Test:
From the client device (Windows or Android), test connectivity to the WireGuard server by running the following command.
ping 10.0.0.1
If successful, it confirms tunnel functionality and routing.
11.5. Verifying Public IP Address:
To confirm that your traffic is routed through the VPN, open https://www.whatismyip.com from the client.
If the tunnel is working, it should show the VPN server’s public IP, not your local ISP IP.
11.6. Running a Traceroute Test:
You can trace the route to an external IP to ensure traffic goes through the VPN.
tracert 8.8.8.8
Expected output example is shown below. The presence of 10.0.0.1 as the first hop indicates successful tunnel routing.
1- 10.0.0.1 ← WireGuard Server IP
2- 192.168.0.1 ← LAN Gateway
3- ...
11.7. Internal Network Access Test:
If the firewall rules permit, clients should have full access to internal LAN devices.
- 
Try pinging an internal device (e.g., 192.168.1.100 ) from the client.
- 
Then ping the client (e.g., 10.0.0.12 ) from within the LAN.
If both directions work, LAN access via VPN is verified.
How to Resolve os-wireguard (missing) Issue?
