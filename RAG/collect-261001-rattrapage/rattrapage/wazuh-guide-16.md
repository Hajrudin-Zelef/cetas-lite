---
id: collect-261001-rattrapage/rattrapage/wazuh-guide-16
title: "Guide Wazuh — SIEM & XDR Open Source en production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "agents", "arr", "incident"]
source: docs/RAG/collect-261001-rattrapage/wazuh_guide.md
source_anchor: ""
source_lines: [2895, 3067]
sha256: ee3a2bd4ebe74ab31d83ac1a67911e8c3fa6cfec9add2152c97253d20cbf9754
---

# Guide Wazuh — SIEM & XDR Open Source en production

### Erreur n°3 — FIM en temps réel sur /var/log ou /tmp
**Symptôme :** explosion d'EPS, manager saturé, alertes en retard.
**Solution :** `realtime` uniquement sur binaires et confs stables ; **jamais** sur les dossiers à forte écriture (section 36).

### Erreur n°4 — Ne pas figer les versions des paquets
**Symptôme :** un `apt upgrade` nocturne casse Wazuh un mardi à 4h.
**Solution :** `apt-mark hold` (section 9) + mises à jour planifiées (section 65).

### Erreur n°5 — Un seul mot de passe admin partagé par toute l'équipe
**Symptôme :** impossible de savoir qui a fait quoi ; mot de passe jamais changé.
**Solution :** un compte par analyste, rôles least-privilege (section 11), rotation après chaque départ.

### Erreur n°6 — Ignorer le tuning (« on regardera les alertes plus tard »)
**Symptôme :** 5 000 alertes/jour, plus personne ne lit, la vraie attaque passe.
**Solution :** rituel de tuning hebdomadaire (section 50). Un SIEM non tuné est un presse-papier coûteux.

### Erreur n°7 — Pas de supervision du SIEM lui-même
**Symptôme :** Filebeat arrêté depuis 3 jours, personne ne l'a vu.
**Solution :** checks Zabbix/Prometheus sur les 4 services + files + espace disque (section 59).

### Erreur n°8 — Sauvegarde jamais testée
**Symptôme :** le jour du sinistre, le snapshot est corrompu ou le dépôt inaccessible.
**Solution :** restauration testée 2×/an sur environnement isolé (sections 63–64).

### Erreur n°9 — Cloner des VM avec un agent déjà inscrit
**Symptôme :** doublons d'ID, agents qui se « volent » la connexion, alertes fantômes.
**Solution :** image avec agent **non inscrit**, inscription au premier boot (section 20, méthode D).

### Erreur n°10 — Exposer le dashboard ou l'API sur Internet sans protection
**Symptôme :** interface d'admin accessible au monde entier.
**Solution :** dashboard/API uniquement via VPN ou IP allowlistées ; TLS partout ; fail2ban ; 2FA si possible via reverse proxy.

---

## 74. Bonnes pratiques SOC pour une petite équipe

Vous êtes 1 à 3 personnes, pas un SOC de 30 analystes en 3×8. Adaptez :

### 74.1 Organisation

- **Un responsable SOC** (vous) : tuning, règles, astreinte niveau 2.
- **Un analyste** (même à temps partiel) : triage quotidien, investigations niveau 1.
- **Un suppléant** : connaît la procédure d'astreinte, peut qualifier une alerte critique.

### 74.2 Les rituels (non négociables)

| Rituel | Fréquence | Durée | Contenu |
|---|---|---|---|
| Triage alertes | Quotidien | 30 min | Qualifier les ≥ 7 de la veille (section 48) |
| Revue vulnérabilités | Hebdomadaire | 1 h | CVE critiques/hautes → plan de patch (section 40) |
| Tuning | Hebdomadaire | 1 h | Top bruit → actions (section 50) |
| Revue MITRE | Mensuelle | 1 h | Couverture par tactique, règles manquantes |
| Test restauration | Semestrielle | ½ journée | Procédure section 64 |
| Exercice incident | Semestrielle | ½ journée | Simulez un cas pratique (sections 75–80) |

### 74.3 Documentation minimale

1. **Procédure d'astreinte** (1 page) : qui appeler, seuils, actions immédiates (tableau section 49).
2. **Cartographie** : quels agents, quels groupes, quelles règles custom (pointez vers le git, section 62).
3. **Journal des incidents** : date, alerte, qualification, action. Même un tableur suffit.
4. **Journal du tuning** : quelle règle modifiée, pourquoi, quand (git log + commentaires XML).

### 74.4 Hygiène des alertes

- **Zéro alerte ≥ 10 « Nouveau » de plus de 24 h** (section 48).
- Chaque faux positif récurrent = **ticket de tuning**, pas « on fera avec ».
- Célébrez les détections : une attaque bloquée par l'active response = à raconter en réunion d'équipe. Ça entretient la vigilance.

