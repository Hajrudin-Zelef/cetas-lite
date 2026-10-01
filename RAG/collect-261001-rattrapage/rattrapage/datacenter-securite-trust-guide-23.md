---
id: collect-261001-rattrapage/rattrapage/datacenter-securite-trust-guide-23
title: "Sécurité datacenter — HSM, Root of Trust, FPGA, sécurité physique, gestion des clés"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-09-27"]
keywords: ["arr", "gpu"]
source: docs/RAG/collect-261001-rattrapage/datacenter_securite_trust_guide.md
source_anchor: ""
source_lines: [2923, 3064]
sha256: c8c993440b7f20a0ec7d72427d6f523856a7fff3c62a5c0737b2f9cefd7b8ac7
---

# Sécurité datacenter — HSM, Root of Trust, FPGA, sécurité physique, gestion des clés

**Q10. Un HSM protège-t-il contre la perte des clés (panne, incendie) ?**
R : Non — il protège contre le **vol et la copie**, pas contre la **perte**. D'où la
sauvegarde chiffrée (HSM backup, export wrappé en MofN), géographiquement séparée et
testée. Sans sauvegarde, une panne = clés perdues = données chiffrées irrécupérables.

# PIÈGES TERRAIN : 25 ERREURS À NE PAS COMMETTRE

1. **Générer la racine d'AC « vite fait » sur un poste admin** — sans cérémonie, sans
   témoin, clé exportable. Refaire une cérémonie après coup ne répare pas les copies.
2. **Importer une clé logicielle dans le HSM** « pour aller plus vite » — la clé a
   existé en clair, elle est potentiellement compromise dès le départ.
3. **Un seul HSM en production** — pas de HA : la première panne arrête les signatures.
4. **HSM sur le VLAN de production** — accessible à tout le réseau interne au lieu
   d'un VLAN crypto filtré.
5. **PIN SO utilisé au quotidien** — le compte d'initialisation devient le compte
   d'exploitation ; le compromettre = tout compromettre.
6. **Batterie du HSM ignorée** — elle s'épuise, le HSM s'efface un jour sans prévenir.
   Superviser et planifier le remplacement.
7. **Automatiser l'admin HSM sans revue** — un playbook Ansible avec les PIN en clair
   dans git annule tout le bénéfice du HSM.
8. **Secure Boot désactivé « pour dépanner » et jamais réactivé** — l'audit trimestriel
   (section 131) existe pour ça.
9. **BMC en DHCP sur le VLAN de production, mot de passe par défaut** — le scénario
   de la section 67, vu et revu.
10. **IPMI 1.5 laissé activé** — hashes craquables offline ; passer en 2.0 chiffré ou
    désactiver.
11. **Firmware téléchargé hors site constructeur** — « trouvé sur un forum » =
    porte dérobée potentielle.
12. **Mise à jour firmware sans re-mesure** — les PCR changent, toute l'attestation
    passe en alerte, et on finit par ignorer les alertes (pire).
13. **Acheter du FPGA pour un workload qui change tous les mois** — le design ne sera
    jamais fini ; c'était un job pour GPU/CPU.
14. **Carte FPGA passive sans vérifier l'airflow** — surchauffe, throttling, panne en
    pleine charge estivale.
15. **Bitstream non signé** — n'importe qui peut reprogrammer la carte réseau.
16. **Photos en salle serveur** — un tableau blanc, un écran, une étiquette : fuite
    d'information. Interdire et faire respecter.
17. **Badge prêté « pour 5 minutes »** — l'anti-passback et la culture (section 98)
    existent pour ça.
18. **Disques « effacés » par simple formatage** — formatage ≠ effacement ; appliquer
    NIST 800-88.
19. **SED acheté mais chiffrement jamais activé** — un SED non activé est un disque
    clair. Vérifier à la mise en service.
20. **Clé codée en dur dans le code** — elle finira dans git, dans une image Docker,
    dans une sauvegarde. Vault/secret manager, toujours.
21. **Vault en Shamir manuel sur infra à redémarrage automatique** — après chaque
    coupure, quelqu'un doit se lever à 3 h du matin. Auto-unseal (section 113).
22. **Rotation « quand on aura le temps »** — c'est jamais. Automatiser (section 186).
23. **Sauvegarde des clés en clair « temporairement »** — le temporaire dure des années
    et finit dans le cloud.
24. **HSM et serveurs critiques sur onduleur non secouru** — la règle d'or (section
    140) : une coupure = PKI à l'arrêt + Vault scellé.
25. **Acheter sur promesse PQC/quantique** — exiger des standards finalisés (FIPS
    203/204/205), des validations datées, des références déployées. Le reste est du
    marketing (section 157).

# CONCLUSION : LA CHECKLIST DU CHEF DE SERVICE

