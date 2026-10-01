---
id: collect-261001-general-networking/general-networking/proxmox-9-1-cluster-ha-99-9-uptime-en-13-etapes-2026-4
title: "Télécharger l'ISO et la somme SHA-256"
domain: general-networking
role: reference
task: reference
actors: ["AMD", "DeepSeek", "Intel", "Nvidia"]
dates: []
keywords: ["amd", "arr", "datacenter", "deepseek", "gpu", "intel", "llama", "llama.cpp", "mai", "nvidia"]
source: docs/RAG/collect-261001-general-networking/proxmox-9-1-cluster-ha-99-9-uptime-en-13-etapes-2026.md
source_anchor: ""
source_lines: [262, 373]
sha256: 3a5044098edb345c8c2e7bf909b694ba15172856dd8b4b4875ee762caed3a411
---

# Télécharger l'ISO et la somme SHA-256

Proxmox Backup Server (PBS) est la solution officielle de sauvegarde, distincte de Proxmox VE mais conçue pour s’y intégrer parfaitement. La branche 3.x s’est arrêtée avec l’ISO PBS 3.4-1, dont la dernière mise à jour officielle remonte au 10 avril 2025 selon la page de téléchargement Proxmox, avant que la 4.x ne prenne le relais. La version 4.2-1 publiée le 29 avril 2026 introduit la déduplication multi-niveau et la vérification incrémentale, ce qui divise par trois la durée des « verify jobs » sur un datastore de plus de 50 Tio. Installez PBS sur une machine séparée – physique de préférence – avec un pool ZFS RAIDZ2 pour la résilience. Une fois installé, ajoutez-le comme stockage dans Proxmox VE.

```
# Sur PBS – créer un utilisateur dédié et un token API
proxmox-backup-manager user create backup@pbs --password "$(openssl rand -base64 24)"
proxmox-backup-manager acl update /datastore/main DatastoreBackup --auth-id backup@pbs
# Générer un token (sortie : valeur secrète à copier)
proxmox-backup-manager user generate-token backup@pbs pve-token
# Sur Proxmox VE – ajouter PBS comme stockage
pvesm add pbs pbs-main \
  --server 192.168.1.20 \
  --datastore main \
  --username backup@pbs \
  --password "TOKEN_SECRET" \
  --fingerprint "$(ssh [email protected] proxmox-backup-manager cert info | grep -i Fingerprint | awk '{print $3}')"
# Lancer un backup manuel de la VM 100
vzdump 100 --mode snapshot --storage pbs-main --notes-template "{{guestname}} {{node}}"
```
L’avantage de PBS sur un simple vzdump local tient à la déduplication par blocs de 4 Mio (chunks). Sauvegarder dix VM Debian quasi-identiques ne consomme guère plus que la place d’une seule VM, car les blocs partagés ne sont stockés qu’une fois. Sur nos clusters de production, le ratio dedup observé varie entre 4,2:1 et 12:1 selon l’homogénéité des invités. Programmez ensuite des jobs quotidiens dans Datacenter → Backup avec une rétention 7 quotidiens + 4 hebdomadaires + 12 mensuels, et activez la vérification automatique toutes les 48 heures.

## Étape 11 : Configurer le passthrough GPU pour une VM IA

Le GPU passthrough est l’un des cas d’usage les plus demandés en 2026, notamment pour faire tourner des modèles locaux Llama 4 ou DeepSeek V3 sur une VM Linux ou Windows. Le principe : retirer la carte graphique du contrôle de l’hôte Proxmox pour l’attribuer exclusivement à une VM via le driver vfio-pci. Cela exige IOMMU activé dans le BIOS et un GPU dans un groupe IOMMU bien isolé.

```
# 1. Activer IOMMU dans GRUB (Intel)
sed -i 's|GRUB_CMDLINE_LINUX_DEFAULT="quiet"|GRUB_CMDLINE_LINUX_DEFAULT="quiet intel_iommu=on iommu=pt"|' /etc/default/grub
update-grub
# Pour AMD : amd_iommu=on iommu=pt
# 2. Activer les modules VFIO au boot
cat > /etc/modules-load.d/vfio.conf << 'EOF'
vfio
vfio_iommu_type1
vfio_pci
EOF
# 3. Identifier le GPU et son ID PCI
lspci -nn | grep -i nvidia
# 01:00.0 VGA compatible controller [0300]: NVIDIA Corporation [10de:2684]
# 01:00.1 Audio device [0403]: NVIDIA Corporation [10de:22ba]
# 4. Lier le GPU à vfio-pci au boot
echo "options vfio-pci ids=10de:2684,10de:22ba disable_vga=1" > /etc/modprobe.d/vfio.conf
# 5. Blacklister les drivers nvidia côté hôte
cat > /etc/modprobe.d/blacklist-nvidia.conf << 'EOF'
blacklist nouveau
blacklist nvidia
blacklist nvidia_drm
blacklist nvidia_modeset
EOF
# 6. Régénérer initramfs et redémarrer
update-initramfs -u -k all
reboot
```
Après redémarrage, vérifiez avec lspci -k -s 01:00.0 que le « Kernel driver in use » indique bien vfio-pci. Vous pouvez alors ajouter le GPU à la VM 100 via l’interface web : Hardware → Add → PCI Device → Raw Device, sélectionnez 0000:01:00.0, cochez « All Functions », « ROM-Bar » et « PCI-Express ». La VM démarre avec la carte exclusivement attribuée et peut installer le pilote NVIDIA propriétaire 565.x ou supérieur, indispensable pour Llama.cpp ou Ollama. Notez que tant que la VM est démarrée, l’hôte ne voit plus la carte – c’est le prix du passthrough.

