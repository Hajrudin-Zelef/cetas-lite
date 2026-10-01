---
id: collect-261001-general-networking/general-networking/nftables-pare-feu-linux-en-13-etapes-2026-1
title: "Debian / Ubuntu"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution", "open source"]
source: docs/RAG/collect-261001-general-networking/nftables-pare-feu-linux-en-13-etapes-2026.md
source_anchor: ""
source_lines: [1, 76]
sha256: bdb299419d3a819512e8c7bdef9c9432472c947777176347b8c9692d66922d94
---

# Debian / Ubuntu

Un serveur Linux exposé sur Internet reçoit des tentatives de connexion non sollicitées en quelques minutes à peine après sa mise en ligne. Sans filtrage réseau correctement configuré, chaque port ouvert devient une porte d’entrée potentielle. C’est là qu’intervient nftables, le sous-système de filtrage de paquets qui a progressivement remplacé iptables dans le noyau Linux. Debian, Ubuntu, Fedora et Rocky Linux s’appuient désormais sur ce moteur en interne, même quand l’administrateur continue de taper des commandes iptables par habitude. Ce tutoriel montre comment configurer nftables depuis zéro, protéger un serveur SSH contre le brute-force, mettre en place du NAT, journaliser les paquets rejetés, puis migrer proprement depuis une ancienne configuration iptables. Treize étapes, un projet complet livré en fin d’article, et une liste de pièges qui font perdre des heures aux administrateurs pressés.

## Pourquoi nftables remplace iptables sur les serveurs Linux

iptables a longtemps été la référence pour filtrer le trafic réseau sous Linux. Le projet Netfilter, qui maintient ces outils, a introduit nftables pour corriger des limites structurelles devenues gênantes avec le temps : quatre outils séparés (iptables, ip6tables, arptables, ebtables), une syntaxe verbeuse, et des performances qui se dégradent avec des milliers de règles. nftables unifie tout cela dans un seul outil, `nft`, capable de gérer IPv4, IPv6, ARP et le pontage réseau avec une seule syntaxe cohérente.

Les distributions majeures ont basculé nftables comme moteur par défaut depuis plusieurs années déjà. Les commandes iptables classiques passent désormais par une couche de compatibilité, `iptables-nft`, qui traduit les anciennes règles vers le format nftables en coulisses. Autrement dit, même un administrateur qui n’a jamais tapé une commande `nft` utilise probablement déjà ce moteur sans le savoir. Le CERT-FR, l’organisme public français qui publie des bulletins de vulnérabilités quotidiens, rappelle régulièrement que le filtrage réseau reste une des premières lignes de défense contre l’exploitation de services exposés, notamment quand un correctif n’est pas encore disponible pour une faille annoncée.

L’ANSSI a également mis à jour sa politique open source en 2026, avec quatre axes qui incluent la publication de logiciels de cybersécurité sous licence libre et l’utilisation de solutions open source dans les infrastructures publiques. nftables, disponible dans le noyau Linux lui-même, coche cette case : pas de licence à acheter, pas de boîte noire, et un code audité par la communauté Netfilter depuis des années.

Au-delà de l’aspect réglementaire, l’argument technique reste le plus convaincant pour un administrateur qui gère plusieurs dizaines de serveurs. iptables devait dupliquer chaque règle IPv4 dans une commande ip6tables séparée pour couvrir IPv6, doublant mécaniquement le travail et le risque d’incohérence entre les deux jeux de règles. nftables résout ce problème avec la famille `inet`, qui applique une seule règle aux deux protocoles simultanément. Sur un parc de serveurs où IPv6 devient progressivement la norme chez les hébergeurs européens, cette unification change concrètement le temps passé à maintenir une configuration réseau cohérente.

## Prérequis : versions, distributions et compétences nécessaires

Avant de commencer, vérifiez que votre environnement correspond aux bases suivantes. nftables fonctionne sur toutes les distributions Linux modernes, mais les commandes d’installation et les chemins de fichiers varient légèrement d’une famille à l’autre.

