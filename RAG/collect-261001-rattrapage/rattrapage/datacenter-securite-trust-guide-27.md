---
id: collect-261001-rattrapage/rattrapage/datacenter-securite-trust-guide-27
title: "Sécurité datacenter — HSM, Root of Trust, FPGA, sécurité physique, gestion des clés"
domain: rattrapage
role: reference
task: reference
actors: ["United States"]
dates: ["2027-01-01"]
keywords: ["arr", "incident", "valuation"]
source: docs/RAG/collect-261001-rattrapage/datacenter_securite_trust_guide.md
source_anchor: ""
source_lines: [3485, 3630]
sha256: 73d048433ca44ee931f8186ad37880a43c3186ef476ebb86df57c110090ab5b9
---

# Sécurité datacenter — HSM, Root of Trust, FPGA, sécurité physique, gestion des clés

| Corrélation | Signification probable | Action |
|---|---|---|
| Échecs auth HSM + connexion BMC inhabituelle | Attaque en cours | Isoler, incident P1 |
| Pic d'opérations HSM ×10 + process inconnu | Ransomware utilisant le HSM | Couper l'accès client, incident P1 |
| Dérive horloge > 1 s + TSA active | Jetons contestables | Resync NTS, vérifier les jetons émis |
| Température baie > 32 °C + charge onduleur > 90 % | Risque coupure thermique | Délestage, intervention |
| Alerte intrusion physique + coupure réseau | Attaque coordonnée | Cellule de crise |

Ces règles se construisent avec le temps — commencer par les 5 ci-dessus, affiner
trimestriellement.

## 213. Gestion des certificats à grande échelle

Au-delà de ~100 certificats, la gestion manuelle ne passe plus :

- **Inventaire automatisé** : scanner le parc (ports 443, PKI, magasins), détecter les
  expirations à J-30/J-7.
- **ACME interne** : EJBCA ou Vault PKI en mode ACME pour le renouvellement
  automatique (90 jours, section 108).
- **Chaînes courtes** : intermédiaire dédiée par usage (serveurs, Wi-Fi, VPN) pour
  limiter l'impact d'une révocation.
