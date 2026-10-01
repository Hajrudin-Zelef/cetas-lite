---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/hermes-skills-devops-opnsense-admin-skill-md-at-9a57353d3717eb68e206b5edd94bf656f87fbbce-k-2
title: "Base URL"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "memory"]
source: docs/RAG/collect-261001-opnsense-pfsense/hermes-skills-devops-opnsense-admin-skill-md-at-9a57353d3717eb68e206b5edd94bf656f87fbbce-kngender5-h.md
source_anchor: ""
source_lines: [278, 432]
sha256: 6c039ddcc7615a3f9d20a4f5b9d6ef4575df56146f7cb457db451e968a8bef7e
---

    # Backup
    def download_backup(self, path="opnsense-backup.xml"):
        auth = self.headers["Authorization"]
        resp = requests.get(
            f"{self.base_url}/api/core/backup/download/this",
            headers={"Authorization": auth},
            verify=self.verify,
        )
        with open(path, "wb") as f:
            f.write(resp.content)
        return path
    
    # System
    def get_status(self):
        return self.get("/api/diagnostics/system/systemInformation")
    
    def reboot(self):
        return self.post("/api/core/system/reboot")
# Usage
fw = OPNsense("opnsense.example.com", "api-key", "api-secret", verify_ssl=False)
# Add firewall rule
fw.add_rule("Allow HTTPS", "10.0.1.10", "443")
fw.apply_rules()
# Add port forward
fw.add_port_forward("Web Server", "443", "10.0.1.10", "443")
# DNS override
fw.add_dns_override("nas", "home.local", "10.0.1.5")
fw.apply_dns()
# Check IDS status
print(fw.get_ids_status())
# Download backup
fw.download_backup("backup.xml")
```
```
# Enter shell
ssh root@opnsense
# View firewall rules
pfctl -sr
# View NAT rules
pfctl -sn
# View states
pfctl -ss
# Restart firewall
pfctl -f /etc/pf.conf
# View Suricata logs
tail -f /var/log/suricata/eve.json
# View system logs
clog /var/log/system.log
# Backup config
cp /conf/config.xml /conf/backup/config-$(date +%Y%m%d).xml
# Restore config
cp /conf/backup/config-good.xml /conf/config.xml
# Reboot to apply
reboot
# Update OPNsense
opnsense-update
# or via GUI: System → Firmware → Update
# Install packages
pkg install os-wireguard
pkg install os-zabbix-agent
```
```
# Via API
fw.add_rule("Block malicious IP", "192.0.2.100", "any", action="block")
fw.apply_rules()
# Via SSH (quick block)
echo "block drop quick from 192.0.2.100 to any" >> /etc/pf.conf
pfctl -f /etc/pf.conf
```
```
# Usually default, but if needed:
fw.add_rule("Allow LAN to WAN", "10.0.1.0/24", "any")
fw.apply_rules()
```
```
# 1. Generate keys
wg genkey | tee privatekey | wg pubkey > publickey
# 2. Add tunnel via API
fw.post("/api/wireguard/server/addItem/{}", {
    "server": {
        "name": "wg0",
        "enabled": "1",
        "instance": "0",
        "pubkey": open("publickey").read().strip(),
        "privkey": open("privatekey").read().strip(),
        "port": "51820",
        "tunneladdress": "10.100.0.1/24",
    }
})
# 3. Add peer
fw.post("/api/wireguard/client/addItem/{}", {
    "client": {
        "name": "laptop",
        "pubkey": "client-pubkey",
        "tunneladdress": "10.100.0.2/32",
    }
})
# 4. Add firewall rule for WireGuard
fw.add_rule("Allow WireGuard", "any", "51820", "UDP")
fw.apply_rules()
```
```
# Via SSH — add to cron
cat >> /etc/crontab << 'EOF'
0 3 * * * root cp /conf/config.xml /conf/backup/config-$(date +\%Y\%m\%d).xml
EOF
# Or via API — create a script that downloads backup
# Add to Hermes cron:
# 0 3 * * * /home/kng/bin/opnsense-backup.sh
```
```
# configuration.yaml
opnsense:
  host: 10.0.1.1
  api_key: !secret opnsense_api_key
  api_secret: !secret opnsense_api_secret
  ssl: false
# Sensors
sensor:
  - platform: opnsense
    host: 10.0.1.1
    api_key: !secret opnsense_api_key
    api_secret: !secret opnsense_api_secret
    monitored_conditions:
      - cpu
      - memory
      - wan_status
```
| Problem | Solution | 
|---|---|
| API returns 401 | Check API key/secret, user has API access permission | 
| Can't connect to web UI | `configctl webgui restart` via SSH | 
| Firewall rules not applying | Run `pfctl -f /etc/pf.conf` or use API apply endpoint | 
| Suricata not starting | Check interface assignment, rule download: `ids/service/reloadRules` | 
| WireGuard tunnel down | Verify keys, check firewall rule for UDP port | 
| Config corrupt | Restore from backup: `cp /conf/backup/config-good.xml /conf/config.xml` | 
| Update fails | Check disk space: `df -h` ; try`opnsense-update -f` | 
| DNS not resolving | Check Unbound: `unbound-control status` ; restart:`unbound-control reload` | 

1. **Use API keys** instead of password auth
2. **Restrict API access** to specific IPs
3. **Enable HTTPS** with valid certificate
4. **Regular backups** — automate daily config backups
5. **Keep updated** — apply security patches promptly
6. **Least privilege** — create dedicated API user with minimal permissions
7. **Enable IDS/IPS** — Suricata on WAN interface
8. **Disable unused services** — turn off what you don't need
9. **Strong passwords** — for all user accounts
10. **2FA** — enable for web UI access

- **API Docs:** https://docs.opnsense.org/development/api.html
- **GitHub:** https://github.com/opnsense/core
- **LobeHub Skill:** https://lobehub.com/skills/openclaw-skills-opnsense-admin
- **Python Client:** https://github.com/O-X-L/opnsense-api-client
- **CLI Tool:** https://github.com/andreas-stuerz/opn-cli
