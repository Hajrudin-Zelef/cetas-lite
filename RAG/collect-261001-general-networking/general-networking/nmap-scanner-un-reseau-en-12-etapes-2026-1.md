---
id: collect-261001-general-networking/general-networking/nmap-scanner-un-reseau-en-12-etapes-2026-1
title: "Nmap version 7.991 ( https://nmap.org )"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["cyber", "open source"]
source: docs/RAG/collect-261001-general-networking/nmap-scanner-un-reseau-en-12-etapes-2026.md
source_anchor: ""
source_lines: [1, 51]
sha256: 54c1a13111c9630eb6eb21e04ccd8423f4a452ddfd09d3b036017218deaa2bd7
---

# Nmap version 7.991 ( https://nmap.org )

Un scan Nmap mal ciblé peut faire tomber un onduleur industriel, déclencher une alerte au SOC d’un client ou, pire, vous exposer aux trois ans de prison prévus par l’article 323-1 du Code pénal. Et pourtant, Nmap reste en 2026 l’outil que la quasi-totalité des analystes réseau ouvrent en premier pour cartographier une infrastructure. La version 7.991, sortie début août 2026, vient encore renforcer sa base cryptographique. Ce tutoriel vous emmène de l’installation jusqu’à un projet complet d’audit réseau automatisé, avec le cadre légal français clarifié dès le départ.

**Avertissement d’usage :** ce guide s’adresse aux administrateurs systèmes, aux équipes sécurité et aux étudiants qui scannent leur propre infrastructure ou un périmètre couvert par une autorisation écrite. Scanner un système tiers sans accord expose à des poursuites pénales, détaillées plus bas. Tech Insider ne cautionne aucun usage non autorisé de **Nmap**.

## Qu’est-ce que Nmap et pourquoi il reste incontournable en 2026

**Nmap** (Network Mapper) est un scanner réseau open source créé en 1997 par Gordon Lyon, alias Fyodor. Sa mission n’a pas changé depuis : découvrir les hôtes actifs sur un réseau, lister leurs ports ouverts, identifier les services qui tournent dessus et, si besoin, deviner le système d’exploitation. Ce qui a changé, en revanche, c’est l’échelle et la profondeur de l’outil. La dernière version stable, **Nmap 7.991**, publiée début août 2026, embarque OpenSSL 3.0.21 et corrige plusieurs failles de la bibliothèque libssh2 qu’elle intègre (CVE-2025-15661, CVE-2026-7598, CVE-2026-55199, CVE-2026-55200, CVE-2026-58050 et CVE-2026-58051), toutes patchées dans une branche interne suivie sous le nom 1.11.1_NMAP1.

Plusieurs guides spécialisés publiés entre 2025 et 2026 classent encore **Nmap** parmi les cinq à dix outils les plus utilisés par les professionnels de la cybersécurité, aux côtés de Wireshark, Metasploit et Burp Suite. Le paquet Python `python3-nmap`, qui permet de piloter l’outil depuis des scripts d’automatisation, cumulait environ 190 000 téléchargements sur le mois de juin 2026 selon PyPI Stats, un signal net de l’usage massif de Nmap dans des pipelines automatisés plutôt qu’en ligne de commande isolée. Sur Windows, un seul miroir de téléchargement (Uptodown) recensait plus de 150 000 installations fin août 2025.

Concrètement, un administrateur réseau utilise Nmap pour trois choses : inventorier les machines connectées à un sous-réseau qu’il ne maîtrise pas totalement, vérifier qu’aucun port sensible n’est resté ouvert après un déploiement, et détecter des services obsolètes ou mal configurés avant qu’un attaquant ne le fasse à sa place. C’est un outil défensif autant qu’offensif : les deux usages emploient exactement la même commande.

Ce qui distingue Nmap de la plupart des scanners commerciaux, c’est sa communauté. Le projet publie son code source ouvertement, documente chaque option dans un manuel consultable en ligne et laisse des centaines de contributeurs indépendants écrire des scripts NSE pour des cas d’usage très spécifiques : détection de ransomware sur des partages SMB, audit de configuration TLS, énumération de bases de données exposées. Cette ouverture explique pourquoi Nmap reste enseigné dans la quasi-totalité des cursus de cybersécurité en France, du BTS SIO aux formations RSSI, et pourquoi il figure systématiquement dans la boîte à outils livrée avec des distributions comme Kali Linux ou Parrot OS.