- **CT logs** : publier les certificats publics en Certificate Transparency
  (détection d'émission frauduleuse).
- **Urgence** : procédure de **révocation massive** testée (compromission
  d'intermédiaire) — CRL/OCSP à jour en < 1 h.

## 214. PRA étendu : le site de secours crypto

Le site de secours n'est pas une simple copie des serveurs :

- **HSM du site distant** : synchronisé (réplication) ou restauré depuis la sauvegarde
  chiffrée — **testé** (section 138).
- **Clés DNS** : la KSK DNSSEC doit exister sur les deux sites (signer depuis le
  secours en cas de bascule).
- **Bastion de secours** : administrer le site distant sans dépendre du site principal.
- **Exercice** : bascule **réelle** une fois par an (pas juste « on a lu la procédure »).
- **Coûts** : le site secours crypto (HSM + 1 serveur) ≈ 25-30 k€ — à comparer au
  coût d'une indisponibilité totale (section 146).

## 215. Documentation d'exploitation : le classeur crypto

Le « classeur » (numérique, chiffré, sauvegardé) que tout nouvel admin doit pouvoir
utiliser :

1. Schéma réseau de la zone crypto (à jour).
2. Inventaire : HSM (S/N, firmware, partitions), clés (usage, dates, custodians).
3. Procédures : démarrage/arrêt, bascule HA, restauration, rotation.
4. Runbooks incidents (sections 136-137, 190).
5. PV : cérémonies, tests DR, revues.
6. Contacts : support constructeur (n° contrat), custodians, astreinte.
7. Journal des changements (qui a modifié quoi, quand).

**Règle** : si l'admin principal est injoignable, un second doit pouvoir tout faire
avec ce classeur + le coffre. Le tester : « exercice du bus » annuel.

# PARTIE 17 — FEUILLE DE ROUTE 2026-2030 ET SYNTHÈSE DÉCISIONNELLE

## 216. Feuille de route : 2026-2027 (fondations)

| Trimestre | Actions | Budget indicatif |
|---|---|---|
| T4 2026 | Inventaire clés + firmwares ; durcissement BMC ; Secure Boot/TPM sur critiques | Temps interne |
| T1 2027 | RFP HSM, POC, achat ; cahier des charges PQC (exiger ML-KEM/ML-DSA logiciels) | 50-80 k€ |
| T2 2027 | Installation, cérémonie racine, intégration AD CS/EJBCA + Vault auto-unseal | Prestations ~15 k€ |
| T3 2027 | VLAN crypto + bastion + SIEM ; premier test DR ; TLS hybride activé où dispo | Temps interne |
| T4 2027 | Revue annuelle ; exercice de crise ; formation équipe | ~5 k€ |

Jalon externe : **01/01/2027** — CNSA 2.0 pour les nouvelles acquisitions US (section
149) : vos fournisseurs s'alignent, profitez-en pour exiger le support PQC.

## 217. Feuille de route : 2028-2029 (montée en puissance)

- **PQC** : pilotes ML-KEM pour l'échange (VPN, TLS internes), évaluation ML-DSA pour
  les signatures ; suivre FIPS 206 (HQC).
- **Attestation** : déployer Keylime sur le parc critique, brancher à l'orchestrateur.
- **Site secours** : HSM distant opérationnel, bascule testée.
- **FPGA** : si le besoin réseau est avéré, POC SmartNIC/offload (partie 3).
- **Conformité** : viser ISO 27001 si exigée par les clients (12-18 mois de projet).
- **Budget** : ~30-50 k€/an (support, extensions, prestations).

## 218. Feuille de route : 2030+ (post-quantique opérationnel)

- **2030** : dépréciation des algos classiques à 112 bits (NIST IR 8547) — vos
  systèmes doivent être **hybrides ou PQC** avant cette date.
- **2030-2031** : migration des signatures vers ML-DSA (le chantier le plus long :
  PKI, code signing, documents).
- **2031** : CNSA 2.0 exclusif côté US — vérifier la conformité de vos équipements.
- **2035** : interdiction des algos vulnérables (IR 8547) — horizon de fin de
  migration.
- **Principe** : ne jamais faire de « big bang » crypto — des **vagues** par usage,
  avec rollback possible, sur 4-5 ans.

## 219. Matrice de maturité : où en êtes-vous ?

| Niveau | Clés | Serveurs | Physique | Énergie |
|---|---|---|---|---|
| 1 — Initial | Clés en fichiers | BMC par défaut | Porte fermée à clé | 1 onduleur |
| 2 — Géré | KMS logiciel, inventaire | Secure Boot/TPM | Badges, vidéo | Onduleur testé |
| 3 — Défini | **HSM**, cérémonies, MofN | Attestation, Redfish | Cages, visiteurs, 800-88 | A+B secourus, PRA testé |
| 4 — Maîtrisé | PQC/hybride, rotation auto | Keylime, SPDM | Tests d'intrusion | Corrélation SIEM |

Objectif réaliste pour une ETI : **niveau 3 en 18 mois**, niveau 4 en 3 ans. Se
positionner honnêtement avant de budgéter.

## 220. KPIs : piloter la fonction crypto

- **% de clés critiques en HSM** (objectif : 100 %).
- **% du parc avec Secure Boot + TPM actifs** (objectif : 100 % serveurs).
- **Délai moyen de rotation** des DEK (objectif : < politique).
- **RTO réel du test DR** vs objectif (écart < 20 %).
- **% de firmwares à jour** (objectif : > 95 %, critiques < 7 jours).
- **Nombre d'alertes HSM non traitées** (objectif : 0 > 24 h).
- **Temps moyen de détection** d'une anomalie physique (test d'intrusion annuel).

Revue **trimestrielle** avec la direction : 1 page, tendances, écarts, décisions.

## 221. Budget pluriannuel type (ETI, ordre de grandeur)

| Année | Investissement | Fonctionnement | Total |
|---|---|---|---|
| N (fondations) | 80-120 k€ (HSM, travaux, presta) | 20 k€ | ~100-140 k€ |
| N+1 | 20-30 k€ (site secours, extensions) | 25 k€ (support) | ~45-55 k€ |
| N+2 | 10-20 k€ | 25 k€ | ~35-45 k€ |
| N+3 | Renouvellement partiel | 25 k€ | ~35-45 k€ |

Soit **~250 k€ sur 4 ans** pour une ETI — dont ~40 % de temps humain interne valorisé.
À mettre en regard : le coût moyen d'une violation de données (plusieurs millions
selon les études IBM — ordre de grandeur public) et d'une journée d'arrêt de production.

## 222. Argumentaire direction : 5 phrases qui font signer

1. « Nos clés critiques sont aujourd'hui dans des fichiers : quiconque lit un disque
   les possède. »
2. « Un HSM coûte 25 k€ ; une journée d'arrêt de la PKI coûte [X] k€ — faites le ratio. »
3. « 80 % des pannes crypto viennent de l'électrique, pas des hackers : le plan couvre
   les deux. »
4. « Le quantique ne cassera pas nos systèmes en 2027, mais nos clients nous demanderont
   notre plan PQC dès 2027 — ayons-le. »
5. « Sans documents (politiques, PV, tests), le prochain audit client échoue quel que
   soit le matériel acheté. »

Adapter [X] avec vos chiffres (section 146) — un argument chiffré bat un argument
technique devant une direction.

## 223. Que faire lundi matin : les 5 premières actions

