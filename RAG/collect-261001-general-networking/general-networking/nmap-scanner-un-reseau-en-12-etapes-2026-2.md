---
id: collect-261001-general-networking/general-networking/nmap-scanner-un-reseau-en-12-etapes-2026-2
title: "Nmap version 7.991 ( https://nmap.org )"
domain: general-networking
role: reference
task: reference
actors: []
dates: ["2026-08-25"]
keywords: ["exploit", "latency"]
source: docs/RAG/collect-261001-general-networking/nmap-scanner-un-reseau-en-12-etapes-2026.md
source_anchor: ""
source_lines: [52, 158]
sha256: 84db92ea533cf6d3fab21888418219accb13488ef7f910f134a28447cec9a9ec
---

# Nmap version 7.991 ( https://nmap.org )

`brew install nmap`
Sur Windows, téléchargez l’installeur officiel depuis nmap.org (version 7.991), qui embarque désormais Npcap avec le pilote de capture réseau nécessaire aux scans bruts. L’installeur a été reconstruit avec NSIS 3.11 depuis la version 7.98, ce qui corrige la faille d’escalade de privilèges CVE-2025-43715 présente dans les anciens installeurs NSIS exécutés en tant que SYSTEM. Ne réutilisez donc jamais un installeur Nmap téléchargé avant août 2025.

Pour ceux qui préfèrent l’interface graphique, Zenmap reste disponible en package Python 3 (zenmap-7.991-py3) et permet de visualiser une topologie réseau sans mémoriser la syntaxe complète.

## Étape 2 : vérifier l’installation et lancer le premier scan

Une fois l’installation terminée, vérifiez la version exacte et lancez un scan minimal sur votre propre machine (localhost) pour confirmer que tout fonctionne :

```
nmap --version
# Nmap version 7.991 ( https://nmap.org )
# Platform: x86_64-pc-linux-gnu
# Compiled with: liblua-5.4.7 openssl-3.0.21 libssh2-1.11.1_NMAP1 ...
nmap 127.0.0.1
```
La sortie typique ressemble à ceci sur un poste Linux fraîchement installé :

```
Starting Nmap 7.991 ( https://nmap.org ) at 2026-08-25 14:02 CEST
Nmap scan report for localhost (127.0.0.1)
Host is up (0.000035s latency).
Not shown: 998 closed tcp ports (reset)
PORT     STATE SERVICE
22/tcp   open  ssh
631/tcp  open  ipp
Nmap done: 1 IP address (1 host up) scanned in 0.19 seconds
```
Si vous obtenez ce type de résultat, votre installation est fonctionnelle et vous pouvez passer aux scans réseau réels, toujours sur un périmètre autorisé.

## Étape 3 : découvrir les hôtes actifs sur le réseau

Avant de scanner des ports, il faut savoir quelles machines répondent. C’est le rôle du **ping scan**, déclenché avec l’option `-sn` (anciennement `-sP`), qui désactive le scan de ports et se contente de la découverte d’hôtes :

`nmap -sn 192.168.1.0/24`
Cette commande balaie les 254 adresses possibles du sous-réseau et renvoie la liste des hôtes qui répondent à un ping ICMP, une requête ARP (sur réseau local) ou une sonde TCP sur les ports 80 et 443. Sur un réseau domestique typique, le résultat prend quelques secondes et liste la box internet, les objets connectés, les ordinateurs et les smartphones actifs. C’est souvent la première surprise d’un audit : la plupart des foyers découvrent deux à trois fois plus d’appareils connectés qu’ils ne l’imaginaient.

## Étape 4 : scanner les ports TCP, SYN contre Connect

Une fois les hôtes actifs identifiés, l’étape suivante consiste à scanner leurs ports ouverts. Nmap propose deux techniques principales pour le TCP :

- **Scan SYN (-sS)** : envoie un paquet SYN et attend la réponse sans terminer la connexion. Rapide, discret, mais nécessite les droits root/administrateur pour manipuler les paquets bruts.
- **Scan Connect (-sT)** : utilise l’appel système`connect()` standard, termine la connexion TCP normalement. Plus lent et plus visible dans les journaux, mais ne demande aucun privilège particulier.

```
# Scan SYN, nécessite sudo
sudo nmap -sS 192.168.1.10
# Scan Connect, sans privilège root
nmap -sT 192.168.1.10
```
Pour cibler l’ensemble des 65 535 ports plutôt que les 1 000 ports les plus courants scannés par défaut, ajoutez `-p-` :

