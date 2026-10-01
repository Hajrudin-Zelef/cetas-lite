---
id: collect-261001-rattrapage/rattrapage/dns-windows-guide-3
title: "DNS sous Windows Server en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/dns_windows_guide.md
source_anchor: ""
source_lines: [200, 367]
sha256: 85d472e785fb1994b7bf06eac9eec99287d7f734e9955e6e8f6456f99aee68e6
---

# DNS sous Windows Server en entreprise — Guide technique ultra-complet

**Règle d'or des changements :** baissez le TTL à 300 s **au moins 2× l'ancien TTL avant** la migration, changez l'IP, puis remontez le TTL. Baisser le TTL *pendant* la migration ne sert à rien : les caches ont déjà l'ancienne valeur.

```powershell
# Lire le TTL d'un enregistrement (TimeToLive est un TimeSpan)
(Get-DnsServerResourceRecord -ZoneName "contoso.local" -Name "www" -RRType A).TimeToLive
# Changer le TTL : on réécrit l'enregistrement
$old = Get-DnsServerResourceRecord -ZoneName "contoso.local" -Name "www" -RRType A
$new = $old.Clone(); $new.TimeToLive = [TimeSpan]::FromMinutes(5)
Set-DnsServerResourceRecord -NewInputObject $new -OldInputObject $old -ZoneName "contoso.local" -PassThru
```

---

# Chapitre 2 — Installation du rôle DNS et prise en main

## 13. Installation du rôle DNS (GUI)

Via le **Gestionnaire de serveur → Ajouter des rôles et fonctionnalités → Rôle « Serveur DNS »**. Sur un contrôleur de domaine, la case est généralement déjà cochée (le rôle DNS s'installe avec AD DS si vous l'avez demandé dans l'assistant de promotion).

Points de vigilance GUI :
- Sur une installation **Server Core**, pas de GUI : utilisez PowerShell (§14).
- Après installation, la console `dnsmgmt.msc` apparaît dans Outils d'administration.
- Le service s'appelle `DNS Server` (nom de service : `DNS`).

## 14. Installation via PowerShell et Server Core

```powershell
# Installer le rôle DNS + outils d'administration (GUI ou Core)
Install-WindowsFeature -Name DNS -IncludeManagementTools

# Vérifier
Get-WindowsFeature -Name DNS | Select-Object Name, Installed, InstallState

# Sur un serveur distant
Install-WindowsFeature -Name DNS -IncludeManagementTools -ComputerName "SRV-DNS-02" -Credential (Get-Credential)

# Démarrage automatique + démarrage immédiat (normalement déjà OK après install)
Set-Service -Name DNS -StartupType Automatic
Start-Service -Name DNS
```

⚠️ Version : sur Windows Server 2025, `Install-WindowsFeature` existe toujours ; l'alternative moderne `Enable-WindowsOptionalFeature` ne concerne pas les rôles serveur.

## 15. La console DNS (dnsmgmt.msc)

`dnsmgmt.msc` (ou `dnsmgmt.msc /s` en autonome) affiche :
- Le serveur → **zones de recherche directe**, **zones de recherche inversée**, **redirecteurs conditionnels**, **points de confiance**.
- Clic droit sur le serveur → Propriétés : les onglets **Interfaces**, **Redirecteurs**, **Indications de racine**, **Avancé** (scavenging, recursion…), **Journalisation de débogage**.

Réflexes console :
- **Actualiser** (F5) après une action PowerShell : la console ne se rafraîchit pas toute seule.
- Affichage → **Avancé** pour voir les enregistrements cachés (type WINS, _msdcs détaillé).
- Ne modifiez jamais les fichiers de zone `.dns` à la main pendant que le service tourne.

## 16. Le module PowerShell DnsServer : tour d'horizon

Le module `DnsServer` (chargé automatiquement) couvre ~140 applets. Les familles à connaître :

| Famille | Exemples | Usage |
|---|---|---|
| Zones | `Get/Add/Remove/Set-DnsServerZone`, `Add-DnsServerPrimaryZone`… | Création, conversion, suppression |
| Enregistrements | `Get/Add/Remove/Set-DnsServerResourceRecord*` | CRUD des RR |
| Transferts | `Get/Set-DnsServerZoneTransfer` | AXFR/IXFR |
| Vieillissement | `Set-DnsServerZoneAging`, `Set-DnsServerScavenging` | Scavenging |
| Résolution | `Set/Get-DnsServerForwarder`, `Get-DnsServerRootHint`, `Set-DnsServerRecursion` | Redirecteurs/racine |
| Diagnostics | `Test-DnsServer`, `Get-DnsServerDiagnostics`, `Clear-DnsServerCache` | Dépannage |
| Stratégies | `Add-DnsServerClientSubnet`, `Add-DnsServerQueryResolutionPolicy` | DNS policies (2016+) |
| DNSSEC | `Add-DnsServerSigningKey`, `Set-DnsServerDnsSecZoneSetting` | Signature |

```powershell
# Lister toutes les applets du module
Get-Command -Module DnsServer | Select-Object Name | Sort-Object Name
# Aide détaillée d'une applet
Get-Help Add-DnsServerResourceRecordA -Full
```

## 17. dnscmd.exe : l'outil historique

`dnscmd` reste utile pour : le debug logging (`/Config /LogLevel`), le vieillissement forcé (`/AgeAllRecords`), et les scripts hérités. Il est disponible dès que le rôle DNS (ou les outils RSAT-DNS) est installé.

```powershell
# Exemples courants
dnscmd SRV-DNS-01 /Info                 # configuration complète du serveur
dnscmd SRV-DNS-01 /ZoneInfo contoso.local
dnscmd SRV-DNS-01 /ClearCache
dnscmd SRV-DNS-01 /ZoneReload contoso.local
```

**Recommandation :** privilégiez PowerShell pour tout ce qui est nouveau ; gardez `dnscmd` pour les 3 cas ci-dessus et le dépannage sur vieux systèmes.

## 18. Vérifier l'état du service DNS

```powershell
# Service + port d'écoute
Get-Service -Name DNS
Get-NetTCPConnection -LocalPort 53 -State Listen | Select-Object LocalAddress, OwningProcess
Get-NetUDPEndpoint -LocalPort 53 | Select-Object LocalAddress

# Le serveur se résout-il lui-même ?
Resolve-DnsName -Name "SRV-DNS-01.contoso.local" -Server "10.0.0.10"

# Test fonctionnel complet depuis le serveur
Test-DnsServer -IPAddress "10.0.0.10" -ZoneName "contoso.local"
```

Si le port 53 n'écoute pas : vérifiez les **interfaces** (§19) et qu'un autre service (un autre DNS, un contrôleur de domaine tiers) ne squatte pas le port.

## 19. Configurer les adresses d'écoute

Par défaut, le DNS écoute sur **toutes** les adresses. Sur un serveur multi-homé (ex : interface de sauvegarde, iLO partagé), restreignez :

```powershell
# N'écouter que sur 10.0.0.10
Set-DnsServer -ListenAddresses "10.0.0.10"
Get-DnsServer | Select-Object -ExpandProperty ListenAddresses

# Revenir à toutes les adresses
Set-DnsServer -ListenAddresses @()
```

GUI : Propriétés du serveur → onglet **Interfaces**. ⚠️ Si vous restreignez, n'oubliez pas `127.0.0.1` si des applications locales interrogent le DNS en loopback.

## 20. Configurer la carte réseau du serveur DNS lui-même

Règle d'or (source de l'erreur n°1, §123) : **un serveur DNS/AD doit se pointer lui-même en premier**.

