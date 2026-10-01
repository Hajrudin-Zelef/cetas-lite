---
id: collect-261001-cisco/cisco/graylog-7-1-centraliser-ses-logs-en-11-etapes-2026-2
title: "Génère un secret de session Graylog (obligatoire, 16+ caractères)"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "agents", "arr", "memory", "open source"]
source: docs/RAG/collect-261001-cisco/graylog-7-1-centraliser-ses-logs-en-11-etapes-2026.md
source_anchor: ""
source_lines: [33, 117]
sha256: fcbcd608da3b7bb69ea05935fd822e2d3a8bf640dfe85e3fae01a64f4cd0bd40
---

# Génère un secret de session Graylog (obligatoire, 16+ caractères)

| Critère | Graylog 7.1 | Wazuh 4.14.7 | ELK / Elastic 9.5 | 
|---|---|---|---|
| Approche | Centralisation de logs + SIEM | Agents HIDS + SIEM | Recherche et analytics + module sécurité | 
| Composants requis | MongoDB + OpenSearch | Indexer + Manager + agents | Elasticsearch + Kibana + Beats | 
| Édition gratuite | Graylog Open | Complet, open source | Basic (fonctions limitées) | 
| Courbe d’apprentissage | Modérée | Modérée à élevée | Élevée | 
| Cas d’usage idéal | Centralisation multi-source, pipelines | Surveillance de postes et serveurs | Recherche full-text à grande échelle | 

Dans une architecture de sécurité mature, ces outils ne s’excluent pas. Beaucoup d’équipes font remonter les alertes de CrowdSec ou les résultats d’un scan OpenVAS/Greenbone vers Graylog pour disposer d’une vue unique. C’est cette approche que ce tutoriel construit pas à pas.

Le choix dépend surtout de la maturité de l’équipe qui exploitera l’outil au quotidien. Une petite structure qui débute dans la gestion de logs gagnera du temps avec Graylog, dont le pipeline de traitement et les dashboards se configurent depuis l’interface web sans écrire de requête complexe. Une équipe qui gère déjà un stack ELK pour d’autres usages (recherche produit, analytics métier) aura intérêt à mutualiser avec Elastic Security plutôt qu’ajouter une quatrième brique. Wazuh reste imbattable dès que la détection doit descendre au niveau du poste de travail, avec surveillance de l’intégrité des fichiers et détection de rootkits, une couche que ni Graylog ni Elastic ne couvrent nativement sans agent dédié.

## Prérequis : matériel, logiciels et versions nécessaires

Cette installation a été testée sur Ubuntu Server 24.04 LTS, mais fonctionne à l’identique sur Debian 12. Voici la configuration minimale recommandée pour un usage réel, pas seulement une démo.

- Système d’exploitation : Ubuntu Server 24.04 LTS ou Debian 12 (64 bits)
- RAM : 8 Go minimum pour un lab, 16 Go recommandés en production (OpenSearch consomme la majorité)
- CPU : 4 vCPU minimum
- Stockage : 50 Go de SSD pour un test, plusieurs centaines de Go en production selon le volume de logs ingérés
- Docker Engine version 27.x ou supérieure
- Docker Compose v2 (intégré au plugin docker-compose-plugin)
- MongoDB 8.0.x (compatible avec la plage 7.x à 8.2.x exigée par Graylog 7.1)
- OpenSearch 2.19.x (dernière version compatible avec Graylog 7.1.x)
- Graylog 7.1.8, image Docker officielle
- Accès root ou sudo sur le serveur
- Ports ouverts : 9000 (interface web), 12201 (GELF UDP/TCP), 1514 (Syslog), 27017 (MongoDB, interne uniquement)

Un point de vigilance mémoire : OpenSearch refuse de démarrer si la limite `vm.max_map_count` du noyau Linux est trop basse. C’est l’une des causes d’échec les plus fréquentes, traitée dans la section dépannage plus bas. La matrice de compatibilité officielle de Graylog reste la référence à consulter avant tout choix de version, car elle change à chaque montée de branche majeure.

## Dimensionner son serveur selon le volume de logs attendu

