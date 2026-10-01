---
id: collect-261001-huawei/huawei/security-onion-surveillance-reseau-nsm-en-13-etapes-1
title: "Calculer le hash SHA256 de l'ISO téléchargée"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution", "incident"]
source: docs/RAG/collect-261001-huawei/security-onion-surveillance-reseau-nsm-en-13-etapes.md
source_anchor: ""
source_lines: [1, 60]
sha256: 6278c42cc47c458bab4a71547eea72a962ff0eb44e74701d006d3b91c883ceb0
---

# Calculer le hash SHA256 de l'ISO téléchargée

Security Onion 3.3.0 est sorti le 11 septembre 2026, dix jours seulement avant la rédaction de ce tutoriel. Cette distribution Linux gratuite empile Suricata, Zeek, Wazuh et la suite Elastic dans un seul système capable de surveiller un réseau d’entreprise ou un lab de formation, sans payer une licence SIEM à cinq chiffres. Ce guide détaille l’installation complète, du téléchargement de l’ISO jusqu’à la première chasse aux menaces dans la console SOC, avec les vrais prérequis matériels, les pièges qui bloquent 90 % des débutants et une checklist de dépannage.

## Qu’est-ce que Security Onion et pourquoi l’installer en 2026

Security Onion est une distribution Linux libre, construite sur Rocky Linux 9, entièrement dédiée à la surveillance de sécurité réseau (Network Security Monitoring, NSM) et à la détection d’intrusion. Contrairement à un pare-feu ou un antivirus classique, l’outil ne bloque rien par défaut : il capture le trafic réseau, l’analyse avec plusieurs moteurs de détection en parallèle, puis centralise les alertes dans une console web unique appelée Security Onion Console (SOC).

La force du projet tient à l’intégration. Plutôt que d’installer et de faire dialoguer séparément Suricata pour la détection d’intrusion réseau, Zeek pour l’analyse de protocoles, Wazuh pour la supervision des postes et des logs, et la suite Elastic (Elasticsearch, Logstash, Kibana) pour l’indexation et la recherche, Security Onion Solutions livre le tout préconfiguré, avec des scripts d’installation qui gèrent les dépendances, les certificats et les flux de données entre composants. Le code source est public sur GitHub, gratuit à l’usage, et sans limite de nombre de capteurs contrairement aux offres SIEM commerciales facturées au volume de données ingérées.

Pour un service informatique en France, l’intérêt dépasse le simple aspect budgétaire. La directive NIS2 impose désormais aux entités essentielles et importantes une capacité de détection et de remontée d’incident dans des délais serrés. Disposer d’une plateforme NSM auto-hébergée, dont les données ne transitent jamais chez un tiers américain, répond directement à cette exigence de souveraineté tout en gardant un budget compatible avec une PME ou une collectivité.

La différence entre un SIEM classique et une plateforme NSM comme Security Onion mérite d’être clarifiée avant de se lancer. Un SIEM traditionnel se contente le plus souvent de collecter et corréler des journaux déjà produits par vos équipements (pare-feu, serveurs, applications). Security Onion va plus loin : il capture directement le trafic réseau brut, ce qui lui permet de détecter des activités que les journaux applicatifs ne mentionnent jamais, comme une exfiltration de données via un protocole détourné ou une communication de commande et contrôle chiffrée mais reconnaissable à ses métadonnées. Cette capacité de capture complète (full packet capture) transforme chaque alerte en dossier d’investigation exploitable, puisque l’analyste peut rejouer littéralement le paquet réseau à l’origine de la détection plutôt que de se contenter d’une ligne de log résumée.

## Prérequis : matériel, logiciels et versions exactes

Avant de lancer quoi que ce soit, vérifiez que votre machine (physique ou machine virtuelle) répond aux exigences officielles. La documentation Security Onion est stricte sur ce point : sous-dimensionner la RAM est la cause numéro un des installations qui échouent au démarrage des services Elasticsearch.

