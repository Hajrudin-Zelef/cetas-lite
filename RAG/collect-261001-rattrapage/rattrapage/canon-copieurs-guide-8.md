---
id: collect-261001-rattrapage/rattrapage/canon-copieurs-guide-8
title: "Canon — Guide ultra-complet copieurs (maintenance au cœur)"
domain: rattrapage
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/canon_copieurs_guide.md
source_anchor: ""
source_lines: [1185, 1386]
sha256: a063cbb4c1f50ec7843729ae0032c5f19445a2544091baaa225a3858a0cd9be3
---

# Canon — Guide ultra-complet copieurs (maintenance au cœur)

### 37.1 Méthode générale
1. Photo **avant** (câbles, nappes, vis).
2. Vis : posez-les sur un plan **dans l'ordre** (scotch + schéma papier).
   Les vis Canon ont des longueurs différentes : une vis trop longue
   dans un trou de nappe = carte percée.
3. Nappes : soulevez le **verrou** du connecteur avant de tirer
   (jamais tirer sur la nappe elle-même).
4. Connecteurs : ne forcez jamais — s'il résiste, il y a un verrou.
5. Remontage : rebranchez **tout** avant de revisser les capots
   (test intermédiaire).

### 37.2 Accès carte principale / DC controller
1. Capot arrière (4–6 vis), repérer la cage des cartes.
2. Débrancher les nappes (photo !), dévisser la carte.
3. **ESD** : bracelet antistatique ou touchez une masse avant de
   manipuler (une décharge tue une carte à 500 €).
