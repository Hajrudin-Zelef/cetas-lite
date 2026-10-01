---
id: collect-250926-servers-hardware/servers-hardware/msi-afterburner-2026-overclock-undervolt-gpu-tuto-2
title: "Cartes NVIDIA : lire le modèle, la conso et les températures"
domain: servers-hardware
role: reference
task: reference
actors: []
dates: []
keywords: ["attention", "benchmark", "gpu", "memory"]
source: docs/RAG/clean4/msi-afterburner-2026-overclock-undervolt-gpu-tuto.md
source_anchor: ""
source_lines: [51, 115]
sha256: 49df5832f2fde9e5866fb3f763dc10042bc1ba01a21357761d41e818754ab96d
---

# Cartes NVIDIA : lire le modèle, la conso et les températures

Au premier lancement, l’interface peut sembler intimidante avec ses cadrans et ses curseurs. En réalité, tout se résume à six réglages principaux et à un système de profils. L’interface « skin » par défaut affiche au centre la fréquence du cœur et la fréquence mémoire, et sur les côtés les jauges de température et d’utilisation. En bas, cinq emplacements numérotés (1 à 5) permettent d’enregistrer des profils. Le bouton **Apply** (la coche) valide vos réglages, le bouton **Reset** les annule, et l’icône en forme d’engrenage ouvre les paramètres avancés.

| Curseur | Nom anglais | Effet | Usage typique | 
|---|---|---|---|
| Limite de puissance | Power Limit (%) | Plafond de watts autorisé | Augmenter pour l’OC, baisser pour l’efficacité | 
| Limite de température | Temp Limit (°C) | Seuil de bridage thermique | Lier à la puissance pour la sécurité | 
| Fréquence du cœur | Core Clock (MHz) | Décalage de fréquence GPU | +100 à +250 MHz selon la carte | 
| Fréquence mémoire | Memory Clock (MHz) | Décalage de fréquence mémoire | +500 à +3000 MHz (GDDR7) | 
| Tension du cœur | Core Voltage (mV) | Tension appliquée au GPU | Déverrouiller puis réduire (undervolt) | 
| Vitesse ventilateurs | Fan Speed (%) | Régime des ventilateurs | Courbe personnalisée ou auto | 

Par défaut, le curseur de tension du cœur (Core Voltage) est verrouillé : nous le déverrouillerons à l’étape 3. Les curseurs de fréquence acceptent des valeurs négatives comme positives, ce qui sera utile pour le sous-voltage. Enfin, gardez à l’esprit que MSI Afterburner applique des **décalages (offsets)** et non des fréquences absolues : « +150 » signifie 150 MHz au-dessus de la courbe d’usine, et non une fréquence cible fixe.

## Étapes 1 à 4 : Installer et configurer MSI Afterburner

**Étape 1 – Télécharger.** Rendez-vous sur la page officielle de Guru3D et récupérez la version stable 4.6.6 Final (build 16757), parue le 29 septembre 2025 (ou, si vous voulez l’éditeur de courbe le plus abouti, la dernière bêta 4.6.7 : la Beta 3, build 17352 du 19 juin 2026 selon iTechGuides, a elle-même été dépassée depuis par la Beta 4, build 17439, arrivée le 8 août 2026 et qui porte enfin la plage de décalage mémoire à +3000 MHz). Le fichier d’installation ne pèse qu’environ 44 Mo et inclut RTSS 7.3.7.

**Étape 2 – Installer.** Lancez l’installateur en tant qu’administrateur. Lorsque l’assistant le propose, **cochez impérativement l’installation de RivaTuner Statistics Server** : sans lui, aucun overlay à l’écran ni compteur de FPS ne fonctionnera. Acceptez l’installation du pilote de bas niveau et redémarrez si le système le demande.

**Étape 3 – Déverrouiller le contrôle de la tension.** Ouvrez les paramètres (engrenage), onglet *General*, puis cochez « Unlock voltage control », « Unlock voltage monitoring » et « Force constant voltage » si vous prévoyez de l’overclocking poussé. Cochez également « Enable hardware control and monitoring ». Validez : le curseur Core Voltage devient actif.

**Étape 4 – Configurer la surveillance.** Toujours dans les paramètres, onglet *Monitoring*, activez les graphiques qui vous intéressent et cochez « Show in On-Screen Display » pour chacun. Les capteurs essentiels sont listés ci-dessous. Vous pouvez exporter votre configuration dans un fichier de profil pour la réutiliser sur une autre machine :

