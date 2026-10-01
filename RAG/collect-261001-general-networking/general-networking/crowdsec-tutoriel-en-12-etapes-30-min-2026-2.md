---
id: collect-261001-general-networking/general-networking/crowdsec-tutoriel-en-12-etapes-30-min-2026-2
title: "Mettre à jour la liste des paquets et le système"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["acquisition", "apache", "distribution", "mai"]
source: docs/RAG/collect-261001-general-networking/crowdsec-tutoriel-en-12-etapes-30-min-2026.md
source_anchor: ""
source_lines: [56, 160]
sha256: 91500cae30f1aa7365a09760ad2c2d9a261204f4925cd5ce955c7846840dc5a1
---

# Mettre à jour la liste des paquets et le système

| Composant | Recommandation | Notes | 
|---|---|---|
| CrowdSec Security Engine | 1.7.8 (mai 2026) | Dernière version stable, paquets officiels | 
| Système d’exploitation | Debian 12/13, Ubuntu 22.04/24.04 | Aussi RHEL, Fedora, Alpine, FreeBSD | 
| RAM | 256 Mo minimum | ~128 Mo consommés en pratique | 
| Processeur | 1 vCPU | Empreinte CPU négligeable au repos | 
| Pare-feu | iptables ou nftables | Requis par le firewall-bouncer | 
| Accès | root ou sudo | Indispensable pour l’installation | 
| Port sortant | 443/tcp | Vers la CAPI et la Console | 

Un détail rassurant : aucune base de données externe n’est nécessaire pour démarrer. CrowdSec utilise par défaut une base SQLite locale (`/var/lib/crowdsec/data/crowdsec.db`), parfaitement adaptée à un serveur unique. Pour une architecture multi-serveur en production, vous pourrez basculer plus tard sur PostgreSQL ou MySQL, mais ce n’est pas requis pour ce tutoriel. Assurez-vous simplement que la date et l’heure du serveur sont correctes (via `chrony` ou `systemd-timesyncd`), car des horloges décalées faussent la corrélation des événements.

## Étapes 1 à 3 – Installer le moteur CrowdSec et le bouncer pare-feu

Les trois premières étapes constituent le socle : ajouter le dépôt officiel, installer le moteur de détection, puis installer le bouncer qui appliquera réellement les bannissements. C’est le point que beaucoup ratent : sans bouncer, CrowdSec *voit* les attaques mais ne *bloque* rien.

### Étape 1 – Ajouter le dépôt officiel CrowdSec

Commencez par mettre votre système à jour, puis exécutez le script d’installation du dépôt fourni par l’éditeur. Ce script POSIX détecte automatiquement votre distribution et configure le dépôt packagecloud adéquat.

```
# Mettre à jour la liste des paquets et le système
sudo apt update && sudo apt upgrade -y
# Ajouter le dépôt officiel CrowdSec (détection automatique de la distro)
curl -s https://install.crowdsec.net | sudo sh
```
Si vous êtes sur une distribution à base de Red Hat (RHEL, Rocky, AlmaLinux), le même script bascule sur `yum`/`dnf`. Pour les puristes méfiants vis-à-vis du *curl | sh*, vous pouvez télécharger le script, l’inspecter, puis l’exécuter manuellement – c’est une bonne habitude de sécurité.

### Étape 2 – Installer le moteur de sécurité

L’installation du paquet `crowdsec` est l’étape clé. Lors de cette opération, l’assistant analyse les services déjà présents sur la machine (SSH, Nginx, Apache…) et installe automatiquement les collections de détection correspondantes. C’est ce qui rend l’installation de CrowdSec si rapide.

```
# Installer le moteur de détection CrowdSec
sudo apt install crowdsec -y
# Vérifier que le service tourne
sudo systemctl status crowdsec --no-pager
```
Sortie attendue : la ligne `Active: active (running)` en vert confirme que le moteur est opérationnel. Dès cet instant, CrowdSec lit déjà vos logs et se connecte à la CAPI pour télécharger la blocklist communautaire.

### Étape 3 – Installer le bouncer pare-feu (remédiation)

Le bouncer est le bras armé du dispositif. Choisissez la variante adaptée à votre pare-feu : `iptables` (classique) ou `nftables` (par défaut sur Debian 11+ et Ubuntu récents). En cas de doute, vérifiez avec `nft list ruleset` : si la commande renvoie un résultat, optez pour nftables.

