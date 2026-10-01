---
id: collect-261001-rattrapage/rattrapage/huawei-licences-tac-guide-3
title: "Licences, support TAC et RMA Huawei — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["attention", "license"]
source: docs/RAG/collect-261001-rattrapage/huawei_licences_tac_guide.md
source_anchor: ""
source_lines: [293, 456]
sha256: 1943f3b5cd1750dd2928acbbfdd2b806e62baf937bfff45e3fb80c17e6af17a8
---

# Licences, support TAC et RMA Huawei — Guide ultra-complet

```
┌─────────────┐    ┌──────────────┐    ┌───────────────┐    ┌──────────────┐    ┌────────────┐
│ 1. Récupérer │───▶│ 2. Preuve de │───▶│ 3. Générer le │───▶│ 4. Transférer │───▶│ 5. Activer │
│    l'ESN     │    │ droit (EID + │    │ fichier .dat  │    │ le .dat sur   │    │ license    │
│ display esn  │    │  mot de passe)│    │ sur le portail│    │ l'équipement  │    │ active     │
└─────────────┘    └──────────────┘    └───────────────┘    └──────────────┘    └────────────┘
                                                                                       │
                                                                              ┌────────▼────────┐
                                                                              │ 6. Vérifier     │
                                                                              │ display license │
                                                                              └─────────────────┘
```

Durée typique : 15 à 30 minutes quand on a tous les éléments ; 2 heures à 2
jours quand il manque la preuve de droit (voir cas vécu n°104).

## 16. Étape 1 — Récupérer l'ESN

Voir section 4. En pratique, sur un équipement en production :

```
<USG6000> display esn
ESN of device: 2102311HSL8H6000123
```

- Copier la valeur **exacte** (copier-coller, pas de recopie manuelle).
- La coller immédiatement dans le registre de licences (section 38) et dans
  le formulaire du portail.
- Sur un châssis à double MPU : relever l'ESN des **deux** cartes
  (`display esn slot X` selon les modèles — **à vérifier sur le portail
  officiel** pour votre châssis).

## 17. Lecture commentée d'un `display esn`

```
<USG6000> display esn
ESN of device: 2102311HSL8H6000123
│               └───────────────────── 19 caractères alphanumériques
│                 • les 7 premiers identifient en général la série/le site de fabrication
│                 • la suite est le numéro unique
└─ Un seul ESN affiché = équipement monobloc (pas de châssis multi-cartes)
```

Si la commande retourne une erreur ou un champ vide : l'équipement n'a pas
d'ESN programmé (rare, matériel très ancien ou carte remplacée sans
reprogrammation) → ouvrir un ticket TAC, seul Huawei peut régénérer un ESN.

## 18. Étape 2 — La preuve de droit (Entitlement ID + mot de passe)

À l'achat d'une licence, le partenaire (ou Huawei) fournit un document
« Proof of Entitlement » (preuve de droit) contenant :

| Champ | Exemple fictif | Usage |
|---|---|---|
| Entitlement ID | `EID-2026-8F3K2P` *(fictif)* | Identifie votre droit sur le portail |
| Activation Password | `A1b2-C3d4-E5f6` *(fictif)* | Mot de passe à saisir pour générer le fichier |
| Part Number | `L-USG-IPS-1Y` *(fictif)* | Référence de la licence achetée |
| Quantité / durée | 1 / 1 an | Ce que vous avez payé |

**Où trouver ce document ?**
1. Dans les pièces jointes du bon de livraison / de la facture du partenaire.
2. Dans l'espace client du partenaire (portail revendeur).
3. En le redemandant au commercial du partenaire (prévoir 24-48h).

> **Piège classique :** le commercial envoie la facture mais pas la preuve de
> droit. Sans Entitlement ID + mot de passe, impossible de générer le `.dat`.
> Exigez ce document **à la commande** (checklist section 113).

## 19. Étape 3 — Le portail de téléchargement des licences

