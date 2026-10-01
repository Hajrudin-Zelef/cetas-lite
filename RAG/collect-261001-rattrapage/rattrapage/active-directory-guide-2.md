---
id: collect-261001-rattrapage/rattrapage/active-directory-guide-2
title: "Active Directory & GPO en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["attention", "parameters"]
source: docs/RAG/collect-261001-rattrapage/active_directory_guide.md
source_anchor: ""
source_lines: [145, 286]
sha256: 7e980636d83d0bd2fc3a9f9bfc59f175248ceab91eb7126443e0c63c481101c8
---

# Active Directory & GPO en entreprise — Guide technique ultra-complet

```
Forêt entreprise.lan
├── Arbre entreprise.lan
│   ├── Domaine racine : entreprise.lan
│   ├── Domaine enfant : prod.entreprise.lan
│   └── Domaine enfant : dev.entreprise.lan
└── Arbre filiale.lan (même forêt, espace de noms disjoint)
    └── Domaine racine : filiale.lan
```

- **Domaine** : unité de réplication. Chaque DC d'un domaine réplique la partition de ce domaine avec les autres DC du même domaine.
- **Arbre** : ensemble de domaines liés par des relations parent-enfant avec un espace de noms DNS contigu (`prod.entreprise.lan` sous `entreprise.lan`).
- **Forêt** : un ou plusieurs arbres. Le **premier domaine créé** est le domaine racine de forêt. Le schéma et la configuration sont communs à toute la forêt.

**Règle d'or du design** : **une seule forêt, un seul domaine** dans 95 % des PME/ETI. Chaque domaine supplémentaire = des DC en plus, des approbations à gérer, des GPO à dupliquer. On ne crée un domaine ou une forêt supplémentaire que pour :

- une vraie frontière de sécurité (exigence réglementaire, isolement total) ;
- un besoin d'autonomie administrative complète (rare) ;
- une entité juridique séparée avec sa propre équipe IT.

