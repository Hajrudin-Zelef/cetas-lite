---
id: collect-261001-huawei/huawei/security-onion-surveillance-reseau-nsm-en-13-etapes-3
title: "Calculer le hash SHA256 de l'ISO téléchargée"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "open source"]
source: docs/RAG/collect-261001-huawei/security-onion-surveillance-reseau-nsm-en-13-etapes.md
source_anchor: ""
source_lines: [124, 196]
sha256: 08066cad64b51f4cbf6decf387b3c2bdeed87a6d7abd90eea668ddd95c21217e
---

# Calculer le hash SHA256 de l'ISO téléchargée

Suricata fonctionne par signatures : il compare chaque paquet réseau à une base de règles qui décrivent des schémas d’attaque connus (exploitation d’une CVE spécifique, communication vers un serveur de commande et contrôle identifié, tentative d’exfiltration par un protocole détourné). C’est un moteur rapide et fiable sur les menaces déjà documentées, mais par nature incapable de repérer une attaque totalement inédite qui ne correspond à aucune signature existante. Zeek, à l’inverse, ne cherche pas de signatures : il journalise en continu tous les métadonnées de chaque connexion réseau (protocole utilisé, volume échangé, certificats TLS vus, requêtes DNS, fichiers transférés) sous forme de logs structurés. C’est ce moteur qui permet la threat hunting rétrospective, en recherchant a posteriori un comportement anormal qu’aucune signature n’aurait détecté au moment des faits.

Wazuh complète ce duo réseau par une surveillance côté hôte : intégrité des fichiers système, analyse des journaux d’authentification, détection de rootkits, conformité à des référentiels de durcissement. C’est le seul des trois moteurs à s’exécuter directement sur les postes et serveurs surveillés, via un agent léger installé sur chaque machine, plutôt qu’en écoute passive sur le réseau. Enfin, la suite Elastic (Elasticsearch pour l’indexation, Logstash pour l’ingestion et la transformation des données, Kibana comme moteur de visualisation sous-jacent à la console SOC) sert de colonne vertébrale technique : c’est elle qui stocke, indexe et rend interrogeables en quelques secondes des dizaines de millions d’événements générés chaque jour par les trois moteurs de détection.

Cette répartition des rôles explique pourquoi une alerte critique dans Security Onion croise très souvent plusieurs sources : une connexion sortante suspecte détectée par Zeek, confirmée par une signature Suricata correspondant à un malware connu, et corrélée à une modification de fichier système détectée par l’agent Wazuh sur le poste à l’origine du trafic. C’est cette corrélation multi-source, impossible à obtenir avec un seul outil pris isolément, qui constitue la véritable valeur ajoutée de la plateforme par rapport à un déploiement de chaque brique en silo.

## Étape 9 : intégrer CyberChef et Playbook pour l’analyse avancée

Security Onion embarque directement CyberChef, accessible depuis le menu principal de la console, pour décoder rapidement des payloads suspects extraits d’une alerte : base64, hexadécimal, chiffrement simple, extraction de fichiers depuis une capture réseau. C’est un gain de temps considérable par rapport à devoir exporter les données vers un outil externe.

Le module Playbook, quant à lui, gère les règles de détection actives. Chaque règle Suricata ou Wazuh peut être activée, désactivée ou modifiée directement depuis l’interface, sans toucher aux fichiers de configuration bruts sur le serveur. C’est particulièrement utile pour ajuster le taux de faux positifs les premières semaines suivant le déploiement, une phase de calibrage que beaucoup d’équipes sous-estiment.

## Étape 10 : configurer la rétention des données et le stockage

Par défaut, Elasticsearch conserve les index selon une politique de rétention basée sur l’espace disque disponible : les données les plus anciennes sont automatiquement purgées lorsque le seuil configuré est atteint. Pour une conformité NIS2 ou un besoin d’investigation forensique, ajustez cette politique dans le fichier de configuration du gestionnaire d’index.

