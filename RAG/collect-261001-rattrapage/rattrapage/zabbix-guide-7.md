---
id: collect-261001-rattrapage/rattrapage/zabbix-guide-7
title: "Guide Zabbix complet — Supervision d'infrastructure en production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "arr"]
source: docs/RAG/collect-261001-rattrapage/zabbix_guide.md
source_anchor: ""
source_lines: [1014, 1223]
sha256: 2dddec4e24e078537e93bcb2a78e13dbdb2c63ef9db9b1946d157994358cc74f
---

# Guide Zabbix complet — Supervision d'infrastructure en production

| Étape | Usage |
|---|---|
| Multiplier / diviser | Octets → Go, ms → s |
| JSONPath / XML XPath | Extraire un champ d'une API |
| Regex | Extraire un motif d'un texte |
| Plage → valeur (Value mapping) | 1→"On battery", 2→"On mains" |
| Différence par seconde | Compteurs SNMP (ifInOctets) → débit |
| Jeter les doublons / les inchangés | Réduire le volume stocké |

Exemple : un onduleur expose la charge en dixièmes de pourcent (`850` = 85,0 %) :

```
Étape 1 : Custom multiplier → 0.1
Résultat : 85 (%)
```

Exemple SNMP : compteur d'octets → Mbit/s :

```
Étape 1 : Change per second
Étape 2 : Custom multiplier → 8
Étape 3 : Custom multiplier → 0.000001   (→ Mbit/s)
```