```
# Capteurs a activer dans l'onglet Monitoring (Show in On-Screen Display)
GPU temperature        -> degC, alerte au-dela de 83 degC
GPU Hot Spot           -> degC, point chaud (souvent +10/15 vs GPU)
GPU usage              -> %
Core clock             -> MHz
Memory clock           -> MHz
Power                  -> W (et % de la limite)
Core voltage           -> mV
Fan speed              -> % / tr/min
Framerate              -> FPS (fourni par RTSS)
Frametime              -> ms (regularite de l'affichage)
```
Réglez l’intervalle d’échantillonnage (Sampling period) sur 1000 ms pour un affichage lisible. Si vous gérez plusieurs profils de surveillance, le fichier de configuration de MSI Afterburner se trouve dans le dossier `Profiles` de l’installation : vous pouvez le sauvegarder avant toute modification importante.

## Étapes 5 à 7 : Surveiller son GPU et configurer l’overlay RTSS (OSD)

**Étape 5 – Vérifier la surveillance en temps réel.** Lancez un jeu ou un benchmark et observez le panneau de surveillance de MSI Afterburner. Vous devez voir évoluer la température, l’utilisation, les fréquences et la consommation. C’est votre tableau de bord : aucun overclocking ne se fait à l’aveugle.

**Étape 6 – Configurer l’overlay RTSS.** Ouvrez RivaTuner Statistics Server (icône dans la zone de notification). Réglez « On-Screen Display support » sur *ON*, choisissez un coin de l’écran et ajustez la taille. RTSS 7.3.7 prend en charge PresentMon V2, ce qui améliore la précision du frametime et la compatibilité avec la génération d’images DLSS 4. Vous pouvez personnaliser l’apparence de l’overlay via un bloc de configuration texte :

```
# Exemple de mise en forme de l'overlay RTSS (onglet Setup, On-Screen Display)
# Balises entre chevrons = variables RTSS ; C0..C3 = couleurs
GPU  : [GT] degC  |  [P] %  |  [F] FPS
VRAM : [MU] Mo    |  CONSO : [PWR] W
1% LOW : [FT] ms (frametime)
# Astuce : activer "Show own statistics" pour le FPS,
# et "Benchmark" pour enregistrer les 1% et 0,1% low.
```
**Étape 7 – Établir une mesure de référence (baseline).** Avant de modifier quoi que ce soit, lancez un benchmark reproductible et notez vos résultats d’origine : FPS moyen, 1 % low, température maximale, consommation maximale et fréquence GPU stabilisée. Sans cette référence, impossible de savoir si vos réglages améliorent réellement les choses. Unigine Superposition et 3DMark Steel Nomad sont parfaits pour cela car ils sont parfaitement répétables.

## Étapes 8 à 10 : Overclocker son GPU en toute sécurité

L’overclocking se fait par **petits incréments testés un par un**. La règle d’or : ne jamais augmenter deux paramètres en même temps, sous peine de ne pas savoir lequel cause une instabilité. Commencez toujours par les limites, puis le cœur, puis la mémoire.

**Étape 8 – Maximiser les limites de puissance et de température.** Poussez le curseur Power Limit à son maximum (souvent +10 à +15 % selon la carte) et liez-le au curseur Temp Limit (icône de cadenas). Cela donne au GPU la marge nécessaire pour tenir des fréquences plus élevées. C’est l’étape la plus sûre : elle ne touche pas aux fréquences, seulement au plafond énergétique.

**Étape 9 – Augmenter la fréquence du cœur.** Ajoutez +50 MHz au Core Clock, appliquez, puis testez 10 à 15 minutes avec un benchmark en boucle. Si c’est stable, ajoutez encore +25 à +50 MHz et recommencez. Continuez jusqu’à voir des artefacts graphiques, un plantage du pilote ou un écran noir. Reculez alors de 30 à 50 MHz pour garder une marge de sécurité.

**Étape 10 – Augmenter la fréquence mémoire.** La mémoire GDDR7 des RTX 50 dispose d’une marge énorme. MSI Afterburner permet désormais de pousser la mémoire jusqu’à **+3000 MHz d’offset**, soit jusqu’à 36 Gbit/s sur une RTX 5080 dont les puces 32 Gbit/s sont bridées à 30 Gbit/s d’usine. Procédez par paliers de +200 MHz. Attention : contrairement au cœur, une mémoire instable ne plante pas toujours – elle déclenche la correction d’erreurs qui **fait chuter les performances en silence**. Surveillez vos FPS : si le score baisse alors que vous montez la fréquence, vous êtes allé trop loin.

