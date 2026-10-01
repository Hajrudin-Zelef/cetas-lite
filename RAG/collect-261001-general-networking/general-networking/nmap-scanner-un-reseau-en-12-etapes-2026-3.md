---
id: collect-261001-general-networking/general-networking/nmap-scanner-un-reseau-en-12-etapes-2026-3
title: "Nmap version 7.991 ( https://nmap.org )"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "cyber", "incident", "open source"]
source: docs/RAG/collect-261001-general-networking/nmap-scanner-un-reseau-en-12-etapes-2026.md
source_anchor: ""
source_lines: [159, 220]
sha256: 4e288ae2361e9f6ca68f9287f1d009bd6acadcbf31dc147374eaf8eb4babfdb6
---

# Nmap version 7.991 ( https://nmap.org )

`sudo nmap -sV --script vulners 192.168.1.10`
Parmi les scripts les plus utilisés en audit défensif en 2025-2026 : `smb-vuln-ms17-010` (détection d’EternalBlue), `ssl-heartbleed`, `ssl-enum-ciphers` (audit de la configuration TLS), `http-enum` (énumération de répertoires web) et `dns-zone-transfer` (détection de transferts de zone DNS mal configurés). La documentation complète est disponible sur le portail NSEDoc officiel.

## Cartographier un environnement Windows et Active Directory

Un cas d’usage très fréquent en entreprise consiste à cartographier un domaine Active Directory avant un audit ou une migration. Nmap dispose de plusieurs scripts NSE dédiés au protocole SMB qui permettent d’identifier les contrôleurs de domaine, les partages accessibles et les comptes exposés sans authentification :

`sudo nmap -p 88,389,445,3268 --script smb-os-discovery,smb-enum-shares 192.168.1.0/24`
Cette commande cible les ports caractéristiques d’un environnement Windows : Kerberos (88), LDAP (389), SMB (445) et le catalogue global (3268). Le script `smb-os-discovery` révèle la version exacte de Windows Server installée sur chaque machine répondante, une information précieuse quand on sait que des failles critiques comme Netlogon (CVE visant les contrôleurs de domaine non patchés) continuent de circuler activement sur des parcs mal maintenus. Le script `smb-enum-shares`, de son côté, liste les partages réseau accessibles, y compris ceux ouverts par erreur à des utilisateurs non authentifiés, un défaut de configuration que l’on retrouve encore régulièrement lors d’audits menés en 2026.

Sur ce type de scan, restez particulièrement prudent : certains scripts de la catégorie `brute` ou `intrusive` peuvent verrouiller des comptes Active Directory après plusieurs tentatives d’authentification échouées, avec un impact direct sur la production. Limitez-vous aux catégories `default`, `discovery` et `safe` lors d’un premier passage, et réservez les scripts plus agressifs à une fenêtre de maintenance validée avec les équipes infrastructure.

## Nmap et les obligations de la directive NIS2

Depuis l’entrée en application de la **directive NIS2** pour les entités essentielles et importantes en France et dans l’Union européenne, la cartographie régulière du parc réseau n’est plus seulement une bonne pratique : c’est une composante attendue de la gestion des risques cyber documentée dans le référentiel de conformité. Un audit Nmap planifié, avec ses rapports XML horodatés et archivés, constitue une preuve tangible de la démarche de gouvernance exigée par le texte, que nous détaillons dans notre guide complet sur la mise en conformité NIS2.

Concrètement, un scan Nmap ne suffit jamais à lui seul à démontrer une conformité NIS2 complète : il doit s’insérer dans un cycle plus large associant gestion des vulnérabilités, plan de réponse à incident et formation des équipes. Mais il reste souvent le point de départ le plus simple à mettre en place pour un service informatique qui découvre ses obligations : avant de documenter des risques, encore faut-il savoir précisément quelles machines et quels services composent le système d’information.

## Étape 10 : exporter et exploiter les résultats de scan

Un scan ponctuel dans un terminal n’a de valeur que sur l’instant. Pour construire un historique ou alimenter un autre outil, Nmap propose plusieurs formats de sortie :

```
# Format normal, lisible
nmap -oN scan_resultat.txt 192.168.1.0/24
# Format XML, exploitable par des scripts et des SIEM
nmap -oX scan_resultat.xml 192.168.1.0/24
# Format "greppable", pratique avec grep/awk
nmap -oG scan_resultat.gnmap 192.168.1.0/24
# Les trois formats en une seule commande
nmap -oA audit_complet 192.168.1.0/24
```
Le format XML (`-oX`) est celui que consomment la majorité des intégrations tierces, y compris la bibliothèque Python `python-nmap` utilisée dans les pipelines SIEM. C’est celui que nous utiliserons dans le projet complet ci-dessous.

## Étape 11 : intégrer Nmap à un SIEM, l’exemple Wazuh

En 2026, la tendance forte n’est plus le scan manuel isolé mais l’intégration de Nmap dans un pipeline de surveillance continue. **Wazuh**, la plateforme SIEM/XDR open source déjà couverte dans notre tutoriel Wazuh, illustre bien ce pattern grâce à sa fonctionnalité de *command monitoring*. On ajoute dans le fichier `ossec.conf` de l’agent Wazuh un bloc qui exécute un script Nmap à intervalle régulier :

```
<localfile>
  <log_format>full_command</log_format>
  <command>python3 /home/audit/nmap_scan.py</command>
  <frequency>604800</frequency> <!-- hebdomadaire -->
</localfile>
```
Le script Python, généralement bâti sur la bibliothèque `python-nmap`, exécute le scan, convertit les résultats en JSON et les ajoute au fichier `/var/ossec/logs/active-responses.log`. Wazuh ingère ensuite ce flux et applique des règles personnalisées pour générer des alertes dès qu’un nouveau port apparaît sur un hôte surveillé. C’est exactement le pattern documenté par l’équipe Wazuh dans son article sur l’audit de sécurité automatisé, qui combine scan Nmap planifié et enrichissement des résultats. À l’inverse, Wazuh peut aussi **détecter des scans Nmap hostiles** dirigés contre vos propres machines, via une règle qui repère les signatures caractéristiques dans les journaux réseau :

```
<rule id="100100" level="10">
  <if_sid>576</if_sid>
  <match>Nmap|Ncat|Nping</match>
  <description>Outil de scan de ports détecté</description>
</rule>
```
Ce double usage, sonde active d’un côté et signal de détection de l’autre, résume bien pourquoi Nmap reste central dans une stratégie de sécurité réseau en 2026 : le même outil sert autant à l’attaque simulée qu’à la défense. Plus de détails dans notre article dédié à l’intégration Nmap et Wazuh.

## Étape 12 : projet complet, un script d’audit réseau automatisé avec alerte

Voici un projet fonctionnel de bout en bout, à adapter à votre propre réseau domestique ou labo de test. L’objectif : scanner un sous-réseau chaque nuit, comparer le résultat avec le scan précédent, et alerter par e-mail si un nouveau port ouvert apparaît, un signe classique de compromission ou de mauvaise configuration.

Créez d’abord le script Python `audit_reseau.py` :

