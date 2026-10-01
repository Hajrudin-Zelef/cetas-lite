---
id: collect-261001-rattrapage/rattrapage/huawei-usg6000-guide-12
title: "Guide ULTRA-COMPLET — Huawei USG6000 (Firewall UTM)"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_usg6000_guide.md
source_anchor: ""
source_lines: [1736, 1891]
sha256: 73cb0b27c19b66c75739145e2271f0e647913214f0f1eea078b2fbf4d9b4d8e2
---

# Guide ULTRA-COMPLET — Huawei USG6000 (Firewall UTM)

- **HRP** (Huawei Redundancy Protocol) : synchronise **sessions, tables NAT, tables de routage, infos VPN** entre les deux USG.
- **VGMP** (VRRP Group Management Protocol) : gère le **basculement** — regroupe les VRRP des interfaces ; si un lien ou un boîtier tombe, tout bascule d'un coup (pas interface par interface).
- Mode **actif/passif** (hot standby) : un seul traite le trafic, l'autre attend **synchronisé**. Bascule en **quelques secondes**, sessions conservées (pas de coupure TCP).

## 98. Prérequis matériels et réseau

📋 Checklist :
- [ ] **2 USG identiques** : même modèle, même version logicielle, mêmes cartes d'extension.
- [ ] **Licences équivalentes** des deux côtés (sinon, après bascule, l'UTM ne protège plus).
- [ ] **Lien heartbeat dédié** : 1 interface (ou 2 pour la redondance) reliée en direct entre les deux USG — **jamais** via un switch partagé avec le trafic.
- [ ] Les interfaces de même nom des deux USG branchées sur les **mêmes réseaux** (même VLAN/switch).
- [ ] Câblage symétrique : GE1/0/1 du A et GE1/0/1 du B sur le même switch WAN, etc.

## 99. Configuration HA complète (les deux nœuds)

```huawei
# ===== Sur USG-A (futur actif) =====
system-view
[USG-A] hrp enable
[USG-A] hrp interface GigabitEthernet 1/0/7 remote 10.0.7.2   # heartbeat vers le pair
[USG-A] interface GigabitEthernet 1/0/7
[USG-A-GigabitEthernet1/0/7] ip address 10.0.7.1 30
[USG-A-GigabitEthernet1/0/7] quit
# VRRP sur chaque interface de production (exemple WAN) :
[USG-A] interface GigabitEthernet 1/0/1
[USG-A-GigabitEthernet1/0/1] vrrp vrid 1 virtual-ip 203.0.113.1
[USG-A-GigabitEthernet1/0/1] vrrp vrid 1 priority 120
[USG-A-GigabitEthernet1/0/1] quit
# ... répéter sur LAN, DMZ ...
# Groupe VGMP :
[USG-A] vrrp group 1
[USG-A-vrrp-group-1] add interface GigabitEthernet 1/0/1 vrid 1
[USG-A-vrrp-group-1] add interface GigabitEthernet 1/0/3 vrid 2
[USG-A-vrrp-group-1] quit
[USG-A] hrp track vrrp group 1
save
```

```huawei
# ===== Sur USG-B (futur passif) : MIROIR avec priorités plus basses =====
system-view
[USG-B] hrp enable
[USG-B] hrp interface GigabitEthernet 1/0/7 remote 10.0.7.1
[USG-B] interface GigabitEthernet 1/0/7
[USG-B-GigabitEthernet1/0/7] ip address 10.0.7.2 30
[USG-B-GigabitEthernet1/0/7] quit
[USG-B] interface GigabitEthernet 1/0/1
[USG-B-GigabitEthernet1/0/1] vrrp vrid 1 virtual-ip 203.0.113.1
[USG-B-GigabitEthernet1/0/1] vrrp vrid 1 priority 100
[USG-B-GigabitEthernet1/0/1] quit
# ... même VGMP ...
[USG-B] vrrp group 1
[USG-B-vrrp-group-1] add interface GigabitEthernet 1/0/1 vrid 1
[USG-B-vrrp-group-1] add interface GigabitEthernet 1/0/3 vrid 2
[USG-B-vrrp-group-1] quit
[USG-B] hrp track vrrp group 1
save
```

⚠️ La syntaxe HRP/VGMP **a évolué selon les versions** — valide chaque commande avec `?` sur ta version. Le principe (heartbeat + VRRP par interface + groupe VGMP) est constant.

## 100. Vérifier que la HA est saine

```huawei
[USG-A] display hrp state verbose     # rôle : active / standby, état du peer
[USG-A] display hrp statistics       # paquets heartbeat, erreurs
[USG-A] display vrrp                 # état des VRID
[USG-A] display hrp sync-status      # état de la synchronisation de config/sessions
```

État sain : A = **active**, B = **standby**, heartbeat OK, sync OK, compteurs d'erreurs à 0.

## 101. Tester le basculement (en heures creuses !)

📋 Procédure de test :
1. Prévenir (même en heures creuses, quelqu'un doit savoir).
2. Ouvrir un **ping continu** et une **session TCP** (ex. SSH) à travers le firewall.
3. Sur l'actif : `hrp switch` (bascule manuelle) ou débrancher le lien WAN.
4. Observer : le ping perd **1 à 3 paquets**, la session SSH **survit** (grâce à la sync HRP).
5. `display hrp state verbose` sur les deux nœuds → les rôles ont permuté.
6. Rebasculer si besoin (`hrp switch` sur le nouveau passif, ou préemption selon config).
7. 📋 **Tester 2 fois par an.** Une HA jamais testée = une HA qui ne marche pas le jour J.

## 102. Session asymétrique en HA : le problème et le remède

**Symptôme** : en temps normal ça marche, mais après un basculement (ou avec un routage asymétrique), des sessions tombent.

**Cause** : le paquet retour arrive sur le nœud qui n'a pas la session (ou la session n'a pas été synchronisée à temps).

**Remèdes** :
- Vérifier que **HRP synchronise bien** (`display hrp sync-status`) — configurer la **synchro automatique** (`hrp auto-sync` / `hrp sync` selon version).
- Vérifier le **câblage symétrique** (section 98) : si le retour passe par un autre chemin physique, c'est mort.
- Éviter les **routes asymétriques** en amont (HSRP/VRRP des routeurs autour).
- 🔧 `display firewall session table` comparé sur les deux nœuds : les tables doivent être **quasi identiques**.

## 103. HA et VPN : ce qui est synchronisé (et ce qui ne l'est pas)

- **Synchronisé** : sessions IPSec/SSL VPN, SA, compteurs — le tunnel **survit** au basculement (renégociation évitée dans la plupart des cas).
- ⚠️ **Non synchronisé** : les **certificats** (à installer sur les deux nœuds), les **licences** (une par nœud), certains états applicatifs.
- 📋 Après tout changement de certificat ou de licence : **le refaire sur le second nœud** et vérifier.

## 104. HA et upgrade firmware : sans coupure (ou presque)

1. Upgrader le **passif** d'abord, le rebooter, vérifier qu'il revient en standby sain.
2. Basculer (`hrp switch`) → l'ancien actif devient passif.
3. Upgrader l'ex-actif, le rebooter.
4. Vérifier la paire, rebasculer si besoin.
5. ⚠️ Les deux nœuds doivent finir sur la **même version** — une paire en versions différentes, c'est un comportement **non supporté** à terme.

## 105. Split-brain : quand les deux se croient actifs

**Cause** : lien heartbeat coupé → chaque nœud pense que l'autre est mort → les deux deviennent actifs → **conflit d'IP virtuelles**, chaos réseau.

**Prévention** :
- **Double lien heartbeat** (deux interfaces, chemins physiques différents).
- Le heartbeat ne doit **jamais** transiter par un équipement unique partagé.
- 🔧 Détection : `display hrp state verbose` affiche **active** sur les deux → couper le lien de production du « faux actif » en urgence, réparer le heartbeat, resynchroniser.

## 106. Web : supervision de la HA

Chemin (New Web UI) : **Monitor > HA** ou System > HA : état des nœuds, historique des bascules, état du heartbeat.
📋 En routine : vérifier **1 fois par semaine** que le standby est bien standby et synchronisé. Ça prend 30 secondes.

---
---

# BLOC J — LOGS, RAPPORTS, SUPERVISION

## 107. Le log center : où vont les logs

L'USG génère : logs de **politique** (permit/deny), logs **UTM** (AV/IPS/URL), logs **système**, logs **VPN**, logs d'**attaque**.

Destinations possibles :
- **Buffer mémoire** (limité, perdu au reboot).
- **Disque dur local** (si installé — recommandé pour garder l'historique).
- **Serveur syslog** externe (recommandé en prod).
- **eLog** (serveur de logs Huawei, voir section 108).

```huawei
system-view
[USG] info-center enable
[USG] info-center logbuffer size 1024
[USG] info-center syslog level informational
save
```

## 108. eLog : le serveur de logs Huawei

**eLog** = le collecteur/analyseur de logs de Huawei (gratuit dans sa version de base, à vérifier selon version). Il reçoit les logs des USG, les indexe, propose des **tableaux de bord** et des **recherches**.

Alternative : n'importe quel **syslog** (rsyslog, Graylog, Splunk...) reçoit les logs USG en standard. eLog apporte les **dashboards pré-câblés** pour firewall Huawei.

📋 Choix terrain : si tu as 1–2 USG, un syslog classique + Grafana suffit. Si tu as un parc Huawei (firewalls + NCE), eLog/Secomea... (vérifier le nom exact selon génération) centralise mieux.

## 109. Envoyer les logs vers un syslog externe

