---
id: collect-261001-general-networking/general-networking/throttlestop-undervolt-cpu-portable-12-etapes-2026-2
title: "Verifier les erreurs materielles WHEA (PowerShell, admin)"
domain: general-networking
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["advisory", "intel"]
source: docs/RAG/collect-261001-general-networking/throttlestop-undervolt-cpu-portable-12-etapes-2026.md
source_anchor: ""
source_lines: [46, 119]
sha256: 1c430c03f0f5057886d8507052e0c0aaea2885392dd94f1ec3caaaf4795148ef
---

# Verifier les erreurs materielles WHEA (PowerShell, admin)

## Prérequis : versions, matériel compatible et sauvegardes

Réunissez les éléments suivants avant de commencer. La compatibilité du processeur est le point le plus important : les puces les plus récentes ont souvent l’undervolt verrouillé (voir la section Plundervolt).

| Élément | Version / recommandation | Rôle | 
|---|---|---|
| ThrottleStop | 9.7.3 (avril 2025) | Undervolt et gestion du throttling | 
| Système d’exploitation | Windows 10 ou Windows 11 (64 bits) | ThrottleStop est Windows uniquement | 
| Processeur | Intel Core mobile (4e à 14e gén.) | Undervolt via FIVR ; à vérifier au cas par cas | 
| HWiNFO (recommandé) | Dernière version | Surveillance fine des tensions et températures | 
| Outil de charge | Cinebench R23 ou TS Bench | Test de stabilité sous charge | 
| .NET Framework | 4.x (fourni avec Windows) | Dépendance d’exécution | 
| Point de restauration | À créer avant de commencer | Filet de sécurité en cas d’instabilité | 

Deux précautions valent de l’or. D’abord, créez un point de restauration Windows : un undervolt trop agressif peut provoquer un écran bleu (BSOD) au démarrage, et il est rassurant de pouvoir revenir en arrière. Ensuite, installez HWiNFO en suivant notre tutoriel : c’est le meilleur moyen de vérifier en temps réel que votre offset est bien appliqué et de surveiller la température, la fréquence et la puissance package pendant les tests.

## ThrottleStop est-il dangereux ? Plundervolt (CVE-2019-11157) et le verrouillage FIVR

C’est la question numéro un, et la réponse est nuancée. Un undervolt qui ne touche qu’à l’Offset Voltage ne présente **aucun risque physique** pour votre matériel : vous réduisez la tension, jamais vous ne l’augmentez au-delà des valeurs d’usine. Le pire scénario est un plantage ou un écran bleu si l’offset est trop agressif – sans conséquence durable, un simple redémarrage annule le réglage non sauvegardé.

Il existe toutefois une limite importante à connaître. En décembre 2019, des chercheurs ont révélé Plundervolt (référencée CVE-2019-11157), une attaque exploitant l’interface d’undervolting pour corrompre les enclaves sécurisées Intel SGX en manipulant la tension. Elle a touché les processeurs Intel Core de la 6e à la 10e génération. Pour la corriger, Intel a diffusé une mise à jour de microcode qui, combinée à une mise à jour du BIOS, permet de *verrouiller* l’interface d’undervolt : la tension est alors figée aux valeurs par défaut. Plus récemment, un avis publié le 6 août 2025 dans la GitHub Advisory Database a référencé une faille distincte, **CVE-2025-7771**, touchant cette fois le pilote ThrottleStop.sys en version 3.0.0.0 : une raison supplémentaire de toujours télécharger l’outil depuis la page officielle TechPowerUp plutôt qu’un miroir douteux, et de garder son antivirus à jour.

Conséquence concrète : sur de nombreux portables récents – souvent à partir de la 11e génération, ou après certaines mises à jour du BIOS – les curseurs de tension du FIVR sont grisés ou restent sans effet. Si vous appliquez un offset et que la température ne bouge pas d’un degré, c’est probablement que l’undervolt est verrouillé au niveau du microcode. Il n’existe pas de contournement logiciel fiable ; seul un BIOS déverrouillé (parfois proposé par le constructeur, ou via des menus cachés) peut rouvrir l’accès. Vous pouvez consulter la fiche Plundervolt sur Wikipédia pour le détail technique. Dans ce cas, tournez-vous vers les limites de puissance (TPL) et le BD PROCHOT, qui restent souvent accessibles.

## Étapes 1 à 4 : installer ThrottleStop et mesurer la référence

### Étape 1 – Télécharger et extraire ThrottleStop

