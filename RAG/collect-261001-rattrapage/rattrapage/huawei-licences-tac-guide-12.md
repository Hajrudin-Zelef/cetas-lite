---
id: collect-261001-rattrapage/rattrapage/huawei-licences-tac-guide-12
title: "Licences, support TAC et RMA Huawei — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["arr", "distribution", "license"]
source: docs/RAG/collect-261001-rattrapage/huawei_licences_tac_guide.md
source_anchor: ""
source_lines: [1809, 1963]
sha256: 3481bce8bb6c4dafb3132a076110797d9bc9d20135a9a7804109b188be488959
---

# Licences, support TAC et RMA Huawei — Guide ultra-complet

**Leçon.** 1) eKit = circuit partenaire d'abord (voir partie J). 2) Faites
écrire les délais **du partenaire** dans le contrat : « le partenaire répond
sous Xh ». 3) Conservez quand même vos accès au portail Huawei (docs,
firmwares selon droits).

## 112. Cas n°12 — Panne un jour férié : NBD vs 4h

**Situation.** Jeudi férié, 10h : le S310 d'un bâtiment tombe (panne
matérielle confirmée par le TAC à 11h, n° RMA généré à 11h30). Contrat
**Standard 9x5xNBD**. Le technicien attend la pièce « demain ». Erreur :
vendredi n'est pas férié mais le RMA a été généré un jour férié → enregistré
le vendredi → pièce reçue le **lundi**. Le bâtiment reste sans réseau 4 jours.

**Gestion.** Contournement d'urgence : un switch de secours (spare froid,
acheté d'avance) est configuré et monté le jeudi après-midi. Leçon intégrée
au plan : pour les sites sans redondance, **un spare froid sur étagère**
coûte moins cher que 4 jours d'arrêt.

**Leçon.** 1) Les jours fériés **ne sont pas** des jours ouvrés : NBD après
un férié = +2-3 jours. 2) Pour les équipements en contrat Standard sur site
non redondé : prévoyez un **spare froid** (switch pré-configuré sous
emballage). 3) Les sites vraiment critiques méritent le palier **Premier
24x7x4** (section 84).

---

# I. CHECKLISTS ET CONTACTS

## 113. Checklist « avant d'acheter »

- [ ] Devis détaillé : références, quantités, **licences incluses vs
      optionnelles**, durées de chaque licence.
- [ ] **Proof of Entitlement** promise par écrit (délai de fourniture).
- [ ] Palier Hi-Care choisi **par équipement** (tableau section 84), avec
      dates de début explicites (section 86).
- [ ] Clause de transfert de licence en cas de RMA.
- [ ] Contact support nommé chez le partenaire (nom, téléphone direct).
- [ ] Délais de réponse du partenaire écrits (surtout pour eKit).
- [ ] Vérification : le modèle commandé est-il bien celui dont la doc et les
      licences existent sur le portail (pas de fin de vie imminente) ?

## 114. Checklist « à la réception du matériel »

- [ ] Contrôle colis + photos (chocs, humidité).
- [ ] Vérifier références livrées = bon de commande.
- [ ] **Photographier toutes les étiquettes** (S/N, ESN, modèle) et archiver.
- [ ] Relever `display esn` / `display version` au premier boot.
- [ ] **Enregistrer les S/N sur support.huawei.com** (section 49).
- [ ] Réclamer immédiatement toute **Proof of Entitlement** manquante.
- [ ] Générer et activer les licences (partie B).
- [ ] Vérifier `display license` : fonctions + dates.
- [ ] Inscrire chaque licence au **registre** (section 38).
- [ ] Inscrire chaque fin de garantie/contrat au **calendrier** (section 34).
- [ ] Tester le matériel avant mise en production (non-DOA).
- [ ] Sauvegarder la configuration initiale.

## 115. Checklist « annuelle » (revue administrative)

- [ ] Audit licences : registre vs `display license` réels (section 40).
- [ ] Lister les expirations (licences + contrats) des 12 prochains mois.
- [ ] Vérifier les garanties par S/N sur le portail (section 88).
- [ ] Chiffrer le budget renouvellements N+1, présenter à la direction.
- [ ] Vérifier les accès au portail (départs, nouveaux arrivants).
- [ ] Tester la procédure de collecte de logs (sections 76–78) sur un
      équipement : le jour de la panne n'est pas le jour de l'apprentissage.
