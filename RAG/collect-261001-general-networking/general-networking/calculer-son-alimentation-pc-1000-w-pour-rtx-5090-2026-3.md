---
id: collect-261001-general-networking/general-networking/calculer-son-alimentation-pc-1000-w-pour-rtx-5090-2026-3
title: "Exemple : RTX 5090 + Ryzen 9 9950X3D"
domain: general-networking
role: reference
task: reference
actors: ["Nvidia"]
dates: []
keywords: ["benchmark", "gpu", "nvidia"]
source: docs/RAG/collect-261001-general-networking/calculer-son-alimentation-pc-1000-w-pour-rtx-5090-2026.md
source_anchor: ""
source_lines: [108, 184]
sha256: 9979ead359273121b445702332a27411020319ca9f94952089a78ad9084eb450
---

# Exemple : RTX 5090 + Ryzen 9 9950X3D

**Étape 7 : comprenez la différence entre single rail et multi rail.** Le rail 12V est le circuit qui alimente le GPU, le CPU et le stockage. Une alimentation single rail regroupe toute la puissance 12V sur un seul circuit protégé par une seule limite de courant (OCP), ce qui simplifie le câblage et convient bien aux GPU très gourmands comme la RTX 5090. Une alimentation multi rail répartit le 12V sur plusieurs circuits indépendants, chacun avec sa propre protection, ce qui coupe plus vite en cas de court-circuit localisé mais impose de bien répartir les connecteurs entre les rails.

**Étape 8 : pour un PC gamer avec une seule carte graphique** (le cas de la grande majorité des configurations en 2026), une alimentation single rail 12V est généralement recommandée. Elle évite les erreurs de répartition et garantit que le connecteur 12V-2×6 peut recevoir toute la puissance nécessaire sans être limité artificiellement.

**Étape 9 : sélectionnez un modèle full-modulaire, ATX 3.1, avec au moins 10 ans de garantie.** Le câblage full-modulaire permet de ne brancher que les câbles réellement utilisés, ce qui améliore l’aération du boîtier. Une garantie longue (Corsair, Seasonic et be quiet! proposent généralement 10 ans sur leurs gammes Gold et supérieures) est un bon indicateur de la qualité des composants internes.

## Étape 10 à 12 : installer, câbler et tester votre alimentation

**Étape 10 : installez physiquement le bloc.** Éteignez et débranchez le PC, patientez 60 secondes pour laisser les condensateurs se décharger. Glissez l’alimentation dans son emplacement (généralement en bas du boîtier), ventilateur orienté vers le bas si votre boîtier a une aération dédiée, ou vers le haut sinon. Fixez avec les 4 vis fournies.

**Étape 11 : câblez dans cet ordre** pour éviter les oublis : connecteur ATX 24 broches sur la carte mère, connecteur(s) EPS 4+4 ou 8 broches pour le CPU, connecteur 12V-2×6 pour le GPU (ou PCIe 8 broches selon votre carte), puis les câbles SATA/périphériques pour le stockage et les ventilateurs. Vérifiez chaque connecteur à la main : vous devez sentir un clic net.

**Étape 12 : premier démarrage et test de stabilité.** Rebranchez le secteur, allumez l’interrupteur à l’arrière du PSU, puis démarrez. Entrez dans le BIOS pour vérifier que les tensions 12V, 5V et 3,3V affichées sont dans une marge de ±5 %. Lancez ensuite un stress test GPU/CPU de 15 minutes pour confirmer qu’aucun redémarrage ni coupure ne se produit sous charge maximale simultanée.

## Automatiser le calcul avec un script Python

Plutôt que de refaire le calcul à la main à chaque changement de composant, voici un petit script Python qui automatise la méthode décrite plus haut, marge de sécurité comprise.

```
calculateur_watts.py
def calculer_psu(gpu_tgp, cpu_watts, plateforme_watts=120, marge=0.35):
    total_base = gpu_tgp + cpu_watts + plateforme_watts
    total_avec_marge = total_base * (1 + marge)
    return round(total_base), round(total_avec_marge)
# Exemple : RTX 5090 + Ryzen 9 9950X3D
base, recommande = calculer_psu(gpu_tgp=575, cpu_watts=230, plateforme_watts=140)
print(f"Consommation de base : {base} W")
print(f"PSU recommandé (marge 35%) : {recommande} W")
# Arrondi à la puissance commerciale la plus proche
puissances_commerciales = [550, 650, 750, 850, 1000, 1200, 1600]
choix = next(p for p in puissances_commerciales if p >= recommande)
print(f"Modèle à choisir : {choix} W minimum")
```
Voici le résultat obtenu en exécutant ce script pour une configuration RTX 5090 :

