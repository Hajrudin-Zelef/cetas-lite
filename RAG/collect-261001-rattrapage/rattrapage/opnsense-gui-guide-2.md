---
id: collect-261001-rattrapage/rattrapage/opnsense-gui-guide-2
title: "OPNsense — Guide ultra-complet de l'interface web"
domain: rattrapage
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-rattrapage/opnsense_gui_guide.md
source_anchor: ""
source_lines: [198, 410]
sha256: 2c1f6a28ebe9010a059800becded1ac8592d4a6af07a41278a0874f7babf462e
---

# OPNsense — Guide ultra-complet de l'interface web

### Services > DHCP Relay
- Relayer vers un serveur DHCP central quand plusieurs VLANs.

### Services > Unbound DNS > General
- Activer, interfaces d'écoute, **DNS over TLS** (forward vers
  9.9.9.9 / 1.1.1.1 en DoT : confidentialité des requêtes).
- **Blocklists** : activer les listes (ads, trackers, malware) —
  l'ad-blocking réseau en 2 clics.
- **Advanced** : cache, ECS, qname minimisation.

### Services > Unbound DNS > Overrides
- **Host Overrides** : noms internes (`intranet.lan → 192.168.1.10`).
- **Domain Overrides** : forwarder conditionnel par domaine
  (ex. : `corp.lan` → DNS Active Directory).

### Services > Unbound DNS > Blocklists / Query / Stats
- Gérer les listes, tester une résolution, voir les stats.

---

## 6. VPN

### VPN > IPsec
- **Tunnels** : Phase 1 (IKEv2 recommandé : AES-GCM, DH 14/19/20,
  auth PSK ou certificats) puis Phase 2 (ESP, réseaux locaux/distants).
- **Mobile Clients** : road-warriors (EAP-MSCHAPv2, certificats).
- **Status Overview** : tunnels UP/DOWN en direct.
- **Security Associations** : détail des SA.
- **Pre-Shared Keys / Advanced**.

### VPN > OpenVPN
- **Servers** : mode Remote Access (SSL/TLS + Auth) ou Peer to Peer.
  Assistant (wizard) disponible pour un serveur road-warrior standard.
- **Clients** : tunnels site-à-site ou client vers fournisseur.
- **Client Export** : génère les `.ovpn` prêts à importer (avec
  certificats) — très pratique.
- **Status** : sessions connectées, kill session.

