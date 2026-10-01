---
id: collect-261001-rattrapage/rattrapage/maintenance-windows-guide-19
title: "Maintenance et exploitation Windows en entreprise"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-26-09"]
keywords: []
source: docs/RAG/collect-261001-rattrapage/maintenance_windows_guide.md
source_anchor: ""
source_lines: [3539, 3732]
sha256: 1bc33c3353e06fd44260597b6463d6be9509cf75e761348eb584bc20ec9db4b0
---

# Maintenance et exploitation Windows en entreprise

`Rapports-Quotidiens.ps1` appelle chaque script, consolide les CSV dans
`C:\Rapports\AAAA-MM-JJ\`, envoie un e-mail de synthèse avec les alertes en
objet. **Testez chaque script à la main** avant de le planifier, et surveillez
`LastTaskResult` (§47) : un rapport qui ne tourne plus est pire que pas de
rapport (faux sentiment de sécurité).

---

## 122. PRA/PCA : RTO, RPO et scénarios

| Notion | Définition | Exemple |
|---|---|---|
| **RTO** (Recovery Time Objective) | Durée max d'interruption acceptable | 4 h pour l'ERP |
| **RPO** (Recovery Point Objective) | Perte de données max acceptable | 1 h pour la comptabilité |
| **PRA** (Plan de Reprise d'Activité) | Reprise après sinistre majeur | Incendie salle serveur |
| **PCA** (Plan de Continuité d'Activité) | Maintien pendant la crise | Mode dégradé |

**Scénarios à couvrir** (a minima) :

1. Perte d'un serveur non critique (restauration simple).
2. Perte d'un contrôleur de domaine (réinstallation + réplication, ou
   restauration d'état système §126).
3. Perte de l'hyperviseur / du cluster (bascule, reconstruction).
4. Ransomware (restauration depuis sauvegardes **hors ligne**, §123).
5. Perte de la salle (site de secours, §123).

**Règle du chef de service** : RTO/RPO se **négocient avec la direction et les
métiers**, pas dans le bureau de l'informaticien. Faites signer le tableau
RTO/RPO par application : c'est votre contrat en cas de crise, et ça dimensionne
le budget sauvegarde.

---

## 123. Stratégie de sauvegarde Windows

| Donnée | Méthode | Fréquence type |
|---|---|---|
| Volumes système serveurs | Image complète (bare metal) | Hebdo + avant changement |
| Données (partages, bases) | Sauvegarde fichiers / applicative | Quotidienne |
| AD (contrôleurs) | État du système | Quotidienne (un DC suffit, en rotation) |
| Postes | Dossiers redirigés / OneDrive / image pour postes critiques | Continue/hebdo |

**Règle 3-2-1** : 3 copies, 2 supports différents, 1 hors site (ou hors ligne).
Contre les ransomwares : au moins une copie **immuable ou déconnectée**
(disque externe rotatif, bande, stockage objet avec verrouillage).

**À tester** : une sauvegarde non testée = pas de sauvegarde (§127).
**À documenter** : quoi, où, fréquence, rétention, chiffrement, responsable,
procédure de restauration pas à pas (§128).

---

## 124. Windows Server Backup en pratique

Windows Server Backup (fonctionnalité à ajouter) suffit pour les petits parcs :

```powershell
# Installer la fonctionnalité (+ outils en ligne de commande)
Install-WindowsFeature Windows-Server-Backup -IncludeManagementTools

# Sauvegarde complète (tous les volumes critiques) vers un disque dédié
wbadmin start backup -backupTarget:E: -include:C:,D: -allCritical -quiet

# Sauvegarde de l'état du système seul (utile pour un DC, voir §126)
wbadmin start systemstatebackup -backupTarget:E: -quiet

