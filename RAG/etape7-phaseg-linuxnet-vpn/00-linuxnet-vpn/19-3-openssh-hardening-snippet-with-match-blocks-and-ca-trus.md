---
id: etape7-phaseg-linuxnet-vpn/00-linuxnet-vpn/19-3-openssh-hardening-snippet-with-match-blocks-and-ca-trus
title: "19.3 OpenSSH hardening snippet with Match blocks and CA trust"
domain: step-7g-linux-networking-and-access-nat-firewalls-ssh-vpns-a
role: deep-dive
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/etape7_phaseG_linuxnet_vpn.md
source_anchor: ""
source_lines: [677, 750]
section: "Step 7G — Linux Networking and Access: NAT, Firewalls, SSH, VPNs, and Security Monitoring"
sha256: 906f3af8790364b84d22a2c765ed7b13235f59cedda700f5762add77f2ad2641
---

# 19.3 OpenSSH hardening snippet with Match blocks and CA trust

### 19.3 OpenSSH hardening snippet with Match blocks and CA trust

```sshd_config
PasswordAuthentication no
PermitRootLogin no
PubkeyAuthentication yes
MaxAuthTries 3
LoginGraceTime 60
X11Forwarding no
ClientAliveInterval 300
ClientAliveCountMax 2
TrustedUserCAKeys /etc/ssh/ca_user_keys.pub

Match Address 100.64.0.0/10
    PasswordAuthentication no
    AllowUsers deploy
```

- `TrustedUserCAKeys` enables CA-signed certificates alongside `authorized_keys` unless `AuthorizedKeysFile none` is set [secondary: section 4 sources].
- `Match` blocks scope stricter policy to tailnet/VPN source ranges [secondary].

### 19.4 WireGuard site-to-site (two gateways)

Site A (`wg0`, 10.200.0.1/24, peer B at 203.0.113.10):

```ini
[Interface]
PrivateKey = <siteA-private>
Address = 10.200.0.1/24
ListenPort = 51820
PostUp = nft add rule inet filter forward iifname "wg0" accept
PostDown = nft delete rule inet filter forward iifname "wg0" accept

[Peer]
PublicKey = <siteB-public>
Endpoint = 203.0.113.10:51820
AllowedIPs = 10.200.0.2/32, 192.168.20.0/24
PersistentKeepalive = 25
```

- `AllowedIPs` is simultaneously the routing table and the access-control list (cryptokey routing); `PersistentKeepalive` keeps NAT mappings alive [secondary: section 7 sources].
- Mirror the stanza on Site B with reversed addresses; add static routes or a routing daemon for the remote LANs behind each gateway [secondary].

### 19.5 strongSwan swanctl road-warrior sketch (IKEv2 + EAP)

```swanctl
connections {
  rw {
    version = 2
    pools = [rw-pool]
    send_cert = never
    local {
      auth = pubkey
      certs = [gateway-cert.pem]
      id = vpn.example.com
    }
    remote {
      auth = eap-mschapv2
      eap_id = %any
    }
    children {
      net {
        local_ts = [10.10.0.0/24]
        inactivity = 3600
      }
    }
  }
}
```

- `swanctl` configuration is the current interface on strongSwan 6.x; legacy `ipsec.conf` is deprecated [secondary: section 10 sources].
- Road-warrior clients require UDP 500 (IKE) and UDP 4500 (NAT-T) reachable on the gateway [secondary: section 10 sources].
- Prefer EAP-TLS or certificate-based client auth over EAP-MSCHAPv2 where the client fleet supports it [secondary].

