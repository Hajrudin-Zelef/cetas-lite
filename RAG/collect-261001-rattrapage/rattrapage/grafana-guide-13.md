---
id: collect-261001-rattrapage/rattrapage/grafana-guide-13
title: "Guide Grafana — Dashboards, visualisation et alerting"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["datacenter", "memory"]
source: docs/RAG/collect-261001-rattrapage/grafana_guide.md
source_anchor: ""
source_lines: [2649, 2854]
sha256: 72fa767509664d3bbc10adda23939c4c25d53c605f1bb41afbaaca9ff708143e
---

# 3. Redémarrage + vérification
sudo systemctl restart grafana-server
sleep 10
curl -s http://localhost:3000/api/health | python3 -m json.tool
sudo tail -n 50 /var/log/grafana/grafana.log | grep -iE "migrat|error|fail"

# 4. Contrôles fonctionnels
# - ouvrir 2-3 dashboards clés, vérifier les panels
# - Alerting → vérifier que les règles s'évaluent (état Normal/Pending)
# - envoyer un Test sur chaque contact point critique
```

Bonnes pratiques :

- **Jamais de mise à jour un vendredi après-midi** ni avant une astreinte.
- Testez d'abord sur une **instance de labo** (ou un snapshot VM).
- Ne sautez pas plus d'une version majeure d'un coup (10 → 11 OK, 9 → 11 :
  passez par 10).
- Les plugins tiers peuvent casser : vérifiez leur compatibilité
  (`grafana-cli plugins ls`, page du plugin).
- Après mise à jour majeure, rouvrez les dashboards provisionnés : le
  schéma JSON migre automatiquement, mais vérifiez visuellement.

Rollback : réinstallez la version précédente (`apt install grafana=11.4.0`)
**et** restaurez la base sauvegardée avant (le schéma a pu migrer).

---

## 67. Usage quotidien — routine de l'exploitant

**Le matin (5 min) — tour de contrôle :**

1. Ouvrir le dashboard « Vue d'ensemble » : tout est vert ?
2. **Alerting → Alert list** : alertes firing/pending non acquittées ?
3. **Silences actifs** : un silence a-t-il expiré cette nuit sans résolution ?
4. Logs Grafana : erreurs inhabituelles (`grep -i error` sur les dernières
   24 h) ?

**En intervention :**

1. **Explore** pour investiguer (requête ad hoc, pas de dashboard jetable).
2. Annoter le dashboard (`Ctrl+clic`) : « début d'investigation 09:12 ».
3. Si maintenance planifiée : créer le **silence** avec ticket de référence.
4. Après résolution : vérifier que l'alerte repasse en Normal, clore le
   silence, noter la cause dans le ticket.

**Hebdomadaire (30 min) :**

- Revoir les alertes de la semaine : faux positifs à ajuster ?
- Vérifier les sauvegardes (log du cron, taille des dumps).
- Nettoyer les dashboards « brouillon » abandonnés.

**Mensuel :**

- Test de la chaîne d'alerte de bout en bout (alerte de test → réception).
- Revue des comptes/tokens (départs, tokens expirés).
- Vérifier l'espace disque de `/var/lib/grafana` et `/var/log/grafana`.

---

## 68. Bonnes pratiques de dashboarding (lisibilité)

1. **Un dashboard = une question.** « Mes serveurs sont-ils sains ? », pas
   « tout l'infra en 48 panels ».
2. **Pas de spaghetti** : 6-8 séries max par graphe. Au-delà, agrégez
   (`sum by`), filtrez, ou divisez.
3. **Hiérarchie visuelle** : en haut les indicateurs (Stat), au milieu les
   tendances (Time series), en bas le détail (Table, Logs).
4. **Seuils et couleurs sobres** : vert/orange/rouge, pas d'arc-en-ciel.
   La couleur doit signifier quelque chose (section 34).
5. **Intervalles adaptés** : `$__rate_interval` dans les `rate()`, Min step
   cohérent avec la plage (pas de point par seconde sur 30 jours).
6. **Titres qui disent la question** : « % CPU par serveur (5 min) », pas
   « Graph 1 ». Description du panel : unité, source, seuil d'alerte.
7. **Variables en haut, dans l'ordre logique** : datacenter → serveur →
   interface. Valeurs par défaut sensées (All ou le périmètre de l'équipe).
8. **Zéro panel vide** : un « No data » permanent = une requête cassée à
   réparer ou un panel à supprimer.
9. **Documentez dans le dashboard** (panel Text) : question, mainteneur,
   runbook, date de dernière revue.
10. **Testez sur écran mural** : lisible à 3 mètres ? Si non, augmentez les
    tailles, réduisez le nombre de panels (voir section 72).

**Anti-patterns à bannir :**

| Anti-pattern | Pourquoi c'est mal | Alternative |
|---|---|---|
| 20 dashboards quasi identiques | Maintenance impossible | Variables + 1 dashboard |
| Jauges sans max | Chiffre sans référence | Min/Max renseignés |
| Courbes qui relient les trous | Masque les pannes de collecte | Connect null → Never |
| Seuils décoratifs jamais alertés | Bruit visuel | Seuil = alerte ou rien |
| `rate()` sur une gauge | Chiffres absurdes | `avg_over_time` |

---

## 69. Exemple : dashboard supervision serveurs (node_exporter)

**Question :** « mes serveurs Linux sont-ils sains ? »

Structure (dossier `Supervision`, variables `$dc`, `$serveur`) :

**Row « Vue d'ensemble » (Stat) :**

```promql
# Serveurs UP
count(up{job="node"} == 1)
# Serveurs DOWN
count(up{job="node"} == 0)
# CPU moyen du parc
100 - (avg(rate(node_cpu_seconds_total{mode="idle"}[5m])) * 100)
# Alertes firing
sum(ALERTS{alertstate="firing"})
```

**Row « Serveur : $serveur » (répétée par serveur si besoin) :**

```promql
# CPU % par mode (Time series, stacked)
sum by (mode) (rate(node_cpu_seconds_total{instance=~"$serveur"}[$__rate_interval]))

