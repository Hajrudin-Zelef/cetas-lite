---
id: collect-261001-general-networking/general-networking/proxmox-9-1-cluster-ha-99-9-uptime-en-13-etapes-2026-6
title: "Télécharger l'ISO et la somme SHA-256"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "benchmarks", "datacenter", "license", "mai", "open source"]
source: docs/RAG/collect-261001-general-networking/proxmox-9-1-cluster-ha-99-9-uptime-en-13-etapes-2026.md
source_anchor: ""
source_lines: [431, 461]
sha256: 6c5c71bb60cd7f5d606b25cfc5d676680c023d65c9ea7357f53073f2f647f45f
---

# Télécharger l'ISO et la somme SHA-256

Officiellement, ce n’est plus tout à fait non : Proxmox Server Solutions GmbH propose désormais sur sa page de téléchargement une image ISO ARM64, la 9.2-1 (1,50 Go), mise à jour en août 2026, mais toujours en statut expérimental et loin de la maturité de la branche x86_64. En attendant sa stabilisation, la communauté maintient depuis plus longtemps un fork nommé pimox qui fonctionne sur Raspberry Pi 5 et Orange Pi 5+. Les performances sont correctes pour des containers LXC légers, mais le passthrough PCI est inopérant et Ceph mono-disque n’a pas de sens. C’est utile uniquement comme nœud QDevice pour donner un quorum à un cluster x86 à deux nœuds.

### Proxmox supporte-t-il Windows 11 et Windows Server 2025 ?

Oui, parfaitement. Windows 11 et Windows Server 2025 demandent UEFI (OVMF), Secure Boot et un TPM virtuel – tous trois pris en charge nativement par Proxmox 9.1. À la création de la VM, ajoutez explicitement `--bios ovmf --machine q35 --tpmstate0 tank-vms:1,version=v2.0`. Installez ensuite les drivers virtio (virtio-win-0.1.262 ou supérieur) pour le réseau et le SCSI. Avec ces réglages, les benchmarks Windows 11 tournent à environ 96 % de la performance native, et l’activation de l’OS fonctionne sans accroc.

### Comment migrer une VM VMware vers Proxmox ?

Depuis Proxmox 8.2, un outil officiel d’import existe via Datacenter → Storage → Add → ESXi. Renseignez l’IP de l’hôte ESXi, vos identifiants vCenter, et Proxmox liste les VM accessibles. Sélectionnez une VM, choisissez le stockage cible, et la conversion VMDK → qcow2 (ou ZFS) se fait en streaming via SSH. Pour 100 Gio de disque sur un lien 1 GbE, comptez 15 à 20 minutes. Ne migrez jamais une VM allumée : exportez-la d’abord à froid pour garantir la cohérence. Pensez ensuite à désinstaller VMware Tools et à installer qemu-guest-agent.

### Quelle est la différence entre LVM-Thin et ZFS pour les disques de VM ?

LVM-Thin est plus simple, plus rapide en écritures séquentielles et consomme moins de RAM. C’est le choix par défaut quand on installe Proxmox sur un seul SSD sans miroir. Depuis Proxmox VE 9.0, sorti en août 2025, le LVM classique (thick) a lui aussi gagné le support des snapshots sur stockage partagé, comblant un manque historique pour les clusters iSCSI qui ne pouvaient pas jusque-là bénéficier des snapshots réservés à LVM-Thin. ZFS apporte de son côté snapshots instantanés, compression LZ4 ou Zstd, checksums end-to-end et tolérance aux erreurs silencieuses, mais consomme environ 1 Gio de RAM par To de pool en ARC. Pour une production avec données critiques, ZFS gagne haut la main grâce à la résilience. Pour un homelab mono-disque, LVM-Thin suffit.

### Comment sauvegarder Proxmox lui-même (pas les VM) ?

L’hôte Proxmox lui-même doit être traité comme un serveur Debian classique. Sauvegardez régulièrement /etc/pve (configuration cluster), /etc/network/interfaces, /root et /var/lib/pve-firewall. Un simple cron quotidien qui tar.gz ces dossiers vers PBS suffit. En cas de panne matérielle, une réinstallation fraîche puis la restauration de /etc/pve via pvecm add suffit à reformer le cluster. Les VM elles-mêmes vivent sur le stockage partagé et restent accessibles dès qu’un nœud est restauré.

### Proxmox Backup Server vs Veeam vs Bacula ?

PBS est conçu spécifiquement pour Proxmox et offre la déduplication par chunks 4 Mio la plus efficace du marché sur des VM Linux quasi-identiques. Veeam reste roi pour les environnements mixtes VMware + Hyper-V + physique mais coûte 30 à 50 euros par VM et par an. Bacula est gratuit et flexible mais demande une configuration complexe en fichiers texte. Pour un cluster 100 % Proxmox, PBS est imbattable sur le rapport fonctionnalités/coût. Pour un environnement hétérogène avec migration progressive, Veeam Universal License couvre les deux.

### Articles complémentaires

## Conclusion : Proxmox VE comme socle européen post-VMware

Proxmox VE 9.1 n’est plus un projet d’amateur : c’est désormais la pierre angulaire de la virtualisation open source en Europe, adoptée par des hébergeurs comme OVHcloud pour ses plateformes managées, par Tranquil IT pour ses offres souveraines, par la majorité des nouvelles écoles d’ingénieurs françaises pour leur cluster pédagogique, et par des milliers de PME en migration depuis VMware. La combinaison Debian 13 + kernel 6.14 + ZFS 2.3.3 + Ceph Squid 19.2.3 + QEMU 10.0.2 forme une stack mature, maintenue, et entièrement gratuite à fonctionnalités égales avec vSphere Foundation – et l’écosystème continue de s’étoffer, entre la sortie de Proxmox Mail Gateway 9.1-1 saluée par la presse spécialisée courant mai-juin 2026 et l’arrivée de VE 9.2 – basée sur Debian 13.5 et le noyau Linux 7.0, publiée en mai 2026 – pour prendre le relais de ce guide, tandis que la branche 8.x, plafonnée à la 8.4 d’avril 2025, sort officiellement de support le 31 août 2026 selon endoflife.date.

Pour aller plus loin, explorez la nouvelle console Proxmox Datacenter Manager pour piloter plusieurs clusters depuis une interface unique, et la SDN EVPN pour offrir du multi-tenant à vos clients. Si vous opérez plus de cinq clusters, l’abonnement Community Proxmox à 110 euros par CPU et par an reste l’investissement le plus rentable du secteur. Les forums officiels et la documentation wiki couvrent désormais 95 % des situations en français ou en anglais.

*Article publié le 17 avril 2026 sur tech-insider.org. Données vérifiées le jour de publication sur la base des notes de version Proxmox VE 9.1-1 (19 novembre 2025) et Proxmox Backup Server 4.2-1 (29 avril 2026).*