```
# Consulter l'espace disque utilisé par les index Elasticsearch
sudo so-elastic-index-size
# Forcer un rafraîchissement de la politique de rétention
sudo salt-call state.apply elasticsearch.ilm_policy
```
Avec 200 Go de disque, comptez généralement une rétention de quelques jours à deux semaines pour un trafic soutenu, ou plusieurs mois pour un petit segment réseau peu chargé. La documentation officielle recommande de dimensionner le stockage en fonction du volume de trafic réel observé pendant les deux premières semaines d’exploitation plutôt que de se fier à une estimation théorique.

## Étape 11 : ajouter un nœud sensor distant (déploiement distribué)

Si votre besoin dépasse un seul site, Security Onion permet d’ajouter des nœuds sensor distants qui remontent leurs données vers le manager central. Sur chaque nouveau site, relancez l’installateur ISO, mais choisissez cette fois le type de nœud Sensor au lieu de Standalone, puis renseignez l’adresse IP et la clé d’enregistrement générée sur le manager.

```
# Sur le manager : générer une clé d'enregistrement pour un nouveau sensor
sudo so-user add-sensor --site "Agence-Lyon"
# Sur le nouveau sensor : rejoindre le cluster existant
sudo so-setup --join-manager 192.168.10.5
```
Cette architecture distribuée est celle recommandée pour toute organisation multi-sites, car elle permet de centraliser la corrélation d’alertes tout en gardant la capture de trafic au plus près de chaque segment réseau surveillé.

## Étape 12 : sécuriser l’accès et gérer les comptes utilisateurs

Une fois l’installation validée, créez des comptes nominatifs pour chaque analyste plutôt que de partager le compte administrateur initial. La console SOC gère nativement des rôles (analyste, administrateur, lecture seule) qui permettent de limiter les actions possibles selon le niveau de responsabilité.

```
# Ajouter un nouvel analyste avec un rôle restreint
sudo so-user add analyste.dupont --role analyst
# Lister les comptes existants et leurs rôles
sudo so-user list
```
Pensez également à restreindre l’accès à la console SOC via un pare-feu ou un VPN plutôt que de l’exposer directement sur Internet. L’outil manipule des données réseau sensibles, et sa console d’administration ne devrait jamais être accessible depuis l’extérieur sans une couche d’authentification supplémentaire.

## Étape 13 : maintenir la plateforme à jour

Security Onion publie des mises à jour régulières qui incluent aussi bien des correctifs de sécurité que des mises à jour de signatures Suricata et Zeek. La commande de mise à jour officielle gère l’ensemble du cycle, y compris la mise à jour des conteneurs Docker sous-jacents.

```
# Lancer la mise à jour complète de la plateforme
sudo soup
# Vérifier la version actuellement installée
sudo so-version
```
Planifiez cette mise à jour lors d’une fenêtre de maintenance : certaines montées de version majeures nécessitent un redémarrage complet des services et peuvent interrompre la capture de trafic pendant quelques minutes.

## Comparatif : Security Onion face aux alternatives

Avant de vous engager, il est utile de situer Security Onion par rapport aux options que vous connaissez peut-être déjà, notamment si votre équipe a testé Suricata, Zeek ou Wazuh séparément.

| Solution | Coût | Composants inclus | Effort d’intégration | 
|---|---|---|---|
| Security Onion | Gratuit, open source | Suricata + Zeek + Wazuh + Elastic Stack + SOC préintégrés | Faible, tout est packagé | 
| Suricata seul | Gratuit, open source | IDS/IPS réseau uniquement | Élevé, corrélation et visualisation à construire | 
| Zeek seul | Gratuit, open source | Analyse de protocoles uniquement | Élevé, pas d’interface d’alerte native | 
| Wazuh seul | Gratuit, open source | HIDS et gestion de logs uniquement | Moyen, pas de capture réseau native | 
| Graylog | Gratuit (édition open) à payant | Centralisation de logs, sans capture réseau intégrée | Moyen | 
| Splunk / Sentinel | Facturé au volume ingéré, souvent 5 à 6 chiffres/an | SIEM complet avec support commercial | Faible mais coûteux | 

