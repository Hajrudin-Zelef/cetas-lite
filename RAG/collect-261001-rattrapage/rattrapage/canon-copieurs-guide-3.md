---
id: collect-261001-rattrapage/rattrapage/canon-copieurs-guide-3
title: "Canon — Guide ultra-complet copieurs (maintenance au cœur)"
domain: rattrapage
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-rattrapage/canon_copieurs_guide.md
source_anchor: ""
source_lines: [384, 554]
sha256: 700b581d93ead216dc2c3853c1905cd9d14f99de45bd3ab35d8c66eb03ec4087
---

# Canon — Guide ultra-complet copieurs (maintenance au cœur)

### 6.4 Extraction d'un bourrage — les règles
1. **Suivre le chemin indiqué à l'écran** (les trappes sont numérotées).
2. Tirer **dans le sens du papier**, doucement, jamais à contre-sens
   (ça casse les capteurs).
3. Vérifier qu'il ne reste **aucun morceau** (un bout de 2 cm suffit
   à re-bourrer).
4. Refermer **toutes** les trappes (un capteur de porte ouvert =
   bourrage fantôme).
5. Si bourrages récurrents au même endroit → noter le jam code et
   traiter la cause (rouleaux, capteur, fusion).

---

## 7. Qualité d'image — diagnostic par symptôme

### 7.1 Mesurer l'intervalle des défauts répétitifs
Un défaut qui se répète à intervalle régulier = un rouleau/tambour
abîmé. Mesurez la distance entre deux défauts :
- ~94 mm → tambour OPC (circonférence typique).
- ~75–80 mm → rouleau de charge.
- ~60 mm → rouleau de développement.
- Intervalle = π × diamètre du composant (vérifiez au manuel).

### 7.2 Tableau des symptômes

| Symptôme | Causes probables | Actions |
|---|---|---|
| **Fond gris/voile** | Développeur usé, toner incompatible, charge primaire faible, bac usagé plein | Ajustement densité, remplacer développeur, toner origine |
| **Image pâle** | Toner bas, laser sale, transfert faible, développeur | Toner, nettoyer optiques laser, vérifier ITB |
| **Lignes verticales blanches** | Corona sale, LED/laser obstrué, tambour rayé | Nettoyer corona, inspecter tambour |
| **Lignes verticales noires** | Raclette abîmée, tambour rayé, rouleau de charge | Remplacer tambour/unité |
| **Points/taches répétitifs** | Tambour piqué, rouleau de fusion abîmé | Mesurer l'intervalle (§7.1) |
| **Bandes horizontales** | Charge, développeur, alimentation laser | Nettoyer, tester par couleur |
| **Couleurs décalées** | ITB, capteurs de registration, vibrations | Calibration auto (AJUST), nettoyer capteurs |
| **Papier ondulé** | Papier humide, fusion trop chaude | Papier sec, vérifier thermistances |
| **Toner qui s'efface au doigt** | Fusion froide (E002 latent), vitesse/papier épais mal réglés | Température fusion, type papier déclaré |
| **Copie ADF floue/décalée** | Vitre d'exposition (bande étroite) sale | Nettoyer la **bande blanche** et la vitre fine ADF |
| **Fond noir total** | Laser/exposition HS, charge primaire | Vérifier laser, E-codes |

### 7.3 Calibrations (à faire après chaque intervention image)
- **Ajustement auto des couleurs** : mode service ou menu utilisateur
  (Réglages > Ajustement > Calibration).
- **Densité** : FUNCTION > DENS (par couleur).
- Toujours calibrer après remplacement tambour/développeur/ITB.

---

## 8. Mode service — le cockpit du technicien

### 8.1 Accès
```
1. Appuyer sur * (ou Paramètres/Enregistrement)
2. Appuyer simultanément sur 2 et 8
3. Rappuyer sur * (ou Paramètres)
→ Menu SERVICE MODE
```
- Niveaux : **COPIER** (moteur), **FEEDER** (ADF), **SORTER** (finisher),
  **FAX**, **BOARD**, etc.
- **Ne touchez jamais** aux réglages d'usine (ADJUST) sans noter
  l'ancienne valeur. Photo avant chaque modif.

### 8.2 Les menus essentiels

| Menu | Usage |
|---|---|
| `COPIER > DISPLAY > ERR` | Historique des erreurs |
| `COPIER > DISPLAY > JAM` | Historique des bourrages |
| `COPIER > DISPLAY > ALARM` | Alarmes |
| `COPIER > FUNCTION > CLEAR > ERR` | Effacer un code erreur |
| `COPIER > FUNCTION > CLEAR > JAM-HIST` | Effacer historique bourrages |
| `COPIER > FUNCTION > PART-CHK` | Tester moteurs/capteurs un par un |
| `COPIER > FUNCTION > DENS` | Ajustement densité |
| `COPIER > COUNTER > TOTAL` | Compteurs |
| `COPIER > COUNTER > PARTS` | Vie des pièces d'usure |
| `COPIER > ADJUST` | Réglages fins (noter avant !) |
| `COPIER > OPTION` | Options machine (à manipuler avec précaution) |

