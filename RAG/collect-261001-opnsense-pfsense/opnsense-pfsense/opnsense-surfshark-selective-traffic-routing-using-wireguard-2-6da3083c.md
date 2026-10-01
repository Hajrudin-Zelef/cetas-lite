---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/opnsense-surfshark-selective-traffic-routing-using-wireguard-2-6da3083c
title: "Use this configuration with WireGuard client"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/opnsense-surfshark-selective-traffic-routing-using-wireguard-2-6da3083c.md
source_anchor: ""
source_lines: [1, 70]
sha256: 0f7c6b42711427488b9735ab7464d13b1aa78bc168e4b33c4bc4f0b393fdba2e
---

# Use this configuration with WireGuard client

WireGuard on OPNsense -> Surfshark - Selective traffic routing
I was trying to route traffic from selected clients over the Surfshark WireGuard tunnel using the official documentation here and ran into issues. Here is how I finally managed to set it up with help from Reddit threads [1] and YouTube [1].
Tested on OPNsense 24.1.10_3-amd64.
Generate Surfshark configuration
- Login to Surfshark and navigate to VPN -> Manual Setup -> Router -> WireGuard
- Under Credentials, you can either import your public key (preferred) or have Surfshark generate a key pair for you.
Key pairs can be generated using multiple ways. One easy option is to use the Windows client.
- Once keys are generated, click 'I have a key pair'. Make sure to save the Private key securely which will be required later
- Give it a name and paste the Public key that was generated and save it
- Now choose a location for the VPN server and download the config.
- Open the config file which looks like the below:
#
# Use this configuration with WireGuard client
#
[Interface]
Address = 10.14.0.2/16
PrivateKey = <redacted>
DNS = 162.252.172.57, 149.154.159.92
[Peer]
PublicKey = <redacted>
AllowedIPs = 0.0.0.0/0
Endpoint = uk-lon.prod.surfshark.com:51820
Configure OPNsense
- Login to OPNsense GUI and navigate to VPN -> WireGuard -> Instances -> Add new (+ sign)
Fill in the details from the [Interface] section of the config file that was downloaded. Private key is the one you generated earlier.
Set a Listen port that is not in use by any other WireGuard Instance you may have.
Tick 'Disable routes' since we are using selective routing of traffic
- Click Save
- Go to Peers tab and add a new peer
Fill the details from Peers section of the config file. Select the Instance that was created.
- Click Save and then Apply.
Add Interface
- Go to Interfaces -> Assignments
- Select the new Instance that was created and click Add.
- Click the Interface that was created and Enable it
- Set the MTU and MSS values to 1420 (!Important!)
Save and Apply changes
Configure Gateway
- Go to System -> Gateways > Configuration > Add new Gateway
Interface - created above
IP address - This will be the Interface IP address from Surfshark config file (10.14.0.2), off by 1 i.e. 10.14.0.1
Far Gateway - selected
- Save and Apply
Configure Firewall Outbound NAT
- Go to Firewall -> NAT -> Outbound. Ensure it is set to 'Hybrid'.
- Add a new entry
Change Interface to the one created.
- Save, Apply changes
Find the MAC address of the client which needs to be routed through VPN
- In Windows, use "ipconfig/all" command
- In Linux, use "ip a" command
Create an alias for the MAC address in OPNsense
- Go to Firewall -> Aliases -> Add
- Save and Apply
Create firewall rule to route traffic
- Go to Firewall -> Rules -> LAN -> Add rule
Set Source to the Alias added in above step
Set Gateway to the one that was created earlier. (You may ignore the Gw IP in screenshot)
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
- Save and Apply changes
