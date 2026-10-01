---
id: collect-261001-rattrapage/rattrapage/zabbix-guide-15
title: "Guide Zabbix complet — Supervision d'infrastructure en production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "agents", "arr"]
source: docs/RAG/collect-261001-rattrapage/zabbix_guide.md
source_anchor: ""
source_lines: [2497, 2612]
sha256: d66ab5d1f1132f347989711acdd9c8d619b8a0b55dea0712739227e577eee713
---

# Guide Zabbix complet — Supervision d'infrastructure en production

🖨️ **Zabbix — Mémo d'astreinte** (une page, à plastifier) :

```
── SERVICES ──────────────────────────────
systemctl status zabbix-server zabbix-agent2 zabbix-proxy
systemctl restart zabbix-server
tail -f /var/log/zabbix/zabbix_server.log

── TESTS RAPIDES ─────────────────────────
zabbix_get -s <IP> -k agent.ping
zabbix_get -s <IP> -k system.cpu.load[all,avg1]
snmpwalk -v2c -c <comm> <IP> 1.3.6.1.2.1.33.1.2.3.1.1.2.1  (batterie %)
snmpwalk -v2c -c <comm> <IP> 1.3.6.1.2.1.33.1.1.4.1.1.2.1 (source: 3=secteur,5=batterie)

── OÙ CHERCHER ───────────────────────────
Pas de données  → Monitoring → Latest data
Trigger inconnu → attendre / vérifier l'item
Pas d'alerte    → Reports → Action log, puis Audit
File d'attente  → Administration → Queue
Santé Zabbix   → Dashboard "Zabbix server health"

── NUMÉROS UTILES ────────────────────────
Passerelle SMS : gammu sendsms TEXT +2250700000000 "message"
Test SMS mensuel : le 1er lundi du mois
Escalade : Exploitation (15 min) → Chef de service (45 min)

── EN CAS DE TEMPÊTE D'ALERTES ───────────
1. Configuration → Maintenance → créer (site concerné)
2. Identifier la cause racine (dépendances, section 39)
3. Acquitter avec commentaire
4. Post-mortem : ajuster seuils/dépendances

── SAUVEGARDE ────────────────────────────
/usr/local/bin/backup_zabbix.sh  (cron 02:00)
Restauration : section 71 — TESTÉE LE : ____/____/________
```

## 86. Glossaire

| Terme | Définition |
|---|---|
| **Action** | Règle "si trigger en PROBLEM alors notifier" (conditions + opérations + escalades) |
| **Agent (actif/passif)** | Programme sur le supervisé ; actif = il envoie, passif = on l'interroge |
| **Agrégat** | Item calculé sur plusieurs hôtes (grpavg, grpsum) |
| **Disponibilité (Availability)** | Icônes ZBX/SNMP/JMX/IPMI vertes/rouges par hôte |
| **Escalade** | Montée en puissance des notifications si pas d'acquittement |
| **Frontend** | Interface web PHP (visualisation + configuration) |
| **Hôte** | Équipement supervisé |
| **Housekeeping** | Nettoyage automatique des vieilles données |
| **Item** | Une métrique collectée (clé + intervalle + type) |
| **LLD** | Low-Level Discovery : découverte automatique (disques, interfaces…) |
| **Macro** | Variable `{$NOM}` (template, hôte, globale) pour seuils et paramètres |
| **Maintenance** | Plage sans notification (mais avec collecte) |
| **Média** | Canal d'alerte (Email, SMS, Telegram, webhook) |
| **MIB** | Dictionnaire SNMP (OID ↔ noms lisibles) |
| **nodata()** | Fonction trigger = 1 si aucune donnée reçue sur la période |
| **OID** | Identifiant numérique d'une donnée SNMP |
| **Poller** | Processus server qui interroge (passif, SNMP…) |
| **Prétraitement** | Transformation de la valeur avant stockage |
| **Proxy** | Collecteur distant avec buffer offline |
| **PSK** | Clé pré-partagée pour chiffrer les flux (TLS) |
| **Template** | Paquet réutilisable (items, triggers, graphes, LLD) |
| **Trapper** | Item alimenté par envoi externe (`zabbix_sender`) |
| **Trends** | Agrégats horaires (min/max/avg/count), longue rétention |
| **Trigger** | Expression logique → OK/PROBLEM |
| **UserParameter** | Clé d'agent personnalisée (script maison) |
| **Value mapping** | Traduction valeur → libellé (5 → "BATTERIE") |
| **vps** | Values per second : valeurs collectées par seconde (dimensionnement) |

