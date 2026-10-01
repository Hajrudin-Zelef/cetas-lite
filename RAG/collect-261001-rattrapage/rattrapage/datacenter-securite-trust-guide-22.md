---
id: collect-261001-rattrapage/rattrapage/datacenter-securite-trust-guide-22
title: "Sécurité datacenter — HSM, Root of Trust, FPGA, sécurité physique, gestion des clés"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "AWS", "CISA", "Intel", "Nvidia", "United States"]
dates: []
keywords: ["amd", "asic", "aws", "dsp", "gpu", "intel", "nvidia"]
source: docs/RAG/collect-261001-rattrapage/datacenter_securite_trust_guide.md
source_anchor: ""
source_lines: [2797, 2922]
sha256: 361079103f3ec10589f7cff59e9cbbe37617d75305ca5d8554f0181ef9de0d2c
---

# GLOSSAIRE (40 TERMES)

1. **HSM (Hardware Security Module)** — Boîtier/carte inviolable qui génère, stocke et
   utilise des clés crypto sans les exposer en clair. Cœur de la partie 1.
2. **KMS (Key Management System)** — Logiciel/service gérant le cycle de vie des clés
   (ex. HashiCorp Vault, AWS KMS). Partie 5.
3. **Root of Trust (RoT)** — Composant implicitement fiable qui amorce la chaîne de
   confiance (silicium, TPM, firmware). Partie 2.
4. **TPM 2.0 (Trusted Platform Module)** — Puce/firmware cryptographique d'une machine :
   mesures (PCR), scellement, attestation. Norme TCG.
5. **PCR (Platform Configuration Registers)** — 24 registres du TPM accumulant les
   hashes des composants démarrés.
6. **Secure Boot** — Vérification des signatures à chaque étape du boot ; bloque si non
   signé.
7. **Measured Boot** — Enregistrement (hash) de chaque étape du boot dans le TPM ;
   détecte sans bloquer.
8. **Attestation (à distance)** — Preuve cryptographique qu'une machine a démarré dans
   un état sain, vérifiée par un tiers.
9. **BMC (Baseboard Management Controller)** — Contrôleur de management d'un serveur
   (iLO, iDRAC, OpenBMC) : ordinateur dans l'ordinateur, à durcir en priorité.
10. **PKCS#11 (Cryptoki)** — API standard d'accès aux HSM/tokens (slots, sessions, objets).
11. **FIPS 140-2 / 140-3** — Certification NIST des modules crypto ; le niveau 3 exige
    une réponse active à l'effraction.
12. **Common Criteria (CC)** — Certification internationale (ISO 15408) par niveaux EAL.
13. **eIDAS** — Règlement UE sur l'identification électronique et les services de
    confiance (signature, horodatage).
14. **QSCD** — Qualified Signature Creation Device : dispositif de création de signature
    qualifiée eIDAS (typiquement un HSM certifié).
15. **AC (Autorité de Certification)** — Émet et révoque les certificats ; sa clé racine
    vit dans un HSM.
16. **CRL / OCSP** — Liste de révocation / protocole de vérification en ligne des
    certificats.
17. **TSA (Time-Stamping Authority)** — Autorité d'horodatage (RFC 3161) ; clé dans HSM,
    horloge fiable requise.
18. **DNSSEC** — Extension sécurisant le DNS par signatures ; la KSK est typiquement en
    HSM.
19. **KEK / DEK** — Key Encryption Key (chiffre les clés) / Data Encryption Key (chiffre
    les données). Hiérarchie de clés.
20. **Envelope encryption** — Pattern : DEK chiffrée par la KEK, données chiffrées par
    la DEK.
21. **MofN (M of N)** — Quorum : M parts sur N nécessaires pour reconstituer un secret.
22. **Cérémonie (key ceremony)** — Procédure formelle et auditée de génération d'une clé
    critique.
23. **Zeroization** — Effacement immédiat des clés en cas d'attaque détectée.
24. **TRNG / DRBG** — Générateur d'aléa vrai (physique) / déterministe (NIST 800-90A).
25. **FPGA** — Circuit logique reprogrammable après fabrication ; parallélisme et
    latence déterministe.
26. **Bitstream** — Fichier de configuration d'un FPGA ; à chiffrer et signer.
27. **LUT / DSP slice** — Briques de base d'un FPGA : table logique / multiplieur câblé.
28. **HLS (High-Level Synthesis)** — Compilation de C/C++ vers circuit FPGA.
29. **SmartNIC / DPU** — Carte réseau programmable déchargeant le CPU (réseau, sécu,
    stockage).
30. **ASIC** — Circuit figé à la fabrication ; optimal en volume, nul en flexibilité.
31. **PQC (Post-Quantum Cryptography)** — Crypto résistante à l'ordinateur quantique
    (ML-KEM, ML-DSA, SLH-DSA).
