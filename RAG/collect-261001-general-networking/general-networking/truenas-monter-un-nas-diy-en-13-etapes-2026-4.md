---
id: collect-261001-general-networking/general-networking/truenas-monter-un-nas-diy-en-13-etapes-2026-4
title: "Identifier la clé USB (attention à bien cibler le bon disque)"
domain: general-networking
role: reference
task: reference
actors: ["Google", "Intel"]
dates: []
keywords: ["ethernet", "intel"]
source: docs/RAG/collect-261001-general-networking/truenas-monter-un-nas-diy-en-13-etapes-2026.md
source_anchor: ""
source_lines: [167, 211]
sha256: 9a2f2dd256a8dd62105dad6be4c2fc8be4c99a3b4271d9dc9f962decfdef6116
---

# Identifier la clé USB (attention à bien cibler le bon disque)

Enfin, ne jamais exposer directement l’interface web de TrueNAS sur Internet en redirigeant un port depuis votre box. Si vous avez besoin d’un accès distant, passez par un VPN vers votre réseau local (WireGuard, disponible directement via une app TrueNAS, ou un routeur compatible), ou par un reverse proxy avec authentification renforcée. Vérifiez également que le pare-feu de votre routeur bloque bien tout trafic entrant non sollicité vers l’IP du NAS.

## Étape 11 : Sauvegardes 3-2-1 et test de restauration

Un pool RAIDZ2 protège contre la panne de disque, pas contre un incendie, un vol, une erreur humaine ou un ransomware qui chiffre vos fichiers avant que le snapshot suivant ne se déclenche. La règle 3-2-1 reste la référence : 3 copies de vos données, sur 2 supports différents, dont 1 copie hors site. Concrètement pour un NAS TrueNAS : les données originales sur le pool ZFS, une réplication vers un second NAS ou un disque externe, et une copie chiffrée envoyée vers un stockage cloud compatible S3 (Backblaze B2, Wasabi, ou équivalent).

```
# Exemple de synchronisation manuelle vers un stockage distant compatible S3 avec rclone
rclone sync /mnt/tank/documents remote-backup:mon-bucket/documents \
  --transfers 4 --checksum --log-file=/var/log/rclone-backup.log
# Vérifier la dernière exécution planifiée dans les logs
tail -n 50 /var/log/rclone-backup.log
```
Le point que presque tout le monde néglige : personne ne teste jamais la restauration avant d’en avoir réellement besoin. Planifiez au moins une fois par trimestre une restauration test d’un snapshot ou d’une sauvegarde distante vers un dossier temporaire, pour vérifier que les fichiers restaurés sont bien intacts et lisibles. Une sauvegarde jamais testée n’est qu’une hypothèse.

## Étape 12 : Cas d’usage concrets pour exploiter votre NAS

Une fois l’infrastructure de base en place (pool, partages, snapshots, sauvegardes), l’intérêt réel d’un NAS DIY se révèle dans les usages qu’il permet d’empiler sur la même machine. Voici quatre scénarios concrets, testés par la communauté homelab, qui montrent pourquoi TrueNAS dépasse largement le simple stockage de fichiers.

- **Serveur multimédia centralisé.** Installez Plex ou Jellyfin depuis le catalogue d’applications, pointez la bibliothèque vers un dataset dédié, et diffusez votre collection de films et séries vers vos téléviseurs, consoles et smartphones sans dépendre d’un abonnement de streaming. Sur un Intel N150, le transcodage logiciel gère confortablement un flux 1080p, mais un flux 4K HDR demande plutôt un CPU avec un iGPU capable de décodage matériel (Quick Sync), à vérifier avant l’achat de la carte mère.
- **Sauvegarde centralisée du parc familial ou professionnel.** Configurez Time Machine sur les Mac (TrueNAS expose nativement un partage SMB compatible Time Machine) et un outil comme Veeam Community Edition ou Duplicati sur les PC Windows, tous pointant vers des datasets séparés sur le même pool. Un seul NAS peut ainsi absorber les sauvegardes de toute une maisonnée ou d’une petite équipe.
- **Photothèque privée avec reconnaissance IA.** Immich, disponible dans le catalogue TrueNAS Apps, reproduit l’essentiel des fonctionnalités de Google Photos (tri automatique, reconnaissance de visages, carte des lieux) mais garde toutes les photos sur votre pool ZFS local plutôt que sur un cloud tiers.
- **Bac à sable pour l’apprentissage DevOps.** TrueNAS Community Edition supporte les machines virtuelles en plus des containers Docker. C’est un terrain d’entraînement réaliste pour tester des déploiements Kubernetes, des bases de données ou des pipelines CI/CD sans payer d’instance cloud à l’heure.

