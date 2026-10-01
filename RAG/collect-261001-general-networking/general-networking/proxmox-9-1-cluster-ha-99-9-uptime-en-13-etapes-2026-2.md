---
id: collect-261001-general-networking/general-networking/proxmox-9-1-cluster-ha-99-9-uptime-en-13-etapes-2026-2
title: "Télécharger l'ISO et la somme SHA-256"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["attention", "datacenter", "mai"]
source: docs/RAG/collect-261001-general-networking/proxmox-9-1-cluster-ha-99-9-uptime-en-13-etapes-2026.md
source_anchor: ""
source_lines: [32, 129]
sha256: f999945fcf5246ec8101454001915ca9523cd9b9d9b1e6e7fef24653c2c98e2f
---

# Télécharger l'ISO et la somme SHA-256

Rendez-vous sur le site officiel proxmox.com/downloads et téléchargez « Proxmox VE 9.1 ISO Installer » (ou directement la 9.2-1, mise à jour le 21 mai 2026 et désormais proposée aussi en variante arm64 depuis le 5 août 2026 pour les architectures ARM). Privilégiez l’un des miroirs européens (notamment celui d’OVHcloud à Roubaix) pour obtenir un débit constant supérieur à 100 Mio/s. Si vous installez plusieurs nœuds, hébergez l’ISO sur un serveur HTTP local ou un partage Samba pour éviter de la retélécharger à chaque fois.

```
# Télécharger l'ISO et la somme SHA-256
wget https://enterprise.proxmox.com/iso/proxmox-ve_9.1-1.iso
wget https://enterprise.proxmox.com/iso/proxmox-ve_9.1-1.iso.sha256
# Vérifier l'intégrité
sha256sum -c proxmox-ve_9.1-1.iso.sha256
# Sortie attendue
# proxmox-ve_9.1-1.iso: OK
```
Une fois la vérification réussie, écrivez l’ISO sur une clé USB. Sous Linux, la commande dd reste la méthode la plus fiable : repérez le périphérique avec lsblk avant chaque exécution, sous peine d’écraser votre disque système. Sous Windows, Rufus en mode « DD » fonctionne sans surprise. Évitez Ventoy avec Proxmox : la 9.1 utilise un installateur graphique GTK qui détecte mal les images chainloadées par Ventoy et provoque des erreurs de montage du pool ZFS racine.

```
# Identifier la clé USB (attention au nom du périphérique)
lsblk
# Écrire l'ISO sur la clé (remplacer /dev/sdX)
sudo dd if=proxmox-ve_9.1-1.iso of=/dev/sdX bs=4M status=progress oflag=sync
sudo sync
```
## Étape 2 : Installer Proxmox VE sur bare metal

Bootez le serveur sur la clé USB et choisissez « Install Proxmox VE (Graphical) » dans le menu. L’installateur démarre une interface GTK très lisible. Acceptez la licence GNU AGPL v3, puis arrivez sur l’écran de sélection du disque. C’est l’étape la plus critique. Si vous possédez deux disques identiques de 480 Gio ou plus, optez pour ZFS RAID-1 dans le menu « Options ». Vous obtiendrez un système racine miroir capable de survivre à une panne disque, avec snapshots, compression LZ4 par défaut et tolérance silencieuse aux erreurs de lecture.

Si vous n’avez qu’un seul SSD NVMe, restez en ext4 sur LVM-Thin. C’est moins puissant mais plus simple à dépanner. Évitez absolument btrfs pour la production : le support reste expérimental dans Proxmox 9.1, et plusieurs rapports de corruption sont remontés sur le bugtracker depuis 2025. Configurez ensuite votre fuseau horaire (Europe/Paris pour la France métropolitaine), votre clavier (« French ») et un mot de passe root robuste de seize caractères minimum.

L’écran réseau vous demande une IP fixe, un masque, une passerelle, un serveur DNS et un nom de domaine pleinement qualifié. Choisissez par exemple pve1.lan.exemple.fr avec l’IP 192.168.1.10/24. Évitez les noms d’hôte génériques comme « proxmox » : si vous formez plus tard un cluster, chaque nœud doit avoir un FQDN unique. Lancez l’installation, comptez sept à douze minutes, puis redémarrez sans la clé USB. Le serveur affiche une URL d’accès du type https://192.168.1.10:8006 – c’est votre point d’entrée web.

## Étape 3 : Configurer les dépôts no-subscription

Au premier login web, une fenêtre vous prévient qu’il n’y a pas d’abonnement valide. C’est normal et sans conséquence sur les fonctionnalités. Pour recevoir les mises à jour, il faut basculer du dépôt « enterprise » au dépôt « no-subscription », sinon apt échoue. Ouvrez un shell SSH ou utilisez la console web intégrée puis éditez les fichiers de sources.

