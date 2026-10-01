---
id: collect-261001-rattrapage/rattrapage/dhcp-windows-guide-17
title: "Guide technique ultra-complet : DHCP sous Windows Server en entreprise"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: ["2026-09-27"]
keywords: ["agent"]
source: docs/RAG/collect-261001-rattrapage/dhcp_windows_guide.md
source_anchor: ""
source_lines: [2642, 2707]
sha256: f1b507062b32bdcbfe2680e802e3361a6f54dc3492955fc03f7a8bd02c7f3e52
---

# Guide technique ultra-complet : DHCP sous Windows Server en entreprise

**Q7. Un copieur (MFP) doit garder une IP fixe. Quelle méthode recommander et pourquoi pas une IP statique configurée sur le copieur ?**
> R : Réservation DHCP par adresse MAC : IP fixe garantie + gestion centralisée (passerelle/DNS modifiables sans toucher le copieur) + traçabilité (inventaire exportable). L'IP statique locale oblige à intervenir sur chaque copieur à chaque changement réseau et n'est pas visible dans l'inventaire DHCP.

**Q8. Qu'est-ce que le DHCP snooping et sur quel équipement se configure-t-il ?**
> R : Fonction des switchs (pas de Windows) : seuls les ports "trusted" (vers les serveurs DHCP légitimes) peuvent émettre des OFFER/ACK ; les ports utilisateurs "untrusted" ne peuvent qu'émettre des requêtes client. Bloque les rogue DHCP au niveau 2. Complété par le rate limiting anti-starvation.

**Q9. Quelles informations ne sont PAS répliquées par le basculement DHCP natif ?**
> R : Les options de niveau serveur, les filtres MAC (Allow/Deny), les classes personnalisées, la configuration DNS dynamique (compte dédié), les paramètres d'audit. À synchroniser manuellement (script section 63).

**Q10. Un poste obtient une IP mais pas de DNS interne, alors que l'option 006 est configurée sur l'étendue. Où chercher ?**
> R : Dans l'ordre de précédence inverse (du plus spécifique au plus général) : 1. Une stratégie applicable à ce client qui redéfinit 006 — 2. Une réservation avec option 006 spécifique — 3. Vérifier que l'option d'étendue est bien appliquée (Get-DhcpServerv4OptionValue) — 4. Côté client : ipconfig /renew après /release (bail ancien conservé).

## A5. Les 15 erreurs classiques (et comment les éviter)

| # | Erreur | Conséquence | Prévention |
|---|--------|-------------|------------|
| 1 | IP statiques dans la plage DHCP sans exclusion | Conflits d'adresses | Exclusions systématiques (§17) |
| 2 | Un seul serveur DHCP en production | Panne = plus de renouvellements | Failover obligatoire (§59) |
| 3 | IP helper vers un seul serveur | VLAN sans DHCP si panne | Helper vers les 2 (§65) |
| 4 | DNS public (8.8.8.8) dans l'option 006 | Échec d'ouverture de session AD | DNS AD uniquement (§34) |
| 5 | Options 066/067 + WDS sur le même serveur | Conflit port UDP 67 | Options DHCP 060 côté WDS (§37) |
| 6 | Deux DHCP indépendants, même étendue | Baux divergents, conflits | Failover natif (§59) |
| 7 | Bail de 8 jours sur Wi-Fi invités | Épuisement de l'étendue | Bail 4 h (§18) |
| 8 | Pas de compte DNS dédié | Échecs de mise à jour DNS, fantômes | svc-dhcp-dns + DnsUpdateProxy (§71) |
| 9 | Filtre Allow activé sur le LAN sans processus MAC | Utilisateurs bloqués, helpdesk saturé | Allow réservé aux VLANs contrôlés (§51) |
| 10 | Modification d'option serveur sans synchro manuelle | Config divergente entre les 2 nœuds | Script §63 |
| 11 | Pas de sauvegarde / jamais testée | Reconstruction manuelle après crash | Export quotidien + test annuel (§73) |
| 12 | DHCP snooping oublié sur un nouveau VLAN | Rogue DHCP possible | Checklist de création de VLAN (§94) |
| 13 | Changement d'IP du serveur sans mise à jour des helpers | VLANs sans DHCP | Procédure §91 |
| 14 | Étendue créée sans relais sur VLAN distant | Broadcast bloqué, pas d'IP | Une étendue par VLAN + helper (§64) |
| 15 | Logs d'audit non archivés | Impossible de tracer une IP dans le passé | Archivage 1 an (§99) |

## A6. Pour aller plus loin

**Documentation officielle Microsoft** :

- *DHCP — Vue d'ensemble* (Windows Server 2022/2025) : `https://learn.microsoft.com/fr-fr/windows-server/networking/technologies/dhcp/dhcp-top`
- *Cmdlets DhcpServer (PowerShell)* : `https://learn.microsoft.com/fr-fr/powershell/module/dhcpserver/`
- *Basculement DHCP — Guide pas à pas* : `https://learn.microsoft.com/fr-fr/windows-server/networking/technologies/dhcp/dhcp-failover`
- *DNS dynamique et DHCP* : `https://learn.microsoft.com/fr-fr/windows-server/networking/technologies/dhcp/dhcp-top`

**RFC de référence** :

- RFC 2131 — Dynamic Host Configuration Protocol (le protocole)
- RFC 2132 — DHCP Options and BOOTP Vendor Extensions (les options)
- RFC 3046 — DHCP Relay Agent Information Option (option 82)
- RFC 4388 — DHCP Lease Query (interrogation des baux)
- RFC 6926 — DHCPv6 Relay Agent (IPv6)

**Sujets connexes à étudier** (guides existants dans `~/workspace/user/files/`) :

- `proxmox_guide.md` — virtualisation des serveurs DHCP (HA VM)
- `debian_ubuntu_guide.md` — alternative ISC DHCP / Kea sur Linux
- `onduleurs_ups_guide.md` — protéger électriquement les serveurs DHCP (un DHCP sans courant = plus de renouvellements)
- `zabbix_guide.md` — supervision des sondes DHCP (section 106)

**Outils complémentaires** :

- Wireshark (filtre `bootp`) — analyse DORA
- Nmap (`--script broadcast-dhcp-discover`) — détection rogue
- DhcpExplorer / dhcpdump — tests manuels
- Yersinia (labo uniquement !) — comprendre les attaques starvation pour mieux s'en protéger

---

*Guide rédigé le 2026-09-27 — versions Windows Server 2019/2022/2025. Les exemples utilisent des plages de documentation (RFC 5737) : adapte à ton plan d'adressage réel avant toute mise en production.*
