---
id: collect-261001-general-networking/general-networking/nas-synology-configurer-dsm-en-12-etapes-2026-5
title: "Exemple de planification de tâche Hyper Backup (via l'interface DSM)"
domain: general-networking
role: reference
task: reference
actors: ["AMD", "Intel"]
dates: []
keywords: ["agent", "amd", "intel"]
source: docs/RAG/collect-261001-general-networking/nas-synology-configurer-dsm-en-12-etapes-2026.md
source_anchor: ""
source_lines: [226, 272]
sha256: f72f6516ea6b5ae6324ce9f8abdb7f1e5efa6a69fb3f396956a2f56ddeb1abfa
---

# Exemple de planification de tâche Hyper Backup (via l'interface DSM)

**Impossible de se connecter après une mise à jour DSM.** Videz le cache du navigateur et essayez en navigation privée avant de suspecter un vrai problème système ; DSM change parfois de port ou de certificat après une mise à jour majeure, ce qui déclenche des avertissements de sécurité du navigateur qu’il faut accepter manuellement une fois.

## Comparatif rapide : quel modèle Synology choisir en 2026

| Modèle | Baies | CPU / RAM | Prix France (boîtier seul) | 
|---|---|---|---|
| DS224+ | 2 | Intel Celeron J4125, 2 Go DDR4 (max 6 Go) | ~400-430 € | 
| DS423+ | 4 + 2 M.2 NVMe (cache) | Intel, 2 Go DDR4 (max 6 Go) | ~430-600 € | 
| DS923+ | 4 (extensible à 9) + 2 M.2 NVMe | AMD Ryzen R1600, 4 Go DDR4 ECC (max 32 Go) | ~600-750 € | 

Pour un premier NAS familial centré sur la sauvegarde photo et documents, le DS224+ suffit largement. Le DS423+ apporte de la marge pour du multimédia et davantage de baies. Le DS923+ vise plutôt les usages plus intensifs : machines virtuelles légères, bases de données, ou charge de travail qui bénéficie de la RAM ECC et du CPU plus puissant.

### Related Coverage

## Questions fréquentes

**Quelle est la dernière version de DSM disponible en août 2026 ?**

DSM 7.4.1-90080, publiée le 23 juillet 2026, est la version générale la plus récente ; elle fait suite à DSM 7.4, lancée le 16 juin 2026 avec l’assistant IA DSM Agent. Pour la maintenance longue durée, Synology a fait basculer le support LTS sur DSM 7.3 depuis octobre 2025 (maintenance jusqu’en octobre 2027), même si des correctifs comme DSM 7.2.2-72806 Update 9 (30 juin 2026) restent disponibles pour les modèles plus anciens.

**Faut-il choisir Btrfs ou ext4 pour un premier NAS Synology ?**

Btrfs, dans la grande majorité des cas, grâce aux snapshots, aux checksums qui détectent la corruption silencieuse et à un historique de versions plus économe en espace avec Synology Drive. ext4 reste pertinent seulement pour du matériel très limité en CPU ou des besoins spécifiques de compatibilité iSCSI.

**Un NAS Synology remplace-t-il une vraie sauvegarde ?**

Non. Un NAS avec RAID protège contre la panne d’un disque, pas contre un vol, un incendie ou un ransomware. Il faut ajouter Hyper Backup vers une destination externe (autre NAS, disque USB ou C2 Storage) pour obtenir une vraie protection selon la méthode 3-2-1-1-0.

**Combien coûte Synology C2 Storage en 2026 ?**

Les paliers Basic en Europe démarrent autour de 9,99 € par an pour 100 Go, 24,99 € par an pour 300 Go et 59,99 € par an pour 1 To, hors TVA. C2 Backup Business se facture environ 11,99 € par To et par mois pour les structures professionnelles.

**Peut-on utiliser des disques durs de PC classiques dans un NAS Synology ?**

Techniquement oui, mais ce n’est pas recommandé pour un usage continu 24/7. Les disques dédiés NAS sont conçus pour tolérer davantage de vibrations et fonctionner en permanence, avec des taux de panne nettement plus bas sur la durée que des disques de bureau grand public.

**QuickConnect est-il sûr pour accéder à son NAS depuis l’extérieur ?**

Avec 2FA activée et Auto Block configuré, QuickConnect offre un niveau de sécurité correct pour un usage domestique. Pour un usage professionnel ou des données sensibles, un accès VPN vers le réseau local reste préférable à une exposition directe de DSM sur internet.

**Quel modèle Synology choisir pour un premier NAS en 2026 ?**

Le DS224+ (2 baies, environ 400-430 € en France) couvre largement les besoins d’un foyer standard pour la sauvegarde photo et documents. Le DS423+ ou le DS923+ apportent plus de baies et de puissance pour du multimédia intensif ou des charges de travail plus lourdes.

**Comment savoir si mon NAS Synology doit être mis à jour pour des raisons de sécurité ?**

Consultez régulièrement la page officielle des avis de sécurité Synology et activez les mises à jour automatiques dans Panneau de configuration > Mise à jour et restauration. Six avis de sécurité au moins ont été publiés entre mars et août 2026, dont un jugé critique en mars concernant telnetd.