Le dimensionnement matériel dépend presque entièrement du volume quotidien de messages ingérés, pas du nombre de serveurs surveillés. Un parc de dix machines Linux qui produit surtout des logs d’authentification et des journaux applicatifs légers reste sous la barre du gigaoctet par jour. À l’inverse, un pare-feu d’entreprise qui journalise chaque connexion réseau, ou une flotte de conteneurs Kubernetes en mode debug, peut facilement dépasser 20 Go par jour à lui seul.

| Volume quotidien de logs | RAM recommandée | vCPU | Stockage SSD (rétention 30 jours) | 
|---|---|---|---|
| Moins de 1 Go/jour (lab, PME < 20 postes) | 8 Go | 4 | 50 Go | 
| 1 à 5 Go/jour (PME, quelques serveurs critiques) | 16 Go | 4 à 6 | 200 Go | 
| 5 à 20 Go/jour (ETI, infrastructure hybride) | 32 Go | 8 | 750 Go, cluster OpenSearch recommandé | 
| Plus de 20 Go/jour | 64 Go et plus, cluster multi-nœuds | 16+ | Plusieurs To, data tiering obligatoire | 

Ces chiffres restent des ordres de grandeur, pas des garanties absolues. Un pipeline mal écrit qui applique des expressions régulières coûteuses sur chaque message peut faire chuter le débit d’ingestion de moitié, même sur un serveur correctement dimensionné selon ce tableau. Il vaut toujours mieux commencer avec une marge de 30 % au-dessus de l’estimation initiale et surveiller la charge CPU d’OpenSearch pendant les deux premières semaines d’exploitation réelle.

## Étape 1 : Préparer le serveur et installer Docker

Sur un serveur fraîchement provisionné, on commence par mettre à jour le système puis installer Docker Engine depuis le dépôt officiel, pas depuis le paquet `docker.io` d’Ubuntu qui traîne souvent une version obsolète.

```
sudo apt update && sudo apt upgrade -y
sudo apt install -y ca-certificates curl gnupg
sudo install -m 0755 -d /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg
sudo chmod a+r /etc/apt/keyrings/docker.gpg
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
  $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt update
sudo apt install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
docker --version
docker compose version
```
Il faut ensuite ajuster deux réglages noyau indispensables à OpenSearch : la limite de zones mémoire mappées et le nombre de fichiers ouverts. Sans ce réglage, le conteneur OpenSearch s’arrête immédiatement au démarrage avec une erreur “max virtual memory areas vm.max_map_count is too low”.

```
echo "vm.max_map_count=262144" | sudo tee -a /etc/sysctl.conf
sudo sysctl -w vm.max_map_count=262144
echo "* soft nofile 65536" | sudo tee -a /etc/security/limits.conf
echo "* hard nofile 65536" | sudo tee -a /etc/security/limits.conf
```
## Étape 2 : Rédiger le fichier Docker Compose pour MongoDB, OpenSearch et Graylog

L’architecture de Graylog repose sur trois briques distinctes qui doivent communiquer entre elles : MongoDB stocke la configuration et les métadonnées (jamais les logs eux-mêmes), OpenSearch indexe et stocke le contenu des messages, et le serveur Graylog orchestre le tout via son interface web et son moteur de traitement. On les déploie ensemble avec un seul fichier `docker-compose.yml`.

Créez d’abord un répertoire de travail et un fichier `.env` pour ne jamais coder les secrets en dur.

```
mkdir -p ~/graylog-stack && cd ~/graylog-stack
# Génère un secret de session Graylog (obligatoire, 16+ caractères)
echo "GRAYLOG_PASSWORD_SECRET=$(openssl rand -hex 32)" > .env
# Génère le hash SHA-256 du mot de passe admin (remplacez MonMotDePasse!2026)
echo "GRAYLOG_ROOT_PASSWORD_SHA2=$(echo -n 'MonMotDePasse!2026' | sha256sum | cut -d' ' -f1)" >> .env
```
Le fichier `docker-compose.yml` suivant déploie la pile complète avec des volumes persistants, essentiels pour ne pas perdre les logs au moindre redémarrage de conteneur.