- [ ] Relire les contrats partenaires : délais, contacts, toujours d'actualité ?
- [ ] Contrôler l'état du stock de spares froids.

## 116. Checklist « avant d'ouvrir un ticket TAC »

- [ ] Modèle, version VRP, ESN/S/N, n° de contrat : réunis.
- [ ] `display diagnostic-information` collecté **pendant** le problème.
- [ ] Logbuffer exporté (avant tout reboot).
- [ ] Base de connaissances consultée (section 58).
- [ ] Description en 5W rédigée (section 80).
- [ ] Sévérité choisie factuellement (sections 63–68).
- [ ] Template anglais rempli (section 70).
- [ ] Contact joignable 24/7 désigné (nom + téléphone direct).
- [ ] Équipe/astreinte prévenue du n° de ticket dès ouverture.

## 117. Contacts utiles (template à compléter)

```text
PARTENAIRE / INTÉGRATEUR
  Société : ............................ [À COMPLÉTER]
  Commercial : ......................... [À COMPLÉTER] (tél. : ............)
  Support N1 : ......................... [À COMPLÉTER] (tél. : ............)
  Email support : ...................... [À COMPLÉTER]
  Délai de réponse contractuel : ....... [À COMPLÉTER]

HUAWEI TAC (canaux officiels — numéros et emails exacts par région
  à vérifier sur https://support.huawei.com > Contact)
  Hotline : ............................ [À COMPLÉTER]
  Email : .............................. [À COMPLÉTER]
  Portail tickets : https://support.huawei.com (Service Request)

CONTRATS
  Hi-Care USG6000 site principal : n° ............, palier ............, fin le ............
  Hi-Care USG6000 site secondaire : n° ............, palier ............, fin le ............
  [À COMPLÉTER pour chaque équipement critique]

EN INTERNE
  Responsable licences/contrats : ..... [À COMPLÉTER]
  Astreinte réseau : ................... [À COMPLÉTER]
  Boîte email alertes d'expiration : ... [À COMPLÉTER]
```

> **Ne notez jamais de mots de passe dans ce document.** Les accès portail
> sont nominatifs (section 48).

---

# J. SPÉCIFICITÉS EKIT

## 118. eKit : le circuit partenaire, pas le direct Huawei

La gamme **eKit** est la marque « distribution » de Huawei pour les PME :
switches, routeurs, AP (AP361, AP761), passerelles de sécurité, IdeaHub...
Le modèle commercial est **100 % indirect** : Huawei vend aux distributeurs,
qui vendent aux partenaires, qui vendent et supportent le client final.

**Conséquence n°1 : le circuit de support.**

```
Client final (vous)
    │ 1. Appel / ticket
    ▼
Partenaire / intégrateur local  ← VOTRE point d'entrée unique (niveau 1)
    │ 2. Escalade si nécessaire
    ▼
Distributeur agréé eKit         ← niveau 2, stocks, RMA logistique
    │ 3. Escalade si nécessaire
    ▼
Huawei                          ← niveau 3 (via le canal distribution)
```

- N'appelez pas directement la hotline Huawei Enterprise pour un produit
  eKit : on vous redirigera vers votre partenaire (cas vécu n°111).
- En revanche, le portail https://support.huawei.com reste utile pour la
  **documentation** et (selon droits) les **firmwares**.
- Les outils eKit (application mobile / cloud eKit) simplifient la gestion
  courante : utilisez-les pour le niveau 0 (diagnostic de base).

**Conséquence n°2 : exigez un contrat de support écrit avec le partenaire.**

| Clause | Pourquoi |
|---|---|
| Délai de réponse du partenaire (ex. : 4h ouvrées) | Sans écrit, c'est « dès que possible » |
| Procédure d'escalade vers le distributeur/Huawei | Savoir qui fait quoi, en combien de temps |
| Gestion des RMA eKit (qui porte le colis, délais) | Éviter le ping-pong du cas n°111 |
| Fourniture des preuves de droit de licence | Indispensable (partie B) |
| Contacts nommés + suppléants | Pas de « M. X est en congés, rappelez » |

## 119. eKit : garantie, licences, firmwares — ce qui change (et ce qui ne change pas)

