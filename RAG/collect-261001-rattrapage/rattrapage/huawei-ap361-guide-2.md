---
id: collect-261001-rattrapage/rattrapage/huawei-ap361-guide-2
title: "Huawei eKit AP361 — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_ap361_guide.md
source_anchor: ""
source_lines: [146, 298]
sha256: 0f3245e13ce14da220c46048959cbe138f56e29edc0c3789bbe2e4df0e1de12a
---

# Huawei eKit AP361 — Guide ultra-complet

- **1 x RJ45 10/100/1000BASE-T**, en dessous / à l'arrière du boîtier selon le sens
  de montage.
- C'est à la fois le **port de données** (uplink vers le switch) et le **port
  d'alimentation PoE-in**.
- Vitesse négociée affichée dans l'app eKit : visez **1000 Mbit/s full duplex**.
  Si vous voyez 100 Mbit/s → câble ou port du switch en cause (voir dépannage §124).

### 5.2. LED d'état

L'AP361 dispose d'une LED d'état (souvent sur la tranche ou près du port).
Signification **typique** des AP Huawei eKit (🔎 codes exacts à vérifier dans le
guide hardware du modèle) :

| État LED (typique) | Signification probable |
|---|---|
| Éteinte | Pas d'alimentation PoE → vérifier injecteur/switch/câble |
| Verte fixe | Fonctionnement normal, associé au cloud / en service |
| Verte clignotante lente | Démarrage en cours, ou en cours d'association |
| Rouge / orange | Défaut : échec boot, pas d'IP, pas de cloud, firmware corrompu |
| Clignotement rapide | Mise à jour firmware en cours → **ne pas couper l'alimentation** |

> ⚠️ Ne vous fiez pas à ma table par cœur : photographiez la page « LED » du guide
> hardware officiel et gardez-la dans le dossier du site. En dépannage nocturne,
> c'est elle qui fait foi.

### 5.3. Bouton Reset

- Trou « Reset » accessible avec un trombone (🔎 à vérifier physiquement).
- Procédure standard Huawei : AP sous tension → appuyer **> 5 secondes** →
  relâcher → l'AP redémarre en configuration d'usine.
- **Effet** : efface SSID, comptes, adresse cloud. Ne touche pas (en général) au
  firmware actif. À réserver aux reprises en main (voir §26).

## 6. Dimensions, poids, discrétion

- **Ø 180 mm, 35 mm d'épaisseur, ~450 g** ✅ (fiches distributeurs).
- Format « galet » blanc, discret au plafond. En zone patrimoniale ou architecte
  exigeant, prévoir une peinture adaptée (🔎 vérifier que la peinture n'est pas
  métallisée — elle atténue le signal — et qu'elle ne fait pas sauter la garantie).
- 450 g : une fixation sur dalle de faux plafond standard tient sans renfort
  particulier, mais **toujours** visser la platine dans l'ossature ou utiliser des
  chevilles adaptées au support (voir §11-14).

## 7. Environnement d'exploitation

| Paramètre | Valeur de référence | Commentaire terrain |
|---|---|---|
| Température | -10 à +50 °C (fiche) / 0 à +40 °C (revendeur) | En intérieur climatisé : aucun souci. Éviter les combles non ventilés, les cuisines pro, les locaux techniques surchauffés |
| Humidité | 5-95 % HR sans condensation | Pas de montage en zone humide (piscine, buanderie) sans caisson adapté |
| Poussière | Usage intérieur propre | En atelier poussiéreux, prévoir nettoyage trimestriel des ouïes |
| CEM | Éviter à moins de 30 cm des câbles de puissance | Les variateurs, ballasts et moteurs sont les pires voisins |
| Altitude | jusqu'à 5 000 m | Sans objet en métropole/Afrique de l'Ouest |

> 🔥 **Point chaud** : un AP collé au plafond d'un local non climatisé sous
> toiture tôle peut dépasser 50 °C l'après-midi. Mesurez avant de valider
> l'emplacement définitif (thermomètre IR à 20 €, ça paie).

## 8. Antennes intelligentes : ce que ça change (et ce que ça ne change pas)

L'AP361 embarque des antennes internes dites « smart » : l'AP commute
électroniquement entre plusieurs diagrammes de rayonnement pour suivre les
clients. Concrètement :

- ✅ Ça **améliore la stabilité** quand les utilisateurs bougent (couloirs, open space).
- ✅ Ça **ne remplace pas** une étude de couverture : 2x2 MIMO à 5 dBi reste un
  petit AP. Comptez un AP pour 80-150 m² en open space, beaucoup moins derrière
  des murs béton.
- ✅ Le gain de 5 dBi est **omnidirectionnel** : le signal part dans toutes les
  directions, y compris vers l'étage du dessus/dessous. En immeuble, baissez la
  puissance plutôt que de « arroser » les voisins (voir §49).

## 9. Gamme eKit : où se situe l'AP361

La gamme **Huawei eKit** vise les PME avec une promesse : *zéro contrôleur à
acheter, gestion cloud depuis un smartphone*. L'AP361 est le modèle d'intérieur
« standard ». Au-dessus, on trouve des modèles 4x4 ou extérieurs ; en dessous,
des AP muraux (wall plate) pour l'hôtellerie.

**Conséquence pratique** : tout l'écosystème (app mobile, portail cloud,
licences) est pensé « simple ». Si vous venez du monde AirEngine + AC (contrôleur
WLAN), vous allez trouver ça dépouillé — c'est normal, c'est le positionnement.
Pour un parc PME de 5 à 50 AP, c'est un avantage. Au-delà, évaluez la gamme
entreprise classique.