| Composant | Minimum absolu (standalone/éval) | Recommandé (trafic réel) | 
|---|---|---|
| Version testée | Security Onion 3.3.0 (build 3.3.0-20260911) | Idem, dernière release stable | 
| CPU | 4 cœurs | 8 cœurs ou plus | 
| RAM | 24 Go (avec swap conseillé) | 32 Go ou plus | 
| Stockage | 200 Go | 500 Go+ selon la rétention visée | 
| Interfaces réseau | 2 NIC (management + monitoring) | 2 NIC, la carte de capture en mode promiscuous | 
| Système de base | Rocky Linux 9 (ISO fournie ou installation minimale) | Rocky Linux 9 à jour | 

Ces chiffres proviennent directement de la documentation officielle (docs.securityonion.net) pour un déploiement standalone, c’est-à-dire un nœud unique qui cumule le rôle de manager et de capteur. Avec seulement 24 Go de RAM, l’installateur avertit qu’il faudra probablement activer de l’espace swap pour éviter que le service Elasticsearch ne soit tué par l’OOM killer du noyau Linux. Pour un lab de test avec un trafic faible, 24 Go suffisent. Pour surveiller ne serait-ce qu’un petit segment réseau de production, comptez 32 Go dès le départ.

Côté logiciels, vous aurez besoin de :

- L’image ISO officielle Security Onion 3.3.0, téléchargée exclusivement depuis le dépôt GitHub Security-Onion-Solutions/securityonion
- Un hyperviseur si vous testez en VM : VMware Workstation/ESXi, VirtualBox ou Proxmox VE
- Une clé USB de 8 Go minimum si vous installez sur du matériel physique
- Un accès Internet stable pendant l’installation (les paquets sont téléchargés depuis les dépôts officiels)
- Un navigateur récent (Chrome, Firefox ou Edge) pour accéder à la console SOC

Le fichier ISO doit systématiquement être vérifié par son empreinte SHA256 avant l’installation. Le dépôt GitHub inclut un fichier DOWNLOAD_AND_VERIFY_ISO.md qui documente la procédure officielle de vérification, une étape que beaucoup de tutoriels non officiels omettent et qui expose à un risque d’image corrompue ou modifiée.

## Étape 1 : télécharger et vérifier l’image ISO

Rendez-vous sur le dépôt GitHub officiel du projet. Ne téléchargez jamais l’ISO depuis un miroir tiers ou un lien partagé sur un forum : c’est un vecteur classique de compromission pour un outil dont le rôle est précisément de détecter des compromissions. Une fois le fichier récupéré, vérifiez son intégrité en ligne de commande.

```
# Calculer le hash SHA256 de l'ISO téléchargée
sha256sum securityonion-3.3.0-20260911.iso
# Comparer avec le hash publié sur la page de release GitHub
# Les deux valeurs doivent être strictement identiques
cat securityonion-3.3.0-20260911.iso.sha256
```
Si les deux empreintes ne correspondent pas, supprimez le fichier et retéléchargez-le. N’installez jamais une ISO dont le hash ne correspond pas : c’est une règle non négociable pour un outil de sécurité qui va ensuite avoir un accès privilégié à tout votre trafic réseau.

## Étape 2 : préparer la machine virtuelle ou le serveur physique

Pour un premier test, une machine virtuelle est largement suffisante et vous permet de faire des snapshots avant chaque étape critique. Dans votre hyperviseur, créez une VM avec les caractéristiques suivantes : 4 vCPU, 24 Go de RAM minimum, un disque de 200 Go en provisionnement fin, et deux cartes réseau virtuelles distinctes.

La première interface réseau (management) doit être connectée à votre réseau administratif classique, celui depuis lequel vous accéderez à la console SOC via HTTPS. La seconde interface (monitoring) doit être branchée sur un port SPAN, un port mirror de votre switch, ou une TAP réseau, de façon à recevoir une copie du trafic que vous souhaitez surveiller. Cette interface de capture ne doit jamais avoir d’adresse IP configurée dessus : elle fonctionne en mode promiscuous, invisible sur le réseau qu’elle observe.

Sur VMware ou Proxmox, pensez à activer le mode promiscuous au niveau du commutateur virtuel pour l’interface de monitoring, sinon Security Onion ne recevra jamais le trafic mirroré, même si la configuration du port SPAN côté switch physique est correcte.

## Étape 3 : démarrer l’installateur et choisir le mode d’installation

