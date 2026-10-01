---
id: collect-261001-general-networking/general-networking/watercooling-aio-installation-en-12-etapes-2026-2
title: "Exemple de script de test (a adapter selon vos outils installes)"
domain: general-networking
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: ["gpu", "mai"]
source: docs/RAG/collect-261001-general-networking/watercooling-aio-installation-en-12-etapes-2026.md
source_anchor: ""
source_lines: [32, 66]
sha256: c34f829778edfbd4b35264047784c0a2899976976368e9107710d1b16e8fe290
---

# Exemple de script de test (a adapter selon vos outils installes)

Un AIO repose sur un circuit fermé et scellé en usine. Le bloc pompe, posé directement sur l’IHS du processeur, capte la chaleur via une plaque de contact en cuivre puis pousse le liquide caloporteur à travers deux tubes souples jusqu’au radiateur. Les ventilateurs fixés sur ce radiateur évacuent la chaleur accumulée par le liquide vers l’extérieur du boîtier, avant que le liquide refroidi ne reparte vers la pompe. Ce cycle tourne en continu tant que le PC est sous tension.

La différence avec un circuit custom se joue sur trois points : un AIO ne se remplit jamais manuellement (à l’exception de rares modèles premium comme le be quiet! Light Loop, qui propose un port de remplissage optionnel et dont la déclinaison IO LCD 360 mm, testée par PC Guide le 15 septembre 2026, a été saluée pour son équilibre entre discrétion acoustique et performance), il n’intègre qu’un seul composant à refroidir la plupart du temps (le CPU, parfois le GPU sur des kits spécifiques), et il ne nécessite aucune purge d’air complexe au-delà d’un léger rodage les premières heures. C’est ce qui explique pourquoi un watercooling AIO se monte en une grosse demi-heure quand un circuit custom peut demander une journée entière.

Le radiateur lui-même mérite un mot d’explication, car deux modèles de même taille en façade peuvent se comporter très différemment. La densité d’ailettes (FPI, fins per inch) conditionne la surface d’échange thermique disponible : plus elle est élevée, plus le radiateur dissipe de chaleur à débit d’air égal, mais plus il impose aussi une pression statique importante aux ventilateurs, qui doivent alors tourner plus vite pour rester efficaces. C’est pour cette raison que les fabricants premium comme ARCTIC ou Corsair associent systématiquement leurs radiateurs à haute densité avec des ventilateurs spécifiquement calibrés pour la pression statique plutôt que pour le débit brut, une distinction qui explique pourquoi mélanger un radiateur haut de gamme avec des ventilateurs génériques donne rarement le résultat espéré.

### Choisir la taille de son radiateur : 240, 280, 360 ou 420 mm

La taille du radiateur détermine directement la capacité de dissipation thermique de votre watercooling AIO. Un radiateur 240 mm (deux ventilateurs de 120 mm) absorbe généralement entre 200 et 245 W avant de saturer, ce qui couvre large la majorité des CPU gaming actuels en usage courant. Un 360 mm (trois ventilateurs de 120 mm) grimpe vers 280 à 300 W et plus, la zone où se situent les puces les plus gourmandes comme les Ryzen X3D poussés en overclocking ou les Core Ultra haut de gamme. Le format 420 mm, plus rare et réservé aux boîtiers grand format, reste pertinent surtout pour les configurations HEDT ou les CPU les plus extrêmes du marché.

| Taille radiateur | TDP CPU recommandé | Emplacement type | Fourchette de prix | Profil sonore | 
|---|---|---|---|---|
| 240 mm | Jusqu’à ~200-245 W | Façade ou toit (boîtiers compacts) | Environ 45-90 € | Souvent le plus silencieux à charge égale | 
| 280 mm | ~230-270 W | Façade ou toit | Environ 90-130 € | Bon compromis silence/performance | 
| 360 mm | ~280-300 W et plus | Façade (recommandé) ou toit | Environ 90-160 € | Format standard pour CPU haut de gamme | 
| 420 mm | 300 W et plus | Façade, boîtiers grand format uniquement | Environ 200-300 € | Réservé aux configurations les plus exigeantes | 