## 10. Documents officiels à garder sous la main

Checklist documentaire du chef de service (à constituer **avant** le chantier) :

- [ ] Datasheet officielle AP361 (PDF, ekit.huawei.com)
- [ ] *WLAN Hardware Installation and Maintenance Guide* — chapitre AP361
      (montage, LED, reset, specs environnementales détaillées)
- [ ] *Configuration Guide* eKit correspondant à votre version logicielle
- [ ] Release notes du firmware que vous déployez
- [ ] Plan de masse du site avec emplacements AP numérotés (AP-01, AP-02…)
- [ ] Tableau d'adressage IP / VLAN du site

> 💡 Stockez tout ça dans le dossier du site sur votre partage d'équipe. Dans
> 2 ans, quand un AP tombera un dimanche, vous bénirez cette discipline.

---
---

# B. MONTAGE MÉCANIQUE

## 11. Le kit de fixation : inventaire avant de monter à l'échelle

Avant de grimper, ouvrez **un** carton et vérifiez le kit (🔎 contenu exact selon lot) :

- Platine de fixation (avec ergots baïonnette ou clips)
- Vis + chevilles (souvent prévues pour béton/placo — **à adapter au support réel**)
- Gabarit de perçage (parfois imprimé sur le carton)
- Éventuellement : collier ou adaptateur pour rail de faux plafond

**Règle d'or** : ne commencez jamais un chantier de 20 AP sans avoir validé le
montage sur **1 AP témoin**. Tous les faux plafonds ne se ressemblent pas.

## 12. Montage au plafond — dalle de faux plafond (cas le plus courant)

Procédure type (adaptez au kit réel) :

1. **Repérez l'ossature** : la platine doit être fixée sur le profilé métallique
   (T-bar), pas seulement sur la dalle minérale qui s'effrite.
2. **Faites passer le câble** : RJ45 caté6a ou cat6, avec **30 cm de mou** en
   réserve au-dessus de la dalle (pour décrocher l'AP sans arracher le câble).
3. **Vissez la platine** sur l'ossature (2 vis minimum).
4. **Clipsez l'AP** sur la platine (quart de tour ou clips selon modèle) jusqu'au
   « clic ».
5. **Branchez le RJ45** (clic du connecteur), vérifiez que la LED s'allume.
6. **Étiquetez** : AP-01, AP-02… au marqueur sur la platine + dans l'app eKit.

> ⚠️ **Sécurité** : un AP de 450 g qui tombe de 3 m sur un collaborateur, c'est un
> accident du travail. Serrage vérifié + test de traction à la main (tirez
> franchement : ça ne doit pas bouger).

## 13. Montage au plafond dur (béton, hourdis)

1. Percez au gabarit (Ø selon chevilles, typiquement 6 mm).
2. Chevilles adaptées : **béton** → chevilles nylon/frappe ; **hourdis creux** →
   chevilles à expansion adaptées (pas de simple cheville nylon dans du creux !).
3. Vissez la platine, clipsez l'AP, câblez avec un chemin de câble ou une goulotte
   jusqu'au point de raccordement.
4. Si le câble court en apparent : goulotte blanche 20x12, propre et pro.

## 14. Montage mural

- L'AP361 se monte aussi au mur (✅ datasheet). Orientez le dôme **vers la zone à
  couvrir** : le rayonnement est optimisé pour un montage plafond (vers le bas) ;
  au mur, la couverture devient asymétrique.
- Hauteur conseillée : **2,5 à 4 m**. En dessous, vandalisme et masquage par le
  mobilier ; au-dessus, le signal s'étale et la puissance utile baisse.
- Évitez le montage mural **derrière** une armoire métallique, un tableau
  électrique ou un pilier béton : c'est le meilleur moyen de créer une zone d'ombre.