> ⚠️ On ne peut pas fusionner deux forêts facilement ni renommer un domaine sans douleur (rendom existe mais c'est risqué, Exchange ne le supporte pas). **Choisissez bien le nom de domaine dès le départ** : utilisez un sous-domaine d'un nom public que vous possédez (ex. `ad.entreprise.fr` ou `interne.entreprise.fr`), jamais un `.local` (conflits mDNS/Bonjour) ni un nom public nu sans réflexion.

---

## 4. Approbations (trusts) entre domaines et forêts

Une **approbation** (trust) permet aux utilisateurs d'un domaine d'accéder aux ressources d'un autre.

| Type | Portée | Création |
|---|---|---|
| Approbation **parent-enfant** | Automatique, transitive, bidirectionnelle | Créée à la promotion du domaine enfant |
| Approbation **racine d'arborescence** | Entre racines d'arbres d'une même forêt, transitive | Automatique |
| Approbation **de raccourci** (shortcut) | Entre deux domaines d'une même forêt pour optimiser l'authentification | Manuelle |
| Approbation **externe** | Entre domaines de forêts différentes, non transitive | Manuelle |
| Approbation **de forêt** | Entre deux forêts entières, transitive | Manuelle |

```powershell
# Lister les approbations du domaine
Get-ADTrust -Filter *

# Créer une approbation de forêt (à exécuter côté initiateur)
# Le mot de passe d'approbation doit être identique des deux côtés
New-ADTrust -SourceForest "entreprise.lan" -TargetForest "partenaire.lan" `
  -TrustType Forest -TrustDirection Bidirectional `
  -TrustPassword (Read-Host "Mot de passe d'approbation" -AsSecureString)

# Vérifier une approbation
netdom trust entreprise.lan /domain:partenaire.lan /verify
```

**Direction** : entrante, sortante ou bidirectionnelle. Une approbation **entrante** vers votre domaine signifie que l'autre domaine vous fait confiance (leurs utilisateurs accèdent à vos ressources).

**Sécurité** : utilisez le **filtrage SID** (activé par défaut sur les approbations inter-forêts) pour empêcher l'usurpation de SID. Pour les environnements sensibles, préférez l'**authentification sélective** : seuls les utilisateurs explicitement autorisés (droit `Allowed to authenticate`) peuvent traverser.

```powershell
# Activer l'authentification sélective sur une approbation
netdom trust entreprise.lan /domain:partenaire.lan /SelectiveAUTH:Yes
```

---

## 5. Catalogue global (Global Catalog)

Le **catalogue global (GC)** est un index partiel de **tous les objets de la forêt** : il contient tous les attributs de son propre domaine + un sous-ensemble d'attributs des autres domaines (ceux marqués dans le schéma comme répliqués vers le GC, ex. `sAMAccountName`, `userPrincipalName`).

**À quoi ça sert** :

1. **Ouverture de session** : pour les groupes universels, le DC doit interroger un GC pour calculer l'appartenance complète.
2. **Recherche inter-domaines** : trouver un utilisateur d'un autre domaine sans connaître son domaine.
3. **Applications** : Exchange/Outlook s'appuient historiquement sur le GC.

```powershell
# Lister les GC de la forêt
Get-ADDomainController -Filter { IsGlobalCatalog -eq $true } |
  Select-Object Name, Domain, Site

# Activer le rôle GC sur un DC
Set-ADObject "CN=NTDS Settings,CN=DC01,CN=Servers,CN=Site-Paris,CN=Sites,CN=Configuration,DC=entreprise,DC=lan" `
  -Add @{ options = 1 }
# Méthode simple via l'interface : Sites et services AD > serveur > Propriétés NTDS Settings > cocher "Catalogue global"
```

**Bonne pratique** : dans un domaine unique, faites de **tous les DC des GC** (aucun inconvénient). En multi-domaines, placez au moins un GC par site. Attention au rôle **maître d'infrastructure** : ne le colocalisez pas avec un GC sauf si tous les DC sont GC ou s'il n'y a qu'un seul domaine (voir section 33).

**Port** : le GC répond sur TCP **3268** (LDAP GC) et **3269** (LDAPS GC).

---

## 6. Partitions d'annuaire

La base AD est découpée en **partitions** (naming contexts), chacune avec sa propre portée de réplication :

| Partition | Contenu | Répliquée vers |
|---|---|---|
| **Domaine** (`DC=entreprise,DC=lan`) | Utilisateurs, groupes, OU, ordinateurs du domaine | Tous les DC du domaine |
| **Configuration** (`CN=Configuration,...`) | Topologie des sites, schéma de réplication | Tous les DC de la forêt |
| **Schéma** (`CN=Schema,...`) | Définitions des classes et attributs | Tous les DC de la forêt |
| **DNS d'application** (`DC=DomainDnsZones,...`) | Zones DNS intégrées à AD (portée domaine) | Tous les DC DNS du domaine |
| **DNS d'application** (`DC=ForestDnsZones,...`) | Zones DNS intégrées à AD (portée forêt) | Tous les DC DNS de la forêt |

```powershell
# Lister les partitions connues par un DC
Get-ADRootDSE | Select-Object -ExpandProperty namingContexts

# Voir les détails via le RootDSE
(Get-ADRootDSE).Properties["namingContexts"]
```

**Implication pratique** : une modification du schéma ou de la configuration se réplique à **toute la forêt**. C'est pour cela que le rôle de maître de schéma est si sensible (section 29).

---

## 7. Contrôleurs de domaine : rôles et architecture

Un **contrôleur de domaine (DC)** héberge une copie inscriptible (ou en lecture seule pour un RODC) de la base AD et fournit :

- **Authentification** : Kerberos (port 88) et NTLM.
- **LDAP** : annuaire sur TCP 389 / 636 (LDAPS).
- **DNS** : quasi toujours intégré à AD (recommandé).
- **Réplication** : avec les autres DC.

**Ports à ouvrir entre DC** (pare-feu) : 53 (DNS), 88 (Kerberos), 135 (RPC endpoint mapper), 389 (LDAP), 445 (SMB), 636 (LDAPS), 3268/3269 (GC), plage RPC dynamique 49152-65535 (ou restreinte, voir ci-dessous).

```powershell
# Restreindre la plage RPC de réplication (bonne pratique pare-feu)
# À faire sur chaque DC, puis redémarrer
Set-ItemProperty -Path "HKLM:\SYSTEM\CurrentControlSet\Services\NTDS\Parameters" `
  -Name "TCP/IP Port" -Value 50000 -Type DWord
Set-ItemProperty -Path "HKLM:\SYSTEM\CurrentControlSet\Services\Netlogon\Parameters" `
  -Name "DCTcpipPort" -Value 50001 -Type DWord
```

⚠️ **Ne jamais** : faire un snapshot/restore sauvage d'un DC virtualisé sans précaution (risque USN rollback, section 43), mettre un DC en pause prolongée, cloner un DC.

**Virtualisation** : depuis 2012, les DC virtualisés sont protégés par le **VM-GenerationID** (Hyper-V, VMware) à condition que l'hyperviseur et l'OS invité soient récents. Règle : un DC virtualisé se sauvegarde et se restaure **comme un DC** (état système), jamais par simple snapshot.

---

## 8. Prérequis avant de promouvoir un contrôleur de domaine

Checklist **avant** `Install-ADDSDomainController` / `Install-ADDSForest` :

