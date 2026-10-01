---
id: collect-261001-rattrapage/rattrapage/datacenter-securite-trust-guide-17
title: "Sécurité datacenter — HSM, Root of Trust, FPGA, sécurité physique, gestion des clés"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "AWS", "Anthropic", "CISA", "Intel", "United States"]
dates: ["2026-09-27", "2027-12-31", "2030-12-31", "2031-12-31"]
keywords: ["datacenter", "amd", "arr", "aws", "compute", "distribution", "intel", "luna", "valuation"]
source: docs/RAG/collect-261001-rattrapage/datacenter_securite_trust_guide.md
source_anchor: ""
source_lines: [2071, 2210]
sha256: 7ec789c3af64cd8731cddf61c7b01a0423987f4e5cbda32baec1ed7d5ceb7a34
---

# Sécurité datacenter — HSM, Root of Trust, FPGA, sécurité physique, gestion des clés

Temps typique : **30-60 min** pour une zone crypto complète — à intégrer au RTO du PRA.

## 146. Cas chiffré : baie crypto complète (élec + froid)

**Contenu** : 2× HSM réseau, 2× serveurs (Vault + PKI), 1 switch, 1 FPGA de test.

| Poste | Valeur |
|---|---|
| Puissance max installée | ~1 600 W (avec marge ×1,5) |
| Onduleur dédié ou quote-part | 3 kVA N+1 (partagé avec la baie) |
| Autonomie batteries | 30 min |
| Clim (quote-part) | ~2 kW froid |
| Coût élec annuel (0,20 €/kWh) | ~2 800 € |
| Coût complet annuel (élec + clim + maintenance onduleur) | ~5 000-7 000 € |

À comparer au **coût d'une heure d'arrêt** des services de clés (signatures bloquées,
bases inaccessibles, VPN down) : pour la plupart des organisations, **quelques heures
d'arrêt coûtent plus cher que l'année d'électricité secourue**. C'est l'argument à
présenter à la direction.

# PARTIE 9 — À VENIR : ANNONCES VÉRIFIÉES AU 27/09/2026

> Règle de cette partie : uniquement des faits **sourcés et vérifiés le 27/09/2026**.
> Le reste est marqué « à vérifier » ou « non trouvé au 27/09/2026 ».

## 147. PQC : l'état des standards NIST au 27/09/2026 (vérifié)

Le NIST a finalisé ses trois premiers standards post-quantiques le **13 août 2024**
(vérifié) :

| Standard | Algorithme | Usage | Statut au 27/09/2026 |
|---|---|---|---|
| **FIPS 203** | **ML-KEM** (ex-Kyber) | Échange de clés / encapsulation | Finalisé, prêt pour la production |
| **FIPS 204** | **ML-DSA** (ex-Dilithium) | Signatures usage général | Finalisé, prêt pour la production |
| **FIPS 205** | **SLH-DSA** (ex-Sphincs+) | Signatures à base de hash | Finalisé, prêt pour la production |

Conséquence pratique : **on peut déjà déployer** ML-KEM/ML-DSA dans les nouveaux
systèmes (avec crypto-agilité, voir section 152). L'argument « attendons les standards »
n'est plus valable depuis août 2024. Menace active : **harvest now, decrypt later** —
les adversaires enregistrent le trafic chiffré aujourd'hui pour le déchiffrer quand un
ordinateur quantique cryptographiquement pertinent existera (fenêtre crédible estimée
**2029-2032** par les acteurs industrie, à prendre comme ordre de grandeur).

## 148. FIPS 206 (HQC) : en attente (à vérifier)

Le NIST a sélectionné **HQC** comme KEM de secours (basé sur les codes correcteurs,
famille différente des réseaux euclidiens — diversification bienvenue). Le standard
**FIPS 206** est **en cours de finalisation, attendue 2026-2027** (suivi communautaire
pqca, septembre 2026) — **à vérifier** sur le site du NIST. En pratique : ne pas
attendre HQC pour démarrer ; ML-KEM est le standard de production aujourd'hui.

## 149. CNSA 2.0 : le calendrier qui s'impose (vérifié)

La **NSA** impose via **CNSA 2.0** (Commercial National Security Algorithm Suite 2.0,
v2.1 de décembre 2024) la migration PQC des systèmes de sécurité nationale US — avec des
effets de bord sur tous les fournisseurs (vérifié, septembre 2026) :

- **1er janvier 2027** : les **nouvelles acquisitions** de systèmes de sécurité nationale
  doivent supporter CNSA 2.0 par défaut (CNSSP 15).
- **31 décembre 2030** : les équipements ne supportant pas CNSA 2.0 doivent être
  **retirés**.
- **31 décembre 2031** : usage **exclusif** des algorithmes CNSA 2.0.
- Algorithmes retenus : **ML-KEM, ML-DSA** (SLH-DSA **non approuvé** pour ces systèmes),
  AES-256, SHA-384/512.

