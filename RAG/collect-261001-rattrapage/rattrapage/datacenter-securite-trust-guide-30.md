---
id: collect-261001-rattrapage/rattrapage/datacenter-securite-trust-guide-30
title: "Sécurité datacenter — HSM, Root of Trust, FPGA, sécurité physique, gestion des clés"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-09-27", "2027-01-01"]
keywords: ["arr"]
source: docs/RAG/collect-261001-rattrapage/datacenter_securite_trust_guide.md
source_anchor: ""
source_lines: [3917, 4006]
sha256: 8f0a2e9c9ee658587cdd1be16b3b687fa519eeeb409d9a6d2e73fd98a02e60db
---

# Sécurité datacenter — HSM, Root of Trust, FPGA, sécurité physique, gestion des clés

```
1. NE PAS redémarrer avant d'avoir les logs (SNMP/syslog).
2. Un seul membre ou tout le cluster ? → supervision.
3. Bascule client sur le 2e membre (procédure §136).
4. Si matériel : appeler le support (n° contrat : ............).
5. Restaurer depuis le HSM backup (procédure testée §138).
6. JAMAIS restaurer depuis une sauvegarde non vérifiée.
ASTREINTE : ............   SUPPORT : ............
```

## Fiche R2 — Coupure électrique, séquence de redémarrage

```
1. Stabiliser le groupe (2-5 min), onduleur en mode normal.
2. Réseau : switches, firewall VLAN crypto.
3. HSM : allumage → vérifier partitions + lire les alertes tamper.
4. Vault/KMS : vérifier l'auto-unseal (sinon Shamir + custodians).
5. Applis : bases TDE, PKI, TSA — dans l'ordre des dépendances.
6. Tests : 1 signature + 1 déchiffrement + CRL/OCSP.
7. PV : durées, écarts, anomalies.
ORDRE INVERSÉ À L'ARRÊT : applis → KMS → HSM → réseau.
```

## Fiche R3 — Clé compromise (premières 30 minutes)

```
1. RÉVOQUER les certificats (CRL/OCSP) — pas de panique, de la méthode.
2. DÉSACTIVER la clé dans le HSM/KMS (ne pas la détruire : forensique).
3. PÉRIMÈTRE : quelles données, quelle période ?
4. RÉGÉNÉRER + RE-CHIFFRER les données exposées.
5. NOTIFIER (direction, juriste, clients selon obligations).
6. POST-MORTEM sous 72 h : comment a-t-elle fuité ?
CELLULE DE CRISE : ............   JURISTE : ............
```

## Fiche R4 — Réception d'un serveur (quarantaine)

```
1. Scellés + S/N = bon de livraison. Photo.
2. Boot en VLAN de quarantaine — PAS en production.
3. UEFI à jour, Secure Boot ON, TPM provisionné.
4. BMC : firmware à jour, mot de passe par défaut CHANGÉ (unique).
5. BMC sur VLAN management, services inutiles OFF, certificat remplacé.
6. Image OS vérifiée (hash), LUKS scellé TPM, SSH par clés.
7. Fiche serveur remplie (§199), PV archivé.
```

## Fiche R5 — Les 5 interdits (à afficher dans le local)

```
⛔ Clé privée en clair hors HSM — JAMAIS, même « temporairement ».
⛔ Mot de passe BMC par défaut — changé à la réception, unique par serveur.
⛔ Équipement crypto sur circuit non secouru — JAMAIS.
⛔ Sauvegarde de clés non testée — test annuel obligatoire.
⛔ Photos en salle serveur — INTERDIT.
En cas de doute : relire le runbook AVANT d'agir.
```

## E. Checklist de relecture avant tout achat (à cocher avec le vendeur)

- [ ] La **référence exacte** (modèle + version firmware) figure sur la liste CMVP du
      NIST pour la certification annoncée — capture d'écran archivée.
- [ ] Le devis détaille : matériel, **support annuel** (taux %, durée, SLA), prestations
      (installation, cérémonie, formation), délais.
- [ ] Le support PQC **logiciel** (ML-KEM/ML-DSA) est confirmé par écrit — pas de
      promesse verbale.
- [ ] Un **POC de 30 jours** est prévu au contrat (protocole §239).
- [ ] La **fin de commercialisation** du modèle est connue (> 3 ans idéalement).
- [ ] Les **consommations** (W) et la plage de température sont compatibles avec vos
      baies et votre clim.
- [ ] Le **circuit de commande** est officiel (distributeur agréé) — facture et
      numéros de série traçables.
- [ ] La **reprise de l'ancien matériel** (si renouvellement) inclut la zeroization
      avec PV.
- [ ] Le **budget temps interne** est chiffré et validé (installation, formation,
      tests, revues) — voir §207 point 6.
- [ ] La **direction** a signé le dossier (besoin, budget, risques résiduels).

> *Cocher ces 10 cases prend une heure et évite 90 % des déconvenues d'achat. Bon courage.*

## F. Journal de maintenance du guide (à remplir par l'équipe)

| Date | Révision | Objet | Par |
|---|---|---|---|
| 27/09/2026 | v1.0 | Création (240 sections, 19 parties, annexes A-F) | Leo |
| _à compléter_ | v1.1 | Revue T1 2027 : FIPS 206, CNSA 2.0 (01/01/2027), évolutions PQC HSM | _équipe_ |
| _à compléter_ | v1.2 | Revue T3 2027 : retour du premier test DR réel, leçons apprises | _équipe_ |
| _à compléter_ | v2.0 | Revue 2028 : état de la migration hybride/PQC, nouveaux produits | _équipe_ |

> *Un guide non maintenu devient faux en 18 mois. Ce journal est le rappel.*