# Planifier : sauvegarde quotidienne à 22:00
wbadmin enable backup -addtarget:E: -schedule:22:00 -include:C:,D: -allCritical -quiet
```

Vérifications d'exploitation :

```cmd
:: Derniers résultats de sauvegarde
wbadmin get versions
:: Détail d'une version
wbadmin get items -version:09/26/2026-22:00
```

**Limites** : pas de déduplication avancée, pas de gestion centralisée
multi-serveurs, restauration granulaire limitée → au-delà d'une dizaine de
serveurs, passez à une vraie solution (Veeam, etc., §138).

---

## 125. VSS : clichés instantanés de volumes

**VSS** (Volume Shadow Copy Service) fige un volume pour sauvegarder des fichiers
ouverts (bases, boîtes partagées) de façon cohérente.

```cmd
:: Lister les clichés existants
vssadmin list shadows
:: Espace utilisé par les clichés
vssadmin list shadowstorage
:: Redimensionner l'espace réservé (ex. 10 % du volume D:)
vssadmin resize shadowstorage /for=D: /on=D: /maxsize=10%
```

**Clichés instantanés pour les partages** (versions précédentes) : activez-les sur
les volumes de données (2 clichés/jour ouvrés typiquement) — les utilisateurs
restaurent eux-mêmes leurs fichiers écrasés (clic droit > Versions précédentes),
ce qui divise les tickets « j'ai écrasé mon fichier ».

Écrivains VSS en erreur = sauvegardes incohérentes :

```cmd
vssadmin list writers
:: Tout écrivain durablement en "Failed" / "Timed out" -> redémarrer le service
:: associé, voire le serveur en fenêtre de maintenance
```

---

## 126. Sauvegarde des contrôleurs de domaine (état du système)

Un DC se sauvegarde via l'**état du système** (AD, SYSVOL, registre, COM+) :

```cmd
wbadmin start systemstatebackup -backupTarget:E: -quiet
```

**Restauration** (deux cas) :

1. **Il reste des DC sains** : ne restaurez pas — **réinstallez** un Windows,
   promovez-le en DC, laissez la réplication repeupler. C'est plus propre qu'une
   restauration (pas de risque d'USN rollback).
2. **Tous les DC sont perdus** : restauration d'état système en mode
   **restauration des services d'annuaire (DSRM)** sur un DC, avec le mot de
   passe DSRM (à conserver sous enveloppe scellée, §128).

> ⚠️ **USN rollback** : ne restaurez jamais un snapshot/clone d'un DC sans
> procédure (restauration non autoritaire) — vous corrompriez la réplication AD
> de tout le domaine. En cas de doute : réinstallation + promotion.

**Bonnes pratiques** : sauvegardez l'état système d'**au moins un DC par jour**
(en rotation), conservez le mot de passe DSRM hors du DC lui-même, et documentez
la procédure de reconstruction complète du domaine (elle servira le jour où…).

---

## 127. Tests de restauration : procédure

**Trimestriel** (partiel) et **annuel** (complet, §9). Procédure type :

1. Choisir le scénario (ex. : restauration du serveur de fichiers sur un hôte
   isolé).
2. Restaurer **sans toucher à la production** (réseau isolé ou VLAN de test).
3. Vérifier : démarrage, services, données (échantillon de fichiers avec
   contrôle d'intégrité), applications (ouverture d'un dossier, impression test).
4. Mesurer : temps réel vs RTO, perte de données vs RPO.
5. **Compte rendu écrit** : ce qui a marché, ce qui a coincé, actions
   correctives, mise à jour du PRA et des procédures.

```powershell
# Exemple de contrôle d'intégrité après restauration d'un partage :
# comparer un échantillon de fichiers (nom + taille + hash)
$ref = Import-Csv C:\PRA\referentiel-fichiers.csv   # généré avant sinistre
foreach ($f in $ref | Select-Object -First 50) {
    $ok = (Test-Path $f.Chemin) -and
          ((Get-FileHash $f.Chemin -Algorithm SHA256).Hash -eq $f.Hash)
    [pscustomobject]@{ Fichier = $f.Chemin; Integre = $ok }
} | Format-Table -AutoSize
```

> Un test qui échoue n'est pas un échec : c'est le test qui a servi. Ce qui est
> un échec, c'est de découvrir le problème le jour du vrai sinistre.

---

## 128. Documentation d'exploitation : le DTI

Le **DTI** (Dossier Technique d'Intervention / d'exploitation) est la bible de
l'équipe. Contenu minimal par système :

| Rubrique | Contenu |
|---|---|
| Fiche d'identité | Rôle, criticité, RTO/RPO, responsable, contrats |
| Architecture | Schéma, IP, VLAN, dépendances (AD, DNS, stockage, onduleur) |
| Installation | Procédure de reconstruction pas à pas (testée !) |
| Exploitation | Checklists, fenêtres, procédures courantes (ce guide en est la base) |
| Supervision | Seuils, alertes, procédures associées |
| Sauvegarde | Quoi/où/fréquence/rétention + procédure de restauration |
| Accès | Comptes (renvoi au coffre-fort), accès physiques, contacts |
| Historique | Incidents majeurs, changements notables |

**Règles** : un DTI par système critique, versionné, relu annuellement, accessible
**hors du système lui-même** (le DTI du serveur de fichiers ne doit pas être
uniquement sur le serveur de fichiers). En crise, on n'a pas le temps de chercher.

---

## 129. Procédures : standard de rédaction

