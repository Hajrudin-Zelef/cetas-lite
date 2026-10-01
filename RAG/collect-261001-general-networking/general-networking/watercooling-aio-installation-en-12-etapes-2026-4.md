---
id: collect-261001-general-networking/general-networking/watercooling-aio-installation-en-12-etapes-2026-4
title: "Exemple de script de test (a adapter selon vos outils installes)"
domain: general-networking
role: reference
task: reference
actors: ["AMD"]
dates: []
keywords: ["amd", "attention"]
source: docs/RAG/collect-261001-general-networking/watercooling-aio-installation-en-12-etapes-2026.md
source_anchor: ""
source_lines: [111, 193]
sha256: 48cec778a016cb9a45fb72cd9b0627ba1084ec9cdb4f1211794f7493f7e03a50
---

# Exemple de script de test (a adapter selon vos outils installes)

Dans le BIOS, cherchez la section généralement nommée Q-Fan Control (ASUS), Smart Fan (MSI/Gigabyte) ou Fan Control selon la marque de carte mère. Définissez une courbe progressive plutôt qu’un mode fixe, avec des paliers proches de ceux-ci pour un usage gaming classique.

```
Courbe ventilateurs radiateur (exemple, a adapter a votre CPU)
Temperature CPU  ->  Vitesse ventilateurs
  40C            ->  30 % (silencieux)
  55C            ->  45 %
  65C            ->  65 %
  75C            ->  85 %
  85C et plus     ->  100 % (securite)
Courbe pompe (recommandation generale)
  Idle             ->  60-70 % minimum (jamais en dessous, la pompe ne doit pas ralentir au repos)
  Charge           ->  100 %
```
La règle la plus importante ici concerne la pompe : contrairement aux ventilateurs, elle ne doit jamais descendre sous un seuil minimal recommandé par le fabricant (souvent 60 à 70 % de sa vitesse maximale), même au repos. Une pompe qui ralentit trop réduit la circulation du liquide et peut provoquer des pics de température CPU en quelques secondes lors d’un pic de charge soudain, typiquement au lancement d’un jeu.

Une fois dans Windows, le logiciel du fabricant prend le relais avec plus de finesse. Corsair iCUE gère le câblage caché des kits LINK et synchronise le RGB sur tout l’écosystème, NZXT CAM pilote l’écran LCD du Kraken Elite en plus des courbes classiques, MSI Center centralise le contrôle des AIO MSI avec le RGB Mystic Light. Voici à quoi ressemble un export de profil de courbe typique dans ce genre de logiciel.

```
{
  "profil": "Gaming silencieux",
  "ventilateurs_radiateur": [
    {"temp_c": 30, "duty_pct": 25},
    {"temp_c": 45, "duty_pct": 40},
    {"temp_c": 60, "duty_pct": 60},
    {"temp_c": 75, "duty_pct": 85},
    {"temp_c": 85, "duty_pct": 100}
  ],
  "pompe": {"mode": "performance", "duty_min_pct": 65}
}
```
Si vous possédez un CPU AMD, pensez à combiner ces réglages avec un passage par Ryzen Master pour ajuster les courbes de tension du processeur, une meilleure évacuation thermique n’a d’intérêt que si le CPU en profite réellement pour tenir ses boosts plus longtemps.

## Tester et valider l’installation : stress test et monitoring des températures

Ne considérez jamais une installation de watercooling AIO comme terminée sans un test de charge. C’est le seul moyen de détecter un mauvais contact thermique, un tilt du waterblock ou une pompe mal branchée avant que le problème ne s’aggrave sur la durée.

Lancez d’abord HWiNFO64 en arrière-plan pour enregistrer les températures, puis démarrez un test de charge CPU de 15 à 30 minutes. Voici un exemple de séquence simple pour qui préfère lancer ses tests en ligne de commande plutôt que via une interface graphique.

```
# Exemple de script de test (a adapter selon vos outils installes)
# Lance un test de charge de 20 minutes puis note l'heure de fin
echo "Debut du stress test : %date% %time%" >> stresstest_log.txt
start "" "C:\Program Files\Cinebench 2026\Cinebench.exe" -b
timeout /t 1200
echo "Fin du stress test : %date% %time%" >> stresstest_log.txt
```
Pendant ce test, gardez un œil sur trois valeurs dans HWiNFO64 : la température CPU package, la vitesse de la pompe (RPM) et la vitesse des ventilateurs du radiateur. Voici un exemple de ce que vous devriez observer sur une configuration saine avec un 360 mm bien installé.

