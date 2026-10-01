---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/opnsense-deploiements-1
title: "OPNsense — 4 déploiements simulés (niveau 1 → hardcore)"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/opnsense_deploiements.md
source_anchor: ""
source_lines: [1, 207]
sha256: 66b2f391ad2af8259874f31d3cc57d450a181f82805581879db2885a09badcb2
---

# OPNsense — 4 déploiements simulés (niveau 1 → hardcore)

> Quatre scénarios complets et réalistes : topologie, plan d'adressage,
> configuration pas à pas (interface web **et** équivalents CLI),
> vérifications et pièges. Chaque niveau réutilise les acquis du
> précédent — lisez dans l'ordre.
>
> **Conventions** : les chemins web sont en anglais tels qu'affichés
> (`Firewall > Rules > LAN`). Les blocs CLI renvoient aux commandes du
> [guide CLI](opnsense_cli_guide.md).

---

# NIVEAU 1 — Petit bureau / domicile (SOHO)

## Objectif
Protéger un petit réseau : 1 box FAI, ~15 postes, WiFi. Internet
fonctionnel en 30 minutes, avec les bases de sécurité.

## Topologie

```
        ┌─────────┐
        │ Box FAI │ 192.168.0.1 (DHCP)
        └────┬────┘
             │ WAN (DHCP)
      ┌──────┴──────┐
      │  OPNsense   │
      └──────┬──────┘
             │ LAN 192.168.1.1/24
      ┌──────┴──────┐
      │  Switch    │── PC, imprimante, borne WiFi (AP mode)
      └────────────┘
```

## Plan d'adressage
| Réseau | Plage | Passerelle |
|---|---|---|
| LAN | 192.168.1.0/24 | 192.168.1.1 (OPNsense) |
| DHCP | 192.168.1.100–200 | — |

## Étapes (web UI)

1. **Installation** : boot sur l'ISO/USB, login `root`/`opnsense`,
   menu console → **1) Assign Interfaces** (em0=WAN, em1=LAN),
   puis **2)** : WAN en DHCP, LAN en `192.168.1.1/24`.
2. **Premier login web** : `https://192.168.1.1`, assistant :
   hostname `fw-maison`, DNS `9.9.9.9`, fuseau horaire, mot de passe
   root fort.
3. **System > Settings > Administration** : HTTPS uniquement,
   activer SSH (clés plus tard).
