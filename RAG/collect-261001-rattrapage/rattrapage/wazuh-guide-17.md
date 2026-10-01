---
id: collect-261001-rattrapage/rattrapage/wazuh-guide-17
title: "Guide Wazuh — SIEM & XDR Open Source en production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "agents", "incident", "intel"]
source: docs/RAG/collect-261001-rattrapage/wazuh_guide.md
source_anchor: ""
source_lines: [3068, 3256]
sha256: 513dbb6ab997282f990396b0f8c250102353fae835a52d41ef7eb21febc94cd2
---

# Guide Wazuh — SIEM & XDR Open Source en production

```xml
<rule id="100900" level="12">
  <decoded_as>nginx-access</decoded_as>
  <regex>GET \/\w+\.php\?cmd=</regex>
  <description>Requête type webshell détectée : $(url)</description>
  <mitre><id>T1505.003</id></mitre>
</rule>
```

> Adaptez `decoded_as` à votre decoder nginx réel (testez avec wazuh-logtest, section 31).

### Investigation

```bash
# 1. Dater le dépôt : FIM donne l'heure exacte de création
# 2. Trouver le point d'entrée : logs nginx autour de cette heure
sudo grep "POST" /var/log/nginx/access.log | grep -B2 -A2 "upload"
# 3. Contenu du webshell (copie forensique AVANT suppression)
sudo cp /var/www/html/x.php /root/forensique/x.php.$(date +%F)
sudo head -50 /root/forensique/x.php.*
# 4. Chercher d'autres fichiers déposés ± 24 h (FIM "new file" sur le webroot)
```

### Réponse

1. Isoler le serveur (pare-feu / VLAN quarantaine).
2. **Ne pas juste supprimer le webshell** : considérez la machine compromise → rebuild depuis une image saine.
3. Corriger la faille (upload non filtré), patcher l'application.
4. Ajouter le hash du webshell à vos listes (CDB) pour détecter sa réapparition ailleurs.

---

## 78. Cas pratique 4 : conformité d'un parc Windows (CIS)

**Objectif :** passer le parc `windows-postes` de 60 % à 90 % de contrôles CIS passants en 3 mois.

### Méthode

```
Semaine 1 : activer SCA, laisser tourner un cycle complet (12 h)
Semaine 2 : exporter les fails par contrôle (dashboard SCA → export CSV)
            → trier par occurrence : les 5 contrôles les plus échoués = 80 % du bruit
Semaines 3-10 : corriger par vague via GPO / Intune :
            - mot de passe complexe + verrouillage (CIS 1.x)
            - désactiver les comptes invités (CIS 2.x)
            - activer l'audit avancé (CIS 17.x)
Semaine 11 : re-mesurer, publier le score à la direction
En continu : alerte si le score d'un poste chute brutalement (dérive)
```

### Suivi via API (script hebdo)

```bash
#!/bin/bash
# Extrait : moyenne des scores SCA du groupe windows-postes
for id in $(curl -sk -H "Authorization: Bearer $TOKEN" \
  'https://localhost:55000/agents?group=windows-postes&select=id&limit=500' \
  | grep -o '"id":"[0-9]*"' | cut -d'"' -f4); do
  curl -sk -H "Authorization: Bearer $TOKEN" \
    "https://localhost:55000/sca/$id" | grep -o '"score":[0-9]*' | head -1
done
```

> Présentez le score CIS comme un **KPI sécurité** en comité de direction : c'est concret, mesurable, et ça justifie les budgets.

---

## 79. Cas pratique 5 : vulnérabilité critique sur le parc (CVE)

**Contexte :** CVE-2026-XXXX, CVSS 9.8, RCE sur un composant présent sur 40 serveurs (annonce un mardi).

### Déroulé avec Wazuh

```
Mardi 10h  : l'annonce tombe.
Mardi 10h05 : Dashboard → Vulnerability Detection → recherche "CVE-2026-XXXX"
             → 40 agents concernés, dont 6 en DMZ (exposés Internet).
Mardi 11h  : patch des 6 DMZ (fenêtre d'urgence) + mitigation (pare-feu) sur les autres.
Mercredi   : patch du reste par vagues, vérification : le CVE disparaît de la liste.
Jeudi      : rapport : 40/40 patchés, preuves (exports dashboard) pour l'audit.
```

### Automatiser la veille

