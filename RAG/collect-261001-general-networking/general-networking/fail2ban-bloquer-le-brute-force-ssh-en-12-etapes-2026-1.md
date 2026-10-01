---
id: collect-261001-general-networking/general-networking/fail2ban-bloquer-le-brute-force-ssh-en-12-etapes-2026-1
title: "Vérifier la version installée"
domain: general-networking
role: reference
task: reference
actors: []
dates: ["2024-04-25"]
keywords: ["apache", "incident", "mai", "open source"]
source: docs/RAG/collect-261001-general-networking/fail2ban-bloquer-le-brute-force-ssh-en-12-etapes-2026.md
source_anchor: ""
source_lines: [1, 59]
sha256: ae8a4632ee3e80ac07dc4bb74528341ea86b5a75c9616514fa9de01b95e000a5
---

# Vérifier la version installée

Un scan automatisé tente une connexion SSH sur votre serveur toutes les quelques secondes. Sur un VPS exposé sans protection, les journaux d’authentification débordent de tentatives de force brute en quelques heures à peine. **Fail2ban** reste en 2026 l’outil open source de référence pour bannir automatiquement les adresses IP malveillantes, avec environ 18 000 étoiles sur GitHub et une présence quasi systématique sur les distributions serveur Debian et Ubuntu. Ce tutoriel détaille l’installation, la configuration et le durcissement de Fail2ban 1.1.0 sur Ubuntu 24.04 LTS et Debian 12/13, avec le backend systemd, l’action nftables moderne, et les pièges les plus fréquents que rencontrent les administrateurs.

## Qu’est-ce que Fail2ban et pourquoi l’utiliser en 2026

Fail2ban est un démon de prévention d’intrusion écrit en Python qui surveille les journaux système (SSH, Nginx, Apache, Postfix, WordPress) et bannit temporairement les adresses IP responsables d’un nombre excessif d’échecs d’authentification. Le principe est simple : dès qu’un filtre détecte un motif suspect (par exemple cinq échecs de connexion SSH en dix minutes), Fail2ban déclenche une action, généralement l’ajout d’une règle de pare-feu qui bloque l’IP source pendant une durée définie.

La dernière version stable officielle est **Fail2ban 1.1.0**, publiée le 25 avril 2024, et c’est toujours cette branche qui reçoit des correctifs actifs sur le dépôt GitHub officiel du projet à l’été 2026 (le dernier commit de maintenance date de juin 2026). Le projet totalise environ **18 007 étoiles GitHub**, ce qui en fait l’outil de bannissement d’IP le plus populaire sur les serveurs Linux orientés hébergement, devant son challenger le plus sérieux, CrowdSec (environ 13 941 étoiles). Contrairement à un pare-feu classique qui filtre selon des règles statiques, Fail2ban ajoute une couche comportementale : il réagit à ce qui se passe réellement dans vos journaux, en temps quasi réel.

Le projet reste activement documenté par la communauté, notamment via l’ArchWiki, qui fait référence pour la configuration avancée, même au-delà des distributions Arch. Sur le plan des performances, un comparatif indépendant publié en mai 2026 mesure une latence d’application du bannissement de l’ordre de 500 millisecondes à 5 secondes pour Fail2ban entre la détection dans le journal et l’ajout effectif de la règle réseau, contre 200 millisecondes à 2 secondes pour CrowdSec sur le même type de scénario. Sur un serveur qui subit un scan automatisé classique, cet écart reste largement suffisant : l’objectif n’est pas de bloquer la première tentative, mais d’empêcher la centième.

Pour un administrateur système, une PME ou un développeur qui gère un VPS OVHcloud, Scaleway ou Hetzner, Fail2ban répond à un besoin concret : réduire la surface d’attaque SSH sans dépendre d’un service tiers payant, et sans avoir à gérer soi-même des listes d’IP à bannir. C’est aussi une brique essentielle d’une stratégie de défense en profondeur, en complément de l’authentification forte et du chiffrement des accès distants.

En France et en Europe, ce besoin est particulièrement concret pour les hébergeurs souverains et les PME qui gèrent elles-mêmes leur infrastructure plutôt que de tout déléguer à un service managé. Un serveur mal protégé qui subit une compromission via SSH devient rapidement un point d’entrée vers des données personnelles soumises au RGPD, ce qui transforme un incident technique en obligation de notification réglementaire. Fail2ban ne remplace évidemment pas une politique de sécurité complète, mais il réduit mécaniquement le nombre de tentatives de connexion qui atteignent réellement le service SSH, ce qui limite d’autant la probabilité qu’un mot de passe faible ou réutilisé finisse par céder.

