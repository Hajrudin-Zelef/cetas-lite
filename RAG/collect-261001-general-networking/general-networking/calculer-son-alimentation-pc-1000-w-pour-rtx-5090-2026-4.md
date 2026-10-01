---
id: collect-261001-general-networking/general-networking/calculer-son-alimentation-pc-1000-w-pour-rtx-5090-2026-4
title: "Exemple : RTX 5090 + Ryzen 9 9950X3D"
domain: general-networking
role: reference
task: reference
actors: ["Nvidia"]
dates: []
keywords: ["acquisition", "benchmark", "gpu", "nvidia"]
source: docs/RAG/collect-261001-general-networking/calculer-son-alimentation-pc-1000-w-pour-rtx-5090-2026.md
source_anchor: ""
source_lines: [185, 260]
sha256: 2959c7fa9b883b01ef91e4d129655008665e1cf36874ade06a90a77d999b88c8
---

# Exemple : RTX 5090 + Ryzen 9 9950X3D

- **GPU** : RTX 5070 Ti, TGP 300 W
- **CPU** : Ryzen 7 9800X3D, environ 130 W en charge
- **Plateforme** : carte mère B650, 32 Go de DDR5, 1 SSD NVMe PCIe 4.0, ventirad air, 4 ventilateurs : environ 90 W
- **Consommation de base cumulée** : 300 + 130 + 90 = 520 W
- **Marge de sécurité appliquée (30 %)** : 520 x 1,30 = 676 W
- **Choix final** : alimentation 750 W, 80 PLUS Gold, ATX 3.0 ou 3.1, semi-modulaire ou full-modulaire
- **Budget alimentation** : environ 90 à 120 € pour ce palier de puissance en France en 2026

Cet exemple illustre un point important : pour une configuration mid-range, la marge de sécurité de 30 % suffit généralement, contre 35-40 % recommandé sur les configurations RTX 5090 où les pics transitoires sont proportionnellement plus élevés. Adaptez toujours la marge à la classe de GPU utilisée plutôt que d’appliquer un pourcentage fixe à toutes les configurations.

## Tableau comparatif des prix d’alimentations en France 2026

Voici une fourchette de prix réaliste pour des modèles populaires 850 W, certifiés Gold, compatibles ATX 3.x, disponibles chez les revendeurs français en 2026.

| Modèle | Puissance | Certification | Norme | Prix indicatif | 
|---|---|---|---|---|
| Corsair RM850e | 850 W | 80 PLUS Gold | ATX 3.x | 120 – 150 € | 
| be quiet! Pure Power 12 M 850 W | 850 W | 80 PLUS Gold | ATX 3.0 | 130 – 160 € | 
| Seasonic Focus GX-850 | 850 W | 80 PLUS Gold | ATX 3.x | 120 – 180 € | 
| MSI MPG A850G | 850 W | 80 PLUS Gold | ATX 3.x / PCIe 5 | 130 – 190 € | 

Ces prix sont des ordres de grandeur observés sur les comparateurs français début 2026 : vérifiez toujours le tarif du jour chez votre revendeur habituel avant achat, les promotions ponctuelles peuvent faire varier ces montants de 20 à 30 €.

## 5 erreurs courantes à éviter lors du choix de son alimentation

Voici les pièges les plus fréquents constatés chez les monteurs de PC, débutants comme confirmés.

1. **Se baser sur le TDP du CPU seul et oublier le GPU.** Le GPU représente souvent 50 à 60 % de la consommation totale sous charge, c’est le poste à calculer en priorité.
2. **Ignorer les pics transitoires.** Un calcul sans marge de sécurité (30-40 %) sous-estime systématiquement le besoin réel, en particulier avec les GPU RTX 50.
3. **Utiliser un adaptateur multi-connecteurs 4×8 broches vers 12V-2×6.** Cette pratique concentre plusieurs câbles anciens sur un connecteur récent non conçu pour ce type d’adaptation, l’une des causes principales de connecteurs fondus documentées en 2025-2026.
4. **Choisir une alimentation non modulaire pour économiser 20 €.** Les câbles fixes en surnombre bloquent la circulation d’air, augmentent les températures internes de 3 à 5 °C et compliquent le câblage.
5. **Réutiliser une vieille alimentation de plus de 6-7 ans sur un nouveau GPU haut de gamme.** Les condensateurs vieillissent et la puissance réellement délivrable chute sous charge, même si l’étiquette affiche toujours le wattage d’origine.

## Dépannage : 8 problèmes fréquents et solutions

Si votre PC présente un comportement instable après un changement de composant ou d’alimentation, voici les causes les plus courantes et comment les résoudre. Dans la majorité des cas, un diagnostic méthodique en moins de 15 minutes permet d’identifier la cause exacte sans avoir à démonter l’intégralité de la configuration.

