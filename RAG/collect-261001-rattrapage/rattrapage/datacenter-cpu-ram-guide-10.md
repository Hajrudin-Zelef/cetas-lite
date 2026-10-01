---
id: collect-261001-rattrapage/rattrapage/datacenter-cpu-ram-guide-10
title: "CPU SERVEUR, RAM SERVEUR, CHÂSSIS & VENTILATION"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["datacenter", "gpu"]
source: docs/RAG/collect-261001-rattrapage/datacenter_cpu_ram_guide.md
source_anchor: ""
source_lines: [1448, 1628]
sha256: a77ab0dd7a461e117a06aa6c07b10457d74991605dda006241256d1622727daa
---

# CPU SERVEUR, RAM SERVEUR, CHÂSSIS & VENTILATION

Contraintes : poids (40-70 kg équipé → **2 personnes minimum** ou lève-
serveur), profondeur souvent > 800 mm (baies 1200 mm), consommation de
3 à 10 kW par serveur (section 85). Un 4U GPU se dimensionne comme une
**charge électrique industrielle**, pas comme un serveur.

## 73. Profondeur des châssis et des baies

Profondeurs typiques (indicatif, vérifier par modèle) :

| Châssis | Profondeur |
|---|---|
| 1U/2U standard | 600-750 mm |
| 2U stockage dense | 750-850 mm |
| 4U GPU | 800-900 mm |

