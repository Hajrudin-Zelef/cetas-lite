---
id: collect-261001-rattrapage/rattrapage/huawei-licences-tac-guide-7
title: "Licences, support TAC et RMA Huawei — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: ["2026-09-27"]
keywords: ["license", "licenses"]
source: docs/RAG/collect-261001-rattrapage/huawei_licences_tac_guide.md
source_anchor: ""
source_lines: [966, 1158]
sha256: 8d018818b4321799e15abeb2e5d6f50d4c90166731cefbff7575888fea04e4a8
---

# Licences, support TAC et RMA Huawei — Guide ultra-complet

- [ ] Le problème est-il reproductible ? (oui / non / intermittent)
- [ ] Avez-vous le modèle exact, la version VRP, l'ESN/S/N ? (section 69)
- [ ] Avez-vous collecté `display diagnostic-information` ? (section 76)
- [ ] Avez-vous cherché dans la base de connaissances ? (section 58)
- [ ] Quel est l'impact métier ? (qui est touché, depuis quand)
- [ ] Avez-vous un contournement temporaire en place ?
- [ ] Qui est joignable 24h/24 si le TAC rappelle ? (nom + téléphone direct)

## 63. Les niveaux de sévérité S1 à S4 : tableau général

D'après la documentation officielle des services Hi-Care :

| Sévérité | Nom usuel | Fenêtre de couverture | Temps de réponse contractuel* |
|---|---|---|---|
| **S1** | Critique | 24x7 | **30 minutes** |
| **S2** | Majeure | 24x7 | **60 minutes** |
| **S3** | Moyenne | 24x7 (traitement) | **2 heures** |
| **S4** | Faible / demande | 24x7 (traitement) | **Jour ouvré suivant (NBD)** |

*\* Temps entre l'acceptation de la demande par le TAC et le premier contact
d'un ingénieur. Les temps de résolution dépendent de la complexité.*

## 64. S1 — Critique : quand le réseau est à terre

**Définition :** panne avec **impact critique sur l'activité** — réseau ou
environnement hors service.

Exemples :
- L'USG6000 du site principal est HS, tout le site est isolé.
- Les deux membres d'un cluster HA sont tombés.
- Faille de sécurité en cours d'exploitation avérée.

**Ce que Huawei s'engage à faire :**
- Ressources **dédiées**, travail en **24x7 jusqu'à résolution ou
  contournement**.
- Tous les efforts commercialement raisonnables pour fournir un contournement.

**Ce que VOUS devez faire :**
- Ressources disponibles **24x7** (quelqu'un doit décrocher quand le TAC
  rappelle, même à 3h du matin).
- Fournir les infos de diagnostic nécessaires.

**Points de vigilance :**
- Si un contournement est mis en place, la sévérité est **abaissée à S3**
  pour la recherche de la cause racine.
- Si l'ingénieur TAC ne parvient pas à vous joindre **dans l'heure**, la
  sévérité est **temporairement abaissée** jusqu'au rétablissement du contact.
  → En S1, le téléphone d'astreinte doit être sur vous, chargé, avec du réseau.

## 65. S2 — Majeure : dégradation importante

**Définition :** fonctionnalité majeure dégradée, impact significatif mais le
réseau n'est pas totalement à terre.

Exemples :
- Un des deux USG du cluster HA est tombé (le second tient la charge).
- La moitié des AP d'un bâtiment ne montent plus.
- Débit divisé par deux sur un lien critique, sans coupure totale.

**Engagements :** ressources disponibles 24x7, réponse en 60 minutes,
efforts raisonnables pour un contournement.

## 66. S3 — Moyenne : sans impact critique

**Définition :** problème qui **n'affecte pas les fonctions critiques** —
dégradation de capacité, de mesure, ou gêne limitée.

Exemples :
- Un AP isolé sur 50 ne répond plus (couverture redondante).
- Les logs ne remontent plus vers le syslog (le réseau fonctionne).
- Question de configuration sans urgence.

**Engagements :** traitement pendant les heures ouvrées locales jusqu'à
résolution ou contournement, réponse en 2 heures.

## 67. S4 — Faible : demandes et questions

**Définition :** demande d'information, documentation, question de
fonctionnement, anomalie cosmétique.

Exemples :
- « Comment activer telle fonction sur l'AR720 ? »
- Coquille dans une page web d'administration.
- Demande de documentation.

**Engagements :** réponse le jour ouvré suivant (NBD).

