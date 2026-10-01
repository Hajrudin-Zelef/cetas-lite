---
id: collect-261001-rattrapage/rattrapage/datacenter-securite-trust-guide-20
title: "Sécurité datacenter — HSM, Root of Trust, FPGA, sécurité physique, gestion des clés"
domain: rattrapage
role: reference
task: reference
actors: ["Intel", "Nvidia"]
dates: []
keywords: ["agents", "asic", "benchmark", "intel", "nvidia"]
source: docs/RAG/collect-261001-rattrapage/datacenter_securite_trust_guide.md
source_anchor: ""
source_lines: [2501, 2640]
sha256: b77207757646a408a21648276a22c19983a47c1c5e59ea75e2667808c04816e1
---

# Sécurité datacenter — HSM, Root of Trust, FPGA, sécurité physique, gestion des clés

| Poste | Détail | Budget indicatif |
|---|---|---|
| dTPM 2.0 (5 serveurs critiques) | Modules/options OEM | ~500 € |
| Temps : durcissement BMC (50) | 1 h/serveur | ~2 semaines-homme |
| Temps : Secure Boot + TPM + LUKS | 2 h/serveur | ~3 semaines-homme |
| Keylime (vérificateur + agents) | 1 semaine-homme + infra | ~5 000 € |
| Formation équipe (3 pers.) | 2 jours | ~4 000 € |
| Prestation d'audit initial | 5 jours | ~7 500 € |
| **Total** | | **~25-35 k€ + temps interne** |

**Planning** : semaine 1-2 audit et baseline ; 3-6 durcissement par vagues de 10 ;
7 Keylime ; 8 tests d'intrusion et recettes. Le poste « temps interne » est le vrai coût
— le prévoir dans la charge de l'équipe.

## 175. FAQ RoT (8 questions)

1. **fTPM ou dTPM ?** dTPM sur les serveurs critiques, fTPM acceptable ailleurs **si**
   patché. Le dTPM coûte quelques dizaines d'euros — ne pas lésiner sur les critiques.
2. **Secure Boot casse-t-il le dual-boot / les modules DKMS ?** Les modules non signés
   sont bloqués : signer ses modules (MOK) ou utiliser des noyaux/modules signés.
3. **L'attestation remplace-t-elle l'antivirus ?** Non — elle vérifie l'**intégrité du
   boot**, pas le comportement à l'exécution. Complémentaire.
4. **TPM + cloud public ?** Les vTPM existent (attestation par l'hyperviseur) — modèle
   de confiance différent (on fait confiance au provider, voir partie 9).
5. **Que faire d'un serveur sans TPM ?** Le prévoir au renouvellement ; en attendant,
   durcissement OS + chiffrement avec passphrase au coffre + surveillance renforcée.
6. **BMC compromis = serveur à jeter ?** Non : reflash complet du BMC depuis l'image
   officielle, changement de tous les secrets, vérification des firmwares — mais avec
   une procédure écrite et un contrôle.
7. **SPDM est-il déployé ?** En 2026 : sur les plateformes récentes et les composants
   neufs — **à vérifier** à l'achat (c'est une case du cahier des charges).
8. **Le measured boot ralentit-il le démarrage ?** De quelques secondes (hash des
   composants) — négligeable sur un serveur.

# PARTIE 12 — APPROFONDISSEMENTS FPGA : CHIFFRER LE ROI

## 176. Étude ROI : offload IPsec sur FPGA (calcul complet)

**Contexte** : 40 Gb/s d'IPsec site-à-site à traiter, actuellement sur CPU (8 cœurs
dédiés sur 4 serveurs).

| Poste | CPU seul | Avec FPGA (2× cartes) |
|---|---|---|
| Cœurs CPU mobilisés | 32 cœurs (~4 serveurs partiellement) | ~2 cœurs (contrôle) |
| Cartes FPGA | 0 € | ~19 000 $ (2× V80 MSRP — à vérifier) |
| Serveurs évités / réalloués | — | ~2 serveurs (~16 000 €) |
| Énergie : CPU 32 cœurs (~400 W) vs FPGA (380 W + 30 W contrôle) | ~400 W | ~410 W |
| Latence IPsec | ~50-200 µs (variable) | ~5-20 µs (déterministe) |
| Développement/intégration | 0 € | ~30-60 k€ (presta ou interne) |

**Verdict** : à débit égal, le FPGA ne fait pas forcément économiser l'énergie ici —
il **libère des cœurs** et **garantit la latence**. Le ROI se calcule en **serveurs
évités** et en **SLA réseau**, pas en watts. Si les 32 cœurs sont réalloués à du
revenu (VM clients), le FPGA se paie en **12-18 mois**. Sinon, rester sur CPU.

## 177. SmartNIC vs DPU vs FPGA brut : quoi choisir

