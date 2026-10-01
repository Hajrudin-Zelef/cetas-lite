---
id: collect-261001-rattrapage/rattrapage/zabbix-guide-14
title: "Guide Zabbix complet — Supervision d'infrastructure en production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "agents", "arr"]
source: docs/RAG/collect-261001-rattrapage/zabbix_guide.md
source_anchor: ""
source_lines: [2330, 2496]
sha256: a2acfb95c02a9ff9c2d9481f928c8b6225c4baacaa6e15ecec8556caf4072ae9
---

# Guide Zabbix complet — Supervision d'infrastructure en production

```bash
snmpwalk -v2c -c communaute_lecture_ici 192.168.10.90 1.3.6.1.2.1.33.1.1.1.1.2.0
# Attendu : STRING: "Easy UPS 3S 40 kVA" (ou modèle réel)
```

**Étape 2 — Hôte** : groupe `Onduleurs`, interface SNMP, macros `{$SNMP_COMMUNITY}`, template `Onduleur SNMP générique (RFC1628)` (section 52).

**Étape 3 — Seuils adaptés au 40 kVA** (macros au niveau hôte) :

```
{$UPS_LOAD_WARN} = 70      # 28 kVA : anticiper l'extension
{$UPS_LOAD_CRIT} = 85      # 34 kVA : proche de la limite
{$UPS_BATT_WARN} = 60
{$UPS_BATT_CRIT} = 30
```

**Étape 4 — Action dédiée** (section 40) : Disaster/High du groupe `Onduleurs` → Email + Telegram immédiats, SMS à 5 min, escalade chef de service à 20 min.

**Étape 5 — Dashboard** : widget "Item value" géant sur `upsEstimatedMinutesRemaining` + courbe charge 24 h + température batterie.

**Étape 6 — Test réel** : passage sur batterie (test programmé avec maintenance !) → vérifier que toute la chaîne se déclenche en < 2 min.

## 77. Cas pratique n°3 : alerte coupure électrique avec escalade SMS

**Scénario** : coupure secteur à 14:03, onduleur sur batterie.

```
14:03:30  Item upsOutputSource → 5 (poll 30 s)
14:03:35  Trigger "COUPURE SECTEUR : ups-salle-01 sur batterie" → PROBLEM (Disaster)
14:03:35  Action : Email + Telegram → groupe Energie  ✓
14:08:35  Pas d'acquittement → SMS → groupe Energie : "COUPURE ups-salle-01, batterie 82 %, autonomie ~35 min"
14:23:35  Toujours pas d'acquittement → SMS → Chef de service
14:41:00  Secteur rétabli → upsOutputSource → 3
14:41:05  Recovery → Email + Telegram "RÉSOLU : retour sur secteur à 14:41"
```

**Points de vigilance** :
- Le SMS contient **l'autonomie restante** (`{ITEM.VALUE}` de l'item autonomie) : l'astreinte sait si elle a 10 ou 40 min.
- Si la coupure touche aussi le server Zabbix : c'est le **watchdog externe** (section 69) qui alerte via la passerelle GSM autonome.
- **Ne jamais** mettre l'envoi SMS sur le même onduleur que la salle supervisée sans batterie de secours pour la clé 4G.

## 78. Cas pratique n°4 : supervision d'un lien VSAT/fibre instable

**Contexte** : site distant relié en VSAT (latence 600 ms, micro-coupures).

1. **Proxy actif** sur site (`ProxyOfflineBuffer=48h`, `HeartbeatFrequency=120`).
2. **Intervalles adaptés** : ping 5 min, SNMP 5 min, pas de LLD agressive.
3. **Triggers tolérants** :
```
# Site down : 3 échecs de ping sur 15 min (pas 1 seul)
{proxy-site:icmpping.count(15m,0)}>2
```
4. **Dépendance** : tout le site dépend du trigger "lien VSAT down".
5. **Dashboard** : disponibilité 30 jours du lien (pour renégocier le SLA opérateur avec des preuves).

> 💡 Les données de disponibilité Zabbix sont des **preuves contractuelles** : exportez le rapport mensuel avant chaque réunion avec l'opérateur.

## 79. Cas pratique n°5 : capacity planning disque avec prédiction

**Objectif** : être alerté **avant** la panne, pas pendant.

```
Trigger "Disque {#FSNAME} plein dans < 14 jours" (Warning) :
  {srv:vfs.fs.size[{#FSNAME},pfree].forecast(#100,14d,,avg)}<10
  and {srv:vfs.fs.size[{#FSNAME},pfree].avg(1h)}<20
```

- `forecast(#100,14d,,avg)` : régression sur les 100 dernières valeurs, projection à 14 jours.
- La 2e condition évite les faux positifs sur disques déjà petits mais stables.
- Dashboard "Capacity" : Top hosts par `pfree` croissant + courbes 90 jours (trends).

**Rituel mensuel** : relecture du dashboard capacity → plan d'extension avant saturation.

## 80. Cas pratique n°6 : supervision d'une salle serveur (température, hygrométrie)

