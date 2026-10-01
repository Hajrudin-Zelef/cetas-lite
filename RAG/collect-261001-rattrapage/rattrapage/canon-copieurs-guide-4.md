---
id: collect-261001-rattrapage/rattrapage/canon-copieurs-guide-4
title: "Canon — Guide ultra-complet copieurs (maintenance au cœur)"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/canon_copieurs_guide.md
source_anchor: ""
source_lines: [555, 654]
sha256: aae002c9b07cd4671d0b496c7295d304a024d206d224bdd3254ccd554df752de
---

# Canon — Guide ultra-complet copieurs (maintenance au cœur)

- L'outil Canon **SST** (PC Windows + câble USB/réseau) flashe le
  firmware : indispensable après remplacement HDD/carte.
- **Règles** :
  1. Toujours sauvegarder (carnet d'adresses, réglages) avant.
  2. Ne jamais couper pendant le flash (onduleur !).
  3. Prendre le firmware **exact** du modèle et de la région.
  4. Après flash : vérifier la version, restaurer les réglages, tester.
- Firmwares récents = correctifs SMB, sécurité (failles), compatibilité
  Windows. Un copieur de 2018 non flashé **ne scannera plus** vers un
  Windows 11 à jour.

---

## 12. Compteurs et gestion

- **Compteurs** : total, par format, couleur/N&B, scan, fax.
  Relevé mensuel pour la facturation au coût-page.
- **Department ID** : codes par service/utilisateur (quotas,
  imputation). À activer pour les clients qui veulent contrôler.
- **Rapports** : la machine peut envoyer le compteur par email
  (utile pour la facturation sans déplacement).

---

## 13. Sécurité des données

- Les copieurs **stockent tout** : copies, scans, impressions sur le
  disque dur. Un copieur réformé = une fuite de données.
- **Avant revente/recyclage** : mode service > FUNCTION > SYSTEM >
  **HDD erase** (effacement complet, plusieurs passes) ou remplacement
  physique du disque.
- **Chiffrement disque** : à activer (avec sauvegarde de la clé !).
- **McAfee embarqué** (séries DX) : vérifie l'intégrité au boot.
- Politique : effacer les carnets d'adresses (emails, dossiers SMB
  avec mots de passe !) avant tout départ de machine.

---

## 14. Dépannage express par symptôme

| Symptôme | Vérifier d'abord |
|---|---|
| Ne s'allume pas | Prise, interrupteur, fusible interne, carte alim |
| Écran noir mais ventilo tourne | Carte contrôleur, barrette RAM, HDD |
| Bourrage à chaque copie | Rouleaux pickup, papier humide, guides |
| Copies pâles | Toner, densité, laser sale |
| Copies trop foncées / fond gris | Développeur, charge, toner compatible |
| Lignes verticales | Tambour, raclette, corona |
| Ne scanne plus vers dossier | SMB (v1/v2), mot de passe, réseau |
| N'imprime plus réseau | IP (DHCP ?), pilote, file d'attente, pare-feu |
| Bruit anormal | Ventilateur, engrenages, unité de fusion |
| Surchauffe / odeur | Fusion, ventilateurs → **éteindre, intervenir** |
| Code E persistant | §5, CLEAR > ERR, puis pièce |
| Lent / se bloque | HDD (E602 latent), mémoire, firmware |

---

## 15. Boîte à outils du technicien Canon

- Tournevis cruciforme/plat (jeux), pince fine, pince coupante.
- **Aspirateur à toner** (filtre HEPA — jamais d'aspirateur normal,
  le toner traverse les filtres et encrasse le moteur).
- Alcool isopropylique + chiffons non pelucheux (optiques, rouleaux).
- Air comprimé sec + masque FFP2.
- Multimètre (tensions, continuité lampes).
- Kit rollers universel Canon, agrafes d'origine, toner d'origine.
- PC portable + **SST** + firmwares des modèles suivis (sur clé USB).
- Lampe frontale (l'intérieur d'un copieur est sombre).
- Carnet : noter **chaque** intervention (date, compteur, pièces,
  codes) — c'est votre mémoire et votre argument commercial.

---

## 16. Glossaire

| Terme | Signification |
|---|---|
| OPC | Tambour photoconducteur organique |
| ITB | Courroie de transfert intermédiaire (couleur) |
| Fixing / Fusion | Unité qui fixe le toner par chaleur+pression |
| Developer | Poudre magnétique qui transporte le toner |
| Corona | Fil/charge de la charge primaire |
| ADF / DADF | Chargeur automatique (recto / recto-verso) |
| Finisher | Module de finition (tri, agrafage, pliage) |
| SST | Service Support Tool (flash firmware Canon) |
| UFR II | Langage d'impression Canon |
| MEAP | Plateforme d'applications embarquées Canon |
| Department ID | Gestion par codes utilisateurs |
| Pickup roller | Rouleau de prise papier |

---

---

## 17. Dictionnaire E-codes étendu (E012 → E8xx)

> Complément du §5. Format condensé : code → composant → action de
> première intention. Les détails (YYYY) précisent le sous-ensemble :
> notez-les toujours.

