---
id: etape7-phaseg-linuxnet-vpn/00-linuxnet-vpn/19-appendix-a-annotated-configuration-references
title: "19. Appendix A — annotated configuration references"
domain: step-7g-linux-networking-and-access-nat-firewalls-ssh-vpns-a
role: deep-dive
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/etape7_phaseG_linuxnet_vpn.md
source_anchor: ""
source_lines: [619, 676]
section: "Step 7G — Linux Networking and Access: NAT, Firewalls, SSH, VPNs, and Security Monitoring"
sha256: 8d5400511903c6cdb06dc6642f22ca16d04b8a23298255a1b3870be4de192c6f
---

# 19. Appendix A — annotated configuration references

## 19. Appendix A — annotated configuration references

All examples below are operational patterns consolidated from the sources cited in their sections; adapt interface names, addresses, and ports before use.

### 19.1 nftables gateway with NAT, filtering, and rate-limited logging

```nft
#!/usr/sbin/nft -f
flush ruleset

table inet filter {
  set ssh_allow { type ipv4_addr; flags timeout; }
  chain input {
    type filter hook input priority 0; policy drop;
    ct state established,related accept
    iif "lo" accept
    icmp type echo-request limit rate 5/second accept
    tcp dport 22 ct state new add @ssh_allow { ip saddr } accept
    log prefix "nft-drop-in: " limit rate 10/minute
  }
  chain forward {
    type filter hook forward priority 0; policy drop;
    ct state established,related accept
    iifname "lan0" oifname "eth0" accept
  }
  chain output { type filter hook output priority 0; policy accept; }
}

table inet nat {
  chain prerouting {
    type nat hook prerouting priority dstnat; policy accept;
    tcp dport 8443 dnat to 192.168.1.20:443
  }
  chain postrouting {
    type nat hook postrouting priority srcnat; policy accept;
    ip saddr 192.168.1.0/24 oifname "eth0" masquerade
  }
}
```

- Dynamic sets with `flags timeout` are the native nftables replacement for temporary fail2ban-style bans [secondary: section 2 sources].
- Rate-limited logging (`limit rate`) prevents log floods on internet-facing chains [secondary: section 2 sources].

### 19.2 fail2ban sshd jail

```ini
[sshd]
enabled  = true
port     = ssh
filter   = sshd
logpath  = /var/log/auth.log
maxretry = 5
findtime = 600
bantime  = 3600
```

- Tune `maxretry`/`findtime`/`bantime` to the threat model; whitelist admin and monitoring IPs via `ignoreip` before enabling [secondary: section 5 sources].

