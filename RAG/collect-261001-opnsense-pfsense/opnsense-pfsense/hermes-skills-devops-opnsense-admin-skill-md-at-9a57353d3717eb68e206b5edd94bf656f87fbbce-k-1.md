---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/hermes-skills-devops-opnsense-admin-skill-md-at-9a57353d3717eb68e206b5edd94bf656f87fbbce-k-1
title: "Base URL"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["decode"]
source: docs/RAG/collect-261001-opnsense-pfsense/hermes-skills-devops-opnsense-admin-skill-md-at-9a57353d3717eb68e206b5edd94bf656f87fbbce-kngender5-h.md
source_anchor: ""
source_lines: [1, 277]
sha256: 60bb501cb45490835b1155e7649efb6eaa4b3b224c595914a001e5310bf03033
---

# Base URL

| name | opnsense-admin | 
|---|---|
| description | Manage OPNsense firewall — rules, NAT, VLANs, DNS, IDS/IPS, VPN, backups via REST API and SSH. Use when administering OPNsense firewall, configuring network security, managing Suricata, Unbound DNS, WireGuard/OpenVPN, or automating firewall operations. | 

Manage OPNsense firewall through REST API and SSH. Covers firewall rules, NAT, VLANs, DNS, IDS/IPS, VPN, backups, and service management.

OPNsense is an open-source FreeBSD-based firewall and routing platform. Fork of pfSense, with modern UI and frequent updates.

**Key Components:**

- **Firewall:** Stateful packet filtering, NAT, rules
- **IDS/IPS:** Suricata (Snort-compatible)
- **DNS:** Unbound (resolver), DNS forwarding
- **VPN:** WireGuard, OpenVPN, IPsec
- **VLAN:** 802.1Q tagging
- **HA:** CARP, failover
- **API:** REST API for automation

