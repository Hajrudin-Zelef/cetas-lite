---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/docs-network-security-tutorials-how-to-setup-wireguard-on-opnsense-112d7c47-2
title: "docs-network-security-tutorials-how-to-setup-wireguard-on-opnsense-112d7c47"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-opnsense-pfsense/docs-network-security-tutorials-how-to-setup-wireguard-on-opnsense-112d7c47.md
source_anchor: ""
source_lines: [71, 179]
sha256: ec83b02390ccd4ad0ff5e05b4ecae8fe78ec38c79e0cfa16d8b42bfdd6c5fe43
---

# docs-network-security-tutorials-how-to-setup-wireguard-on-opnsense-112d7c47

Endpoint Address and Port (optional): Use this if you're adding a static client (e.g., remote Windows PC).
- 
- 
Leave other settings as default unless specified.
- 
Click Save.
4. Assign Interface to WireGuard:
To ensure the WireGuard traffic can be properly routed and managed, you need to assign it to an interface in OPNsense with the following steps.
- 
Go to Interfaces → Assignments .
- 
From the drop-down menu, select your WireGuard instance (e.g., wg0).
- 
Click Add, then click on the newly created interface (e.g., OPT1).
- 
Enable the interface, rename it to something like WireGuardVPN, and click Save.
5. Configure Firewall Rules:
Firewall rules must be configured to allow incoming and outgoing traffic through the WireGuard interface with the following steps.
- 
Go to Firewall → Rules → WireGuardVPN (or the interface name you assigned).
- 
Click + Add to create a new rule.
- 
Set Action to Pass, Protocol to any, Source and Destination to any.
- 
Click Save and then Apply Changes.
6. Test the VPN Connection:
After all configurations are complete, it's time to test if the VPN tunnel works correctly with the following steps.
- 
On your client device, activate the WireGuard connection.
- 
On the OPNsense dashboard, go to VPN → WireGuard → Status to check if a handshake is established.
- 
You may also view logs under System → Log Files → General orFirewall Logs to verify successful traffic.
How to Configure the WireGuard VPN Server in OPNsense?
WireGuard is a modern, high-performance VPN protocol that offers superior speed, security, and simplicity compared to traditional VPN solutions like OpenVPN and IPsec. OPNsense includes built-in support for WireGuard, allowing users to set up a secure VPN server with minimal effort. In this section, we will walk through the complete process of configuring the WireGuard VPN server in OPNsense from verifying its availability to adding peers and applying firewall rules. Whether you are connecting from a Windows PC, mobile device, or remote office, this guide will help you build a reliable VPN infrastructure using WireGuard on OPNsense.
Step 1. Verify WireGuard Availability in OPNsense
As of OPNsense version 24.1 and later, WireGuard is included as a core componentof the system. This means you no longer need to manually install the os-wireguard plugin. To check if WireGuard is available follow these steps below.
- 
Log in to your OPNsense web interface.
- 
Navigate to VPN → WireGuard from the main menu.
- 
If the WireGuard configuration page opens successfully, it is already installed and ready to use. Figure 1. WireGuard Service on OPNsense tipIf you don’t see the WireGuard menu, make sure your OPNsense system is up to date by going to System → Firmware → Status , and clicking Check for Updates.
If you're using OPNsense 21.x or 22.x, you will need to install the WireGuard plugin manually with the following steps.
- 
Go to System → Firmware → Plugins .
- 
Search for os-wireguard .
- 
Click the + icon to install the plugin.
- 
After installation, refresh the page and access WireGuard from VPN → WireGuard .
Consider upgrading to the latest OPNsense version to benefit from built-in WireGuard support and better security.
Step 2: WireGuard VPN Server Configuration on OPNsense
After confirming that WireGuard is available on your OPNsense system (version 24.1 or later), you can now configure the WireGuard server. This includes setting the server’s IP address, listening port, and generating cryptographic keys for secure communication.
Follow the steps below to create your WireGuard server instance.
2.1. Enable WireGuard Service:
Before creating your VPN server instance, you must enable the WireGuard service. To do this, log in to the OPNsense web interface and follow these steps.
- 
Go to VPN → WireGuard → Instances .
- 
At the bottom of the page, check the Enable WireGuard option.
- 
Click Apply to activate the backend service.
This step starts the WireGuard backend service that is required for creating and managing VPN tunnels.
Figure 2. Enabling WireGuard Service on OPNsense
2.2. Create and Configure WireGuard VPN Server Instance on OPNsense:
Once WireGuard is enabled on your OPNsense firewall, the next step is to create and configure your WireGuard server instance. This instance will define how the VPN server listens for connections, what IP addresses it uses, and which cryptographic keys are assigned to it.
Follow the steps below to complete the server setup.
- 
Click the + icon in the Instances tab to open the configuration form. This instance will define the server’s listening port, tunnel address, and cryptographic identity. Figure 3. Adding a New WireGuard Instance
- 
After clicking the "+" button, a configuration window titled Edit instance will appear. Here, you will define the basic parameters for your WireGuard server. Fill in the following fields. Field Description Enabled Leave checked to activate this instance Name Any descriptive label (e.g., MyWireGuard) Public Key / Private Key Click the gear icon to auto-generate keys Listen Port Default is 51820, or set a custom unused UDP port Tunnel Address Example: 10.0.0.1/24 (used for VPN IP assignments) Peers Leave empty for now – will be configured later Disable Routes Leave unchecked unless needed for advanced routing infoThe tunnel address should be unique and reserved only for VPN communication. For example, 10.0.0.1/24 is common for small-scale setups. tipIf you plan to connect multiple clients, ensure that the subnet (/24) has enough room to assign IP addresses to all peers. Figure 4. Filling out the Edit Instance form in WireGuard
- 
After filling out all the necessary fields, click Save to create the instance. The server is now configured and ready to accept incoming WireGuard connections (once peers are added).
2.3. View and Copy the Server’s Public Key:
To complete the server-side setup, you need to copy the server’s public key, which will later be required when setting up clients. Follow these steps:
- 
Go to VPN → WireGuard → Instances .
- 
Click the Edit icon next to the instance you just created.
- 
Find the Public Key field and copy the value.
Never share the private key. Only the public key should be exchanged with client peers.
Figure 5. Viewing and copying the server’s public key
2.4. Close the Edit Window:
Once you’ve noted the public key, click Cancel or close the window. Your WireGuard server is now fully configured and listening on the defined port (e.g., 51820/UDP). You’re ready to proceed with peer (client) configuration in the next step.
This completes the initial server-side setup.
Step 3. WireGuard VPN Client Setup on Windows
In this section, we’ll walk through how to install and configure the WireGuard client on a Windows system. This setup will allow your Windows PC to securely connect to your OPNsense WireGuard VPN server using encrypted tunnels.
You may follow the steps below to install and configure WireGuard as a VPN client on a Windows platform.
3.1. Download and Install Windows WireGuard Client
Download and install the Windows installer from the WireGuard website. This option chooses the most recent version for your hardware, downloads it, and installs it.
- 
Go to the official WireGuard website: https://www.wireguard.com/install
- 
Under Windows, click Download Windows Installer
- 
Run the installer and follow the on-screen instructions
Figure 6. Downloading WireGuard Windows Installer
After the installation, you should see the WireGuard icon in the notification area on the taskbar.
Figure 7. WireGuard icon on taskbar
3.2. Configuring WireGuard Windows Client
- 
Launch the WireGuard application. In the Tunnels tab, click the down arrow next to the Add Tunnel button. Figure 8. Configuring WireGuard on Windows Client
- 
Choose Add empty tunnel. Figure 9. Adding empty tunnel
- 