Même hors marché US : les équipementiers alignent leurs roadmaps sur CNSA 2.0 — c'est le
calendrier qui tire l'industrie.

## 150. EO 14412 et OMB M-26-15 : les jalons US (vérifié)

- **Executive Order 14412** (22 juin 2026, « Securing the Nation Against Advanced
  Cryptographic Attacks ») : les agences fédérales doivent migrer leurs systèmes les plus
  sensibles vers PQC au **31/12/2030**, et vers l'authentification post-quantique au
  **31/12/2031** ; pilote de migration terminé au **31/12/2027**.
- **OMB M-26-15** : plan de migration en **5 phases de 2026 à 2035**, plans de migration
  des agences à remettre en **octobre 2026**.
- **NIST IR 8547** : calendrier de retrait des algos classiques — RSA/ECDSA/EdDSA et
  Diffie-Hellman à 112 bits de sécurité **dépréciés après 2030, interdits en 2035**.

Pour un datacenter européen : pas d'obligation directe, mais les **exigences clients**
et les **fournisseurs** suivront ces calendriers — anticiper 2027-2028 comme fenêtre
d'alignement (recommandation ENISA en ce sens).

## 151. TLS hybride X25519MLKEM768 : déjà en production (vérifié)

Le mode **hybride** combine l'échange classique (X25519) et post-quantique (ML-KEM-768) :
si l'un des deux est cassé, l'autre protège encore. Vérifié septembre 2026 : supporté en
production par **Chrome, Firefox, Cloudflare, AWS, Windows Schannel**. Concrètement : une
partie du trafic TLS mondial est **déjà** résistante au harvest-now-decrypt-later côté
échange de clés. Reste classique : **l'authentification** (certificats) — la migration des
signatures (ML-DSA) est le chantier 2027-2030. Action : vérifier que vos terminaisons TLS
négocient bien les groupes hybrides, et planifier la PKI PQC.

## 152. PQC dans les HSM : Quantum Protect, Luna PQC FM (vérifié)

Les HSM suivent (vérifié le 27/09/2026, suivi pqca/wg-readiness-tracking) :

- **Utimaco u.trust GP HSM** : package **Quantum Protect**, statut **« Ready »** —
  supporte **ML-KEM, ML-DSA, XMSS, LMS, HSS**, avec un **simulateur gratuit** pour tester
  la migration. Argument fort : support PQC **sans changement de matériel**
  (crypto-agilité matérielle).
- **Thales Luna** : **PQC Functionality Module** documenté (mécanismes post-quantiques
  dans le HSM).
- **YubiHSM 2** (firmware 2.4+) : PQC **en test pilote**, support complet non prévu —
  pour une racine PQC future, prévoir un autre matériel.

Leçon : à l'achat d'un HSM en 2026, exiger le **support PQC logiciel** (au minimum par
mise à jour firmware) — un HSM incapable de ML-KEM/ML-DSA sera obsolète avant sa fin de
vie mécanique.

## 153. L'affaire HAWK : juillet 2026 (vérifié)

En **juillet 2026**, une équipe incluant un modèle d'IA d'Anthropic a trouvé une
**faiblesse réelle** dans **HAWK**, un schéma de signature à base de réseaux encore en
évaluation au NIST. L'équipe HAWK l'a **retiré** de la compétition. Aucun système en
production ne l'utilisait — rien n'a cassé. La leçon pour ce guide : **« post-quantique »
ne veut pas dire « définitivement sûr »**. D'où l'exigence de **crypto-agilité** : la
capacité à changer d'algorithme sans ré-architecturer (c'est le critère n°1 d'un achat
HSM/KMS en 2026, devant le choix de l'algorithme lui-même).

## 154. Confidential computing : Intel TDX, AMD SEV-SNP (vérifié)

Le **confidential computing** chiffre la **mémoire vive** des VM : ni l'hyperviseur, ni
l'opérateur cloud, ni un accès physique à la RAM ne voient les données en clair
(vérifié, documentation 2026) :

- **Intel TDX** (Trust Domain Extensions) : VM confidentielles sur Xeon Scalable
  (4e gén. et suivantes).
- **AMD SEV-SNP** (Secure Encrypted Virtualization – Secure Nested Paging) : sur EPYC
  (Milan/Genoa et suivants), avec protection d'intégrité et attestation.
- **Arm CCA** (Confidential Compute Architecture, Armv9) : Realm isolés gérés par le
  RMM — edge et serveurs Arm (vérifié).
- Usage : **attestation à distance** (section 56) + distribution de secrets (le secret
  n'est délivré qu'à une VM attestée — projet **Trustee**/Confidential Containers).

Pour un datacenter : TDX/SEV-SNP protègent les workloads sensibles même contre
l'administrateur de l'hyperviseur — complément du chiffrement disque et du HSM.

## 155. DDRop : l'attaque de septembre 2026 (vérifié)

