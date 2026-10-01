---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/ideafieldpro-opnsense-install-config-vpn-blob-head-readme-md-c179bd92-3
title: "ideafieldpro-opnsense-install-config-vpn-blob-head-readme-md-c179bd92"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/ideafieldpro-opnsense-install-config-vpn-blob-head-readme-md-c179bd92.md
source_anchor: ""
source_lines: [421, 474]
sha256: fe69ed300cc6041035676141b0ce6d1212767aad19f26e0f05d36e265e0529b7
---

# ideafieldpro-opnsense-install-config-vpn-blob-head-readme-md-c179bd92

1. Verify LAN IP configuration:
  - Console: Option 2 → View current IP
2. Check client network settings:
  - Must be on same subnet (192.168.1.x)
  - Gateway should point to OPNsense LAN IP
3. Clear browser cache/try incognito mode
4. Disable client firewall temporarily

1. Verify WAN interface has IP address
  - Console → Option 2 → Select WAN → View
2. Check gateway configuration:
  - **System → Gateways → Single**
3. Verify DNS settings:
  - **System → Settings → General**
4. Test from OPNsense:
  - Console → Option 7 → Ping 8.8.8.8

1. Verify DHCP service is running:
  - **Services → Dnsmasq DNS & DHCP → Settings** → Ensure**Enable Dnsmasq** is checked
2. Restart DHCP service:
  - **Services → Dnsmasq DNS & DHCP → Settings** → Click**Apply** to restart
3. Check DHCP leases:
  - **Services → Dnsmasq DNS & DHCP → Leases**
4. Verify firewall rules allow DHCP (UDP ports 67/68) on the LAN interface

- Navigate to **System → Configuration → Backups**
- Download configuration XML weekly/monthly
- Store securely off-site

- Check: **System → Firmware → Updates**
- Enable automatic security updates
- Review changelog before major updates
- Schedule updates during maintenance windows

- Change default SSH port
- Enable two-factor authentication
- Implement firewall rule logging
- Regular log review via **System → Log Files**
- Consider IDS/IPS (Suricata) for advanced threat detection

- Enable dashboard widgets for at-a-glance status
- Monitor DHCP lease utilization
- Track VPN connection logs
- Set up email alerts for critical events

- Maintain network diagram
- Document all static mappings
- Keep firewall rule inventory
- Record configuration changes

- OPNsense Forum
- DHCP and DNSmasq deep dive: How to set up DHCP server on OPNsense (Zenarmor)

**Author**: Craig Sheffield (ideafieldpro)
