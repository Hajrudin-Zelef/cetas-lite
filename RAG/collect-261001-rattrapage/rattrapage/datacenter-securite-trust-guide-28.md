---
id: collect-261001-rattrapage/rattrapage/datacenter-securite-trust-guide-28
title: "Sécurité datacenter — HSM, Root of Trust, FPGA, sécurité physique, gestion des clés"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-09-27"]
keywords: []
source: docs/RAG/collect-261001-rattrapage/datacenter_securite_trust_guide.md
source_anchor: ""
source_lines: [3631, 3767]
sha256: 4a227f522a469f6fc8abb5c952e8b77390cd7f3b1f28a9a3f2aad42f76298151
---

# Sécurité datacenter — HSM, Root of Trust, FPGA, sécurité physique, gestion des clés

1. **Lister** les 10 clés les plus critiques et où elles sont (fichier ? HSM ? tête
   de quelqu'un ?).
2. **Vérifier** 5 BMC pris au hasard : mot de passe par défaut ? firmware à jour ?
   VLAN dédié ?
3. **Contrôler** que les sauvegardes de clés existent **et** qu'on sait les restaurer
   (demander la date du dernier test).
4. **Vérifier** que les équipements crypto sont sur circuits secourus (suivre les
   câbles, pas les plans).
5. **Planifier** la réunion de cadrage (modèle de menaces, section 184) avec la
   direction — 2 h, ordre du jour écrit.

Ces 5 actions ne coûtent rien, prennent une semaine, et révèlent 90 % des problèmes.
Le reste du guide sert à les corriger durablement.

---

> **Fin du guide.** 240 sections, glossaire 40 termes, quiz 10 Q/R, 25 pièges terrain,
> 19 parties + annexes A-F + 5 fiches réflexes. Rédigé le 27/09/2026 — références produits vérifiées par
> recherche web le 27/09/2026, prix et détails non sourcés marqués « à vérifier »,
> références introuvables marquées « non trouvée au 27/09/2026 ». Aucune invention.
> Revue conseillée : T1 2027.

# PARTIE 18 — DICTIONNAIRE DES ATTAQUES ET CONTRE-MESURES

> Pourquoi cette partie : on ne défend bien que ce qu'on comprend. Chaque attaque est
> décrite avec sa contre-mesure dans ce guide. Faits publics établis ; détails
> techniques pointus marqués « à vérifier ».

## 224. Attaques contre les HSM : canaux auxiliaires et glitch

- **Analyse de consommation / électromagnétique** : déduire la clé en observant la
  consommation ou les émissions pendant les opérations crypto. Contre-mesure : les HSM
  certifiés FIPS 140-2 L3/L4 implémentent des contre-mesures (masquage, blindage) —
  d'où l'importance de la certification (section 5), pas d'un « boîtier maison ».
- **Glitch (tension/horloge)** : perturber la puce pour sauter une vérification.
  Contre-mesure : capteurs de tension/fréquence avec zeroization (section 3).
- **Attaque au froid (cold boot adaptée)** : refroidir la SRAM pour prolonger la
  rémanence après coupure. Contre-mesure : capteur de température + effacement.
- **Extraction par le logiciel** : la voie la plus fréquente en pratique — pas le
  boîtier, mais une **API mal configurée** (clé marquée extractible, PIN faible).
  Contre-mesure : clés non extractibles, PIN forts, moindre privilège (sections 39, 44).

## 225. Attaques contre le firmware : rootkits UEFI

Cas publics établis : **LoJax** (2018, premier rootkit UEFI observé in the wild),
**TrickBoot** (2020, module UEFI), **CosmicStrand** (2022), **BlackLotus** (2023,
contournement du Secure Boot via binaire signé révoqué tardivement). Le pattern :

1. L'attaquant obtient un accès admin/OS.
2. Il écrit dans la flash SPI (variable UEFI, Option ROM, ou bootloader).
3. Le malware **survit** à la réinstallation de l'OS et au changement de disque.

Contre-mesures (ce guide) : Secure Boot + measured boot + attestation (55-57), Boot
Guard/PSB (58), mises à jour firmware (62), scellement TPM (53). Et surtout : **détecter**
— un rootkit UEFI ne se voit pas depuis l'OS, seulement par mesure/attestation.

## 226. Attaques contre les BMC : l'historique parle

Les BMC ont connu des vagues de failles critiques (exécution de code à distance,
contournement d'authentification) chez tous les constructeurs — les bulletins iLO/iDRAC/
XCC en témoignent chaque année. Le pattern d'exploitation :

1. Scan du réseau à la recherche des ports BMC (623/UDP IPMI, 443/HTTPS).
2. Connexion avec identifiants par défaut ou exploitation d'une CVE non patchée.
3. Montage d'une image ISO via KVM virtuel → boot sur l'image → contrôle total.
4. **Persistance** : le BMC n'est pas réinstallé avec l'OS.

Contre-mesures : checklist section 60 (VLAN dédié, mots de passe uniques, patch < 30
jours, IPMI 1.5 désactivé). En 2026, un BMC exposé sur Internet est compromis en
**quelques heures** (les botnets scannent en permanence) — ne jamais l'exposer.

## 227. Attaques réseau contre la zone crypto

- **Interception** : si le canal client↔HSM n'est pas en TLS mutuel, un attaquant
  réseau peut observer ou rejouer. Contre-mesure : NTLS/TLS mutuel (section 35).
- **Déni de service** : inonder le HSM de requêtes — il ne « casse » pas, mais les
  applis légitimes sont ralenties. Contre-mesure : firewall avec rate-limiting, HA.
- **Usurpation de client** : se faire passer pour un serveur applicatif autorisé.
  Contre-mesure : authentification mutuelle par certificats (pas seulement par IP).
- **Attaque de l'homme du milieu sur l'administration** : intercepter la session
  d'admin du HSM. Contre-mesure : poste d'admin dédié, sur le VLAN crypto uniquement,
  pas de Wi-Fi.

## 228. Attaques physiques : evil maid, cold boot, DMA

- **Evil maid** : accès physique bref (femme/homme de ménage, hôtel) pour modifier le
  bootloader. Contre-mesure : Secure Boot + scellement TPM (le boot modifié ne
  déverrouille plus le disque), contrôle d'accès physique (partie 4).
- **Cold boot** : refroidir la RAM, la transplanter, lire les clés. Contre-mesure :
  chiffrement mémoire (confidential computing, partie 9), extinction complète plutôt
  que mise en veille pour les machines sensibles.
- **DMA (Thunderbolt/PCIe)** : un périphérique malveillant lit la RAM via DMA.
  Contre-mesure : IOMMU activé (section 171), ports DMA désactivés sur les machines
  sensibles, contrôle physique des ports.
- **Vol de disque** : sans chiffrement, lecture directe. Contre-mesure : LUKS/BitLocker
  + clé scellée au TPM (section 170).

## 229. Attaques supply chain : cas publics et parades

Cas publics établis (patterns) : composants réseau avec firmwares modifiés découverts à
la réception, disques reconditionnés vendus comme neufs avec des firmwares anciens et
vulnérables, interceptions de colis. Parades (sections 63-66) : circuit officiel,
vérification des scellés et numéros de série, SBOM, attestation SPDM, premier boot en
quarantaine, reflash systématique du reconditionné. Le coût de la parade (quelques
heures par serveur) est négligeable face au coût d'un composant compromis en production.

## 230. Harvest now, decrypt later : le dossier complet

L'attaque la plus **stratégique** de la décennie :

1. **Aujourd'hui** : l'adversaire enregistre le trafic chiffré (TLS, VPN, archives) —
   coût du stockage dérisoire.
2. **Demain (2029-2032+, incertain)** : un ordinateur quantique cryptographiquement
   pertinent casse RSA/ECDH (Shor) et déchiffre rétroactivement.
3. **Cible** : les données à **longue durée de confidentialité** (secrets industriels,
   données de santé, communications sensibles, clés à longue vie).

Défense **dès 2026** : TLS hybride X25519MLKEM768 (section 151), inventaire crypto et
priorisation par durée de confidentialité, crypto-agilité des HSM/KMS (section 152),
ne pas attendre FIPS 206. Ce n'est pas de la paranoïa : c'est de la **gestion du risque
à horizon 10 ans**, comme on provisionne une garantie décennale.

## 231. Attaques contre les sauvegardes

- **Vol de bandes/disques** : transport non sécurisé, prestataire peu regardant.
  Contre-mesure : chiffrement (section 208) + transport scellé + traçabilité.
- **Ransomware sur les sauvegardes en ligne** : chiffrement des dépôts accessibles.
  Contre-mesure : immuabilité + air gap (section 209).
- **Sauvegarde non chiffrée « temporaire »** : devient permanente. Contre-mesure :
  chiffrement **par défaut**, pas en option.
- **Restauration avec une clé détruite** : la sauvegarde est illisible. Contre-mesure :
  coordonner rétention et cycle de vie des clés (sections 109, 208), tester.

## 232. Ingénierie sociale sur les custodians

Le MofN (section 43) repose sur des **humains** : ils sont la cible.

