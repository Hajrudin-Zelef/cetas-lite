---
id: collect-261001-rattrapage/rattrapage/canon-copieurs-guide-2
title: "Canon — Guide ultra-complet copieurs (maintenance au cœur)"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-rattrapage/canon_copieurs_guide.md
source_anchor: ""
source_lines: [178, 383]
sha256: 069d69f9c4f3fdc99b70d24047c7d4161e4adc3b5c98240fabbabbe91c1b0b3d
---

# Canon — Guide ultra-complet copieurs (maintenance au cœur)

**E002-0000 — Température trop basse**
- La fusion ne monte pas : lampe/résistance chauffante coupée,
  connecteur, carte d'alimentation. Tester la continuité de la
  lampe (hors tension !).

**E003-0000 — Basse température après veille**
- Variante de E002 au sortir de veille. Souvent lampe en fin de vie.

**E007-0000 — Erreur rotation film de fusion**
- Le film ne tourne pas : engrenages, moteur, capteur. Démonter,
  vérifier l'entraînement, graisser si préconisé (graisse spéciale
  fixing, jamais de graisse standard qui cuit).

**E009-0000 — Erreur pression de fusion**
- Capteur de pression came : vérifier le mécanisme, le capteur,
  les connecteurs.

### 5.2 Bloc TONER / DÉVELOPPEUR

**E020-0x81 — Erreur densité toner (noir)**
- Le capteur de densité (ou le patch de test) ne lit pas la bonne
  densité. Causes : toner vide, développeur usé, capteur sale,
  trémie bouchée.
- Action : vérifier toner, nettoyer le capteur de densité/patch,
  lancer un ajustement (FUNCTION > DENS), si échec → développeur.

**E021 à E028** — variantes couleur par couleur (C/M/J).
- Même logique, par station couleur.

**E025-0000 — Moteur d'alimentation toner**
- Trémie/bouteille bloquée, moteur HS, toner agglutiné (chaleur).
  Démonter la trémie, vérifier la rotation à la main.

**E013-0000 — Bac toner usagé plein**
- Vider/remplacer le bac, nettoyer le détecteur. Si le code persiste
  après vidage : capteur sale ou connecteur.

### 5.3 Bloc SCANNER / LECTEUR

**E110-0000 — Erreur moteur scanner**
- Le chariot ne bouge pas : courroie, moteur, capteur home.
  Vérifier le déplacement à la main (hors tension), la courroie.

**E202-0001 / 0002 — Erreur position home (CIS)**
- Le capteur de contact ne trouve pas sa position : butée, capteur,
  nappe. Nettoyer le rail, vérifier la nappe.

**E225-0000 — Intensité lampe CIS**
- Lampe/ligne CIS faible : nettoyer la vitre et la bande blanche,
  si persistant → unité CIS.

**E301-0001 — Erreur intensité lumineuse**
- Même famille : lampe d'exposition ou CIS selon modèle.

### 5.4 Bloc CONTRÔLEUR / COMMUNICATION

**E240-0000 — Erreur communication contrôleur**
- Carte principale ↔ sous-cartes : connecteurs, nappes, puis carte.
  Souvent un connecteur mal remis après intervention.

**E351-0000 — Erreur communication contrôleur principal**
- Redémarrer ; si persistant, carte principale (vérifier les
  alimentations d'abord).

**E602-0001 / 0002 / 0003 — Erreur disque dur**
- **Le classique après micro-coupure.** 0001 = boot impossible,
  0002 = erreur lecture/écriture, 0003 = secteurs défectueux.
- Action : mode service > FUNCTION > SYSTEM > CHK-HDD (vérification),
  si échec → remplacer le HDD/SSD, réinstaller le firmware (SST),
  restaurer la sauvegarde.
- **Prévention** : onduleur obligatoire, arrêt propre
  (pas de coupure sauvage).

**E604-0000 — Mémoire insuffisante**
- Travail trop lourd (fichier énorme) : augmenter la RAM si possible,
  simplifier le document, redémarrer.

**E610-xxxx — Erreur chiffrement disque**
- Clé TPM/pile : vérifier la pile de sauvegarde, réinitialiser le
  chiffrement (perte des données !).

**E710-0000 / E711-0000 — Erreur IPC (communication interne)**
- Initialisation ou communication entre cartes : connecteurs, puis
  cartes (DC controller, main controller).

**E732-0001 — Erreur communication lecteur**
- Nappe lecteur ↔ contrôleur : vérifier/rebrancher.

**E733-0001 — Erreur communication imprimante**
- Côté moteur d'impression : idem, connecteurs puis cartes.

**E744 / E747 / E748 / E749 — Cartes option**
- Langue, carte UFR, carte contrôleur, carte PDL : réinstaller le
  firmware du module, vérifier la carte.

### 5.5 Bloc ADF (chargeur)

**E400-0001 / E401-xxxx — Erreur communication ADF**
- Nappe ADF : vérifier le connecteur (souvent pincé à la charnière).

**E402-0000 — Moteur courroie ADF**
**E404-0000 — Moteur de sortie ADF**
**E405-0000 — Moteur de séparation ADF**
- Tester chaque moteur en mode service (FUNCTION > PART-CHK),
  vérifier les courroies et rouleaux ADF (usure = bourrages 00xx).

**E412-0000 — Ventilateur ADF**
- Souffler/nettoyer, remplacer si bloqué.

### 5.6 Bloc FINISHER (trieur/agrafage)

**E500-0001 / E501 — Erreur communication finisher**
- Câble finisher ↔ copieur : rebrancher (machine éteinte).

**E514-8000 — Moteur de sortie du finisher**
**E530-xxxx — Alignement arrière**
**E540-xxxx — Déplacement agrafeuse**
**E542-xxxx — Moteur agrafeuse**
- Vérifier les butées mécaniques (souvent une agrafe coincée),
  les capteurs (poussière de papier), puis les moteurs en PART-CHK.
- **Agrafes** : utiliser les agrafes Canon d'origine (les compatibles
  se coincent et faussent les capteurs).

