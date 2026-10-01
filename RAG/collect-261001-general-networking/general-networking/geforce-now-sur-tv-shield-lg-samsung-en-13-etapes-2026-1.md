---
id: collect-261001-general-networking/general-networking/geforce-now-sur-tv-shield-lg-samsung-en-13-etapes-2026-1
title: "Vérifier la version Fire OS installée en SSH/ADB (si le débogage est activé)"
domain: general-networking
role: reference
task: reference
actors: ["AWS", "Google", "Nvidia", "Samsung"]
dates: []
keywords: ["blackwell", "datacenter", "ethernet", "gpu", "luna", "nvidia"]
source: docs/RAG/collect-261001-general-networking/geforce-now-sur-tv-shield-lg-samsung-en-13-etapes-2026.md
source_anchor: ""
source_lines: [1, 45]
sha256: a5d13d22f1f65f9a27b930c4c4e2130b8943debba0fa179ec6cffa46e70e292a
---

# Vérifier la version Fire OS installée en SSH/ADB (si le débogage est activé)

Vous avez un abonnement GeForce NOW Ultimate à 19,99 $/mois — le tarif fixé par NVIDIA depuis août 2025 pour l’accès aux serveurs RTX 5080-class —, un salon équipé d’une télé 4K et vous voulez jouer sans allumer le PC. Bonne nouvelle : en 2026, GeForce NOW tourne nativement sur la Shield TV, les téléviseurs LG et Samsung, les boîtiers Android TV/Google TV et, depuis le CES 2026, sur certains Fire TV Stick d’Amazon. Mauvaise nouvelle : chaque plateforme a ses propres limites de résolution, ses propres bugs de configuration et ses propres pièges Bluetooth qui peuvent transformer une soirée jeu en séance de dépannage. Ce tutoriel détaille, appareil par appareil, comment installer, appairer et régler GeForce NOW sur une télé pour obtenir le meilleur flux possible avec le palier RTX 5080-class disponible en Ultimate.

## Pourquoi jouer à GeForce NOW sur téléviseur change la donne en 2026

Le cloud gaming sur grand écran n’est plus une expérience au rabais. Depuis le passage des serveurs Ultimate à l’architecture Blackwell (rendu RTX 5080-class), NVIDIA a aussi ouvert la porte du salon : l’application GeForce NOW est désormais native sur les téléviseurs LG sous webOS 5.0 et versions ultérieures, sur les Samsung Smart TV via le Samsung Gaming Hub, sur tous les appareils Android TV/Google TV (Shield TV, Chromecast avec Google TV, boîtiers tiers) et, nouveauté du CES 2026, sur les Fire TV Stick 4K Plus (2e génération) et 4K Max d’Amazon. Le principe reste le même que sur PC : aucune installation de jeu, aucun téléchargement, juste un flux vidéo streamé depuis un datacenter NVIDIA. Côté tarifs, NVIDIA a confirmé dès septembre 2025 les formules annuelles à 199,99 $ pour Ultimate et 99,99 $ pour Performance (soit 9,99 $/mois en mensuel contre 19,99 $/mois pour Ultimate), deux paliers qui donnent tous deux accès à l’app TV, la différence portant surtout sur la priorité serveur et le plafond de résolution. Mais sur télé, la contrainte change de nature. Ce n’est plus la puissance du GPU local qui compte, c’est la qualité du décodage vidéo de l’appareil, la latence Bluetooth de la manette, la gestion HDMI-CEC et le débit réellement disponible sur le port Ethernet ou le Wi-Fi de la box.

Beaucoup d’utilisateurs qui migrent du PC vers la télé se plantent sur ce point précis : ils appliquent les réglages qui fonctionnaient sur leur PC (résolution, bitrate, datacenter) sans se rendre compte que le client TV a ses propres plafonds. Un Fire TV Stick 4K Plus reste bridé à 1080p 60 FPS même sur une télé 4K, alors qu’une LG OLED récente peut grimper jusqu’à 4K HDR, voire 120 Hz sur certains modèles Micro RGB. Ce guide couvre les 13 étapes pour configurer correctement chaque famille d’appareils, tester le réseau, appairer les manettes et éviter les pièges les plus fréquents.

## Prérequis : compte, matériel et versions avant de commencer

Avant de vous lancer, réunissez ces éléments. Le temps total de configuration complète, tests réseau et calibration inclus, tourne autour de 90 minutes si vous partez de zéro sur un appareil neuf.