Un principe simple pour trancher : si votre boîtier accepte un 360 mm en façade, prenez-le plutôt qu’un 240 mm, même pour un CPU milieu de gamme. Les chiffres confirment l’intuition : selon un comparatif SunbeamTech de juillet 2026, un radiateur 360 mm tourne entre 8 et 18°C plus frais qu’un 240 mm en pleine charge. La marge thermique supplémentaire se traduit directement par des ventilateurs qui tournent moins vite à charge égale, donc moins de bruit au quotidien.

## Comparatif 2026 : les meilleurs watercooling AIO du marché

Le marché 2026 se structure clairement en trois segments : l’entrée de gamme portée par Thermalright, le milieu de gamme équilibré (ARCTIC, Cooler Master, Lian Li, et désormais Gigabyte avec son Eagle 360 annoncé le 15 septembre 2026 : une pompe donnée pour 3 200 RPM et jusqu’à 82,31 CFM de débit d’air, des ventilateurs plafonnant à 37,6 dBA pour 2,52 mmH₂O de pression statique selon Times of India Tech, commercialisé dès septembre 2026) et le haut de gamme orienté esthétique et logiciel (NZXT, Corsair, be quiet!), auquel s’ajoute désormais Noctua. La marque avait publié sa feuille de route dès septembre 2025, Tom’s Hardware rapportant alors un report du premier trimestre au deuxième trimestre 2026 pour son AIO, avant qu’Asetek ne confirme le 31 mars 2026 que le kit avait validé ses tests de production (PVT) et qu’il arrive finalement en rayon à la mi-juin 2026. Le segment milieu de gamme n’est pas en reste côté affichage : Lian Li a dévoilé en juin 2026 le HydroShift II, un AIO à écran OLED incurvé lancé le 22 mai 2026 selon Wccftech, tandis qu’ASUS a présenté dès le 7 janvier 2026 au CES son ROG Strix LC IV doté d’un écran de 5,08 pouces, preuve que l’affichage embarqué ne reste plus l’apanage du haut de gamme. Voici les références qui reviennent le plus souvent dans les configurations montées cette année, avec les prix constatés et les caractéristiques communiquées par les fabricants.

| Modèle | Radiateur | Prix constaté | Pompe / Bruit | Point fort | 
|---|---|---|---|---|
| ARCTIC Liquid Freezer III Pro 360 | 360 mm | ~90-105 $ (MSRP 89,99 $) | 3000 RPM / 22,5 dBA | VRM heatsink intégré, meilleur rapport qualité-prix du segment | 
| NZXT Kraken Elite RGB 360 | 360 mm | ~220-250 $ (MSRP 249,99 $) | 28 dBA | Écran LCD 60 Hz avec overlay CPU/GPU | 
| Corsair iCUE LINK TITAN 420 RX | 420 mm | ~270-300 $ (MSRP 299,99 $) | Non communiqué | Plus gros radiateur du comparatif, câblage caché iCUE LINK | 
| be quiet! Light Loop 360 | 360 mm | ~140-160 $ (MSRP 159,99 $) | Très silencieux | Port de remplissage, garantie 5 ans | 
| Lian Li Galahad II Trinity Performance 360 | 360 mm | ~160 $ | Non communiqué | Forte densité RGB, montage jugé facile | 
| Cooler Master MasterLiquid 360 Atmos | 360 mm | ~110 $ | 29 dBA | Gère 300 W+ pour un tarif contenu | 
| Thermalright Aqua Elite 240 V3 | 240 mm | ~45 $ (MSRP 44,90 $) | 1800 RPM / 25,6 dBA | Le ticket d’entrée le moins cher du marché | 

Ces tarifs sont exprimés en dollars US, base sur laquelle communiquent la plupart des fabricants. Comptez généralement 15 à 20 % de plus sur le marché français une fois la TVA et la marge distributeur appliquées, un 360 mm ARCTIC qui affiche 90 $ aux États-Unis se retrouvant donc plutôt autour de 100 à 110 € chez un revendeur comme Materiel.net ou LDLC. Les colonnes marquées « non communiqué » reflètent l’absence de donnée officielle chiffrée au moment de la rédaction, mieux vaut l’indiquer que d’inventer un chiffre.

