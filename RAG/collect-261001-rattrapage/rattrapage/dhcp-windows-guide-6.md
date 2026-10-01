---
id: collect-261001-rattrapage/rattrapage/dhcp-windows-guide-6
title: "Guide technique ultra-complet : DHCP sous Windows Server en entreprise"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/dhcp_windows_guide.md
source_anchor: ""
source_lines: [822, 1001]
sha256: d62b1e869410923fef33ed926bc31407b9b27f59f64f73016c830ba4c3f5d8ac
---

# Guide technique ultra-complet : DHCP sous Windows Server en entreprise

Le **démarrage PXE** permet d'installer des OS via le réseau (WDS — Windows Deployment Services, MDT). Le DHCP indique au client PXE **où** trouver le serveur de démarrage :

| Option | Rôle | Valeur type |
|--------|------|-------------|
| 066 | Nom du serveur de démarrage | `srv-wds.contoso.local` ou IP |
| 067 | Nom du fichier de démarrage | `boot\x64\wdsnbp.com` (BIOS) / `boot\x64\wdsmgfw.efi` (UEFI) |

```powershell
# Configurer PXE sur l'étendue (si WDS sur un autre serveur que le DHCP)
Set-DhcpServerv4OptionValue -ScopeId 192.0.2.0 -OptionId 66 -Value "srv-wds.contoso.local"
Set-DhcpServerv4OptionValue -ScopeId 192.0.2.0 -OptionId 67 -Value "boot\x64\wdsnbp.com"
```

> ⚠️ **Si WDS et DHCP sont sur le même serveur** : NE PAS configurer 066/067 ! Cocher plutôt dans WDS : *"Ne pas écouter sur le port 67"* + *"Configurer les options DHCP 060"*. Sinon conflit de ports (les deux écoutent en UDP 67).
>
> ⚠️ **UEFI vs BIOS** : le fichier de boot diffère (`wdsnbp.com` pour BIOS, `wdsmgfw.efi` pour UEFI). Avec un parc mixte, utiliser des **stratégies DHCP** (section 52) pour servir le bon fichier selon l'architecture annoncée par le client (option 93).

Vérification côté client PXE : au boot, le client affiche l'IP du serveur de démarrage. Si le client PXE n'obtient pas d'IP → voir cas pratique n°6 (section 82).

## 38. Précédence des options : serveur < étendue < réservation < stratégie < classe

Quand une même option est définie à plusieurs niveaux, **le niveau le plus spécifique gagne** :

```text
Options de SERVEUR  (toutes les étendues)
        │
        ▼  écrasé par
Options d'ÉTENDUE   (un sous-réseau)
        │
        ▼  écrasé par
Options de RÉSERVATION (un équipement)
        │
        ▼  écrasé par
STRATÉGIE / CLASSE (filtrage fin)
```

```powershell
# Option de serveur : s'applique à toutes les étendues qui ne la redéfinissent pas
Set-DhcpServerv4OptionValue -DnsServer 192.0.2.53, 192.0.2.54 -DnsDomain "contoso.local"
# (sans -ScopeId = niveau serveur)

# Lister les options de niveau serveur
Get-DhcpServerv4OptionValue | Format-Table OptionId, Name, Value -AutoSize
```

**Stratégie recommandée** :

- **Niveau serveur** : DNS (006), domaine (015) — identiques partout.
- **Niveau étendue** : routeur (003) — différent par VLAN.
- **Niveau réservation** : exceptions (un MFP avec SMTP dédié...).
- **Stratégies** : cas fins (PXE par architecture, classes de fournisseurs).

## 39. Options personnalisées (définies par l'utilisateur)

Pour les options non présentes dans le catalogue (ex. : option 150 pour Cisco TFTP, 156 pour Alcatel) :

```powershell
# 1. Créer la définition d'option (une fois, niveau serveur)
Add-DhcpServerv4OptionDefinition `
    -Name "Cisco TFTP Server" `
    -OptionId 150 `
    -Type IPAddress `
    -Description "Serveurs TFTP pour téléphones Cisco" `
    -MultiValued

# 2. Assigner la valeur sur l'étendue voix
Set-DhcpServerv4OptionValue -ScopeId 192.0.4.0 -OptionId 150 -Value 192.0.4.10, 192.0.4.11