- Un compte NVIDIA actif, idéalement lié à un abonnement GeForce NOW Ultimate (accès prioritaire aux serveurs RTX 5080-class) ou Performance — ou, pour un essai ponctuel, un pass journalier à 7,99 $ (Ultimate) ou 3,99 $ (Performance), disponible depuis août 2026
- Un des appareils compatibles : NVIDIA Shield TV / Shield TV Pro, Chromecast avec Google TV, boîtier Android TV/Google TV récent, téléviseur LG webOS 5.0 ou plus récent (2020+), téléviseur Samsung Tizen 2020-2025 via Samsung Gaming Hub, ou Fire TV Stick 4K Plus (2e génération)/4K Max
- Une connexion internet stable : 25 Mbps minimum pour du 1080p60, 35-45 Mbps pour du 4K, idéalement en Ethernet filaire jusqu’au boîtier
- Une manette Bluetooth compatible : DualSense, DualShock 4, manette Xbox sans fil (modèle Bluetooth) ou un pad tiers certifié
- Le firmware ou l’OS du téléviseur à jour (webOS 24/25 pour LG, dernière version Tizen pour Samsung, dernière mise à jour Android TV/Google TV)
- Un adaptateur Ethernet USB si votre Chromecast ou Fire TV Stick n’a pas de port réseau filaire natif
- Un compte Google Play (pour les appareils Android TV/Google TV) ou un accès à l’App Store LG/Samsung selon votre téléviseur

Vérifiez aussi votre position par rapport aux datacenters NVIDIA. En France, les flux transitent généralement par les points de présence de Paris ou Francfort ; une latence inférieure à 40 ms est recommandée, le seuil maximal accepté par NVIDIA étant fixé à 80 ms de round-trip, comme le détaille la page officielle de configuration requise GeForce NOW.

## Étape 1 à 3 : choisir le bon appareil TV pour GeForce NOW

**Étape 1 — Identifiez votre scénario.** Si votre téléviseur est une LG ou une Samsung récente (2020 ou plus récent), vous n’avez besoin d’aucun boîtier externe : l’application est préinstallée ou disponible directement dans le LG Content Store / Game Portal ou le Samsung Gaming Hub. Si votre télé est plus ancienne ou d’une autre marque, il vous faut un boîtier Android TV, une Shield TV ou un Fire TV Stick compatible.

**Étape 2 — Comparez les plateformes disponibles.** Chaque famille d’appareils a un plafond de résolution et de fréquence différent sur GeForce NOW. Le tableau ci-dessous résume les capacités réelles constatées en 2026.

| Appareil | Résolution max GeForce NOW | FPS max | Ethernet natif | Prix indicatif | 
|---|---|---|---|---|
| NVIDIA Shield TV Pro | 4K HDR | 60 FPS (jusqu’à 120 Hz sortie TV) | Oui (Gigabit) | ≈ 199,99 $ | 
| NVIDIA Shield TV (tube) | 4K HDR | 60 FPS | Non (Wi-Fi natif) | ≈ 149,99 $ | 
| Chromecast avec Google TV (4K) | 4K HDR | 60 FPS | Non (adaptateur USB requis) | ≈ 49-59 € | 
| LG OLED webOS 24/25 (native) | 4K HDR, 120 Hz sur modèles Micro RGB/OLED select | 60-120 FPS selon modèle | Oui | Variable selon modèle | 
| Samsung Tizen (Gaming Hub) | 4K HDR | 60 FPS | Oui | Variable selon modèle | 
| Fire TV Stick 4K Plus (2e gén.) / 4K Max | 1080p (plafond imposé par l’app) | 60 FPS | Non (adaptateur USB requis) | ≈ 39-59 € | 

**Étape 3 — Tranchez selon votre priorité.** Pour la meilleure qualité d’image et la latence la plus faible, la Shield TV Pro reste la référence NVIDIA : port Ethernet Gigabit natif, décodeur vidéo optimisé, mises à jour logicielles fréquentes. Si vous avez déjà une LG OLED récente ou une Samsung compatible Gaming Hub, inutile d’acheter un boîtier : l’app native suffit. Si votre budget est serré et que vous acceptez un plafond 1080p60, le Fire TV Stick 4K Plus fait le travail pour moins de 60 €, mais gardez en tête que même sur une télé 4K, l’image restera limitée à 1080p tant que NVIDIA n’aura pas levé cette restriction sur les Fire TV Stick.

## Étape 4 à 6 : installer l’application selon votre plateforme

**Étape 4 — Android TV, Google TV, Shield TV.** Ouvrez le Google Play Store depuis l’appareil, recherchez « GeForce NOW », installez l’application officielle NVIDIA. Sur Shield TV, l’app est parfois préinstallée ou proposée en mise à jour automatique lors du premier démarrage.

**Étape 5 — LG webOS et Samsung Tizen.** Sur LG, ouvrez le LG Content Store (ou Game Portal selon la version webOS), recherchez GeForce NOW dans la catégorie jeux, installez. Sur Samsung, ouvrez le Samsung Gaming Hub depuis l’écran d’accueil, GeForce NOW y figure directement aux côtés de Xbox Cloud Gaming et Amazon Luna, sans installation séparée nécessaire : il s’agit d’un lien direct vers le service.

