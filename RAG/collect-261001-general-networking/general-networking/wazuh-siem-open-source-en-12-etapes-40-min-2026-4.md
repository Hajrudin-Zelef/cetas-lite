---
id: collect-261001-general-networking/general-networking/wazuh-siem-open-source-en-12-etapes-40-min-2026-4
title: "Redémarre l'agent et relance les scans SCA"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "agents", "attention", "incident"]
source: docs/RAG/collect-261001-general-networking/wazuh-siem-open-source-en-12-etapes-40-min-2026.md
source_anchor: ""
source_lines: [219, 293]
sha256: 006fd491d7915b18a8d7c8a5d83c1f03985b19b7e344beba444cf78c7b22725c
---

# Redémarre l'agent et relance les scans SCA

```
<localfile>
  <log_format>syslog</log_format>
  <location>/var/log/auth.log</location>
</localfile>
```
Pour simuler une attaque, tentez plusieurs connexions SSH échouées depuis une autre machine (par exemple `ssh [email protected]` avec un mauvais mot de passe, répété rapidement). Wazuh corrèle les échecs successifs et déclenche la règle **5712** (« brute force ») dès que le seuil est franchi. Voici une alerte réelle telle qu’elle apparaît dans les journaux d’alertes (`/var/ossec/logs/alerts/alerts.log`) :

```
** Alert 1751709322.148763: - syslog,sshd,authentication_failures,
2026 Jul 05 10:15:22 (web01) 10.0.0.21->/var/log/auth.log
Rule: 5712 (level 10) -> 'sshd: brute force trying to get access to the system.'
Src IP: 203.0.113.45
Src User: root
Jul  5 10:15:22 web01 sshd[20481]: Failed password for root from 203.0.113.45 port 54122 ssh2
```
Décryptage : la règle 5712 est de niveau 10 (élevé), l’IP source est identifiée, l’utilisateur ciblé aussi. Dans le tableau de bord, module *Threat Hunting*, l’alerte est enrichie et interrogeable. Vous venez de *voir* une attaque. Reste à la *stopper* automatiquement.

## Étape 9 – Automatiser la réponse : bannir l’attaquant

La réponse active (Active Response) de Wazuh exécute automatiquement une commande lorsqu’une alerte franchit un seuil. Le script intégré `firewall-drop` ajoute l’IP malveillante aux règles de pare-feu local (iptables/nftables) pour une durée déterminée. Configurons le bannissement automatique sur détection de force brute SSH. Sur le manager, dans `/var/ossec/etc/ossec.conf`, ajoutez :

```
<active-response>
  <command>firewall-drop</command>
  <location>local</location>
  <rules_id>5710,5712,5763</rules_id>
  <timeout>600</timeout>
</active-response>
```
Ici, dès qu’une des règles 5710, 5712 ou 5763 se déclenche, l’agent bannit l’IP source pendant 600 secondes (10 minutes), puis la débloque automatiquement. Redémarrez le manager avec `sudo systemctl restart wazuh-manager`, puis relancez votre attaque de test. Surveillez le journal de réponse active sur l’agent :

`sudo tail -f /var/ossec/logs/active-responses.log````
2026/07/05 10:15:23 wazuh-execd: Starting firewall-drop
2026/07/05 10:15:23 active-response/bin/firewall-drop: Adding rule for 203.0.113.45
2026/07/05 10:25:23 active-response/bin/firewall-drop: Removing rule for 203.0.113.45
```
Félicitations : vous avez une chaîne détection → réponse entièrement automatisée. L’attaquant est neutralisé en quelques secondes, sans intervention humaine. Pour une défense en profondeur, combinez cette réponse active avec un blocage réseau en amont : notre guide CrowdSec mutualise les listes d’IP malveillantes à l’échelle communautaire, un excellent complément à Wazuh.

## Étape 10 – Cartographier les alertes sur MITRE ATT&CK

Détecter, c’est bien ; comprendre *quelle technique adverse* a été employée, c’est mieux. Wazuh mappe nativement ses règles sur le référentiel MITRE ATT&CK, la base de connaissances mondiale des tactiques et techniques d’attaquants. Notre attaque SSH par force brute, par exemple, est associée à la technique **T1110 (Brute Force)**.