### 8.3 PART-CHK — tester sans démonter
- Active chaque moteur/embrayage/capteur individuellement :
  idéal pour isoler un E-code (ex. : E404 → tester le moteur de
  sortie ADF, l'entendre tourner ou pas).
- **Attention** : certains tests font bouger des mécanismes —
  mains hors de la machine.

---

## 9. Procédures de remplacement (grandes classiques)

> Machine **éteinte et débranchée**, EPI (masque, gants), toner =
> ne pas inhaler, ne pas jeter au feu.

### 9.1 Unité de fusion
1. Ouvrir la trappe latérale/arrière (selon modèle), débrancher les
   connecteurs (thermistances, lampe).
2. Dévisser (2–4 vis), sortir l'unité en la tenant par les poignées.
3. Monter la neuve, rebrancher, revisser.
4. CLEAR > ERR, **remettre le compteur à zéro** (COUNTER > PARTS >
   FIXING > clear), test 20 copies, vérifier la température
   (DISPLAY > ANALOG > température fixing).

### 9.2 Tambour / unité tambour
1. Sortir la cartouche de toner, puis l'unité tambour (ne jamais
   toucher la surface verte/bleue, ne pas l'exposer à la lumière >
   5 min — la couvrir d'un chiffon).
2. Monter la neuve, remettre le toner.
3. Reset compteur PARTS > DRUM, calibration couleurs, test mire.

### 9.3 Développeur
1. Sortir l'unité de développement, vider l'ancien développeur
   (sac étanche, ne pas respirer).
2. Verser le neuf **lentement et uniformément**, remonter.
3. **Initialisation obligatoire** : FUNCTION > INSTALL > STIR ou
   TONER-S (selon modèle) — agite et calibre le mélange.
   Sans init, densité fausse garantie.
4. Reset PARTS > DEVELOPER, calibration.

### 9.4 Courroie de transfert (ITB)
1. Sortir l'unité ITB (souvent par l'avant après retrait des
   tambours), débrancher le moteur.
2. Ne pas toucher la courroie avec les doigts (graisse = défauts).
3. Remonter, reset PARTS > ITB, **registration couleurs**
   (FUNCTION > ADJUST > REGIST).

### 9.5 Rouleaux d'entraînement (kit rollers)
1. Cassettes dehors, localiser pickup/feed/separation rollers.
2. La plupart se déclipsent sans outils (languette).
3. Remplacer par kit complet (pickup + feed + separation),
   reset si compteur associé, test multi-pages.

---

## 10. Réseau, scan, impression

### 10.1 Installation réseau
- IP fixe recommandée (DHCP = l'imprimante « disparaît » au
  renouvellement). Réglages > Réseau > TCP/IP.
- Pilotes : **UFR II** (standard Canon), **PCL6** (universel),
  **PostScript** (option sur certains modèles, pour la PAO).
- Toujours prendre le pilote sur **canon.com** (pas le CD fourni,
  obsolète).

### 10.2 Scan vers dossier (SMB) — la panne n°1 en entreprise
- Prérequis : dossier partagé Windows, utilisateur dédié
  (ex. `scan` / mot de passe **qui n'expire pas**), droits en écriture.
- Sur le copieur : carnet d'adresses > nouveau > dossier SMB >
  `\\SERVEUR\scans`, identifiants.
- **Pannes classiques** :
  - Windows a mis à jour → SMBv1 désactivé : le vieux copieur ne
    parle que SMBv1 → **activer SMBv2/3 sur le copieur** (firmware
    récent) ou forcer SMBv2 côté Windows.
  - Mot de passe expiré → erreur d'authentification.
  - Antivirus/pare-feu bloque le partage.
- Alternative robuste : **Scan to FTP** ou **Scan to Email**.

### 10.3 Scan to Email
- SMTP du FAI ou Office 365 / Google (avec **mot de passe d'application**,
  pas le mot de passe normal).
- Ports : 587 (STARTTLS) ou 465 (SSL). Tester depuis le copieur
  (bouton Test).
- Limite de taille : régler la résolution (300 dpi suffit pour du
  bureau, pas 600).

### 10.4 Impression mobile / cloud
- **Canon PRINT** (appli), AirPrint, Mopria : activer dans les réglages
  réseau. Pratique, mais à sécuriser (impression suiveuse avec badge
  = **uniFLOW**, la solution Canon entreprise).

---

## 11. Firmware — SST (Service Support Tool)

