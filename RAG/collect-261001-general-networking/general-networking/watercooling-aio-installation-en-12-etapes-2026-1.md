---
id: collect-261001-general-networking/general-networking/watercooling-aio-installation-en-12-etapes-2026-1
title: "Exemple de script de test (a adapter selon vos outils installes)"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["gpu"]
source: docs/RAG/collect-261001-general-networking/watercooling-aio-installation-en-12-etapes-2026.md
source_anchor: ""
source_lines: [1, 31]
sha256: 7e48fcff9d811cfa4f8cf215a83df60cbd61328af4f5ab8a19e8e8606c253e1b
---

# Exemple de script de test (a adapter selon vos outils installes)

Un Ryzen 9 9800X3D ou une RTX 5080 poussés à fond dégagent une chaleur qu’un simple ventirad peine désormais à évacuer sans monter en régime, et donc en bruit. C’est cette réalité qui pousse de plus en plus de joueurs français à franchir le pas du **watercooling AIO** (All-In-One), un système de refroidissement liquide scellé et prêt à l’emploi, sans les contraintes d’un circuit custom. Le mouvement est déjà largement engagé : selon une étude MarketIntelo, les AIO représentaient 34,6 % du marché des refroidisseurs CPU dès août 2025, une part qui aurait semblé irréaliste il y a encore cinq ans. Une autre étude, signée DataInsightsMarket, chiffre pour sa part le marché mondial du watercooling AIO à 489 millions de dollars en 2025, avec une croissance annuelle moyenne de 9,6 % projetée d’ici 2032 selon les données publiées en février 2026 — deux méthodologies différentes, mais le même constat d’un segment en pleine expansion. Longtemps réservé aux configurations haut de gamme, le watercooling PC est devenu accessible : Thermalright vend désormais un modèle 240 mm à 44,90 $, un tarif qui aurait semblé impensable pour du liquide il y a cinq ans. Signe que la bascule est désormais actée jusque chez les puristes de l’air, Noctua a lancé le 16 juin 2026 son tout premier AIO, le NL-LC1, décliné en 240, 360 et 420 mm autour d’une pompe Asetek Emma G8 V2 : quand la marque autrichienne la plus attachée au ventirad traditionnel se met elle aussi au liquide, difficile de continuer à voir le watercooling comme une niche.

Ce guide détaille les 12 étapes pour installer un watercooling AIO sur un PC de bureau, du choix de la taille de radiateur jusqu’au premier stress test. Comptez environ 45 minutes pour un boîtier déjà ouvert et un socket compatible, un peu plus si c’est votre première installation. Nous verrons aussi les erreurs qui reviennent le plus souvent sur les forums d’entraide, les pannes classiques et comment les résoudre, ainsi qu’un exemple de configuration complète testée de bout en bout.

## Pourquoi installer un watercooling AIO sur votre PC en 2026

Trois raisons expliquent l’engouement actuel pour le watercooling AIO. La première tient aux puces elles-mêmes : les CPU haut de gamme actuels tirent régulièrement plus de 250 W sous charge soutenue, un niveau que la plupart des ventirads à air gèrent en sacrifiant l’acoustique. La seconde tient au prix : la gamme s’est étoffée vers le bas, avec des modèles 240 mm sous les 50 € qui rendent le liquide compétitif face à un bon ventirad double tour, une dynamique renforcée par des acteurs comme ASRock, qui a dévoilé le 31 décembre 2025 une gamme complète de nouveaux AIO répartie sur six séries (Taichi, Phantom Gaming, Steel Legend, Challenger, Pro et WS) en prévision du CES, tenu début janvier 2026, et par Thermaltake, qui a doublé la mise sur la période : la marque a d’abord lancé le 14 janvier 2026 sa gamme TH V3 Series ARGB Sync, déclinée en trois variantes de bloc pompe, avant d’annoncer pour le Japon, dès le 11 septembre 2026, une plateforme complète de 14 produits incluant les AIO AW360 et AW420 de sa nouvelle gamme AW-series à partir de 79 980 ¥, selon TechTimes. La troisième est esthétique : boîtiers à panneau vitré, ARGB synchronisé et même écrans LCD intégrés (le NZXT Kraken Elite RGB 360 en est l’exemple le plus poussé) ont transformé le refroidissement en élément de personnalisation à part entière, une tendance que confirme le TRYX Holo 360 mm, dévoilé en juin 2026 avec un panneau réfléchissant qui rompt avec les habituels écrans plats.