# 3. Vérifier
Get-DhcpServerv4OptionValue -ScopeId 192.0.4.0 -OptionId 150
```

Types disponibles : `Byte`, `Word`, `DWord`, `String`, `IPAddress`, `BinaryData`, `EncapsulatedData`.

## 40. Options utiles — tableau de référence élargi

| Option | Nom | Usage entreprise |
|--------|-----|------------------|
| 001 | Masque de sous-réseau | Envoyé automatiquement |
| 003 | Routeur | Passerelle par défaut |
| 006 | Serveurs DNS | DNS AD |
| 012 | Nom d'hôte | Peut imposer un nom (rare) |
| 015 | Nom de domaine | Suffixe DNS |
| 042 | Serveurs NTP | Serveurs de temps (ex. : contrôleur de domaine) |
| 044/046 | WINS | Héritage NetBIOS |
| 051 | Durée du bail | Automatique |
| 058/059 | T1/T2 renouvellement | Automatique (50 % / 87,5 %) |
| 060 | Class ID fournisseur | PXEClient (WDS) |
| 066 | Serveur de boot | PXE / WDS |
| 067 | Fichier de boot | PXE / WDS |
| 093 | Architecture client | PXE (BIOS=0, UEFI x64=7, UEFI x64 HTTP=16) |
| 119 | Suffixes de recherche DNS | Liste de domaines |
| 121 | Routes statiques sans classe | Routes spécifiques poussées par DHCP |
| 150 | Serveurs TFTP (Cisco) | Téléphonie Cisco |
| 156 | Alcatel | Téléphonie Alcatel-Lucent |
| 252 | WPAD | URL du fichier PAC proxy |

```powershell
# Exemple : pousser des serveurs NTP (option 042) au niveau serveur
Set-DhcpServerv4OptionValue -OptionId 42 -Value 192.0.2.53

# Exemple : WPAD pour la configuration proxy automatique
Set-DhcpServerv4OptionValue -ScopeId 192.0.2.0 -OptionId 252 -Value "http://proxy.contoso.local/wpad.dat"
```

## 41. Option 121 — Routes statiques poussées par DHCP

Permet de distribuer des **routes spécifiques** sans toucher chaque poste (ex. : joindre un réseau distant via un routeur secondaire) :

```powershell
# Format : "destination/masque,passerelle" — syntaxe PowerShell :
Set-DhcpServerv4OptionValue -ScopeId 192.0.2.0 -OptionId 121 `
    -Value "192.0.9.0/24,192.0.2.254", "192.0.10.0/24,192.0.2.254"
```

> ⚠️ À manier avec précaution : une route erronée ici = trafic détourné pour **tout le VLAN**. Tester sur une réservation d'abord (précédence, section 38).

## 42. Supprimer / réinitialiser une option

```powershell
# Supprimer une option d'une étendue (retombe sur le niveau serveur s'il existe)
Remove-DhcpServerv4OptionValue -ScopeId 192.0.2.0 -OptionId 252

# Supprimer une option de niveau serveur
Remove-DhcpServerv4OptionValue -OptionId 42

# Supprimer une définition d'option personnalisée
Remove-DhcpServerv4OptionDefinition -OptionId 150
```

## 43. Exporter la configuration des options (documentation)

```powershell
# Documenter toutes les options de toutes les étendues en CSV
Get-DhcpServerv4Scope | ForEach-Object {
    $scopeId = $_.ScopeId
    Get-DhcpServerv4OptionValue -ScopeId $scopeId | ForEach-Object {
        [PSCustomObject]@{
            Étendue  = $scopeId.IPAddressToString
            OptionId = $_.OptionId
            Nom      = $_.Name
            Valeur   = ($_.Value -join ", ")
        }
    }
} | Export-Csv "C:\Admin\dhcp_options.csv" -NoTypeInformation -Encoding UTF8
```

## 44. Bonnes pratiques — Options (récap)

1. ✅ DNS (006) + domaine (015) au **niveau serveur** : un seul endroit à maintenir.
2. ✅ Routeur (003) au **niveau étendue** : un par VLAN.
3. ✅ Exceptions au **niveau réservation** uniquement.
4. ✅ Documenter les options personnalisées (qui les utilise ? pourquoi ?).
5. ❌ Ne jamais mettre de DNS public sur les postes du domaine.
6. ❌ Ne pas configurer 066/067 si WDS est sur le même serveur que DHCP.

---

# Bloc D — Scopes avancés : superscopes, multicast, stratégies, filtres

## 45. Superscopes — regrouper des étendues (multinet)

Un **superscope** regroupe plusieurs étendues IPv4 logiques sur le **même segment physique**. Cas d'usage : migration d'adressage (ancienne plage + nouvelle plage cohabitent), ou besoin de plus d'adresses qu'un /24 sur un même VLAN.

```powershell
# 1. Créer deux étendues sur le même segment
Add-DhcpServerv4Scope -Name "LAN-Ancien"  -StartRange 192.0.2.50  -EndRange 192.0.2.200  -SubnetMask 255.255.255.0 -LeaseDuration (New-TimeSpan -Days 8) -State Active
Add-DhcpServerv4Scope -Name "LAN-Nouveau" -StartRange 192.0.11.50 -EndRange 192.0.11.200 -SubnetMask 255.255.255.0 -LeaseDuration (New-TimeSpan -Days 8) -State Active

# 2. Créer le superscope et y ajouter les étendues
Add-DhcpServerv4Superscope -SuperscopeName "Superscope-LAN" -ScopeId 192.0.2.0, 192.0.11.0

# 3. Vérifier
Get-DhcpServerv4Superscope -SuperscopeName "Superscope-LAN"
```

Comportement : le serveur attribue d'abord depuis la **première** étendue ; quand elle est pleine, il pioche dans la suivante.

> ⚠️ Prérequis réseau : le routeur du segment doit router les **deux** sous-réseaux (adresses secondaires sur l'interface, ou inter-VLAN). Sans ça, les clients de la 2e plage n'ont pas de passerelle fonctionnelle.

