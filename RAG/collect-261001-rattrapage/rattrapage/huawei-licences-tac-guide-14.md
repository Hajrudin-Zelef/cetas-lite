---
id: collect-261001-rattrapage/rattrapage/huawei-licences-tac-guide-14
title: "Licences, support TAC et RMA Huawei — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["incident", "license"]
source: docs/RAG/collect-261001-rattrapage/huawei_licences_tac_guide.md
source_anchor: ""
source_lines: [2083, 2253]
sha256: 0f98796e9e7e9b45ca8188d2607101eb727c4477fc198b89d9ad077814663c9a
---

# Licences, support TAC et RMA Huawei — Guide ultra-complet

**Pratiques d'équipe recommandées :**
- Mettre en place le script nocturne `display license` + alertes J-90/J-60/J-30.
- Centraliser les syslogs de tout le parc (prérequis du support efficace).
- Maintenir un **spare froid** pour chaque équipement non redondé critique.
- Faire une **répétition annuelle** : un technicien ouvre un ticket « à blanc »
  (S4) pour vérifier que toute la chaîne (accès portail, collecte de logs,
  contacts) fonctionne.
- Formations : cursus HCIA/HCIP Datacom pour l'équipe (le vocabulaire TAC
  et les procédures y sont couverts).

**Autres guides de la série (même dossier) :**
- `onduleurs_ups_guide.md`, `proxmox_guide.md`, `huawei_*_guide.md`
  (AP361, AP761, S310, AR720, USG6000, eKit, eSight).

---

## Conclusion du guide

Licences, TAC et RMA ne sont pas des corvées administratives : ce sont les
**filets de sécurité** de votre production. Un chef de service qui tient son
registre de licences, son calendrier d'échéances et ses checklists RMA dort
mieux — et fait dormir son astreinte. Les pannes arriveront toujours ; la
différence entre une nuit blanche et un incident maîtrisé, c'est ce qui a été
**préparé avant**.

> *Dernière mise à jour : septembre 2026. Vérifiez les valeurs marquées
> « à vérifier sur le portail officiel » avant tout engagement.*

---

# L. ANNEXES OPÉRATIONNELLES

## 124. Script de contrôle automatique des licences (exemple)

Principe : chaque nuit, un script se connecte en SSH aux USG, lance
`display license`, extrait les dates d'expiration et alerte. Exemple de
logique (pseudo-code, à adapter à votre outillage — Ansible, Python/Netmiko,
ou simple expect) :

```text
POUR CHAQUE équipement_critique DANS inventaire:
    sortie = ssh(equipement, "display license")
    POUR CHAQUE ligne "Expire date : AAAA-MM-JJ" DANS sortie:
        jours_restants = date_expiration - aujourd_hui
        SI jours_restants <= 90: alerter("J-90", équipement, fonction)
        SI jours_restants <= 60: alerter("J-60", équipement, fonction)
        SI jours_restants <= 30: alerter("J-30 URGENT", équipement, fonction)
        SI jours_restants <= 0:  alerter("EXPIREE", équipement, fonction)
ENVOYER le rapport à la boîte d'alertes (section 117)
```

Exigences :
- Compte SSH **dédié** en lecture seule sur les équipements.
- La boîte destinataire doit être **surveillée** (pas une boîte générique).
- Tester le script après chaque changement de version VRP (le format de
  sortie peut évoluer).

## 125. Modèle de mail : demande de devis de renouvellement au partenaire

```text
Objet : Demande de devis — renouvellement licences UTM USG6000 (échéance JJ/MM/AAAA)

Bonjour [Nom],

Merci de nous chiffrer le renouvellement suivant :

- Équipement : USG6000 [modèle], site [nom], ESN [ESN]
- Licences : IPS + AV + URL Filtering (références actuelles : [part numbers])
- Durée souhaitée : 3 ans (merci de chiffrer aussi l'option 1 an pour comparaison)
- Date de début souhaitée : [JJ/MM/AAAA] (alignée sur la fin des droits actuels)

Merci de joindre à votre offre :
- les part numbers exacts et les dates de début/fin de droits,
- le délai de fourniture de la Proof of Entitlement après commande.

Échéance interne : devis attendu avant le [JJ/MM/AAAA] (J-60).

Cordialement,
[Nom] — [Fonction] — [Téléphone]
```

