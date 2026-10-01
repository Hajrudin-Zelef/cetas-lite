---
id: collect-261001-rattrapage/rattrapage/canon-copieurs-guide-9
title: "Canon — Guide ultra-complet copieurs (maintenance au cœur)"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/canon_copieurs_guide.md
source_anchor: ""
source_lines: [1387, 1515]
sha256: ba679499e13a6430cef2c964825436f7db730090b2972d94cd409fd11b67a5c6
---

# Canon — Guide ultra-complet copieurs (maintenance au cœur)

| Pièce | iR 2520 (N&B A3) | iR-ADV C5500 (coul. A3) | DX C5800 (coul. A3) |
|---|---|---|---|
| Toner (cartouche std) | ~15k pages | ~20–30k (N) / ~15–20k (C/M/J) | idem C5500 |
| Tambour | ~150k | ~150–200k par couleur | ~150–200k |
| Développeur | ~300k | ~300–400k par couleur | ~300–400k |
| Unité de fusion | ~300k | ~300–400k | ~400–500k |
| ITB | — (transfert direct) | ~300k | ~300–400k |
| Kit rollers | ~100k | ~120k | ~150k |
| Bac toner usagé | ~50k | ~50–80k | ~50–80k |

## 49. Les 10 fautes professionnelles à ne jamais commettre

1. Monter un toner compatible chez un client sous contrat.
2. Oublier le reset PARTS après remplacement (corrections d'usure fausses).
3. Toucher la surface du tambour à mains nues.
4. Couper la machine en arrachant la prise (HDD).
5. Modifier un réglage ADJUST sans noter l'ancienne valeur.
6. Remonter sans rebrancher une nappe (E240 garanti).
7. Jeter un disque dur de copieur sans effacement.
8. Promettre un SLA intenable pour signer.
9. Intervenir sur la HVT sans qualification (haute tension).
10. Partir sans faire signer le rapport d'intervention.

## 50. Dépannage par téléphone — le script

Avant de vous déplacer, 5 minutes au téléphone économisent 2 heures :
1. « Que dit **exactement** l'écran ? » (code ou message)
2. « Depuis quand ? Après quel événement ? » (orage, déménagement, nouveau papier)
3. « Le papier utilisé a-t-il changé ? » (50 % des cas)
4. « Les trappes sont-elles bien fermées ? » (bourrage fantôme)
5. « Éteignez, attendez 2 minutes, rallumez — que se passe-t-il ? »
→ Si le code persiste : préparez la pièce probable **avant** de partir
(consultez §5/§17 avec le code donné).

---

---

## 51. Magasins papier grande capacité (LCT) et options

- **LCT (Large Capacity Tray)** : 2 000–3 000 feuilles. Ses pannes
  propres : moteur de montée du plateau (le plateau ne monte plus →
  0108), capteur de niveau, rouleaux spécifiques.
- **Ponte de passage (pass-through)** : entre copieur et finisher —
  bourrages 0Dxx fréquents si mal aligné après déplacement.
- **Module d'insertion** : bourrages si papier différent non déclaré.
- Règle : après tout déplacement de la machine, **réalignez les
  modules** (LCT, finisher) — 1 cm de décalage = bourrages en série.

## 52. Mots de passe et verrouillages à connaître

| Accès | Défaut usine (à changer !) |
|---|---|
| Mode service | `* 2 8 *` (pas de mot de passe, accès physique) |
| Admin système (Settings/Registration) | ID : `7654321` / PIN : `7654321` (séries récentes) |
| Anciennes séries | ID : `1` / PIN : `1111` ou sans |

- **Changez les mots de passe admin** à l'installation (sinon
  n'importe qui modifie le réseau/SMTP).
- Notez-les dans votre dossier client (chiffré) — un admin perdu =
  reset usine = perte de tous les réglages.

---

---

## 53. Que mettre dans la valise avant de partir (par symptôme)

> Le technicien qui part sans la pièce probable perd une demi-journée.
> Préparez d'après l'appel (§50).

| Symptôme annoncé | Emporter |
|---|---|
| E000/E001/E002 | Unité de fusion (ou thermistances + lampe) |
| E020/E025 | Toner origine + développeur |
| E202/E225 | Kit de nettoyage optiques + nappe ADF |
| E400–E413 | Kit rollers ADF + courroies |
| E514/E540/E590 | Agrafes origine + kit de nettoyage capteurs |
| E602 | HDD/SSD compatible + clé USB SST + firmware |
| E805/E824 | Ventilateurs (2 modèles courants) |
| Bourrage 0108 | Kit rollers complet + rame de papier de test |
| Bourrage 0301 | Unité de fusion ou doigts de séparation |
| Fond gris | Toner origine + développeur + outil nettoyage corona |
| Lignes verticales | Tambour de la couleur concernée |
| Scan SMB mort | PC portable (test réseau, firmware) |
| N'imprime plus | Câble réseau, pilote sur clé USB |

---

---

## 54. Plan de formation — technicien débutant (4 semaines)

- **Semaine 1** : théorie (§2), démontage/remontage (§37), nettoyages
  (§4), changement toner et bourrages (§6) — sur machine d'atelier.
- **Semaine 2** : mode service (§8), E-codes courants (§5),
  remplacement rollers et tambour (§9), qualité d'image (§7).
- **Semaine 3** : réseau et scan (§10, §33), firmware SST (§11),
  fusion et ITB (§9), sécurité (§26).
- **Semaine 4** : tournées accompagnées, rapports d'intervention
  (§43), script téléphonique (§50), puis autonomie sur les cas
  simples avec un senior en backup téléphonique.

---

---

## 55. Pense-bête de poche (à garder sur soi)

```
Mode service : * 2 8 *
Clear erreur : COPIER > FUNCTION > CLEAR > ERR
Historique   : DISPLAY > ERR / JAM / ALARM
Compteurs    : COPIER > COUNTER > TOTAL / PARTS
Test moteurs : FUNCTION > PART-CHK
Capteurs     : DISPLAY > SENSOR
Effacer HDD  : FUNCTION > SYSTEM > HDD erase
Jam 01xx pickup / 02xx registration / 03xx fusion / 0Axx duplex
E0xx fusion / E02x toner / E40x ADF / E5xx finisher / E6xx disque
94 mm = tambour / 60 mm = développeur (intervalles de défauts)
Règle d'or : photo avant, reset PARTS après, rapport signé.
```

---

*Fin du guide — 1500+ lignes. La maintenance Canon tient en une phrase :
**préventif régulier, pièces d'origine, papier de qualité, firmware à jour,
et tout noter.** Le reste, c'est de l'expérience — et ce guide est fait
pour l'accélérer.*