# Mémoire (Time series)
node_memory_MemTotal_bytes{instance=~"$serveur"} - node_memory_MemAvailable_bytes{instance=~"$serveur"}
# + ligne de référence : node_memory_MemTotal_bytes

# Disque % par mountpoint (Time series + seuils 80/90)
100 * (1 - (node_filesystem_avail_bytes{fstype!~"tmpfs|overlay",instance=~"$serveur"}
            / node_filesystem_size_bytes{fstype!~"tmpfs|overlay",instance=~"$serveur"}))

# Réseau (Time series, 2 axes ou 2 panels)
rate(node_network_receive_bytes_total{device!~"lo|veth.*",instance=~"$serveur"}[$__rate_interval]) * 8
rate(node_network_transmit_bytes_total{device!~"lo|veth.*",instance=~"$serveur"}[$__rate_interval]) * 8

# Load (Time series)
node_load1{instance=~"$serveur"}
node_load5{instance=~"$serveur"}
node_load15{instance=~"$serveur"}
```

**Row « Détail » (Table) :** top filesystems, processus top CPU via
`topk(5, ...)`, dernières alertes (Alert list).

Annotations : déploiements (section 40). Liens : drill-down vers les logs
Loki du serveur (section 41).

---

## 70. Exemple : dashboard supervision réseau

**Question :** « mon réseau transporte-t-il correctement ? »

Sources : SNMP via snmp_exporter → Prometheus, ou flows, ou Zabbix.

**Row « Liens WAN/LAN » :**

```promql
# Débit par interface (bits/s), top 8
topk(8, rate(ifHCInOctets[$__rate_interval]) * 8)
# Légende : {{ifDescr}} ({{instance}})

# Utilisation % vs capacité (nécessite la vitesse en label ou en metric)
100 * rate(ifHCInOctets[$__rate_interval]) * 8 / ifHighSpeed * 1e6

# Erreurs/discards (signe d'un lien dégradé)
rate(ifInErrors[$__rate_interval]) + rate(ifOutErrors[$__rate_interval])
rate(ifInDiscards[$__rate_interval]) + rate(ifOutDiscards[$__rate_interval])
```

**Row « Disponibilité » (State timeline) :**

```promql
up{job="snmp"}
# ou ping : probe_success{job="blackbox-icmp"}
```

**Row « Inventaire » (Table) :** équipement, interface, description,
débit actuel, erreurs — via transformations (Merge + Organize fields).

Bonnes pratiques réseau :

- **Nommez les interfaces** (`ifDescr`/`ifAlias` en légende, pas `ifIndex`).
- Seuils : utilisation > 80 % sustained = capacité à planifier ; erreurs >
  0 = investiguer (un lien sain n'a quasiment pas d'erreurs).
- **Dépendances** : un switch d'agrégation down explique 20 alertes
  serveurs — regroupez les notifications par site (`group_by: [site]`).

---

## 71. Exemple : dashboard supervision onduleur (lien métier)

**Question :** « mes onduleurs protègent-ils correctement la charge ? »
*Pour un chef de service systèmes & énergies, c'est un dashboard vital.*

Sources possibles : carte réseau de l'onduleur (SNMP → snmp_exporter →
Prometheus), NUT (`nut_exporter`), ou supervision constructeur (Eaton IPP,
Schneider/APC PowerChute, Vertiv) qui expose SNMP ou Modbus.

**Row « État instantané » (Stat + Gauge) :**

```promql
# Charge de l'onduleur en % (Gauge, seuils 70/90)
ups_load_percent{ups="$onduleur"}

# Tension d'entrée / sortie (Stat)
ups_input_voltage{ups="$onduleur"}
ups_output_voltage{ups="$onduleur"}

