---
id: collect-261001-mikrotik/mikrotik/deploiement-entreprise-ubiquiti-mikrotik-4
title: "Déploiement entreprise — Ubiquiti & MikroTik (référence complète)"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["preemption"]
source: docs/RAG/collect-261001-mikrotik/deploiement_entreprise_ubiquiti_mikrotik.md
source_anchor: ""
source_lines: [497, 680]
sha256: a981d7a3611b46c25d21d76c9d36c7b81927bbc4fc8c016a6a3eeed2256d56e3
---

# Déploiement entreprise — Ubiquiti & MikroTik (référence complète)

```shell
MT> /interface wireguard add name=wg0 listen-port=51820 private-key="<cle-privee>"
MT> /ip address add address=10.8.0.1/24 interface=wg0
MT> /interface wireguard peers add interface=wg0 public-key="<cle-publique-client>" allowed-address=10.8.0.2/32 persistent-keepalive=25s comment="Teletravailleur-1"
MT> /ip firewall filter add chain=input action=accept protocol=udp dst-port=51820 comment="WireGuard"
MT> /ip firewall filter add chain=forward action=accept in-interface=wg0 out-interface=vlan10-data
MT> /interface wireguard print
MT> /interface wireguard peers print
! --- IPsec site-à-site (IKEv2) ---
MT> /ip ipsec profile add name=prof-agence hash-algorithm=sha256 enc-algorithm=aes-256 dh-group=modp2048
MT> /ip ipsec peer add name=peer-agence address=198.51.100.2/32 profile=prof-agence auth-method=pre-shared-key secret="CleIPsec"
MT> /ip ipsec proposal add name=prop-agence auth-algorithms=sha256 enc-algorithms=aes-256-cbc
MT> /ip ipsec policy add src-address=192.168.0.0/16 dst-address=192.168.100.0/24 peer=peer-agence proposal=prop-agence tunnel=yes
MT> /ip ipsec active-peers print
```

### 6.8 VRRP (HA)

```shell
MT> /interface vrrp add name=vrrp-data interface=vlan10-data vrid=10 priority=200 preemption-mode=yes
MT> /ip address add address=192.168.10.254/24 interface=vrrp-data
MT> /interface vrrp print
! (second nœud : priority=100, même vrid, sync via /system backup + export)
```

### 6.9 QoS : simple queues + queue tree (PCQ)

```shell
! --- Simple : limiter un client ---
MT> /queue simple add name=clientA target=192.168.10.50/32 max-limit=50M/50M
! --- Queue tree : prioriser la VoIP (avec mangle) ---
MT> /ip firewall mangle add chain=forward action=mark-packet new-packet-mark=voip dscp=46 comment="Marque VoIP"
MT> /queue tree add name=q-voip parent=global packet-mark=voip priority=1 max-limit=30M
MT> /queue tree add name=q-defaut parent=global priority=8 max-limit=1000M
MT> /queue tree print
! --- PCQ : équité entre N utilisateurs ---
MT> /queue type add name=pcq-down kind=pcq pcq-classifier=dst-address pcq-rate=20M
MT> /queue simple add name=partage target=192.168.30.0/24 queue=pcq-down/pcq-down max-limit=100M/100M
```

### 6.10 WiFi : CAPsMAN (contrôleur intégré)

```shell
MT> /interface wifi capsman set enabled=yes
MT> /interface wifi provisioning add master-configuration=cfg-corp slave-configurations=cfg-invites
MT> /interface wifi configuration add name=cfg-corp ssid=CORP security.authentication-types=wpa2-psk,wpa3-psk security.passphrase="MotDePasseFort" datapath.vlan-id=10
MT> /interface wifi configuration add name=cfg-invites ssid=INVITES security.authentication-types=wpa2-psk security.passphrase="Invite2026" datapath.vlan-id=30
MT> /interface wifi capsman remote-cap add address=192.168.99.0/24
MT> /interface wifi registration-table print
```

### 6.11 Outils de diagnostic MikroTik

```shell
MT> /ping 8.8.8.8 count=5
MT> /tool traceroute 8.8.8.8
MT> /tool torch interface=ether1-WAN
MT> /tool sniffer quick interface=vlan10-data ip-address=192.168.10.50
MT> /tool profile
MT> /interface print stats
MT> /ip firewall connection print count-only
MT> /log print where topics~"firewall"
MT> /system resource print
MT> /system health print
```

### 6.12 Scripts & automatisation

```shell
MT> /system script add name=backup-quotidien source="/system backup save name=auto-backup; /export file=auto-export"
MT> /system scheduler add name=sched-backup on-event=backup-quotidien start-time=03:00:00 interval=1d
MT> /tool netwatch add host=9.9.9.9 interval=30s up-script="/log info \"FAI-1 OK\"" down-script="/log warning \"FAI-1 DOWN\""
MT> /system script print
```

---

## 7. WiFi entreprise — comparatif

