---
id: collect-261001-general-networking/general-networking/proxmox-9-1-cluster-ha-99-9-uptime-en-13-etapes-2026-5
title: "Télécharger l'ISO et la somme SHA-256"
domain: general-networking
role: reference
task: reference
actors: ["AMD", "Broadcom", "Intel"]
dates: []
keywords: ["agent", "amd", "benchmarks", "datacenter", "intel", "mai"]
source: docs/RAG/collect-261001-general-networking/proxmox-9-1-cluster-ha-99-9-uptime-en-13-etapes-2026.md
source_anchor: ""
source_lines: [374, 430]
sha256: 1b1d3d2a833ae772078a39493ba94bef30c54c21859e4acbb2c29ef30dae82df
---

# Télécharger l'ISO et la somme SHA-256

## Comparaison Proxmox VE 9.1 vs VMware ESXi 9 vs XCP-ng 8.3

Le marché de l’hyperviseur s’est redessiné en 2025-2026. Voici un comparatif factuel des trois principales solutions sur les critères qui pèsent en homelab et en PME française. Les tarifs VMware reflètent la grille Broadcom de mars 2026 après la dernière revalorisation de 8 %.

| Critère | Proxmox VE 9.1 | VMware ESXi 9 + vCenter | XCP-ng 8.3 | Hyper-V Server | 
|---|---|---|---|---|
| Licence | AGPLv3 gratuit | Abonnement par cœur | GPLv2 gratuit | Inclus Windows Server | 
| Prix par cœur/an | 0 € (ou 110 € support) | 350 USD minimum | 0 € (ou Vates Pro) | ≈ 65 € via licence WS | 
| Cluster gratuit | Oui, illimité | Non, vCenter requis | Oui via Xen Orchestra | Failover Cluster gratuit | 
| HA + migration live | Oui natif | Oui mais via vCenter | Oui natif | Oui via SCVMM | 
| Stockage intégré | ZFS + Ceph + LVM | vSAN payant | XOSAN payant | Storage Spaces Direct | 
| Containers natifs | LXC inclus | Non | Non | Windows Containers | 
| API REST | Oui complète | Oui | Oui (Xen Orchestra) | WMI + PowerShell | 
| Communauté | 100 000+ sujets forum | Réduite post-Broadcom | Active mais petite | Documentation MS | 
| Migration depuis VMware | Outil natif esxi-import | – | Outil v2v | Limitée | 

L’outil esxi-import livré avec Proxmox 9.1 lit directement les VMDK depuis une datastore VMware via SSH et reconstruit les VM avec les bons drivers virtio. Sur nos tests, la migration d’un cluster ESXi 8 de 24 VM Windows Server vers Proxmox 9.1 prend environ six heures de travail effectif, en incluant la conversion des disques. Le retour sur investissement se calcule en mois et non en années : sur un parc moyen de 32 cœurs, l’économie annuelle dépasse 10 000 euros, hors gains sur vSAN et autres add-ons.

## Pièges courants et messages d’erreur fréquents

Au fil de plusieurs centaines d’installations Proxmox, certaines erreurs reviennent avec une régularité quasi mathématique. En voici huit que vous croiserez probablement, et leur résolution éprouvée.

- **« No subscription key »** au login web : c’est un simple rappel, pas une erreur. Le warning disparaît si vous achetez un abonnement, ou peut être masqué avec un patch javascript non-officiel – déconseillé en production.
- **« KVM: entry failed, hardware error »** au démarrage d’une VM : vérifiez que VT-x/AMD-V est bien activé dans le BIOS. Sur Dell PowerEdge, le menu s’appelle « Virtualization Technology » sous « Processor Settings ».
- **« Cannot find a valid network configuration »** à l’installation : symptôme d’un câble réseau non branché ou d’un VLAN tagué sans configuration. Branchez sur un port d’accès du switch ou ajoutez –net-vlan-id à l’installateur.
- **« Corosync TOTEM RETRANSMIT »** dans les logs cluster : votre lien Corosync sature. Séparez Corosync du trafic Ceph et VM en utilisant une carte gigabit dédiée, ou ajoutez un second lien Corosync via pvecm.
- **« ZFS pool can’t import »** après un reboot : généralement un disque externe non détecté à temps. Ajoutez ZFS_IMPORT_TIMEOUT=60 dans /etc/default/zfs.
- **« qmp command ‘guest-ping’ failed: Got timeout »** : l’agent QEMU n’est pas installé dans la VM. Installez qemu-guest-agent depuis la VM et redémarrez-la.
- **« HA: fence node failed »** en cas de bascule HA : votre nœud n’a pas de mécanisme de fencing matériel. Configurez un watchdog softdog ou IPMI dans /etc/default/pve-ha-manager.
- **« PVE Backup: VM lock timeout »** : un job vzdump précédent s’est mal terminé. Effacez le verrou avec qm unlock VMID puis relancez.