4. Au remontage : vérifier chaque nappe est bien enfoncée **jusqu'au
   clic** (la moitié des E240/E711 viennent d'une nappe mal remise).

---

## 38. Refroidissement et gestion thermique

- Un copieur A3 dégage **1–2 kW** de chaleur en production : le local
  doit être ventilé/climatisé (sinon E000 en cascade l'été).
- Circuit : ventilateurs d'admission → cartes → ventilateur
  d'évacuation + ventilateur de la fusion.
- Maintenance : dépoussiérage des grilles **à chaque visite**
  (poussière = isolation thermique), vérifier le sens de rotation
  des ventilateurs après remplacement.
- Symptômes de surchauffe : E000/E001 sans cause fusion, ralentissement,
  erreurs aléatoires (E315, E240 fantômes).

---

## 39. Les capteurs — types et tests

| Type | Rôle | Test |
|---|---|---|
| Photo-interrupteur | Détection papier | DISPLAY > SENSOR, passer une feuille |
| Thermistance | Température fusion | Résistance vs température (manuel) |
| Capteur de densité | Patch toner | Nettoyer, DENS |
| Capteur de niveau toner | Fin de toner | Secouer doucement la cartouche |
| Capteur de porte | Sécurité | Ouvrir/fermer, état SENSOR |
| Capteur d'humidité/température | Environnement | DISPLAY > ANALOG |

- **90 % des « capteurs HS » sont des capteurs sales** : nettoyez
  avant de commander.
- Les capteurs optiques vieillissent (LED qui faiblit) : si le
  nettoyage ne suffit plus après 5 ans, remplacez.

---

## 40. Bonus — scanners imageFORMULA

- Pas de toner ni de tambour : la maintenance = **rouleaux + vitres**.
- Kit de rouleaux (pickup/feed/separation) à changer selon le compteur
  (souvent 200k–600k passages).
- Nettoyage : vitres CIS + rouleaux à l'alcool isopropylique,
  **kit de nettoyage Canon** (lingettes).
- Bourrages : guides mal réglés, documents agrafés/froissés,
  rouleaux lisses.
- Pilote : **CaptureOnTouch** (Canon) — à maintenir à jour.

---

## 41. Modèles A4 récents (imageRUNNER 1600/2600 series)

- N&B A4, remplace les 1435 : simples, pour petits bureaux.
- Maintenance allégée : cartouche **tout-en-un** (toner + tambour),
  peu de pièces détachées.
- Points faibles : chargeur ADF (rouleaux), bac papier (guides).
- En contrat : coût-page simple à calculer (une seule cartouche).

---

## 42. Lexique service étendu

| Abréviation | Signification |
|---|---|
| DC-CON | Carte contrôleur DC (moteur) |
| MAIN-CON | Carte contrôleur principale |
| HVT | Transformateur haute tension |
| LSU | Unité laser (Laser Scanner Unit) |
| CIS | Capteur d'image par contact (scanner) |
| CCD | Capteur d'image (anciens scanners) |
| ITB | Courroie de transfert intermédiaire |
| ETB | Courroie de transfert électrostatique (N&B) |
| OPC | Tambour photoconducteur organique |
| DADF | Chargeur recto-verso automatique |
| RADF | Chargeur recto-verso (ancien terme) |
| SST | Service Support Tool |
| MEAP | Plateforme applicative Canon |
| UFR II | Langage d'impression Canon |
| PCL | Langage d'impression universel |
| PS | PostScript (option) |
| SMB | Partage de fichiers Windows |
| LDAP | Annuaire (carnet d'adresses d'entreprise) |
| TPM | Puce de sécurité (chiffrement) |
| E-RDS | Diagnostic à distance Canon |

---

---

## 43. Fiches d'intervention types (templates)

### 43.1 Rapport de visite préventive
```
Date : ___  Client : ___  Machine : ___ (n° série : ___)
Compteur total : ___  Couleur : ___  N&B : ___
Erreurs relevées : ___
Bourrages relevés : ___
Consommables : toner N ___% C ___% M ___% J ___% / bac usagé ___%
Pièces (PARTS) : tambour ___%  fusion ___%  ITB ___%  développeur ___%
Actions : nettoyage optiques / rouleaux / calibration / firmware ___
Préconisations : ___
Prochaine visite : ___
Signature technicien : ___  Signature client : ___
```

### 43.2 Rapport de dépannage
```
Date : ___  Appel client : ___
Symptôme décrit : ___
Code erreur / jam : ___
Diagnostic : ___
Pièces remplacées (références) : ___
Réglages modifiés (anciennes valeurs notées) : ___
Tests effectués : ___
Résultat : OK / à suivre : ___
```

---

## 44. Garanties et litiges

- Garantie constructeur : 1 an pièces (hors consommables), main-d'œuvre
  selon contrat. **Le toner compatible annule la garantie** sur le
  circuit image.
- Garantie de votre intervention : 3 mois minimum sur la pièce
  remplacée (standard pro).
- Litige « ça marchait avant votre passage » : votre **rapport écrit**
  avec compteurs et photos = votre protection. Toujours faire signer.
- Pièce DOA (morte à l'arrivée) : ne jamais la monter de force —
  photo, retour fournisseur immédiat.

---

## 45. Ce qui change (2024-2026)

- **Sécurité** : TPM + McAfee de série (DX), chiffrement disque par
  défaut → les procédures HDD changent (sauvegarde de clé obligatoire).
- **SMB** : SMBv1 mort partout → firmwares récents indispensables.
- **Cloud** : impression/scannage cloud natif (OneDrive, Google Drive,
  Dropbox) → moins de serveurs locaux, plus de config réseau/web.
- **Consommables** : cartouches tout-en-un sur de plus en plus de
  modèles (maintenance simplifiée, mais coût-page à recalculer).
- **E-Maintenance** : la remontée auto des compteurs devient le standard
  des contrats → investissez dans la plateforme, c'est votre vigie.

---

## 46. Ressources

- **Manuels de service** : via le distributeur Canon agréé (accès
  partenaire). Indispensables pour les vues éclatées et les valeurs
  de réglage exactes.
- **SST + firmwares** : portail partenaire Canon.
- **Communautés** : forums de techniciens (Copytechnet, groupes
  Facebook/WhatsApp de techniciens Afrique) — entraide sur les pannes
  tordues.
- **Formation** : académies Canon régionales, formations fabricants
  (Deye/Victron pour le solaire, Canon pour le print — même logique).

---

## 47. Index rapide des codes

- **E000–E009** : fusion → §5.1
- **E010–E019** : moteurs, toner usagé → §5.1/§5.2
- **E020–E028** : densité toner → §5.2
- **E110–E301** : scanner/lecteur → §5.3
- **E240–E749** : contrôleur/communication/cartes → §5.4
- **E400–E490** : ADF → §5.5
- **E500–E5F5** : finisher → §5.6
- **E602–E630** : disque/mémoire/firmware → §5.4
- **E674–E699** : fax/cartes → §5.4
- **E700–E799** : modules/options → §17
- **E800–E84x** : ventilateurs/moteurs → §5.7/§17
- **Jam 00xx/01xx/02xx/03xx/0Axx/0Dxx** : bourrages → §6/§18

---

---

## 48. Annexe — durées de vie par famille (repères terrain)

> Ordres de grandeur constatés en Afrique de l'Ouest (chaleur,
> poussière, papier variable). Divisez par ~1,5 les valeurs
> constructeur « conditions idéales ».

