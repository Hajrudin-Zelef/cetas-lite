---
id: collect-261001-general-networking/general-networking/throttlestop-undervolt-cpu-portable-12-etapes-2026-1
title: "Verifier les erreurs materielles WHEA (PowerShell, admin)"
domain: general-networking
role: reference
task: reference
actors: ["AMD", "Intel"]
dates: []
keywords: ["agent", "amd", "benchmark", "gpu", "intel", "mai"]
source: docs/RAG/collect-261001-general-networking/throttlestop-undervolt-cpu-portable-12-etapes-2026.md
source_anchor: ""
source_lines: [1, 45]
sha256: b90696a6335ed77d4137e5e2348c3bce3154baec96cf04f4cf269e1aca64c7b6
---

# Verifier les erreurs materielles WHEA (PowerShell, admin)

*Dernière mise à jour : 04 juin 2026.* Votre PC portable gaming hurle, grimpe à 100 °C et perd des FPS dès que l’action s’intensifie ? Le coupable porte un nom : le bridage thermique, ou *thermal throttling*. La bonne nouvelle, c’est qu’un utilitaire gratuit d’à peine 1,71 Mo selon la fiche FileCroco mise à jour le 2 mai 2026, la version 9.7.3 (bêta) publiée en avril 2025 par TechPowerUp et qui ajoute, selon TechSpot, le réglage par cœur (per-core adjustment) pour les puces Intel de 12e génération et plus récentes – un outil déjà crédité de 207 962 téléchargements sur Uptodown au compteur d’août 2026 –, peut souvent régler le problème sans tournevis, sans démontage et sans pâte thermique. Ce tutoriel **ThrottleStop** vous montre, en 12 étapes et une trentaine de minutes, comment undervolter (sous-volter) le processeur Intel de votre portable pour gagner jusqu’à 15 °C, stopper le bridage et récupérer les images par seconde perdues.

Nous couvrons l’installation de **ThrottleStop** 9.7.3 (bêta, avril 2025), qui succède à la 9.7.2 bêta du 16 janvier 2025 – une version que le fil NotebookTalk situe à la même date et qui retire l’ancien code de multiplicateur maximal tout en corrigeant la sauvegarde de l’offset PROCHOT, en plus d’ajouter la prise en charge des puces Arrow Lake selon UltrabookReview –, elle-même publiée après la 9.7 du 26 décembre 2024 (référencée par Uptodown) et la 9.6 du 26 mai 2023. Cette version 9.7, toujours proposée par TechPowerUp au format ZIP de 2 Mo selon MajorGeeks, cumule déjà 103 488 téléchargements recensés en septembre 2026, et FileHorse la qualifiait encore de sortie « il y a 10 mois » en octobre 2025 – preuve qu’elle est restée la version de référence tout au long de 2025. Le guide 2026 d’UltrabookReview continue même de citer la 9.7 comme version stable de référence, avec un réglage FIVR affiné pour les puces HX/K de 10e génération, tandis que la page SourceForge du projet affiche une dernière mise à jour au 6 janvier 2025. Nous détaillons aussi la mesure de référence avec le TS Bench, l’undervolting du CPU Core et du CPU Cache via la fenêtre FIVR, le réglage des limites de puissance (TPL), la gestion du BD PROCHOT, puis la sauvegarde d’un profil qui se lance au démarrage de Windows. Vous trouverez aussi 8 pièges courants, 8 problèmes de dépannage et une FAQ complète. Ce guide s’adresse aux processeurs Intel : si vous avez un CPU AMD Ryzen, l’undervolt passe par Ryzen Master (voir plus bas).

## Qu’est-ce que ThrottleStop et pourquoi undervolter un CPU portable ?

**ThrottleStop** est un petit logiciel Windows gratuit créé par Kevin Glynn, plus connu sous le pseudonyme « UncleWebb », et distribué par TechPowerUp. Sa mission tient en une phrase : rendre le contrôle du processeur à l’utilisateur. Sur un ordinateur portable, le constructeur applique des réglages de puissance et de tension volontairement conservateurs pour garantir la stabilité de toutes les puces sorties d’usine, y compris les plus mauvaises. Résultat : la plupart des CPU reçoivent bien plus de tension qu’ils n’en ont réellement besoin, ce qui produit une chaleur inutile.

Cette chaleur excédentaire déclenche le *thermal throttling* : dès que le processeur atteint sa température maximale (le Tjmax, généralement 100 °C sur les puces Intel mobiles), il abaisse brutalement sa fréquence pour ne pas se détériorer. Vos FPS chutent, le portable devient bruyant et les temps de compilation ou d’export s’allongent. L’undervolting consiste à réduire la tension fournie au CPU tout en gardant les mêmes fréquences : le processeur consomme moins, chauffe moins, et peut donc rester à haute fréquence plus longtemps avant d’être bridé.