```bash
# Cron quotidien : lister les CVE critiques apparues dans les dernières 24 h
# (via l'API indexer : recherche sur wazuh-states-vulnerabilities-*)
curl -sk -u admin:'<MDP>' 'https://localhost:9200/wazuh-states-vulnerabilities-*/_search?pretty' \
  -H 'Content-Type: application/json' -d '{
    "query": { "range": { "vulnerability.published_at": { "gte": "now-1d" } } },
    "size": 50,
    "_source": ["vulnerability.id", "vulnerability.score.base", "agent.name", "package.name"]
  }'
```

> Le module Vulnerability Detector met à jour ses bases **toutes les heures** (section 39) : un CVE publié le matin est visible l'après-midi. C'est votre radar.

---

## 80. Cas pratique 6 : exfiltration de données via DNS

**Contexte :** un poste `pc-compta-03` envoie des requêtes DNS anormalement longues vers un domaine rare — technique classique d'exfiltration (DNS tunneling).

### Détection

Wazuh ne fait pas nativement de la détection DNS fine, mais :

1. Collectez les logs DNS (serveur DNS interne en syslog, section 52, ou logs du resolver du poste).
2. Decoder + règle sur la **longueur** et l'**entropie** des requêtes :

```xml
<!-- Exemple : alerter sur des requêtes DNS > 60 caractères vers l'extérieur -->
<rule id="101000" level="10">
  <decoded_as>dns-query</decoded_as>
  <regex>\w{60,}\.</regex>
  <description>Requête DNS anormalement longue : exfiltration possible ($(query))</description>
  <mitre><id>T1048.003</id></mitre>
</rule>
```

### Investigation

```
1. Volume : combien de requêtes ? quelle taille totale ? (dashboard : agg par query)
2. Domaine : nouvellement enregistré ? (whois, threat intel)
3. Poste : processus à l'origine ? (syscollector processes, EDR)
4. Données : quelles données auraient pu partir ? (DCP / classification)
```

### Réponse

1. Bloquer le domaine au DNS/pare-feu.
2. Isoler le poste, image forensique.
3. Si exfiltration confirmée : procédure incident + notifications légales (données personnelles ?).

> Ce cas montre la limite ET la force de Wazuh : pas de NDR natif, mais avec un bon decoder sur vos logs DNS existants, vous couvrez 80 % du besoin pour 20 % de l'effort.

---

## 81. Checklist quotidienne / hebdomadaire / mensuelle

### Quotidienne (30 min, analyste)

- [ ] Dashboard → *Overview* : alertes niveau ≥ 10 des dernières 24 h qualifiées ?
- [ ] Agents `disconnected` : justifiés (maintenance) ou à investiguer ?
- [ ] Santé : manager, indexer (`_cluster/health`), Filebeat, espace disque
- [ ] Active responses de la nuit : légitimes ? (log `/var/ossec/logs/active-responses.log`)
- [ ] Zéro alerte ≥ 10 restée « Nouveau »

### Hebdomadaire (2 h, responsable SOC)

- [ ] Revue CVE critiques/hautes → avancement patch (section 40)
- [ ] Top 10 règles bruyantes → 1–2 actions de tuning (section 50)
- [ ] Revue des comptes : nouveaux utilisateurs dashboard/API ?
- [ ] Vérifier que les sauvegardes Borg + snapshots se sont bien déroulées
- [ ] Journal d'incidents à jour

### Mensuelle (½ journée)

- [ ] Revue MITRE : couverture par tactique, règles manquantes (section 47)
- [ ] Score SCA moyen par groupe : dérive ? (section 42)
- [ ] Relecture des règles `overwrite` : toujours pertinentes ? (section 32)
- [ ] Rotation des logs/archivage conforme à la rétention légale
- [ ] Test d'une alerte de bout en bout (fausse attaque en lab → alerte → notification)
- [ ] Mise à jour planifiée si version corrective disponible (section 65)

### Trimestrielle / semestrielle

- [ ] Test de restauration (section 64)
- [ ] Exercice incident type ransomware (section 76)
- [ ] Revue des accès (comptes dashboard, API, clés d'agents)
- [ ] Rotation des clés d'agents critiques (section 21)

---

## 82. Pense-bête de poche : commandes essentielles

```bash
# ── Services ──────────────────────────────────────────────
sudo systemctl status wazuh-manager wazuh-indexer wazuh-dashboard filebeat
sudo systemctl restart wazuh-manager        # après modif de conf
sudo /var/ossec/bin/wazuh-control status    # statut interne du manager

# ── Agents ────────────────────────────────────────────────
sudo /var/ossec/bin/agent_control -l        # liste + états
sudo /var/ossec/bin/agent_control -i 007    # détail d'un agent
sudo /var/ossec/bin/agent_groups -l         # groupes
sudo /var/ossec/bin/manage_agents -l        # clés

