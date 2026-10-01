---
id: collect-250926-servers-hardware/servers-hardware/ram-ddr5-activer-xmp-et-expo-ull-en-12-etapes-2026-3
title: "Séquence de validation recommandée pour un profil RAM DDR5 overclocké"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD"]
dates: []
keywords: ["amd", "gpu", "latency", "memory"]
source: docs/RAG/clean4/ram-ddr5-activer-xmp-et-expo-ull-en-12-etapes-2026.md
source_anchor: ""
source_lines: [86, 191]
sha256: 1fa76ee3c42abf5d779c0db96dacaf2e6c2e7bc121c337f9c5834f8ab14cd203
---

# Séquence de validation recommandée pour un profil RAM DDR5 overclocké

```
--- Avant (EXPO standard) ---
Fréquence: DDR5-6000, CL30-38-38-96
tREFI: 32768 (valeur JEDEC standard)
tRRDS: 8
tWR: 48
ULL Enable: Désactivé
VDDP: 1.00V (auto)
--- Après (EXPO ULL activé) ---
Fréquence: DDR5-6000, CL30-38-38-96
tREFI: 22000 (rafraîchissement plus fréquent, latence réduite)
tRRDS: 6
tWR: 40
ULL Enable: Activé
VDDP: 1.05V
Latence mesurée (AIDA64): -6.2 ns vs profil EXPO standard
```
La fréquence brute (6000 MT/s) et les timings principaux (CL30-38-38-96) ne bougent pas : c’est bien la couche de sous-timings qui change, avec un rafraîchissement mémoire plus fréquent (tREFI plus bas) qui réduit les micro-latences au prix d’une charge légèrement supérieure sur le contrôleur mémoire. C’est précisément ce compromis qui explique pourquoi EXPO ULL reste réservé aux kits certifiés capables d’encaisser ce régime sans perdre en stabilité.

## Étape 10 à 12 : stabiliser, tester et valider son overclock

Un profil qui démarre sous Windows n’est pas nécessairement stable sur la durée. Ces trois dernières étapes valident que votre configuration tiendra la charge en jeu pendant des heures sans plantage ni corruption de données.

1. **Lancez un test de stabilité mémoire dédié.** Windows Memory Diagnostic est trop léger pour détecter des erreurs d’overclocking mémoire. Préférez TestMem5 avec la configuration Anta777 (extrême) ou Karhu RAM Test, deux outils gratuits reconnus par la communauté d’overclocking pour leur capacité à détecter des erreurs en quelques minutes plutôt qu’en heures.
2. **Faites tourner le test au minimum 30 à 60 minutes sans erreur.** Une erreur détectée, même une seule, signifie que le profil actuel n’est pas fiable pour un usage quotidien : mieux vaut revenir à l’étape précédente et assouplir légèrement les sous-timings (remonter tREFI d’un cran, par exemple) plutôt que de risquer une corruption de sauvegarde de jeu ou un crash en pleine partie compétitive.
3. **Validez avec un stress test combiné CPU+RAM.** Une fois le test mémoire pur validé, complétez avec un stress test qui sollicite processeur et mémoire simultanément (OCCT en mode “Large Data Set”, ou le protocole détaillé dans notre guide de stress test CPU et GPU) pendant au moins 20 minutes, en surveillant les températures du contrôleur mémoire intégré au CPU via HWiNFO.

Voici un exemple de séquence de test recommandée, à exécuter dans l’ordre pour valider un profil du plus rapide au plus exigeant :

```
# Séquence de validation recommandée pour un profil RAM DDR5 overclocké
1. TestMem5 (config Anta777) - 30 cycles minimum (~45 min)
   -> 0 erreur requis avant de continuer
2. Karhu RAM Test - 2000% couverture minimum (~60 min)
   -> Aucune "error" dans le log, sinon reset immédiat du profil
3. OCCT Large Data Set (CPU + RAM combinés) - 20 min
   -> Surveiller temp. IMC (Integrated Memory Controller) < 85°C
4. Session de jeu réelle - 2h minimum sur le titre le plus exigeant
   -> Aucun crash, aucun artefact visuel, aucun freeze
Si un test échoue: revenir au BIOS, assouplir tREFI ou tRRDS d'un cran,
relancer la séquence depuis l'étape 1.
```
Un profil qui passe cette séquence complète peut raisonnablement être considéré comme stable pour un usage quotidien intensif, y compris en sessions de jeu de plusieurs heures. Ne sautez jamais l’étape du test de jeu réel : certains titres sollicitent la mémoire différemment des outils de stress test synthétiques, et un profil “propre” sur TestMem5 peut occasionnellement révéler une instabilité uniquement visible en conditions de jeu prolongées.

