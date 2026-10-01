---
id: collect-261001-rattrapage/rattrapage/glpi-guide-8
title: "Guide GLPI ultra-complet — Ticketing, gestion de parc et pilotage SAV"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "agents"]
source: docs/RAG/collect-261001-rattrapage/glpi_guide.md
source_anchor: ""
source_lines: [1121, 1314]
sha256: 4d610647edfdc21227c6987062601d512c60043bcd2ab2a0e5d309e27388087c
---

# Guide GLPI ultra-complet — Ticketing, gestion de parc et pilotage SAV

Même logique : marque, modèle, série, liaison (« connecté à » tel ordinateur),
consommables compatibles (pour les imprimantes : **quelles cartouches ?** —
lien direct vers le stock, section 30).

**Cas des imprimantes réseau :** renseignez l'adresse IP, la file d'impression,
le pilote. L'agent FusionInventory peut les découvrir en SNMP (section 35) :
compteur de pages remonté automatiquement — **données de facturation sans relevé
manuel**.

---

## 28. Fiches copieurs : le métier (marque, modèle, série, compteur, contrat)

C'est **le cœur du métier** : chaque copieur client est une fiche « Imprimante »
(enrichie) ou un type d'objet dédié. Voici la fiche type à normaliser pour tout
le parc :

### Fiche copieur normalisée

| Rubrique | Champs |
|---|---|
| **Identification** | Marque (Kyocera, Canon, Ricoh...), modèle (TASKalfa 4054ci), n° de série, n° d'inventaire, adresse MAC, IP |
| **Localisation** | Client (entité), site, bâtiment, étage/bureau, contact sur place |
| **Compteurs** | N&B, couleur, total, date du dernier relevé (champs personnalisés) |
| **Contrat** | Contrat de maintenance lié (section 37), type (forfait, coût/copie), date de fin, SLA applicable |
| **Consommables** | Toners compatibles (liens vers stock), kit de fusion, tambour — avec seuils d'alerte |
| **Historique** | Tous les tickets liés, interventions, pièces remplacées |
| **Documents** | Contrat signé (PDF), PV d'installation, manuel |

### Relevés de compteurs : processus

1. **Mensuel** (ticket récurrent, section 19) ou **automatique** (SNMP, section 35).
2. Saisie dans les champs personnalisés de la fiche + suivi dans le ticket.
3. **Alerte** si écart anormal (ex. +50 % vs mois précédent = fuite ? erreur de
   saisie ? à contrôler).
4. Les compteurs alimentent la **facturation au coût/copie** (export CSV pour
   la comptabilité).

> **Lien avec le guide Kyocera** : les codes d'accès maintenance (ex. 10871087),
> les U-codes et les C-codes documentés dans le guide copieurs doivent exister
> en articles de la base de connaissances (section 24) et être liés aux fiches
> copieurs concernées.

---

## 29. Périphériques réseau : switchs, routeurs, bornes Wi-Fi, onduleurs

- **Matériels réseau** (`Parc > Matériels réseau`) : switchs, routeurs, pare-feux.
  Renseignez les ports, VLAN, firmware, liaisons (« ce switch alimente ces PC »).
  Découverte SNMP via FusionInventory (section 35).
- **Onduleurs** : GLPI n'a pas de type natif « onduleur » en 10.x — utilisez
  `Parc > Autres` ou créez un type via le plugin **Champs personnalisés** /
  **Objects**. Champs métier : puissance (kVA), autonomie, date des batteries,
  prochaine maintenance (lien avec le guide onduleurs !).
- **Liaisons** : branchez chaque équipement critique sur son onduleur dans la
  CMDB (« alimenté par »). En cas de coupure, vous savez **exactement** ce qui
  est protégé ou non.

---

## 30. Cartouches et consommables : gérer le stock de toner

`Parc > Cartouches` (modèles) + `Parc > Consommables` : le **magasin** de GLPI.

### Organisation du stock

1. Créez les **modèles de cartouches** : référence constructeur (ex. TK-8345K),
   couleur, capacité (pages), seuil d'alerte.
2. Liez chaque modèle aux **imprimantes/copieurs compatibles**.
3. **Entrées** : à chaque livraison fournisseur (bon de livraison → quantité).
4. **Sorties** : le technicien « donne » une cartouche depuis le ticket ou la
   fiche imprimante → le stock décrémente, le coût s'impute au client/contrat.

### Alertes et réassort

- Seuil d'alerte par modèle (ex. < 3 toners noirs TK-8345) → notification au
  magasinier + au chef de service.
