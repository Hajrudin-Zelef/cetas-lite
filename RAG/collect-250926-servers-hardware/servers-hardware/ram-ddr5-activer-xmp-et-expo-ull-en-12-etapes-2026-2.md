---
id: collect-250926-servers-hardware/servers-hardware/ram-ddr5-activer-xmp-et-expo-ull-en-12-etapes-2026-2
title: "Séquence de validation recommandée pour un profil RAM DDR5 overclocké"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Intel"]
dates: []
keywords: ["amd", "arr", "benchmark", "intel", "latency", "memory", "training"]
source: docs/RAG/clean4/ram-ddr5-activer-xmp-et-expo-ull-en-12-etapes-2026.md
source_anchor: ""
source_lines: [40, 85]
sha256: ffb227379dbb86a554e2dd783a89e535cb3d539d76a543e20ded780c9acae2d0
---

# Séquence de validation recommandée pour un profil RAM DDR5 overclocké

Côté kits disponibles, G.Skill a présenté sa gamme Trident Z5 NeoX RGB en 32 Go (2×16 Go) certifiée EXPO ULL lors du Computex 2026, déclinée en DDR5-6000 dans les variantes CL36, CL30, CL28 et CL26. Silicon Power a de son côté montré ses kits XPower Cyclone R certifiés ROG au même salon, tandis que Lexar a dévoilé dès le CES 2026 sa série THOR Z RGB en 32 Go à 6000 MT/s, compatible à la fois EXPO et XMP 3.0 pour couvrir les deux plateformes avec un seul produit. Côté Intel, Corsair a mis à jour en août 2026 sa gamme Vengeance RS DDR5 certifiée XMP 3.0, déclinée en quatre références (finitions Grey/Black et White) toutes cadencées à 6000 MHz, avec un kit 32 Go affiché à 505,99 $ et un 16 Go à 292,99 $, des tarifs qui donnent une idée concrète de l’ampleur de la flambée des prix outre-Atlantique. Pour les acheteurs qui partent de zéro, Kingston Fury Beast Black EXPO en 32 Go DDR5-6000 CL30 reste une référence citée comme choix de tête de gamme accessible dans les comparatifs français de 2026.

## Étape 1 à 3 : préparer son PC avant de toucher au BIOS

La préparation évite 90 % des mauvaises surprises. Ces trois premières étapes prennent dix minutes et vous donnent une base de comparaison indispensable pour mesurer les gains réels une fois le réglage terminé.

1. **Identifiez votre kit RAM exact.** Ouvrez CPU-Z (onglet Memory et SPD) ou HWiNFO64 pour relever la fréquence JEDEC actuelle, les timings par défaut et le nombre de profils XMP/EXPO stockés sur la puce SPD. Notez la référence commerciale complète du kit (visible sur l’étiquette physique ou dans Thaiphoon Burner) : elle vous servira si vous devez chercher le profil EXPO ULL sur le site du fabricant.
2. **Mesurez vos performances de référence.** Lancez un benchmark mémoire (AIDA64 Cache & Memory Benchmark en version d’essai, ou le module intégré à HWiNFO) et notez la bande passante en lecture/écriture ainsi que la latence en nanosecondes. Lancez aussi deux ou trois runs dans votre jeu habituel pour noter le FPS moyen actuel. Sans ces chiffres de départ, impossible de savoir si votre réglage a réellement apporté quelque chose.
3. **Sauvegardez vos réglages BIOS actuels.** La plupart des BIOS UEFI récents (MSI, ASUS, Gigabyte, ASRock) proposent une fonction de profil utilisateur permettant d’exporter la configuration actuelle sur une clé USB. Faites-le avant toute modification : en cas de blocage au démarrage, vous pourrez restaurer l’état initial en quelques secondes plutôt que de tout reconfigurer à la main.

Voici à quoi ressemble une lecture SPD typique dans HWiNFO64 avant activation d’un profil, pour un kit DDR5-6000 CL30 encore réglé sur ses valeurs JEDEC par défaut :

```
Module: DIMM_A2 - Kingston Fury Beast Black EXPO
Capacité: 16384 MB, Type: DDR5-4800 (JEDEC actif)
Timings actuels: 40-39-39-77 @ 4800 MT/s
Profil XMP disponible: Non détecté
Profil EXPO #1: DDR5-6000 CL30 (30-38-38-96), 1.35V
Profil EXPO ULL: Disponible (nécessite AGESA 2026+)
Tension VDDQ_TX: 1.10V (JEDEC), 1.25V (EXPO)
Tension VPP: 1.80V
```
Ce relevé confirme deux choses essentielles : le kit tourne encore à 4800 MT/s au lieu des 6000 MT/s annoncés sur la boîte, et un profil EXPO ULL est bien présent sur la puce mais reste inactif tant que le BIOS n’a pas été mis à jour vers une révision AGESA compatible. C’est exactement ce type de diagnostic qu’il faut poser avant de passer à l’étape suivante.

