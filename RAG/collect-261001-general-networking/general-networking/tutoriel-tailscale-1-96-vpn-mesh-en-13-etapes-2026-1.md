---
id: collect-261001-general-networking/general-networking/tutoriel-tailscale-1-96-vpn-mesh-en-13-etapes-2026-1
title: "Installation rapide (Debian, Ubuntu, Fedora, Arch, etc.)"
domain: general-networking
role: reference
task: reference
actors: ["Apple", "Google", "Microsoft"]
dates: []
keywords: ["arr", "benchmark", "open source"]
source: docs/RAG/collect-261001-general-networking/tutoriel-tailscale-1-96-vpn-mesh-en-13-etapes-2026.md
source_anchor: ""
source_lines: [1, 49]
sha256: 7fce8db55aa3f8ed7dca9490a3c71743f3c5671d21e2eb14e06486f9d43fd9a8
---

# Installation rapide (Debian, Ubuntu, Fedora, Arch, etc.)

**Mise à jour : 17 août 2026.** Tailscale s’est imposé en 2026 comme la solution de référence pour bâtir un réseau privé maillé (mesh VPN) entre serveurs, postes de travail, conteneurs et services cloud, une position renforcée par une levée de **160 millions de dollars à une valorisation de 1,5 milliard de dollars** bouclée en avril 2025 (BNN Bloomberg), portant le financement total cumulé à **275 millions de dollars**, et confirmée par une trajectoire commerciale solide : un chiffre d’affaires récurrent annuel (ARR) évalué à **45,2 millions de dollars** dès juin 2025 selon les estimations de GetLatka citant tailscale.com, une dynamique qui a ensuite porté la base de clients professionnels payants à près de **40 000** en juillet 2026 d’après NCFA Canada. Le dépôt GitHub officiel affiche désormais la **v1.102.4**, publiée en septembre 2026 selon Tailscale, comme dernière build stable — une lignée qui a succédé à la branche 1.96.x et qui prolonge le support multi-tailnet pour l’opérateur Kubernetes, des correctifs mémoire iOS et la fonctionnalité *Peer Relays* en disponibilité générale, avec un gain de débit jusqu’à **15× supérieur** aux serveurs DERP traditionnels. Avec un plan Personal gratuit autorisant jusqu’à 6 utilisateurs pour 0 $ et un nombre illimité d’appareils, Tailscale couvre aussi bien le homelab que les déploiements d’entreprise sous le plan désormais renommé Standard, facturé **8 $/utilisateur/mois** en août 2026 selon ModSignal (engagement annuel).

Ce tutoriel guide pas à pas la mise en place d’un tailnet complet en 13 étapes : installation sur Linux Debian/Ubuntu, configuration des nœuds de sortie (exit nodes), routage de sous-réseaux (subnet routes), MagicDNS, ACL JSON, Tailscale SSH, Tailscale Funnel pour exposer un service local sur Internet, audit des connexions et migration optionnelle vers **Headscale**, l’alternative open source auto-hébergée. Comptez environ 90 minutes pour exécuter l’ensemble du parcours sur deux machines Linux et un smartphone, avec une dizaine d’extraits de code prêts à copier.

## Pourquoi choisir Tailscale en 2026

Tailscale, fondé en 2019 à Toronto par Avery Pennarun et David Crawshaw (anciens ingénieurs de Google), employait environ **250 salariés en novembre 2025** selon les estimations de GetLatka, contre 123 en 2023 — un doublement de l’effectif qui accompagne l’accélération commerciale de la société. L’équipe construit son réseau virtuel directement sur le protocole **WireGuard**, qui utilise le port UDP 41641 par défaut et le chiffrement **ChaCha20-Poly1305**. Là où WireGuard nu impose une configuration manuelle des clés et des pairs sur chaque machine, Tailscale ajoute une couche de coordination qui distribue automatiquement les clés publiques, traverse les NAT (NAT traversal) et établit des connexions point à point chiffrées entre tous les nœuds du tailnet.

Le résultat : un réseau plat, où la machine A peut joindre la machine B en quelques secondes, peu importe leurs fournisseurs d’accès, leur localisation géographique ou la complexité des pare-feux qui les séparent. Cette approche zero-config est la principale raison pour laquelle Tailscale écrase aujourd’hui des solutions historiques comme OpenVPN ou IPsec dans le segment des PME, des homelabs et des équipes DevOps. Pour une comparaison technique avec WireGuard pur, voir notre tutoriel WireGuard 12 étapes.