## Projet complet : le profil DDR5-6000 CL30 EXPO ULL pas à pas

Pour rendre ce guide directement exploitable, voici un profil complet, testé selon la séquence décrite plus haut, pour un kit DDR5-6000 CL30 certifié EXPO ULL sur plateforme Ryzen 9000 / AM5. Ces valeurs servent de point de départ réaliste : elles doivent être validées sur votre propre matériel avant tout usage prolongé, chaque contrôleur mémoire silicium ayant sa propre marge de tolérance.

```
# Profil RAM DDR5-6000 CL30 - EXPO ULL - Ryzen 9000 / AM5
# À reproduire manuellement dans le BIOS, section Overclocking/OC Tweaker
[Timings principaux]
Fréquence         = DDR5-6000 (3000 MHz réel)
CAS Latency (tCL) = 30
tRCD               = 38
tRP                = 38
tRAS               = 96
[Sous-timings EXPO ULL]
tREFI              = 22000
tRRDS              = 6
tWR                = 40
ULL Enable         = Activé
[Tensions]
VDDQ_TX             = 1.25V
VPP                 = 1.80V
VDDP                = 1.05V
VSOC (Ryzen)        = 1.15V (ne pas dépasser 1.25V en usage prolongé)
[Résultat attendu après validation]
Bande passante lecture : +3 à 5% vs profil EXPO standard
Latence AIDA64          : -6 ns vs profil EXPO standard
Gain FPS moyen en jeu   : environ +4% (donnée constructeur AMD)
```
Ce profil correspond aux réglages appliqués en usine sur des kits comme le G.Skill Trident Z5 NeoX RGB en variante CL30, présenté au Computex 2026. Si votre kit affiche des timings de base différents (CL36, CL28 ou CL26 selon la variante achetée), n’appliquez jamais ces valeurs de sous-timings telles quelles : chargez d’abord le profil EXPO natif correspondant à votre kit, puis activez uniquement l’option ULL Enable, en laissant le BIOS appliquer automatiquement les sous-timings validés par le fabricant pour votre référence exacte. Le tableau des prérequis plus haut rappelle les logiciels nécessaires pour vérifier, à chaque étape, que les valeurs affichées correspondent bien à celles attendues.

### Variante pour un kit CL36 d’entrée de gamme

Les kits CL36, un peu moins onéreux que les CL30 lorsqu’ils restent disponibles, suivent la même logique avec des timings de base plus détendus. Voici le profil équivalent pour cette variante, toujours à valider avec la séquence de test complète avant tout usage quotidien :

```
# Profil RAM DDR5-6000 CL36 - EXPO ULL - Ryzen 9000 / AM5
[Timings principaux]
Fréquence         = DDR5-6000 (3000 MHz réel)
CAS Latency (tCL) = 36
tRCD               = 38
tRP                = 38
tRAS               = 96
[Sous-timings EXPO ULL]
tREFI              = 24000
tRRDS              = 7
tWR                = 44
ULL Enable         = Activé
[Tensions]
VDDQ_TX             = 1.25V
VPP                 = 1.80V
VDDP                = 1.05V
```
## Résultats attendus : gains de FPS et de latence mesurés

Les gains d’un réglage RAM ne se ressentent jamais de la même façon selon le jeu et la résolution. Sur un GPU haut de gamme en 1080p, où le processeur devient souvent le facteur limitant, l’écart entre RAM à fréquence JEDEC et RAM optimisée peut représenter plusieurs dizaines de FPS dans les titres compétitifs. En 4K avec un GPU saturé, l’écart se réduit fortement puisque c’est la carte graphique qui devient le goulot d’étranglement, pas la mémoire vive.

| Étape d’optimisation | Fréquence effective | Latence relative | Gain FPS estimé (1080p, jeu CPU-bound) | 
|---|---|---|---|
| JEDEC par défaut (aucun profil) | 4800 MT/s | Référence | Référence (0%) | 
| XMP ou EXPO activé | 6000 MT/s | Réduite (fréquence plus élevée) | Gain le plus important du parcours | 
| EXPO ULL activé | 6000 MT/s | -5 à -7 ns vs EXPO standard | +4% supplémentaire (donnée AMD) | 

