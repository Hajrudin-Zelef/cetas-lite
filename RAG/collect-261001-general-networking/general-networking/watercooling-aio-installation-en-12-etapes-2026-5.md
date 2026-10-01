---
id: collect-261001-general-networking/general-networking/watercooling-aio-installation-en-12-etapes-2026-5
title: "Exemple de script de test (a adapter selon vos outils installes)"
domain: general-networking
role: reference
task: reference
actors: ["AMD", "Intel"]
dates: []
keywords: ["amd", "arr", "gpu", "intel"]
source: docs/RAG/collect-261001-general-networking/watercooling-aio-installation-en-12-etapes-2026.md
source_anchor: ""
source_lines: [194, 253]
sha256: e023f311d7a7e04df26043e90799901af66b1612165325686d23a8d1539ef7ee
---

# Exemple de script de test (a adapter selon vos outils installes)

- **Bruit de gargouillis ou de pompe qui grince :** presque toujours une bulle d’air résiduelle du transport. Redémarrez le PC et laissez-le tourner une dizaine de minutes, le bruit disparaît généralement de lui-même une fois l’air remonté et dissipé.
- **Températures élevées au repos malgré un AIO neuf :** contact thermique imparfait entre le waterblock et l’IHS. Démontez, nettoyez à l’alcool isopropylique, réappliquez une fine couche de pâte thermique et resserrez en croix.
- **Les ventilateurs du radiateur ne tournent pas du tout :** câble mal enfoncé ou branché sur un en-tête désactivé dans le BIOS. Vérifiez le connecteur et l’en-tête utilisé dans le Hardware Monitor du BIOS.
- **Le RGB ne s’allume pas :** câble ARGB inversé ou branché sur le mauvais type d’en-tête. Vérifiez le détrompeur du connecteur et le type exact d’en-tête (3 broches 5 V contre 4 broches 12 V).
- **Le logiciel de contrôle ne détecte pas l’AIO :** pilote USB interne manquant ou service Windows arrêté. Redémarrez le logiciel en tant qu’administrateur, ou réinstallez-le après un redémarrage complet du PC.
- **Fuite de liquide visible autour des raccords :** arrêtez immédiatement le PC et débranchez l’alimentation. Un circuit AIO est scellé en usine et ne se répare pas soi-même, contactez le service après-vente du fabricant pour un remplacement sous garantie.
- **Écran LCD intégré qui clignote ou plante sous charge GPU (cas rapporté sur certains NZXT Kraken Elite) :** il s’agit d’un problème de firmware déjà identifié par certains utilisateurs, une mise à jour du firmware via NZXT CAM corrige généralement le comportement.
- **Températures qui remontent progressivement plusieurs mois après l’installation :** pâte thermique qui commence à sécher ou pompe qui perd en efficacité. Surveillez l’évolution via HWiNFO64 sur la durée et prévoyez un contrôle si l’écart dépasse 10°C par rapport aux valeurs initiales.
- **Le PC ne démarre plus juste après le montage :** vérifiez en priorité les connecteurs d’alimentation principaux (24 broches carte mère et EPS 12V du CPU), souvent débranchés par erreur pendant la manipulation autour du socket.

## Astuces avancées pour aller plus loin

Une fois l’installation de base validée, plusieurs réglages permettent d’aller chercher les derniers degrés ou les derniers décibels. Combiner un watercooling AIO avec un undervolt CPU reste le geste le plus rentable : en réduisant la tension du processeur à fréquence égale, vous baissez la chaleur générée à la source, ce qui laisse au radiateur une marge supplémentaire pour tourner plus doucement. Notre guide sur ThrottleStop, bien que pensé pour les PC portables, explique la logique de l’undervolt qui s’applique tout aussi bien côté desktop via Ryzen Master ou l’Intel XTU.

Pour aller plus loin sur le monitoring, un script PowerShell simple permet de journaliser la température CPU à intervalle régulier, utile pour repérer une dérive lente sur plusieurs semaines d’usage.

Côté acoustique, remplacer les ventilateurs fournis par des modèles à courbe de pression statique plus douce (be quiet! Silent Wings ou Noctua NF-A12x25 par exemple) réduit sensiblement le bruit perçu à charge moyenne, au prix d’un investissement supplémentaire de 20 à 30 € par ventilateur. C’est un réglage qui a plus d’impact sur le ressenti au quotidien qu’un gain de quelques degrés supplémentaires, la plupart des utilisateurs étant plus sensibles au bruit qu’à un écart de température de 2-3°C invisible sans thermomètre.

