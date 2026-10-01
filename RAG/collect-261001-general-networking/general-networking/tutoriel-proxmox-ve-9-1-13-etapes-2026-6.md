---
id: collect-261001-general-networking/general-networking/tutoriel-proxmox-ve-9-1-13-etapes-2026-6
title: "Téléchargement de l'ISO Proxmox VE 9.1"
domain: general-networking
role: reference
task: reference
actors: ["AWS", "Intel", "Microsoft"]
dates: []
keywords: ["aws", "benchmark", "intel", "open source"]
source: docs/RAG/collect-261001-general-networking/tutoriel-proxmox-ve-9-1-13-etapes-2026.md
source_anchor: ""
source_lines: [458, 531]
sha256: 6a4f56ec39deaa1bf534a6155ef3f96f4a8a24b2ed379e36e5c80ec09fa3aa3e
---

# Téléchargement de l'ISO Proxmox VE 9.1

### Erreur 2 : « TASK ERROR: command ‘apt-get update’ failed »

Cause : le dépôt pve-enterprise est encore actif sans abonnement. Solution : commentez les lignes `deb https://enterprise.proxmox.com` dans `/etc/apt/sources.list.d/pve-enterprise.list` et ajoutez le dépôt no-subscription comme à l’étape 5.

### Erreur 3 : VM Windows en boucle de redémarrage après import ESXi

Cause : drivers VirtIO absents dans Windows. Solution : démarrez d’abord la VM avec un contrôleur SATA (`qm set 110 --sata0 local-zfs:vm-110-disk-0 --boot order=sata0`), installez VirtIO via le `virtio-win.iso` attaché en CD-ROM, puis rebasculez sur le contrôleur VirtIO SCSI.

### Erreur 4 : « Quorum failed » après reboot d’un nœud

Cause : sur un cluster à 2 nœuds, la perte d’un nœud fait tomber le quorum (1 vote sur 2). Solution : ajoutez un troisième nœud ou un QDevice externe avec `pvecm qdevice setup <ip-qdevice>`. Le QDevice peut être un Raspberry Pi 5 sous Debian 13 avec le paquet `corosync-qnetd`.

### Erreur 5 : « device-mapper: create ioctl on… failed: Device or resource busy »

Cause : un disque a une signature LVM ou ZFS résiduelle. Solution : avant `pveceph osd create` ou `zpool create`, lancez `wipefs -a /dev/nvmeXn1` et `sgdisk --zap-all /dev/nvmeXn1`.

### Erreur 6 : performances réseau VM en dessous de 1 Gbit/s

Cause : modèle de carte VirtIO non utilisé, ou multiqueue désactivé. Solution : `qm set 100 --net0 virtio,bridge=vmbr0,queues=4` ; vérifiez côté invité avec `ethtool -L eth0 combined 4`.

### Erreur 7 : sauvegarde PBS très lente (< 50 Mo/s)

Cause : disque PBS en HDD ou checksum SHA-256 calculé sur CPU lent. Solution : utilisez un SSD pour le datastore PBS, vérifiez avec `proxmox-backup-client benchmark` que le débit chunk-write atteint au moins 200 Mo/s.

### Erreur 8 : interface web inaccessible après upgrade

Cause : service `pveproxy` ou `pvedaemon` non redémarré. Solution : SSH dans le nœud, lancez `systemctl restart pveproxy pvedaemon pvestatd`. Si l’erreur persiste, regardez `journalctl -u pveproxy -n 100` pour identifier un éventuel certificat expiré ou conflit de port 8006.

## Cas d’usage en France et en Europe en 2026

Plusieurs administrations françaises ont publiquement annoncé leur migration vers Proxmox dans le cadre de la stratégie « cloud de confiance » et de la trajectoire SecNumCloud. Le ministère de l’Économie et des Finances, à travers Bercy, et plusieurs collectivités territoriales évaluent l’hyperviseur viennois pour réduire leur dépendance à VMware. À l’échelle européenne, la cour des comptes des Pays-Bas et plusieurs Länder allemands ont migré une partie de leur parc en 2024-2025. La société allemande Siltronic et l’éditeur danois Tradeshift ont publiquement témoigné de leur bascule en conférence Proxmox Day Vienna en juin 2025.

Pour les hébergeurs européens souverains, Proxmox est devenu une brique standard. OVHcloud, Scaleway et Hetzner proposent depuis 2024-2025 des serveurs dédiés livrés avec Proxmox VE pré-installé via leur portail self-service, parfois en moins de 2 minutes pour un Hetzner AX52 à 50 € HT mensuels. Ces déploiements rapides facilitent les architectures hybrides où Proxmox sert d’hyperviseur de bureau (lab) et de tampon avant migration vers du cloud public AWS ou Azure.