Si vous ne retenez que dix actions :

1. **Inventorier** les clés et les firmwares (on ne sécurise pas l'inconnu).
2. **Mettre les clés critiques en HSM** — même un YubiHSM 2 à 650 € change tout pour
   une PME.
3. **Cérémonie + MofN** pour chaque clé racine, PV archivé.
4. **Sauvegarder et tester** la restauration chaque année.
5. **Durcir les BMC** : mots de passe uniques, VLAN dédié, firmware à jour.
6. **Activer Secure Boot + TPM** sur 100 % du parc serveurs, auditer trimestriellement.
7. **Segmenter** : VLAN crypto pour les HSM, jamais d'Internet, TLS mutuel.
8. **Secourir électriquement** toute la chaîne crypto (HSM, Vault, PKI) et tester le
   redémarrage après coupure.
9. **Préparer le PQC** : TLS hybride maintenant, crypto-agilité exigée à chaque achat.
10. **Écrire les politiques** : sans documents, l'audit échoue quel que soit le matériel.

> *Guide rédigé le 27/09/2026. Références produits vérifiées par recherche web le
> 27/09/2026 ; prix et détails non sourcés marqués « à vérifier ». Relire avant tout
> achat : les gammes, les firmwares et les validations évoluent vite.*

# PARTIE 14 — DOSSIERS D'ACHAT, GOUVERNANCE ET MODÈLES DE DOCUMENTS

## 192. Cahier des charges HSM : template

Un RFP HSM sérieux contient :

1. **Contexte** : organisation, usages (PKI, TDE, TSA...), volumes (ops/s moyens et
   pics — voir section 36).
2. **Exigences fonctionnelles** : form factor, nombre de partitions/tenants, APIs
   (PKCS#11 obligatoire, JCE/CNG selon écosystème), clustering HA, sauvegarde.
3. **Exigences de sécurité** : FIPS 140-2/3 niveau 3 minimum (référence CMVP exigée),
   Common Criteria si marchés publics, support **PQC logiciel** (ML-KEM/ML-DSA).
4. **Exigences d'exploitation** : SNMP/syslog, API REST d'admin, MTBF, batterie
   remplaçable, support J+1 ou 4 h.
5. **Exigences électriques** : conso max, plage de température, double alimentation.
6. **Prestations** : installation, cérémonie, formation, POC 30 jours.
7. **Prix** : matériel, support annuel (5 ans), prestations — **détaillés par ligne**.
8. **Références** : 3 clients comparables déployés.

Envoyer le **même** RFP aux 3 finalistes (Thales, Utimaco, Entrust) pour comparer à
périmètre égal.

## 193. Grille de notation fournisseurs (exemple)

| Critère | Poids | Thales | Utimaco | Entrust |
|---|---|---|---|---|
| Conformité au besoin fonctionnel | 25 % | /20 | /20 | /20 |
| Certifications (FIPS, CC, PCI) | 15 % | | | |
| PQC / crypto-agilité | 10 % | | | |
| Performance mesurée en POC | 15 % | | | |
| TCO 5 ans | 15 % | | | |
| Support et SLA | 10 % | | | |
| Références clients | 10 % | | | |
| **Total pondéré** | 100 % | | | |

Noter **pendant** le POC, pas après ; faire noter par 2 personnes indépendantes ;
archiver la grille (pièce d'audit d'achat).

## 194. Planning projet PKI/HSM type (12 semaines)

| Semaines | Phase | Livrable |
|---|---|---|
| 1-2 | Cadrage : modèle de menaces, RFP | Dossier de cadrage |
| 3-5 | POC : 2 finalistes, tests d'intégration | Grille de notation, choix |
| 6-7 | Commande, réception, installation physique | PV de réception |
| 8 | Configuration : partitions, réseau, supervision | Dossier d'exploitation |
| 9 | Cérémonie de génération des clés | PV de cérémonie |
| 10 | Intégration applicative (AD CS/EJBCA, Vault) | Recette fonctionnelle |
| 11 | Formation équipe + runbooks | Attestations, runbooks |
| 12 | Test DR + bascule en production | PV de test, GO |

Prévoir **+4 semaines** de marge (délais fournisseurs, indisponibilités) — un projet
crypto ne se fait jamais « entre deux urgences ».

## 195. Gouvernance : RACI de la fonction crypto

| Activité | RSSI | Admin HSM | Exploitation | Direction | Auditeur |
|---|---|---|---|---|---|
| Politique crypto | A | C | C | R | I |
| Cérémonie | A | R | I | C | Témoin |
| Exploitation quotidienne | I | R | C | I | I |
| Rotation planifiée | A | R | C | I | Vérifie |
| Test DR annuel | A | R | R | I | Vérifie |
| Achat/renouvellement | C | C | I | A/R | I |