## Étape 4 à 6 : activer un profil XMP ou EXPO en un clic

Cette section couvre l’activation basique, celle qui rapproche déjà votre RAM de ses spécifications annoncées. Comptez cinq minutes.

1. **Entrez dans le BIOS.** Redémarrez le PC et appuyez sur la touche dédiée (Suppr ou F2 sur la plupart des cartes mères, F10 sur certains portables) dès l’écran de démarrage du constructeur. Basculez en mode avancé si votre carte mère démarre par défaut en interface simplifiée (bouton F7 chez MSI, F7 chez ASUS, souvent un bouton “Advanced Mode” visible en bas de l’écran).
2. **Localisez le menu mémoire.** Sur les cartes MSI, cherchez “Overclocking” puis “A-XMP” ou “DDR Memory Profile”. Chez ASUS, direction “Ai Tweaker” puis “AI Overclock Tuner” avec les options XMP I / XMP II / EXPO. Chez Gigabyte, l’onglet “Tweaker” contient une liste déroulante “Memory Profile”. Chez ASRock, cherchez “OC Tweaker” puis “Load XMP Setting” ou “Load EXPO Setting”.
3. **Sélectionnez le profil et sauvegardez.** Choisissez le premier profil disponible (souvent le plus stable, parfois nommé Profile 1 ou EXPO I), appuyez sur F10 pour sauvegarder et redémarrer. Le PC peut effectuer un ou deux redémarrages automatiques pendant la phase de “training” mémoire : c’est normal, ne coupez pas l’alimentation à ce moment.

Si l’écran reste noir après plusieurs tentatives, la plupart des cartes mères modernes disposent d’un système de récupération automatique (Q-Flash, MemOK!, Boot Failure Guard) qui réinitialise les réglages mémoire après trois à cinq échecs de démarrage consécutifs. Laissez le PC gérer ces cycles sans intervenir : couper l’alimentation manuellement pendant le training peut parfois nécessiter un reset CMOS complet à la pile.

Une fois de retour sous Windows, reprenez CPU-Z et vérifiez que la fréquence affichée correspond bien à celle annoncée sur la boîte du kit (6000 MT/s dans notre exemple, affiché comme 3000 MHz en fréquence réelle DDR). Si c’est le cas, la première étape critique est validée : votre RAM tourne enfin à sa vitesse commerciale. C’est déjà, à ce stade, la totalité du gain que la majorité des utilisateurs recherchent, et il est parfaitement possible de s’arrêter ici si l’overclocking plus poussé ne vous intéresse pas.

## Étape 7 à 9 : passer à EXPO ULL et régler les sous-timings avancés

Sur plateforme AMD Ryzen 9000, avec un kit certifié et un BIOS à jour, EXPO ULL ajoute un second palier de réglage. Ces trois étapes demandent un peu plus de rigueur, mais restent accessibles sans connaissances poussées en électronique.

1. **Vérifiez la présence de l’option ULL.** Une fois le profil EXPO de base activé (étape 6), retournez dans le même menu mémoire du BIOS. Si votre AGESA est à jour et votre kit certifié, une nouvelle case “ULL Enable” ou “EXPO Ultra Low Latency” apparaît, généralement grisée tant que le profil EXPO standard n’est pas déjà sélectionné.
2. **Activez ULL et laissez les sous-timings en automatique dans un premier temps.** Cochez “ULL Enable” sans toucher manuellement à tREFI, tRRDS ou tWR : le BIOS applique alors les valeurs certifiées par le fabricant du kit (les mêmes que celles validées en usine chez G.Skill ou Kingston sur les kits EXPO ULL). Sauvegardez et redémarrez.
3. **Ajustez la tension VDDP si l’option est disponible.** EXPO ULL introduit un réglage de tension séparé pour le contrôleur VDDP, distinct du VDDQ classique. Une valeur légèrement plus élevée (dans la plage validée par le fabricant, généralement affichée en info-bulle dans le BIOS) peut stabiliser les sous-timings les plus agressifs. Ne dépassez jamais la plage recommandée affichée par le BIOS lui-même : ce n’est pas un réglage à improviser à l’aveugle.

Voici un exemple de comparaison de timings, tel qu’il apparaît dans HWiNFO64 avant et après activation d’EXPO ULL sur un kit DDR5-6000 CL30 certifié :