```
$ python3 calculateur_watts.py
Consommation de base : 945 W
PSU recommandé (marge 35%) : 1276 W
Modèle à choisir : 1600 W minimum
```
Ce résultat peut surprendre, mais il illustre bien pourquoi les guides constructeurs recommandent 1000 à 1200 W pour une RTX 5090 : la marge de 35 % appliquée à une base déjà élevée pousse rapidement vers le palier supérieur. Vous pouvez ajuster la variable `marge` à 0,25 si vous ne comptez pas overclocker, ou la laisser à 0,35-0,40 si vous voulez une marge confortable pour l’avenir.

Ce script reste volontairement simple, mais vous pouvez l’étendre facilement. Par exemple, ajoutez un paramètre `nb_ssd` qui multiplie le nombre de SSD NVMe par 6-8 W chacun, ou un paramètre `overclock` qui augmente la marge de 10 points supplémentaires si vous prévoyez de pousser le GPU au-delà de ses réglages d’usine. Si vous gérez plusieurs configurations (un PC personnel et un PC pour un proche, par exemple), transformez la fonction en une petite bibliothèque réutilisable avec un dictionnaire de profils GPU prédéfinis, en reprenant les valeurs du tableau TGP présenté plus haut dans cet article.

## Vérifier la consommation réelle avec nvidia-smi et HWiNFO

Le calcul théorique est une base de départ, mais rien ne remplace une mesure réelle une fois le PC monté. Sur une carte NVIDIA, la commande `nvidia-smi` permet d’interroger la consommation instantanée du GPU directement depuis un terminal.

`nvidia-smi --query-gpu=power.draw,power.limit,temperature.gpu --format=csv -l 1`
Cette commande affiche, chaque seconde, la puissance instantanée consommée par le GPU, sa limite de puissance configurée et sa température. Voici un exemple de sortie pendant une session de jeu sur une RTX 5090 :

```
power.draw [W], power.limit [W], temperature.gpu
542.18 W, 600.00 W, 68
571.90 W, 600.00 W, 70
563.44 W, 600.00 W, 71
575.02 W, 600.00 W, 71
```
Pour une vue globale du système (CPU, carte mère, disques), installez HWiNFO64 en mode capteurs seuls, activez la journalisation CSV via le menu *Settings > Logging*, puis lancez un benchmark de 10 minutes. Ouvrez ensuite le fichier CSV généré dans un tableur pour repérer le pic maximal sur la colonne “CPU Package Power” et “GPU Power” : additionnez ces deux valeurs de pic (pas les moyennes) pour valider que votre alimentation dispose d’une marge suffisante par rapport au pic mesuré réel, et non uniquement par rapport au TGP théorique.

## Projet complet : budget d’alimentation pour un PC RTX 5090

Pour rendre cette méthode concrète, voici un exemple de configuration complète haut de gamme avec le détail du calcul, du choix de l’alimentation jusqu’au budget final.

- **GPU** : RTX 5090, TGP 575 W
- **CPU** : Ryzen 9 9950X3D, environ 230 W en charge avec PBO
- **Plateforme** : carte mère X870E, 64 Go de DDR5, 2 SSD NVMe PCIe 5.0, watercooling AIO 360 mm, 6 ventilateurs, RGB : environ 140 W
- **Consommation de base cumulée** : 575 + 230 + 140 = 945 W
- **Marge de sécurité appliquée (35 %)** : 945 x 1,35 = 1276 W
- **Choix final** : alimentation 1200 W, 80 PLUS Platinum, ATX 3.1, connecteur 12V-2×6 natif, full-modulaire
- **Budget alimentation** : environ 220 à 280 € pour ce palier de puissance et de certification en France en 2026

Ce budget alimentation représente environ 6 à 8 % du coût total d’un PC gamer haut de gamme équipé d’une RTX 5090, une proportion qu’il ne faut surtout pas rogner : une alimentation sous-dimensionnée ou d’entrée de gamme sur ce type de configuration est la cause la plus fréquente d’instabilité et, dans les cas les plus graves, de dommages matériels.

Pour comparer avec un budget plus mesuré, voici le même calcul appliqué à une configuration milieu-haut de gamme, très représentative des PC gamer vendus en France en 2026.