## Conseils avancés pour gagner en performance

Une fois votre cluster fonctionnel, plusieurs réglages fins boostent les performances de 20 à 60 % selon les charges. D’abord, activez le numa=on sur toutes les VM dont la RAM dépasse 16 Gio : QEMU exposera la topologie NUMA réelle de l’hôte au lieu d’un seul nœud virtuel, ce qui réduit drastiquement la latence d’accès mémoire sur les CPU Xeon Gold à plusieurs sockets. Sur un Dell R740 bi-socket, on observe un gain de 28 % sur SQL Server 2022.

Ensuite, configurez systématiquement `iothread=1` et `aio=io_uring` sur les disques virtio-scsi. Le moteur io_uring du kernel Linux 6.14 surpasse l’ancien aio=native sur tous les benchmarks fio aléatoires 4K (jusqu’à 2,3× sur SSD NVMe). Pour les containers LXC, utilisez systemd-cgroup-v2 (activé par défaut sur Debian 13) et fixez explicitement les CPU avec `pct set 200 --cores 2 --cpulimit 2` pour éviter qu’un container CPU-bound n’affame les autres.

Enfin, n’oubliez pas l’IRQ affinity. Sur les serveurs avec plusieurs cartes 10 GbE, distribuez manuellement les interruptions des NIC sur des cores différents de ceux dédiés à Ceph ou à ZFS. La commande `set_irq_affinity.sh ethX` du paquet Intel ixgbe-utils automatise la tâche. Pour une charge Ceph intense, gagnez encore 15 % en activant le tunable `osd_op_num_threads_per_shard = 4` dans /etc/pve/ceph.conf.

## Cas d’usage complet : hébergement WordPress haute disponibilité

Voici un projet complet typique pour conclure ce tutoriel. Vous hébergez plusieurs sites WordPress à forte croissance et vous voulez une plateforme privée HA. Sur un cluster trois nœuds Proxmox + Ceph, déployez l’architecture suivante : trois containers LXC Debian 13 avec Nginx + PHP-FPM 8.4 derrière HAProxy, un container MariaDB Galera Cluster en trois nœuds répliqués synchrones, un container Redis pour le cache d’objets, un container Proxmox Backup Client qui envoie chaque heure des snapshots incrémentaux vers PBS, et un container Caddy 2.10 en frontal avec Let’s Encrypt automatique.

L’ensemble tient sur 12 vCPU et 24 Gio de RAM, soit moins d’un quart d’un cluster trois nœuds Dell R650. Avec Ceph en réplication 3, vous survivez à la panne complète d’un nœud sans perte de données ni interruption client. Et l’addition reste sous 6 000 euros pour le matériel d’occasion, contre 24 000 euros pour une stack VMware équivalente vSphere Foundation. C’est cette équation économique qui explique pourquoi le baromètre IT-Connect d’avril 2026 mesure une part de marché Proxmox passée de 8 % à 19 % en France sur l’année écoulée.

## FAQ Proxmox VE 9.1 – Questions fréquentes 2026

### Proxmox VE est-il vraiment gratuit pour un usage commercial ?

Oui, sans limitation fonctionnelle. La licence AGPLv3 autorise l’usage commercial illimité, y compris l’hébergement de clients tiers. Le dépôt « no-subscription » est totalement libre. L’abonnement payant (Community à partir de 115 euros par CPU et par an depuis la grille tarifaire introduite avec le lancement de VE 9.0 en août 2025, Standard à 330 euros, Premium à 990 euros) n’apporte que le support technique officiel et l’accès aux paquets « enterprise » plus testés. La distinction est purement contractuelle, pas technique. C’est l’une des grandes différences avec le modèle Broadcom où une mauvaise licence bloque vMotion ou vSAN.

### Combien de nœuds maximum dans un cluster Proxmox ?

La limite théorique du quorum Corosync se situe à 32 nœuds dans un même cluster, et c’est la valeur supportée officiellement par Proxmox. En pratique, au-delà de 16 nœuds, les latences de Corosync demandent un réseau dédié 10 GbE avec timing strict. Pour les très grandes infrastructures, l’approche recommandée est de découper en plusieurs clusters de 8 à 12 nœuds, gérés ensuite par une couche fédérative comme Proxmox Datacenter Manager, qui est sorti de bêta pour atteindre la version 1.1 le 21 mai 2026.

### Peut-on faire tourner Proxmox sur un Raspberry Pi ?