```
# Désactiver le dépôt enterprise
sed -i 's|^deb|#deb|' /etc/apt/sources.list.d/pve-enterprise.list
sed -i 's|^deb|#deb|' /etc/apt/sources.list.d/ceph.list 2>/dev/null
# Ajouter le dépôt no-subscription
cat > /etc/apt/sources.list.d/pve-no-subscription.list << 'EOF'
deb http://download.proxmox.com/debian/pve trixie pve-no-subscription
EOF
# Ajouter le dépôt Ceph Squid no-subscription
cat > /etc/apt/sources.list.d/ceph-no-subscription.list << 'EOF'
deb http://download.proxmox.com/debian/ceph-squid trixie no-subscription
EOF
# Mettre à jour le système
apt update && apt full-upgrade -y
reboot
```
Au redémarrage, vérifiez la version installée. La commande pveversion vous renvoie un résumé court, tandis que pveversion -v détaille chaque paquet. Vous devez voir au minimum proxmox-ve: 9.1.0, pve-manager: 9.1.x et un noyau 6.14.8-2-pve. Si vous restez sur Proxmox 8.4 pour un cluster existant, la procédure de migration est documentée dans le wiki officiel mais ne s’improvise pas : prévoyez une fenêtre de maintenance, sauvegardez tout via Proxmox Backup Server, et migrez nœud par nœud après avoir vidé chaque hôte par live migration.

```
# Vérifier la version installée
pveversion -v
# Sortie attendue (extrait)
# proxmox-ve: 9.1.0 (running kernel: 6.14.8-2-pve)
# pve-manager: 9.1.3
# pve-kernel-6.14: 6.14.8-2
# qemu-server: 10.0.2
# pve-container: 6.0.4
# zfsutils-linux: 2.3.3
# ceph: 19.2.3
```
## Étape 4 : Créer un pool ZFS pour les machines virtuelles

ZFS reste le différenciateur majeur de Proxmox face aux hyperviseurs concurrents. La version 2.3.3 livrée avec Proxmox 9.1 ajoute RAIDZ Expansion stable, dnodesize auto par défaut et une compression Zstd niveau 3 désormais accélérée par les instructions AVX2 sur les CPU récents. Pour ce tutoriel nous créons un pool « tank » sur deux disques de 4 To en mode miroir. Cette configuration tolère la perte d’un disque sans interruption de service et offre de meilleures performances aléatoires qu’un RAIDZ1 sur trois disques.

```
# Lister les disques disponibles
lsblk -d -o NAME,SIZE,MODEL,SERIAL
# Effacer toute table de partition existante (attention : irréversible)
wipefs -a /dev/sdb /dev/sdc
# Créer un pool ZFS miroir avec ashift=12 pour disques 4Kn
zpool create -f -o ashift=12 \
  -O compression=zstd \
  -O atime=off \
  -O xattr=sa \
  -O acltype=posixacl \
  tank mirror /dev/sdb /dev/sdc
# Vérifier l'état
zpool status tank
zfs list
# Activer la rétention de snapshots
zfs set snapshot_limit=20 tank
```
Une fois le pool créé en ligne de commande, il faut l’enregistrer dans la couche de stockage Proxmox pour qu’il apparaisse comme cible d’installation de VM. Allez dans Datacenter → Storage → Add → ZFS, sélectionnez « tank » dans la liste déroulante, cochez les types de contenu « Disk image » et « Container », laissez Thin provision activé. Vous pouvez aussi le faire en CLI avec pvesm. Le stockage devient immédiatement disponible sur tous les nœuds du cluster, à condition que chaque hôte possède un pool homonyme – c’est une exigence stricte de Proxmox.

```
# Ajouter le pool tank comme stockage Proxmox
pvesm add zfspool tank-vms --pool tank --content images,rootdir --sparse 1
# Vérifier la liste des stockages
pvesm status
```
## Étape 5 : Créer votre première machine virtuelle Debian

Téléchargez d’abord une image ISO d’invité dans le stockage « local ». Depuis l’interface web, naviguez vers pve1 → local (pve1) → ISO Images → Upload. Vous pouvez aussi pointer vers une URL avec « Download from URL ». Pour ce tutoriel, prenons l’ISO netinst de Debian 13 (debian-13.0.0-amd64-netinst.iso) – 758 Mio, dépôt français mirror.debian.org à l’INRIA Lyon.

Créez ensuite la VM avec la commande qm. L’ID 100 est conventionnellement réservé aux VM système, on l’utilise donc pour cette VM de démonstration. Choisissez le mode UEFI (OVMF) plutôt que SeaBIOS : il est requis par Windows 11 et les futures Linux 7.x, et n’apporte aucune pénalité de performance sur les invités modernes. Le disque scsi0 sur tank-vms profite directement de la compression Zstd et des snapshots ZFS.