Dans le tableau de bord, le module *MITRE ATT&CK* présente une matrice interactive : chaque colonne est une tactique (Accès initial, Exécution, Persistance, Exfiltration…), et les alertes s’y positionnent en temps réel. Vous visualisez ainsi la *progression* d’un attaquant dans votre système d’information, et non plus une simple liste d’événements isolés. C’est le langage commun des analystes SOC et des rapports d’incident.

Concrètement, cette cartographie transforme Wazuh en outil de *threat hunting* : filtrez par technique, identifiez les tactiques les plus sollicitées contre votre parc, et priorisez vos défenses. Pour un rapport NIS 2 ou une investigation post-incident, pouvoir dire « l’attaquant a utilisé T1110 puis T1078 (Valid Accounts) » apporte une rigueur que les décideurs et les autorités attendent désormais.

## Étape 11 – Sécuriser et fiabiliser votre déploiement Wazuh

Un SIEM est une cible de choix : il concentre les journaux et les secrets de toute l’infrastructure. Le durcir est impératif. Voici les mesures prioritaires après une installation Wazuh.

- **Changer tous les mots de passe par défaut** avec`wazuh-passwords-tool.sh` , y compris ceux des comptes internes de l’indexeur.
- **Remplacer les certificats auto-signés** par des certificats de confiance (Let’s Encrypt ou PKI interne) via l’outil`wazuh-certs-tool.sh` , pour supprimer les avertissements navigateur et chiffrer proprement.
- **Filtrer les ports** 55000, 9200 et 443 au pare-feu : n’exposez jamais l’API ni l’indexeur à Internet. Réservez l’accès au tableau de bord à un VPN.
- **Isoler les flux inter-sites** dans un tunnel chiffré. Pour relier des agents distants au manager, faites transiter le trafic par un VPN WireGuard ou un réseau maillé Tailscale.
- **Centraliser l’authentification** du tableau de bord via SSO. Wazuh s’intègre à un fournisseur d’identité OIDC/SAML comme Keycloak, pour appliquer MFA et politiques d’accès.
- **Planifier les sauvegardes** de`/var/ossec/etc` , des règles personnalisées et des index critiques.

Pensez également à la rétention. L’indexeur peut saturer le disque : mettez en place une politique de gestion du cycle de vie des index (Index State Management) pour supprimer ou archiver automatiquement les données au-delà de votre durée légale de conservation. Nous y revenons dans la section conformité et dans les astuces avancées.

## Étape 12 – Aligner Wazuh sur NIS 2 et le RGPD

Dernière étape, et non des moindres pour une organisation française ou européenne : transformer votre déploiement Wazuh en actif de conformité. La directive NIS 2 impose aux entités essentielles et importantes de détecter les incidents, de conserver des preuves et de notifier l’ANSSI sous 24 heures (alerte précoce) puis 72 heures (notification complète). Wazuh adresse directement ces obligations.

- **Détection et journalisation** : la corrélation en temps réel et l’archivage des événements (`/var/ossec/logs/archives/` ) fournissent la traçabilité exigée.
- **Gestion des vulnérabilités** : le module dédié documente en continu votre exposition, un pilier de NIS 2.
- **Preuves d’incident** : les alertes horodatées, enrichies MITRE et exportables constituent un dossier de notification prêt à l’emploi.
- **Souveraineté et RGPD** : parce que Wazuh est auto-hébergé, les journaux – qui contiennent souvent des données personnelles (IP, identifiants) – restent sur votre infrastructure, en Europe, sous votre seul contrôle. Aucun transfert vers un cloud tiers.

Attention à la contrepartie RGPD : les journaux étant des données personnelles, définissez une durée de conservation proportionnée et documentée (souvent 6 à 12 mois pour les journaux de sécurité), et restreignez leur accès. La politique de rétention des index n’est donc pas qu’une question d’espace disque : c’est aussi une obligation légale. Bien configuré, Wazuh devient la colonne vertébrale de votre conformité, à coût de licence nul. Le cadre européen est détaillé par l’ENISA et l’ANSSI.

## 6 pièges courants à éviter avec Wazuh

La plupart des déploiements Wazuh qui échouent butent sur les mêmes erreurs. Les connaître vous fera gagner des heures.

