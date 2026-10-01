---
id: collect-261001-rattrapage/rattrapage/datacenter-securite-trust-guide-19
title: "Sécurité datacenter — HSM, Root of Trust, FPGA, sécurité physique, gestion des clés"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "AWS", "Intel"]
dates: []
keywords: ["datacenter", "agent", "amd", "aws", "capex", "intel", "open source"]
source: docs/RAG/collect-261001-rattrapage/datacenter_securite_trust_guide.md
source_anchor: ""
source_lines: [2362, 2500]
sha256: 8ac244827454f5306224e380e0789c6434740708ca7ea77626d8e6228aac3eba
---

# Sécurité datacenter — HSM, Root of Trust, FPGA, sécurité physique, gestion des clés

| Poste (5 ans) | A : 2× appliance achetée | B : AWS CloudHSM (2 HSM) | C : 4× YubiHSM 2 + Vault |
|---|---|---|---|
| CAPEX matériel | ~50 000 € | 0 € | ~2 600 € |
| Support/licences 5 ans | ~45 000 € | 0 € | 0 € |
| Abonnement cloud 5 ans | 0 € | ~130 000 $ | 0 € |
| Prestation install. | ~8 000 € | ~3 000 € | ~5 000 € |
| Électricité 5 ans | ~1 500 € | 0 € | ~100 € |
| Temps humain (estimé) | ~30 k€ | ~15 k€ | ~35 k€ |
| **Total ordre de grandeur** | **~135 k€** | **~145 k$** | **~43 k€** |

Lecture : le scénario C (YubiHSM) ne convient que si le débit réel est faible
(dizaines d'ops/s) et sans besoin de clustering natif — sinon c'est A ou B. Entre A et
B, le **seuil** est à ~3-4 ans (section 48) ; B gagne en flexibilité et en absence de
gestion matérielle. Tous les chiffres sont **indicatifs, à vérifier** sur devis.

## 166. Migrer une PKI existante vers le HSM

1. **Inventaire** : où sont les clés aujourd'hui ? (fichiers, keystores Java, Windows...)
   — c'est souvent la découverte la plus pénible.
2. **Nouvelle hiérarchie** : générer la nouvelle racine **dans** le HSM (cérémonie).
   Ne pas « importer » l'ancienne racine sauf contrainte absolue (et dans ce cas,
   cérémonie d'import traçée, puis destruction des copies logicielles avec PV).
3. **Coexistence** : faire signer une intermédiaire par la nouvelle racine, migrer les
   émettrices par vagues, maintenir l'ancienne chaîne le temps de l'expiration.
4. **Nettoyage** : détruire les anciennes clés logicielles (PV, section 100), révoquer
   les anciens certificats en fin de vie.
5. **Durée typique** : 3-6 mois pour une PKI d'entreprise — prévoir la communication
   (les applications doivent basculer de chaîne).

## 167. FAQ HSM (10 questions)

1. **Un HSM ralentit-il les applications ?** La latence réseau (~1-5 ms) ne se sent que
   sur des opérations à très haute fréquence ; le cache et la résumption TLS absorbent
   le reste. Mesurer avant d'optimiser.
2. **Faut-il un HSM pour Let's Encrypt ?** Non — les certificats publics à 90 jours se
   gèrent en logiciel ; le HSM sert aux clés **longues et critiques** (racines, KEK).
3. **HSM + virtualisation ?** Appliance réseau : aucun problème (le HSM est distant).
   Carte PCIe : PCI passthrough vers une VM dédiée.
4. **Que se passe-t-il si on perd le PIN SO ?** Selon les modèles : réinitialisation
   d'usine (= perte des clés) ou procédure constructeur. D'où le MofN (section 43).
5. **Peut-on virtualiser un HSM ?** Non — c'est du matériel. On virtualise l'**accès**
   (partitions, clients multiples).
6. **HSM et sauvegardes Veeam ?** Les clés ne sont pas dans les VM sauvegardées si
   elles sont dans le HSM — c'est le but. Sauvegarder **séparément** (section 41).
7. **Durée de vie d'un HSM ?** 7-10 ans typiques (batterie, support constructeur).
   Planifier le renouvellement à 5 ans.
8. **HSM d'occasion ?** Possible via reconditionneurs sérieux **avec PV de reset** ;
   jamais sans (section 66). Budget divisé par 2-3, mais support à vérifier.
9. **Un seul HSM suffit-il ?** Pour une racine offline : oui (avec backup au coffre).
   Pour de la production : **non** — cluster de 2 minimum.
10. **Le quantique casse-t-il mon HSM ?** Pas le boîtier : les algorithmes. D'où la
    crypto-agilité et le PQC (partie 9). Le HSM protège les clés **quel que soit**
    l'algorithme.

# PARTIE 11 — APPROFONDISSEMENTS SERVEURS / ROOT OF TRUST

## 168. Inventaire Redfish : script et automatisation

Redfish (HTTPS/JSON) permet d'inventorier firmware, Secure Boot et TPM sans agent :

```
# Exemple de principe (à adapter, identifiants en coffre-fort)
curl -sk -u admin:$PASS https://bmc-serveur01/redfish/v1/Systems/1 \
  | jq '.BiosVersion, .SecureBoot'
curl -sk -u admin:$PASS https://bmc-serveur01/redfish/v1/UpdateService/FirmwareInventory
```

Automatisation : un script (Ansible/Python) qui parcourt les BMC, compare les versions à
la **baseline validée**, et alerte sur les écarts. Fréquence : hebdomadaire. C'est la
base du registre de firmware (section 62) et de l'audit Secure Boot (section 131).
Stocker les identifiants BMC dans un **coffre** (Vault), jamais dans le script.

## 169. Keylime : déployer l'attestation en pratique

**Keylime** (open source, CNCF) : l'agent sur chaque serveur envoie les mesures (quotes
TPM) au vérificateur, qui compare aux valeurs de référence et peut **révoquer** l'accès
(ex. retirer un nœud du cluster). Déploiement type :