## Étape 12 : Déployer un stockage Ceph hyperconvergé

Ceph transforme votre cluster Proxmox en stockage distribué hyperconvergé : chaque nœud apporte ses disques au pool global, les données sont répliquées en trois copies (ou en erasure coding 4+2), et les VM bougent entre nœuds sans déplacement de blocs. La version Squid 19.2.3 livrée avec Proxmox 9.1 stabilise enfin BlueStore v2, qui réduit l’amplification d’écriture de 35 % sur les SSD NVMe par rapport à BlueStore v1 – sachant que Proxmox VE 9.2, sorti en mai 2026, bascule déjà la version par défaut vers Ceph Tentacle 19.x, selon ComputingForGeeks. Pour un cluster de trois nœuds avec deux SSD NVMe 1 Tio chacun, vous obtenez environ 1,8 Tio utile en réplication 3.

```
# Sur chaque nœud – installer les paquets Ceph
pveceph install --repository no-subscription
# Sur pve1 – initialiser Ceph avec le réseau dédié
pveceph init --network 10.0.0.0/24
# Créer le premier moniteur
pveceph mon create
# Créer le manager
pveceph mgr create
# Sur pve2 et pve3 – créer moniteurs et managers
pveceph mon create
pveceph mgr create
# Sur chaque nœud – créer les OSD (Object Storage Daemons)
pveceph osd create /dev/nvme1n1
pveceph osd create /dev/nvme2n1
# Créer un pool RBD pour les VM
pveceph pool create vm-pool --size 3 --min_size 2 --pg_autoscale_mode on --add_storages
# Vérifier la santé du cluster
ceph -s
```
La sortie de ceph -s doit afficher « HEALTH_OK ». Si vous voyez « HEALTH_WARN: not enough PGs per OSD », ajustez le nombre de placement groups : `ceph osd pool set vm-pool pg_num 128`. Pour un cluster trois nœuds avec six OSD au total, la formule recommandée par la documentation Ceph est `(OSD × 100) / size`, soit 200 PGs arrondi à la puissance de deux la plus proche, donc 256. Une fois Ceph stable, déplacez vos VM critiques dessus avec `qm move_disk 100 scsi0 ceph-vm-pool`. Vous bénéficierez d’une bascule HA en moins de quinze secondes en cas de panne d’un nœud.

## Étape 13 : Sécuriser et superviser votre cluster en production

Un cluster Proxmox en production doit être verrouillé. Première mesure : interdire l’authentification root par mot de passe via SSH, n’autoriser que les clés publiques. Deuxième mesure : créer un utilisateur d’administration dédié dans le royaume PAM avec rôle PVEAdmin, et désactiver le login web du compte root@pam. Troisième mesure : ne jamais exposer le port 8006 directement sur Internet – placez un reverse-proxy comme Caddy ou Nginx en frontal avec authentification mTLS ou WireGuard.

```
# Créer un utilisateur admin et un token
pveum useradd admin@pve --password "$(openssl rand -base64 24)"
pveum aclmod / -user admin@pve -role PVEAdmin
pveum user token add admin@pve api-token --privsep 0
# Activer le firewall du datacenter (mode permissif d'abord)
cat > /etc/pve/firewall/cluster.fw << 'EOF'
[OPTIONS]
enable: 1
policy_in: DROP
policy_out: ACCEPT
[RULES]
IN ACCEPT -source +management -p tcp -dport 8006 -log nolog
IN ACCEPT -source +management -p tcp -dport 22 -log nolog
IN ACCEPT -source +ceph-net -p tcp -dport 3300 -log nolog
EOF
# Définir les ipsets
cat >> /etc/pve/firewall/cluster.fw << 'EOF'
[IPSET management]
192.168.99.0/24
[IPSET ceph-net]
10.0.0.0/24
EOF
# Activer le firewall sur chaque nœud
pvenode config set --firewall 1
```
Côté supervision, Proxmox expose nativement des métriques au format InfluxDB ou Graphite via Datacenter → Metric Server. Pointez-les vers un container LXC hébergeant InfluxDB 3 et Grafana – qui s’installe en cinq minutes via la stack Prometheus/Grafana. Vous obtenez des dashboards prêts à l’emploi (Grafana ID 10048) avec CPU, RAM, IO, latence Ceph, débit ZFS et état du quorum. Ajoutez Alertmanager pour recevoir des notifications par email ou par webhook Mattermost dès qu’un OSD passe en down ou qu’un nœud quitte le quorum.

