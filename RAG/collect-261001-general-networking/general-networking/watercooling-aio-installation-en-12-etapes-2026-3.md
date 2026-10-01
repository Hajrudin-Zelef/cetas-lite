---
id: collect-261001-general-networking/general-networking/watercooling-aio-installation-en-12-etapes-2026-3
title: "Exemple de script de test (a adapter selon vos outils installes)"
domain: general-networking
role: reference
task: reference
actors: ["AMD", "Intel", "Nvidia"]
dates: []
keywords: ["amd", "attention", "gpu", "intel", "mai", "nvidia"]
source: docs/RAG/collect-261001-general-networking/watercooling-aio-installation-en-12-etapes-2026.md
source_anchor: ""
source_lines: [67, 110]
sha256: 0ee80c54c52cf4c46bb3209957b4a566759a4571b0ab18d4de16337b298fcd5d
---

# Exemple de script de test (a adapter selon vos outils installes)

Le Liquid Freezer III Pro d’ARCTIC mérite une mention à part : c’est aujourd’hui la référence citée en premier dans la plupart des comparatifs matériel spécialisés de 2026, notamment chez Gamers Nexus, dont les protocoles de test thermique font référence dans la presse hardware anglophone. Sa combinaison prix contenu, VRM heatsink intégré (rare à ce tarif) et pompe capable de 3000 RPM en fait un choix quasiment par défaut pour qui veut un 360 mm sans complication. Côté tarifs, MaxMyBuild relevait en juin 2026 un Liquid Freezer III 240 autour de 65 $ et la version 360 autour de 105 $, des prix qui restent compétitifs face au reste du segment malgré une légère hausse depuis le lancement. Le nouveau venu à surveiller reste toutefois le Noctua NL-LC1, entrevu pour la première fois le 20 mai 2025 sur OC3D puis officialisé à Computex 2026 avant sa mise en vente mondiale le 16 juin 2026 à 11h00 CEST selon TechTimes : la caution acoustique de Noctua appliquée au liquide pourrait rebattre les cartes du haut de gamme dès que les premiers comparatifs indépendants seront disponibles.

## Vérifier la compatibilité socket avant l’achat : LGA1851, LGA1700, AM5, AM4

Chaque plateforme processeur impose son propre système de fixation, et confondre les brackets reste une source classique d’installation ratée. Les CPU Intel Core Ultra 200 récents utilisent le socket LGA1851, qui impose un contact frame dédié pour éviter le bombement du socket sous la pression de fixation, un problème qui avait touché certains montages sur les générations LGA1700 précédentes. Côté AMD, le socket AM5 conserve un bracket à ressorts classique, mais réclame un couple de serrage régulier des quatre vis pour éviter que le waterblock ne se pose de travers sur l’IHS.

| Socket | Plateforme | Type de fixation | Point d’attention | 
|---|---|---|---|
| LGA1851 | Intel Core Ultra 200 (2025-2026) | Contact frame dédié | Vérifiez que l’AIO inclut le kit LGA1851, tous les anciens kits LGA1700 ne sont pas rétrocompatibles | 
| LGA1700 | Intel Core 12ᵉ à 14ᵉ génération | Bracket à ressorts standard | Compatible avec la majorité des kits récents, parfois via adaptateur fourni | 
| AM5 | AMD Ryzen 7000 à 9000 | Bracket à ressorts | Serrage en croix impératif pour éviter un tilt du waterblock | 
| AM4 | AMD Ryzen 1000 à 5000 | Bracket à ressorts (legacy) | Certains fabricants ne livrent plus le kit AM4 en standard, à confirmer avant achat | 

Avant tout achat, un coup d’œil rapide à l’utilitaire système de Windows permet de confirmer le socket exact de votre carte mère si vous avez un doute. Ouvrez une invite de commandes et tapez la commande suivante.

```
wmic baseboard get product,manufacturer,version
systeminfo | findstr /C:"Processeur" /C:"Processor"
```
Ces deux lignes affichent respectivement le modèle exact de votre carte mère et le processeur installé, de quoi croiser l’information avec la liste de compatibilité publiée sur la fiche produit de l’AIO avant de sortir la carte bancaire.

## Installation étape par étape : les 12 étapes complètes

Voici la procédure complète, dans l’ordre où elle se déroule concrètement sur un boîtier ouvert et posé à plat. Chaque étape part du principe que le PC est hors tension et débranché du secteur.

