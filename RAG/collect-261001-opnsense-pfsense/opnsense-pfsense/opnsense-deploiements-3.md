---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/opnsense-deploiements-3
title: "OPNsense — 4 déploiements simulés (niveau 1 → hardcore)"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["intel", "memory"]
source: docs/RAG/collect-261001-opnsense-pfsense/opnsense_deploiements.md
source_anchor: ""
source_lines: [394, 534]
sha256: 049599df442c25057ee57648cd7b32510809f4c966bfa796aebfb095d1e1ad9c
---

# OPNsense — 4 déploiements simulés (niveau 1 → hardcore)

### A. Routage BGP (plugin FRR)
1. Installer **os-frr**, activer : **Routing > FRR > General**.
2. **Routing > FRR > BGP** : AS 65001, router-id, neighbors vers les
   deux transits (remote-as du FAI, password MD5, prefix-lists
   d'import/export strictes).
3. **Prefix-lists** : n'annoncer **que** vos préfixes (jamais de
   `0.0.0.0/0` en sortie vers un transit, jamais de full-table
   acceptée sans filtre — risque de devenir transit par accident).
4. **Route-maps** : prepend/communities selon politique.
5. Vérifier : `vtysh -c "show bgp summary"` (FRR), préfixes reçus/
   annoncés conformes.
6. **BFD** si supporté par les transits (détection de panne < 1 s).

### B. IPv6 natif dual-stack
7. **Interfaces > [WAN]** : IPv6 = DHCPv6 ou statique selon transit.
8. **Services > DHCPv6 + Router Advertisements** : RA `Managed` sur
   les VLANs clients (ou SLAAC + stateless DHCPv6 selon client).
9. **Firewall > Rules** : **dupliquer la politique en IPv6**
   (les règles IPv4 ne couvrent PAS IPv6 — erreur classique :
   IPv6 grand ouvert par défaut !).
10. **Firewall > NAT > NPTv6** si besoin (éviter le NAT66, préférer
    le routage pur).
11. Tester : `test-ipv6.com` depuis un client → 10/10.

### C. Filtrage applicatif et threat intel
12. **Zenarmor** (plugin) : déployer sur les interfaces clients —
    politiques par catégorie (réseaux sociaux, streaming, malware),
    rapports par IP/client, **TLS inspection** avec CA déployée
    (clients managés uniquement, avec consentement/charte).