## Questions fréquentes (FAQ)

### Quelle différence entre Proxmox VE et Proxmox Backup Server ?

Proxmox VE est l’hyperviseur qui exécute les VM et conteneurs ; Proxmox Backup Server (PBS) est le serveur de sauvegarde déduplicquée et chiffrée dédié, qui s’utilise comme cible des jobs vzdump. Les deux produits sont open source et gratuits, indépendants, et peuvent fonctionner sur des machines distinctes.

### Peut-on installer Proxmox VE 9.1 sur un mini-PC ou un NUC ?

Oui, à condition que le CPU supporte la virtualisation matérielle. Les Intel NUC 12, 13, 14 et les ASUS NUC 15 fonctionnent parfaitement, de même que les Beelink, Minisforum et GMKtec à base de Ryzen. Le seul écueil concerne les NIC Realtek 2.5GbE qui demandent parfois un driver hors arbre Linux : l’extension *r8125-dkms* du dépôt Proxmox community resout généralement le problème.

### Proxmox VE 9 est-il compatible avec Windows 11 et Windows Server 2025 ?

Oui. Proxmox VE 9 supporte les machines virtuelles Windows 11 (TPM 2.0 émulé via le state qcow2 introduit en 9.1, Secure Boot, Q35 + UEFI) et Windows Server 2025. Microsoft maintient l’éligibilité au support tant que les drivers VirtIO signés sont installés, condition documentée dans la base SVVP (Server Virtualization Validation Program).

### Quelle est la durée de support de Proxmox VE 8 ?

Proxmox VE 8 reçoit des mises à jour de sécurité jusqu’au 31 août 2026. Au-delà, seuls les abonnés Premium pourraient obtenir des correctifs ponctuels. Migrer vers PVE 9 avant cette date est fortement recommandé : la procédure officielle `pve8to9` détecte automatiquement les incompatibilités.

### Faut-il choisir Proxmox ou XCP-ng comme alternative à VMware ?

Les deux sont open source, basés sur des hyperviseurs Linux et matures. Proxmox VE l’emporte sur la richesse des fonctionnalités intégrées (Ceph natif, conteneurs LXC + OCI, PBS dédié) et sur la base utilisateurs européenne. XCP-ng (basé sur XenServer) plaît aux migrants Citrix et offre Xen Orchestra comme console centralisée. Pour un greenfield 2026, Proxmox concentre l’essentiel de la dynamique.

### Combien de VM peut héberger un nœud Proxmox ?

La limite logicielle est très haute (plus de 10 000 VMID par cluster). En pratique, c’est le matériel qui limite : un Xeon Gold 6526Y bi-socket avec 512 Go de RAM ECC peut faire tourner 80 à 120 VM Linux légères ou 30 à 50 VM Windows Server. Pour un VDI bureautique, les ratios CPU 1:6 à 1:8 et RAM 1:1,5 sont raisonnables.

### Le HA Manager fonctionne-t-il sans Ceph ?

Oui, à condition d’utiliser un stockage partagé : NFS, iSCSI, CIFS ou ZFS over iSCSI. Le seul critère est que la VM puisse être démarrée sur n’importe quel nœud du cluster. Si vos disques sont sur un stockage local exclusif, le HA Manager refuse l’ajout de la VM.

### Comment automatiser entièrement le provisioning de mon cluster ?

Combinez trois outils : **Terraform** avec le provider `bpg/proxmox` pour décrire l’infrastructure (VM, ressources, réseaux) ; **cloud-init** dans des templates pour configurer chaque VM au premier démarrage (utilisateurs, SSH keys, paquets) ; **Ansible** pour les opérations idempotentes ultérieures (mises à jour, déploiement applicatif). Cette stack reproduit ce que VMware appelle vRealize Automation, sans licence supplémentaire.

### Related Coverage

## Conclusion : un hyperviseur ouvert, mature et compétitif

Proxmox VE 9.1 incarne en 2026 la promesse d’un hyperviseur de classe entreprise sans verrou propriétaire ni explosion tarifaire. Pour un coût plancher de zéro euro et un coût plafond de 1 100 € HT par CPU et par an, vous obtenez un produit qui supporte la virtualisation KVM, les conteneurs LXC et OCI, le stockage Ceph hyper-convergé, la sauvegarde déduplicquée via PBS 4.1, le HA Manager, la migration à chaud, l’intégration Active Directory, le firewall et l’API REST avec providers Terraform et Ansible. Aucune licence vSphere n’offre cet éventail dans cette gamme de prix en 2026.