`sudo nmap -sS -p- 192.168.1.10`
Ce scan complet prend nettement plus de temps (plusieurs minutes contre quelques secondes) mais révèle des services qui tournent sur des ports non standards, une technique de dissimulation fréquente chez les attaquants qui déplacent un service SSH du port 22 vers un port à quatre chiffres.

## Étape 5 : scanner les ports UDP et composer avec les pare-feu

Le protocole UDP, utilisé par le DNS, le SNMP ou encore les services VoIP, se scanne différemment avec l’option `-sU`. Contrairement au TCP, l’absence de réponse ne signifie pas forcément qu’un port est ouvert : UDP n’accuse pas systématiquement réception.

`sudo nmap -sU --top-ports 100 192.168.1.10`
Un scan UDP complet sur 65 535 ports peut prendre plusieurs heures selon la configuration du pare-feu cible, c’est pourquoi on se limite en général aux ports les plus courants (`--top-ports`) en contexte d’audit régulier. Si la cible se trouve derrière un pare-feu qui bloque les échos ICMP, Nmap peut la considérer comme éteinte alors qu’elle répond bien. L’option `-Pn` force le scan sans phase de découverte préalable :

`sudo nmap -Pn -sS 192.168.1.10`
## Étape 6 : détecter les services et leurs versions

Savoir qu’un port est ouvert ne dit rien sur ce qui tourne dessus. L’option `-sV` interroge chaque port ouvert pour identifier le service exact et sa version :

`sudo nmap -sV 192.168.1.10````
PORT    STATE SERVICE VERSION
22/tcp  open  ssh     OpenSSH 9.6p1 Ubuntu 3ubuntu13.5
80/tcp  open  http    nginx 1.26.1
3306/tcp open  mysql   MySQL 8.0.39
```
Cette information est précieuse pour croiser les versions détectées avec des bases de vulnérabilités connues : un OpenSSH ou un MySQL non patché depuis des mois est le premier signal qu’un auditeur va relever. Combinez `-sV` avec un niveau d’intensité plus élevé (`--version-intensity 9`) si les sondes par défaut ne suffisent pas à identifier un service exotique.

## Étape 7 : identifier le système d’exploitation

L’option `-O` déclenche la détection d’OS, basée sur l’analyse des caractéristiques de la pile TCP/IP de la cible (taille de fenêtre, options TCP, TTL initial). Elle nécessite les droits root :

`sudo nmap -O 192.168.1.10`
La commande combinée la plus utilisée par les praticiens en 2026, selon plusieurs cheat sheets publiés cette année, reste `-A`, qui active en une seule fois la détection d’OS, la détection de version, les scripts NSE par défaut et un traceroute :

`sudo nmap -A 192.168.1.10`
Gardez à l’esprit que la détection d’OS reste probabiliste : Nmap indique un pourcentage de certitude et peut se tromper sur des systèmes durcis, des conteneurs ou des machines derrière un pare-feu qui normalise les paquets sortants.

## Étape 8 : ajuster la vitesse avec les modèles de timing

Nmap propose six profils de vitesse, de `-T0` (paranoïaque, un paquet toutes les cinq minutes) à `-T5` (insane, aussi rapide que possible au risque de fausses négatives). Le profil par défaut est `-T3`.

| Template | Nom | Cas d’usage typique | 
|---|---|---|
| -T0 | Paranoid | Contournement d’IDS très sensible, scan sur plusieurs jours | 
| -T1 | Sneaky | Scan discret sur réseau surveillé | 
| -T2 | Polite | Limite la charge réseau sur des liens fragiles | 
| -T3 | Normal | Profil par défaut, équilibre vitesse/fiabilité | 
| -T4 | Aggressive | Réseau local rapide, audit interne planifié | 
| -T5 | Insane | Labo de test uniquement, risque de résultats incomplets | 

En environnement professionnel, `-T4` est le choix le plus courant pour un audit interne planifié en dehors des heures de production, tandis que `-T2` convient mieux à un scan mené en journée sur une infrastructure sensible où la latence compte.

## Étape 9 : automatiser l’audit avec le moteur de scripts NSE

Le **Nmap Scripting Engine (NSE)** transforme Nmap en véritable plateforme d’audit. Les scripts sont regroupés en catégories : `default` (sûrs, activés par `-sC`), `vuln` (vérification active de vulnérabilités connues), `auth`, `brute`, `intrusive` et `exploit`. Pour lancer une recherche de vulnérabilités documentées sur une cible :

`sudo nmap -sV --script vuln 192.168.1.10`
Le script `vulners` pousse la logique plus loin en interrogeant la base Vulners.com pour chaque service identifié, et renvoie directement les identifiants CVE et les scores CVSS correspondants :