13. **CrowdSec** (plugin) : bannissement collaboratif (les IP qui
    attaquent ailleurs sont bloquées ici avant d'essayer).
14. **Suricata** : IPS sur TRANSIT + règles ET Pro si budget.
15. **Firewall > Aliases > URL Tables** : blocklists (Spamhaus DROP,
    Emerging Threats) rafraîchies toutes les heures.

### D. QoS stricte (Shaper)
16. **Firewall > Shaper** : pipes par client (débit contractuel),
    queues prioritaires (VoIP EF > tout), règle « par VLAN ».
17. Exemple : pipe `CLIENT-A-DOWN` 500 Mbit, `CLIENT-A-UP` 500 Mbit ;
    queue VoIP `priority 7`, queue default `weight 1`.
18. Vérifier sous charge : `ping` stable pendant un speedtest
    (bufferbloat maîtrisé).

### E. Automatisation (API)
19. **System > Access > Users** : créer un compte `api-bot` avec clé
    API (privilèges restreints : aliases, diagnostics).
20. Scripts : ajouter/retirer des IP d'alias en masse, snapshots de
    config avant chaque changement, checks de supervision.
21. Exemple (conceptuel) :
    `POST /api/firewall/alias/addItem` — voir la doc API intégrée
    (`https://<fw>/api/docs` selon version / wiki OPNsense).
22. **Git** : versionner les exports `config.xml` (cron + push).

### F. Durcissement extrême
23. Tout le §15 du guide web UI, **plus** :
    - Web UI + SSH : écoute sur **MGMT/VLAN 99 uniquement**, via VPN
      admin (WireGuard dédié) pour l'accès distant.
    - `root` : login console uniquement ; admins = comptes nominatifs
      + 2FA + sudo ciblé.
    - **System > Firmware > Audit** : vérification d'intégrité
      hebdomadaire (cron).
    - Désactiver **tous** les services non utilisés (proxy, captive
      portal, etc.).
    - Logs : syslog TLS vers SIEM **hors-site**, rétention 1 an.
    - **Backups chiffrés** : local + hors-site, test de restauration
      trimestriel (pas de backup non testé = pas de backup).
    - Plan de **disaster recovery** écrit : RTO/RPO, procédure de
      rebuild from scratch, contacts FAI/transits.

### G. Supervision
24. **SNMPv3** → Zabbix : CPU, mémoire, états pf, trafic par interface,
    statut CARP, tunnels.
25. **NetFlow/Insight** : rétention longue (dimensionnement, facturation
    au 95e percentile si revente de bande passante).
26. Alertes : gateways down, CARP MASTER change, Suricata blocks en
    masse, disque > 80 %, certificats < 30 jours (ACME auto normalement).

## Équivalents CLI

```shell
# BGP (FRR)
vtysh -c "show bgp summary"
vtysh -c "show bgp ipv6 unicast summary"
# IPv6
ifconfig | grep inet6
netstat -rn -f inet6 | head
# États pf sous charge
pfctl -s info | grep -i "state-table\|insert"
pfctl -s memory
# Queues shaper
pfctl -s queue -v
# API (exemple lecture d'alias)
# curl -k -u "key:secret" https://fw-mgmt/api/firewall/alias/get
```

## Vérifications
- [ ] `show bgp summary` : 2 sessions Established, préfixes annoncés
      = vos blocs uniquement.
- [ ] Coupure d'un transit : convergence BGP < 2 min, pas de blackhole.
- [ ] IPv6 : 10/10 sur test-ipv6.com, règles IPv6 = miroir des IPv4.
- [ ] Speedtest client : débit = contrat, latence stable (QoS).
- [ ] Attaque simulée (scan nmap) : CrowdSec/Suricata bannissent,
      alerte reçue.
- [ ] Restore : rebuild complet d'un nœud depuis le backup en < 1 h
      (chronométré en exercice).

## Pièges hardcore
- Devenir **transit accidentel** (réannoncer les routes d'un FAI à
  l'autre) → prefix-lists strictes, toujours.
- IPv6 oublié dans les règles → backdoor grande ouverte.
- TLS inspection sans base légale/charte → problème juridique.
  En entreprise : informer et faire signer.
- API sans restriction de privilèges = RCE déguisé. Compte dédié,
  réseau MGMT uniquement.
- pfsync sur lien saturé → états désynchronisés → coupures au
  failover. Lien sync dédié, jamais sur le trunk de production.

---

# Tableau récapitulatif

| Aspect | Niv. 1 | Niv. 2 | Niv. 3 | Hardcore |
|---|---|---|---|---|
| WAN | 1 (DHCP) | 1 (fixe) | 2 + failover | 2 transits BGP |
| HA | non | non | CARP 2 nœuds | CARP + pfsync |
| VLANs | 1 LAN | 3 | 4+ | N clients |
| VPN | non | WireGuard RA | + IPsec sites | + VPN admin dédié |
| DNS | Unbound + DoT | + overrides internes | idem | idem |
| Filtrage | NAT + base | inter-VLAN strict | + Suricata IPS | + Zenarmor + CrowdSec |
| IPv6 | non | optionnel | recommandé | natif dual-stack |
| Routage | statique | statique | statique | BGP (FRR) |
| QoS | non | non | basique | shaper contractuel |
| Supervision | dashboard | dashboard | SNMP + NetFlow | Zabbix + SIEM |
| Backups | manuel | manuel + historique | auto chiffré | auto + test restore |
| 2FA admin | non | oui | oui | oui + VPN admin |

---

*Fin des scénarios. Chaque niveau est déployable tel quel en suivant
les étapes dans l'ordre — et chaque niveau supérieur suppose les
réflexes du précédent (backup avant, tester après, monitorer toujours).*