- **Le PC redémarre seul en jeu, jamais au bureau** : quasi toujours un sous-dimensionnement de l’alimentation par rapport aux pics du GPU. Recalculez avec une marge de 40 % et vérifiez les watts réels via`nvidia-smi` .
- **Odeur de plastique brûlé près du GPU** : éteignez immédiatement le PC, débranchez, et inspectez le connecteur 12V-2×6 pour des traces de fonte ou de noircissement. Ne rallumez pas avant d’avoir remplacé le câble.
- **Le PC ne démarre pas du tout, aucun ventilateur ne tourne** : vérifiez l’interrupteur à l’arrière du PSU, testez une autre prise secteur, puis testez le connecteur ATX 24 broches avec la méthode du trombone (PSU débranché du reste du PC).
- **Écran noir avec bips au démarrage** : vérifiez que le connecteur EPS 4+4 broches du CPU est bien enfoncé jusqu’au clic des deux côtés.
- **Le GPU plafonne ses performances en dessous des scores attendus** : ouvrez le BIOS, vérifiez la tension du rail 12V (doit rester au-dessus de 11,4 V sous charge). Une tension qui chute sous 11,2 V indique une alimentation insuffisante.
- **Coupure brutale uniquement pendant un benchmark ou un rendu 3D lourd** : c’est le signe classique d’un déclenchement de la protection OCP. Passez au palier de puissance supérieur.
- **Bruit de “coil whine” (sifflement aigu) sous forte charge** : généralement inoffensif, lié à la qualité des bobines internes. Un modèle Platinum/Titanium de meilleure qualité réduit ce phénomène par rapport à une entrée de gamme Bronze.
- **Ventilateur du PSU qui tourne à plein régime en permanence** : vérifiez que l’alimentation n’est pas positionnée trop près d’une source de chaleur (GPU, watercooling) et que les grilles d’aération du boîtier ne sont pas obstruées par la poussière.

## Astuces avancées pour les configurations haut de gamme

Si vous montez une configuration extrême (RTX 5090, overclocking CPU/GPU, multi-stockage), quelques ajustements supplémentaires permettent de sécuriser davantage votre installation et d’en tirer le meilleur parti sur la durée, au-delà du simple calcul initial du wattage.

- **Limitez le power limit du GPU à 90-95 %** via l’application constructeur si vous êtes proche du palier de puissance supérieur : la perte de performance est généralement inférieure à 3 %, pour une baisse de consommation de pointe non négligeable.
- **Privilégiez les câbles natifs plutôt que les câbles sleevés tiers** pour le connecteur 12V-2×6, sauf si le fabricant du câble alternatif garantit explicitement la compatibilité 600 W et la certification ATX 3.1.
- **Sur une configuration multi-GPU ou avec carte d’acquisition/eGPU** , envisagez une alimentation multi-rail pour isoler les circuits et limiter le risque en cas de défaut sur un composant.
- **Surveillez le vieillissement** : au-delà de 5 ans d’utilisation intensive, la capacité réelle d’une alimentation baisse de 5 à 10 %. Recalculez votre marge de sécurité en conséquence avant d’ajouter un GPU plus gourmand.
- **Pensez à l’onduleur (UPS)** pour les configurations dépassant 1000 W : une micro-coupure secteur sur ce type de système peut endommager des composants coûteux, un onduleur avec régulation de tension ligne interactive protège l’ensemble.

## FAQ : questions fréquentes sur le calcul de l’alimentation PC

**Quelle alimentation pour une RTX 5090 ?**

Un minimum de 1000 W est recommandé par NVIDIA, avec 1200 W conseillé pour une configuration confortable incluant un CPU haut de gamme et de la marge pour l’overclocking.

**Faut-il toujours prendre une alimentation plus puissante que le calcul théorique ?**

Oui, dans une limite raisonnable. La marge de 30-40 % couvre les pics transitoires réels, mais inutile de viser 1600 W pour une configuration RTX 5070 : cela augmente le coût sans bénéfice, une alimentation trop peu chargée n’est pas plus efficace.

**Le connecteur 12V-2×6 est-il compatible avec les anciens câbles 12VHPWR ?**

Mécaniquement oui via un adaptateur, mais ce n’est pas recommandé pour les GPU les plus gourmands (RTX 5090). Privilégiez un câble natif 12V-2×6 fourni avec une alimentation ATX 3.1.

**80 PLUS Gold suffit-il pour un PC gamer en 2026 ?**

Pour la majorité des configurations jusqu’à 850 W, oui. Au-delà de 1000 W, Platinum ou Titanium limitent mieux la chaleur dissipée et la facture d’électricité sur le long terme.

**Single rail ou multi rail pour un seul GPU ?**

