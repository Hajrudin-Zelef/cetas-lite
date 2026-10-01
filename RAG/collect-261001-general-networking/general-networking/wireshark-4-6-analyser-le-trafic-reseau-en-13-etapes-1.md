---
id: collect-261001-general-networking/general-networking/wireshark-4-6-analyser-le-trafic-reseau-en-13-etapes-1
title: "Sur Fedora / RHEL / Rocky Linux"
domain: general-networking
role: reference
task: reference
actors: ["JFrog"]
dates: []
keywords: ["distribution", "ethernet", "exploit", "gpu", "incident", "open source"]
source: docs/RAG/collect-261001-general-networking/wireshark-4-6-analyser-le-trafic-reseau-en-13-etapes.md
source_anchor: ""
source_lines: [1, 39]
sha256: fccece096a4497a4a0d678c8c4df773b92e9247e69ef31da76ca8b4beff96044
---

# Sur Fedora / RHEL / Rocky Linux

Un pare-feu bloque ce qu’il reconnaît. Un antivirus détecte ce qu’il a déjà catalogué. Mais quand un serveur rame sans raison apparente, qu’une application refuse de se connecter ou qu’un poste envoie des paquets vers une adresse inconnue à 3h du matin, un seul outil permet de voir ce qui circule vraiment sur le câble : Wireshark. Sorti en version 4.6.8 stable cet été, l’analyseur de paquets open source reste, vingt-huit ans après ses débuts sous le nom d’Ethereal, la référence pour quiconque doit diagnostiquer un réseau ou enquêter sur un incident. Ce tutoriel détaille l’installation, la capture, le filtrage et l’analyse de trafic en 13 étapes concrètes, avec des exemples de commandes, un projet complet de détection de brute-force SSH, et les pièges qui font perdre des heures aux débutants.

## Pourquoi apprendre Wireshark en 2026

Le 1er septembre 2026, le CERT-FR (rattaché à l’ANSSI) publiait encore cinq avis de sécurité en une seule journée, portant sur SPIP, Kaspersky Endpoint Security Windows et JFrog Artifactory, entre autres. La veille, le portail cyberveille.esante.gouv.fr recensait des failles critiques activement exploitées sur PaperCut (CVSS 9,1) et Gitea (CVSS 9,8), ainsi que des vulnérabilités notées 9,8 sur Redis et WatchGuard. Ce rythme n’est pas une anomalie : c’est le quotidien de la cybersécurité en France depuis le lancement de la stratégie nationale 2026-2030, qui pousse les entreprises et les administrations à renforcer la surveillance de leurs infrastructures exposées, notamment via le portail de signalement 17Cyber.gouv.fr.

Un scanner de vulnérabilités comme OpenVAS/Greenbone ou un IDS comme Suricata repèrent des signatures connues. Wireshark, lui, montre le trafic brut, sans filtre ni interprétation automatique. C’est justement ce qui en fait un complément indispensable : quand une alerte tombe, c’est souvent une capture Wireshark qui permet de confirmer si l’incident est réel, d’identifier la machine compromise, ou de comprendre pourquoi une application tierce échoue silencieusement. Les administrateurs réseau, les étudiants en cybersécurité, les développeurs qui débuggent une API et les analystes en réponse à incident l’utilisent tous, pour des raisons différentes mais avec les mêmes commandes de base.

## Qu’est-ce que Wireshark et à quoi ça sert vraiment

Wireshark est un analyseur de protocoles réseau open source sous licence GPLv2, disponible sur Windows, macOS, Linux et BSD. Il capture le trafic qui transite par une interface réseau, puis le décode paquet par paquet jusqu’au dernier octet, pour des centaines de protocoles (Ethernet, IP, TCP, UDP, HTTP, DNS, TLS, SMB, et bien d’autres). Le projet, fondé par Gerald Combs en 1998 sous le nom d’Ethereal, a été renommé Wireshark en 2006 après un différend avec un ancien employeur sur la marque. Il tourne aujourd’hui sous la Wireshark Foundation et reste maintenu activement, avec une branche stable (4.6.8) et une ancienne branche de maintenance (4.4.18) toujours supportée pour les environnements qui ne peuvent pas migrer immédiatement.

