---
id: etape7-phaseg-linuxnet-vpn/00-linuxnet-vpn/19-6-tailscale-acl-policy-least-privilege-example
title: "19.6 Tailscale ACL policy (least-privilege example)"
domain: step-7g-linux-networking-and-access-nat-firewalls-ssh-vpns-a
role: deep-dive
task: regulation
actors: []
dates: ["2026-09-22"]
keywords: ["agents", "research"]
source: docs/RAG/etape7_phaseG_linuxnet_vpn.md
source_anchor: ""
source_lines: [751, 834]
section: "Step 7G — Linux Networking and Access: NAT, Firewalls, SSH, VPNs, and Security Monitoring"
sha256: 54e4c9538412fc2a678bb58b5481f6bc0a479e2a22aa2a6f1701dc7e770e2a98
---

# 19.6 Tailscale ACL policy (least-privilege example)

### 19.6 Tailscale ACL policy (least-privilege example)

```jsonc
{
  "tagOwners": {
    "tag:prod":   ["autogroup:admin"],
    "tag:worker": ["autogroup:admin"]
  },
  "acls": [
    { "action": "accept", "src": ["tag:worker"], "dst": ["tag:prod:6379", "tag:prod:27017"] },
    { "action": "accept", "src": ["autogroup:admin"], "dst": ["*:*"] }
  ],
  "ssh": [
    { "action": "check", "src": ["autogroup:admin"], "dst": ["tag:prod"], "users": ["deploy", "root"] }
  ]
}
```

- `"action": "check"` in the `ssh` section requires per-connection approval — suited to privileged access on sensitive systems [secondary: section 8 sources].
- Define `tagOwners` before nodes join with `--advertise-tags`, or joins are rejected [secondary: section 8 sources].

### 19.7 Wazuh custom decoder and rule sketch

```xml
<decoder name="myapp">
  <program_name>myapp</program_name>
</decoder>

<rule id="100200" level="7">
  <decoded_as>myapp</decoded_as>
  <match>authentication failure</match>
  <description>MyApp authentication failure</description>
</rule>
```

- Decoders parse logs; rules match decoded fields to raise alerts; custom rules, decoders, CDB lists, and SCA policies are the artifacts to version in Git for disaster recovery [secondary: section 6 sources].

### 19.8 UFW and firewalld quick reference

```bash
ufw default deny incoming
ufw default allow outgoing
ufw allow 22/tcp
ufw allow from 192.168.1.0/24 to any port 5432
ufw --dry-run enable

firewall-cmd --permanent --add-service=ssh
firewall-cmd --permanent --zone=internal --add-source=192.168.1.0/24
firewall-cmd --reload
firewall-cmd --list-all --zone=public
```

- Always allow SSH before `ufw enable`; remember firewalld's `--permanent` vs runtime split [secondary: section 3 sources].

---

## 20. Appendix B — selection flowcharts (text form)

### 20.1 "Which remote-access VPN?"

1. Team of people needing access to servers/services, identity provider available -> **Tailscale** (hosted) or **Headscale** (self-hosted control) [secondary].
2. Need Layer-2 adjacency (broadcast, legacy discovery protocols) -> **ZeroTier** [secondary].
3. Simple point-to-point or site-to-site links, full config control -> **plain WireGuard** [secondary].
4. Must authenticate against RADIUS/AD with EAP, or interoperate with vendor appliances -> **strongSwan IKEv2** [secondary].
5. Egress firewall only allows TCP/443, or legacy client compatibility required -> **OpenVPN** (2.7 + DCO where possible) [secondary].
6. Self-hosted certificate-based mesh, no SaaS, own CA acceptable -> **Nebula** [secondary].

### 20.2 "Which firewall frontend?"

1. Debian/Ubuntu single-purpose host -> **UFW** [secondary].
2. RHEL-family or multi-interface host needing zones -> **firewalld** [secondary].
3. Complex policy, high-cardinality sets, atomic updates -> **raw nftables** [secondary].
4. Fleet-wide firewall policy as code -> nftables rulesets managed by Ansible (see Step 7C) [secondary].

### 20.3 "Which intrusion-response layering?"

1. Harden first: key-only SSH, no root login, default-deny firewall [secondary].
2. Add reactive banning: fail2ban (single host, simple) or CrowdSec (collaborative, multi-service) — not both on the same logs [secondary].
3. Add detection: Wazuh agents (host/FIM/vuln) and/or Security Onion (network/NSM) [secondary].
4. Centralize: forward alerts to the SIEM of record; keep custom rules/decoders in Git [secondary].

---

*End of Step 7G research file. Observation cutoff 2026-09-22.*