ThrottleStop agit sur les trois grands types de bridage identifiés par UncleWebb : le bridage thermique (température), le bridage par limite de puissance (Power Limit / TDP) et le bridage lié au régulateur de tension (VRM). Contrairement à un overclock qui augmente les fréquences et les risques, un undervolt bien mené ne réduit que la tension : il n’y a aucune usure supplémentaire, et le gain se lit immédiatement au thermomètre. C’est l’une des rares optimisations logicielles qui améliore à la fois les performances, les nuisances sonores et l’autonomie.

## Undervolt, offset de tension et bridage thermique : les concepts clés

Avant de plonger dans le tutoriel **ThrottleStop**, il faut maîtriser quatre notions. Elles reviennent dans toutes les étapes et évitent 90 % des erreurs de débutant.

### L’offset de tension (Offset Voltage)

Vous ne fixez pas une tension absolue : vous appliquez un *décalage négatif* (offset) par rapport à la courbe de tension d’origine. Un offset de -100 mV signifie « donne 100 millivolts de moins que d’habitude, à chaque fréquence ». La courbe garde sa forme, elle est simplement translatée vers le bas. C’est ce qui rend l’undervolt sûr : le CPU continue de demander plus de tension quand il monte en fréquence, mais toujours 100 mV en dessous de la valeur d’usine.

### FIVR, CPU Core et CPU Cache

La fenêtre FIVR (Fully Integrated Voltage Regulator) est le cœur de l’undervolt dans ThrottleStop. Sur la majorité des processeurs Intel mobiles, deux rails comptent : **CPU Core** (les cœurs) et **CPU Cache** (le cache et l’anneau interne). Sur la plupart des architectures, ces deux rails sont liés : il faut leur appliquer le *même* offset, sous peine d’instabilité. On peut aussi undervolter l’iGPU (Intel GPU) et le System Agent, mais ces gains sont marginaux et plus risqués – on les laisse de côté pour un premier réglage.

### PL1, PL2 et le TPL

La fenêtre TPL (Turbo Power Limits) définit combien de watts le CPU a le droit de consommer. PL1 est la limite soutenue (le TDP réel après quelques secondes), PL2 la limite en pic (turbo court). Beaucoup de portables imposent un PL1 très bas pour rester silencieux, ce qui bride les performances même quand la température le permettrait. Ajuster ces valeurs – avec prudence – complète parfaitement l’undervolt.

## ThrottleStop vs Intel XTU vs Ryzen Master : le bon outil selon votre CPU

ThrottleStop n’est pas le seul logiciel d’undervolt, mais c’est le plus léger et le plus complet pour les portables Intel. Point crucial : **ThrottleStop ne fonctionne que sur les processeurs Intel**. Si votre machine embarque un CPU AMD Ryzen, il faut passer par Ryzen Master (voir notre tutoriel dédié à l’undervolt du GPU avec MSI Afterburner pour la partie carte graphique). Le tableau ci-dessous résume les différences.

| Critère | ThrottleStop 9.7.3 | Intel XTU | AMD Ryzen Master | 
|---|---|---|---|
| Éditeur | UncleWebb / TechPowerUp | Intel | AMD | 
| CPU compatibles | Intel uniquement | Intel uniquement | AMD Ryzen uniquement | 
| Taille / poids | < 1 Mo, portable | Installeur lourd | Installeur lourd | 
| Undervolt (offset) | Oui (FIVR) | Oui | Oui (Curve Optimizer) | 
| Limites de puissance | Oui (TPL : PL1/PL2) | Oui | Oui (PPT/TDC/EDC) | 
| Benchmark intégré | Oui (TS Bench) | Oui | Non | 
| BD PROCHOT / C-States | Oui | Non | Non | 
| Prix | Gratuit | Gratuit | Gratuit | 

Intel XTU fait globalement le même travail que ThrottleStop côté undervolt, mais il est plus lourd, moins réactif et offre moins de contrôle sur le BD PROCHOT et les C-States. La plupart des passionnés de portables Intel préfèrent ThrottleStop pour sa légèreté et sa gestion fine des profils. Pour les tests de stabilité, XTU et ThrottleStop se complètent bien : on peut valider un offset dans XTU puis piloter le tout depuis ThrottleStop. Le test comparatif de Tom’s Hardware France sur l’undervolting d’un portable gaming confirme les gains thermiques mesurables des deux approches.

