---
id: collect-261001-rattrapage/rattrapage/dhcp-windows-guide-7
title: "Guide technique ultra-complet : DHCP sous Windows Server en entreprise"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["attribution", "ethernet", "valuation"]
source: docs/RAG/collect-261001-rattrapage/dhcp_windows_guide.md
source_anchor: ""
source_lines: [1002, 1195]
sha256: 07da950eba79381b80998837abc71f433fd19ae849a6479f49576001f46f1ae5
---

# Guide technique ultra-complet : DHCP sous Windows Server en entreprise

```powershell
# Retirer une étendue d'un superscope
Remove-DhcpServerv4Superscope -ScopeId 192.0.11.0

# Supprimer le superscope (les étendues sont conservées)
Remove-DhcpServerv4Superscope -SuperscopeName "Superscope-LAN"
```

## 46. Étendues de multidiffusion (multicast scopes)

Pour les applications de **multicast** (déploiement d'images WDS en multicast, visioconférence), le DHCP peut attribuer des adresses multicast (224.0.0.0 – 239.255.255.255, RFC 2365).

```powershell
# Créer une étendue multicast
Add-DhcpServerv4MulticastScope `
    -Name "Multicast-WDS" `
    -Description "Multicast pour deploiement WDS" `
    -StartRange 239.0.1.1 `
    -EndRange 239.0.1.254 `
    -TTL 32 `
    -LeaseDuration (New-TimeSpan -Days 1) `
    -State Active

# Lister les étendues multicast
Get-DhcpServerv4MulticastScope
```

> 💡 Usage réel : **rare** hors WDS/multicast vidéo. La plupart des entreprises n'en ont pas besoin. Le TTL limite la portée (32 = site local typiquement).

## 47. Stratégies DHCP — le filtrage fin (introduction)

Les **stratégies** permettent d'attribuer des IP/options différentes **au sein d'une même étendue** selon des critères : adresse MAC (préfixe constructeur), classe utilisateur, classe fournisseur, nom d'hôte...

Cas d'usage typiques :

- Servir le bon **fichier PXE** selon BIOS/UEFI (option 93).
- Réserver une sous-plage aux **téléphones IP** (préfixe MAC du constructeur).
- Appliquer des options différentes aux **MFP** (lien métier copieurs).

```powershell
# Exemple : stratégie PXE — clients UEFI x64 (option 93 = 7)
Add-DhcpServerv4Policy `
    -Name "PXE-UEFI-x64" `
    -ScopeId 192.0.2.0 `
    -Condition OR `
    -VendorClass Hex, "PXEClient:Arch:00007" `
    -Description "Clients PXE en UEFI x64"

# Assigner le fichier de boot UEFI à cette stratégie
Set-DhcpServerv4OptionValue -ScopeId 192.0.2.0 -PolicyName "PXE-UEFI-x64" `
    -OptionId 67 -Value "boot\x64\wdsmgfw.efi"
```

## 48. Stratégies par classe fournisseur (vendor class)

La **classe fournisseur** est déclarée par le client dans son DISCOVER (option 60). Valeurs courantes :

| Classe | Équipement |
|--------|------------|
| `PXEClient:Arch:00000` | PXE BIOS |
| `PXEClient:Arch:00007` | PXE UEFI x64 |
| `PXEClient:Arch:00016` | PXE UEFI x64 HTTP boot |
| `MSFT 5.0` | Client Windows |

```powershell
# Lister les classes fournisseurs connues du serveur
Get-DhcpServerv4Class -Type Vendor

# Stratégie complète : BIOS vs UEFI
Add-DhcpServerv4Policy -Name "PXE-BIOS" -ScopeId 192.0.2.0 `
    -Condition OR -VendorClass Hex, "PXEClient:Arch:00000"
Set-DhcpServerv4OptionValue -ScopeId 192.0.2.0 -PolicyName "PXE-BIOS" `
    -OptionId 67 -Value "boot\x64\wdsnbp.com"

Add-DhcpServerv4Policy -Name "PXE-UEFI" -ScopeId 192.0.2.0 `
    -Condition OR -VendorClass Hex, "PXEClient:Arch:00007", "PXEClient:Arch:00016"
Set-DhcpServerv4OptionValue -ScopeId 192.0.2.0 -PolicyName "PXE-UEFI" `
    -OptionId 67 -Value "boot\x64\wdsmgfw.efi"
```

## 49. Stratégies par classe utilisateur (user class)

La **classe utilisateur** (option 77) est configurée **côté client** — utile pour distinguer des populations :

```powershell
# Côté client Windows : définir sa classe utilisateur
ipconfig /setclassid "Ethernet" "Imprimantes"

# Côté serveur : créer la classe puis la stratégie
Add-DhcpServerv4Class -Name "MFP" -Type User -Data "Imprimantes" -Description "Multifonctions"

Add-DhcpServerv4Policy -Name "Policy-MFP" -ScopeId 192.0.6.0 `
    -Condition OR -UserClass Hex, "Imprimantes"

# Options spécifiques aux MFP : serveur SMTP dédié (option 69), par exemple
Set-DhcpServerv4OptionValue -ScopeId 192.0.6.0 -PolicyName "Policy-MFP" `
    -OptionId 69 -Value 192.0.6.25
```

> **Lien métier copieurs** : avec une classe utilisateur `MFP`, tu peux pousser automatiquement aux copieurs : serveur SMTP de scan-to-mail dédié, serveur LDAP du carnet d'adresses, ou un DNS interne spécifique — sans toucher chaque machine.

## 50. Stratégies par plage d'adresses (sous-plage dédiée)

Une stratégie peut restreindre l'attribution à une **sous-plage** de l'étendue :

```powershell
# Les téléphones (préfixe MAC 00:1B:2A = constructeur fictif) prennent dans .180-.200
Add-DhcpServerv4Policy `
    -Name "Telephones-IP" `
    -ScopeId 192.0.4.0 `
    -Condition OR `
    -MacAddress EQ, "001B2A*" `
    -StartRange 192.0.4.180 `
    -EndRange 192.0.4.200 `
    -Description "Telephones IP du constructeur X"
```

Ordre d'évaluation : les stratégies sont évaluées **dans l'ordre** (champ `ProcessingOrder`). La première qui correspond gagne.

```powershell
# Voir l'ordre des stratégies et le modifier
Get-DhcpServerv4Policy -ScopeId 192.0.4.0 | Select-Object Name, ProcessingOrder, Enabled
Set-DhcpServerv4Policy -Name "Telephones-IP" -ScopeId 192.0.4.0 -ProcessingOrder 1
```

## 51. Filtres MAC : liste d'autorisation (Allow) et de refus (Deny)

Les **filtres** agissent au **niveau serveur** (toutes les étendues) : autoriser ou refuser des MAC.

```powershell
# 1. Activer les listes (par défaut : Allow activée, Deny désactivée)
Set-DhcpServerv4FilterList -Allow $true -Deny $true

# 2. Ajouter une MAC en liste blanche (ex : seuls les équipements connus obtiennent une IP)
Add-DhcpServerv4Filter -List Allow -MacAddress "00-1B-A9-3F-2C-7D" -Description "MFP-Compta-RDC"

# 3. Bannir une MAC (ex : équipement personnel non autorisé)
Add-DhcpServerv4Filter -List Deny -MacAddress "AA-BB-CC-DD-EE-FF" -Description "Poste perso - refuse"

# 4. Lister
Get-DhcpServerv4Filter -List Allow
Get-DhcpServerv4Filter -List Deny
```

> ⚠️ **Allow = mode strict** : si la liste Allow est activée **et non vide**, seules les MAC listées obtiennent un bail. À n'activer que sur des VLANs contrôlés (ex. : VLAN imprimantes, VLAN serveurs de test). **Jamais** sur le LAN utilisateurs sans un processus d'enregistrement des MAC — sinon helpdesk submergé.
>
> 💡 Alternative moderne : **802.1X** sur les switchs (authentification avant DHCP) — plus robuste que le filtrage MAC (MAC spoofable en 30 secondes).

## 52. Stratégie PXE complète : BIOS + UEFI + HTTP boot

```powershell
# Nettoyer d'abord les options 066/067 globales si WDS est sur un autre serveur
# (elles restent utiles comme valeurs par défaut)

# BIOS legacy
Add-DhcpServerv4Policy -Name "PXE-BIOS" -ScopeId 192.0.2.0 -Condition OR `
    -VendorClass Hex, "PXEClient:Arch:00000"
Set-DhcpServerv4OptionValue -ScopeId 192.0.2.0 -PolicyName "PXE-BIOS" `
    -OptionId 66 -Value "srv-wds.contoso.local"
Set-DhcpServerv4OptionValue -ScopeId 192.0.2.0 -PolicyName "PXE-BIOS" `
    -OptionId 67 -Value "boot\x64\wdsnbp.com"

# UEFI x64 (TFTP)
Add-DhcpServerv4Policy -Name "PXE-UEFI" -ScopeId 192.0.2.0 -Condition OR `
    -VendorClass Hex, "PXEClient:Arch:00007"
Set-DhcpServerv4OptionValue -ScopeId 192.0.2.0 -PolicyName "PXE-UEFI" `
    -OptionId 66 -Value "srv-wds.contoso.local"
Set-DhcpServerv4OptionValue -ScopeId 192.0.2.0 -PolicyName "PXE-UEFI" `
    -OptionId 67 -Value "boot\x64\wdsmgfw.efi"

# UEFI HTTP boot (moderne, plus rapide que TFTP)
Add-DhcpServerv4Policy -Name "PXE-UEFI-HTTP" -ScopeId 192.0.2.0 -Condition OR `
    -VendorClass Hex, "PXEClient:Arch:00016"
Set-DhcpServerv4OptionValue -ScopeId 192.0.2.0 -PolicyName "PXE-UEFI-HTTP" `
    -OptionId 66 -Value "srv-wds.contoso.local"
Set-DhcpServerv4OptionValue -ScopeId 192.0.2.0 -PolicyName "PXE-UEFI-HTTP" `
    -OptionId 67 -Value "boot\x64\wdsmgfw.efi"
```

## 53. Classes personnalisées : créer ses propres classes

```powershell
# Classe fournisseur personnalisée (ex : un constructeur de MFP)
Add-DhcpServerv4Class -Name "MFP-ConstructeurX" -Type Vendor `
    -Data "ConstructeurX-MFP" -Description "MFP du constructeur X"

# Classe utilisateur personnalisée
Add-DhcpServerv4Class -Name "Visiteurs" -Type User `
    -Data "Guest" -Description "Equipements invites"

# Lister
Get-DhcpServerv4Class | Select-Object Name, Type, Data
```