## 126. Modèle de mail : signalement au partenaire (circuit eKit)

```text
Objet : [eKit] AP361 site [nom] — panne suspectée, demande de prise en charge

Bonjour [Nom],

Nous rencontrons le problème suivant sur un équipement eKit :

- Modèle : [ex. AP361], S/N : [numéro], site : [nom/adresse]
- Symptôme : [description factuelle, méthode 5W]
- Depuis : [date/heure]
- Actions tentées : [liste]
- Pièces jointes : [diagnostic, photos, logs]

Merci de nous confirmer la prise en charge et le délai d'intervention
conformément à notre contrat de maintenance.

Joignable : [nom, téléphone direct].

Cordialement,
[Nom]
```

## 127. Fiche réflexe « panne à 2h du matin » (à imprimer)

```
┌─────────────────────────────────────────────────────────┐
│ PANNE RÉSEAU — FICHE RÉFLEXE (2h du matin)              │
├─────────────────────────────────────────────────────────┤
│ 1. RESPIRER. Ne rien rebooter en panique.               │
│ 2. Qualifier : quoi / où / depuis quand / qui impacté.  │
│ 3. Y a-t-il un contournement ? (lien secours, HA ?)     │
│    → OUI : l'activer, puis traiter en S2/S3.            │
│    → NON + site à terre : S1.                           │
│ 4. Collecter AVANT tout reboot :                        │
│    display diagnostic-information + display logbuffer    │
│ 5. Ouvrir le ticket (template section 70) ou appeler    │
│    la hotline avec le n° de contrat sous les yeux.      │
│ 6. Prévenir : [nom astreinte N+1 — À COMPLÉTER].         │
│ 7. Noter l'heure de chaque action (journal de bord).    │
│ 8. Ne pas fermer le ticket sans cause racine comprise.  │
├─────────────────────────────────────────────────────────┤
│ N° contrat Hi-Care : ............ [À COMPLÉTER]          │
│ Hotline TAC : ................... [À COMPLÉTER]          │
└─────────────────────────────────────────────────────────┘
```

## 128. Matrice de décision : réparer, RMA, ou acheter neuf ?

| Situation | Décision |
|---|---|
| Sous garantie / Hi-Care, panne matérielle confirmée | **RMA** (partie G) |
| Hors garantie, pièce disponible en spare, panne simple (alim, ventilo) | **Réparer** avec pièce d'origine via le partenaire |
| Hors garantie, carte mère / châssis, devis réparation > 60 % du neuf | **Acheter neuf** (et en profiter pour monter en gamme) |
| Équipement en fin de vie (EoS), panne matérielle | **Acheter neuf** : le RMA devient impossible à terme |
| DOA (défectueux à réception) | **Échange DOA** immédiat via le vendeur (section 106) |
| Dommage exclu de garantie (surtension, choc) | **Acheter neuf** + corriger la cause (parafoudre...) |

## 129. Gérer un parc multi-sites : organisation des dossiers

Arborescence recommandée (sur le partage d'équipe + sauvegarde) :

```text
parc_reseau/
├── site_paris/
│   ├── inventaire.md            (équipements, S/N, ESN, versions)
│   ├── licences/
│   │   ├── USG6000_21023..._IPS.dat
│   │   └── registre_licences.csv (section 39)
│   ├── contrats/
│   │   ├── hicare_2026-2029.pdf
│   │   └── preuves_de_droit/
│   ├── configs/
│   │   └── sauvegardes_AAAAMMJJ/
│   └── tickets/
│       └── SR-XXXX_2026-09-27/
├── site_lyon/
│   └── ... (même structure)
└── _transverse/
    ├── calendrier_echeances.csv (section 34)
    ├── contacts.md (section 117)
    └── procedures/ (ce guide + fiches réflexes)
```

## 130. Sauvegarde des configurations : politique minimale

Sans sauvegarde, un RMA devient une **reconstruction** (erreurs, oublis,
heures perdues). Politique minimale :

