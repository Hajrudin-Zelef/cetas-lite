---
id: collect-261001-rattrapage/rattrapage/datacenter-securite-trust-guide-25
title: "Sécurité datacenter — HSM, Root of Trust, FPGA, sécurité physique, gestion des clés"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["datacenter", "asic", "gpu", "luna", "open source"]
source: docs/RAG/collect-261001-rattrapage/datacenter_securite_trust_guide.md
source_anchor: ""
source_lines: [3238, 3350]
sha256: 383e0210e3c81ebd82711f94dfac429025cd9f1ebaa3ef430bf7942fa72cd292
---

# Sécurité datacenter — HSM, Root of Trust, FPGA, sécurité physique, gestion des clés

**Traduction** : 2 PDU 16 A (A+B), 1 disjoncteur 10 A par PDU, onduleur : quote-part
3 kVA avec 30 min d'autonomie (~1 kWh utile pour cette zone), clim : ~2,5 kW froid.
**Test** : coupure simulée semestrielle avec chronométrage de la séquence de
redémarrage (section 145). Coût annuel de la continuité électrique de cette zone :
~3 500 € (élec + maintenance onduleur) — dérisoire face au risque.

## 206. Planning PRA crypto (reprise après sinistre)

| Étape | Action | RTO cible | Responsable |
|---|---|---|---|
| 0 | Déclenchement : sinistre site principal | — | Direction |
| 1 | Mise en route site secours (élec, réseau) | 1 h | Infra |
| 2 | Allumage HSM backup, vérification partitions | 30 min | Admin HSM |
| 3 | Restauration KMS/Vault (auto-unseal ou MofN) | 30 min | Admin HSM |
| 4 | Restauration AD CS/EJBCA + CRL/OCSP | 1 h | Admin PKI |
| 5 | Restauration applis (TDE : déchiffrement OK ?) | 2 h | Exploitation |
| 6 | Tests : signature, déchiffrement, certificats | 30 min | RSSI |
| 7 | Communication clients | — | Direction |
| **Total** | | **~5 h 30** | |

Testé **une fois par an** en conditions réelles (section 138) — le RTO réel est
toujours supérieur au RTO papier la première fois.

## 207. Leçons apprises : ce que les incidents enseignent

