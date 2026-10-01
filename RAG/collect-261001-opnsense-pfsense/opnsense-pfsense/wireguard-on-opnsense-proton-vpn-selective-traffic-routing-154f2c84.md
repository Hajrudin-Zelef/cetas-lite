---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/wireguard-on-opnsense-proton-vpn-selective-traffic-routing-154f2c84
title: "Key for opnsense-wg-GB"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["accelerator"]
source: docs/RAG/collect-261001-opnsense-pfsense/wireguard-on-opnsense-proton-vpn-selective-traffic-routing-154f2c84.md
source_anchor: ""
source_lines: [1, 69]
sha256: 87e98fa2812aed0b8d9864d83c3282014a2d44eab2e4d9246ccba7c80b397b68
---

# Key for opnsense-wg-GB

WireGuard on OPNsense -> Proton VPN- Selective traffic routing
This is a replica of the Surfshark VPN post from earlier, adapted for Proton VPN.
Tested on OPNsense 24.7.10_2-amd64.
Generate Proton VPN Wireguard configuration
Login to Proton VPN and navigate to Downloads
Select the options that you need
Scroll down to the server location you need and pick a server with low load. Sadly, there is no option to select the best server based on load here. I'm selecting a UK server here.
Hit Create button and then download the config to local machine which we will use to set up the VPN connection.
[Interface]
# Key for opnsense-wg-GB
# Bouncing = 20
# NetShield = 2
# Moderate NAT = off
# NAT-PMP (Port Forwarding) = off
# VPN Accelerator = on
PrivateKey = <redacted>
Address = 10.2.0.2/32
DNS = 10.2.0.1
[Peer]
# UK#225
PublicKey = <redacted>
AllowedIPs = 0.0.0.0/0
Endpoint = <redacted>
Configure OPNsense
- Login to OPNsense GUI and navigate to VPN -> WireGuard -> Instances -> Add new (+ sign)
Fill in the details from the [Interface] section of the config file that was downloaded. Private key is the one you generated earlier.
Set a Listen port that is not in use by any other WireGuard Instance you may have.
Tick 'Disable routes' since we are using selective routing of traffic
- Click Save
- Go to Peers tab and add a new peer
Fill the details from Peers section of the config file. Select the Instance that was created. Make sure to fill the Public key, endpoint details.
Click Save and then Apply.
Add Interface
- Go to Interfaces -> Assignments
- Select the new Instance that was created and click Add.
- Click the Interface that was created and Enable it
Save and Apply changes
Configure Gateway
- Go to System -> Gateways > Configuration > Add new Gateway
Interface - created above
IP address - Use an IP address that is not the gateway or DNS. I'm using 10.2.0.4 here.
Far Gateway - selected
Save and Apply
Configure Firewall Outbound NAT
- Go to Firewall -> NAT -> Outbound. Ensure it is set to 'Hybrid'.
- Add a new entry
Change Interface to the one created.
Find the MAC address of the client which needs to be routed through VPN
- In Windows, use "ipconfig/all" command
- In Linux, use "ip a" command
Create an alias for the MAC address in OPNsense
- Go to Firewall -> Aliases -> Add
Save and Apply
Create firewall rule to route traffic
- Go to Firewall -> Rules -> LAN -> Add rule
Set Source to the Alias added in above step
Set Gateway to the one that was created earlier.
- Save and Apply changes
Rules are processed in order, so ensure that the newly added rule is hit before any other.
- Reboot OPNsense
- Confirm that the traffic from the client is now routed over VPN!
Add a killswitch
If the VPN interface goes down for some reason, the killswitch will ensure that the client traffic will not exit via the default WAN interface.
- Go to Firewall -> Rules -> LAN -> Edit the rule created to route traffic
Set local tag t0 'NO_WAN_EGRESS'
- Save and Apply changes
- Go to Firewall -> Rules -> Floating -> Add new
Show 'Advanced features'
Save and Apply changes
