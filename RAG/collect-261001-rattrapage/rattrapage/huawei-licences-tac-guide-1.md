---
id: collect-261001-rattrapage/rattrapage/huawei-licences-tac-guide-1
title: "Licences, support TAC et RMA Huawei — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["license"]
source: docs/RAG/collect-261001-rattrapage/huawei_licences_tac_guide.md
source_anchor: ""
source_lines: [1, 141]
sha256: 1c61ba00420685550beec4f140990620dccdad4a29a26adebba9df844ac155d2
---

# Licences, support TAC et RMA Huawei — Guide ultra-complet

> **Public :** chefs de service systèmes & énergies, administrateurs réseau, responsables d'exploitation.
> **Parc de référence :** AP361, AP761 (eKit), S310, AR720, USG6000, eSight, gamme eKit.
> **Langue :** français. Les modèles de tickets TAC sont en anglais (le TAC travaille en anglais).
> **Avertissement :** les procédures, délais et tarifs évoluent. Toute valeur marquée
> **« à vérifier sur le portail officiel »** doit être confirmée sur
> https://support.huawei.com avant engagement contractuel. Aucun identifiant réel
> ne figure dans ce guide : les exemples sont fictifs.
> **Date de rédaction :** septembre 2026.

---

## Table des matières

| Partie | Sections | Contenu |
|---|---|---|
| A — Le système de licences | 1 – 14 | Principes, ESN, types de licences, par produit |
| B — Activer une licence | 15 – 30 | Pas à pas, portail, CLI, erreurs |
| C — Transférer, renouveler, révoquer | 31 – 45 | RMA, renouvellements, registre |
| D — Le portail support.huawei.com | 46 – 60 | Compte, enregistrement, firmwares, HedEx |
| E — Ouvrir un ticket TAC | 61 – 75 | Sévérités, template, suivi, escalade |
| F — Collecter les infos + Garantie | 76 – 90 | diagnostic-information, logs, Hi-Care |
| G — Procédure RMA | 91 – 100 | De la panne au retour du défectueux |
| H — Cas vécus | 101 – 112 | 12 situations réelles commentées |
| I — Checklists et contacts | 113 – 117 | Avant d'acheter, réception, annuel |
| J — Spécificités eKit | 118 – 119 | Le circuit partenaire |
| K — Références | 120 – 123 | Glossaire, quiz, pour aller plus loin |
| L — Annexes opérationnelles | 124 – 143 | Templates, scripts, antisèches, FAQ |

---

# A. LE SYSTÈME DE LICENCES HUAWEI

## 1. Pourquoi la partie administrative compte autant que la technique

Un chef de service peut configurer un USG6000 les yeux fermés ; s'il oublie que la
licence IPS expire un vendredi à 18h00, c'est toute la sécurité périmétrique qui
tombe en mode dégradé pendant le week-end. La partie administrative — licences,
contrats de support, garanties — n'est pas de la paperasse : c'est de la
**disponibilité**. Trois constats de terrain :

- **80 % des « pannes » remontées au TAC un lundi matin sont des licences
  expirées** pendant le week-end (constat empirique des équipes support).
- Un RMA sans contrat Hi-Care, c'est 30 jours ouvrés de délai de réparation
  (Return For Repair) contre 4 heures avec un contrat Premier.
- Un firmware non téléchargeable parce que le contrat logiciel a expiré peut
  bloquer un projet de mise à jour de tout un parc.

Ce guide traite donc les licences, le TAC et le RMA comme des composants
d'exploitation à part entière, avec leurs procédures, leurs SLA et leurs pièges.

## 2. Le triangle : licences, support, garantie

Ces trois notions sont liées mais distinctes. Les confondre est l'erreur n°1.

| Notion | Question à laquelle elle répond | Gérée par | Exemple |
|---|---|---|---|
| **Licence** | « Ai-je le droit d'utiliser cette fonction ? » | Fichier `.dat` lié à l'ESN | Licence IPS sur USG6000 |
| **Support (TAC)** | « Qui m'aide quand ça ne marche pas ? » | Contrat Hi-Care ou garantie | Ticket S2 un dimanche |
| **Garantie** | « Qui remplace le matériel défectueux ? » | Garantie standard ou Hi-Care | RMA d'un S310 HS |

