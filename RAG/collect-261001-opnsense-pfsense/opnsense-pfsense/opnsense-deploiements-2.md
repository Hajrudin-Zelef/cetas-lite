---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/opnsense-deploiements-2
title: "OPNsense — 4 déploiements simulés (niveau 1 → hardcore)"
domain: opnsense-pfsense
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["datacenter", "exploit", "intel", "latency"]
source: docs/RAG/collect-261001-opnsense-pfsense/opnsense_deploiements.md
source_anchor: ""
source_lines: [208, 393]
sha256: 911775a05bed00699950e1cf793e1e5525ec1b6a00dc5ebd8f415df02278ead2
---

# OPNsense — 4 déploiements simulés (niveau 1 → hardcore)

## Vérifications
- [ ] Un invité (VLAN 20) navigue mais `ping 192.168.10.1` → **échec**.
- [ ] Un employé ouvre `https://intranet.pme.lan`.
- [ ] Depuis l'extérieur : `https://203.0.113.10` → site ; `http://`
      → rien (pas de règle 80).
- [ ] Télétravailleur WireGuard : handshake récent (`wg show`),
      accès intranet OK.
- [ ] **Live View** : voir les Block sur INVITES→EMPLOYES quand on teste.

## Pièges niveau 2
- Règle INVITES « Pass → any » au lieu de « Pass → !RFC1918 » :
  les invités voient tout le réseau interne.
- Oublier la règle firewall WAN pour WireGuard (le peer n'a jamais
  de handshake).
- AllowedIPs du peer mal configuré (`0.0.0.0/0` route tout le trafic
  du client via le VPN — voulu ou non ?).