1. **Éteignez le PC et débranchez l’alimentation secteur.** Posez le boîtier sur une surface stable, plane et bien éclairée. Retirez le panneau latéral et, si besoin, le panneau opposé pour accéder plus facilement à l’arrière de la carte mère.
2. **Démontez l’ancien système de refroidissement.** Débranchez le ventirad ou l’AIO existant, retirez son bracket, puis nettoyez l’ancienne pâte thermique sur l’IHS du CPU avec de l’alcool isopropylique et un chiffon non pelucheux. Ne réutilisez jamais l’ancienne pâte séchée.
3. **Identifiez le socket et installez le bracket adapté.** Repérez si votre carte mère est en LGA1851, LGA1700, AM5 ou AM4, puis montez le backplate (si nécessaire) et le contact frame ou bracket fourni avec votre AIO en suivant scrupuleusement la notice, chaque plateforme a sa propre géométrie de vis.
4. **Fixez les ventilateurs sur le radiateur avant de le monter dans le boîtier.** Vérifiez le sens des pales imprimé sur le cadre du ventilateur : une flèche indique le sens du flux d’air, une autre le sens de rotation. Montez-les à l’envers et vous inverserez tout le flux thermique de votre configuration.
5. **Positionnez le radiateur à l’emplacement choisi.** En façade avant, orientez les ventilateurs en intake (aspiration vers l’intérieur) pour la majorité des configurations gaming, c’est la configuration qui refroidit le mieux le CPU dans la plupart des tests. Vissez le radiateur en croix, comme pour une roue de voiture, sans jamais forcer un coin avant les autres.
6. **Appliquez la pâte thermique si nécessaire.** La plupart des blocs pompe modernes arrivent pré-appliqués d’usine, dans ce cas ignorez cette étape. Sinon, déposez un point de la taille d’un grain de riz au centre de l’IHS pour un CPU standard, ou un fin trait en X pour les IHS plus grands.
7. **Positionnez le bloc pompe sur le CPU.** Alignez les trous de fixation du bloc avec les vis du bracket, puis serrez progressivement et en croix, jamais un coin à fond avant d’entamer le suivant. Un serrage inégal incline le waterblock et crée un point chaud d’un seul côté du die.
8. **Branchez le connecteur de la pompe.** Reliez-le à l’en-tête AIO_PUMP ou CPU_FAN de la carte mère, jamais à un en-tête SYS_FAN générique qui pourrait réduire sa vitesse en veille et créer un point de surchauffe silencieux.
9. **Branchez les ventilateurs du radiateur.** Utilisez le hub fourni ou les en-têtes SYS_FAN disponibles, puis raccordez le câble ARGB ou RGB sur l’en-tête correspondant. Ne confondez jamais un en-tête ARGB 3 broches (5 V) avec un en-tête RGB 4 broches (12 V), l’erreur peut griller le contrôleur LED.
10. **Refermez le boîtier et redémarrez.** Rebranchez l’alimentation secteur, allumez le PC et entrez immédiatement dans le BIOS pour vérifier que la pompe et les ventilateurs sont bien détectés avec une vitesse de rotation cohérente, généralement affichée dans l’onglet Hardware Monitor ou équivalent.
11. **Installez le logiciel de contrôle du fabricant.** Une fois dans Windows, installez iCUE, CAM, MSI Center ou l’utilitaire correspondant à votre marque, puis configurez une courbe de ventilation et de pompe adaptée à votre usage (voir la section dédiée plus bas).
12. **Lancez un test de charge de 15 à 30 minutes.** Surveillez les températures avec HWiNFO64 pendant tout le test pour confirmer que l’installation tient la charge sans dérive anormale (voir la section sur les tests plus bas pour les seuils à surveiller).

Si vous remontez un PC complet plutôt que de simplement changer le refroidissement, c’est aussi le bon moment pour repartir sur des pilotes GPU propres. Notre guide sur Display Driver Uninstaller (DDU) détaille comment supprimer un ancien pilote NVIDIA ou AMD avant d’en réinstaller un neuf, une étape qui évite bien des plantages après un montage.

## Configurer le BIOS et le logiciel de contrôle : iCUE, NZXT CAM, MSI Center

Une fois l’AIO physiquement installé, la configuration logicielle détermine si vous obtenez un système silencieux au repos et efficace en charge, ou l’inverse. Deux niveaux de réglage existent : la courbe native du BIOS, qui fonctionne même sans Windows installé, et le logiciel du fabricant, plus fin mais dépendant de l’OS.