1. Installer le vérificateur + registrar sur une machine dédiée (VLAN management).
2. Enregistrer les EK des serveurs (prouver l'authenticité des TPM).
3. Capturer les **mesures de référence** sur un serveur sain (golden).
4. Déployer l'agent, brancher la décision à l'orchestrateur (ex. : un nœud non attesté
   ne reçoit pas ses secrets / est exclu du scheduling).
5. **Processus** : à chaque mise à jour firmware/noyau, re-capturer les références —
   sinon alerte massive (le piège n°1 de l'attestation).

## 170. Sceller LUKS au TPM : pas-à-pas (principe)

1. TPM provisionné, Secure Boot activé (sinon le scellement ne veut rien dire).
2. Enroller la clé LUKS dans le TPM liée aux PCR 0+7 (firmware + secure boot) :
   `systemd-cryptenroll --tpm2-device=auto --tpm2-pcrs=0+7 /dev/sdX`.
3. Tester : reboot → déverrouillage **sans mot de passe** si les PCR sont conformes.
4. Tester le cas d'échec : modifier un réglage UEFI → le déverrouillage doit **échouer**
   (puis déblocage manuel avec la passphrase de secours).
5. Conserver la **passphrase de secours** au coffre (c'est elle qui sauve en cas de
   changement de carte mère).

Limite : protège contre le **vol du disque**, pas contre un attaquant avec accès au
serveur allumé (la clé est en RAM).

## 171. Durcir GRUB et le boot Linux

- Mot de passe GRUB (empêche l'édition des paramètres au boot — le contournement
  `init=/bin/bash`).
- `shim` + GRUB + noyau **signés** (Secure Boot), modules noyau signés
  (`CONFIG_MODULE_SIG_FORCE`).
- Paramètres noyau durcis : `slab_nomerge`, `pti=on`, `vsyscall=none`, IOMMU activé
  (`intel_iommu=on` / `amd_iommu=on` — protège aussi contre le DMA malveillant, en lien
  avec les attaques type DDRop).
- `/boot` en lecture seule si possible, vérification périodique des hashes.

## 172. BMC : centraliser la gestion d'un parc

Au-delà de 10 serveurs, gérer les BMC un par un ne passe plus à l'échelle :

- **Inventaire centralisé** : CMDB avec IP BMC, versions, mots de passe (coffre).
- **Mises à jour groupées** : via les outils constructeur (HPE OneView, Dell OME,
  Supermicro SSM) ou Redfish scripté — par vagues, avec validation.
- **Comptes** : un compte d'administration central (LDAP/AD) + un compte local de
  secours par BMC (mot de passe unique au coffre).
- **Supervision** : tous les syslog BMC vers le SIEM, tableau de bord des versions.
- **Règle** : aucun BMC accessible sans passer par le **bastion** du VLAN management.

## 173. Microcodes CPU : pourquoi et comment les maintenir

Les **microcodes** corrigent des bugs CPU, dont des failles de sécurité (Spectre/Meltdown
et leurs descendants). Ils se chargent au boot (via le firmware UEFI ou l'OS). En
datacenter :

- Suivre les bulletins Intel/AMD, appliquer via mise à jour UEFI **ou** package OS
  (`intel-microcode` / `amd64-microcode` sous Debian/Ubuntu).
- Après mise à jour : **re-mesurer** (les PCR peuvent changer) et re-tester les perfs
  (certains mitigations coûtent des % de performance — à mesurer, pas à subir).
- Un parc avec des microcodes de 3 ans = des failles connues exploitables localement.

## 174. Cas chiffré : sécuriser 50 serveurs (BOM + planning)

**Périmètre** : 50 serveurs 2U, dont 5 critiques (PKI, Vault, hyperviseurs).