| Aspect | UniFi (contrôleur) | MikroTik CAPsMAN |
|---|---|---|
| Gestion | GUI centralisée, adoption auto | CLI `/interface wifi` |
| SSID/VLAN | Profils réseau en 2 clics | `datapath.vlan-id` |
| Portail invité | Hotspot intégré + vouchers | Hotspot `/ip hotspot` |
| Roaming | 802.11r/k/v auto | 802.11r à configurer |
| Idéal pour | Simplicité, sites multi-AP | Budget, contrôle fin |

---

## 8. Sécurité : durcissement

### Checklist commune
- [ ] Services inutiles désactivés (telnet/ftp/www/api ouverts au monde).
- [ ] Admin par comptes nominatifs, mots de passe forts, SSH par clés.
- [ ] Firewall input : **drop par défaut**, n'ouvrir que MGMT + services.
- [ ] WinBox : restreint au MGMT (`/ip service set winbox address=…`),
      jamais exposé sur WAN (vecteur d'attaque historique).
- [ ] Mises à jour : RouterOS **stable** (pas de testing en prod),
      EdgeOS : vérifier les CVE.
- [ ] Backups chiffrés + export, hors-site, test de restore.

### MikroTik — protections spécifiques
```shell
MT> /ip firewall filter add chain=input action=add-src-to-address-list address-list=port-scanners address-list-timeout=2w in-interface=ether1-WAN psd=21,3s,3,1 comment="Detecte scan"
MT> /ip firewall filter add chain=input action=drop src-address-list=port-scanners
MT> /ip ssh set strong-crypto=yes
MT> /ip neighbor discovery-settings set discover-interface-list=MGMT
```

---

## 9. Supervision

| Outil | Ubiquiti | MikroTik |
|---|---|---|
| SNMP | UniFi : limité ; EdgeRouter : complet | SNMP complet (v1/2c/3) |
| Syslog | EdgeOS/UniFi → SIEM | `/system logging` → SIEM |
| Natif | UniFi Insights / DPI | **The Dude** (cartographie, sondes) |
| NetFlow | EdgeRouter : `set system flow-accounting` | Traffic Flow `/ip traffic-flow` → collecteur |

```shell
! EdgeRouter — flow accounting
UBNT> configure
UBNT# set system flow-accounting interface eth0
UBNT# set system flow-accounting netflow server 192.168.99.70 port 9996
UBNT# set system flow-accounting netflow version 9
UBNT# commit; save
```

```shell
! MikroTik — Traffic Flow
MT> /ip traffic-flow set enabled=yes interfaces=ether1-WAN
MT> /ip traffic-flow target add dst-address=192.168.99.70 port=9996 version=9
```

---

## 10. Interopérabilité Ubiquiti ↔ MikroTik

| Protocole | État | Notes |
|---|---|---|
| Trunk 802.1Q | ✅ | PVID/native cohérents des 2 côtés |
| LACP 802.3ad | ✅ | `mode 802.3ad` ↔ `mode=802.3ad` |
| RSTP | ✅ | Bridge MikroTik : `stp=yes`, même priority root |
| OSPF | ✅ | Mêmes timers, MTU identique |
| BGP | ✅ | eBGP standard |
| VRRP | ✅ | Même VRID ; préemption cohérente |
| WireGuard | ✅ | Clés/AllowedIPs miroir, keepalive 25 s |
| IPsec IKEv2 | ✅ | Proposals identiques (AES-256/SHA-256/DH14) |
| DHCP | ⚠️ | Un seul serveur par VLAN (sinon conflits) |

---

## 11. Checklist de mise en production

### Avant le jour J
- [ ] Configs exportées : EdgeRouter `show configuration` ; MikroTik
      `/export` + `/system backup save` — hors-site.
- [ ] Plan de rollback écrit (Netinstall MikroTik / reset EdgeRouter
      en dernier recours).
- [ ] Accès console/SSH testés sur chaque équipement.
- [ ] Firmware : version **stable** validée (pas de release candidate).
- [ ] Étiquetage des câbles et ports.

### Recette
- [ ] Chaque VLAN : DHCP, passerelle, Internet, isolation conformes.
- [ ] NAT : `show nat translations` / `/ip firewall connection` OK.
- [ ] Failover WAN : débrancher FAI-1 → bascule, rebrancher → retour.
- [ ] VRRP : extinction du master → VIP sur le backup < 10 s.
- [ ] VPN : WireGuard handshake récent, IPsec SA établies, trafic OK.
- [ ] WiFi : roaming entre AP, débit, portail invité.
- [ ] QoS : appel VoIP clair pendant saturation du lien.
- [ ] Firewall : scan nmap depuis WAN → bloqué + loggé.
- [ ] Supervision : SNMP/syslog/NetFlow reçus par les collecteurs.
- [ ] NTP synchronisé partout.

### Ordre de déploiement
1. Edge/routeurs → 2. Switchs/VLANs → 3. WiFi → 4. Firewall/NAT →
   5. VPN → 6. QoS → 7. Supervision. Valider chaque couche avant la suivante.

---

## 12. Maintenance

