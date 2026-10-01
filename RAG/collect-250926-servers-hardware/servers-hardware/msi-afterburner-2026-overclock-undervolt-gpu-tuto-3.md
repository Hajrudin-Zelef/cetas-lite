---
id: collect-250926-servers-hardware/servers-hardware/msi-afterburner-2026-overclock-undervolt-gpu-tuto-3
title: "Cartes NVIDIA : lire le modèle, la conso et les températures"
domain: servers-hardware
role: reference
task: reference
actors: ["Intel", "Nvidia"]
dates: []
keywords: ["nvidia", "benchmark", "benchmarks", "blackwell", "gpu", "intel", "memory"]
source: docs/RAG/clean4/msi-afterburner-2026-overclock-undervolt-gpu-tuto.md
source_anchor: ""
source_lines: [116, 196]
sha256: 5118abee987a525592ef28dcacd1e1217bc834b953db3aaf042273aec050ff23
---

# Cartes NVIDIA : lire le modèle, la conso et les températures

```
# Plan d'overclocking incremental, a tester palier par palier
# (valeurs indicatives, chaque exemplaire de carte differe)
Etape A  Power Limit  -> maximum disponible       [appliquer + tester]
Etape B  Core Clock   -> +50, +75, +100, +125 MHz [+1 palier / 15 min]
Etape C  Memory Clock -> +200, +400, +600, +800   [+1 palier / 15 min]
Etape D  Validation   -> benchmark 1 h + 3 jeux reels
# Si plantage / artefacts -> revenir au dernier palier stable -50 MHz
```
| Carte | Offset cœur (départ) | Offset mémoire (départ) | Limite puissance | Remarque | 
|---|---|---|---|---|
| RTX 5090 | +100 à +150 MHz | +1000 à +2000 MHz | Max (575 W TGP) | Privilégier l’undervolt | 
| RTX 5080 | +100 à +200 MHz | +1500 à +3000 MHz | Max | GDDR7 très tolérante | 
| RTX 5070 Ti | +100 à +200 MHz | +1000 à +2000 MHz | Max | Bon rendement | 
| RX 9070 XT | +50 à +150 MHz | Marge limitée (RDNA 4) | Max | Undervolt très efficace | 
| Intel Arc B580 | Via Power + Voltage | Marge modérée | Max | Pilotes en évolution | 

## Étapes 11 à 12 : Undervolter avec l’éditeur de courbe (Ctrl+F)

L’undervolting est, pour beaucoup de joueurs européens, le réglage le plus rentable de tout ce tutoriel. Plutôt que de chercher plus de puissance, on cherche le **meilleur rendement** : la même performance avec moins de tension, donc moins de chaleur, moins de bruit et moins de watts. C’est aussi le réglage le plus sûr, car on réduit la tension au lieu de l’augmenter.

**Étape 11 – Ouvrir l’éditeur de courbe.** Appuyez sur **Ctrl+F** dans MSI Afterburner pour ouvrir l’éditeur de courbe tension/fréquence (Voltage/Frequency Curve Editor). Chaque point relie une tension (axe horizontal, en mV) à une fréquence (axe vertical, en MHz). L’objectif : forcer le GPU à atteindre sa fréquence cible à une tension plus basse que celle d’usine.

**Étape 12 – Aplatir la courbe.** Choisissez une tension cible (par exemple 900 mV pour une RTX 50). Cliquez sur le point correspondant à cette tension, montez-le à la fréquence souhaitée (par exemple 2700 MHz), puis maintenez **Maj (Shift)** et sélectionnez tous les points situés à droite pour les aplatir au même niveau. Appliquez avec la coche. Le GPU ne dépassera plus cette tension. Testez ensuite la stabilité : des artefacts ou un plantage signifient que la tension est trop basse pour la fréquence visée – remontez de 15 à 25 mV.

```
# Exemple de courbe undervolt sur une RTX 50 (a adapter a VOTRE carte)
# Tension (mV)  ->  Frequence cible (MHz)
875 mV  ->  2580 MHz   # profil "silence / efficacite"
900 mV  ->  2700 MHz   # profil "equilibre" (recommande)
925 mV  ->  2790 MHz   # profil "performance"
# Methode :
# 1) Ctrl+F pour ouvrir l'editeur
# 2) Cliquer-glisser le point de tension cible a la frequence voulue
# 3) Shift + selectionner les points a droite -> aplatir
# 4) Appliquer (coche) puis tester 30 min
# 5) Instable ? remonter la tension de 15-25 mV
```
Un undervolt bien réglé combiné à une courbe de ventilation adaptée peut faire **chuter la température de 10 à 20 °C** et réduire la consommation de plusieurs dizaines de watts, sans perte de performance perceptible. Sur une carte qui chauffe l’été, c’est la différence entre un PC silencieux et une turbine. Si vous comparez le coût d’usage d’un PC à celui du jeu dématérialisé, notre dossier meilleur cloud gaming 2026 chiffre l’alternative sans GPU local.