## Nmap est-il légal en France ? Ce que dit précisément le Code pénal

C’est la question que devrait se poser quiconque tape `nmap` suivi d’une adresse IP qui n’est pas la sienne. En France, l’infraction de référence est l’**article 323-1 du Code pénal**, qui réprime « le fait d’accéder ou de se maintenir, frauduleusement, dans tout ou partie d’un système de traitement automatisé de données » (STAD). La peine de base est de **trois ans d’emprisonnement et 100 000 € d’amende**. Si le scan entraîne une suppression ou une altération de données, la peine grimpe à cinq ans et 150 000 €. Et si le système visé est un STAD contenant des données à caractère personnel géré par l’État, la sanction peut atteindre sept ans d’emprisonnement et 300 000 € d’amende, selon les chiffres actualisés cités par le cabinet ACI et le rapport 2026 de l’ICLG sur le droit de la cybersécurité en France.

La nuance juridique porte sur le scan de ports lui-même. Des analyses de l’OSSIR datant d’un colloque JSSI, encore citées dans la doctrine actuelle, expliquent qu’un scan SYN classique n’établit jamais de connexion TCP complète : le paquet SYN part, la cible répond par SYN/ACK, mais Nmap coupe la poignée de main avant le dernier ACK. Techniquement, il n’y a donc pas toujours d’« accès » au sens strict de l’article 323-1. Cela dit, plusieurs analyses juridiques plus récentes, dont un rapport 2026 sur la législation cyber française, rappellent que scanner un système tiers sans autorisation peut être retenu comme un **acte préparatoire** à une intrusion, ce qui suffit à justifier des poursuites, surtout si le scan est suivi d’une tentative d’exploitation.

En pratique, la règle à retenir est simple : ne scannez jamais un périmètre que vous ne possédez pas ou pour lequel vous n’avez pas une **autorisation écrite** (contrat, lettre de mission, ordre de mission pentest). Pour vos propres tests, restez sur votre réseau domestique, un labo virtualisé (Metasploitable, DVWA) ou un environnement cloud dédié aux tests d’intrusion autorisés. Le site de Cybermalveillance.gouv.fr publie des rappels réguliers sur ce cadre légal pour les professionnels comme pour les particuliers.

## Prérequis : versions, matériel et environnement de test

Avant de lancer la première commande, vérifiez que votre environnement correspond à ces prérequis. Ils sont volontairement modestes : Nmap tourne aussi bien sur un Raspberry Pi que sur un poste de travail professionnel.

| Élément | Version / configuration recommandée | Remarque | 
|---|---|---|
| Nmap | 7.991 (août 2026) ou 7.99 minimum | Versions antérieures à 7.98 exposées à CVE-2025-43715 sur l’installeur Windows | 
| Système d’exploitation | Linux (Ubuntu 24.04+, Debian 12+), Windows 10/11, macOS 14+ | Comportement légèrement différent selon l’OS pour les scans SYN | 
| Droits d’exécution | root / sudo (Linux, macOS), administrateur (Windows) | Obligatoire pour les scans SYN (-sS) et la détection d’OS (-O) | 
| Zenmap (interface graphique) | zenmap-7.991-py3 | Nécessite Python 3 ; facultatif si vous restez en CLI | 
| Cible de test | Machine virtuelle Metasploitable2/3 ou votre propre réseau local | Ne jamais cibler un tiers sans autorisation écrite | 
| Espace disque | 200 Mo minimum | Plus si vous stockez des exports XML volumineux | 

Comptez environ 30 minutes pour l’installation et les premiers scans, et jusqu’à 90 minutes si vous allez au bout du projet complet d’audit automatisé présenté à l’étape 12.

## Étape 1 : installer Nmap sur Linux, Windows et macOS

L’installation varie selon la plateforme, mais reste triviale partout. Sur les distributions Linux basées sur Debian ou Ubuntu :

```
sudo apt update
sudo apt install nmap -y
nmap --version
```
Sur Fedora, RHEL ou CentOS :

`sudo dnf install nmap -y`
Sur macOS, la méthode la plus fiable passe par Homebrew :