Retenez la formule : **la licence autorise, le support assiste, la garantie
remplace.** Un équipement peut être sous garantie mais sans licence UTM valide :
le TAC vous aidera pour le hardware, mais la fonction IPS restera inactive.

## 3. Vocabulaire de base : ESN, S/N, PN, LAC

| Terme | Signification | Où le trouver |
|---|---|---|
| **ESN** (Equipment Serial Number) | Numéro de série « logique » de l'équipement, utilisé par le système de licences | `display esn`, étiquette, interface web |
| **S/N** (Serial Number) | Numéro de série physique/matériel | Étiquette au dos, `display device` |
| **PN** (Part Number) | Référence commerciale du produit | Bon de commande, étiquette |
| **LAC** (License Authorization Code) | Code d'autorisation de licence (selon les gammes) | Document d'achat de licence |
| **Entitlement ID** | Identifiant de droit d'utilisation, avec mot de passe d'activation | Preuve de droit (Proof of Entitlement) fournie à l'achat |
| **Fichier `.dat`** | Fichier de licence chiffré, lié à un ESN précis | Téléchargé depuis le portail de licences |

> **Point crucial :** l'ESN et le S/N sont souvent identiques ou dérivés l'un de
> l'autre, mais le système de licences ne connaît que l'**ESN**. Toujours copier
> l'ESN affiché par `display esn`, jamais le recopier à la main depuis une
> étiquette si on peut l'éviter.

## 4. L'ESN en détail : l'ancre de tout le système

L'ESN (Equipment Serial Number) est l'identifiant auquel chaque fichier de
licence est cryptographiquement lié. Concrètement :

- Un fichier `.dat` généré pour l'ESN `2102311XYZ...` **ne s'activera jamais**
  sur un équipement d'ESN différent. Le message sera :
  `Error: The ESN of the license file does not match with the device.`
- L'ESN est stable pour la vie de l'équipement, sauf remplacement de la carte
  mère / carte de contrôle (cas RMA : voir section 32).
- Sur les châssis modulaires, il y a un ESN par carte (MPU active/standby) :
  la licence doit correspondre à la carte active.

**Récupérer l'ESN — trois méthodes :**

```
Méthode 1 (recommandée) — en CLI :
<USG6000> display esn
ESN of device: 2102311HSL8H6000123

Méthode 2 — interface web :
System > License Management : l'ESN est affiché en haut de page.

Méthode 3 — étiquette physique :
au dos / dessous de l'équipement, champ « ESN » ou « SN ».
(À n'utiliser qu'en dernier recours : risque d'erreur de recopie.)
```

**Bonnes pratiques ESN :**
- [ ] Toujours copier-coller depuis `display esn`, jamais retaper.
- [ ] Stocker les ESN dans le registre de licences (section 38).
- [ ] Vérifier l'ESN **avant** de générer le fichier sur le portail : le fichier
      est à usage unique, une erreur d'ESN = un nouveau fichier à demander.

## 5. Le fichier de licence `.dat` : anatomie

Le fichier `.dat` est un fichier chiffré généré par Huawei à partir de trois
ingrédients : votre droit d'achat (Entitlement), l'ESN cible, et le contenu de
la licence (fonctions, capacité, durée). Propriétés :

| Propriété | Détail |
|---|---|
| Format | Binaire chiffré, extension `.dat` |
| Liaison | Un fichier = un ESN = un équipement |
| Usage | **Unique** : une fois activé, il est marqué comme consommé côté portail |
| Taille | Quelques kilo-octets |
| Contenu lisible | Aucun (ne pas tenter de l'ouvrir avec un éditeur : il paraîtra corrompu, c'est normal) |
| Transport | TFTP, FTP, SFTP, USB, ou upload via l'interface web |

> **Ne jamais** renommer le fichier avec des caractères spéciaux ou des espaces,
> et ne jamais l'ouvrir/modifier avec un éditeur de texte : le moindre octet
> modifié invalide la signature et provoque `Error: Failed to activate the license`
> ou un diagnostic de « fichier corrompu ».

## 6. Types de licences (1/3) : les licences de fonctionnalités

Elles déverrouillent des fonctions logicielles. Exemples concrets sur le parc
de référence :

