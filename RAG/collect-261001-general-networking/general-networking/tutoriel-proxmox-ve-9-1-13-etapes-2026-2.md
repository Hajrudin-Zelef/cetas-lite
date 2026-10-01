---
id: collect-261001-general-networking/general-networking/tutoriel-proxmox-ve-9-1-13-etapes-2026-2
title: "Téléchargement de l'ISO Proxmox VE 9.1"
domain: general-networking
role: reference
task: reference
actors: ["AMD", "Intel", "Microsoft"]
dates: []
keywords: ["amd", "attention", "gpu", "intel", "mai"]
source: docs/RAG/collect-261001-general-networking/tutoriel-proxmox-ve-9-1-13-etapes-2026.md
source_anchor: ""
source_lines: [30, 86]
sha256: a36d204fa9527756899298ae5a2b74ed63d56e562cb3f7c379563c964c1d827c
---

# Téléchargement de l'ISO Proxmox VE 9.1

Rendez-vous sur la page officielle *Downloads* de Proxmox — dont l’entrée « USB stick » du wiki a elle aussi été mise à jour en mai 2026 — pour récupérer l’ISO. L’ISO 9.1-1 datée du 19 novembre 2025 pèse environ 1,6 Go, tout comme le fichier `proxmox-ve_9.2-1.iso` publié en mai 2026 (également 1,6 Go) et son pendant ARM64 référencé 9.2-1-arm64 ; la page de checksums officielle, qui publie le hash SHA-256 de chaque image, a été actualisée en août 2026. Vérifiez systématiquement cette somme de contrôle publiée à côté du lien : un fichier corrompu durant le téléchargement génère des erreurs aléatoires durant l’installation, parfois plusieurs étapes plus loin, et fait perdre de longues minutes de diagnostic — une bonne pratique que détaille aussi le guide TechFuelHQ du 19 juin 2026 consacré à la création de clé USB via `dd`, ainsi que le tutoriel de credativ GmbH publié en février 2026, qui rappelle qu’une clé USB de 8 Go suffit largement pour flasher l’ISO stable.

```
# Téléchargement de l'ISO Proxmox VE 9.1
curl -O https://enterprise.proxmox.com/iso/proxmox-ve_9.1-1.iso
# Vérification de l'empreinte SHA-256
sha256sum proxmox-ve_9.1-1.iso
# Comparez la sortie avec la valeur affichée sur proxmox.com/downloads
# Exemple attendu : a1b2c3d4...  proxmox-ve_9.1-1.iso
# Création de la clé USB sous Linux (remplacez /dev/sdX par votre périphérique)
sudo dd if=proxmox-ve_9.1-1.iso of=/dev/sdX bs=4M status=progress oflag=sync
# Sous macOS (vérifiez le numéro de disque avec diskutil list)
diskutil unmountDisk /dev/disk4
sudo dd if=proxmox-ve_9.1-1.iso of=/dev/rdisk4 bs=1m
diskutil eject /dev/disk4
```
Sous Windows, préférez balenaEtcher ou Rufus en mode « DD Image » : Rufus en mode ISO produit parfois une clé qui démarre mais affiche un écran noir au boot. Pour les installations déployées à grande échelle, utilisez le mode PXE : Proxmox a posé les bases de l’installateur réseau dès la version 8.2, un outil d’installation automatisée pour le bare-metal que la newsletter officielle du projet mettait encore en avant début mars 2026, avant que la 8.4 ne généralise l’installation automatisée/sans surveillance (« unattended installation ») — une trajectoire vers un déploiement rationalisé amorcée dès avril 2025 et documentée par le blog LeGoutDuLibre le 19 novembre 2025. La feuille de route s’est encore étoffée avec l’ajout des options `--pxe` et `--pxe-loader`, ainsi que d’une nouvelle sous-commande `inspect-iso`, à la commande `proxmox-auto-install-assistant prepare-iso` — des ajouts documentés par Datazone dès la sortie de la 9.2 le 21 mai 2026 — qui affinent le choix du chargeur d’amorçage réseau pour les déploiements PXE à grande échelle ; la page officielle des téléchargements liste par ailleurs, depuis août 2026, un installeur ISO 9.2 dédié aux architectures ARM64, élargissant Proxmox VE au-delà du x86_64.

## Étape 2 : Configurer le BIOS/UEFI et démarrer sur la clé

Branchez la clé USB sur la machine cible et entrez dans le firmware (touche F2, F12, Suppr ou Esc selon le constructeur). Trois réglages conditionnent la suite : **activez la virtualisation matérielle** (Intel VT-x ou AMD-V), **activez l’IOMMU** (Intel VT-d ou AMD IOMMU) si vous prévoyez du passthrough GPU ou NIC, et **désactivez le Secure Boot** car Proxmox n’est pas signé Microsoft. Sur les serveurs Dell PowerEdge, ces options se trouvent dans System BIOS > Processor Settings ; sur HPE ProLiant, dans System Configuration > BIOS/Platform > System Options ; sur Supermicro, dans Advanced > CPU Configuration.