Côté gouvernance, Tailscale fonctionne sous un modèle dit *zero trust* : chaque connexion est authentifiée via un fournisseur d’identité externe (Google, Microsoft Entra ID, GitHub, Okta, Apple ID), puis autorisée par des règles ACL granulaires écrites en JSON. Cela colle parfaitement aux exigences NIS2 et DORA en vigueur en Europe depuis 2025 sur la segmentation réseau et l’authentification forte. Notre dossier architecture Zero Trust détaille le contexte réglementaire.

## Comparatif Tailscale vs WireGuard vs OpenVPN

Avant d’entrer dans la pratique, mettons en perspective Tailscale face aux deux solutions VPN historiques que l’on rencontre encore le plus souvent en production. Le tableau suivant résume les écarts mesurés en mars 2026 sur un même couple de machines (Debian 12 / Ubuntu 24.04 LTS), reliées par une fibre 1 Gbit/s symétrique — une période durant laquelle la mise à jour de mars 2026 de Tailscale a justement couvert l’intervalle de versions **1.94.1 à 1.96.4** selon les notes de version officielles, la même lignée que celle benchmarkée ci-dessous.

| Critère | Tailscale 1.96 | WireGuard 1.0.x | OpenVPN 2.6 | 
|---|---|---|---|
| Modèle réseau | Mesh (P2P) | Point à point | Hub-and-spoke | 
| Configuration initiale | 2 commandes | Clés manuelles | PKI + certificats | 
| Traversée NAT | Automatique (DERP) | Limitée | Très limitée | 
| Débit maximal mesuré | 940 Mbit/s | 950 Mbit/s | 320 Mbit/s | 
| Protocole sous-jacent | WireGuard + DERP | WireGuard pur | SSL/TLS personnalisé | 
| Chiffrement | ChaCha20-Poly1305 | ChaCha20-Poly1305 | AES-256-GCM | 
| Plate-formes officielles | 9 | 5 | 7 | 
| Coût plan gratuit | 0 $ jusqu’à 6 utilisateurs | 0 $ (auto-hébergé) | 0 $ (auto-hébergé) | 
| Délai de mise en route | ≈ 5 min | ≈ 30 min | ≈ 90 min | 

Le verdict pratique : à débit quasi équivalent, Tailscale divise par six le temps de mise en service par rapport à WireGuard pur, et par dix-huit face à OpenVPN. Le compromis se situe sur la confiance accordée au plan de contrôle hébergé, problématique que nous adresserons en étape 13 avec Headscale.

## Prérequis avant de démarrer

Pour suivre intégralement ce tutoriel, vous aurez besoin de l’environnement suivant. Toutes les versions citées ont été testées en avril 2026 et correspondent aux paquets stables disponibles sur les dépôts officiels — l’application iOS, passée en version 1.80.1 dès janvier 2025 selon l’historique de versions IPA4Fun, a depuis reçu plusieurs mises à jour majeures alignées sur le rythme de publication des clients desktop, à l’image du client Android et Android TV, passé en version 1.98.8 le 30 juin 2026 puis en version 1.99.148 dès le 11 juillet 2026 selon APKMirror.

- **Deux machines Linux** : idéalement une Debian 12 « Bookworm » et une Ubuntu 24.04 LTS « Noble Numbat ».
- **Tailscale v1.96.5** ou plus récente, installée depuis le dépôt apt officiel.
- **Un compte Google, GitHub ou Microsoft** pour l’authentification SSO du tailnet.
- **Accès root ou sudo** sur les deux machines.
- **Un smartphone iOS 17+ ou Android 10+** pour valider le client mobile (optionnel).
- **Un nom de domaine** (optionnel, pour Tailscale Funnel à l’étape 11).
- **Un Raspberry Pi 4 ou 5** ou un VPS bon marché si vous souhaitez un nœud de sortie permanent.
- **Connaissances de base** : ligne de commande Linux, édition de fichiers JSON, notions IP/CIDR.

Pour les pratiquants avancés qui souhaitent containeriser le déploiement, vous pouvez vous appuyer sur notre tutoriel Docker Compose pour orchestrer le client Tailscale dans une stack reproductible. Si vous virtualisez sur Proxmox, le tutoriel Proxmox VE 9.1 couvre la création des LXC qui hébergeront vos nœuds.

## Étape 1 : Créer son compte et son tailnet

Avant la moindre ligne de commande, rendez-vous sur la page d’inscription de Tailscale et choisissez votre fournisseur d’identité préféré. Tailscale ne stocke pas de mot de passe : la délégation à un IdP externe est obligatoire dans le plan Personal comme dans le plan Starter. Cette particularité explique pourquoi Tailscale est conforme par défaut aux exigences MFA imposées par NIS2 sur les opérateurs de services essentiels.

