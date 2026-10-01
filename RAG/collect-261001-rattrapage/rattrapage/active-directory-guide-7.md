---
id: collect-261001-rattrapage/rattrapage/active-directory-guide-7
title: "Active Directory & GPO en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-rattrapage/active_directory_guide.md
source_anchor: ""
source_lines: [985, 1163]
sha256: 2ee511255c14707697e2541c53fc6bac2cabb262e8d5ef017ff8be98767795e8
---

# Active Directory & GPO en entreprise — Guide technique ultra-complet

**Bonnes pratiques** : placez-le sur un DC **bien répliqué et disponible** ; surveillez l'événement 16650 ; ne le saisissez (seize) qu'en dernier recours car un ancien maître RID qui reviendrait en ligne pourrait distribuer des pools en double → **SID en double** (catastrophe).

---

## 32. Émulateur PDC (PDC Emulator)

Le rôle le plus **critique** au quotidien :

- **Référence de temps** du domaine (hiérarchie NTP : PDC de la forêt → DC → clients).
- **Verrouillages de compte** : c'est lui qui comptabilise les échecs et verrouille.
- **Changements de mot de passe urgents** : répliqués en priorité vers lui.
- **Édition des GPO** : l'éditeur (GPMC) cible le PDC par défaut pour éviter les conflits.
- **Scripts/objets** sensibles au « dernier écrivain ».

```powershell
(Get-ADDomain).PDCEmulator

# Configurer la source de temps du PDC (ex. serveurs NTP externes)
w32tm /config /manualpeerlist:"0.fr.pool.ntp.org 1.fr.pool.ntp.org" /syncfromflags:manual /reliable:yes /update
Restart-Service w32time
w32tm /query /status
w32tm /query /peers

# Sur les autres DC : synchronisation depuis la hiérarchie du domaine (par défaut)
w32tm /config /syncfromflags:domhier /update
```

**Bonnes pratiques** : PDC sur le DC le plus **robuste et central** (souvent le premier DC du site principal) ; **ne le placez jamais sur un RODC** (impossible de toute façon) ; surveillez-le en priorité ; prévoyez son transfert rapide en cas de maintenance.

---

## 33. Maître d'infrastructure (Infrastructure Master)

**Rôle** : maintient les références inter-domaines (quand un objet d'un autre domaine est membre d'un groupe local, il met à jour le DN fantôme/phantom).

**Règle de placement** (multi-domaines) : **ne le mettez PAS sur un DC qui est aussi catalogue global**, sauf si :