- **Rapport mensuel** : consommation par client, par modèle → ajustement des
  commandes et détection des anomalies (un client qui consomme 3× plus = fuite ?
  usage détourné ?).
- Inventaire physique **trimestriel** : comptage réel vs GLPI, régularisation
  avec motif (écart justifié dans l'historique).

> **Le saviez-vous ?** Lier les cartouches aux tickets d'intervention permet de
> calculer le **coût réel par client** (pièces + consommables + main-d'œuvre) —
> indispensable pour renégocier les contrats déficitaires.

---

## 31. Logiciels, licences et versions

`Parc > Logiciels` : catalogue (nom, éditeur, versions), puis **licences**
(type : OEM, volume, SaaS ; nombre de postes ; dates d'expiration ; clé).

- L'agent FusionInventory remonte les logiciels installés → rapprochement
  automatique avec les licences → **alertes de sous/sur-licence**.
- Dates d'expiration (antivirus, certificats, abonnements SaaS) : notifications
  à J-60 / J-30 / J-7.
- Dictionnaire des logiciels (`Configuration > Dictionnaires`) : fusionne
  « Acrobat Reader », « Adobe Acrobat Reader DC »... en une seule entrée propre.

---

## 32. Réservations de matériel

`Assistance > Réservations` : prêt de matériel (vidéoprojecteur, PC portable de
dépannage, copieur de remplacement). Le déclarant réserve sur un créneau, le
technicien valide, le retour est tracé. Fini le « qui a pris le portable n°12 ? ».

> **Usage SAV :** gérez ainsi votre **parc de prêt** (copieurs de remplacement
> pendant une réparation longue) : chaque prêt = réservation liée au ticket
> d'intervention.

---

# PARTIE V — INVENTAIRE AUTOMATIQUE

## 33. FusionInventory : présentation et architecture

**FusionInventory** = le plugin GLPI + l'**agent** installé sur les postes (ou
serveur dédié pour le SNMP). Il remonte : matériel, logiciels, réseau, imprimantes.

```
Poste Windows/Linux ──agent──▶ https://glpi/plugins/fusioninventory/
Serveur de découverte ──SNMP──▶ switchs, imprimantes, copieurs ──▶ GLPI
```

> **Note :** le plugin historique « FusionInventory for GLPI » évolue ; en
> GLPI 10.x, l'inventaire natif (`Administration > Inventaire`, agent GLPI)
> reprend l'essentiel. Vérifiez la compatibilité exacte plugin ↔ version GLPI
> **avant** d'installer. Les principes ci-dessous restent valables.

---

## 34. Installation de l'agent FusionInventory

### Windows (GPO ou script)

```powershell
# Installation silencieuse de l'agent GLPI/FusionInventory
msiexec /i "glpi-agent-1.x.msi" /qn `
  SERVER="https://glpi.entreprise.lan" `
  RUNNOW=1
```

Fichier de config `C:\Program Files\GLPI-Agent\etc\glpi-agent.cfg` :

```ini
server = https://glpi.entreprise.lan
tag = AGENCE-NORD
delaytime = 24
```

### Linux (Debian/Ubuntu)

```bash
sudo apt install -y glpi-agent   # ou fusioninventory-agent selon dépôt
sudo tee /etc/glpi-agent/conf.d/serveur.cfg <<'EOF'
server = https://glpi.entreprise.lan
tag = SIEGE
EOF
sudo systemctl enable --now glpi-agent
```

### Vérification côté GLPI

`Administration > Inventaire > Agents` : l'agent apparaît après son premier
contact. Forcez un inventaire : `glpi-agent --debug` côté poste, ou tâche
« Forcer l'inventaire » côté serveur.

---

## 35. Inventaire automatique : réseau et SNMP

### Découverte réseau

`Plugins > FusionInventory > Découverte réseau` : définissez des plages IP
(ex. `192.168.10.0/24`), une tâche planifiée lance le scan via l'agent.
Résultat : liste des équipements répondant (ping/ARP), à importer ou ignorer.

### Inventaire SNMP (imprimantes, copieurs, switchs)

1. `Configuration > Authentification SNMP` : créez la communauté
   (`public` → **changez-la !**, ou SNMPv3 avec authPriv).
2. Tâche « Inventaire réseau » : plage IP + authentification SNMP.
3. L'agent interroge chaque équipement : modèle, série, **compteurs de pages**,
   niveaux de toner, firmware.

> **Jackpot métier :** les compteurs des copieurs réseau remontent **sans
> déplacement**. Programmez l'inventaire SNMP la nuit du 30 de chaque mois →
> relevés de facturation automatiques (export API, section 53).

### Bonnes pratiques SNMP

