---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/blog-it-2022-02-02-wireguad-and-split-vpn-on-unifi-dream-machine-pro-se-24e3a435-3
title: "This script downloads the latest split-vpn and installs it"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["kill switch"]
source: docs/RAG/collect-261001-unifi-ubiquiti/blog-it-2022-02-02-wireguad-and-split-vpn-on-unifi-dream-machine-pro-se-24e3a435.md
source_anchor: ""
source_lines: [153, 254]
sha256: 4fb266fbd401eb1849e03ac1285a8853be92fc7b99c00b517fc193768e1da0fd
---

# This script downloads the latest split-vpn and installs it

- Endpoint : Set here IP of Wireguard service provider. ForIVPN provider to obtainIP go to https://www.ivpn.net/status/ and for example selectFrance inCountry list.
Now, go to your Terminal and type nslookup fr1.wg.ivpn.net
So, the Wiregard Endpoint is IP 185.246.211.185
- What ports do you use for WireGuard? UDP 53
UDP 80
UDP 443
UDP 1194
UDP 2049
UDP 2050
UDP 30587
UDP 41893
UDP 48574
UDP 58237
In my Wireguard config (wg0.conf) file, I uses 2049 as UDP port
- PublicKey : Set here the public key of your Wireguard IP VPN serverEndpoint . So in this examplePublicKey egual tog7BuMzj3r<redacted>
3. Edit the vpn.conf file with your desired settings. Make sure that:
- The option DNS_IPV4_IP and/orDNS_IPV6_IP is set to the DNS server you want to force for your clients, or set them to empty if you do not want to force any DNS.
  - IVPN provider DNS entry:
    - 10.0.254.1 = regular DNS with no blocking (OpenVPN only) (use10.0.254.101 for Multi-hop connections).
    - 10.0.254.2 = standard AntiTracker to block advertising and malware domains (OpenVPN + WireGuard) (use10.0.254.102 for Multi-hop connections).
    - 10.0.254.3 = Hardcore Mode AntiTracker to also block Google and Facebook domains (OpenVPN + WireGuard) (use10.0.254.103 for Multi-hop connections).
- IVPN provider DNS entry:
- The option VPN_PROVIDER is set to"external" for WireGuard.
- The option VPN_ENDPOINT_IPV4 orVPN_ENDPOINT_IPV6 is set to your WireGuard server’s IP as defined inwg0.conf ’s Endpoint variable (e.g.172.20.xxx.yyy inwg0.conf ).
- The option ROUTE_TABLE is the same number (101 ) as Table in yourwg0.conf file.
- The option DEV is set to"wg0" or your interface’s name if different (i.e.: the name of your.conf file).
- The option MSS_CLAMPING_IPV4 is set for me to"1380"
4. Run wg-quick to start WireGuard with your configuration and test if the connection worked. Replace wg0 with your interface name if different.
/mnt/data/wireguard/setup_wireguard.sh
wg-quick up ./wg0.conf
- You can skip the first line if you already setup the WireGuard kernel module previously as instructed at wireguard-kmod .
- Type wg to check your WireGuard connection and make sure you received a handshake. No handshake indicates something is wrong with your WireGuard configuration. Double check your configuration’s Private and Public key and other variables.
interface: wg0
  public key: xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
  private key: (hidden)
  listening port: 56092
peer: yyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyy
  endpoint: 185.246.211.185:2049
  allowed ips: 0.0.0.0/1, 128.0.0.0/1
  latest handshake: 1 minute, 25 seconds ago
  transfer: 26.71 GiB received, 1.20 GiB sent
