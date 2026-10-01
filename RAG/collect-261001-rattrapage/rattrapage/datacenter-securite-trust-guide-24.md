---
id: collect-261001-rattrapage/rattrapage/datacenter-securite-trust-guide-24
title: "Sécurité datacenter — HSM, Root of Trust, FPGA, sécurité physique, gestion des clés"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Intel"]
dates: []
keywords: ["datacenter", "amd", "arr", "gpu", "intel"]
source: docs/RAG/collect-261001-rattrapage/datacenter_securite_trust_guide.md
source_anchor: ""
source_lines: [3065, 3237]
sha256: f9fce7acd965db7925f89f2ff5a2a8611ab38714b1046f2027f88a53b986efd4
---

# Sécurité datacenter — HSM, Root of Trust, FPGA, sécurité physique, gestion des clés

R = Responsible, A = Accountable, C = Consulté, I = Informé. L'adapter, l'afficher, le
revoir à chaque changement d'équipe. Le cumul R+A sur une seule personne pour la
cérémonie est **interdit**.

## 196. Former l'équipe : plan minimal

- **Admins HSM** (2 pers. min.) : formation constructeur (3-5 jours) + cérémonie
  assistée + habilitation.
- **Exploitants** : runbooks (sections 136-137), exercice de crise annuel
  (section 190).
