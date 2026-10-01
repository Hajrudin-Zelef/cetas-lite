---
id: collect-261001-rattrapage/rattrapage/canon-copieurs-guide-5
title: "Canon — Guide ultra-complet copieurs (maintenance au cœur)"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-rattrapage/canon_copieurs_guide.md
source_anchor: ""
source_lines: [655, 813]
sha256: d30c06d7deb9b3e735ff161df14cd0300952fc87d6cbfcddc078f89b3cc8e304
---

# Canon — Guide ultra-complet copieurs (maintenance au cœur)

| Code | Composant / sens | Première action |
|---|---|---|
| E012-0000 | Moteur tambour | PART-CHK moteur, engrenages, connecteurs |
| E014-0000 | Moteur fixing | Vérifier charge mécanique de la fusion |
| E019-0001 | Pression toner usagé | Bac plein / capteur (voir E013) |
| E026–E028 | Densité toner couleur (M/C/J) | Voir E020, par station |
| E032-0001 | Compteur NE (option) | Réinitialiser le compteur du contrôleur NE |
| E121-0000 | Erreur rotation (générique moteur) | Identifier le moteur via le détail YYYY |
| E196-0001 | Erreur ROM / firmware | Reflash SST du module concerné |
| E197-0000 | Erreur moteur (laser/scanner) | Selon YYYY : laser ou scanner |
| E260-0000 | Capteur de densité | Nettoyer le capteur, ajustement DENS |
| E261-0000 | Signal zero-cross | Alimentation AC, carte d'alimentation |
| E315-xxxx | Erreur données image | Redémarrer, RAM/HDD, firmware |
| E317-xxxx | Erreur communication (module) | Selon YYYY, connecteurs puis carte |
| E350-0000 | Erreur communication (générique) | Câbles internes, cartes |
| E354-xxxx | Erreur contrôleur (sous-code) | Manuel de service du modèle |
| E390-xxxx | Erreur DDI (interface) | Redémarrer, carte interface |
| E407-0000 | Moteur rouleau ADF | PART-CHK, courroie |
| E408-0000 | Moteur séparation ADF | Rouleaux, moteur |
| E409-0000 | Moteur lecture ADF | Capteur + moteur |
| E413-0000 | Moteur décalage ADF | Butée, capteur |
| E420-0000 | EEPROM ADF | Réécrire les données, remplacer carte ADF |
| E490-xxxx | Mauvais type DADF | Vérifier la compatibilité du chargeur |
| E503-xxxx | Communication finisher (détail) | Câble, carte finisher |
| E505-xxxx | Mémoire finisher | Carte finisher |
| E514-xxxx | Moteur sortie finisher (détail) | Butée, capteur, moteur |
| E518-xxxx | Moteur décalage (finisher) | Idem |
| E520-xxxx | Moteur agrafage (détail) | Agrafe coincée, capteur |
| E531-xxxx | Alignement (détail) | Poussière papier, capteur |
| E540-xxxx | Déplacement agrafeuse | Rail, capteur, moteur |
| E544-xxxx | Réserve agrafes | Recharger, détecteur |
| E551-xxxx | Ventilateur finisher | Nettoyer/remplacer |
| E575-xxxx | Empileur (stacker) | Butée, capteur |
| E584-xxxx | Volet (shutter) finisher | Mécanisme, capteur |
| E590-xxxx | Perforatrice | Bac confettis plein ! capteurs |
| E5F0–E5F5 | Module piqûre à cheval (saddle) | Agrafes, butées, capteurs |
| E601-xxxx | Erreur communication (disque/carte) | Connecteurs, HDD |
| E603-xxxx | Erreur clavier/panneau | Nappe panneau |
| E605-xxxx | Erreur mémoire (détail) | Barrette RAM, réinstaller |
| E606-xxxx | Erreur HDD (détail) | Voir E602 |
| E609-xxxx | Erreur carte (détail) | Selon YYYY |
| E611-xxxx | Erreur communication (option) | Carte option |
| E614-xxxx | Firmware incompatible | Reflash version cohérente |
| E616-xxxx | Erreur données (détail) | Restaurer réglages |
| E674-xxxx | Carte fax | Réinstaller, vérifier ligne |
| E677-xxxx | Carte imprimante (option) | Firmware, carte |
| E676-xxxx | Erreur mémoire fax | Carte fax |
| E700-xxxx | Erreur finisher (générique) | Câble, carte |
| E701-xxxx | Erreur communication (module) | Selon YYYY |
| E712-xxxx | Erreur option (détail) | Carte option concernée |
| E719-xxxx | Monnayeur/lecteur cartes | Voir §5.7 |
| E720-xxxx | Erreur carte (détail) | Selon YYYY |
| E730-xxxx | Erreur PDL | Réinstaller firmware PDL |
| E734-xxxx | Erreur communication (détail) | Connecteurs |
| E737-xxxx | Erreur mémoire (détail) | RAM |
| E740-xxxx | Erreur carte réseau | Carte NIC, firmware |
| E741-xxxx | Erreur PCI | Reseat carte, slot |
| E743-xxxx | Erreur DDI-S | Redémarrer |
| E745-xxxx | Erreur carte (détail) | Selon YYYY |
| E746-xxxx | Erreur contrôleur d'impression | Carte, firmware |
| E750-xxxx | Erreur communication (détail) | Connecteurs |
| E804-xxxx | Ventilateur alimentation | Nettoyer/remplacer |
| E806-xxxx | Ventilateur finisher | Idem |
| E808-xxxx | Entraînement fixing | Voir §5.7 |
| E810-xxxx | Moteur nettoyeur | Vérifier mécanisme |
| E812-xxxx | Ventilateur (détail) | Selon YYYY |
| E815-xxxx | Moteur (détail) | PART-CHK |
| E817-xxxx | Rotatif développement (couleur) | Engrenages du carrousel |
| E819-xxxx | Erreur moteur (détail) | Selon YYYY |
| E821-xxxx | Nettoyeur (détail) | Raclette, bac |
| E823-xxxx | Ventilateur (détail) | Selon YYYY |
| E825-xxxx | Ventilateur (détail) | Selon YYYY |
| E830-xxxx | Erreur refroidissement | Ventilateurs, filtres |
| E840-xxxx | Volet (détail) | Mécanisme |

