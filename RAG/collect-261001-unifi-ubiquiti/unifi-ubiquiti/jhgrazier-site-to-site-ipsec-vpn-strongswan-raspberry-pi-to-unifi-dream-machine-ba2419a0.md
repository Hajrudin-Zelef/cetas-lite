---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/jhgrazier-site-to-site-ipsec-vpn-strongswan-raspberry-pi-to-unifi-dream-machine-ba2419a0
title: "jhgrazier-site-to-site-ipsec-vpn-strongswan-raspberry-pi-to-unifi-dream-machine-ba2419a0"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/jhgrazier-site-to-site-ipsec-vpn-strongswan-raspberry-pi-to-unifi-dream-machine-ba2419a0.md
source_anchor: ""
source_lines: [1, 159]
sha256: 78f2df6ca5d997209bf1c3a726d665be319ca680a8ad92fb86b1b94aa14b2e4f
---

# jhgrazier-site-to-site-ipsec-vpn-strongswan-raspberry-pi-to-unifi-dream-machine-ba2419a0

This guide documents a working site-to-site IPsec VPN using strongSwan on a Raspberry Pi and a UniFi Dream Machine.
Local (UDM side):
- 192.168.32.0/24
- 10.10.1.0/24
- 10.10.2.0/24
- 10.20.30.0/24
- 10.100.253.0/24
- 10.100.254.0/24
Remote (Pi side):
- 192.168.40.0/24
- Raspberry Pi IP: 192.168.40.250
Public IP / FQDN:
- UDM WAN: x.x.x.x
Install strongSwan on the Pi:
sudo apt update
sudo apt install -y strongswan strongswan-pki
Enable at boot:
sudo systemctl enable strongswan-starter
Enable IPv4 forwarding:
sudo sysctl -w net.ipv4.ip_forward=1
echo "net.ipv4.ip_forward=1" | sudo tee /etc/sysctl.d/99-ipsec-forwarding.conf
sudo sysctl --system
config setup
    charondebug="ike 1, knl 1, cfg 1"
conn udm-site
    keyexchange=ikev2
    ike=aes256-sha256-modp2048
    esp=aes256-sha256-modp2048
    authby=psk
    left=%any
    leftid=@far-pi
    leftsubnet=192.168.40.0/24
    right=x.x.x.x
    rightid=@x.x.x.x
    rightsubnet=192.168.32.0/24,10.10.1.0/24,10.10.2.0/24,10.20.30.0/24,10.100.253.0/24,10.100.254.0/24
    dpddelay=10s
    dpdtimeout=40s
    dpdaction=restart
    closeaction=restart
    keyingtries=%forever
    reauth=no
    auto=start
Important:
- Use straight ASCII double quotes only.
- leftid and rightid must match UDM configuration exactly.
- Remove trailing spaces inside subnet lists.
@far-pi @x.x.x.x : PSK "REPLACE_WITH_SHARED_KEY"
Quotes must be ASCII double quotes.
Create a Site-to-Site VPN:
VPN Type: Manual IPsec
IKE Version: IKEv2
Authentication: Pre-Shared Key
IKE Proposal:
- AES-256
- SHA-256
- DH Group 14
ESP Proposal:
- AES-256
- SHA-256
- PFS disabled or Group 14
Local Networks:
- 192.168.32.0/24
- 10.10.1.0/24
- 10.10.2.0/24
- 10.20.30.0/24
- 10.100.253.0/24
- 10.100.254.0/24
Remote Networks:
- 192.168.40.0/24
Remote Gateway:
- Pi public IP or dynamic DNS
Keep NAT-T enabled unless both sides have public IPs.
Restart strongSwan:
sudo systemctl restart strongswan-starter
Check status:
ipsec statusall
You should see:
- IKE_SA: ESTABLISHED
- CHILD_SA: INSTALLED
- Traffic counters increasing
DPD configuration:
dpddelay=10s
dpdtimeout=40s
dpdaction=restart
closeaction=restart
keyingtries=%forever
reauth=no
Create watchdog script:
sudo tee /usr/local/sbin/ipsec-watchdog.sh >/dev/null <<'EOF'
#!/usr/bin/env bash
set -euo pipefail
CONN="udm-site"
S="$(ipsec statusall 2>/dev/null || true)"
if ! grep -qE "^ *${CONN}\[[0-9]+\]: ESTABLISHED" <<<"$S" || \
   ! grep -qE "^ *${CONN}\{[0-9]+\}: +INSTALLED, +TUNNEL" <<<"$S"; then
  logger -t ipsec-watchdog "Tunnel unhealthy, restarting strongswan-starter"
  systemctl restart strongswan-starter
fi
EOF
Make executable:
sudo chmod +x /usr/local/sbin/ipsec-watchdog.sh
Create service:
sudo tee /etc/systemd/system/ipsec-watchdog.service >/dev/null <<'EOF'
[Unit]
Description=Watchdog for strongSwan site-to-site tunnel
After=network-online.target strongswan-starter.service
Wants=network-online.target
[Service]
Type=oneshot
ExecStart=/usr/local/sbin/ipsec-watchdog.sh
EOF
Create timer:
sudo tee /etc/systemd/system/ipsec-watchdog.timer >/dev/null <<'EOF'
[Unit]
Description=Run strongSwan tunnel watchdog every minute
[Timer]
OnBootSec=30
OnUnitActiveSec=60
AccuracySec=5
[Install]
WantedBy=timers.target
EOF
Enable:
sudo systemctl daemon-reload
sudo systemctl enable --now ipsec-watchdog.timer
Verify:
systemctl list-timers | grep ipsec-watchdog
journalctl -t ipsec-watchdog -n 50
The Pi acts as a router.
Do not SNAT VPN traffic.
Ensure forwarding policy allows traffic:
sudo iptables -F
sudo iptables -P FORWARD ACCEPT
From Pi:
ping 192.168.32.10
From UDM LAN:
traceroute 192.168.40.88
Expected:
- Hop through Pi
- Replies from remote hosts
Live logs:
journalctl -u strongswan-starter -f
Common issues:
AUTH_FAILED
PSK mismatch or ID mismatch.
NO_PROP
IKE or ESP proposal mismatch.
Tunnel established but no traffic
Routing or firewall issue on remote host.
- UDM periodically renegotiates CHILD_SA.
- Multiple remote subnets work inside a single tunnel.
- strongSwan starts automatically at boot.
- Watchdog ensures automatic recovery.
Verified working:
- IKEv2
- NAT traversal
- Multi-subnet routing
- Bidirectional ICMP and TCP traffic
- Automatic tunnel recovery
