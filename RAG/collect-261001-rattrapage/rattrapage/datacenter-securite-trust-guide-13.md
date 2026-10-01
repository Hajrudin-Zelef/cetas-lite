---
id: collect-261001-rattrapage/rattrapage/datacenter-securite-trust-guide-13
title: "Sécurité datacenter — HSM, Root of Trust, FPGA, sécurité physique, gestion des clés"
domain: rattrapage
role: reference
task: reference
actors: ["AWS", "Microsoft"]
dates: []
keywords: ["datacenter", "aws", "open source"]
source: docs/RAG/collect-261001-rattrapage/datacenter_securite_trust_guide.md
source_anchor: ""
source_lines: [1526, 1662]
sha256: 0b7af920089f63392499fff02a7403a780e29162ae8afa6015bea1a8a6650c62
---

# Sécurité datacenter — HSM, Root of Trust, FPGA, sécurité physique, gestion des clés

Aucune personne ne cumule **génération + utilisation + audit**. Les cumuls sont
documentés comme exceptions temporaires, avec date de fin.

## 111. Intégration HSM↔KMS : les 3 architectures

1. **KMS logiciel + HSM en backend** : le KMS (ex. Vault) stocke sa clé maîtresse dans le
   HSM (seal/unseal), les opérations courantes restent logicielles. Bon compromis
   coût/sécurité — le plus courant en PME/ETI.
2. **HSM direct** : les applications appellent le HSM en PKCS#11, pas de KMS
   intermédiaire. Simple, mais pas de gestion de cycle de vie (rotation, politiques) —
   à réserver aux usages simples (PKI, TSA).
3. **KMS cloud + Cloud HSM** : le KMS du provider (AWS KMS, Azure Key Vault Managed HSM)
   adossé à des HSM dédiés. Idéal cloud-native ; la souveraineté dépend du modèle
   (BYOK/HYOK, section 116).

## 112. HashiCorp Vault : panorama

**Vault** est le KMS logiciel de référence (open source + entreprise) : stockage de
secrets, PKI intégrée, chiffrement as-a-service (Transit), identités dynamiques
(credentials à usage unique pour bases de données). Points clés pour ce guide :

- Le **seal/unseal** : Vault démarre « scellé », sa clé maîtresse est découpée en parts
  (Shamir) ou confiée à un HSM/cloud-KMS (**auto-unseal**).
- Le backend de stockage doit être **chiffré et sauvegardé** (Raft intégré ou Consul).
- La **PKI intégrée** peut émettre des certificats internes avec l'AC racine sur HSM.

## 113. Vault : auto-unseal avec HSM

Sans auto-unseal, chaque redémarrage de Vault exige que des humains apportent les parts
Shamir — **incompatible avec une infrastructure qui redémarre seule** (panne électrique,
voir partie 8). L'**auto-unseal** confie le descellement à un HSM (via PKCS#11) ou à un
KMS cloud : au démarrage, Vault demande au HSM de déchiffrer sa clé maîtresse, **sans
intervention humaine**. Conditions : le HSM doit être **disponible avant Vault**
(ordre de démarrage documenté), sur **alimentation secourue** (section 140), et l'accès
réseau HSM↔Vault sécurisé (VLAN crypto, section 35). C'est le chaînon qui rend la
continuité électrique critique pour la crypto (voir section 141).

## 114. EJBCA : la PKI open source

**EJBCA** (PrimeKey/Cybertrust) est l'AC logicielle open source de référence : émission de
certificats X.509, OCSP, CMP/SCEP pour l'enrôlement, profils de certificats. Branchée en
PKCS#11 sur n'importe quel HSM du marché, elle couvre : PKI interne d'entreprise, AC pour
IoT/industrie, certificats serveurs. Alternative : **AD CS** (Microsoft) pour les
environnements Windows — le choix dépend de l'écosystème, pas de la sécurité (les deux
s'adossent au HSM).

## 115. Sauvegardes chiffrées des clés

Les sauvegardes du KMS/HSM suivent les mêmes règles que les clés elles-mêmes :

- **Chiffrées** (par une KEK dédiée, stockée en MofN), jamais en clair.
- **Testées** : restauration annuelle sur un environnement isolé (section 138).
- **Géographiquement séparées** : une copie sur le site principal (coffre), une sur le
  site distant — le vol/incendie d'un site ne doit pas tout emporter.
- **Rétention** : alignée sur la durée de vie des données chiffrées (une sauvegarde
  chiffrée avec une clé détruite = données perdues — coordonner avec la section 109).
- **Journalisées** : chaque export/import de clé est un événement d'audit (section 120).

## 116. BYOK et HYOK : apporter ses clés au cloud

- **BYOK** (Bring Your Own Key) : vous générez la clé dans **votre** HSM on-prem, vous
  l'importez (wrappée) dans le KMS cloud. Vous gardez la **génération**, le provider
  gère l'**usage**. Révocation possible en supprimant la clé (avec les conséquences :
  données illisibles).
