---
id: collect-261001-general-networking/general-networking/morningreis-eeda36e8bb07dcb750d77e9a744776e8-067af069-1
title: "Key for OPNsense Demo"
domain: general-networking
role: reference
task: reference
actors: ["United States"]
dates: ["2022-12"]
keywords: ["accelerator"]
source: docs/RAG/collect-261001-general-networking/morningreis-eeda36e8bb07dcb750d77e9a744776e8-067af069.md
source_anchor: ""
source_lines: [1, 142]
sha256: 61453496f755f6d304b34e52c6a5f5538902efd3033a10d2407299d7e38d14e5
---

# Key for OPNsense Demo

Published: 16 December 2022
Reference: https://docs.opnsense.org/manual/how-tos/wireguard-selective-routing.html
Goal: Set up one or more Wireguard connections from ProtonVPN on OPNsense, with policy based routing, and optional Killswitch.
I'm writing this guide first as a reference for my future self for when I inevitably forget how to do this, but also to help others out. I found there were not many guides on this specific configuration, particularly not with multiple concurrent connections, and these were some steps which were not at all obvious. I did begin with the guide in the official OPNsense documentation, but even that was missing info to make ProtonVPN work. If you are a pfSense user, it is very similar to OPNsense, and you should be able to follow along with some success, but I have not tested it myself.
- Log into https://account.protonvpn.com/downloads
- Scroll down half way to the Wireguard Configuration section
- Type in a name for the certificate you will generate, select Router, and configure your VPN options (Malware/ad/tracker blocking, Moderate NAT, NAT-PMP, VPN Accelerator)
- You can hit the create button right under these settings and Proton will select the server closest to you, or you can manually select a server that you prefer.
- Hit create
For the purpose of this demo, this is the configuration I will be the following configuration and referring back to this:
[Interface]
# Key for OPNsense Demo
# Bouncing = 5
# NetShield = 1
# Moderate NAT = off
# NAT-PMP (Port Forwarding) = off
# VPN Accelerator = on
PrivateKey = yEwTFeoVRgB/6pU8sTzyLseLnv8U7ZO/jAyVd4Uiym8=
Address = 10.2.0.2/32
DNS = 10.2.0.1
[Peer]
# US-CA#33
PublicKey = 4v/dB/ha+PGL0jihNVlVj81NGAFh6VndO9s4giDZEUw=
AllowedIPs = 0.0.0.0/0
Endpoint = 185.230.126.18:51820
Before we get started, there is one more parameter we will need!
ProtonVPN provides the Public Key for the Peer, and the Private Key for the Interface, however you will also need the Public Key for the interface. We will have to generate this ourselves You will need access to a terminal with wireguard installed for this, and the simplest way is to SSH into your OPNsense machine with the Wireguard plugin already installed.
If you do not already have the Wireguard plugin installed:
- Log into OPNsense
- Navigate to System > Firmware > Plugins
- Search for os-wireguard
- Press the + to install
If SSH is not already enabled on your OPNsense server:
- Navigate to System > Settings > Administration
- Check Enable Secure Shell
SSH is disabled by default, and unless you have a good reason to leave it enabled, you should disable it again when you are done
Now we need to generate the Interface Public Key:
- 
SSH into your OPNsense server ssh root@10.0.0.1
- 
Select option 8 to access the Shell
- 
Enter the following command: echo <PrivateKey> | wg pubkey eg. `echo yEwTFeoVRgB/6pU8sTzyLseLnv8U7ZO/jAyVd4Uiym8= | wg pubkey`
- 
Your Interface Public Key will be output on screen  eg. Interface Public Key: `IQjcNrCTlbtVrlkdvvPaV0NZK9wK5VofoiZeHfYzuA4=`
You can disable SSH on your OPNsense server at this time. We are ready to continue with actually configuring the VPN connection(s).
- Navigate to VPN > WireGuard > Endpoints
- Select the + to add a new endpoint
- Fill it out the form as follows:
Name:                <Anything you want>
Public Key:          <Peer Public Key>
		       eg.   4v/dB/ha+PGL0jihNVlVj81NGAFh6VndO9s4giDZEUw=
Shared Secret:       <Blank>
Allowed IPs:         0.0.0.0/0
Endpoint Address:    <Peer Endpoint IP>
               eg.   185.230.126.18