32. **ML-KEM / ML-DSA / SLH-DSA** — Standards NIST FIPS 203/204/205 (échange de clés,
    signatures, signatures hash).
33. **CNSA 2.0** — Suite d'algorithmes imposée par la NSA aux systèmes US (calendrier
    2027-2031).
34. **Confidential computing** — Chiffrement de la mémoire des VM (Intel TDX, AMD
    SEV-SNP, Arm CCA, NVIDIA H100 CC).
35. **TEE (Trusted Execution Environment)** — Environnement d'exécution isolé matériellement.
36. **NIST 800-88** — Norme de sanitization des médias (Clear / Purge / Destroy).
37. **SED (Self-Encrypting Drive)** — Disque à chiffrement intégré (TCG Opal) ;
    mise au rebut par crypto-erase.
38. **SBOM** — Software Bill of Materials : inventaire des composants logiciels d'un
    produit.
39. **SPDM** — Protocole DMTF d'attestation des composants matériels (cartes, disques).
40. **Redfish** — API REST standard (DMTF) de gestion des serveurs ; remplace IPMI.

# QUIZ : 10 QUESTIONS + RÉPONSES

**Q1. Pourquoi une clé privée ne doit-elle jamais exister en clair hors d'un HSM ?**
R : Parce que toute copie en clair (fichier, sauvegarde, snapshot VM) est une
compromission potentielle : quiconque lit le support possède la clé. Le HSM garantit
que la clé naît, vit et meurt à l'intérieur du périmètre inviolable — seules les
opérations (signer, déchiffrer) sortent.

**Q2. Quelle différence entre secure boot et measured boot ?**
R : Le secure boot **vérifie** la signature et **bloque** si non signé ; le measured
boot **enregistre** le hash dans le TPM sans bloquer. Il faut les deux : le premier
empêche le connu-mauvais, le second permet de **détecter** (via attestation) ce qui a
vraiment démarré.

**Q3. Pourquoi le BMC est-il le talon d'Achille d'un serveur ?**
R : C'est un ordinateur indépendant (SoC, OS, réseau) avec accès total au serveur
(allumage, KVM, ISO), **invisible depuis l'OS**. Compromis (mot de passe par défaut,
firmware non patché), il donne un contrôle persistant que ni la réinstallation ni
l'antivirus ne voient.

**Q4. Dans quel cas choisir un FPGA plutôt qu'un GPU ?**
R : Quand le workload est **stable** (pas de changement fréquent), que la **latence
déterministe** est critique (réseau, trading, protocoles), et que l'écosystème logiciel
du GPU n'apporte rien. Pour l'IA générative et tout ce qui évolue vite : GPU.

**Q5. Qu'est-ce qu'un quorum MofN et pourquoi ne pas mettre toutes les parts dans le
même coffre ?**
R : La clé est découpée en N parts dont M sont nécessaires pour la reconstituer :
personne seul ne peut la restaurer. Si toutes les parts sont au même endroit, un seul
vol/accès les récupère toutes — le quorum ne vaut que par la **séparation physique et
organisationnelle**.

**Q6. Pourquoi l'auto-unseal de Vault est-il lié à l'alimentation électrique ?**
R : Sans auto-unseal, Vault redémarre scellé et exige des humains (Shamir). Avec
auto-unseal sur HSM, le descellement est automatique **si le HSM est disponible avant
Vault**. Après une coupure, l'ordre de redémarrage (HSM → Vault → applis) et
l'alimentation secourue du HSM conditionnent la reprise — d'où la règle « jamais sur
onduleur non secouru ».

**Q7. Que signifie « harvest now, decrypt later » et que faire dès 2026 ?**
R : Les adversaires enregistrent le trafic chiffré aujourd'hui pour le déchiffrer avec
un futur ordinateur quantique. Action : déployer le **TLS hybride** (X25519MLKEM768,
déjà supporté par les navigateurs et clouds), inventorier la crypto, exiger la
**crypto-agilité** des HSM/KMS, planifier la migration ML-DSA des signatures.

**Q8. Quelle est la méthode de destruction d'un disque contenant des données critiques
en fin de vie ?**
R : Selon NIST 800-88, niveau **Destroy** (broyage/désintégration) pour le top
critique, ou **Purge** (crypto-erase sur SED) avec **PV de destruction** tracé par
numéro de série. Un disque « jeté » lisible = fuite de données.

**Q9. Pourquoi tester la restauration des sauvegardes de clés une fois par an ?**
R : Parce qu'une sauvegarde jamais testée n'est pas une sauvegarde : mots de passe
perdus, KEK détruites, procédures obsolètes, quorum injoignable ne se révèlent qu'au
test. Le test DR crypto (section 138) mesure aussi le **RTO réel**.