### 74.5 Articulation avec l'équipe systèmes

Le SOC ne patch pas, ne redémarre pas les serveurs de prod : il **détecte et qualifie**, puis transmet avec un niveau de sévérité et un délai (tableau section 40). Formalisez ce contrat : qui fait quoi, sous quel délai, avec quelle escalade.

---

## 75. Cas pratique 1 : intrusion SSH par force brute

**Contexte :** `srv-web-01` (10.0.1.15) exposé en SSH sur Internet (à corriger !).

### Ce que Wazuh voit

```
1. Règle 5716 (niveau 5) : "SSHD: authentication failed." × 40 en 2 min, srcip 203.0.113.7
2. Règle 5763 (niveau 10) : "sshd: brute force attack." → ACTIVE RESPONSE firewall-drop
3. L'IP est bannie 10 min (active-responses.log)
4. 2 h plus tard : règle 5718 (niveau 8) : "SSHD: authentication success." depuis 198.51.100.99 !
```

### Investigation pas à pas

```bash
# 1. Historique de l'IP sur 7 jours : a-t-elle touché d'autres machines ?
# Dashboard Discover : data.srcip : "198.51.100.99"

# 2. Sur srv-web-01 : que s'est-il passé autour du succès ?
sudo grep "198.51.100.99" /var/log/auth.log | tail -20
sudo last -i | head -20

# 3. L'utilisateur compromis a-t-il fait quelque chose ?
sudo grep "jdupont" /var/log/auth.log | grep -E "sudo|su " | tail -20
```

### Réponse

1. **Isoler** : coupez le SSH entrant au pare-feu (ou `firewall-drop` manuel via API, section 58).
2. **Bloquer** l'utilisateur : `usermod -L jdupont`, forcez la rotation des clés/mots de passe.
3. **Chercher la persistance** : FIM sur `/root/.ssh/authorized_keys`, cron, services (règles 550-557).
4. **Corriger la cause** : SSH derrière VPN/bastion, clés uniquement, fail2ban + Wazuh active response en ceinture-bretelles.

### Leçons

- Le niveau 8 « premier succès après échecs » est **plus important** que le brute force lui-même : c'est lui qui doit réveiller l'astreinte.
- Sans allowlist correcte, l'active response aurait pu bannir votre bastion : vérifiez la section 45.

---

## 76. Cas pratique 2 : ransomware simulé (FIM + active response)

**Contexte :** exercice sur une VM de lab `lab-win-01`. On simule un ransomware qui chiffre `C:\Data`.

### Détection attendue

```
1. FIM temps réel : dizaines de fichiers modifiés en quelques secondes (règles 550/551, niveau 7)
   → règle custom "rafale FIM" (fréquence 50 / 60 s) niveau 13
2. Nouveau processus avec extension suspecte (syscollector / command wodle)
3. Active response : quarantine.sh → coupe le réseau de la VM
```

### Règle « rafale FIM » (à préparer AVANT l'exercice)

```xml
<rule id="100800" level="13" frequency="50" timeframe="60">
  <if_matched_sid>550</if_matched_sid>
  <same_field>id</same_field>   <!-- même agent -->
  <description>Rafale FIM : 50+ fichiers modifiés en 60s sur $(agent_name) — ransomware probable</description>
  <mitre><id>T1486</id></mitre>  <!-- Data Encrypted for Impact -->
</rule>
```

```xml
<active-response>
  <command>quarantine-net</command>
  <location>local</location>
  <rules_id>100800</rules_id>
</active-response>
```

### Déroulé de l'exercice

1. Lancez le simulateur (script qui renomme/chiffre des fichiers de test).
2. Chronométrez : temps entre le début du chiffrement et l'alerte niveau 13.
3. Vérifiez la quarantaine réseau : la VM ne parle plus ?
4. **Restaurez** depuis snapshot : mesurez le RTO.
5. Débriefez : qu'est-ce qui a manqué ? (souvent : la règle rafale n'existait pas, ou le FIM n'était pas en temps réel sur `C:\Data`).

> Refaites cet exercice **2×/an**. C'est le meilleur test de votre chaîne détection → réponse.

---

## 77. Cas pratique 3 : serveur web compromis (PHP webshell)

**Contexte :** `srv-web-02` héberge une application PHP. Un webshell `x.php` est déposé via une faille d'upload.

### Signaux Wazuh (dans l'ordre)

```
1. FIM temps réel /var/www/html : "New file" x.php (règle 554, niveau 7)
2. Logs nginx : POST /upload.php puis GET /x.php?cmd=id (règle custom, niveau 10)
3. Command wodle : processus php-fpm lançant /bin/sh (anormal → niveau 12)
```

### Règle custom « accès webshell »

