---
id: collect-261001-rattrapage/rattrapage/dhcp-windows-guide-2
title: "Guide technique ultra-complet : DHCP sous Windows Server en entreprise"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "attribution"]
source: docs/RAG/collect-261001-rattrapage/dhcp_windows_guide.md
source_anchor: ""
source_lines: [128, 279]
sha256: ee0407f79f14d768685af99febda6be2cad3d044dbca43e7b655000a88cefc17
---

# Guide technique ultra-complet : DHCP sous Windows Server en entreprise

```text
                    ┌─────────────────────────┐
                    │   Active Directory      │
                    │  + DNS (dynamique)      │
                    └────────────┬────────────┘
                                 │ autorisation
              ┌──────────────────┴──────────────────┐
              │                                     │
   ┌──────────▼──────────┐              ┌───────────▼──────────┐
   │  DHCP primaire      │◄────────────►│  DHCP secondaire     │
   │  SRV-DHCP-01        │  basculement │  SRV-DHCP-02         │
   │  192.0.2.10         │  (failover)  │  192.0.2.11          │
   └──────────┬──────────┘              └───────────┬──────────┘
              │                                     │
     ┌────────┴────────┐                   ┌────────┴────────┐
     │ Étendues :      │                   │ Réplicas        │
     │ - LAN 192.0.2.0 │                   │ (hot standby    │
     │ - WiFi 192.0.3.0│                   │  ou load bal.)  │
     │ - Voix 192.0.4.0│                   └─────────────────┘
     │ - Guest 192.0.5.0
     └────────┬────────┘
              │ agent de relais (IP helper)
              │ sur chaque VLAN / routeur
     ┌────────┴─────────────────────────────────┐
     │ VLAN 10 LAN │ VLAN 20 WiFi │ VLAN 30 Voix │
     └──────────────────────────────────────────┘
```

Principes :

1. **Toujours 2 serveurs DHCP** en basculement (section 59) — jamais un seul en production.
2. **Une étendue par VLAN/sous-réseau**.
3. **Agent de relais** sur chaque routeur/VLAN (section 64) — le broadcast DHCP ne traverse pas les routeurs.
4. **DNS dynamique** activé pour que chaque bail crée/met à jour l'enregistrement DNS (section 69).

## 6. Glossaire éclair des termes DHCP (version courte — le glossaire complet est en annexe A3)

- **Étendue (scope)** : plage d'adresses IP qu'un serveur peut distribuer sur un sous-réseau.
- **Bail (lease)** : attribution temporaire d'une IP à un client, avec une durée.
- **Réservation** : bail permanent lié à une adresse MAC.
- **Exclusion** : plage d'adresses **jamais** distribuée (équipements en IP statique).
- **Option** : paramètre réseau transmis avec le bail (passerelle, DNS...).
- **Superscope** : regroupement de plusieurs étendues (multinet sur un même segment).
- **Basculement (failover)** : réplication entre 2 serveurs DHCP (hot standby ou load balance).
- **Agent de relais** : service qui relaie les broadcasts DHCP entre VLANs.
- **Rogue DHCP** : serveur DHCP illégitime/pirate sur le réseau.

## 7. Installation du rôle DHCP — via le Gestionnaire de serveur (GUI)