- If you need to bring down the WireGuard tunnel, run wg-quick down ./wg0.conf in this folder (replacewg0.conf with your interface configuration if different).
- Note that wg-quick up/down commands need to be run from this folder so the script can pick up the correct configuration file.
5. If the connection works, check each client to make sure they are on the VPN. See the Tips How do I check my clients are on the VPN? below.
6. If everything is working, create a run script called run-vpn.sh and stop-vpn.sh in the current directory so you can easily run this WireGuard configuration. Fill the script with the following contents:
- run-vpn.sh
#!/bin/sh
# Set up the WireGuard kernel module and tools
/mnt/data/wireguard/setup_wireguard.sh
# Load configuration and run wireguard
cd /etc/split-vpn/wireguard/ivpn
. ./vpn.conf
# /etc/split-vpn/vpn/updown.sh ${DEV} pre-up >pre-up.log 2>&1
wg-quick up ./${DEV}.conf >wireguard.log 2>&1
cat wireguard.log
- stop-vpn.sh
#!/bin/sh
cd /mnt/data/split-vpn/wireguard/ivpn/
wg-quick down /mnt/data/split-vpn/wireguard/ivpn//wg0.conf
- Modify the cd line to point to the correct directory. Make sure that theDEV variable in thevpn.conf file is set to the WireGuard interface name (which should the same as the WireGuard configuration filename without.conf ).
- Optional: If you want to block Internet access to forced clients if the WireGuard tunnel is brought down via wg-quick , setKILLSWITCH=1 andREMOVE_KILLSWITCH_ON_EXIT=0 in thevpn.conf file.
- Optional: Uncomment the pre-up line by removing the# at the beginning of the line if you want to block Internet access for forced clients if WireGuard fails to run. Keeping it commented out doesn’t enable the iptables kill switch until after WireGuard runs successfully.
7. Give the script executable permissions. You can run this script next time you want to start this WireGuard configuration.
chmod +x /etc/split-vpn/wireguard/ivpn/run-vpn.sh
chmod +x /etc/split-vpn/wireguard/ivpn/stop-vpn.sh
How do I run this at boot? Permalink
On the UDM Pro SE, boot scripts are supported natively via systemd. The boot script survives across firmware upgrades and reboots.
1. Create a master run script under /etc/split-vpn/run-vpn.sh that will be used to run your VPNs. In this master script, call the run script of each VPN client that you want to run at boot. For example, here we are running a WireGuard client and an OpenVPN client.
#!/bin/sh
/etc/split-vpn/wireguard/ivpn/run-vpn.sh
/etc/split-vpn/openvpn/protonvpn/run-vpn.sh
2. Give the master run script executable permissions.
chmod +x /etc/split-vpn/run-vpn.sh
3. Install the boot service for your device.
For the UDM-SE or UDR, run the following commands to install a systemd boot service.
curl -o /etc/systemd/system/run-vpn.service https://raw.githubusercontent.com/peacey/split-vpn/main/examples/boot/run-vpn.service
systemctl daemon-reload && systemctl enable run-vpn
[Unit]
Description=Run split-vpn
Wants=network-online.target
After=network-online.target
[Service]
Type=forking
ExecStartPre=sh -c 'VPN_DIR="$(find /mnt/data/split-vpn /data/split-vpn -maxdepth 1 -type d -name "split-vpn" 2>/dev/null | head -n1)"; rm -f /etc/split-vpn; ln -sf "$VPN_DIR" /etc/split-vpn'
ExecStart=/etc/split-vpn/run-vpn.sh
Restart=on-failure
RestartSec=5
[Install]
WantedBy=multi-user.target
The default systemd service is set to restart automatically on failure. If you do not want this behaivour, modify /etc/systemd/system/run-vpn.service and remove the Restart=... line.
That’s it. Now the VPN will start at every boot.
Tips Permalink
Tips: Configuration variables for vpn.conf file  Permalink   
Settings are modified in vpn.conf. Multiple entries can be entered for each setting by separating the entries with spaces. Click here to see all the settings.
Tips: Customise MTU, MSS and MSS clamping Permalink
- What is MTU?
In networking, Maximum Transmission Unit (MTU) is a measurement representing the largest data packet that a network-connected device will accept. Imagine it as being like a height limit for freeway underpasses or tunnels: Cars and trucks that exceed the height limit cannot fit through, just as packets that exceed the MTU of a network cannot pass through that network.
However, unlike cars and trucks, data packets that exceed MTU are broken up into smaller pieces so that they can fit through. This process is called fragmentation. Fragmented packets are reassembled once they reach their destination.
MTU is measured in bytes — a “byte” is equal to 8 bits of information, meaning 8 ones and zeroes. 1500 bytes is the maximum MTU size.
For Wireguard, the MTU is egal to 1420 byte
- What is Maximum Segment Size (MSS)?
MSS (maximum segment size) limits the size of packets, or small chunks of data, that travel across a network, such as the Internet. All data that travels over a network is broken up into packets. Packets have several headers attached to them that contain information about their contents and destination. MSS measures the non-header portion of a packet, which is called the payload.