| Solution | Exemples | Contrôle | Effort | Idéal pour |
|---|---|---|---|---|
| **FPGA brut** (Alveo V80) | Carte PCIe + votre design | Total (RTL/HLS) | Élevé (mois) | Protocole propriétaire, latence extrême |
| **SmartNIC FPGA** | N6001-PL, cartes ODM | Moyen (firmware + partie programmable) | Moyen | Offload réseau standard (OVS, IPsec) |
| **DPU/ASIC** | NVIDIA BlueField, Intel Mount Evans | Faible (SDK) | Faible (semaines) | Cas standards, volumes |

Règle : **DPU** si votre besoin est standard (c'est du produit), **SmartNIC FPGA** si
vous avez besoin de logique custom modérée, **FPGA brut** si le différenciateur est
dans le silicium que vous dessinez. Ne pas payer le prix du FPGA brut pour un usage
DPU.

## 178. Sécurité du bitstream en profondeur

Le bitstream est un **actif critique** : il contient votre propriété intellectuelle et
définit le comportement matériel.

- **Chiffrement** : les FPGA modernes (Versal, Agilex) chiffrent le bitstream stocké en
  flash (AES), déchiffré à l'intérieur de la puce avec une clé **en eFuses/BBRAM**.
- **Authentification** : signature (RSA/ECDSA) vérifiée avant configuration — un
  bitstream modifié ne charge pas.
- **Chaîne d'approvisionnement** : générer les clés de bitstream en interne, les
  stocker en MofN (même discipline que les clés crypto, partie 5).
- **Mise à jour** : canal sécurisé (TLS mutuel), version vérifiée, rollback possible.
- **Menace** : un attaquant qui remplace le bitstream d'une carte réseau FPGA contrôle
  le trafic — traiter les cartes FPGA comme des **équipements de sécurité**, pas comme
  des périphériques banals.

## 179. Tester un FPGA avant achat : la méthode

1. **Définir le benchmark** : votre workload réel (pas un benchmark synthétique du
   vendeur) — débit, latence p99, conso mesurée à la prise.
2. **POC 30 jours** : les distributeurs/constructeurs prêtent des cartes ; exiger le
   kit de développement et un exemple proche de votre cas.
3. **Mesurer** : latence (hardware timestamping), débit soutenu 24 h (stabilité
   thermique), erreurs.
4. **Chiffrer le TCO** : carte + serveur + développement + formation (section 88).
5. **Décider** sur les chiffres, pas sur la démo — et documenter la décision
   (si le FPGA ne passe pas le POC, c'est une économie, pas un échec).

## 180. FPGA et PQC : préparer la transition

Les algorithmes PQC (ML-KEM, ML-DSA) ont des **clés et signatures plus grosses** et des
profils de calcul différents de RSA/ECC — les accélérateurs actuels devront évoluer :

- Le FPGA est **idéal** pour cette transition : reprogrammable quand les paramètres
  changent (cf. affaire HAWK, section 153), alors qu'un ASIC crypto serait figé.
- Prévoir dans les designs crypto-FPGA une **marge de ressources** (LUT/BRAM) pour les
  futurs algos PQC.
- Les moteurs câblés (ex. 400G crypto du V80) gèrent le **symétrique** (AES) — non
  impacté par le quantique en pratique (AES-256 résiste) ; c'est l'**asymétrique**
  qui migrera.

## 181. Refroidissement : dimensionner l'airflow (calcul)

Une carte passive de 190 W (V80) dans un serveur 2U :

- Air nécessaire (ordre de grandeur) : P = ṁ × Cp × ΔT → pour ΔT = 15 °C,
  ṁ ≈ 190 / (1005 × 15) ≈ 0,0126 kg/s ≈ **~25 CFM** à travers la carte.
- Vérifier que le serveur fournit ce débit **au niveau du slot** (courbes constructeur,
  pas de valeur « moyenne châssis »).
- Température d'entrée d'air ≤ 30 °C recommandée (au-delà, les ventilateurs montent en
  régime = bruit + conso + usure).
- **Ne jamais** obstruer le slot adjacent, ne jamais laisser de câble devant la carte.
- En baie : 4 serveurs × 2 FPGA = **1,5 kW** de plus à évacuer — à intégrer au bilan
  thermique de la baie (section 142).

## 182. Comparatif perf/watt chiffré (tableau indicatif)

Ordres de grandeur pour du **chiffrement AES-GCM à haut débit** (à mesurer sur votre
workload — les chiffres varient fortement) :

| Plateforme | Débit AES-GCM typique | Puissance | Efficacité indicative |
|---|---|---|---|
| CPU (AES-NI, 1 cœur) | ~5-10 Gb/s | ~15 W/cœur | ~0,5 Gb/s/W |
| CPU (multi-cœurs, 16) | ~80-150 Gb/s | ~250 W | ~0,5 Gb/s/W |
| FPGA (moteurs câblés) | ~400 Gb/s (3× 400G sur V80) | ~190 W (carte) | ~2 Gb/s/W |
| ASIC (moteur dédié) | Ligne-rate | ~10-30 W | > 10 Gb/s/W |

Le FPGA gagne **~4×** en efficacité sur le CPU pour le chiffrement en ligne — c'est là
que l'argument énergie est le plus fort. Pour du calcul irrégulier, l'écart se réduit.

## 183. FAQ FPGA (8 questions)