Baies : 600, 800, 1000, 1100, 1200 mm de profondeur utile. Règles :
- **Profondeur baie ≥ profondeur serveur + 150 mm** (câbles arrière +
  circulation d'air + porte).
- Un serveur de 800 mm dans une baie de 600 mm = porte arrière impossible
  à fermer = recirculation d'air chaud (section 63) + câbles écrasés.
- En rénovation : mesurez vos baies **avant** de commander. C'est
  l'erreur la plus bête et la plus fréquente.

## 74. Rails : types et pièges

| Type de rail | Caractéristique | Piège |
|---|---|---|
| Rails à friction | le serveur coulisse, pas de roulement | effort important, 1U/2U légers |
| Rails télescopiques (ball-bearing) | sortie complète, roulements | vérifier la longueur vs baie |
| Rails « tool-less » | montage sans vis | compatibilité trous carrés/ronds |
| Bras de câbles (CMA) | guide les câbles en sortie | gêne la ventilation arrière si mal réglé |

Points de vigilance :
- Les rails sont **spécifiques au châssis** (ou à la famille) : pas de
  rail générique fiable.
- Charge max des rails (indicatif : 30-60 kg) : un 4U de 60 kg sur des
  rails de 40 kg = accident programmé.
- **Serrez les vis avant** : un serveur mal fixé glisse et arrache les
  câbles d'alimentation — panne franche le jour J.

## 75. Poids : chiffres et manutention (indicatif)

| Serveur équipé | Poids typique |
|---|---|
| 1U 2S, 12 DIMM, 8 NVMe | 15-20 kg |
| 2U 2S, 24 DIMM, 12 NVMe | 25-35 kg |
| 2U stockage 24× 3,5" | 35-45 kg |
| 4U GPU 8× GPU | 60-90 kg |

Règles de manutention :
- > 25 kg : 2 personnes. > 50 kg : lève-serveur ou 3 personnes.
- Sortez les alimentations et les disques pour alléger avant montage
  (gagne 5-10 kg).
- **Ne transportez jamais un serveur sur ses rails déployés** : les
  roulements ne sont pas conçus pour les chocs latéraux.

## 76. Tours vs rack : quand choisir quoi

| Critère | Tour (tower) | Rack |
|---|---|---|
| Bruit | 30-45 dB(A) : bureau possible | 60-80 dB(A) : local dédié |
| Refroidissement | gros ventilateurs lents, efficace | dense, bruyant |
| Densité | 1 serveur = 1 tour | 42 serveurs par baie |
| Prix à config égale | souvent -10 à -20 % (indicatif) | + rails, + baie |
| Évolution | spacieux, simple | contrainte par le format |

Verdict : **PME sans local technique → tour silencieuse** (ou hébergement
externe). Dès qu'il y a un local dédié ou plus de 3-4 serveurs, le rack
gagne sur tous les plans (câblage, PDU, brassage, sécurité).

## 77. Blade : principe, avantages, limites (bref)

Les lames (blade) mutualisent alimentations, ventilateurs et réseau dans un
châssis : densité extrême, câblage minimal, administration unifiée.

| Avantage | Inconvénient |
|---|---|
| densité max (16-32 lames/châssis) | verrouillage constructeur total |
| câblage réduit | prix d'entrée élevé (châssis) |
| provisioning rapide | panne châssis = N serveurs impactés |

Positionnement 2026 : le blade recule face aux **serveurs rack denses 1U/2U**
et à l'hyperconvergé, sauf niches (VDI massif, telco). Si vous n'avez pas
déjà un parc blade, **n'y entrez pas** : le verrouillage et le coût du
châssis ne se justifient plus pour la plupart des usages.

## 78. Densité de baie : le kW par rack

C'est **le** chiffre qui dimensionne la salle : additionnez les puissances
max des serveurs, divisez par 42U.

| Remplissage type | Puissance/rack (indicatif) |
|---|---|
| 20× serveurs 1U généralistes (~500 W) | ~10 kW |
| 14× serveurs 2U 2S (~1000 W) | ~14 kW |
| 10× serveurs 2S HPC (~2000 W) | ~20 kW |
| 8× serveurs 4U GPU (~8 kW) | ~64 kW |

Repères salle :
- Salle climatisée classique (détente directe) : **3-5 kW/rack**.
- Salle moderne (eau glacée, confinement) : **10-20 kW/rack**.
- Au-delà de 20 kW/rack : confinement strict + DLC ou portes arrière à eau.

**Avant d'acheter 10 serveurs, vérifiez que votre salle peut évacuer leur
puissance.** C'est le premier calcul de la partie H (section 92).

## 79. Câblage et PDU en baie : principes

- **Séparez les chemins** : alimentation à gauche, réseau à droite (ou
  l'inverse, mais constant). Un câble réseau qui pend devant une
  alimentation = intervention impossible sans coupure.
- **Longueurs justes** : 0,5-1 m en baie ; les câbles de 3 m enroulés
  bloquent le flux d'air arrière.
- **PDU** : de préférence **monitorées** (mesure par prise) : c'est votre
  wattmètre permanent par serveur (section 90).
- **Étiquetage** : chaque câble, chaque prise PDU, chaque serveur. Le jour
  de la panne à 3h du matin, les étiquettes valent de l'or.
- **Bras de câbles (CMA)** : utiles pour sortir un serveur sans
  débrancher, mais ils retiennent la chaleur à l'arrière : à réserver aux
  serveurs peu denses.

## 80. Checklist avant achat d'un châssis

1. [ ] TDP CPU max supporté ≥ TDP de mes CPU (avec marge 20 %).
2. [ ] Redondance ventilateurs N+1 confirmée sur **ma** configuration.
3. [ ] Nombre de slots DIMM = canaux × DPC prévus (12/24 selon carte).
4. [ ] Baies disque : nombre, format (2,5"/3,5"), NVMe/SAS, fond de panier.
5. [ ] Slots PCIe : nombre, génération, largeur, hauteur (pleine/demi).
6. [ ] Profondeur du serveur ≤ profondeur de mes baies − 150 mm.
7. [ ] Rails inclus ? Compatibles avec mes baies (trous carrés/ronds) ?
8. [ ] Alimentations : puissance, redondance, rendement (section 84).
9. [ ] BMC/IPMI : licence incluse ou option payante ?
10. [ ] QVL mémoire et CPU : mes références y figurent ?
11. [ ] Garantie : 3 ans sur site J+1 minimum ; 5 ans pour l'amortissement.
12. [ ] Poids équipé ≤ capacité de mes rails et de mon dos.

---

## 81. PUE expliqué simplement

PUE (Power Usage Effectiveness) = **énergie totale du datacenter /
énergie des équipements IT**. Un PUE de 1,5 signifie que pour 1 W utile
aux serveurs, 0,5 W part en climatisation, éclairage, pertes.

```
 PUE = (IT + froid + pertes elec + auxiliaires) / IT

 Exemple : 100 kW IT, 40 kW froid, 10 kW pertes -> PUE = 150/100 = 1,5
```

Échelle de lecture :
- 1,05-1,2 : excellent (hyperscalers, free cooling, DLC).
- 1,3-1,5 : bon datacenter moderne bien conçu.
- 1,6-2,0 : salle classique, marge de progrès.
- > 2,0 : petite salle mal optimisée — chaque serveur y coûte double.

**Le PUE ne mesure pas l'efficacité des serveurs**, seulement celle de
l'enveloppe. Un serveur efficace dans une salle à PUE 2,0 reste une
mauvaise affaire énergétique globale.

## 82. PUE : valeurs typiques et pièges de mesure

Pièges classiques :
1. **PUE instantané vs annuel** : le PUE varie avec la saison (free cooling
   l'hiver). Seul le PUE **annualisé** a un sens économique.
2. **PUE partiel** : mesurer sans l'éclairage, les bureaux ou les pertes
   transformateur flatte le chiffre. Exigez le périmètre de mesure.
3. **PUE à faible charge** : une salle à 20 % de remplissage a un PUE
   catastrophique (le froid tourne pour rien). Le PUE se juge **à charge
   nominale de conception**.
4. **Le PUE ne dit rien du carbone** : 1,3 au charbon > 1,6 au
   renouvelable, en CO2. Complétez par le CUE (Carbon Usage Effectiveness)
   si le reporting RSE compte.

Pour un chef de service : suivez **deux** chiffres — le PUE annualisé de
la salle et la **consommation IT absolue** (kWh). Réduire le PUE en
laissant des serveurs à vide allumés est une victoire comptable et une
défaite énergétique.

## 83. Du TDP à la facture : la chaîne de calcul

