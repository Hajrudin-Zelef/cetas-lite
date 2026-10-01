---
id: collect-261001-rattrapage/rattrapage/dns-windows-guide-13
title: "DNS sous Windows Server en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/dns_windows_guide.md
source_anchor: ""
source_lines: [1642, 1777]
sha256: 8fbdfb8d7f2b60de02bd718f5d12eaaae63d5ee95f9add072035a424e8bf744d
---

# DNS sous Windows Server en entreprise — Guide technique ultra-complet

**Résolution :** repassez la zone en `Secure`, corrigez le DNS du client, vérifiez la synchro horaire (`w32tm /query /status`), puis `nltest /dsregdns` / `ipconfig /registerdns`.

## 106. Cas pratique n°7 : transfert de zone refusé

**Symptômes :** le secondaire reste avec un vieux serial ; Event 7697 côté secondaire.

**Diagnostic :**
```powershell
# Côté primaire : qui est autorisé ?
Get-DnsServerZoneTransfer -Name "contoso.local" | Select-Object SecureSecondaries, SecondaryServers
# Le port TCP 53 est-il ouvert du secondaire vers le primaire ?
Test-NetConnection -ComputerName "10.0.0.10" -Port 53
# Comparer les serials
(Resolve-DnsName "contoso.local" -Type SOA -Server "10.0.0.10").SerialNumber
(Resolve-DnsName "contoso.local" -Type SOA -Server "192.168.99.10").SerialNumber
```
**Causes fréquentes :** IP du secondaire absente de la liste blanche (§50), pare-feu qui bloque TCP 53 (l'AXFR/IXFR est en TCP, pas UDP !), ou maître injoignable.

**Résolution :** ajoutez l'IP (`Set-DnsServerZoneTransfer`), ouvrez TCP 53, forcez `Sync-DnsServerZone` côté secondaire.

## 107. Cas pratique n°8 : lenteurs de résolution (timeouts)

**Symptômes :** chaque première résolution prend 3–6 s ; navigation « qui hésite ».

**Diagnostic :**
```powershell
# Mesurer : où part le temps ?
Measure-Command { Resolve-DnsName "www.contoso.local" -Server "10.0.0.10" -DnsOnly }
Measure-Command { Resolve-DnsName "www.example.com"   -Server "10.0.0.10" -DnsOnly }
# Les redirecteurs répondent-ils vite ?
1..3 | ForEach-Object { Measure-Command { Resolve-DnsName "www.example.com" -Server "1.1.1.1" -DnsOnly } }
Get-DnsServerForwarder
```
**Causes fréquentes :** redirecteur lent ou injoignable (timeout 3 s × N redirecteurs avant repli), NS fantôme, EDNS bloqué par un vieux pare-feu (les grosses réponses UDP sont droppées → repli TCP lent), ou cache froid après un `Clear-DnsServerCache` en heure de pointe.

**Résolution :** remplacez le redirecteur lent, vérifiez `-Timeout`, testez avec `-DnssecOk` pour isoler un problème EDNS, et ne videz le cache qu'en heure creuse (§95).

## 108. Cas pratique n°9 : SRV d'AD manquants après promotion d'un DC

**Symptômes :** `dcdiag /test:dns` en échec, les clients n'utilisent pas le nouveau DC, `nltest /dnsgetdc` ne le liste pas.

**Diagnostic :**
```powershell
# Le DC s'est-il enregistré ?
Get-DnsServerResourceRecord -ZoneName "_msdcs.contoso.local" -RRType SRV |
    Where-Object { $_.RecordData.DomainName -match "SRV-DC-03" }
# Que dit Netlogon ?
Get-WinEvent -LogName "System" | Where-Object Id -in 5774,5775 | Select-Object -First 5 TimeCreated, Message
```
**Causes fréquentes :** la carte réseau du nouveau DC ne pointe pas vers un DNS interne (§20), la zone `_msdcs` n'accepte pas les mises à jour, ou le service Netlogon n'a pas encore tourné (attendre ~1 h ou forcer).

**Résolution :** corrigez le DNS de la carte, puis `nltest /dsregdns` sur le nouveau DC, et revérifiez les SRV + le CNAME GUID (§7).

## 109. Cas pratique n°10 : le DC ne s'enregistre plus (Netlogon 5774/5775)

**Symptômes :** Event 5774/5775 réguliers sur un DC ; ses SRV disparaissent progressivement (scavenging) → les clients ne le trouvent plus.

**Diagnostic :**
```powershell
# Forcer puis observer
nltest /dsregdns; Start-Sleep 5
Get-WinEvent -LogName "System" -MaxEvents 20 | Where-Object Id -in 5774,5775
# Horloge (Kerberos) ?
w32tm /query /status
# La zone est-elle en Secure + AD-integrated ?
Get-DnsServerZone -Name "contoso.local" | Select-Object DynamicUpdate, IsDsIntegrated
```
**Causes fréquentes :** mot de passe du compte machine désynchronisé, horloge dérivée > 5 min, zone passée en `None`, ou ACL de la zone cassées.

**Résolution :** `w32tm /resync`, `nltest /dsregdns`, en dernier recours `netdom resetpwd` + redémarrage Netlogon. Traitez la cause, pas le symptôme : un DC qui ne s'enregistre plus est souvent un DC dont l'AD va mal.

## 110. Cas pratique n°11 : DNSSEC — validation qui échoue (SERVFAIL)

**Symptômes :** une zone signée ne se résout plus (`SERVFAIL`), les autres zones fonctionnent.

**Diagnostic :**
```powershell
# Sans validation, ça répond ?
Resolve-DnsName "www.contoso.com" -Server "10.0.0.10" -DnssecOk:$false
# Avec validation ?
Resolve-DnsName "www.contoso.com" -Server "10.0.0.10" -DnssecOk
# La chaîne est-elle intacte ? (DS chez le parent, signatures valides, horloge)
```
**Causes fréquentes :** rollover KSK sans mise à jour du DS chez le registrar (§72), signatures expirées (horloge du serveur), ou redirecteur qui **casse** DNSSEC (ne transmet pas les enregistrements) → dans ce cas la validation échoue à cause du redirecteur, pas de la zone (§74).

**Résolution :** vérifiez le DS public (`Resolve-DnsName contoso.com -Type DS -Server 8.8.8.8`), l'horloge, et contournez temporairement le redirecteur suspect en interrogeant la racine.

## 111. Cas pratique n°12 : boucle de redirecteurs

**Symptômes :** timeouts généralisés sur certaines zones ; le journal analytique montre des requêtes en boucle entre deux IP.

**Diagnostic :**
```powershell
Resolve-DnsName -Name "srv.partenaire.local" -Trace | Select-Object -Last 25
Get-DnsServerForwarder -ComputerName "SRV-DNS-01"
Get-DnsServerZone -Name "partenaire.local"   # conditional forwarder ?
```
**Cause :** A redirige vers B qui redirige vers A (classique après une fusion, §62).

**Résolution :** cassez la boucle au bon endroit (le redirecteur doit pointer vers un serveur **autoritaire** ou un résolveur Internet, jamais vers un pair qui vous renvoie la balle), documentez la chaîne complète.

## 112. Cas pratique n°13 : zone AD-integrated qui ne se réplique pas

**Symptômes :** un enregistrement créé sur `SRV-DNS-01` n'apparaît pas sur `SRV-DNS-02`, même après des heures.

**Diagnostic :**
```powershell
# La réplication AD fonctionne-t-elle ? (le DNS n'est que le messager)
repadmin /showrepl SRV-DNS-02
dcdiag /test:replications
# La zone est-elle AD-integrated des DEUX côtés ?
Get-DnsServerZone -Name "contoso.local" -ComputerName "SRV-DNS-01" | Select-Object IsDsIntegrated, DirectoryPartitionName
Get-DnsServerZone -Name "contoso.local" -ComputerName "SRV-DNS-02" | Select-Object IsDsIntegrated, DirectoryPartitionName
# Doublon fichier + AD ? (Event 4515)
Get-WinEvent -LogName "DNS Server" | Where-Object Id -eq 4515 | Select-Object -First 3
```
**Causes fréquentes :** la réplication AD est cassée (le DNS est innocent), portée de réplication différente entre serveurs, ou **zone en double** (fichier `.dns` + AD : le serveur charge le fichier et ignore l'AD → Event 4515).

**Résolution :** réparez d'abord AD (`repadmin`, `dcdiag`), harmonisez les portées (§29), supprimez le doublon fichier si 4515.

## 113. Cas pratique n°14 : cache poisoning suspecté

**Symptômes :** un nom résout vers une IP inconnue de façon intermittente ; plusieurs postes impactés (donc ce n'est pas le cache d'un seul poste).

**Diagnostic :**
```powershell
# D'où vient la réponse ? Cache serveur ou autoritaire ?
Resolve-DnsName "banque.contoso.com" -Server "10.0.0.10" -CacheOnly
Resolve-DnsName "banque.contoso.com" -Server "10.0.0.10" -DnsOnly
# Le serveur autoritaire légitime dit quoi ? (court-circuitez votre infra)
Resolve-DnsName "banque.contoso.com" -Type A -Server "8.8.8.8" -DnsOnly
```
**Conduite à tenir :** 1) `Clear-DnsServerCache` immédiat ; 2) vérifiez le cache locking (§119) et que la récursion n'est pas ouverte (§117) ; 3) journal analytique pour identifier la réponse forgée et sa source ; 4) si zone critique : envisagez DNSSEC (§69). Un « poisoning » est souvent en réalité un **vieux redirecteur compromis** ou un équipement d'interception (proxy, portail captif) — vérifiez avant d'incriminer le protocole.

## 114. Cas pratique n°15 : après migration / changement d'IP d'un DC