## 68. Bien choisir sa sévérité (et ne pas en abuser)

| Bonne pratique | Mauvaise pratique |
|---|---|
| S1 uniquement si l'activité est réellement bloquée | Mettre S1 « pour aller plus vite » sur un S3 |
| Décrire l'impact métier factuellement | Exagérer l'impact pour forcer la priorité |
| Accepter la requalification proposée par le TAC si justifiée | S'obstiner sur S1 quand un contournement existe |

**Pourquoi c'est important :** abuser des S1 décrédibilise vos vrais S1.
Les équipes TAC notent l'historique. Le jour où vous aurez un vrai S1 un
dimanche à 2h du matin, vous voulez que le mot « critique » venant de vous
soit pris au sérieux immédiatement (voir cas vécu n°108).

## 69. Les informations OBLIGATOIRES du ticket

Un ticket sans ces éléments fera des allers-retours. Fournissez **d'emblée** :

1. **Modèle exact** de l'équipement (ex. : `USG6680`, `S310-24P`).
2. **Version logicielle** (ex. : `VRP V600R024C10` — via `display version`).
3. **ESN et/ou numéro de série**.
4. **Contrat de support** (n° Hi-Care, ou « sous garantie standard »).
5. **Description du problème** : quoi, quand, fréquence, reproductible ?
6. **Impact métier** : qui est touché, depuis quand, contournement en place ?
7. **Logs** : `display diagnostic-information` en pièce jointe (section 76).
8. **Topologie** : schéma simplifié du site concerné.
9. **Contact joignable** : nom, téléphone direct, email, fuseau horaire,
   disponibilités.
10. **Actions déjà tentées** : reboot ? changement de câble ? rollback ?

## 70. Template de ticket TAC en anglais (vierge, prêt à copier)

```text
Subject: [S2] USG6000 - IPS license expired, UTM blades down - Site Paris

1. Customer information
   Company: [COMPANY NAME]
   Contact name: [FULL NAME]
   Phone (24/7 reachable): [PHONE NUMBER]
   Email: [EMAIL]
   Time zone / availability: [e.g. CET, available 24/7 for S1/S2]

2. Equipment information
   Product model: [e.g. USG6680]
   Software version: [e.g. VRP V600R024C10 - from "display version"]
   ESN: [from "display esn"]
   Serial Number: [from label or "display device"]
   Support contract: [Hi-Care contract number or "standard warranty"]

3. Problem description
   Summary: [one sentence]
   First occurrence: [date + time with time zone]
   Frequency: [permanent / intermittent - how often / reproducible: yes-no]
   Error messages (exact text):
   [paste exact messages]

4. Business impact
   Impacted users/services: [who / what]
   Workaround in place: [yes - describe / no]
   Requested severity: [S1/S2/S3/S4 + one-line justification]

5. Troubleshooting already performed
   - [action 1 + result]
   - [action 2 + result]

6. Attachments
   - diagnostic-information output: [filename]
   - network topology diagram: [filename]
   - screenshots: [filenames]

7. Requested action
   [e.g. "Please provide a workaround and root cause analysis."]
```

## 71. Template rempli — exemple fictif n°1 (licence UTM)

```text
Subject: [S2] USG6000 - IPS/AV licenses expired, UTM inspection stopped - Site Paris

1. Customer information
   Company: ACME Industries (fictitious)
   Contact name: Jean Dupont (fictitious)
   Phone (24/7 reachable): +33 6 00 00 00 00 (fictitious)
   Email: jean.dupont@example.com (fictitious)
   Time zone / availability: CET, available 24/7 for S1/S2

2. Equipment information
   Product model: USG6680 (fictitious model reference)
   Software version: VRP V600R024C10 (from "display version")
   ESN: 2102311HSL8H6000123 (fictitious)
   Serial Number: 2102311HSL8H6000123 (fictitious)
   Support contract: Hi-Care Standard 9x5xNBD, contract #HC-2026-4451 (fictitious)

3. Problem description
   Summary: IPS, AV and URL filtering blades stopped working this morning.
   First occurrence: 2026-09-27 08:15 CET
   Frequency: permanent
   Error messages (exact text):
   "License of IPS will expire in 0 days" then "The license file has expired."

4. Business impact
   Impacted users/services: 350 users on Paris site without UTM inspection.
   Basic firewalling still works.
   Workaround in place: no - renewal .dat file requested from partner,
   pending reception.
   Requested severity: S2 - security function down on main site, no workaround.

