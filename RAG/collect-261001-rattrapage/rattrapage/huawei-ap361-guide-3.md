---
id: collect-261001-rattrapage/rattrapage/huawei-ap361-guide-3
title: "Huawei eKit AP361 — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-rattrapage/huawei_ap361_guide.md
source_anchor: ""
source_lines: [299, 449]
sha256: 15508798ca408b0d50c7b12bd5338aa19febd1d06b448f65939b4bd9ec3da7d0
---

# Huawei eKit AP361 — Guide ultra-complet

## 15. Positionnement radio : les règles qui paient

| Règle | Pourquoi |
|---|---|
| Un AP par zone de vie, pas par « surface théorique » | Les murs béton tuent le 5 GHz ; comptez les cloisons, pas les m² |
| Évitez les extrémités de couloir | Mieux vaut centrer sur la zone d'usage |
| Distance entre AP : 10-20 m en open space | Pour un roaming propre et un réemploi de canaux |
| Jamais dans un placard / local technique fermé | Cage de Faraday + surchauffe |
| Jamais à côté (< 1 m) d'un four micro-ondes | Le 2,4 GHz se fait hacher (voir §121) |
| Hauteur homogène sur un même plateau | Sinon le roaming devient imprévisible |
| Évitez l'aplomb des détecteurs incendie/sprinklers | Contrainte réglementaire + maintenance |

**Test terrain** : après montage, faites le tour avec votre téléphone (analyseur
Wi-Fi type « WiFi Analyzer ») : visez **-60 dBm ou mieux** en 5 GHz dans les zones
de travail, **-67 dBm minimum** pour la voix. En dessous de -70 dBm : ajoutez ou
déplacez un AP.

## 16. Précautions CEM et thermiques

- **Distance aux câbles de puissance** : 30 cm minimum en parcours parallèle
  (norme usuelle de séparation courants forts/faibles). Croisement à 90° si
  inévitable.
- **Ne pas poser l'AP sur** : un ballast, un variateur de lumière, un moteur, un
  onduleur. Les harmoniques et le champ magnétique perturbent l'électronique.
- **Ventilation** : ne pas enfermer l'AP dans un caisson étanche sans
  déclassement thermique. Si caisson obligatoire (extérieur abrité, atelier),
  mesurez la température interne en charge.
- **Foudre** : en intérieur c'est le switch et le câblage qui prennent ; prévoyez
  parafoudre sur l'arrivée électrique de la baie et, en zone orageuse, des
  protections PoE/Ethernet sur les liens exposés.

## 17. Checklist de fin de montage (à cocher par AP)

- [ ] Platine fixée sur support solide (test de traction OK)
- [ ] AP clipse, ne bouge pas
- [ ] RJ45 branché (clic), 30 cm de mou au-dessus de la dalle
- [ ] LED allumée (couleur notée : _______)
- [ ] Étiquette AP-XX posée sur la platine
- [ ] Photo de l'emplacement (pour le DOE — dossier des ouvrages exécutés)
- [ ] Emplacement reporté sur le plan de masse
- [ ] Câble testé au testeur (continuité + plan de câblage)

---
---

# C. ALIMENTATION POE

## 18. Ce que dit le constructeur : 802.3af, 8,8 W

✅ Datasheet : alimentation **PoE 802.3af**, consommation **8,8 W** (certains
revendeurs annoncent 9,4 W en pointe — retenez **10 W** pour vos calculs, marge
incluse).

Rappel des standards PoE :

| Standard | Nom courant | Puissance max par port (PSE) | Puissance utile (PD) | Tension |
|---|---|---|---|---|
| 802.3af | PoE | 15,4 W | 12,95 W | 44-57 V |
| 802.3at | PoE+ | 30 W | 25,5 W | 50-57 V |
| 802.3bt type 3 | PoE++ / 4PPoE | 60 W | 51 W | 50-57 V |
| 802.3bt type 4 | PoE++ | 90/100 W | 71,3 W | 52-57 V |

L'AP361 étant **802.3af**, un switch **PoE (15,4 W/port) suffit** : pas besoin de
PoE+ pour l'alimenter. MAIS (voir §21) prévoyez large si le switch doit aussi
alimenter des caméras ou des AP plus gourmands demain.

## 19. Calcul du budget PoE — la méthode

**Formule** : `Budget à prévoir = Σ (consommation max de chaque PD) × 1,2`
(20 % de marge pour les pertes câble et les pics).

**Exemple : 12 x AP361 sur un switch 24 ports PoE**