```
# Journalisation simple de la temperature CPU toutes les 5 secondes
# Necessite un capteur expose via WMI (namespace variable selon la carte mere)
while ($true) {
    $date = Get-Date -Format "yyyy-MM-dd HH:mm:ss"
    $temp = Get-WmiObject MSAcpi_ThermalZoneTemperature -Namespace "root/wmi" |
            Select-Object -ExpandProperty CurrentTemperature
    $celsius = ($temp / 10) - 273.15
    Add-Content -Path "temp_log.csv" -Value "$date,$celsius"
    Start-Sleep -Seconds 5
}
```
### Passer au custom loop : la suite logique ?

Pour les passionnés qui trouvent les limites de l’AIO trop contraignantes (radiateur unique, pompe fixe, aucune personnalisation de trajet de tubes), le circuit custom reste l’étape suivante. Il permet d’intégrer un bloc GPU en plus du CPU sur la même boucle, de choisir librement la taille et le nombre de radiateurs, et d’opter pour des raccords rapides qui simplifient la maintenance. La contrepartie est réelle : un budget qui grimpe facilement au-delà de 300-400 € rien que pour les composants du circuit, un entretien régulier du liquide (purge, remplacement tous les 6 à 12 mois selon le liquide utilisé) et un risque de fuite non nul en cas d’erreur de montage. Pour une première expérience de refroidissement liquide, l’AIO reste très largement recommandé.

## Projet complet : notre configuration testée de A à Z

Pour illustrer l’ensemble de la procédure sur un cas concret, voici la configuration montée pour la rédaction de ce guide, avec le détail des pièces et les résultats obtenus après installation.

- **CPU :** AMD Ryzen 7 9800X3D (socket AM5, TDP configurable jusqu’à 120-142 W selon profil PBO)
- **Carte mère :** chipset X870, en-têtes AIO_PUMP et 3 ports SYS_FAN disponibles
- **GPU :** une RTX série 50 milieu-haut de gamme, refroidissement d’origine conservé
- **Watercooling AIO :** ARCTIC Liquid Freezer III Pro 360, monté en façade avant en intake
- **Boîtier :** tour moyenne avec façade mesh, deux ventilateurs d’origine en exhaust arrière et supérieur
- **Pâte thermique :** celle pré-appliquée sur le bloc pompe ARCTIC, aucune application manuelle nécessaire
- **Temps total de montage :** 40 minutes, boîtier déjà ouvert, en partant d’un ventirad démonté

Résultat après un stress test Cinebench de 20 minutes suivi d’une session de jeu de deux heures : température CPU package stabilisée à 71-74°C en charge soutenue, contre 89-92°C avec le ventirad simple tour d’origine sur le même boîtier et la même pièce. Les ventilateurs du radiateur ne dépassent 1400 RPM qu’en charge lourde, restant sous les 700 RPM au repos, ce qui rend la configuration quasiment silencieuse en usage bureautique. Le coût total de l’opération s’est limité au prix de l’AIO, sans achat de pâte thermique ni d’outillage supplémentaire.

Le câblage a demandé un peu plus de soin que le montage lui-même. Avec la pompe, les trois ventilateurs du radiateur et le câble ARGB à faire cheminer jusqu’à la carte mère, mieux vaut passer les câbles à l’arrière du boîtier dès le début plutôt que de tout tirer au dernier moment. Un serre-câble tous les 10-15 cm suffit à garder un rendu propre sans bloquer le retour du panneau latéral, un détail qui ne change rien aux températures mais qui évite de devoir tout redémonter une semaine plus tard pour ajouter un composant.

Si votre budget GPU fait aussi partie de vos arbitrages pour ce genre de montage, notre suivi des prix des cartes graphiques en 2026 aide à situer où investir en priorité entre carte graphique et refroidissement.

## Foire aux questions

**Un watercooling AIO est-il plus fiable qu’un ventirad ?**

Non, pas au sens strict. Un ventirad n’a ni pompe ni liquide, donc mécaniquement moins de pièces susceptibles de tomber en panne. Un AIO bien installé reste toutefois fiable sur plusieurs années, la plupart des fabricants proposant des garanties de 5 à 6 ans sur leurs modèles récents.

**Combien de temps dure un watercooling AIO avant de devoir le remplacer ?**