## 87. Quiz : 10 questions pour valider (avec réponses)

**Q1.** Quelle est la différence entre un check agent **actif** et **passif** ?
> *R : Passif : le server interroge l'agent (port 10050 côté agent). Actif : l'agent récupère sa liste d'items puis envoie les valeurs au server (port 10051 côté server). L'actif traverse NAT/pare-feu en flux sortant uniquement.*

**Q2.** Un proxy Zabbix calcule-t-il les triggers et envoie-t-il des alertes ?
> *R : Non. Le proxy collecte et bufferise (utile sur lien instable), puis relaie au server. Seul le server évalue les triggers et déclenche les actions.*

**Q3.** Que signifient les valeurs 3 et 5 de `upsOutputSource` (RFC 1628) ?
> *R : 3 = alimentation normale (secteur), 5 = sur batterie. Le trigger Disaster "coupure secteur" teste `=5`.*

**Q4.** Pourquoi faut-il une *recovery expression* différente du seuil de déclenchement ?
> *R : Pour l'hystérésis : éviter le flapping (oscillation PROBLEM/OK) quand la valeur oscille autour du seuil. Ex. déclenche à 90 %, ne se résout qu'à 80 %.*

**Q5.** Dans quel ordre met-on à jour server, proxies et agents lors d'une montée de version majeure ?
> *R : Server d'abord, puis proxies (même version majeure requise), puis agents (peuvent rester en N-1 temporairement). Jamais l'inverse, jamais de saut de 2 majeures.*

**Q6.** À quoi sert `ProxyOfflineBuffer` ?
> *R : Durée pendant laquelle le proxy conserve localement les données quand le lien vers le server est coupé (ex. 24h sur liaison 4G instable). Au retour du lien, les données sont renvoyées : pas de trou dans les graphes.*

**Q7.** Citez 3 moyens de réduire le bruit des alertes.
> *R : Hystérésis (recovery expression), seuils avec durée (avg/count au lieu de last), dépendances entre triggers, maintenances planifiées, plages horaires, acquittement obligatoire. (3 parmi ces réponses.)*

**Q8.** Pourquoi choisir PostgreSQL + TimescaleDB plutôt que MySQL pour un nouveau déploiement ?
> *R : Recommandation officielle Zabbix ; compression native des historiques (÷5 à ÷10 d'espace disque), retention policies automatiques, meilleures performances sur gros volumes. Le partitionnement MySQL est manuel et plus fragile.*

**Q9.** Que vérifiez-vous en premier si un item SNMP est "Not supported" ?
> *R : En CLI depuis le server : `snmpwalk` avec la même communauté et le même OID. Si snmpwalk échoue (timeout → réseau/communauté ; No Such Object → mauvais OID), inutile de chercher côté Zabbix.*

**Q10.** Pourquoi les alertes "Zabbix en panne" doivent-elles partir par un canal indépendant ?
> *R : Parce que si le server ou le réseau principal tombe, l'email classique (qui dépend de la même infrastructure) ne partira pas. Il faut un canal hors dépendance : SMS via passerelle GSM autonome, watchdog externe.*

## 88. Pour aller plus loin

**Documentation officielle** (toujours la référence pour votre version exacte) :
- Doc Zabbix 7.0 : `https://www.zabbix.com/documentation/7.0`
- Matrice de compatibilité (PHP, bases, navigateurs) : page "Requirements"
- Templates officiels : dépôt Git `zabbix/zabbix` (dossier `templates/`)

**À explorer ensuite** :
- **API Zabbix** : automatisez la création d'hôtes (inventaire CMDB → Zabbix), les maintenances, les exports Git.
- **TimescaleDB avancé** : continuous aggregates pour des rapports instantanés sur 2 ans.
- **CA interne + certificats** (plutôt que PSK) si vos flux traversent des réseaux tiers.
- **Zabbix + NUT** : corrélez les événements onduleur Zabbix (SNMP) avec les arrêts automatiques NUT de vos serveurs (voir votre guide onduleurs, section NUT).
- **Tests de charge** : `zabbix_sender` en boucle pour valider le dimensionnement avant la mise en production.
- **Communauté** : forum Zabbix, Zabbix Summit (conférence annuelle), blog officiel.