| Poste | Calcul | Résultat |
|---|---|---|
| 1 AP361 | 8,8 W → arrondi | 10 W |
| 12 AP361 | 12 × 10 W | 120 W |
| Marge 20 % | 120 × 1,2 | **144 W** |
| Budget switch à choisir | ≥ 144 W | **≥ 150 W** (palier commercial courant : 185 W, 250 W, 370 W) |

> 💡 **Piège classique** : le « budget PoE total » du switch n'est pas le nombre
> de ports × 15,4 W. Un switch « 24 ports PoE » avec 185 W de budget ne peut pas
> alimenter 24 équipements à 15,4 W (ça ferait 370 W). Lisez **toujours** la ligne
> « PoE power budget » de la fiche du switch.

**Tableau de dimensionnement rapide (AP361 seuls, marge 20 % incluse)** :

| Nb d'AP361 | Budget à prévoir | Palier switch conseillé |
|---|---|---|
| 4 | 48 W | 60-130 W |
| 8 | 96 W | 130-185 W |
| 12 | 144 W | 185 W |
| 16 | 192 W | 250 W |
| 24 | 288 W | 370 W |

## 20. Distances : jusqu'où peut aller le câble ?

- **Limite Ethernet : 100 m** (90 m de câble fixe + 10 m de jarretières). Au-delà :
  pas de PoE fiable, pas de gigabit garanti.
- **Chute de tension** : en 802.3af à 10 W sur 100 m de Cat5e, la perte est de
  l'ordre de 1 à 2 W — acceptable. En câble bas de gamme (CCA — aluminium cuivré),
  la résistance est plus forte : **exigez du 100 % cuivre**.
- **Règle terrain** : si la distance dépasse 80 m, tirez du **Cat6A** (meilleure
  section, moins de pertes) et mesurez la tension PoE à l'arrivée si vous avez un
  testeur PoE (recommandé : ~200 €, amorti au premier dépannage).

| Distance | Câble conseillé | Remarque |
|---|---|---|
| < 50 m | Cat6 100 % cuivre | Cas standard |
| 50-80 m | Cat6 100 % cuivre | OK sans souci |
| 80-100 m | Cat6A 100 % cuivre | Marge de sécurité |
| > 100 m | **Interdit en cuivre** | Ajouter un switch intermédiaire ou passer en fibre + convertisseur PoE local |

## 21. Switch PoE vs injecteur : comment choisir

| Critère | Switch PoE manageable | Injecteur PoE individuel |
|---|---|---|
| Nb d'AP | ≥ 4 : le switch gagne | 1-3 AP isolés |
| Supervision | Oui (SNMP, état par port, reset PoE à distance) | Non (sauf modèles manageables, rares) |
| Redémarrage à distance | Oui (cycle PoE du port) | Non (il faut aller sur place) |
| Coût | Mutualisé, dégressif | ~20-40 €/pièce, mais N alimentations = N points de panne |
| Secours électrique | Centralisé sur l'onduleur de la baie | Dispersé, souvent oublié |
| Recommandation | ✅ **Choix par défaut** | Dépannage, site distant, ajout ponctuel |

**Choix du switch** (exemples de gammes, pas de pub : vérifiez les fiches) :
un switch manageable **24 ports GE avec budget PoE ≥ 185 W** couvre 12 AP361 +
caméras. Vérifiez : budget PoE, ventilateurs (bruit en open space !), garantie,
et compatibilité 802.3af (tous les switchs PoE+ / bt sont rétrocompatibles af).

**Choix de l'injecteur** : **802.3af minimum, 15,4 W**, 44-57 V. Un injecteur
802.3at (30 W) convient aussi (rétrocompatible). Évitez les injecteurs « passifs »
24 V non normalisés (type vieux Ubiquiti passif) : **incompatibles et dangereux**
pour l'AP361.

## 22. Câblage : les 5 règles non négociables

1. **100 % cuivre, pas de CCA** (Copper Clad Aluminum). Le CCA a ~60 % de
   résistance en plus → échauffement + chute de tension PoE. Exigez la mention
   sur le touret.
2. **Cat6 minimum** pour du neuf (Cat6A si > 80 m ou environnement perturbé).
3. **Pas de rallonge RJ45 bricolée** : un seul tenant de la baie à l'AP.
   Les raccords intermédiaires sont la cause n°1 des « PoE qui ne monte pas ».
4. **Rayon de courbure** : ≥ 4x le diamètre du câble. Un câble plié à angle droit
   derrière la dalle = paire fragilisée = 100 Mbit/s au lieu du gigabit.
5. **Testez chaque lien** au testeur (wiremap + longueur). Un lien non testé est un
   lien en panne qui s'ignore.

## 23. Sécurisation électrique de la baie PoE

Vous êtes chef de service Systèmes & Énergies : c'est votre terrain.