Endpoint Port:       51820
Keepalive Interval:  25
- Save
- Select the Local tab
- Select the + to add a new Local connection
- Select advanced mode
- Fill it out the form as follows:
Name:                <Anything you want>
Public Key:          <Leave Blank!!!>
Private Key:         <Leave Blank!!!>
Listen Port:         <Select a unique port for each tunnel>
MTU:                 1412
DNS Server:          <Leave Blank!!!>
Tunnel Address:      10.2.0.2/28
Peers:               <Select the Endpoint you just created>
Disable Routes:      Checked
Gateway:             <Select a unique IP for each tunnel>
This step is very important!
When we activate the connection, Wireguard will initially populate the Public and Private Keys. Allow it to do this. We will go back and change it to the correct values later.
Not setting the MTU to 1412 or 1420 will not prevent a Wireguard connection, but will cause many lost packets and severe performance degradation.
Setting a DNS Server at this stage will override all of OPNsense's DNS configurations. We will continue to use OPNsense's DNS configs by leaving this blank, and we will take care of DNS leaks later on.
The Tunnel Address provided by ProtonVPN is a /32 subnet. /28 is used in this guide with success. With a single tunnel, /32 may work, however with multiple, you will need to use /28 or some other range.
- Start Wireguard from the General Tab , or if it is already started, then turn it off and then on again. You can also restart it from the dashboard.
- Go to the Status Tab and then theHandshakes Tab . You should see new entries corresponding to your tunnel.
- Go back to the Local Tab and edit your tunnel
- You will see entries in there for the Public and Private Key. Delete these and replace them with your own values.
The Public Key will be the string you generated
eg. IQjcNrCTlbtVrlkdvvPaV0NZK9wK5VofoiZeHfYzuA4= The Private Key will be the string provided by ProtonVPN
eg.yEwTFeoVRgB/6pU8sTzyLseLnv8U7ZO/jAyVd4Uiym8=
- Save and restart the Wireguard
- Navigate to Interfaces > Assignments
- Add a new interface with your newly created tunnel
- Select that interface
- Check Enable Interface
- Change the name to your preference
- Leave everything else untouched
- Save and Apply Changes
- Navigate to System > Gateways > Single
- Add a new Gateway
- Fill out the form as follows:
Name:                <Anything you want>
Description:         <Your preference>
Interface:           <Select the interface you created in the pervious section>
Address Family:      IPv4
IP Address:          <Enter the Peer Endpoint IP>
               eg.   185.230.126.18
Upstream Gateway:    Unchecked
Far Gateway:         Checked
Disable Gateway Monitoring: Unchecked
Monitor IP:          Blank
- Save and Apply Changes
In this step we will create Aliases to essentially create labels for our various hosts and networks to base routing rules off of. While not strictly required, this is going to be significantly simpler than manually making Firewall rules directly with IPs.
- Navigate to Firewall > Aliases
- Add new Alias
- Fill out the form as follows:
Name:                <Something descriptive>
Type:                <Hosts for individual IPs or Networks for IP Ranges>
Categories:          <Blank>
Content:             <Enter your IPs or IP Ranges in CIDR Notation>
Description:         <Your preference>
- Save
- Create another Alias except with the following settings:
Name:                RFC1918_Networks
Type:                Networks
Categories:          <Blank>
Content:             192.168.0.0/16 10.0.0.0/8 172.16.0.0/12
Description:         All local (RFC1918) networks
- Save
- Navigate to Firewall > NAT > Outbound
- Select Hybrid outbound NAT rule generation
- Add a new rule
- Make the following changes:
Interface:                <Select your Wireguard Interface>
Source Address:           <Select the Alias you wish to route>
- Leave everything else default
- Save and Apply Changes
- Navigate to Firewall > Rules > LAN
- Add new rule
- Make the following changes:
Source:                <Select the Alias you wish to route>
Destination / Invert:  Checked
Destination:           RFC1918_Networks
Gateway:               <Select your Wireguard Gateway>
Click Advanced Options Show/Hide
Set local tag:         NO_WAN_EGRESS
The NO_WAN_EGRESS local tag will be used for the killswitch to prevent traffic leaking out if a tunnel is down