- **Direction** : sensibilisation 2 h (risques, coûts d'arrêt, responsabilités) —
  c'est elle qui signe les budgets et les PV de cérémonie.
- **Recyclage** : 1 jour/an (nouveautés firmware, retours d'incidents, évolutions PQC).
- **Documentation** : tout est écrit (runbooks, schémas, mots de passe au coffre) —
  l'équipe doit survivre au départ d'un membre clé (**bus factor** ≥ 2 sur chaque
  compétence critique).

## 197. Veille sécurité : s'abonner aux bonnes sources

- **Bulletins constructeurs** : Thales, Utimaco, Entrust, HPE, Dell, Supermicro, AMD,
  Intel — alertes firmware critiques sous 30 jours (section 62).
- **ANSSI** : bulletins, guides d'hygiène (référence FR).
- **NIST** : évolutions FIPS 140-3, PQC (csrc.nist.gov).
- **Communautés** : listes Keylime, EJBCA, Vault (changements, failles).
- **Rituel** : revue mensuelle de 1 h (quoi de neuf ? quoi appliquer ?), décisions
  tracées. La veille non ritualisée n'existe pas.

## 198. Modèle de PV de cérémonie (template)

```
PROCÈS-VERBAL DE CÉRÉMONIE DE GÉNÉRATION DE CLÉS
Date/heure : ............  Lieu : ............
Objet : génération de la clé ............ (usage : ............)
Script version : ............

Participants :
- Officiant 1 : ............ (rôle : ............)
- Officiant 2 : ............ (rôle : ............)
- Témoin/auditeur : ............
- Scribe : ............

Matériel : HSM ............ (S/N : ............), firmware ............,
scellés vérifiés : OK / NOK (photos jointes).

Opérations :
1. ............ (heure : ....)
2. ............ (heure : ....)

Clé générée : algorithme ............, taille ............
Empreinte clé publique (SHA-256) : ............
  (relevée indépendamment par : ............ et ............)

Sauvegarde : méthode ............, vérifiée : OK / NOK
Supports remis : ............ (n°, custodians, signatures)

Incidents / écarts au script : ............

Signatures : ............ ............ ............
```

## 199. Modèle de fiche serveur (template)

```
FICHE SERVEUR — mise en service sécurisée
Nom : ............  S/N châssis : ............  S/N CM : ............
Baie/U : ............  VLAN : ............  IP BMC : ............

Firmware : UEFI ............, BMC ............, NIC ............, disques ............
Secure Boot : activé / clés ............
TPM : présent (dTPM/fTPM), EK vérifié : oui/non
Boot Guard/PSB : ............

BMC : mot de passe changé le ............ (coffre réf. ............),
      services : HTTPS/Redfish OK, Telnet/HTTP/IPMI1.5 désactivés : oui/non
      certificat remplacé : oui/non

OS : image ............ (hash ............), LUKS scellé TPM : oui/non
Comptes : ............

Date de mise en service : ............  Opérateur : ............
Prochaine revue firmware : ............
```

## 200. Modèle de registre firmware (template)

| Serveur | UEFI | BMC | NIC | Disques | FPGA/bitstream | Baseline ? | Écart | Action |
|---|---|---|---|---|---|---|---|---|
| srv-pki-01 | 2.14 | 7.02 | 21.5 | 4.01 | — | ✅ | — | — |
| srv-vault-02 | 2.14 | 6.98 | 21.5 | 4.01 | — | ❌ | BMC 6.98 < 7.02 | Planifier MAJ |
| srv-gpu-01 | 2.14 | 7.02 | 21.5 | 4.01 | bitstream v3.2 signé | ✅ | — | — |

Revu **mensuellement**, écarts traités sous 30 jours (critiques : sous 7 jours).
Source : inventaire Redfish automatisé (section 168).

## 201. Revue annuelle de la fonction crypto

Ordre du jour (1 journée) :

1. Bilan incidents de l'année (tentatives, pannes, alertes).
2. Résultats du **test DR** (section 138) et du test d'intrusion physique.
3. État des clés : inventaire, rotations faites/restantes, expirations à venir.
4. Firmwares : parc à jour ? écarts ?
5. Veille : nouveautés (PQC, failles, EOL produits) — impacts ?
6. Budget N+1 : renouvellements, extensions, formations.
7. Mise à jour des politiques et du modèle de menaces (section 184).

Compte-rendu à la direction : 2 pages, chiffrées (coûts, risques résiduels).

# PARTIE 15 — CAS D'ÉCOLE COMPLETS : DE LA THÉORIE AU DEVIS

## 202. Cas 1 : datacenter régional, ETI 400 personnes (BOM complète)

**Contexte** : ETI industrielle, 1 datacenter 20 baies, PKI interne (2 000 utilisateurs),
TDE sur 3 bases, Vault pour 40 applis, TSA pour les journaux, site de secours à 30 km.

**BOM sécurité/crypto** :

| Poste | Choix | Qté | Coût indicatif |
|---|---|---|---|
| HSM réseau (prod + HA) | Appliance FIPS 140-2 L3, ~5 000 tps | 2 | ~50 000 € |
| HSM backup (site secours) | Même gamme, modèle inférieur | 1 | ~18 000 € |
| YubiHSM 2 (racine offline) | Standard + FIPS | 2 | ~1 600 € |
| Serveurs PKI/Vault/TSA | 2U durcis, dTPM | 3 | ~24 000 € |
| Coffre-fort ignifuge | Clés, PV, backup | 1 | ~2 500 € |
| Contrôle d'accès (3 portes) | Badges DESFire + biométrie salle | 1 lot | ~12 000 € |
| Vidéosurveillance | 8 caméras + NVR 90 j | 1 lot | ~9 000 € |
| Câblage zone crypto + PDU metered | — | 1 lot | ~4 000 € |
| Prestations (install, cérémonie, formation) | — | 1 lot | ~18 000 € |
| Support HSM 5 ans (18 %/an) | — | — | ~61 000 € |
| **Total 5 ans** | | | **~200 000 €** |

Soit **~40 k€/an** — à comparer au coût d'une fuite de données industrielles ou d'une
journée d'arrêt de production (souvent supérieur). Le découpage en phases (année 1 :
HSM + PKI ; année 2 : site secours) lisse le budget.

## 203. Cas 2 : PME industrielle, 60 personnes, sans datacenter

**Contexte** : pas de salle serveur (hébergé chez un prestataire), besoin : PKI interne
(Wi-Fi, VPN), chiffrement des sauvegardes, signature des plans/contrats.

**Solution** : pas d'appliance — **2× YubiHSM 2** (650 € pièce) + EJBCA sur VM hébergée
+ Vault OSS pour les secrets + sauvegardes chiffrées (envelope encryption, DEK
enveloppées par une KEK du YubiHSM). Cérémonie simplifiée (2 officiants + témoin).
**Budget total : ~8-12 k€** (voir le détail section 118). Le point dur n'est pas le
matériel mais la **discipline** : cérémonie, sauvegardes testées, rotations. Quand
l'activité croît (> 100 signatures/heure, audit externe) : migrer vers l'appliance
(section 166).

## 204. Cas 3 : hébergeur / MSP, mutualisation

**Contexte** : hébergeur avec 30 clients, besoin de séparer les clés par client.

**Solution** : 2× appliance réseau en HA avec **partitions par client** (jusqu'à 20 par
boîtier chez Thales, 31 conteneurs chez Utimaco u.trust — sections 20, 24). Chaque
client a son CO et ses politiques ; facturation au client (le HSM devient un **centre
de profit**). Points de vigilance : le **partage du débit** (dimensionner le boîtier
pour la somme des pics), la **réversibilité** (export wrappé des clés si un client
part), et le **contrat** (qui est responsable en cas de compromission d'une partition
voisine ? — à cadrer juridiquement).

## 205. Dimensionnement électrique complet d'une salle crypto

Reprenons le cas 1 (section 202), zone crypto d'une baie :

| Équipement | Puissance max |
|---|---|
| 2× HSM réseau | 2× 110 W = 220 W |
| 3× serveurs 2U (PKI, Vault, TSA) | 3× 350 W = 1 050 W |
| 1× switch management | 60 W |
| 1× KVM/écran (ponctuel) | 50 W |
| **Sous-total** | **1 380 W** |
| Marge ×1,5 | **2 070 W** |

