---
id: collect-261001-rattrapage/rattrapage/opnsense-cli-guide-3
title: "OPNsense — Guide d'administration en CLI (sans l'interface web)"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/opnsense_cli_guide.md
source_anchor: ""
source_lines: [552, 625]
sha256: e61f485ea72b0121556622f73ccd06775daff87aa5c4e488337a7ac25df47af5
---

# Sync XMLRPC : section <hasync> de config.xml
grep -A10 "<hasync>" /conf/config.xml
```

- Les deux nœuds doivent avoir la **même version** d'OPNsense.
- Testez le failover en heures creuses : coupez le MASTER et
  vérifiez que le BACKUP passe MASTER (`ifconfig carp0`).

---

## 24. Dépannage express en CLI

```shell
# === Pas d'Internet depuis le LAN ===
ifconfig                          # interfaces UP ?
netstat -rn | grep default        # passerelle par défaut ?
drill www.google.com @127.0.0.1   # DNS local OK ?
pfctl -sr | head -30              # règles chargées ?
pfctl -sn                         # NAT présent ?
clog -f /var/log/filter.log       # qui est bloqué, en direct ?
# Test depuis OPNsense lui-même :
ping -c3 8.8.8.8
traceroute 8.8.8.8

# === Un port-forward ne marche pas ===
pfctl -sn | grep <port>           # règle NAT présente ?
pfctl -sr | grep <port>           # règle firewall WAN présente ?
pfctl -ss | grep <ip_serveur>     # états créés ?
# Le serveur a-t-il OPNsense comme passerelle ? (arp -a)

# === VPN IPsec qui ne monte pas ===
ipsec status                      # SA visibles ?
clog -f /var/log/ipsec.log        # erreurs phase 1 / phase 2 ?
# PSK identique ? Proposals compatibles ? UDP 500/4500 autorisés ?

# === WireGuard silencieux ===
wg show wg0 latest-handshakes     # handshake récent ?
# Non -> réseau/clé/endpoint. Oui mais pas de trafic -> AllowedIPs / firewall.

# === Web UI inaccessible ===
configctl webgui restart
sockstat -4 | grep :443            # le port écoute ?
# En dernier recours : option 11 (reload services) ou reboot.
```

---

## 25. Fiche réflexe : les commandes à connaître par cœur

| Besoin | Commande |
|---|---|
| Config complète | `cat /conf/config.xml` (+ backup avant modif) |
| Appliquer la conf | `configctl <service> restart` / `configctl filter reload` |
| Règles firewall actives | `pfctl -sr` |
| NAT actif | `pfctl -sn` |
| Connexions actives | `pfctl -ss` |
| Tuer les états d'une IP | `pfctl -k <ip>` |
| Ajouter une règle simple | `easyrule block\|pass ...` |
| Table de routage | `netstat -rn` |
| Interfaces | `ifconfig` |
| DHCP baux | `cat /var/dhcpd/var/db/dhcpd.leases` |
| DNS test | `drill <nom> @127.0.0.1` |
| IPsec | `ipsec status` / `swanctl --list-sas` |
| WireGuard | `wg show` |
| Logs firewall live | `clog -f /var/log/filter.log` |
| Logs système | `clog /var/log/system.log` |
| Mise à jour | `opnsense-update -t opnsense` |
| Version | `cat /usr/local/opnsense/version/opnsense.version` |

---

*Fin du guide. Règle d'or OPNsense en CLI : `config.xml` est la vérité,
`configctl` applique, `pfctl` vérifie. Sauvegardez avant chaque
modification manuelle, et ne désactivez jamais `pf` à distance.*