Synthèse des retours d'expérience de l'industrie (pas d'invention : des patterns) :

1. **La panne arrive toujours pendant les congés** — d'où l'astreinte formée, les
   runbooks écrits pour un lecteur fatigué, et les mots de passe au coffre (pas « dans
   la tête de Karim »).
2. **C'est l'électrique qui tue la crypto** — 80 % des indisponibilités de HSM en PME
   viennent de coupures, pas d'attaques. La partie 8 n'est pas un bonus.
3. **La compromission vient de l'intérieur** — insider ou prestataire : d'où la
   séparation des rôles, les journaux, le MofN.
4. **Le firmware oublié** — BMC non patché depuis 3 ans : la faille était publique.
   Le registre (section 200) et la veille (section 197) sont des mesures de sécurité
   à part entière.
5. **La sauvegarde non testée** — découverte le jour du sinistre. Tester, tester,
   tester.
6. **Le projet crypto sous-budgété en temps humain** — le matériel est 30 % du coût,
   l'humain 70 % (installation, cérémonies, tests, formation, revues). Le dire à la
   direction **avant**, pas après.

# ANNEXES

## A. Index des sections

**Partie 1 — HSM (1-48)** : 1 principe, 2 TRNG, 3 tamper resistance, 4 tamper evidence,
5 FIPS 140, 6 Common Criteria, 7 PCI HSM/eIDAS, 8 PKI, 9 TDE, 10 signature, 11 horodatage,
12 DNSSEC, 13 paiement, 14 5G/blockchain/IoT, 15 form factors, 16 appliance réseau,
17 carte PCIe, 18 USB/nano, 19 cloud HSM, 20 Thales Luna Network, 21 Thales Luna PCIe,
22 Luna USB/DPoD, 23 Utimaco Se Gen2, 24 Utimaco u.trust, 25 Atalla, 26 Entrust nShield,
27 Futurex, 28 Marvell LiquidSecurity, 29 YubiHSM 2, 30 comparatif général,
31 performances, 32 certifications, 33 schéma datacenter, 34 schéma PME, 35 VLAN/firewall,
36 dimensionnement méthode, 37 exemples chiffrés, 38 partitions, 39 PKCS#11, 40 autres APIs,
41 sauvegardes, 42 cérémonie, 43 MofN, 44 custodians, 45 HA, 46 batterie/MTBF, 47 prix,
48 TCO.
**Partie 2 — RoT (49-70)** : 49 définition, 50 chaîne de confiance, 51 lien guide Zelef,
52 TPM serveur, 53 PCR, 54 secure boot serveurs, 55 measured boot, 56 attestation protocole,
57 vérificateur, 58 Boot Guard/PSB, 59 BMC talon d'Achille, 60 durcissement BMC,
61 OpenBMC/Redfish/SPDM, 62 firmware signés, 63 contrefaçons, 64 SBOM, 65 réception,
66 reconditionné, 67 cas BMC compromis, 68 tableau RoT, 69 limites TPM, 70 TPM+HSM.
**Partie 3 — FPGA (71-90)** : 71 principe, 72 vocabulaire, 73 flot dev, 74 HLS, 75 Alveo V80,
76 autres Alveo, 77 Agilex 7, 78 SmartNIC N6000, 79 autres acteurs, 80 FPGA vs GPU vs ASIC,
81 réseau, 82 crypto, 83 trading, 84 IA, 85 stockage, 86 conso/refroidissement,
87 grille de décision, 88 coûts, 89 outils, 90 production.
**Partie 4 — Physique (91-104)** : 91 4 cercles, 92 périmètre, 93 contrôle d'accès,
94 cages/racks, 95 vidéo, 96 intrusion, 97 visiteurs, 98 anti-tailgating, 99 personnel,
100 NIST 800-88, 101 SED, 102 déclassement, 103 élec/incendie, 104 budgets.
**Partie 5 — KMS (105-120)** : 105 cycle de vie, 106 hiérarchie, 107 envelope encryption,
108 rotation, 109 durées de vie, 110 rôles, 111 HSM↔KMS, 112 Vault, 113 auto-unseal,
114 EJBCA, 115 sauvegardes chiffrées, 116 BYOK/HYOK, 117 comparatif KMS, 118 BOM PME,
119 erreurs, 120 journaux.
**Partie 6 — Conformité (121-128)** : 121 ISO 27001, 122 contrôles crypto, 123 PCI DSS,
124 RGPD, 125 eIDAS, 126 questionnaires clients, 127 journaux/audit, 128 politiques.
**Partie 7 — Checklists (129-138)** : 129 serveur, 130 BMC, 131 firmware/secure boot,
132 HSM, 133 cérémonie, 134 FPGA, 135 baie, 136 runbook HSM down, 137 runbook clé
compromise, 138 test DR.
**Partie 8 — Énergie (139-146)** : 139 consos, 140 règle d'or onduleur, 141 architecture
élec, 142 refroidissement, 143 dimensionnement, 144 monitoring, 145 séquence groupe,
146 cas baie crypto.
**Partie 9 — À venir (147-157)** : 147 NIST PQC, 148 FIPS 206, 149 CNSA 2.0, 150 EO/M-26-15,
151 TLS hybride, 152 PQC dans HSM, 153 HAWK, 154 TDX/SEV-SNP, 155 DDRop, 156 H100/CCA/CoCo,
157 pas encore là.
**Partie 10 — Approfondissements HSM (158-167)** : 158 AD CS, 159 OpenSSL/nginx,
160 EJBCA, 161 OpenTSA, 162 DNSSEC/BIND, 163 Luna HA, 164 négociation, 165 TCO 3 scénarios,
166 migration PKI, 167 FAQ HSM.
**Partie 11 — Approfondissements RoT (168-175)** : 168 Redfish, 169 Keylime, 170 LUKS+TPM,
171 GRUB, 172 parc BMC, 173 microcodes, 174 cas 50 serveurs, 175 FAQ RoT.
**Partie 12 — Approfondissements FPGA (176-183)** : 176 ROI IPsec, 177 SmartNIC/DPU,
178 bitstream, 179 POC, 180 FPGA+PQC, 181 airflow, 182 perf/watt, 183 FAQ FPGA.
**Partie 13 — Approfondissements KMS/physique (184-191)** : 184 modèle de menaces,
185 Vault durci, 186 rotation auto, 187 3-2-1 clés, 188 câblage, 189 SIEM, 190 exercice
de crise, 191 FAQ KMS.
**Partie 14 — Achat/gouvernance (192-201)** : 192 RFP, 193 grille notation, 194 planning,
195 RACI, 196 formation, 197 veille, 198 PV cérémonie, 199 fiche serveur, 200 registre
firmware, 201 revue annuelle.
**Partie 15 — Cas d'école (202-207)** : 202 ETI, 203 PME, 204 hébergeur, 205 dimensionnement
élec, 206 PRA, 207 leçons.
**Partie 16 — Sauvegardes/réseau/PRA (208-215)** : 208 chiffrement sauvegardes, 209
ransomware, 210 segmentation, 211 bastion, 212 supervision corrélée, 213 certificats à
grande échelle, 214 site secours, 215 classeur d'exploitation.
**Partie 17 — Feuille de route (216-223)** : 216 2026-2027, 217 2028-2029, 218 2030+,
219 maturité, 220 KPIs, 221 budget, 222 argumentaire, 223 lundi matin.
**Partie 18 — Attaques (224-233)** : 224 HSM, 225 rootkits UEFI, 226 BMC, 227 réseau,
228 physique, 229 supply chain, 230 HNDL, 231 sauvegardes, 232 ingénierie sociale,
233 synthèse.
**Partie 19 — Pour aller plus loin (234-240)** : 234 labo, 235 outils open source,
236 livres et références, 237 formations et certifications, 238 communautés, 239 maquette
POC, 240 mot de la fin.

## B. Table des tableaux et schémas