4. **Services > DHCPv4 > LAN** : Enable, Range `192.168.1.100–200`,
   DNS `192.168.1.1` (l'OPNsense lui-même).
5. **Services > Unbound DNS > General** : Enable, écoute sur LAN,
   **DNS over TLS** coché (forward TLS vers Quad9/Cloudflare),
   **Blocklists** : activer `ads` + `malware`.
6. **Firewall > Rules > LAN** : la règle « Default allow LAN to any »
   existe déjà → la garder pour l'instant (niveau 1).
7. **Firewall > Rules > WAN** : vérifier qu'il n'y a **aucune**
   règle Pass (par défaut : rien ne rentre — parfait).
8. **System > Settings > General** : NTP par défaut OK.
9. **System > Configuration > Backups** : télécharger le premier
   `config.xml`.

## Équivalents CLI

```shell
# Vérifier interfaces et DHCP
ifconfig
cat /var/dhcpd/etc/dhcpd.conf | head -20
# Règles actives
pfctl -sr | head -20
pfctl -sn
# Test DNS chiffré
drill www.example.com @127.0.0.1
# Backup
cp /conf/config.xml /root/niveau1-$(date +%Y%m%d).xml
```

## Vérifications
- [ ] Un PC obtient `192.168.1.x` en DHCP, navigue sur Internet.
- [ ] `drill pub.malware.exemple @192.168.1.1` → bloqué (blocklist).
- [ ] Depuis Internet (4G du téléphone) : `https://<ip-publique>` →
  **aucune réponse** (web UI non exposée).
- [ ] **Firewall > Log Files > Live View** : voir les blocages WAN
  (scans) en direct.

## Pièges niveau 1
- Box FAI en routeur (double NAT) : acceptable au niveau 1, mais
  mettre la box en **mode bridge** si possible (sinon port-forwards
  à faire 2 fois plus tard).
- Oublier le backup après chaque étape qui marche.

---

# NIVEAU 2 — PME (3 VLANs + VPN road-warrior)

## Objectif
Bureau de 40 personnes : séparer **employés**, **invités** et
**serveurs** ; accès VPN pour les télétravailleurs ; un serveur web
publié ; DNS interne.

## Topologie

```
                    ┌────────────┐
                    │    FAI     │ IP publique fixe 203.0.113.10
                    └─────┬──────┘
                          │ WAN (PPPoE ou statique)
                   ┌──────┴──────┐
                   │  OPNsense   │
                   └──┬──┬───┬───┘
                      │  │   │   trunk 802.1Q
              ┌───────┘  │   └───────┐
              │          │           │
        VLAN 10    VLAN 20     VLAN 30
   EMPLOYES 192.168.10.0/24  INVITES 192.168.20.0/24  SERVEURS 192.168.30.0/24
   .1 = OPNsense             .1 = OPNsense            .1 = OPNsense
                                                  .10 = serveur web
                                                  .11 = NAS
```

## Plan d'adressage
| VLAN | Réseau | DHCP | Notes |
|---|---|---|---|
| 10 EMPLOYES | 192.168.10.0/24 | .100–200 | accès Internet + serveurs |
| 20 INVITES | 192.168.20.0/24 | .100–200 | Internet seul, isolé |
| 30 SERVEURS | 192.168.30.0/24 | réservations MAC | pas d'Internet direct* |

\* Les serveurs sortent via un proxy/updates autorisés uniquement
(voir règles).

## Étapes (web UI)

### A. VLANs et interfaces
1. **Interfaces > Other Types > VLAN** : créer VLAN 10, 20, 30 sur
   l'interface physique LAN (parent `em1`).
2. **Interfaces > Assignments** : assigner `VLAN 10 → EMPLOYES`,
   `VLAN 20 → INVITES`, `VLAN 30 → SERVEURS`, Enable + IP statique
   `.1/24` sur chacune.
3. Switch : trunk 802.1Q vers OPNsense, ports access dans les bons
   VLANs (côté switch, hors OPNsense).

### B. DHCP + DNS
4. **Services > DHCPv4** : activer sur les 3 VLANs (ranges adaptées).
   **Static Mappings** sur SERVEURS : `.10` (web), `.11` (NAS).
5. **Services > Unbound DNS > Overrides > Host Overrides** :
   `intranet.pme.lan → 192.168.30.10`, `nas.pme.lan → 192.168.30.11`.

### C. Règles firewall (le cœur du niveau 2)
6. **Firewall > Aliases** : créer `VLAN_EMPLOYES` (192.168.10.0/24),
   `VLAN_INVITES`, `VLAN_SERVEURS`, `SRV_WEB` (192.168.30.10),
   `PORTS_WEB` (80, 443).
7. **Firewall > Rules > INVITES** :
   - Pass INVITES → `! RFC1918` (alias « réseaux privés » : tout
     sauf les réseaux internes) ports 80/443/53 — Internet seul.
   - (implicite : pas d'accès aux autres VLANs.)
8. **Firewall > Rules > EMPLOYES** :
   - Pass EMPLOYES → SERVEURS ports 80/443 (intranet).
   - Pass EMPLOYES → any (Internet).
9. **Firewall > Rules > SERVEURS** :
   - Pass SERVEURS → any port 80/443 (mises à jour uniquement).
   - Block SERVEURS → LAN nets (log).
10. **Firewall > Rules > WAN** : aucune règle Pass pour l'instant.

### D. Publication du serveur web
11. **Firewall > NAT > Port Forward** : WAN, TCP, dest `WAN address`,
    port 443 → `192.168.30.10:443`, cocher « Associated filter rule »
    (crée la règle WAN automatiquement).
12. Tester depuis l'extérieur : `https://203.0.113.10` → site OK.

### E. VPN road-warrior (WireGuard)
13. **VPN > WireGuard > Instances** : Add, port `51820`, générer les clés.
14. **VPN > WireGuard > Peers** : un peer par télétravailleur
    (clé publique du client, AllowedIPs `10.8.0.x/32`, keepalive 25).
15. **Interfaces > Assignments** : assigner `wg0` → interface `WG`
    (pour les règles firewall).
16. **Firewall > Rules > WG** : Pass WG → EMPLOYES + SERVEURS
    (selon profil).
17. **Firewall > Rules > WAN** : Pass UDP 51820 (source any).
18. Distribuer les configs clients (QR code / fichier).

### F. Finitions
19. **System > Access > Users** : créer un compte `admin` (désactiver
    le login web `root` direct si souhaité), activer **2FA TOTP**.
20. **System > Firmware** : vérifier les mises à jour.
21. **System > Configuration > Backups** : télécharger + activer
    l'historique (déjà actif par défaut).

## Équivalents CLI

```shell
# VLANs
ifconfig | grep vlan
# Règles par interface
pfctl -sr | grep -A2 "label"
# NAT port-forward actif
pfctl -sn | grep 443
# WireGuard
wg show
wg show wg0 latest-handshakes
# États d'un invité
pfctl -ss | grep 192.168.20.
# Backup
cp /conf/config.xml /root/niveau2-$(date +%Y%m%d).xml
```