Réglez aussi l’ordre de démarrage pour placer l’USB devant le disque dur. Sur les machines avec contrôleur RAID matériel (PERC, MegaRAID, SmartArray), basculez le contrôleur en mode *HBA* ou *JBOD* si vous comptez utiliser ZFS — l’usage de ZFS au-dessus d’un RAID matériel est explicitement déconseillé par la documentation Proxmox car la double abstraction empêche ZFS de gérer correctement les checksums et la reconstruction. À l’inverse, si vous utilisez ext4 ou XFS, le RAID matériel reste pertinent.

## Étape 3 : Lancer l’installateur graphique et choisir le filesystem

Au démarrage de la clé, l’écran d’accueil propose plusieurs entrées. Sélectionnez « Install Proxmox VE (Graphical) » sauf si votre carte graphique pose problème, auquel cas optez pour le mode terminal. Acceptez la licence EULA, puis arrivez à l’écran clé : le choix du système de fichiers cible. Quatre options s’offrent à vous.

| Filesystem | Cas d’usage | Avantages | Inconvénients | 
|---|---|---|---|
| ext4 | Mono-disque, lab, station | Simple, rapide, journalisé | Pas de snapshot natif, pas de RAID logiciel | 
| XFS | Gros fichiers, débit séquentiel | Performant sur grandes VM | Pas de snapshot natif | 
| ZFS RAID0 | Cache rapide, scratch | Snapshots, compression LZ4 | Aucune redondance | 
| ZFS RAID1/Z1/Z2/Z3 | Production, redondance | Snapshots, replication, scrubs, compression, déduplication | Coût RAM (1 Go/To) | 

Pour un serveur de production avec deux disques NVMe, le ZFS RAID1 (mirror) est la valeur la plus sûre. La nouveauté 2025 majeure est le support de l’expansion RAIDZ : vous pouvez désormais ajouter un disque à un pool RAID-Z existant sans tout reconstruire, fonctionnalité réclamée depuis dix ans par la communauté ZFS et finalement disponible avec OpenZFS 2.3 inclus dans Proxmox 9.0. Configurez ashift=12 pour les disques 4K modernes et compression=lz4 dans les options avancées.

## Étape 4 : Définir mot de passe, hostname et réseau

L’écran suivant demande le pays (sélectionnez France pour le fuseau Europe/Paris), un mot de passe administrateur et une adresse e-mail. Ce mot de passe est celui du compte `root` Linux ; vous pourrez créer ensuite des utilisateurs Proxmox dédiés via PAM, LDAP ou Microsoft Entra ID. L’e-mail reçoit les notifications de pannes ZFS, de mises à jour disponibles et d’échecs de sauvegardes : ne mettez pas une adresse fictive.

L’écran réseau réclame une attention particulière. Définissez un **FQDN** complet du type `pve01.lab.example.fr` et non simplement `pve01` : la commande `pvecm` qui crée le cluster utilise ce hostname pour générer les certificats. Renseignez l’IPv4, le masque CIDR (24, 22, etc.), la passerelle et un serveur DNS atteignable — depuis la version 9.2 sortie le 21 mai 2026, l’installateur accepte aussi une configuration IPv6 seule, avec auto-configuration SLAAC et annonces de routeur (Router Advertisement), et les fichiers de réponse (answer files) peuvent désormais activer automatiquement la clé d’abonnement, selon la documentation Datazone de mai 2026. Sur un cluster, choisissez idéalement un réseau dédié pour Corosync (le bus de cluster), distinct du réseau de management et du réseau de stockage.

```
# Vérification post-installation : fichier /etc/hosts
cat /etc/hosts
# Sortie attendue :
# 127.0.0.1       localhost.localdomain localhost
# 192.168.10.11   pve01.lab.example.fr pve01
# Si le hostname est mal configuré, corrigez :
hostnamectl set-hostname pve01.lab.example.fr
nano /etc/hosts  # remplacez la ligne 127.0.1.1 si presente
```
## Étape 5 : Premier accès à l’interface web et désactivation du repo Enterprise

Après le redémarrage, l’écran console affiche l’URL du panneau d’administration : `https://192.168.10.11:8006`. Le port 8006 est par défaut depuis Proxmox VE 1.0. Connectez-vous avec l’utilisateur `root`, le realm *Linux PAM standard authentication* et le mot de passe défini à l’étape 4. Au premier accès, un avertissement de certificat auto-signé apparaît : c’est normal, vous le remplacerez plus tard par un certificat Let’s Encrypt.

