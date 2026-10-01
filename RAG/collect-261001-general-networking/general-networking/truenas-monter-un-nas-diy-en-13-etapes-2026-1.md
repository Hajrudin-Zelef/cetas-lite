---
id: collect-261001-general-networking/general-networking/truenas-monter-un-nas-diy-en-13-etapes-2026-1
title: "Identifier la clé USB (attention à bien cibler le bon disque)"
domain: general-networking
role: reference
task: reference
actors: ["AMD", "Intel", "Nvidia"]
dates: []
keywords: ["amd", "cost", "ethernet", "gpu", "intel", "mai", "nvidia", "open source"]
source: docs/RAG/collect-261001-general-networking/truenas-monter-un-nas-diy-en-13-etapes-2026.md
source_anchor: ""
source_lines: [1, 37]
sha256: 57b4dd36e8402c06ab9f602726e87762342ad12d0ec455c210af656840b84005
---

# Identifier la clé USB (attention à bien cibler le bon disque)

Les prix des disques durs ont grimpé de 46 % en quatre mois selon Club386, une carte QNAP 4 baies diskless dépasse maintenant les 695 € chez les revendeurs allemands, et les inquiétudes sur la confidentialité des données hébergées chez des tiers ne faiblissent pas. Résultat : de plus en plus de particuliers et de petites structures en France et en Europe se tournent vers le NAS DIY (fait maison) plutôt que vers un boîtier Synology ou QNAP tout prêt. Le logiciel qui domine cette scène s’appelle TrueNAS, et son nom a changé en 2026 : ce qu’on appelait TrueNAS SCALE est devenu **TrueNAS Community Edition**. Signe de cette popularité grandissante, la branche TrueNAS 25.10 a franchi la barre des 100 000 déploiements actifs dès février 2026, selon les chiffres communiqués par iXsystems. Ce guide vous montre, en 13 étapes concrètes, comment transformer un PC ou un mini-ITX en serveur de stockage ZFS complet avec TrueNAS Community Edition 25.10.6 « Goldeye », la version stable au 22 août 2026.

On y couvre le choix du matériel, l’installation, la création du pool ZFS, les partages réseau, les snapshots, les applications Docker comme Plex ou Nextcloud, la sécurisation de l’accès et un vrai plan de sauvegarde. On termine avec les pièges classiques, un dépannage en 8 points et un comparatif de coûts face à Synology et QNAP.

## Qu’est-ce que TrueNAS Community Edition en 2026 ?

Avant de commencer, il faut clarifier un point qui perd beaucoup de monde en 2026 : TrueNAS SCALE n’existe plus sous ce nom. Avec la sortie de TrueNAS 25.04 « Fangtooth » en avril 2025, iXsystems a fusionné ses gammes et rebaptisé SCALE en **TrueNAS Community Edition (CE)**. Cette branche a ensuite reçu une première mise à jour de maintenance, la 25.04.1, dès mai 2025, tandis que l’ancienne branche 24.10 « Electric Eel » recevait sa dernière mouture terminale, la 24.10.2.4, en août 2025, avant d’être définitivement abandonnée au profit de la nouvelle génération. La page produit officielle le confirme noir sur blanc : « TrueNAS SCALE is now officially known as TrueNAS Community Edition. » L’ancien nom SCALE ne survit plus que dans un contexte historique, y compris sur les forums où d’anciens utilisateurs demandent encore pourquoi leur installation « SCALE » s’affiche désormais comme « Community Edition ».

Concrètement, TrueNAS Community Edition reste un système d’exploitation basé sur Debian Linux, gratuit et open source, conçu pour transformer n’importe quel PC x86-64 en serveur de stockage ZFS. Il coexiste avec TrueNAS Enterprise, la version payante vendue avec le support commercial d’iXsystems et destinée aux entreprises. Les deux partagent le même socle logiciel : Community Edition sert de base gratuite pour les particuliers, les homelabs et les petites structures, tandis qu’Enterprise ajoute le support et des fonctionnalités de scalabilité.

La branche majeure 25.10 « Goldeye » est passée en disponibilité générale dès octobre 2025, avant de recevoir une mise à jour 25.10.1 en décembre 2025, puis une 25.10.2 en février 2026 qui a corrigé plus de 100 bugs selon iXsystems. La version stable au moment de la rédaction est **TrueNAS 25.10.6 « Goldeye »**, sortie le 12 août 2026. Cette mise à jour de maintenance corrige des failles de sécurité dans le noyau Linux et le pilote NVIDIA GPU, selon les notes de version officielles. Une branche 26.0.0 existe déjà en version bêta (BETA.3 au 20 août 2026), mais elle n’est pas destinée à la production : pour ce tutoriel, on utilise exclusivement la branche 25.10 stable.