**Équipement** : sonde SNMP (ex. Sonde USB/Modbus via passerelle, ou boîtier type AKCP/APC NetBotz).

1. Hôte `sonde-salle-01`, template custom avec items :
```
salle.temp   (OID sonde, °C)   intervalle 2m
salle.hygro  (OID sonde, %HR)  intervalle 5m
```
2. Triggers :
```
# Warning 27 °C, High 32 °C, Disaster 35 °C (+ hysteresis recovery -2 °C)
{salle:salle.temp.avg(10m)}>32
# Hygrométrie > 70 % : risque condensation/corrosion
{salle:salle.hygro.avg(30m)}>70
# Sonde muette
{salle:salle.temp.nodata(10m)}=1
```
3. Corrélation : superposez température salle + charge onduleur + état climatisation sur un graphe.

> 💡 Pour vous : la température salle est un **indicateur avancé** de problème de climatisation, elle-même souvent liée à l'énergie. Mettez-la sur le dashboard "Énergie".

## 81. Cas pratique n°7 : déploiement multi-sites avec proxies

**Architecture** : 1 server central + 3 proxies (sites Nord, Sud, Est).

```
Site Nord (fibre)     Site Sud (4G)         Site Est (VSAT)
proxy-nord (actif)    proxy-sud (actif)     proxy-est (actif)
Buffer 6h             Buffer 24h            Buffer 48h
```

1. Créez les 3 proxies dans le frontend (*Administration → Proxies*), mode actif, PSK par proxy (section 20).
2. Chaque hôte de site : *Monitored by* = son proxy.
3. **Nommage** : préfixe par site (`sud-srv-01`, `est-ups-01`) → filtres et permissions par groupe `Sites/Sud`.
4. **Supervision des proxies eux-mêmes** : hôtes `proxy-sud` avec template `Zabbix proxy health` (officiel) → alerte si `zabbix[proxy,lastaccess]` > 10 min.
5. **Mise à jour** : planifiez par site (jamais les 3 en même temps), ordre server → proxies → agents.

## 82. Cas pratique n°8 : supervision d'un cluster Proxmox VE

1. Sur chaque nœud : Agent 2 + template `Linux by Zabbix agent`.
2. Items spécifiques (UserParameter ou plugin) :
```ini
# VMs en cours (nécessite droits lecture sur l'API ou qm)
UserParameter=pve.vm.running, qm list 2>/dev/null | grep -c running
UserParameter=pve.node.quorum, pvecm status 2>/dev/null | grep -q "Quorate: Yes" && echo 1 || echo 0
```
3. Triggers :
```
# Perte du quorum → Disaster (le cluster ne migre plus)
{pve:pve.node.quorum.last()}=0
# Stockage Ceph/ZFS > 85 %
```
4. Alternative sans agent : API Proxmox via **HTTP agent** + JSONPath (token API dédié).

> 💡 Le quorum Proxmox et l'état Zabbix HA (section 58) : deux clusters à superviser avec la même rigueur — ce sont vos fondations.

## 83. Cas pratique n°9 : blackout test — valider toute la chaîne d'alerte

**Objectif** : prouver que l'alerte arrive vraiment quand tout s'écroule. À faire **2 fois par an**, en heures ouvrées, avec l'équipe prévenue.

🖨️ **Protocole** :

- [ ] Prévenir l'équipe (ceci est un TEST, pas une vraie panne)
- [ ] Couper le secteur de la salle de test (ou simuler : arrêter l'agent d'un hôte pilote)
- [ ] Chronométrer : T+0 coupure → T+? email → T+? Telegram → T+? SMS
- [ ] Vérifier l'escalade : ne pas acquitter, attendre le SMS chef de service
- [ ] Vérifier le watchdog externe (section 69)
- [ ] Rétablir, vérifier le message "RÉSOLU"
- [ ] Compte-rendu : délais mesurés, points de défaillance, actions correctives

**Objectifs** : email < 2 min, SMS < 10 min, escalade < 30 min. Si ce n'est pas tenu : revoir intervalles et actions **avant** la vraie panne.

## 84. Cas pratique n°10 : rapport mensuel de disponibilité pour la direction

*Reports → Availability* ou dashboard Grafana (section 67) :

```
Disponibilité Septembre 2026
─────────────────────────────────────────
Service / équipement        Dispo    Indisponibilité
Onduleur salle (énergie)    99,98 %  9 min (coupure 12/09)
Lien VSAT site Est          97,20 %  20 h (opérateur)
Serveur fichiers            100 %    —
Infra globale (moyenne)     99,40 %
─────────────────────────────────────────
Incidents majeurs : 2 (détail en annexe)
Alertes traitées : 34 (dont 3 Disaster)
```

> 💡 Un rapport mensuel régulier, même d'une page, **justifie le budget supervision** et objective les discussions avec les opérateurs/fournisseurs. C'est aussi votre mémoire en cas d'audit.

## 85. Pense-bête de poche (cheat sheet imprimable)

