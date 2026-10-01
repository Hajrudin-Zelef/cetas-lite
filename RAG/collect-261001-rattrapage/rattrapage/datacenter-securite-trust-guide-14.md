---
id: collect-261001-rattrapage/rattrapage/datacenter-securite-trust-guide-14
title: "Sécurité datacenter — HSM, Root of Trust, FPGA, sécurité physique, gestion des clés"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["datacenter", "agent", "distribution"]
source: docs/RAG/collect-261001-rattrapage/datacenter_securite_trust_guide.md
source_anchor: ""
source_lines: [1663, 1800]
sha256: 022c121d92b210b3e4714934559cf51d6103776b1c9fe6c7025e461efd29234a
---

# Sécurité datacenter — HSM, Root of Trust, FPGA, sécurité physique, gestion des clés

L'**ISO/IEC 27001** est la norme de référence pour un **SMSI** (système de management de la
sécurité de l'information). Principe : on identifie les risques, on applique des mesures
(l'Annexe A : 93 contrôles), on **améliore en continu** (roue de Deming). La certification
par un organisme accrédité est un **argument commercial** fort (exigée dans de nombreux
appels d'offres) et un **cadre** qui force la rigueur. Pour un datacenter/PME : viser la
certification est un projet de **6 à 18 mois** ; en attendant, appliquer les contrôles
pertinents (dont ceux de la section 122) apporte déjà 80 % de la valeur.

## 122. Annexe A : les contrôles crypto (A.10 et associés)

Les contrôles directement liés à ce guide (ISO 27001:2022, Annexe A) :

- **A.10 — Cryptographie** : politique d'utilisation de la cryptographie, **gestion des
  clés** sur tout leur cycle de vie (c'est exactement la partie 5 de ce guide).
- **A.8.10** : suppression des informations (cf. NIST 800-88, section 100).
- **A.8.13** : sauvegarde des informations (cf. section 115).
- **A.5.33** : protection des procédures d'exploitation — dont les **cérémonies**
  (section 42).
- **A.8.15** : journalisation (cf. section 120), **A.8.16** : surveillance.
- **A.5.14** : contact avec les autorités / groupes spécialisés (utile pour remonter
  les vulnérabilités firmware).
- **A.8.9** : gestion de la configuration (inventaire firmware, section 62).

En audit, l'auditeur demandera : la **politique crypto écrite**, l'**inventaire des clés**,
les **PV de cérémonie**, les **preuves de rotation** et les **journaux**. Ce guide fournit
le fond ; il reste à l'écrire dans vos documents.

## 123. PCI DSS v4 : si vous touchez au paiement

Si vous stockez, traitez ou transmettez des données de cartes bancaires, **PCI DSS**
s'applique — et ses exigences crypto sont strictes : chiffrement des données carte au
repos et en transit, **gestion des clés documentée** (génération, distribution,
stockage, rotation au moins annuelle, destruction), et pour les opérations PIN : HSM
certifié **PCI HSM** (voir sections 13, 25, 28). La v4.x insiste sur l'approche
personnalisée et la validation continue. Point pratique : la conformité PCI est
**périmétrique** — isolez le périmètre carte du reste (segmentation réseau prouvée)
pour réduire le coût d'audit.

## 124. RGPD et chiffrement : l'article 32

Le RGPD n'impose pas le chiffrement nommément, mais son **article 32** exige des mesures
« appropriées » dont le chiffrement des données personnelles est l'exemple cité. En
pratique : chiffrement des données personnelles **au repos** (disques, bases, sauvegardes)
et **en transit** (TLS), avec gestion des clés sérieuse (partie 5). Bénéfice concret :
des données personnelles **chiffrées avec des clés non compromises** peuvent éviter la
notification de violation (considérant 83) — c'est un argument ROI direct pour le HSM.

## 125. eIDAS : signature et horodatage

Le règlement **eIDAS** (UE) encadre la signature électronique et les services de confiance.
Niveaux : simple, avancée, **qualifiée** (la plus forte, équivalent manuscrit). Pour le
qualifié : certificat qualifié + **QSCD** (dispositif de création de signature qualifié —
typiquement un HSM certifié, voir section 7) + prestataire de services de confiance
qualifié. Si vos clients exigent du « qualifié », le choix du HSM et du prestataire est
contraint — à cadrer **avant** l'achat du matériel.

## 126. Exigences clients : répondre aux questionnaires

Les grands clients envoient des **questionnaires de sécurité** (type SIG, CAIQ). Les
questions crypto reviennent toujours : « où sont stockées les clés ? » (→ HSM, partie 1),
« comment sont-elles sauvegardées ? » (→ section 41), « quelle rotation ? » (→ section
108), « qui y a accès ? » (→ section 110), « quels journaux ? » (→ section 120).
**Préparer un dossier standard** avec ces réponses (sans divulguer de secrets
opérationnels) fait gagner des semaines par appel d'offres. Les certifications (ISO 27001,
rapports d'audit de l'hébergeur) répondent à 50 % des questions d'un coup.

## 127. Journaux et audit : ce qu'il faut tracer

Synthèse des exigences de traçabilité de ce guide :

| Source | Événements à tracer | Rétention indicative |
|---|---|---|
| HSM | Génération, utilisation, export, destruction de clés ; échecs d'auth | 1-3 ans |
| KMS (Vault...) | Unseal, lecture de secrets, émissions de certificats | 1-3 ans |
| BMC/serveurs | Connexions, montages ISO, mises à jour firmware | 1 an |
| Contrôle d'accès physique | Entrées/sorties salles critiques | 1 an |
| Vidéosurveillance | Enregistrements zones critiques | 90 jours (section 95) |
| Cérémonies | PV signés (papier + numérique) | Durée de vie de la clé + 5 ans |

Rétentions **indicatives** — à aligner sur vos obligations contractuelles et légales
(avec le juriste, voir l'avertissement en tête de partie).

## 128. Politiques à écrire : la liste minimale

Sans ces documents, l'audit échoue quel que soit le matériel :

1. **Politique de cryptographie** : algos autorisés/interdits, tailles de clés minimales
   (ex. RSA ≥ 2048, interdiction de DES/3DES sauf legacy documenté).
2. **Politique de gestion des clés** : cycle de vie, rotation, MofN, rôles (partie 5).
3. **Procédure de cérémonie** : script pas-à-pas (section 42).
4. **Politique de contrôle d'accès physique** : zonage, visiteurs, habilitations
   (partie 4).
5. **Procédure de destruction des médias** : NIST 800-88 adaptée (section 100).
6. **Plan de réponse aux incidents** : dont les runbooks crypto (sections 136-137).
7. **Politique de mise à jour des firmwares** : registre, tests, rollback (section 62).

Modèles : partir des exemples de l'**ANSSI** (guides d'hygiène) et des templates ISO 27001
— puis les **adapter** à votre réalité, jamais du copier-coller pur.

# PARTIE 7 — CHECKLISTS PRATIQUES : LA MISE EN SERVICE SÉCURISÉE

## 129. Checklist : réception et mise en service d'un serveur

**À la réception** (voir aussi section 65) :

- [ ] Emballage/scellés intacts, n° de série = bon de livraison.
- [ ] Inventaire physique : CPU, RAM, disques, cartes, alimentations.
- [ ] Premier boot en **zone de quarantaine** (VLAN isolé, pas de production).

**Firmware & plateforme** :

- [ ] UEFI à la version validée, **Secure Boot activé** (clés propres si images maison).
- [ ] **TPM 2.0** présent et provisionné (`tpm2_pcrread` OK, EK certifié).
- [ ] Boot Guard / PSB activé si disponible (**à vérifier** selon OEM).
- [ ] Option ROM des cartes : versions notées, Secure Boot les vérifie.

**BMC** (détail section 130) :

- [ ] Firmware BMC à jour, mot de passe par défaut **changé** (unique par serveur).
- [ ] BMC sur **VLAN management** dédié, IP statique, DNS correct.
- [ ] Services inutiles désactivés (Telnet, HTTP, IPMI 1.5).
- [ ] Certificat du BMC remplacé par un certificat de la PKI interne.

**OS & durcissement** :

- [ ] Installation depuis une **image signée/vérifiée** (hash contrôlé).
- [ ] Chiffrement disque (LUKS/BitLocker) avec clé **scellée au TPM** (PCR 0-7).
- [ ] Comptes par défaut supprimés/renommés, SSH par clés uniquement, root désactivé.
- [ ] Agent de supervision + syslog vers SIEM, NTP (NTS si possible).
- [ ] Pare-feu local : uniquement les ports nécessaires.

**Documentation** :

- [ ] Fiche serveur : n° de série, versions firmware, empreintes de clés, VLAN, rôle.
- [ ] PV de mise en service signé et archivé.

## 130. Checklist : durcissement BMC (iLO/iDRAC/OpenBMC)