```
# Base URL
https://opnsense.example.com
# Auth: API Key + Secret
# Generate in: System → Access → Users → API Keys
Header: Authorization: Basic $(echo -n "api-key:api-secret" | base64)
```
```
ssh root@opnsense.example.com
# or
ssh -i ~/.ssh/opnsense admin@opnsense.example.com
```
```
curl -s -u "api-key:api-secret" \
  "https://opnsense.example.com/api/firewall/filter/getRule/{}" | jq .
```
```
curl -s -X POST -u "api-key:api-secret" \
  "https://opnsense.example.com/api/firewall/filter/addRule/{}" \
  -H "Content-Type: application/json" \
  -d '{
    "rule": {
      "description": "Allow HTTPS to web server",
      "source_net": "any",
      "destination_net": "10.0.1.10",
      "destination_port": "443",
      "protocol": "TCP",
      "action": "pass",
      "enabled": "1",
      "sequence": "1",
      "direction": "in"
    }
  }'
```
```
curl -s -X POST -u "api-key:api-secret" \
  "https://opnsense.example.com/api/firewall/filter/apply"
```
```
curl -s -u "api-key:api-secret" \
  "https://opnsense.example.com/api/firewall/source_nat/getRule/{}" | jq .
```
```
curl -s -X POST -u "api-key:api-secret" \
  "https://opnsense.example.com/api/firewall/source_nat/addRule/{}" \
  -H "Content-Type: application/json" \
  -d '{
    "rule": {
      "description": "Port forward HTTPS to internal server",
      "source": "any",
      "destination": "any",
      "destination_port": "443",
      "target": "10.0.1.10",
      "target_port": "443",
      "protocol": "TCP",
      "interface": "WAN"
    }
  }'
```
```
curl -s -u "api-key:api-secret" \
  "https://opnsense.example.com/api/interfaces/vlan/getItem/{}" | jq .
```
```
curl -s -X POST -u "api-key:api-secret" \
  "https://opnsense.example.com/api/interfaces/vlan/addItem/{}" \
  -H "Content-Type: application/json" \
  -d '{
    "vlan": {
      "if": "igb0",
      "tag": "100",
      "descr": "IoT Network",
      "mode": "normal"
    }
  }'
```
```
curl -s -u "api-key:api-secret" \
  "https://opnsense.example.com/api/unbound/settings/get" | jq .
```
```
curl -s -X POST -u "api-key:api-secret" \
  "https://opnsense.example.com/api/unbound/settings/addHost/{}" \
  -H "Content-Type: application/json" \
  -d '{
    "host": {
      "hostname": "nas",
      "domain": "home.local",
      "rr": "A",
      "server": "10.0.1.5",
      "description": "NAS server"
    }
  }'
```
```
curl -s -X POST -u "api-key:api-secret" \
  "https://opnsense.example.com/api/unbound/service/reconfigure"
```
```
curl -s -u "api-key:api-secret" \
  "https://opnsense.example.com/api/ids/service/status" | jq .
```
```
curl -s -X POST -u "api-key:api-secret" \
  "https://opnsense.example.com/api/ids/settings/set" \
  -H "Content-Type: application/json" \
  -d '{
    "ids": {
      "enabled": "1",
      "ips": "1",
      "promisc": "0",
      "interfaces": ["wan", "lan"]
    }
  }'
```
```
curl -s -X POST -u "api-key:api-secret" \
  "https://opnsense.example.com/api/ids/service/reloadRules"
```
```
curl -s -u "api-key:api-secret" \
  "https://opnsense.example.com/api/wireguard/client/get" | jq .
```
```
curl -s -X POST -u "api-key:api-secret" \
  "https://opnsense.example.com/api/wireguard/client/addItem/{}" \
  -H "Content-Type: application/json" \
  -d '{
    "client": {
      "name": "laptop",
      "pubkey": "client-public-key-here",
      "tunnel-address": "10.100.0.2/32",
      "serveraddress": "10.100.0.1",
      "serverport": "51820"
    }
  }'
```
```
curl -s -u "api-key:api-secret" \
  "https://opnsense.example.com/api/core/backup/download/this" \
  -o opnsense-backup-$(date +%Y%m%d).xml
```
```
# Via SSH
ssh root@opnsense "ls -la /conf/backup/"
```
```
curl -s -X POST -u "api-key:api-secret" \
  "https://opnsense.example.com/api/core/system/reboot"
```
```
# Check for updates
curl -s -u "api-key:api-secret" \
  "https://opnsense.example.com/api/core/firmware/status" | jq .
# Upgrade
curl -s -X POST -u "api-key:api-secret" \
  "https://opnsense.example.com/api/core/firmware/upgrade" \
  -H "Content-Type: application/json" \
  -d '{"upgrade": "1"}'
```
```
curl -s -u "api-key:api-secret" \
  "https://opnsense.example.com/api/diagnostics/system/systemInformation" | jq .
```
```
import requests, json
from base64 import b64encode
class OPNsense:
    def __init__(self, host, api_key, api_secret, verify_ssl=True):
        self.base_url = f"https://{host}"
        auth = b64encode(f"{api_key}:{api_secret}".encode()).decode()
        self.headers = {
            "Authorization": f"Basic {auth}",
            "Content-Type": "application/json",
        }
        self.verify = verify_ssl
    
    def get(self, path):
        resp = requests.get(
            f"{self.base_url}{path}",
            headers=self.headers,
            verify=self.verify,
        )
        return resp.json()
    
    def post(self, path, data=None):
        resp = requests.post(
            f"{self.base_url}{path}",
            headers=self.headers,
            json=data or {},
            verify=self.verify,
        )
        return resp.json()
    
    # Firewall rules
    def list_rules(self):
        return self.get("/api/firewall/filter/getRule/{}")
    
    def add_rule(self, description, dest, port, protocol="TCP", action="pass"):
        return self.post("/api/firewall/filter/addRule/{}", {
            "rule": {
                "description": description,
                "source_net": "any",
                "destination_net": dest,
                "destination_port": port,
                "protocol": protocol,
                "action": action,
                "enabled": "1",
            }
        })
    
    def apply_rules(self):
        return self.post("/api/firewall/filter/apply")
    
    # NAT
    def list_nat(self):
        return self.get("/api/firewall/source_nat/getRule/{}")
    
    def add_port_forward(self, desc, ext_port, int_ip, int_port, protocol="TCP"):
        return self.post("/api/firewall/source_nat/addRule/{}", {
            "rule": {
                "description": desc,
                "source": "any",
                "destination": "any",
                "destination_port": ext_port,
                "target": int_ip,
                "target_port": int_port,
                "protocol": protocol,
                "interface": "WAN",
            }
        })
    
    # DNS
    def add_dns_override(self, hostname, domain, ip, rr="A"):
        return self.post("/api/unbound/settings/addHost/{}", {
            "host": {
                "hostname": hostname,
                "domain": domain,
                "rr": rr,
                "server": ip,
            }
        })
    
    def apply_dns(self):
        return self.post("/api/unbound/service/reconfigure")
    
    # IDS/IPS
    def get_ids_status(self):
        return self.get("/api/ids/service/status")
    
    def enable_ids(self, interfaces=None):
        return self.post("/api/ids/settings/set", {
            "ids": {
                "enabled": "1",
                "ips": "1",
                "interfaces": interfaces or ["wan"],
            }
        })
    