### VPN > WireGuard
- **Instances** : créer (port d'écoute, clés générées en 1 clic).
- **Peers** : endpoint, AllowedIPs, keepalive (25 s derrière un NAT).
- **Status** : handshakes et volumes — le diagnostic en un coup d'œil.

> Rappel valable pour les 3 : une règle **Firewall > Rules**
> (WAN pour IPsec/WireGuard/OpenVPN, + l'interface VPN elle-même)
> est obligatoire — le VPN ne bypass pas le firewall.

---

## 7. IDS/IPS et sécurité avancée

### Services > Intrusion Detection (Suricata)
- **Download** : télécharger les jeux de règles (ET Open, OTX…).
- **Policy** : activer en mode **IDS (détection)** d'abord, analyser
  les alertes, puis basculer en **IPS (blocage)** sur les interfaces
  WAN (ou LAN selon stratégie).
- **Rules** : activer/désactiver par catégorie, seuils.
- **Alerts** : les alertes avec payload — le tuning se fait ici
  (désactiver les faux positifs un par un).

### Plugins de sécurité (à installer si besoin)
- **Zenarmor (Sensei)** : filtrage applicatif/web avancé, contrôle
  par catégories, rapports par utilisateur.
- **CrowdSec** : IPS collaboratif (bannit les IP agressives vues
  par la communauté).
- **ACME Client** : certificats Let's Encrypt automatiques
  (web UI, VPN).

---

## 8. Routage

### System > Routes > Configuration
- Routes statiques IPv4/IPv6 (destination, passerelle).
- « Disable all packet filtering » n'existe pas ici : le firewall
  reste actif, pensez aux règles.

### System > Routes > Status
- Table effective.

### Plugin FRR (à installer)
- OSPF, BGP, RIP via **Routing > FRR** avec VTYSH intégré.
  Pour les réseaux dynamiques multi-sites.

---

## 9. Services réseau

### Services > NTP
- **General** : serveurs (pool.ntp.org par défaut), interfaces d'écoute.
  Une horloge juste = IPsec, logs et 2FA fiables.

### Services > SNMP
- Activer l'agent, communities (v2c) ou **v3** (authPriv recommandé),
  modules MIB-II.

### Services > Captive Portal
- Portail captif (hôtellerie, invités) : zones, page de login,
  vouchers, limites de bande passante.

### Services > Web Proxy (Squid, plugin)
- Proxy transparent ou explicite, cache, filtrage SSL bump
  (avec CA déployée sur les postes), antivirus ICAP (ClamAV).

### Reporting > NetFlow / Insight
- **Insight** : historique du trafic par hôte/service (qui consomme
  quoi) — activez la capture sur les interfaces internes.
- Données précieuses pour le dimensionnement et l'audit.

### Services > Syslog (destinations)
- Déjà vu en §2 : cible distante, TLS.

---

## 10. Utilisateurs, accès, 2FA

- **System > Access > Users** : créer/modifier, groupes, shell,
  clés SSH, **OTP (TOTP)** : scannez le QR code avec votre appli
  d'authentification.
- **Groups** : privilèges par pages (ex. : groupe « support » avec
  accès Diagnostics + Logs seulement).
- **Servers** : brancher LDAP/AD ou RADIUS pour centraliser les logins
  admin et VPN.

---

## 11. Diagnostics (le menu à connaître)

### Interfaces > Diagnostics
- **Ping**, **Traceroute**, **DNS Lookup**, **Test Port** :
  avec choix de l'interface **source** (crucial en multi-WAN :
  tester via WAN1 puis WAN2).
- **Packet Capture** : capture avec filtres (interface, host, port),
  **téléchargeable en .pcap** (ouverture dans Wireshark).
  L'équivalent web de tcpdump, sans SSH.
- **ARP/NDP Tables**, **Routes**.

### Firewall > Diagnostics (§4)
- States, pfTables, Statistics, Sockets.

### System > Diagnostics
- **Activity** (top), **System Activity**, **Disk Usage**.

### VPN > Diagnostics
- Status IPsec/OpenVPN/WireGuard (§6).

---

## 12. Sauvegarde, historique, restauration

### System > Configuration > Backups
- **Download** : télécharger `config.xml` (chiffrable par mot de passe).
- **Restore** : restaurer un fichier.
- **History** : chaque « Apply » crée une révision horodatée avec
  **diff** — cliquez pour voir ce qui a changé, restaurez en 1 clic.
  C'est le « git » de votre firewall : en cas de bêtise, revenez en
  arrière sans stress.
- **Google Drive / Nextcloud** (plugins) : backups automatiques
  chiffrés hors-site.

---

## 13. Mises à jour via l'interface

### System > Firmware > Status / Updates
- **Check for updates** : liste des correctifs (type : security,
  bugfix, major).
- **Update** : installe (redémarre les services, rarement le système
  complet sauf mise à jour majeure).
- Lisez le changelog affiché avant une **mise à jour majeure**
  (ex. 24.x → 25.x) : plugins à vérifier, backup avant.
- **System > Firmware > Audit** : vérifie l'intégrité des fichiers
  (détection de corruption / compromission).

---

## 14. Haute disponibilité via l'interface

1. **Firewall > Virtual IPs** : créer les VIP **CARP** (une par
   interface concernée : WAN, LAN…), même VHID des deux côtés,
   advskew 0 (master) / 100 (backup).
2. **System > High Availability** : activer la sync XMLRPC
   (IP du pair, mot de passe, éléments à synchroniser).
3. **Firewall > Virtual IPs > Status** : vérifier MASTER/BACKUP.
4. Test : débranchez le master → les VIP passent en MASTER sur le
   second en quelques secondes (coupure brève des états pf :
   activez « state synchronization » / pfsync si besoin de continuité
   totale).

---

## 15. Durcissement (hardening) recommandé

- [ ] Mot de passe root fort + **2FA TOTP** (System > Access).
- [ ] Web UI : HTTPS uniquement, écoute sur **LAN uniquement**,
  port non standard (ex. 8443).
- [ ] SSH : clés uniquement (pas de mot de passe), port non standard.
- [ ] WAN : « Block private networks » + « Block bogon networks »
  cochés ; **aucune règle Pass sur WAN** sauf VPN/port-forwards
  nécessaires.
- [ ] Règles LAN : principe du moindre privilège (pas de « allow all »
  permanent si possible) ; log sur les règles sensibles.
- [ ] Mises à jour firmware : vérifier mensuellement.
- [ ] Backups chiffrés automatiques hors-site (History + plugin cloud).
- [ ] Suricata en IPS sur WAN après tuning.
- [ ] DNS over TLS + blocklists (anti-malware au DNS).
- [ ] Désactiver les services inutilisés (chaque service = surface
  d'attaque).

---

## 16. Dépannage via l'interface (méthode)