Un AIO reste un circuit fermé et pré-rempli en usine. Contrairement à un circuit custom, vous n’ajoutez jamais de liquide, ne purgez rien et ne choisissez pas vos propres tubes. C’est précisément ce qui en fait une porte d’entrée réaliste vers le liquide pour quelqu’un qui n’a jamais dépassé le montage d’un ventirad. Si vous cherchez déjà à optimiser la partie logicielle de votre refroidissement, sachez que nous avons aussi détaillé comment undervolter votre GPU avec MSI Afterburner pour réduire la chaleur à la source, une étape qui se combine très bien avec un nouveau système de refroidissement liquide.

## Prérequis : matériel, outils et compatibilité avant de commencer

Avant de sortir le tournevis, réunissez ce qui suit. La liste paraît longue mais la plupart des éléments se trouvent déjà dans votre boîte à outils PC ou dans la boîte de l’AIO.

- **Le kit AIO complet** : bloc pompe, radiateur, ventilateurs déjà fournis, câble d’alimentation et notice papier ou QR code vers la notice numérique.
- **Un tournevis cruciforme PH2** , idéalement avec une tige magnétique pour ne pas perdre de vis dans le boîtier.
- **De l’alcool isopropylique à 90 % ou plus** et un chiffon non pelucheux ou un filtre à café, pour retirer l’ancienne pâte thermique.
- **De la pâte thermique de rechange** , uniquement si le bloc pompe de votre AIO n’est pas livré pré-appliqué (la majorité des modèles récents le sont).
- **Un boîtier compatible** avec la taille de radiateur choisie (240, 280, 360 ou 420 mm), à vérifier dans la fiche technique du fabricant du boîtier avant l’achat.
- **Une carte mère avec les en-têtes adaptés** : un en-tête CPU_FAN ou AIO_PUMP dédié à la pompe, des en-têtes SYS_FAN pour les ventilateurs du radiateur, et un en-tête ARGB 3 broches (5 V) ou RGB 4 broches (12 V) selon l’éclairage de votre kit.
- **Le logiciel de contrôle du fabricant** en version à jour : Corsair iCUE, NZXT CAM, MSI Center ou l’équivalent be quiet!/Lian Li selon votre modèle. Téléchargez toujours la dernière version depuis le site officiel plutôt qu’un exécutable fourni sur une clé USB qui accompagne la boîte, souvent périmé.
- **Un outil de monitoring** comme HWiNFO64, pour vérifier les températures avant et après l’installation. Nous avons consacré un guide complet à HWiNFO si vous ne l’avez jamais utilisé.
- **Un outil de stress test** : Cinebench 2026, FurMark ou Prime95 selon que vous testez le CPU, le GPU ou les deux.
- **45 à 60 minutes devant vous** , sans être interrompu en plein serrage des vis du waterblock.

Un point mérite une vérification préalable : la longueur des tubes. Un AIO 360 mm avec des tubes de 400 mm ne se monte pas de la même façon dans une tour compacte que dans un boîtier full tower. Mesurez la distance entre l’emplacement prévu du radiateur et celui du socket CPU avant de valider votre achat, un retour d’AIO pour cause de tubes trop courts reste l’un des motifs de réclamation les plus fréquents chez les revendeurs.

Le type de façade du boîtier compte aussi plus qu’on ne le pense. Une façade mesh (perforée) laisse circuler l’air librement et permet au radiateur de respirer sans effort supplémentaire pour les ventilateurs. Une façade en verre trempé, très répandue sur les boîtiers premium pour l’esthétique, restreint mécaniquement le flux d’air disponible en façade, ce qui pousse les ventilateurs du radiateur à tourner plus vite pour obtenir le même résultat thermique, donc plus de bruit à performance égale. Si l’acoustique compte pour vous autant que l’apparence, ce critère pèse au moins autant que la marque du radiateur au moment de choisir le boîtier.

## Comment fonctionne un watercooling AIO : pompe, radiateur et ventilateurs