La documentation officielle est claire sur ce que l’outil n’est pas : “Wireshark isn’t an intrusion detection system. It will not warn you when someone does strange things on your network that he/she isn’t allowed to do.” (Wireshark n’est pas un système de détection d’intrusion. Il ne vous avertira pas si quelqu’un fait des choses anormales sur votre réseau sans y être autorisé), précise le guide utilisateur officiel. Le rôle de Wireshark commence après l’alerte, pas avant. Un forum communautaire officiel résume la nuance en réponse à la question de savoir si l’outil peut faire office d’IDS : “It would come into play once IDS alarms have been raised and network logs have been preserved to investigate the occurrence in more detail” (il entre en jeu une fois que des alertes IDS ont été déclenchées et que les journaux réseau ont été conservés pour approfondir l’investigation), selon une réponse publiée sur le forum d’assistance Wireshark. En clair : Suricata ou Wazuh sonnent l’alarme, Wireshark sert à comprendre ce qui s’est réellement passé.

Les usages concrets couvrent quatre familles : le diagnostic réseau classique (latence, pertes de paquets, mauvaise configuration DNS ou DHCP), le débogage applicatif (une API qui timeout, une requête HTTP mal formée, un certificat TLS invalide), la formation et la pédagogie (comprendre comment fonctionne réellement un handshake TCP ou une résolution DNS), et l’investigation de sécurité (confirmer une exfiltration de données, analyser un fichier PCAP fourni par un client lors d’un test d’intrusion, ou documenter une preuve pour un rapport d’incident).

## Prérequis : versions, matériel et droits nécessaires

Avant de commencer, vérifiez que votre environnement correspond aux versions actuelles. Wireshark évolue vite côté sécurité : les anciennes versions 3.x ne reçoivent plus de correctifs et embarquent des bibliothèques de décodage obsolètes, potentiellement vulnérables.

| Composant | Version recommandée (sept. 2026) | Rôle | Remarque | 
|---|---|---|---|
| Wireshark | 4.6.8 (stable) ou 4.4.18 (maintenance) | Interface graphique et moteur de décodage | Évitez toute version antérieure à 4.4 | 
| Npcap | 1.88 | Pilote de capture pour Windows | Remplace WinPcap, abandonné depuis 2013 | 
| libpcap | Fournie par la distribution Linux/macOS | Capture bas niveau sur Unix | Installée via le paquet système, pas de téléchargement séparé | 
| TShark | Identique à Wireshark | Version ligne de commande | Installée avec le paquet principal | 
| Espace disque | 10 Go minimum recommandés | Stockage des captures | Un lien 1 Gbit/s saturé génère plusieurs Go/minute | 
| Droits système | Administrateur (Windows) ou groupe “wireshark” (Linux) | Accès à l’interface réseau en mode promiscuous | Ne jamais lancer Wireshark en root en permanence | 

Sur le plan matériel, un PC de bureau standard suffit largement : Wireshark n’a pas besoin de GPU ni de RAM excessive pour du trafic courant. La contrainte vient surtout du stockage et du CPU quand on analyse des captures de plusieurs gigaoctets avec des filtres complexes. Côté réseau, capturer sur un switch nécessite soit un accès physique à un port mirroring (SPAN), soit une capture directe sur la machine cible : sur un réseau commuté moderne, vous ne voyez par défaut que votre propre trafic, pas celui du voisin.

Dernier prérequis, non technique celui-là : l’autorisation. Capturer du trafic sur un réseau qui ne vous appartient pas, ou sans le consentement des personnes concernées, expose à des poursuites au titre de l’article 226-15 du Code pénal (atteinte au secret des correspondances) et du RGPD si des données personnelles sont interceptées. Ce tutoriel part du principe que vous travaillez sur votre propre réseau, un labo isolé, ou dans le cadre d’un mandat de test d’intrusion écrit.

## Étapes 1 à 3 : installer Wireshark 4.6 et configurer la capture

### Étape 1 : télécharger et installer Wireshark

Téléchargez toujours Wireshark depuis le site officiel, jamais depuis un site tiers ou un lien raccourci : un faux installeur de Wireshark est un vecteur classique de malware, précisément parce que l’outil inspire confiance. Sur Windows, l’installeur .exe propose d’office l’installation de Npcap. Sur macOS, privilégiez le paquet .dmg officiel ou `brew install wireshark` si vous utilisez Homebrew. Sur Debian, Ubuntu et dérivés, la commande est directe :