## Étapes 13 à 14 : Courbe de ventilation et profils automatiques

**Étape 13 – Créer une courbe de ventilation personnalisée.** Dans les paramètres, onglet *Fan*, cochez « Enable user defined software automatic fan control ». Vous obtenez un graphique reliant la température (axe horizontal) au régime des ventilateurs (axe vertical). Tracez une courbe progressive : ventilateurs silencieux en dessous de 50 °C, montée régulière jusqu’à 100 % vers 83-85 °C. Sur les RTX 5090 et 5080, MSI Afterburner 4.6.6 permet de piloter **jusqu’à quatre ventilateurs indépendamment**, voire de manière asynchrone pour cibler les points chauds.

**Étape 14 – Enregistrer et automatiser les profils.** Une fois vos réglages validés, cliquez sur l’icône de sauvegarde puis sur un emplacement numéroté (1 à 5) pour mémoriser le profil. Activez « Apply overclocking at system startup » dans les paramètres pour charger automatiquement votre profil au démarrage de Windows. Vous pouvez aussi associer des profils différents à des jeux précis via RTSS, ou lancer un profil avec un raccourci. Pour les utilisateurs avancés, un script de démarrage permet de forcer un profil :

```
# Lancer MSI Afterburner et appliquer le profil n2 au demarrage (.bat)
@echo off
cd /d "C:\Program Files (x86)\MSI Afterburner"
start "" "MSIAfterburner.exe" -Profile2
# -Profile1 a -Profile5 : charge le profil correspondant
# Placer le raccourci dans le dossier de demarrage :
# Win+R -> shell:startup -> y deposer le raccourci
```
Astuce : créez deux profils complémentaires – un profil « performance » (overclock) pour les jeux exigeants et un profil « silence » (undervolt agressif + ventilateurs bas) pour les sessions de soirée ou les jeux légers. Vous basculerez de l’un à l’autre en un clic.

## Tester la stabilité : Kombustor, Unigine et benchmarks

Un overclock ou un undervolt n’est validé qu’après des tests de stabilité sérieux. Un réglage peut sembler stable cinq minutes puis planter au bout d’une heure de jeu. La méthode fiable combine **tests synthétiques** (charge maximale constante) et **jeux réels** (charge variable, plus représentative). Comptez au minimum une heure de test continu avant de considérer un profil comme stable.

| Outil | Type | Durée conseillée | Ce qu’il révèle | 
|---|---|---|---|
| MSI Kombustor | Charge maximale (FurMark) | 15-30 min | Stabilité thermique, throttling | 
| Unigine Superposition | Benchmark répétable | 3-5 boucles | Artefacts, score, régularité | 
| 3DMark Steel Nomad | Test de stress | 20 boucles | Stabilité sur la durée (%) | 
| Jeu réel exigeant | Charge variable | 1-2 h | Plantages en conditions réelles | 
| OCCT | Test GPU + détection d’erreurs | 30-60 min | Erreurs mémoire silencieuses | 

Interprétez les signaux : des **artefacts graphiques** (points, lignes, textures corrompues) indiquent une fréquence trop haute ou une tension trop basse. Un **écran noir suivi d’une récupération du pilote** signale une instabilité du cœur. Une **baisse de score** malgré une fréquence plus élevée trahit une mémoire au-delà de sa limite (correction d’erreurs active). Dans tous les cas, reculez d’un palier. La plupart des benchmarks se pilotent aussi en ligne de commande pour automatiser une boucle de test :

```
# Boucle de test de charge GPU avec MSI Kombustor (exemple)
MSI-Kombustor-x64.exe /benchmark /width=2560 /height=1440 /fullscreen /duration=900000
# /duration en millisecondes (900000 = 15 min)
# Surveiller en parallele : Hot Spot < 90 degC, aucun artefact
```
## Spécificités RTX 50 (Blackwell) et Radeon RX 9000 (RDNA 4) en 2026

La génération 2026 change la donne pour MSI Afterburner. Côté NVIDIA, les **GeForce RTX 50 « Blackwell »** embarquent de la mémoire GDDR7 dont le potentiel d'overclocking est exceptionnel. La mise à jour de MSI Afterburner permet de pousser la mémoire jusqu'à 36 Gbit/s et étend la plage de fréquence à +3000 MHz d'offset. Particularité notable : les puces des RTX 5080 sont des modèles 32 Gbit/s volontairement bridés à 30 Gbit/s d'usine, d'où une marge confortable. La RTX 5090, avec son enveloppe thermique de **575 W**, est la candidate idéale à l'undervolting : on peut souvent lui retirer 80 à 100 W sans perte de performance.