1. **Gestionnaire de serveur** → **Gérer** → **Ajouter des rôles et fonctionnalités**.
2. Type d'installation : **Installation basée sur un rôle ou une fonctionnalité**.
3. Sélectionner le serveur (membre du domaine, IP statique déjà configurée).
4. Cocher **Serveur DHCP** → **Ajouter les fonctionnalités** requises (outils d'administration).
5. Suivant jusqu'à **Installer**. Fermer l'assistant.
6. **Post-installation obligatoire** : cliquer sur le drapeau jaune → **Terminer la configuration DHCP** → crée les groupes de sécurité (`DHCP Administrators`, `DHCP Users`) et **autorise** le serveur dans AD (demande des droits d'administrateur d'entreprise).

> ⚠️ Sans l'étape 6, le service DHCP **ne distribue aucune adresse** dans un domaine AD (protection anti-rogue, voir section 9).

## 8. Installation du rôle DHCP — via PowerShell (méthode recommandée)

```powershell
# 1. Installer le rôle + outils d'administration (RSAT DHCP)
Install-WindowsFeature -Name DHCP -IncludeManagementTools

# Vérifier l'installation
Get-WindowsFeature -Name DHCP | Select-Object Name, Installed, InstallState

# 2. Redémarrer le service si besoin
Restart-Service -Name DHCPServer

# 3. Vérifier que le service est en automatique et démarré
Get-Service -Name DHCPServer | Select-Object Name, Status, StartType

# 4. Autoriser le serveur dans Active Directory (droits Enterprise Admin requis)
Add-DhcpServerInDC -DnsName "srv-dhcp-01.contoso.local" -IPAddress 192.0.2.10

# 5. Vérifier l'autorisation
Get-DhcpServerInDC

# 6. Créer les groupes de sécurité (fait normalement par la post-installation GUI)
# Si install 100% PowerShell, les créer à la main :
# net localgroup "DHCP Administrators" /add
# net localgroup "DHCP Users" /add
```

Installation sur **Server Core** (sans GUI) — identique, c'est même le cas d'usage idéal :

```powershell
# Sur Server Core, tout se fait en PowerShell : le module DhcpServer est complet
Import-Module DhcpServer
Get-Command -Module DhcpServer | Measure-Object   # ~120 cmdlets disponibles
```

Désinstallation propre (si besoin) :

```powershell
# D'abord retirer l'autorisation AD, puis le rôle
Remove-DhcpServerInDC -DnsName "srv-dhcp-01.contoso.local" -IPAddress 192.0.2.10
Uninstall-WindowsFeature -Name DHCP -IncludeManagementTools -Restart
```

## 9. Autorisation DHCP dans Active Directory — le garde-fou anti-rogue

Dans un domaine AD, un serveur DHCP **non autorisé ne distribue pas d'adresses**. C'est la première défense contre les rogue DHCP (section 92).

Fonctionnement :

1. Au démarrage, le service DHCP envoie un broadcast `DHCPINFORM` pour détecter les autres serveurs DHCP du domaine.
2. Il interroge AD (`CN=NetServices,CN=Services,CN=Configuration,DC=...`) pour vérifier s'il est dans la liste des serveurs autorisés.
3. Si absent → le service **reste démarré mais ne répond à aucun DISCOVER**. L'observateur d'événements loggue l'événement **1046** ("Le serveur DHCP n'est pas autorisé").

```powershell
# Lister les serveurs DHCP autorisés dans le domaine
Get-DhcpServerInDC | Format-Table DnsName, IPAddress -AutoSize

# Autoriser (à exécuter avec un compte Enterprise Admin)
Add-DhcpServerInDC -DnsName "srv-dhcp-02.contoso.local" -IPAddress 192.0.2.11

# Retirer l'autorisation (serveur décommissionné)
Remove-DhcpServerInDC -DnsName "srv-dhcp-02.contoso.local" -IPAddress 192.0.2.11
```

> ⚠️ **En workgroup** (sans AD), il n'y a pas d'autorisation : n'importe quel serveur DHCP fonctionne. En entreprise avec AD, c'est obligatoire.

Vérification de l'état d'autorisation en console : **DHCP** → clic droit sur le serveur → si "Autoriser" est proposé, il ne l'est pas encore.

## 10. La console DHCP (dhcpmgmt.msc) — tour d'horizon

```powershell
# Ouvrir la console
dhcpmgmt.msc
```

Arborescence :

```text
DHCP
└── SRV-DHCP-01.contoso.local
    ├── Pool d'adresses IPv4
    │   └── Étendue [192.0.2.0] LAN-Bureaux
    │       ├── Pool d'adresses   (plage + exclusions)
    │       ├── Baux d'adresses   (baux actifs)
    │       ├── Réservations      (IP fixes par MAC)
    │       ├── Options d'étendue (003, 006, 015...)
    │       └── Stratégies        (filtrage par classe/VLAN)
    ├── Options de serveur        (s'appliquent à toutes les étendues)
    ├── Stratégies
    └── Filtres (Autoriser / Refuser)
```

Actions essentielles (clic droit) :