- tous les DC du domaine sont des GC, ou
- il n'y a qu'un seul domaine dans la forêt (cas le plus courant : la règle ne s'applique pas).

```powershell
(Get-ADDomain).InfrastructureMaster
```

En domaine unique, oubliez cette contrainte : mettez-le où vous voulez.

---

## 34. Transférer un rôle FSMO (Move-ADDirectoryServerOperationMasterRole)

Le **transfert** est l'opération **propre** : l'ancien détenteur est en ligne et consent.

```powershell
# Transférer les 5 rôles vers DC02 (noms ou numéros 0-4)
Move-ADDirectoryServerOperationMasterRole -Identity "DC02" `
  -OperationMasterRole SchemaMaster, DomainNamingMaster, RIDMaster, PDCEmulator, InfrastructureMaster `
  -Force

# Avec les numéros (0=Schéma, 1=Nommage, 2=RID, 3=PDC, 4=Infrastructure)
Move-ADDirectoryServerOperationMasterRole -Identity "DC02" -OperationMasterRole 0,1,2,3,4 -Force

# Transférer un seul rôle (ex. PDC avant maintenance du DC01)
Move-ADDirectoryServerOperationMasterRole -Identity "DC02" -OperationMasterRole PDCEmulator -Force
```

**Quand transférer** : maintenance planifiée du détenteur, rééquilibrage, décommission d'un DC (obligatoire avant rétrogradation si le DC détient des rôles — sinon la rétrogradation les transfère automatiquement, mais mieux vaut le faire explicitement avant).

---

## 35. Saisir un rôle FSMO (seize via ntdsutil) — procédure d'urgence

La **saisie** (seize) s'utilise quand le détenteur est **mort et ne reviendra jamais**. C'est une opération **destructrice** : l'ancien détenteur ne doit **jamais** être reconnecté au réseau (risque de RID en double, de schéma divergent).

```powershell
# Procédure ntdsutil (le seul moyen de "seize" ; la cmdlet PowerShell ne fait que des transferts)
ntdsutil
roles
connections
connect to server DC02
quit
seize schema master
seize naming master
seize rid master
seize pdc
seize infrastructure master
quit
quit
```

**Après une saisie** :

1. Vérifiez : `netdom query fsmo`.
2. **Ne rallumez jamais** l'ancien DC sur le réseau. Réinstallez-le from scratch s'il doit revenir.
3. Nettoyez ses métadonnées (section 12).
4. Surveillez la réplication pendant 48 h.

> ⚠️ Erreur classique : saisir un rôle alors que le détenteur est juste **temporairement injoignable** (panne réseau). Résultat : deux maîtres potentiels → corruption. La saisie = détenteur **définitivement perdu**.

---

## 36. Bonnes pratiques de placement des rôles FSMO

**Petite structure (1-2 DC, un domaine)** : laissez les 5 rôles sur le premier DC. Simplicité > micro-optimisation.

**Structure moyenne (3+ DC, un site)** :

| Rôle | Placement recommandé |
|---|---|
| Schéma + Nommage (forêt) | DC principal du site central |
| RID + PDC + Infrastructure (domaine) | DC principal (souvent le même) |

**Multi-sites** : les 5 rôles sur un DC du site central (le mieux connecté), **jamais** sur un RODC ni sur un DC de site distant fragile.

**Règles** :

- Documentez quel DC détient quoi (dans votre wiki + étiquette physique/virtuelle).
- Avant toute maintenance du détenteur : transférez le PDC a minima.
- Ne placez jamais un rôle FSMO sur un DC que vous prévoyez de décommissionner.
- Surveillez avec un script quotidien (section 73).

---

## 37. Sites et services : concepts

Un **site AD** représente un **emplacement réseau bien connecté** (LAN ou liaisons rapides), pas forcément un site géographique. AD s'en sert pour :

- **Optimiser la réplication** : réplication fréquente intra-site (toutes les ~15 s après notification), planifiée inter-sites (défaut 180 min).
- **Diriger les clients** vers le DC le plus proche (via les enregistrements DNS `_ldap._tcp.<Site>._sites`).
- **Cibler les GPO** et DFS.

**Vocabulaire** :

| Terme | Définition |
|---|---|
| **Site** | Regroupement de sous-réseaux bien connectés. |
| **Sous-réseau** | Objet `10.10.20.0/24` associé à un site. **Tout sous-réseau client doit être déclaré**, sinon les clients tombent dans `Default-First-Site-Name`. |
| **Liaison de site (Site Link)** | Chemin logique entre sites : coût, fréquence, planification. |
| **Bridgehead (tête de pont)** | DC désigné pour la réplication inter-sites (choisi par le KCC, ou manuel). |
| **ISTG** | Serveur générateur de topologie inter-sites (un par site, rôle KCC). |

---

## 38. Créer et configurer des sites, sous-réseaux

```powershell
# Créer les sites
New-ADReplicationSite -Name "Site-Paris"
New-ADReplicationSite -Name "Site-Lyon"

# Déclarer les sous-réseaux (TOUS les réseaux : clients, serveurs, DMZ, Wi-Fi)
New-ADReplicationSubnet -Name "10.10.10.0/24" -Site "Site-Paris" -Description "LAN Paris"
New-ADReplicationSubnet -Name "10.10.20.0/24" -Site "Site-Lyon"  -Description "LAN Lyon"
New-ADReplicationSubnet -Name "10.10.30.0/24" -Site "Site-Paris" -Description "Wi-Fi Paris"
New-ADReplicationSubnet -Name "192.168.99.0/24" -Site "Site-Paris" -Description "DMZ"

# Déplacer un serveur (DC) dans son site : via la console
# "Sites et services AD" > Servers > clic droit > Déplacer
# Ou en PowerShell (objet serveur) :
Get-ADDomainController "DC-Lyon" | Move-ADObject -TargetPath "CN=Servers,CN=Site-Lyon,CN=Sites,CN=Configuration,DC=ad,DC=entreprise,DC=fr"

# Renommer le site par défaut si vous ne l'utilisez plus
Rename-ADObject -Identity "CN=Default-First-Site-Name,CN=Sites,CN=Configuration,DC=ad,DC=entreprise,DC=fr" -NewName "Site-Paris"
```

**Vérification côté client** : `nltest /dsgetsite` doit retourner le bon site. Sinon : sous-réseau manquant ou mal déclaré.

---

## 39. Liaisons de sites : coût, planification, bridgehead

```powershell
# Créer une liaison de site
New-ADReplicationSiteLink -Name "Lien-Paris-Lyon" `
  -SitesIncluded "Site-Paris","Site-Lyon" `
  -Cost 100 -ReplicationFrequencyInMinutes 60

# Modifier le coût (plus le coût est bas, plus la liaison est préférée)
Set-ADReplicationSiteLink -Identity "Lien-Paris-Lyon" -Replace @{ cost = 50 }

