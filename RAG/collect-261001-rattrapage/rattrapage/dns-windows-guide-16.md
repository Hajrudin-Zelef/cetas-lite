---
id: collect-261001-rattrapage/rattrapage/dns-windows-guide-16
title: "DNS sous Windows Server en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/dns_windows_guide.md
source_anchor: ""
source_lines: [2041, 2181]
sha256: c8b94b4f88dea0d721ab8d4e241dc904c4c5a38ebe59fa77df37434b716c0c78
---

# DNS sous Windows Server en entreprise — Guide technique ultra-complet

**Avant :**
- [ ] Architecture validée : 2+ DNS par site, AD-integrated, portées de réplication définies (§85-86)
- [ ] Zones créées (directes + inversées), SOA/NS cohérents (§31-32, §43)
- [ ] Redirecteurs testés un par un (`Test-DnsServer -Context Forwarder`) ou root hints vérifiés (§60)
- [ ] Aging/scavenging paramétrés **et** documentés (qui scavenger, quels intervalles) (§54-56)
- [ ] DHCP : compte dédié + DnsUpdateProxy + modèle A/B choisi (§80)
- [ ] Transferts en liste blanche, récursion restreinte si exposé (§50, §117)
- [ ] TTL de migration baissés si bascule prévue (§12)

**Le jour J :**
- [ ] `Test-DnsServer` complet sur chaque serveur (§94)
- [ ] Résolution testée depuis chaque VLAN/site (interne, externe, SRV AD)
- [ ] Supervision branchée (sondes §89)
- [ ] Journal d'audit DNS activé (§96)

**Après (semaine 1) :**
- [ ] Revue des Event 2501/4004/4015/4515 (§99)
- [ ] Vérification des serials SOA entre serveurs (§51)
- [ ] Dossier d'exploitation à jour (schéma, chaînes de redirecteurs, scopes)

## 142. Checklist d'audit trimestriel

- [ ] Groupe **DnsAdmins** : membres légitimes uniquement (§120)
- [ ] Groupe **DnsUpdateProxy** : pas de DC dedans (§139)
- [ ] Zones : pas de doublon fichier+AD (Event 4515), NS à jour (pas de fantômes)
- [ ] Scavenging : Event 2501 plausibles, intervalles toujours cohérents avec les baux DHCP
- [ ] Enregistrements statiques critiques : inventaire à jour (serveurs, imprimantes, équipements)
- [ ] Redirecteurs : temps de réponse toujours OK ; root hints à jour
- [ ] DNSSEC : clés valides, prochains rollovers planifiés (§72)
- [ ] Sauvegarde des zones testée (restauration sur maquette, §143)
- [ ] Revue des stratégies DNS / scopes : la matrice split-brain est-elle toujours vraie ? (§68)
- [ ] Correctifs Windows installés sur les DNS (CVE DNS, §121)

## 143. Sauvegarde et restauration des zones

**Zones AD-integrated :** sauvegardées avec l'**état système** de l'AD (Windows Server Backup). Pas de fichier à copier : la restauration passe par l'AD.

**Zones standard (fichier) :** copiez `%SystemRoot%\System32\dns\*.dns` + exportez la config :

```powershell
# Sauvegarde : export de la zone + config du serveur
Export-DnsServerZone -Name "dmz.contoso.local" -FileName "sav-dmz.contoso.local.dns"
dnscmd SRV-DNS-01 /ZoneExport dmz.contoso.local sav-dmz2.txt
Copy-Item "C:\Windows\System32\dns\*.dns" "\\sauvegarde\dns\SRV-DNS-01\"

# Sauvegarde de la config complète (à garder avec)
Get-DnsServerZone | Export-Csv "\\sauvegarde\dns\zones.csv" -NoTypeInformation
Get-DnsServerForwarder | Export-Csv "\\sauvegarde\dns\forwarders.csv" -NoTypeInformation
```

**Restauration d'une zone standard :** recopiez le `.dns`, recréez la zone en pointant vers le fichier existant (`Add-DnsServerPrimaryZone -ZoneFile`), `Sync-DnsServerZone`.

**Restauration d'un enregistrement supprimé par erreur :** pas de corbeille AD pour le DNS → recréez-le (`Add-DnsServerResourceRecord*`) ou restaurez l'état système (lourd). D'où l'importance de l'**audit** (§96) pour savoir *quoi* recréer.

## 144. Supervision : compteurs et alertes

Compteurs de performance Windows utiles (`\DNS\...`) :

| Compteur | Ce qu'il dit | Alerte si |
|---|---|---|
| `Total Query Received/sec` | Charge | Variation brutale inexpliquée |
| `Recursive Queries/sec` | Part de récursion | Explosion = boucle ou attaque |
| `Recursive Query Failures/sec` | Échecs de récursion | > 1 % des récursives |
| `Zone Transfer Requests Received` | Demandes AXFR | Toute demande hors secondaires connus |
| `Dynamic Updates Received/sec` | Mises à jour dynamiques | Chute brutale = DHCP/AD en panne |