**E590-xxxx — Perforatrice**
- Bac de confettis plein (le classique oublié !), capteurs sales.

### 5.7 Bloc VENTILATEURS / DIVERS

**E805-0000 — Ventilateur d'évacuation chaleur**
**E824-0000 — Ventilateur de charge primaire**
- Nettoyer (poussière bloquante), tester, remplacer. Un ventilateur
  mort = surchauffe = E000 en cascade.

**E808-0000 — Entraînement fixing**
- Moteur/engrenages de la fusion : vérifier la charge mécanique
  (unité de fusion grippée ?).

**E719-0000 — Monnayeur / lecteur de cartes**
- Périphérique de paiement : vérifier la connexion, désactiver si
  non utilisé (le code peut bloquer la machine).

### 5.8 Méthode générale face à un E-code

```
1. Noter le code COMPLET (EXXX-YYYY) + circonstances
2. Éteindre, attendre 2 min, rallumer → si disparu, surveiller
3. Si retour : mode service > COPIER > FUNCTION > CLEAR > ERR
4. Si persistant : consulter le manuel de service du modèle
5. Tester le composant (PART-CHK), vérifier connecteurs/nappes
6. Remplacer la pièce (toujours machine ÉTEINTE et débranchée)
7. CLEAR > ERR, test complet, noter l'intervention
```

---

## 6. Bourrages papier — jam codes

> Format type : `0AXX` ou `01XX`. Les 2 premiers chiffres = **zone**,
> les 2 derniers = **type**. Exemple : `0108` = zone pickup, bourrage
> au retrait.

### 6.1 Zones (2 premiers chiffres)

| Code zone | Localisation |
|---|---|
| `01xx` | Pickup (cassette → rouleaux d'entraînement) |
| `02xx` | Alimentation / registration (avant transfert) |
| `03xx` | Sortie / fusion (fixing → bac de sortie) |
| `0Axx` | Duplex (recto-verso) |
| `0Dxx` | Finisher / sortie option |
| `00xx` | ADF (chargeur de documents) |

### 6.2 Types (2 derniers chiffres, les plus courants)

| Code | Signification | Action |
|---|---|---|
| `xx01` | Retard (delay jam) : le papier n'arrive pas au capteur à temps | Rouleaux usés, papier humide, guides mal réglés |
| `xx02` | Stationnaire : le papier reste trop longtemps sur un capteur | Capteur sale/défectueux, papier coincé |
| `xx08` | Échec de prise (pickup retry) | Rouleaux pickup lisses → remplacer/nettoyer |
| `xx0B`–`xx0E` | Multi-prise, longueur anormale | Ventiler le papier, séparation usée |

### 6.3 Dépannage par zone

**Zone 01 (pickup)** — le plus fréquent :
- Rouleaux pickup lisses → nettoyer à l'alcool isopropylique ou
  remplacer (kit rollers). Test : frottez le doigt, si ça glisse
  sans accrocher, c'est mort.
- Séparateur (separation pad/roller) usé → multi-prises.
- Papier humide ou guides lâches.

**Zone 02 (registration)** :
- Rouleaux de registration sales → nettoyer.
- Capteurs poussiéreux → souffler.

**Zone 03 (fusion/sortie)** :
- Unité de fusion : doigts de séparation (picker fingers) usés,
  film abîmé → le papier s'enroule.
- Température : vérifier les thermistances (E002 en arrière-plan ?).

**Zone 0A (duplex)** :
- Chemin duplex obstrué (morceau de papier), rouleaux inverseurs.

**Zone 00 (ADF)** :
- Rouleaux ADF usés, guides du chargeur mal réglés, document
  froissé/agrafé (interdire les agrafes !).