- **HYOK** (Hold Your Own Key) : la clé **ne quitte jamais** votre HSM ; le cloud
  appelle votre HSM pour chaque opération (ex. Double Key Encryption chez Utimaco/
  Microsoft). Sécurité maximale, **latence** et **dépendance réseau** à assumer.
- Choix : BYOK pour la conformité standard, HYOK pour les données les plus sensibles
  ou les exigences de souveraineté strictes.

## 117. Tableau comparatif KMS

| Solution | Modèle | HSM backend | Idéal pour | Coût indicatif |
|---|---|---|---|---|
| HashiCorp Vault (OSS) | Logiciel on-prem | Oui (PKCS#11, auto-unseal) | ETI/datacenter, contrôle total | 0 € + infra + expertise |
| HashiCorp Vault (Enterprise) | Logiciel on-prem | Oui + HSM as-a-service | Multi-sites, support | Licence (à vérifier) |
| AWS KMS | Cloud managé | Oui (CloudHSM backend en option) | Workloads AWS | ~1 $/mois/clé + requêtes |
| Azure Key Vault / Managed HSM | Cloud managé | Oui (Managed HSM = dédié) | Workloads Azure | Voir section 19 |
| GCP Cloud KMS (+ Cloud HSM) | Cloud managé | Oui (niveau HSM) | Workloads GCP | ~1 $/mois/version + requêtes |
| EJBCA | Logiciel on-prem (PKI) | Oui (PKCS#11) | AC interne, IoT | 0 € (OSS) + support en option |

## 118. Cas chiffré PME : BOM KMS + HSM complète

**Besoin** : PME 80 personnes, PKI interne (Wi-Fi EAP-TLS, VPN), chiffrement des
sauvegardes, signature des documents RH.

| Poste | Choix | Coût indicatif |
|---|---|---|
| HSM racine | 2× YubiHSM 2 (1 prod + 1 backup coffre) | ~1 300 € |
| Serveur PKI | VM existante + EJBCA (OSS) | 0 € |
| KMS applicatif | Vault OSS sur 3 VM (Raft) | 0 € + infra existante |
| Auto-unseal | YubiHSM 2 n°3 dédié (ou Shamir manuel si redémarrages rares) | ~650 € |
| Coffre-fort | Ignifuge, pour backup HSM + PV cérémonie | ~1 500 € |
| Prestation cérémonie + formation | 3 jours | ~5 000 € |
| **Total** | | **~8 500 €** |

C'est le **ticket d'entrée réaliste** d'une crypto bien gérée en PME. Le poste le plus
sous-estimé : le **temps humain** (cérémonie, revues, tests de restauration).

## 119. Erreurs classiques de gestion des clés

1. Clé maîtresse **dans un fichier** sur le serveur applicatif (« en attendant le HSM »).
2. **Même clé** pour chiffrer et signer, ou pour la prod et les tests.
3. Sauvegarde des clés **en clair** « parce que c'est plus simple à restaurer ».
4. Rotation **jamais** faite (« on n'a jamais eu le temps »).
5. Quorum MofN dont toutes les parts sont **dans le même coffre**.
6. Vault en **Shamir manuel** sur une infra qui redémarre seule après coupure.
7. Clés **codées en dur** dans le code ou les playbooks Ansible (secrets → Vault/Ansible
   Vault chiffré, jamais en clair dans git).
8. **Anciennes clés jamais détruites** : elles traînent dans les sauvegardes pendant
   des années.
9. Pas d'**inventaire** : personne ne sait combien de clés existent ni où.
10. Certificats **expirés** en production parce que le renouvellement était manuel.

## 120. Journaux d'audit des opérations de clés

Tout KMS/HSM sérieux journalise : **qui** (identité), **quoi** (génération, utilisation,
export, destruction), **quand** (horodatage fiable — voir section 11), **résultat**
(succès/échec). Exigences :

- Journaux **protégés en écriture** (append-only, WORM ou SIEM distant) — un attaquant
  qui efface ses traces dans les logs a gagné.
- **Alertes** : échecs d'authentification répétés, export de clé (événement rare =
  alerte immédiate), utilisation hors heures ouvrées.
- **Rétention** : 1 an minimum en ligne, archivage selon les obligations (voir
  section 127).
- **Revue** : les logs que personne ne lit ne servent à rien — revue hebdo automatisée
  + revue mensuelle humaine.

# PARTIE 6 — CONFORMITÉ : L'ESSENTIEL SANS JARGON JURIDIQUE

> Avertissement : cette partie est **généraliste et informative**. Ce n'est pas un conseil
> juridique. Pour une certification, un audit ou un contentieux, faites valider par un
> juriste/auditeur qualifié.

## 121. ISO 27001:2022 : l'essentiel en bref