```
# Variante iptables (la plus répandue)
sudo apt install crowdsec-firewall-bouncer-iptables -y
# OU variante nftables (Debian/Ubuntu récents)
# sudo apt install crowdsec-firewall-bouncer-nftables -y
# Confirmer l'enregistrement du bouncer auprès de la LAPI
sudo cscli bouncers list
```
Le bouncer s’enregistre tout seul auprès de la LAPI grâce à une clé API générée à l’installation. Dans la sortie de `cscli bouncers list`, vous devez voir une ligne du type `cs-firewall-bouncer-...` avec un statut valide. Si le bouncer apparaît, votre chaîne détection → remédiation est désormais complète.

## Étape 4 – Vérifier l’installation et lire les métriques

Avant d’aller plus loin, validez que tout communique correctement. L’outil en ligne de commande `cscli` est votre couteau suisse : il interroge la LAPI pour afficher l’état du moteur, les collections actives, les bouncers connectés et les métriques de traitement.

```
# Vue d'ensemble des métriques (lignes lues, alertes, décisions)
sudo cscli metrics
# Collections de détection installées
sudo cscli collections list
# Bouncers connectés à la LAPI
sudo cscli bouncers list
# Moteurs (machines) enregistrés
sudo cscli machines list
# État de la connexion à la blocklist communautaire (CAPI)
sudo cscli capi status
```
Dans la sortie de `cscli metrics`, concentrez-vous sur la section **Acquisition Metrics** : elle indique combien de lignes ont été lues par source de logs. Si une source affiche zéro ligne lue alors que le service génère bien des journaux, c’est le signe d’un problème de chemin dans votre fichier d’acquisition – nous y reviendrons dans la section dépannage. La commande `cscli capi status` doit, elle, renvoyer un message confirmant que vous pouvez interagir avec la Central API (CAPI).

## Étape 5 – Maîtriser les collections, parseurs et scénarios du Hub

Le Hub est ce qui distingue CrowdSec d’un simple bloqueur d’IP. Une **collection** regroupe les parseurs (qui transforment une ligne de log brute en événement structuré) et les scénarios (qui décrivent un comportement malveillant, par exemple « plus de 5 échecs SSH en 10 secondes »). Vous pilotez tout cela via `cscli`.

```
# Parcourir tout le contenu du Hub disponible
sudo cscli hub list
# Installer la collection SSH (si elle n'a pas été ajoutée automatiquement)
sudo cscli collections install crowdsecurity/sshd
# Mettre à jour l'index du Hub puis les éléments installés
sudo cscli hub update
sudo cscli hub upgrade
# Recharger le moteur pour prendre en compte les changements
sudo systemctl reload crowdsec
```
Pensez à automatiser la mise à jour du Hub. Les scénarios évoluent constamment pour suivre les nouvelles techniques d’attaque ; un Hub à jour, c’est une détection à jour. Une tâche `cron` hebdomadaire exécutant `cscli hub update && cscli hub upgrade` suivie d’un `systemctl reload crowdsec` est une excellente pratique. Vous pouvez inspecter en détail un scénario installé avec `cscli scenarios inspect crowdsecurity/ssh-bf` pour comprendre ses seuils de déclenchement.

## Étape 6 – Sécuriser l’accès SSH contre le brute-force

SSH est la cible numéro un de tout serveur Linux exposé. La collection `crowdsecurity/sshd` détecte les attaques par force brute et les tentatives d’énumération d’utilisateurs en analysant `/var/log/auth.log` (ou le journal systemd selon votre distribution). Vérifions qu’elle est bien active.

```
# Confirmer que la collection SSH est chargée
sudo cscli collections list | grep sshd
# Lister les scénarios SSH actifs
sudo cscli scenarios list | grep ssh
# Voir les décisions actives (IP actuellement bannies)
sudo cscli decisions list
# Voir l'historique des alertes (déclenchements de scénarios)
sudo cscli alerts list
```
La distinction **decisions** / **alerts** est fondamentale et source de confusion fréquente. Une *alerte* est l’enregistrement historique d’un scénario qui s’est déclenché (« telle IP a tenté un brute-force SSH à telle heure »). Une *décision* est l’action en cours qui en découle (« telle IP est bannie pour 4 heures »). Une décision expire ; une alerte reste tracée. Pour bannir manuellement une IP de test et observer le mécanisme :