```powershell
# Relever les compteurs DNS en une ligne (échantillon 30 s)
Get-Counter "\DNS\Total Query Received/sec", "\DNS\Recursive Query Failures/sec" -SampleInterval 5 -MaxSamples 6

# Statistiques du service (alternative intégrée)
Get-DnsServerStatistics -ComputerName "SRV-DNS-01" | Select-Object -ExpandProperty QueryStatistics
```

Intégrez au minimum les sondes du §89 dans votre supervision (Zabbix/PRTG/SCOM) : **port 53, témoin interne, récursion, SRV AD, cohérence des serials**.

---

## 145. Pense-bête des commandes

### Zones
```powershell
Add-DnsServerPrimaryZone -Name "contoso.local" -ReplicationScope "Domain" -DynamicUpdate "Secure"
Add-DnsServerPrimaryZone -NetworkID "10.1.0.0/24" -ReplicationScope "Domain" -DynamicUpdate "Secure"
Add-DnsServerSecondaryZone -Name "contoso.local" -ZoneFile "contoso.local.dns" -MasterServers "10.0.0.10"
Add-DnsServerStubZone -Name "partenaire.local" -MasterServers "192.168.50.10"
Add-DnsServerConditionalForwarderZone -Name "partenaire.local" -MasterServers "192.168.50.10"
Get-DnsServerZone | Format-Table ZoneName, ZoneType, IsDsIntegrated
Suspend-DnsServerZone / Resume-DnsServerZone / Remove-DnsServerZone -Name "zone"
Sync-DnsServerZone -Name "contoso.local"
Export-DnsServerZone -Name "contoso.local" -FileName "export.txt"
```

### Enregistrements
```powershell
Add-DnsServerResourceRecordA -Name "www" -ZoneName "contoso.local" -IPv4Address "10.0.1.50" -TimeToLive 01:00:00
Add-DnsServerResourceRecordCName -Name "intranet" -ZoneName "contoso.local" -HostNameAlias "srv-web-01.contoso.local."
Add-DnsServerResourceRecordMX -Name "@" -ZoneName "contoso.local" -MailExchange "mx1.contoso.local." -Preference 10
Add-DnsServerResourceRecordSrv -Name "_ldap._tcp" -ZoneName "contoso.local" -DomainName "srv-ad-01.contoso.local." -Port 389 -Priority 10 -Weight 100
Add-DnsServerResourceRecordPtr -Name "50" -ZoneName "1.10.in-addr.arpa" -PtrDomainName "www.contoso.local."
Add-DnsServerResourceRecordTXT -Name "@" -ZoneName "contoso.local" -DescriptiveText "v=spf1 mx -all"
Get-DnsServerResourceRecord -ZoneName "contoso.local" -RRType A
Remove-DnsServerResourceRecord -ZoneName "contoso.local" -Name "www" -RRType A -RecordData "10.0.1.50" -Force
```

### Transferts, aging, redirecteurs
```powershell
Set-DnsServerZoneTransfer -Name "contoso.local" -SecondaryServers @("10.0.0.11")
Set-DnsServerZoneAging -Name "contoso.local" -Aging $true -RefreshInterval 7.00:00:00 -NoRefreshInterval 7.00:00:00
Set-DnsServerScavenging -ScavengingState $true -ScavengingInterval 7.00:00:00
Start-DnsServerScavenging -Force
Set-DnsServerForwarder -IPAddress @("1.1.1.1","9.9.9.9") -UseRootHint $false
Get-DnsServerRootHint | Format-Table NameServer, IPAddress
```

### Diagnostic
```powershell
Resolve-DnsName -Name "www.contoso.local" -Type A -Server "10.0.0.10" -DnsOnly
Resolve-DnsName -Name "www.contoso.local" -CacheOnly
Resolve-DnsName -Name "_ldap._tcp.dc._msdcs.contoso.local" -Type SRV
Test-DnsServer -IPAddress "10.0.0.10" -ZoneName "contoso.local"
Clear-DnsServerCache -Force; Clear-DnsClientCache
dnscmd SRV-DNS-01 /ClearCache; dnscmd SRV-DNS-01 /ZoneReload contoso.local
nltest /dsregdns; nltest /dnsgetdc:contoso.local; dcdiag /test:dns
Get-DnsServerDiagnostics | Format-List
```

### Stratégies DNS & DNSSEC
```powershell
Add-DnsServerClientSubnet -Name "Subnet-LAN" -IPv4Subnet "10.0.0.0/8"
Add-DnsServerZoneScope -ZoneName "contoso.com" -Name "Scope-Interne"
Add-DnsServerQueryResolutionPolicy -Name "P1" -Action ALLOW -ClientSubnet "EQ,Subnet-LAN" -ZoneName "contoso.com" -ZoneScope "Scope-Interne,1"
Add-DnsServerSigningKey -ZoneName "contoso.com" -KeyType KeySigningKey -CryptoAlgorithm RsaSha256
Set-DnsServerDnsSecZoneSetting -ZoneName "contoso.com" -SignWithNSEC3 $true
Get-DnsServerTrustAnchor
```

---

## 146. Quiz : 10 questions