| Serveur | DNS préféré | DNS auxiliaire |
|---|---|---|
| `SRV-DNS-01` (10.0.0.10) | `127.0.0.1` (ou 10.0.0.10) | `10.0.0.11` |
| `SRV-DNS-02` (10.0.0.11) | `127.0.0.1` (ou 10.0.0.11) | `10.0.0.10` |
| Serveur membre / poste | `10.0.0.10` | `10.0.0.11` |

```powershell
# Configurer via PowerShell (interface "Ethernet0" par ex.)
Set-DnsClientServerAddress -InterfaceAlias "Ethernet0" -ServerAddresses "127.0.0.1","10.0.0.11"
Get-DnsClientServerAddress -InterfaceAlias "Ethernet0" | Select-Object -ExpandProperty ServerAddresses
```

Ne mettez **jamais** un DNS externe (FAI, 8.8.8.8) dans la carte d'un DC : le DC doit résoudre l'AD en priorité, l'externe passe par les redirecteurs (§57).

## 21. Checklist d'installation

- [ ] Rôle DNS installé (`Get-WindowsFeature DNS` → Installed = True)
- [ ] Service `DNS` en démarrage automatique et démarré
- [ ] Interfaces d'écoute restreintes si multi-homé (§19)
- [ ] Carte réseau : lui-même en premier, partenaire en second (§20)
- [ ] Pare-feu : UDP/TCP 53 ouverts vers les clients et les autres DNS (§122)
- [ ] Redirecteurs configurés et testés (§60), ou indications de racine vérifiées
- [ ] Zone(s) créée(s), SOA/NS corrects (§31)
- [ ] Test bout-en-bout : `Test-DnsServer` + `Resolve-DnsName` depuis un client
- [ ] Supervision : le port 53 et le service sont monitorés (§144)

## 22. Premier test de résolution

```powershell
# 1. Le serveur répond-il ?
Resolve-DnsName -Name "contoso.local" -Server "10.0.0.10" -Type SOA

# 2. Récursion vers Internet OK ?
Resolve-DnsName -Name "www.example.com" -Server "10.0.0.10"

# 3. Un client du site résout-il un DC ?
Resolve-DnsName -Name "_ldap._tcp.dc._msdcs.contoso.local" -Type SRV -Server "10.0.0.10"

# 4. Temps de réponse (mesure simple)
Measure-Command { Resolve-DnsName -Name "www.contoso.local" -Server "10.0.0.10" -DnsOnly | Out-Null }
```