## Pourquoi construire un NAS DIY plutôt qu’acheter un Synology ou QNAP ?

La réponse courte : le contrôle et le rapport performance/prix. Un NAS DIY sous TrueNAS Community Edition tourne sur du matériel PC standard, ce qui veut dire que vous choisissez vous-même le processeur, la quantité de RAM, le nombre de baies et le type de réseau (2,5 GbE ou 10 GbE) sans être limité aux specs figées d’un boîtier fermé. Sur les forums TrueNAS, le blogueur homelab Brian Moses a documenté en détail sa propre construction dans un billet largement cité dans la communauté : « With drives, my out of pocket cost was a little over $1,750 and without drives the price comes in at right around $617 », précise-t-il à propos de sa configuration compacte à 5 baies avec réseau 10GbE.

À titre de comparaison, un NAS Synology 4 baies diskless démarre autour de 659 € chez les revendeurs allemands (idealo.de, août 2026), et un QNAP TS-464-8G équivalent se négocie à partir de 695,26 €, disques non inclus dans les deux cas. Le DIY ne gagne pas systématiquement sur le prix pur (une carte mère mini-ITX + boîtier + alimentation peut coûter cher aussi), mais il gagne sur la flexibilité : upgrade de RAM à volonté, choix libre des disques, accès root complet, et surtout la possibilité de faire tourner des containers Docker sans les limitations d’un NAS propriétaire.

Il y a aussi un argument de résilience aux prix. Selon Club386, les prix des disques durs ont grimpé de 46 % en seulement quatre mois début 2026, un Toshiba MG11ACA 24 To passant de 386,09 € à 602,00 €. Avec un NAS DIY, vous pouvez ajouter des disques progressivement, au fil des baisses de prix, plutôt que d’acheter un pack complet imposé par le constructeur.

Pour les foyers et les indépendants en France, il existe un troisième argument, plus discret mais tout aussi réel : la question de la localisation des données. Un NAS DIY reste physiquement chez vous, sous votre propre toit, ce qui évite de dépendre des conditions générales d’un cloud tiers pour vos photos de famille, vos documents administratifs ou vos sauvegardes professionnelles. Ce n’est pas un argument juridique en soi, mais c’est une garantie pratique que personne d’autre ne peut modifier unilatéralement les conditions d’accès à vos fichiers du jour au lendemain.

Enfin, un NAS DIY sous TrueNAS s’adapte à des usages que la plupart des boîtiers grand public ne couvrent pas nativement : héberger un serveur Plex qui transcode plusieurs flux 4K en simultané, faire tourner un petit cluster de VM pour tester des configurations, ou servir de cible de sauvegarde pour tout un parc de PC via Time Machine (macOS) ou Veeam Community Edition (Windows). Ces usages restent possibles sur un Synology ou un QNAP haut de gamme, mais à un coût d’achat souvent double, pour une puissance CPU comparable à une carte mini-ITX à 150 €.

## Prérequis : matériel, logiciel et budget avant de commencer

Avant de sortir le tournevis, voici ce qu’il vous faut réellement. Cette liste correspond aux exigences officielles du guide matériel de TrueNAS, complétées par les recommandations pratiques de la communauté homelab.

- Un processeur x86-64 (Intel ou AMD), 2 cœurs minimum selon la doc officielle, mais 4 cœurs recommandés si vous comptez lancer des applications Docker
- 8 Go de RAM minimum pour une installation basique jusqu’à 8 disques (+1 Go par disque supplémentaire selon iXsystems), 16 Go pour un usage réel, 32 Go ou plus si vous prévoyez des VM ou plusieurs containers
- Un support d’amorçage dédié : SSD ou clé USB de 16 à 20 Go minimum (un SSD SATA ou NVMe de 120 Go est largement suffisant et plus fiable dans la durée)
- Au moins deux disques identiques pour créer un premier pool de stockage ZFS (HDD ou SSD selon l’usage)
- Une carte mère avec suffisamment de ports SATA ou un contrôleur HBA en mode passthrough (IT mode) pour plus de baies
- Une clé USB de 8 Go minimum pour créer le support d’installation
- Un deuxième ordinateur pour télécharger l’image ISO et accéder à l’interface web de TrueNAS pendant la configuration
- Une connexion Ethernet filaire (le Wi-Fi n’est pas recommandé ni officiellement supporté pour le stockage principal)