Le portail officiel de gestion des licences électroniques est accessible via
le support Huawei (adresse constatée : `https://app.huawei.com/isdp` — **à
vérifier sur le portail officiel**, l'URL peut évoluer ; en cas de doute,
partez de https://support.huawei.com > rubrique Licences).

Prérequis :
- Un compte support.huawei.com (voir section 47).
- L'Entitlement ID et le mot de passe d'activation.
- L'ESN exact de l'équipement cible.

## 20. Génération unitaire : le formulaire ESN

Procédure type (les intitulés exacts peuvent varier) :

1. Connectez-vous au portail de licences.
2. Saisissez l'**Activation Password** de votre preuve de droit, cliquez sur
   **Activate**.
3. Le portail affiche les informations du droit (part number, quantité, durée,
   fonctions). **Vérifiez-les** : c'est votre dernière chance de détecter une
   erreur de commande (mauvais produit, mauvaise durée).
4. Saisissez l'**ESN** (copié depuis `display esn`).
5. Validez : le portail génère le fichier `.dat`.
6. Téléchargez-le immédiatement et archivez-le (voir section 22).

> **Rappel :** le fichier est à usage unique et lié à l'ESN saisi. Une faute
> de frappe sur l'ESN = fichier inutilisable = nouvelle demande à formuler.

## 21. Génération en masse : le Batch Import

Pour un parc (ex. : 40 AP gérés = licences contrôleur, ou plusieurs USG) :

1. Sur le portail, téléchargez le **template** (modèle de fichier) proposé
   pour l'import en masse.
2. Remplissez une ligne par équipement : **Password** (mot de passe
   d'activation) + **ESN**.
3. Cliquez sur **Batch Import** et téléversez le fichier rempli.
4. Le portail génère les `.dat` en lot ; téléchargez l'archive.

**Conseils :**
- Préparez le fichier depuis votre registre de licences (section 38) : une
  ligne = un équipement, zéro recopie manuelle.
- Vérifiez le rapport d'import ligne par ligne avant de quitter la page.

## 22. Télécharger le `.dat` : les vérifications immédiates

Dès le téléchargement :

- [ ] Le nom du fichier mentionne-t-il le produit / l'ESN ? (selon le portail)
- [ ] La taille est-elle plausible (quelques Ko, pas 0 octet) ?
- [ ] Archivez-le dans `licences/<année>/<équipement>_<ESN>_<fonction>.dat`
      **et** dans la sauvegarde hors site.
- [ ] Notez dans le registre : date de génération, date d'expiration,
      Entitlement ID d'origine.

Ne supprimez jamais l'original téléchargé : en cas de RMA, il servira de
preuve pour demander le transfert (section 32).

## 23. Étape 4 — Transférer le fichier sur l'équipement

Quatre méthodes, de la plus simple à la plus « terrain » :

| Méthode | Commande / action | Quand l'utiliser |
|---|---|---|
| **Web** | System > License Management > Upload | Accès HTTPS disponible, petit fichier |
| **TFTP** | `tftp <serveur> get licence.dat` | Réseau local, serveur TFTP dispo |
| **FTP/SFTP** | `ftp` / `sftp` puis `get` | Transfert authentifié |
| **USB** | Copier sur clé, `copy usb:/licence.dat flash:/` | Site isolé, pas de réseau |

Vérifiez la présence après transfert :
```
<USG6000> dir
...
licence.dat   (taille cohérente avec l'original)
```

## 24. Étape 5 — Activer : `license active` (CLI)

```
<USG6000> system-view
[USG6000] license active licence.dat
Info: License file activated successfully.
```

Points d'attention :
- Le nom du fichier doit être **exact** (sensible à la casse selon les versions).
- Ne lancez **jamais** l'activation pendant un basculement actif/standby :
  en cas de switchover durant l'activation, le fichier peut échouer à se
  « dispatcher » (erreur documentée : *Failed to dispatch the license file*).
  Si cela arrive : attendre la fin du basculement, puis réactiver.
- Sur les châssis double MPU : vérifier l'espace disque de la MPU standby
  avant d'activer (erreur documentée : disque plein côté standby).

## 25. Étape 5 bis — Activer via l'interface web