Rendez-vous sur la page officielle TechPowerUp et téléchargez ThrottleStop 9.7.3, une archive ZIP mise en ligne en avril 2025 et compatible Windows 7 à 11 ; la page officielle affiche désormais une empreinte SHA256 (7B5E…7098A) que vous pouvez comparer après téléchargement pour vérifier que le fichier n’a pas été altéré. Si un miroir vous propose plutôt l’ancienne archive 9.7 du 26 décembre 2024, l’empreinte à contrôler change : SHA256 d83d…800d, selon les fiches TechPowerUp et Uptodown. C’est une archive ZIP : il n’y a pas d’installation. Extrayez-la dans un dossier permanent, par exemple *C:\ThrottleStop*. Évitez le Bureau ou le dossier Téléchargements, car vous pointerez plus tard le démarrage automatique vers ce chemin. Lancez *ThrottleStop.exe* ; au premier démarrage, un avertissement rappelle que l’outil est destiné aux utilisateurs avertis – acceptez pour continuer.

### Étape 2 – Comprendre l’interface principale

La fenêtre principale se lit en trois zones. À gauche, les cases de contrôle globales (Speed Shift EPP, Disable Turbo, BD PROCHOT). Au centre, les profils (numérotés de 1 à 4) et les boutons FIVR, TPL et Options. À droite, la lecture en direct de la fréquence, de la température (C0 %) et de l’état du throttling. En bas, la barre d’état affiche des drapeaux comme PROCHOT, POWER ou THERMAL quand un bridage se produit – c’est votre tableau de bord.

### Étape 3 – Mesurer la température de référence avec le TS Bench

Avant tout réglage, mesurez l’état actuel. Cliquez sur **TS Bench**, choisissez la taille 960M et lancez le test. Notez la température maximale, la fréquence moyenne et surtout la présence de drapeaux THERMAL ou POWER. C’est votre point de comparaison. Voici un exemple de relevé typique sur un portable gaming bridé thermiquement :

```
=== TS Bench 960M – AVANT undervolt ===
Processeur : Intel Core i7 (13e gen, TDP 45 W)
Duree du test        : 00:00:58
Temperature max      : 100 C  <-- Tjmax atteint
Frequence moyenne    : 3,12 GHz (bridee)
Puissance package    : 45 W (PL1)
Drapeaux actifs      : THERMAL, PROCHOT
Erreurs de calcul    : 0
Verdict              : bridage thermique confirme
```
### Étape 4 – Ouvrir la fenêtre FIVR

Cliquez sur le bouton **FIVR** pour ouvrir la fenêtre de contrôle de tension. Si les curseurs sont grisés ou inaccessibles, votre undervolt est verrouillé (relisez la section Plundervolt) : passez directement aux limites de puissance. S’ils sont actifs, félicitations, vous pouvez undervolter. Sélectionnez d’abord **CPU Core** dans la liste « FIVR Control », puis cochez *Unlock Adjustable Voltage* : c’est la case qui déverrouille le curseur d’offset.

## Étapes 5 à 8 : undervolter le CPU Core et le CPU Cache

### Étape 5 – Appliquer un premier offset sur le CPU Core

Toujours dans FIVR, avec CPU Core sélectionné et *Unlock Adjustable Voltage* coché, descendez le curseur **Offset Voltage** à -80 mV. C’est une valeur de départ prudente que la quasi-totalité des puces encaissent sans broncher. Ne partez jamais directement à -150 mV : vous risqueriez un plantage immédiat sans savoir à quel palier votre CPU décroche. La fenêtre FIVR doit ressembler à ceci :

```
FIVR Control : CPU Core
[x] Unlock Adjustable Voltage
Offset Voltage      : -80,1 mV
[x] CPU Core         (rail actif)
Default VID         : 1,250 V
Idle / Turbo VID    : suivent la courbe -80 mV
[ ] CPU Cache        (a regler a l'etape 6)
```
### Étape 6 – Appliquer le même offset sur le CPU Cache

Sélectionnez maintenant **CPU Cache** dans la liste FIVR, cochez *Unlock Adjustable Voltage* et appliquez exactement le même offset : -80 mV. Sur la plupart des processeurs Intel, les rails Core et Cache sont électriquement liés : un déséquilibre entre les deux est l’une des causes les plus fréquentes d’instabilité. Gardez-les toujours identiques, sauf indication contraire propre à votre modèle de puce.

### Étape 7 – Valider avec « OK – Save voltages immediately »