```
Exemple de releve HWiNFO64 - Ryzen 7 9800X3D + AIO 360 mm
                          Idle        Charge (20 min Cinebench)
CPU Package Temp          38C         72C
CPU Package Power         18 W        135 W
Pompe (RPM)                1450 RPM    2600 RPM
Ventilateurs radiateur      650 RPM    1550 RPM
Frequence CPU (moy.)       3.8 GHz     5.1 GHz
```
Au-delà de 85-90°C en charge soutenue sur un CPU moderne, c’est le signe qu’il faut revoir l’installation avant toute chose : reserrage du waterblock, vérification du branchement de la pompe, ou tout simplement re-application de la pâte thermique. En dessous de 80°C sur un stress test de 20 minutes avec un 360 mm, l’installation peut être considérée comme validée pour un usage gaming normal.

## Watercooling AIO vs ventirad : lequel choisir en 2026

La question revient à chaque montage de PC, et la réponse a évolué avec la baisse des prix des AIO d’entrée de gamme. Un Thermalright Aqua Elite 240 V3 à 44,90 $ coûte désormais moins cher que certains ventirads double tour premium, ce qui change la donne pour un budget serré.

**Quand un ventirad suffit encore largement.** Un bon ventirad reste imbattable sur trois points : aucune pompe ne peut tomber en panne, aucun liquide ne peut fuir, et la durée de vie dépasse largement celle d’un AIO dans la majorité des cas. Pour un usage bureautique ou un CPU milieu de gamme sans overclocking, un ventirad simple tour à 30-40 € couvre largement les besoins sans aucune des contraintes de montage propres au watercooling.

À l’inverse, le watercooling AIO prend l’avantage dès que le TDP dépasse 200-250 W en charge soutenue, ou dès que l’esthétique du boîtier entre en jeu : un radiateur avec ventilateurs RGB en façade transforme visuellement une configuration qu’un ventirad, même performant, ne peut pas égaler dans un boîtier à panneau vitré. Le compromis reste celui du risque : un circuit scellé peut, dans de rares cas, présenter un défaut de pompe ou une fuite après plusieurs années d’usage, quand un ventirad ne tombe quasiment jamais en panne au sens mécanique du terme.

## Les erreurs les plus fréquentes à éviter lors du montage

Ces erreurs reviennent le plus souvent dans les retours d’expérience et les forums d’entraide matériel. Toutes sont évitables avec un peu d’attention au moment du montage.

- **Ventilateurs montés dans le mauvais sens.** Une flèche imprimée sur le cadre du ventilateur indique le sens du flux d’air, l’ignorer inverse tout le refroidissement du radiateur sans qu’aucun message d’erreur ne vous prévienne.
- **Pompe branchée sur le mauvais en-tête.** Un en-tête SYS_FAN générique peut réduire la vitesse de la pompe en veille selon la courbe définie, alors qu’elle doit rester proche de son régime nominal en permanence. Utilisez toujours l’en-tête AIO_PUMP ou CPU_FAN dédié.
- **Film plastique de protection oublié sous le bloc pompe.** Certains fabricants protègent la plaque de contact en cuivre d’un film avant expédition, retirez-le systématiquement avant de poser le bloc sur le CPU, sans quoi le contact thermique est nul.
- **Serrage inégal des vis du waterblock.** Un vissage en diagonale mal réparti incline légèrement le bloc pompe sur l’IHS, créant un point chaud localisé qui fausse les lectures de température d’un seul cœur ou d’un seul côté du die.
- **Confusion entre en-têtes ARGB et RGB.** Un connecteur ARGB 3 broches (5 V) branché par erreur sur un en-tête RGB 4 broches (12 V), ou l’inverse, peut endommager définitivement le contrôleur LED du kit.
- **Radiateur monté en toit sur un boîtier d’entrée de gamme sans vérifier le poids supporté.** Un 360 mm avec ses trois ventilateurs pèse son poids, un montage en toit sur une tôle trop fine peut, à terme, provoquer un léger fléchissement visible du panneau supérieur.

## Dépannage : les problèmes les plus courants et leurs solutions

Voici les pannes et symptômes les plus signalés après l’installation d’un watercooling AIO, avec la cause probable et la marche à suivre pour chacun.