| Élément | Version ou configuration recommandée | Remarque | 
|---|---|---|
| Distribution | Debian 12/13, Ubuntu 22.04 LTS et supérieur, Fedora récente, Rocky Linux 9 et supérieur | nftables est le moteur par défaut sur ces versions | 
| Noyau Linux | Série 5.x ou 6.x | Le support nftables complet nécessite un noyau relativement récent | 
| Paquet | nftables (dépôts officiels de chaque distribution) | Toujours installer depuis les dépôts, jamais depuis une source tierce non vérifiée | 
| Accès | Droits root ou sudo | Indispensable pour modifier les règles réseau | 
| Accès physique ou console | Recommandé pour un serveur distant | Une mauvaise règle peut couper l’accès SSH | 
| Connaissances | Bases de TCP/IP, ports, protocoles | Comprendre ce qu’on filtre évite les erreurs de configuration | 

Un point mérite d’être souligné avant de toucher à quoi que ce soit : sur un serveur distant accessible uniquement par SSH, une règle mal écrite peut couper la connexion et vous laisser dehors. Gardez toujours une session console ouverte (via l’hébergeur, une VM locale ou un accès KVM) pendant les premiers tests, et prévoyez un mécanisme de rollback automatique. Ce tutoriel montre comment faire les deux.

## Étape 1 : Vérifier le noyau et la compatibilité nftables

Commencez par confirmer que votre système prend en charge nftables nativement. La plupart des noyaux récents l’intègrent déjà, mais mieux vaut vérifier avant d’installer quoi que ce soit.

```
uname -r
grep -i nf_tables /boot/config-$(uname -r) 2>/dev/null
lsmod | grep nf_tables
```
Si la commande `lsmod` ne remonte rien, ce n’est pas forcément un problème : le module se charge automatiquement au premier usage de `nft`. Vérifiez aussi si iptables fonctionne déjà en mode legacy ou en mode nft sur votre machine, car les deux backends ne doivent jamais coexister sur les mêmes règles au risque de créer des conflits difficiles à diagnostiquer.

```
update-alternatives --display iptables 2>/dev/null
iptables --version
```
Une sortie mentionnant `nf_tables` à côté du numéro de version confirme que vous êtes déjà sur le bon backend. Sur Debian et Ubuntu récents, c’est le comportement par défaut depuis longtemps.

## Étape 2 : Installer nftables sur Debian, Ubuntu, Fedora et Rocky Linux

L’installation change peu d’une distribution à l’autre, mais quelques détails comptent, notamment la désactivation propre d’un ancien service iptables qui pourrait entrer en conflit.

```
# Debian / Ubuntu
sudo apt update
sudo apt install nftables -y
sudo systemctl enable --now nftables
# Fedora / RHEL / Rocky Linux
sudo dnf install nftables -y
sudo systemctl enable --now nftables
# Arch Linux
sudo pacman -S nftables
sudo systemctl enable --now nftables
```
Sur les systèmes qui utilisaient auparavant firewalld ou ufw comme frontend, désactivez-les avant de démarrer nftables directement, sous peine de voir deux gestionnaires de règles se marcher dessus. Un conflit entre firewalld et nftables actif en parallèle est l’une des causes les plus fréquentes de comportements imprévisibles signalées sur les forums d’administration système.

```
sudo systemctl disable --now firewalld
sudo systemctl disable --now ufw
```
Vérifiez ensuite que le service tourne et que le fichier de configuration principal existe déjà, généralement vide ou avec une politique par défaut minimale.

```
sudo systemctl status nftables
cat /etc/nftables.conf
```
## Étape 3 : Comprendre les tables, chaînes, règles et priorités

nftables organise le filtrage en trois niveaux hiérarchiques, et comprendre cette structure évite de nombreuses erreurs par la suite. Une table regroupe des chaînes par famille d’adresses (ip, ip6, inet pour les deux à la fois, arp, bridge). Une chaîne contient une liste ordonnée de règles et s’accroche à un point d’ancrage du noyau comme `input`, `output` ou `forward`. Une règle définit une condition de correspondance suivie d’une action, appelée verdict : accept, drop, reject, ou un renvoi vers une autre chaîne.