## Prérequis et versions testées

Avant de commencer, assurez-vous de disposer de l’environnement suivant. Ce tutoriel a été testé et validé sur les configurations ci-dessous, à jour au 22 août 2026.

| Composant | Version testée | Remarque | 
|---|---|---|
| Fail2ban | 1.1.0 (25/04/2024, branche de maintenance active) | Paquet officiel des dépôts Debian/Ubuntu | 
| Système d’exploitation | Ubuntu 24.04 LTS ou Debian 12 (Bookworm) / 13 (Trixie) | Support LTS jusqu’en 2029 pour Ubuntu 24.04 | 
| Pare-feu backend | nftables (banaction par défaut) | iptables-legacy encore supporté en repli | 
| Backend de journalisation | systemd/journald ou fichiers texte classiques | python3-systemd requis pour le backend systemd | 
| Accès serveur | SSH avec droits sudo ou root | Accès console de secours recommandé (voir pièges) | 
| Python | 3.11 ou supérieur | Inclus nativement sur Ubuntu 24.04 et Debian 12/13 | 

Vous aurez aussi besoin d’un accès à la console web de votre hébergeur (KVM, VNC, ou console de secours façon OVHcloud/Scaleway). C’est votre filet de sécurité si vous vous bannissez vous-même par erreur, un incident classique que nous détaillons plus bas. Comptez environ 90 minutes pour suivre l’ensemble des étapes, tester chaque jail et valider la configuration.

Ce tutoriel part du principe que vous administrez un serveur Linux exposé sur Internet, que ce soit un VPS personnel, un serveur d’entreprise ou une instance de développement accessible à distance. Si vous gérez déjà plusieurs dizaines de serveurs et cherchez une solution avec intelligence de menaces mutualisée entre utilisateurs, la section dédiée à CrowdSec plus bas vous aidera à trancher entre les deux approches selon votre contexte.

## Étape 1 : Installer Fail2ban sur Ubuntu ou Debian

Fail2ban est disponible directement dans les dépôts officiels Ubuntu et Debian, aucun dépôt tiers n’est nécessaire. Commencez par mettre à jour la liste des paquets, puis installez Fail2ban et la bibliothèque Python pour systemd, indispensable si vous voulez utiliser le backend journald plutôt que des fichiers de log classiques.

```
sudo apt update
sudo apt install -y fail2ban python3-systemd
# Vérifier la version installée
fail2ban-client --version
fail2ban-server --version
```
Une fois l’installation terminée, le service se lance et s’active automatiquement au démarrage sur les paquets Debian/Ubuntu récents. Vérifiez tout de même son état avant de continuer.

```
sudo systemctl status fail2ban
sudo systemctl enable fail2ban
```
Si le service refuse de démarrer immédiatement après l’installation, c’est normal : Fail2ban n’a encore aucun jail actif tant que vous n’avez pas créé de configuration locale. Passez à l’étape suivante.

## Étape 2 : Comprendre l’arborescence de configuration

C’est le point que la majorité des tutoriels bâclent, et pourtant c’est la source numéro un de confusion chez les débutants. Fail2ban lit ses fichiers de configuration dans un ordre précis, et chaque fichier suivant peut surcharger le précédent.

- **/etc/fail2ban/jail.conf** — le modèle fourni par le projet en amont. Ne le modifiez jamais directement : il sera écrasé à chaque mise à jour du paquet.
- **/etc/fail2ban/jail.d/defaults-debian.conf** — fichier ajouté par le mainteneur Debian/Ubuntu du paquet. C’est ici que se trouve la vraie valeur par défaut de`banaction` , à savoir`nftables` , alors que`jail.conf` affiche encore`iptables-multiport` en amont.
- **/etc/fail2ban/jail.local** — votre fichier de surcharge globale. C’est ici que vous placez vos réglages personnalisés qui s’appliquent à tous les jails.
- **/etc/fail2ban/jail.d/*.conf** — vos définitions de jails individuels (un fichier par service protégé, par exemple`sshd.local.conf` ou`nginx.local.conf` ).