> 💡 **Toujours** vérifier le prétraitement avec *Test* (bouton dans l'item) avant de sauvegarder.

## 34. Intervalles, périodes de collecte et flex intervals

- **Intervalle** (`1m`, `30s`, `5m`) : fréquence de collecte. Plus c'est court, plus c'est précis — et plus ça coûte (base, réseau, CPU).
- **Flex intervals** : collectez plus souvent à certaines heures (`Interval: 30s, Period: 1-5,08:00-18:00`) et moins la nuit.
- **Scheduling** (`wd{1-5}h8-18`) : syntaxe avancée pour des créneaux complexes (⚠️ 6.0+).

Repères :

| Métrique | Intervalle conseillé |
|---|---|
| Ping / disponibilité | 1 min (5 min sur lien VSAT) |
| CPU / RAM | 1 min |
| Espace disque | 5 min (avec LLD) |
| Température salle | 2–5 min |
| Charge onduleur | 30 s–1 min (critique pour vous) |
| Inventaire (packages) | 1 jour |

> 💡 **Règle** : l'intervalle doit être ≤ 1/3 du délai d'alerte souhaité. Alerte "disque plein dans l'heure" → collecte toutes les 5 min minimum.

## 35. Triggers : syntaxe et logique

Un trigger = **expression** évaluée à chaque nouvelle valeur → états `OK` / `PROBLEM`.

```
{<server>:<clé>.<fonction>(<paramètres>)}<opérateur><seuil>
```

Exemples :

```
# Load > 4 sur 5 min
{serveur:system.cpu.load[all,avg1].avg(5m)}>4

# Disque / > 90 %
{srv:vfs.fs.size[/,pused].last()}>90

# Pas de données depuis 10 min (agent mort ?)
{srv:agent.ping.nodata(10m)}=1

# Service SSH arrêté
{srv:systemd.unit.is_active[ssh].last()}=0
```

**Opérateurs** : `> < >= <= = <>`, `and`, `or`, `not`. **Sévérités** : Not classified, Information, Warning, Average, High, Disaster.

### Assistant d'expression (recommandé)

Dans le frontend, *Configuration → Triggers → Create* : le bouton **Add** ouvre le constructeur (item + fonction + paramètres + seuil). **Utilisez-le** : il évite 99 % des erreurs de syntaxe.

### Nommage

```
# Bon : précis, avec macros
"Charge CPU > {$CPU_LOAD_CRIT} sur {HOST.NAME} depuis 5 min"
# Mauvais : vague
"Alerte CPU"
```

> 💡 Incluez `{HOST.NAME}` et `{ITEM.VALUE}` dans le nom : l'email d'alerte devient lisible sans ouvrir le frontend.

## 36. Fonctions de trigger : le catalogue indispensable

🖨️ **Fiche mémo** — les fonctions à connaître par cœur :

| Fonction | Signification | Exemple d'usage |
|---|---|---|
| `last()` | Dernière valeur | Seuil simple |
| `avg(5m)` | Moyenne sur 5 min | Lisser les pics CPU |
| `min(15m)` / `max(15m)` | Min/max sur période | Pics mémoire |
| `nodata(10m)` | =1 si aucune donnée depuis 10 min | Agent/équipement muet |
| `change()` | Différence avec la valeur précédente | Compteur qui bouge |
| `diff()` | =1 si la valeur a changé (0 sinon) | Détection de reboot via uptime |
| `date()` / `dayofweek()` | Date/jour courant | Alertes ouvrées uniquement |
| `time()` | Heure courante HHMMSS | Plages horaires |
| `count(10m,,"gt")` | Nombre de valeurs > seuil sur 10 min | "3 dépassements en 10 min" |
| `forecast()` | Prédiction linéaire | "Disque plein dans 7 jours" |
| `trendavg(1h,7d)` | Moyenne des trends | Comparaison semaine/semaine |

Exemples commentés :

```
# CPU > 90 % pendant 10 minutes d'affilée (évite les pics)
{srv:system.cpu.util.avg(10m)}>90

# Au moins 5 dépassements de 95 % CPU en 15 min
{srv:system.cpu.util.count(15m,95,"gt")}>5

# Reboot détecté (uptime a diminué)
{srv:system.uptime.change()}<0

# Pas de ping depuis 5 min → hôte injoignable
{sw:icmpping.nodata(5m)}=1

# Prédiction : /var plein dans moins de 7 jours
{srv:vfs.fs.size[/var,pfree].forecast(#10,7d,,avg)}<10
```

> ⚠️ `forecast()` et les fonctions sur trends exigent assez d'historique : ne vous étonnez pas d'un `Unknown` les premiers jours.

## 37. Hystérésis et seuils intelligents

**Problème** : un seuil unique à 90 % fait osciller l'alerte (PROBLEM/OK/PROBLEM…) quand la valeur oscille autour de 90 — c'est du **flapping**, source n°1 du bruit.

**Solution — hystérésis** : seuil de déclenchement ≠ seuil de résolution.

```
# Déclenche à 90, ne se résout qu'en dessous de 80
Expression problème : {srv:vfs.fs.size[/,pused].last()}>90
Expression de récupération : {srv:vfs.fs.size[/,pused].last()}<80
```

Dans le frontend : *Trigger → Recovery expression*. Autres techniques :

- **Durée** : `.avg(10m)>90` au lieu de `.last()>90`.
- **Occurrences** : `.count(15m,90,"gt")>=3`.
- **Macros** : `{$DISK_PUSED_CRIT}` = 90, `{$DISK_PUSED_CRIT_RECOVERY}` = 80.

> 💡 Chaque trigger critique de production devrait avoir une expression de récupération. C'est 30 secondes de travail qui divisent le bruit par deux.

## 38. Sévérités et bonnes pratiques de nommage

| Sévérité | Quand l'utiliser | Exemple |
|---|---|---|
| Disaster | Service vital coupé, action immédiate | Onduleur sur batterie + charge > 95 % |
| High | Dégradation forte | Disque > 95 %, site injoignable |
| Average | À traiter sous 24 h | Disque > 85 %, température > 27 °C |
| Warning | Informatif, tendance | Certificat expire dans 30 j |
| Information | Purement informatif | Reboot détecté |

Règles de nommage :

```
[Équipement] Métrique > seuil (contexte)
"Charge onduleur ups-salle-01 > 95 % depuis 2 min"
```

- Toujours `{HOST.NAME}` (ou `{HOST.HOST}`) dans le nom.
- Toujours l'unité et le seuil.
- **Jamais** de "ALERTE !!!" en majuscules : la sévérité porte déjà l'urgence.

## 39. Dépendances entre triggers (éviter les tempêtes d'alertes)

**Scénario** : le switch `sw-coeur-01` tombe → 40 hôtes derrière deviennent "injoignables" → 40 emails. **Inutile.**

**Solution** : *Trigger → Dependencies* : le trigger "hôte injoignable" **dépend** du trigger "switch injoignable". Zabbix **masque** les dépendants (ils passent quand même en PROBLEM mais ne déclenchent pas d'action).

```
Trigger "Switch sw-coeur-01 injoignable"   (cause racine)
    └── Trigger "srv-fichiers-01 injoignable" dépend du précédent
    └── Trigger "ups-salle-01 injoignable" dépend du précédent
```

Configuration : sur le trigger dépendant, onglet *Dependencies → Add*, sélectionnez le trigger du switch.

> 💡 Dessinez votre **arbre de dépendances** par site : routeur → switch → serveurs/onduleurs. 10 minutes de réflexion = des nuits d'astreinte paisibles.

Alternative/complément : les **actions avec conditions** sur les groupes d'hôtes (section 40).

## 40. Actions : le moteur de notification

*Configuration → Actions → Trigger actions*. Une action = **conditions** (quand) + **opérations** (quoi faire) + **escalades** (quand répéter/monter en puissance).

### Exemple : action "Alertes critiques énergie"

**Conditions** (toutes doivent matcher — logique AND entre types différents) :

```
Trigger severity >= High
AND Host group = Onduleurs
AND Trigger value = PROBLEM
```

**Opérations** :

```
Step 1 (0s)      : Send to user group "Energie" via Email + Telegram
Step 2 (15 min)  : Send to user "Chef de service" via SMS   (escalade)
Step 3 (1h, répéter toutes les 30 min) : répéter l'escalade
```

**Recovery operation** : envoyer "RÉSOLU" aux mêmes destinataires (toujours configurer !).

Paramètres clés d'une opération :

