---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/blog-it-2022-02-02-wireguad-and-split-vpn-on-unifi-dream-machine-pro-se-24e3a435-1
title: "This script downloads the latest split-vpn and installs it"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: ["2020-03"]
keywords: ["exploit", "kill switch", "open source"]
source: docs/RAG/collect-261001-unifi-ubiquiti/blog-it-2022-02-02-wireguad-and-split-vpn-on-unifi-dream-machine-pro-se-24e3a435.md
source_anchor: ""
source_lines: [1, 70]
sha256: f45653998136c7b073eae48f17a9d08b83a7c48d4f51226fde25cc6a92ace0db
---

# This script downloads the latest split-vpn and installs it

Wireguad® and Split VPN on Unifi Dream Machine Pro SE (UDM PRO SE)
How to set up a helper script for multiple VPN clients on the UDM PRO SE that creates a split tunnel for the VPN connection, and forces configured clients through the VPN instead of the default WAN.
In this post we will see how to set up a helper script for multiple VPN clients on the UDM PRO SE that creates a split tunnel for the VPN connection, and forces configured clients through the VPN instead of the default WAN. This is accomplished by marking every packet of the forced clients with an iptables firewall mark (fwmark), adding the VPN routes to a custom routing table, and using a policy-based routing rule to direct the marked traffic to the custom table. This script works with OpenVPN, WireGuard®, OpenConnect, StrongSwan, or an external nexthop VPN client on your network.
What is VPN Split Tunneling ? Permalink
Split tunneling is a VPN feature that divides your internet traffic and sends some of it through an encrypted virtual private network (VPN) tunnel, but routes the rest through a separate tunnel on the open network. Typically, split tunneling will let you choose which apps, host, vlan to secure and which can connect normally.
This is a useful feature when you need to keep some of your traffic private, while still maintaining access to local network devices. So you can access foreign networks and local networks at the same time. It’s also great if you want to save some bandwidth.
How does VPN Split Tunneling work ? Permalink
Split tunneling is a clever VPN feature that gives you much more control over what data you encrypt and send through a VPN server, and what data travels through the faster, unencrypted open web.
So, how does it work? Well, in order to understand what VPN split tunneling is, you first need to understand the basics of a VPN server.
By default, your device will probably have a single, direct connection to the internet, through which your data will be sent and received. But, when you use a VPN, this creates a secure connection between your device and a VPN server. That VPN server then accesses the internet on your behalf. So, every single bit of data gets sent and received through the secure VPN server.
On the plus side, this keeps all your data completely encrypted. But, because everything needs to travel through the VPN, it can slow your internet speeds.
Split tunneling works by giving you two connections at the same time: the secure VPN connection and an open connection to the internet. So, you can protect your sensitive data without slowing down your other internet activities.
Why use WireGuard® ? Permalink
WireGuard® uses state-of-the-art cryptography, which makes it faster, more secure, and more friendly to mobile and IoT (Internet of Things) devices than other VPN (Virtual Private Network) technologies like OpenVPN or IPsec.
WireGuard® was developed in the last decade, using modern cryptographic primitives and protocols like ChaCha20/Poly1305, Curve25519, BLAKE2, SipHash24, HKDF, and the Noise protocol. Because this crypto is (relatively) easy to implement and understand, the standard C WireGuard implementation is only about 6,000 lines of code. For comparison, that’s around 100 times smaller than other VPN implementations that are saddled with 90s-era cryptography, like OpenVPN (OpenSSL) or strongSwan (IPsec).
- High performance
WireGuard’s modern crypto means that it’s faster than other VPN technologies at establishing connections (and re-establishing connections on flaky networks). And it’s lightweight, adding minimal overhead when encrypting and decrypting network traffic. Plus, WireGuard has been built into the Linux kernel since March 2020, allowing it to run even faster on systems with modern Linux kernels.
While speed tests may vary from network to network and implementation to implementation, several recently-published comparisons of WireGuard to OpenVPN and IPsec show WireGuard to be the clear performance champion:
- Security
WireGuard’s modern crypto also makes it more secure. Instead of offering system administrators a million different cryptographic configuration combinations, like OpenVPN or IPsec do, WireGuard has just one. This means that WireGuard is always secured with the industry’s best practices, right out-of-the-box — you can’t shoot yourself in the foot with the infamous null cipher suite (or a million other configuration pitfalls) like you can with OpenVPN or IPsec.
This straightforward cryptographic design also leads to a much smaller attack surface. Without a million cryptographic options, and with a small, readable code-base, it’s easy for a defender to audit the WireGuard source code — and difficult for an attacker to find any hidden issues to exploit.
- Open Source
WireGuard is open source (and free software — the standard C implementation is GPLv2), so everyone is free to download the source code, audit it, tinker with it, and deploy it to any number of servers or endpoints, completely free of charge. And because the source code is open, it has been audited, probed, and formally verified by a number of teams and techniques.
For more details you can read this comparison of VPN protocols from IVPN provider.
Split VPN on UDM Pro SE Permalink
In this part we show : How to install the helper Split-VPN script, made by Peacey, on Unifi Dream Machine Pro SE. And how to configure WireGuarde® protocol with the IVPN.
Features Permalink
- Works with UDM-Pro, UDM base, and UDM-Pro-SE, UDR, and UXG-Pro.
- Force traffic to the VPN based on source interface (VLAN), MAC address, IP address, or IP sets.
- Exempt sources from the VPN based on IP, MAC address, IP:port, MAC:port combinations, or IP sets. This allows you to force whole VLANs through by interface, but then selectively choose clients from that VLAN, or specific services on forced clients, to exclude from the VPN.
- Exempt destinations from the VPN by IP. This allows VPN-forced clients to communicate with the LAN or other VLANs.
- Force domains to the VPN or exempt them from the VPN (only supported with dnsmasq or pihole).
- Port forwarding on the VPN side to local clients (not all VPN providers give you ports).
- Redirect DNS for VPN traffic to either an upstream DNS server or a local server like pihole, or block DNS requests completely.
- Built-in kill switch via iptables and blackhole routing.
- Works across IP changes, network restarts, and the UDM’s WAN Failover.
- Can be used with multiple openvpn instances with separate configurations for each. This allows you to force different clients through different VPN servers.
- IPv6 support for all options.
- Run on boot support via UDM-Utilities boot script.
- Supports OpenVPN, WireGuard kernel module, WireGuard-go docker container, OpenConnect docker container (AnyConnect), StrongSwan docker container (IKEv2 and IPSec), and external VPN clients on your network (nexthop).
Install Split-VPN helper script on UDM-Pro-SE  Permalink   
1. SSH into the Unifi Dream Machine
ssh root@<udm IP>
2. Download and run the installation script
curl -LSsf https://raw.githubusercontent.com/peacey/split-vpn/main/vpn/install-split-vpn.sh | sh
#!/bin/sh
# This script downloads the latest split-vpn and installs it
# to the data directory (/mnt/data or /data, whichever exists).
set -e
# Get the persistent data directory
if [ -d "/mnt/data" ]; then
	DATA_DIR="/mnt/data"
elif [ -d "/data" ]; then
	DATA_DIR="/data"
else
	echo ERROR: Could not find the data directory.
	exit 1
fi
# Download and install
mkdir -p "${DATA_DIR}/split-vpn"
cd "${DATA_DIR}/split-vpn"
echo Downloading latest split-vpn...
curl -LsSfo split-vpn.zip https://github.com/peacey/split-vpn/archive/main.zip
echo Installing to "${DATA_DIR}/split-vpn"...
unzip -oq split-vpn.zip
cp -rf split-vpn-main/vpn ./
rm -rf split-vpn-main split-vpn.zip
chmod +x vpn/*.sh vpn/hooks/*/*.sh vpn/vpnc-script
# Link to /etc
rm -f /etc/split-vpn