## 5 erreurs courantes à éviter lors d’un montage NAS DIY

- **Mélanger des disques de tailles ou d’âges différents dans un même vdev.** ZFS aligne la capacité utile sur le plus petit disque du vdev, et des disques achetés en même temps ont statistiquement plus de risques de tomber en panne au même moment (défaut de lot).
- **Utiliser le disque de démarrage comme espace de stockage de données.** C’est une erreur fréquente chez les débutants qui veulent économiser un port SATA : en cas de panne du disque système, vous perdez la configuration ET les données au lieu de perdre uniquement la configuration.
- **Sous-dimensionner la RAM en pensant que 8 Go officiels suffisent à un usage réel.** C’est un minimum de démarrage, pas une recommandation de confort, surtout dès que vous ajoutez des apps Docker.
- **Négliger le test S.M.A.R.T. long des disques neufs avant mise en production.** Un disque neuf défectueux se révèle souvent dans les 48 premières heures d’utilisation intensive (courbe de mortalité infantile), autant le savoir avant d’y stocker des données importantes.
- **Confondre snapshots et sauvegardes.** Un snapshot protège contre l’erreur humaine sur le même pool, pas contre la perte du NAS entier. Sans réplication ou copie hors site, vous n’avez toujours qu’une seule copie physique de vos données.

## Dépannage : 8 problèmes fréquents et leurs solutions

1. **L’interface web reste inaccessible après l’installation.** Vérifiez que le câble Ethernet est bien branché avant le démarrage (TrueNAS ne détecte pas toujours une interface branchée à chaud) et confirmez l’adresse IP via la console physique du serveur.
2. **Le pool ZFS affiche l’état DEGRADED.** Un disque a été marqué défaillant. Identifiez-le avec `zpool status -v`, remplacez-le physiquement, puis lancez `zpool replace tank ancien-disque nouveau-disque` pour déclencher le resilver.
3. **Les transferts SMB sont anormalement lents.** Testez d’abord la bande passante réseau brute avec iperf3 entre le client et le NAS avant de blâmer TrueNAS : un câble Cat5e ou un switch limité à 1 GbE plafonne tout, quelle que soit la config logicielle.
4. **Une application Docker refuse de démarrer après une mise à jour.** Consultez les logs du container depuis Apps > sélectionner l’app > Logs, et vérifiez en priorité les chemins de volumes : une mise à jour de TrueNAS peut parfois modifier les points de montage attendus.
5. **Le pool ne monte plus après un redémarrage soudain (coupure de courant).** ZFS est journalisé (copy-on-write) et résiste bien aux coupures brutales, mais lancez un `zpool scrub tank` dès que possible pour vérifier l’intégrité et corriger d’éventuelles erreurs de checksum.
6. **Le certificat ACME/Let’s Encrypt ne se renouvelle pas automatiquement.** Vérifiez que le port 80 (validation HTTP-01) reste bien accessible depuis Internet vers votre NAS au moment du renouvellement, ou basculez vers une validation DNS-01 si votre registrar le supporte.
7. **Le NAS surchauffe ou les ventilateurs tournent à plein régime en permanence.** Vérifiez le mapping des courbes de ventilation dans le BIOS et la propreté des filtres à poussière : un boîtier NAS mal ventilé fait grimper la température des disques au-delà de 40°C, ce qui réduit leur durée de vie.
8. **Impossible de se connecter en SSH après activation du firewall intégré.** Vérifiez dans System Settings > Services que le service SSH est bien démarré ET que la règle de firewall autorise le port 22 depuis votre plage d’IP locale, les deux réglages sont indépendants.

## Astuces avancées pour optimiser les performances de votre NAS