> **Règle** : un code en `E8xx` = souvent ventilateur/moteur
> (poussière), un code en `E02x` = toner/développeur, un code en
> `E40x` = ADF, un code en `E5xx` = finisher, un code en `E6xx` =
> disque/mémoire/firmware. Cette logique par familles accélère le
> diagnostic.

---

## 18. Jam codes — compléments par famille

### 18.1 Lire un jam code complet
L'historique (DISPLAY > JAM) affiche : `date, heure, code, cassette,
format`. Exemple : `0108, C1, A4` = échec de prise, cassette 1, A4.

### 18.2 Les récurrents en maintenance Afrique

| Jam | Zone | Cause n°1 terrain |
|---|---|---|
| 0108 | Pickup C1 | Rouleaux lisses (poussière) |
| 0101 | Pickup | Papier humide (saison des pluies) |
| 0208 | Registration | Guides bac mal réglés |
| 0301 | Sortie fusion | Doigts de séparation usés |
| 0A01 | Duplex | Morceau resté d'un bourrage précédent |
| 0011 | ADF | Document agrafé / rouleaux ADF |

### 18.3 Capteurs : les tester
- DISPLAY > SENSOR : état en temps réel de chaque capteur (0/1).
- Passez une feuille devant le capteur suspect : l'état doit changer.
  Sinon : capteur sale (souffler), débranché, ou HS.
- Les capteurs optiques (photosenseurs) s'encrassent avec la poussière
  de papier : nettoyage à l'air sec à chaque visite des zones à
  bourrages récurrents.

---

## 19. Procédures de remplacement — compléments

### 19.1 Remplacement HDD + réinstallation firmware
1. Sauvegarder : carnet d'adresses, réglages (export).
2. Machine éteinte, remplacer le HDD/SSD (modèle compatible,
   même capacité ou supérieure).
3. Boot en mode download, flasher avec **SST** (firmware complet).
4. Restaurer les réglages, reconfigurer le réseau.
5. Test : 50 copies + scans + impressions réseau.
6. **Prévention** : onduleur + arrêt propre systématique.

### 19.2 Kit de maintenance complet (200–300k pages)
1. Tambour(s), développeur(s), unité de fusion, ITB (couleur),
   kit rollers, bac toner usagé, filtres.
2. Ordre : sortir les anciennes unités, aspirer l'intérieur
   (aspirateur à toner), monter les neuves.
3. **Reset de chaque compteur PARTS** (sinon la machine croit que
   c'est toujours l'ancienne pièce).
4. Initialisations : STIR/TONER-S (développeur), calibration couleurs,
   ajustement densité, registration.
5. Test : 100 copies (N&B + couleur), mire qualité, 50 recto-verso.
6. Noter : date, compteur, pièces (n° de série si garantie).

### 19.3 Courroie et rouleaux ADF
1. Ouvrir le capot ADF, repérer pickup/feed/separation rollers
   (souvent clipsés).
2. Remplacer le kit, nettoyer les capteurs du chemin papier.
3. Test : 20 originaux recto-verso en ADF.

### 19.4 Après chaque remplacement — la règle des 3 R
- **Reset** (compteurs PARTS), **Réglage** (calibrations),
  **Rapport** (noter l'intervention).

---

## 20. Calibrations et réglages fins (ADJUST)

> ADJUST = réglages d'usine. **Photo/noter avant chaque modification.**
> Ne touchez que ce que vous comprenez.

| Réglage | Effet | Quand l'utiliser |
|---|---|---|
| DENS (par couleur) | Densité d'image | Après développeur/tambour |
| REGIST | Superposition des couleurs | Après ITB, couleurs décalées |
| BLANK | Marges blanches | Bords sales |
| FEED-ADJ | Vitesse d'entraînement | Bourrages/décalages récurrents |
| FIX-TEMP | Température de fusion | Toner qui s'efface / papier ondulé |
| LASER | Puissance laser | Image pâle persistante |