- Le switch : trunk oublié = VLANs morts (toujours tester en
  untagged d'abord).

---

# NIVEAU 3 — Entreprise multi-sites (HA + multi-WAN + IPS)

## Objectif
Siège + 2 agences, 200 utilisateurs : **double FAI avec failover**,
**HA active/passive** (2 firewalls), tunnels **IPsec** vers les
agences, **IPS** actif, supervision.

## Topologie

```
   FAI-1 (fibre)              FAI-2 (4G/ADSL secours)
      │                          │
 ┌────┴─────┐              ┌──────┴─────┐
 │ FW-1     │◄─── sync ───►│ FW-2       │  CARP : .2/.3, VIP .1/.254
 │ (MASTER) │  + pfsync   │ (BACKUP)   │
 └────┬─────┘              └────────────┘
      │ trunk (VLANs 10/20/30/40 comme niveau 2 + VLAN 99 MGMT)
      │
   ┌──┴──┐        IPsec IKEv2         ┌────────┐
   │ LAN │◄──────────────────────────►│ Agence │ (OPNsense ou routeur IPsec)
   └──┬──┘   192.168.0.0/16 chiffré  └────────┘
      │
   (DMZ, serveurs, WiFi corporate…)
```

## Plan d'adressage (siège)
| Élément | IP |
|---|---|
| WAN1 (FAI-1) | 203.0.113.10/30, GW 203.0.113.9 |
| WAN2 (FAI-2) | DHCP (ou 198.51.100.50/29) |
| LAN VIP CARP | 192.168.10.254 (FW-1 .253, FW-2 .252) |
| MGMT VLAN 99 | 192.168.99.0/24 |
| Agence A | 192.168.100.0/24 via IPsec |
| Agence B | 192.168.110.0/24 via IPsec |

## Étapes (web UI)

### A. Haute disponibilité (à faire en premier)
1. Sur FW-1 et FW-2 : **même version OPNsense** (prérequis absolu).
2. **Firewall > Virtual IPs** : créer les VIP **CARP** (WAN, LAN,
   chaque VLAN) — même VHID, advskew `0` sur FW-1, `100` sur FW-2.
3. **System > High Availability** : activer sync XMLRPC vers le pair
   (IP, mot de passe, cocher : rules, aliases, NAT, DHCP, VPN…).
4. **Status** : vérifier MASTER/BACKUP sur les deux nœuds.
5. Tester : débrancher FW-1 → FW-2 passe MASTER (coupure < 10 s).

### B. Multi-WAN avec failover
6. **System > Gateways > Single** : GW_WAN1 (monitor `9.9.9.9`),
   GW_WAN2 (monitor `1.1.1.1`). « Disable Gateway Monitoring » =
   **jamais** en prod.
7. **System > Gateways > Group** : `GW_FAILOVER` — WAN1 Tier 1,
   WAN2 Tier 2, trigger `Packet Loss or High Latency`.
8. **Firewall > NAT > Outbound** : passer en **Hybrid**, règle
   manuelle : LAN → GW_FAILOVER (le NAT suit le groupe).
9. **Firewall > Rules > LAN** : règle Pass avec **Gateway =
   GW_FAILOVER** (sinon le trafic sort toujours par la default GW).
10. Tester : débrancher WAN1 → les sessions basculent sur WAN2
    (vérifier **System > Gateways** : WAN1 « down », groupe sur WAN2).

### C. IPsec vers les agences (IKEv2)
11. **VPN > IPsec > Tunnels** : Phase 1 IKEv2 (AES-256-GCM, DH 19,
    auth **certificats** — créez la CA dans System > Certificates),
    Phase 2 ESP (réseaux `192.168.0.0/16` ↔ `192.168.100.0/24`).
12. **Firewall > Rules > IPsec** : Pass IPsec → LAN (et inversement
    sur LAN → IPsec).
13. **Status Overview** : tunnel UP, trafic chiffré.
14. Répéter pour l'agence B.

### D. IPS (Suricata)
15. **Services > Intrusion Detection** : **Download** des règles
    (ET Open), activer sur **WAN** en mode **IDS** (détection).
16. Laisser tourner 1 semaine, analyser **Alerts**, désactiver les
    faux positifs (applis métiers).
17. Basculer en **IPS** (blocage) sur WAN.
18. Politique : bloquer les catégories `exploit`, `malware`,
    `trojan` ; alerter seulement sur `policy`, `info`.

### E. Supervision et finitions
19. **Reporting > NetFlow** : activer la capture sur LAN/VLANs
    (Insight : qui consomme quoi).
20. **Services > SNMP** : v3 authPriv pour Zabbix/Centreon.
21. **System > Settings > Logging** : syslog distant (serveur SIEM).
22. **System > Access** : 2FA pour tous les admins, comptes nominatifs.
23. **System > Configuration > Backups** : backup chiffré + envoi
    automatique (plugin Nextcloud/Google Drive).
24. **System > Firmware** : planifier la vérification mensuelle
    (System > Settings > Cron).

## Équivalents CLI

```shell
# CARP
ifconfig | grep -A2 carp
# Gateways
configctl gateways status 2>/dev/null || grep -A8 "<gateways>" /conf/config.xml
# IPsec
ipsec status
swanctl --list-sas
# Suricata
configctl suricata status
clog -f /var/log/suricata/eve.json | grep alert
# Sync HA
grep -A10 "<hasync>" /conf/config.xml
```

## Vérifications
- [ ] Failover WAN : coupure WAN1 → Internet maintenu via WAN2.
- [ ] Failover HA : extinction FW-1 → FW-2 MASTER, sessions maintenues
      (avec pfsync).
- [ ] Agences : `ping 192.168.100.1` depuis le siège via IPsec.
- [ ] Suricata : lancer un test (ex. : `curl` vers une URL de test
      EICAR-like) → alerte/block visible.
- [ ] Restauration : restaurer un backup sur FW-2 (lab) → config OK.

## Pièges niveau 3
- Versions différentes entre FW-1/FW-2 → sync silencieusement cassée.
- Règle LAN sans « Gateway = groupe » : le failover WAN ne s'applique
  pas au trafic LAN.
- IPsec + NAT : exclure le trafic IPsec du NAT outbound manuel
  (règle `no-nat` / ne pas NATer vers les réseaux distants).
- Suricata en IPS **avant** tuning = applis métiers coupées un lundi
  matin. Toujours IDS d'abord.
- Monitor IP des gateways = la passerelle FAI elle-même : si le FAI
  route mais perd Internet, le monitor reste « up ». Monitorer une IP
  **Internet** (9.9.9.9).

---

# NIVEAU HARDCORE — Datacenter / opérateur

## Objectif
Héberger des clients en datacenter : **routage BGP**, **IPv6 natif**,
**filtrage applicatif** (Zenarmor), **threat intel** (CrowdSec),
**QoS stricte**, **automatisation API**, durcissement maximal.
C'est le niveau « on ne dort que d'un œil ».

## Topologie (simplifiée)

```
        Transit-1 (BGP)         Transit-2 (BGP)
            │                       │
      ┌─────┴───────┐       ┌───────┴───────┐
      │  FW-A       │◄─────►│  FW-B         │  HA CARP + pfsync
      │  (FRR BGP)  │ sync  │  (FRR BGP)    │
      └─────┬───────┘       └───────┬───────┘
            │ trunk (VLANs clients, MGMT, DMZ, TRANSIT)
   ┌────────┼──────────┬───────────┼─────────┐
   │        │          │           │         │
CLIENT-A  CLIENT-B   DMZ-      MGMT     TRANSIT
VLAN 101  VLAN 102  MUTUALISE  VLAN 99  /30 publics
::/48      ::/48    (hébergé)
/24        /24
```

## Plan d'adressage (extrait)
| Élément | IPv4 | IPv6 |
|---|---|---|
| Transit-1 | 203.0.113.0/30 | 2001:db8:1::/64 |
| AS propre | AS 65001 (ou PI RIPE) | — |
| Client A | 198.51.100.0/24 | 2001:db8:100::/48 |
| Client B | 198.51.101.0/24 | 2001:db8:200::/48 |
| MGMT | 192.168.99.0/24 | — |

## Étapes

