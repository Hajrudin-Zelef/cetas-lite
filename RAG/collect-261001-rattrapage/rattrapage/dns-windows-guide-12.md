---
id: collect-261001-rattrapage/rattrapage/dns-windows-guide-12
title: "DNS sous Windows Server en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-rattrapage/dns_windows_guide.md
source_anchor: ""
source_lines: [1503, 1641]
sha256: d62e9c7c0869292bea85148ac9f4998727f2f6bdb017b847ebf7843c3cb9a203
---

# DNS sous Windows Server en entreprise — Guide technique ultra-complet

Filtrez par IP client pour traquer un poste qui spamme (boucle applicative, malware de type DGA) : triez par QNAME distincts.

## 98. Debug logging (dnscmd /LogLevel)

Le debug logging écrit dans `dns.log` avec un niveau de détail réglable par masque hexadécimal.

```powershell
# Activer un debug complet temporaire (valeur courante de debug exhaustif)
dnscmd SRV-DNS-01 /Config /LogLevel 0x8100F331
dnscmd SRV-DNS-01 /Config /LogFilePath "C:\Windows\System32\dns\dns.log"
dnscmd SRV-DNS-01 /Config /LogFileMaxSize 0x4000000   # 64 Mo

# Désactiver après usage (0x0 = aucun log)
dnscmd SRV-DNS-01 /Config /LogLevel 0x0
```

⚠️ Ne laissez **jamais** le debug logging actif en permanence : le fichier grossit vite et chaque requête coûte des E/S disque. Activez, reproduisez, capturez, désactivez.

## 99. Event IDs à connaître

| ID | Source | Signification | Action |
|---|---|---|---|
| 2 / 3 | DNS Server | Démarrage/arrêt du service | Corrélé aux maintenances |
| 2501 | DNS Server | Scavenging : N enregistrements supprimés | Vérifier que N est plausible (§56) |
| 4004 / 4015 | DNS Server | Le serveur n'a pas pu charger les zones AD | Vérifier AD/réplication, puis `nltest /dsregdns` |
| 4013 | DNS Server | Zones AD non chargées au démarrage (AD pas prêt) | Souvent transitoire ; persistant = problème AD |
| 4515 | DNS Server | **Zone en double** (fichier + AD) | Supprimer le doublon (cas n°13, §112) |
| 7697 | DNS Server | Échec de transfert de zone | Liste blanche, pare-feu, maître |
| 5774 / 5775 | Netlogon | Échec d'enregistrement DNS du DC | Droits, zone, DNS du DC lui-même |
| 5781 | Netlogon | Aucun DC localisable (dynamique) | Conséquence, pas la cause : cherchez le DNS |

```powershell
# Balayer les erreurs DNS des dernières 24 h
Get-WinEvent -LogName "DNS Server" |
    Where-Object { $_.TimeCreated -gt (Get-Date).AddHours(-24) -and $_.Level -le 2 } |
    Select-Object TimeCreated, Id, @{n='Message';e={$_.Message.Substring(0,100)}} |
    Sort-Object TimeCreated
```

---

## 100. Cas pratique n°1 : le client ne résout plus rien

**Symptômes :** un poste (ou un VLAN entier) ne résout aucun nom, interne comme externe. Les IP fonctionnent (ping 10.0.1.50 OK).

**Diagnostic :**
```powershell
# Sur le poste : que vaut sa config ?
Get-DnsClientServerAddress -InterfaceAlias "Ethernet0" | Select-Object -ExpandProperty ServerAddresses
# Teste-t-il le bon serveur ?
Resolve-DnsName "www.contoso.local" -Server "10.0.0.10" -ErrorAction Stop
# Le serveur DNS est-il joignable en 53 ?
Test-NetConnection -ComputerName "10.0.0.10" -Port 53
```
**Causes fréquentes :** mauvaise IP DNS distribuée par DHCP (option 006), pare-feu local/VLAN qui bloque l'UDP 53, carte réseau du poste avec DNS en dur obsolète, ou le serveur DNS lui-même arrêté.

**Résolution :** corriger la source (DHCP, GPO, pare-feu), `Clear-DnsClientCache`, puis valider avec `Resolve-DnsName` sans `-Server` (config par défaut du poste).

## 101. Cas pratique n°2 : résolution intermittente

**Symptômes :** ça marche, puis ça ne marche plus, puis ça remarche. Souvent « le matin ça rame ».

**Diagnostic :**
```powershell
# Le poste a-t-il DEUX DNS dont un mort ?
Get-DnsClientServerAddress | Select-Object InterfaceAlias, ServerAddresses
# Tester chaque serveur séparément, plusieurs fois
1..5 | ForEach-Object { Resolve-DnsName "www.contoso.local" -Server "10.0.0.11" -DnsOnly | Select-Object -ExpandProperty IPAddress }
# NS fantôme dans la zone ?
Get-DnsServerResourceRecord -ZoneName "contoso.local" -RRType NS
```
**Causes fréquentes :** DNS secondaire configuré mais hors ligne (le client bascule après timeout), NS fantôme dans la zone (serveur décommissionné), round-robin vers une IP morte (§36), ou scavenging qui a supprimé un enregistrement encore utilisé.

**Résolution :** supprimez les serveurs morts de la config DHCP et des NS de zone ; remplacez le round-robin aveugle par un vrai équilibreur si la dispo compte.

## 102. Cas pratique n°3 : un nom interne résout vers une IP externe

**Symptômes :** `intranet.contoso.com` résout `203.0.113.10` (public) au lieu de `10.0.1.50` sur le LAN → hairpinning, lenteurs, blocages pare-feu.

**Diagnostic :**
```powershell
Resolve-DnsName "intranet.contoso.com" -Server "10.0.0.10" -DnsOnly
# D'où vient la réponse ? De votre zone ou du redirecteur ?
Get-DnsServerResourceRecord -ZoneName "contoso.com" -Name "intranet" -ErrorAction SilentlyContinue
Get-DnsServerQueryResolutionPolicy -ZoneName "contoso.com"
```
**Causes fréquentes :** pas de zone/vue interne pour ce nom (split-brain non géré, §63), stratégie DNS dont le subnet ne couvre pas le VLAN du client (§66), ou enregistrement interne oublié lors de la publication d'un nouveau service public.

**Résolution :** créer l'enregistrement dans la vue/scope interne (ou la zone interne), tester depuis chaque type de réseau (LAN, VPN, Wi-Fi invité).

## 103. Cas pratique n°4 : enregistrements dupliqués / round-robin inattendu

**Symptômes :** un nom résout vers 3 IP dont une obsolète ; une appli tombe une fois sur trois.

**Diagnostic :**
```powershell
# Lister TOUS les A du nom (nslookup n'en montre qu'un !)
Get-DnsServerResourceRecord -ZoneName "contoso.local" -Name "appli" -RRType A |
    Select-Object HostName, Timestamp, @{n='IP';e={$_.RecordData.IPv4Address}}
# Lequel est statique, lequel est dynamique ?
```
**Causes fréquentes :** ancien A statique oublié lors d'une migration + nouveau A dynamique (doublon), ou deux machines qui s'enregistrent sous le même nom (image dupliquée sans sysprep → même nom d'hôte).

**Résolution :** supprimer l'enregistrement obsolète (`Remove-DnsServerResourceRecord` en précisant `-RecordData`), corriger la cause (sysprep, inventaire), et baisser le TTL avant les prochaines migrations (§12).

## 104. Cas pratique n°5 : le scavenging a supprimé des enregistrements valides

**Symptômes :** après un week-end prolongé ou des congés, des machines/équipements « disparaissent » du DNS.

**Diagnostic :**
```powershell
# Le scavenging a-t-il tourné ? Combien de suppressions ?
Get-WinEvent -LogName "DNS Server" | Where-Object Id -eq 2501 |
    Select-Object -First 5 TimeCreated, Message
# Comparer bail DHCP et intervalles
Get-DhcpServerv4Scope -ComputerName "SRV-DHCP-01" | Select-Object Name, LeaseDuration
Get-DnsServerZoneAging -Name "contoso.local" | Select-Object NoRefreshInterval, RefreshInterval
Get-DnsServerScavenging | Select-Object ScavengingState, ScavengingInterval
```
**Cause typique :** `no-refresh + refresh` (ex : 7+7 j) < indisponibilité des machines (congés 3 semaines) ou < bail DHCP (30 j). Le timestamp n'a pas été rafraîchi → périmé → supprimé.

**Résolution :** restaurez les enregistrements (recréation ou restauration §143), **augmentez** les intervalles (ex : 14+14 j) ou réduisez les baux DHCP, et pour les équipements critiques : enregistrements **statiques**.

## 105. Cas pratique n°6 : les mises à jour dynamiques échouent

**Symptômes :** les nouveaux postes n'apparaissent pas dans le DNS ; erreurs Netlogon 5774/5775 sur les DC.

**Diagnostic :**
```powershell
# La zone accepte-t-elle les mises à jour sécurisées ?
Get-DnsServerZone -Name "contoso.local" | Select-Object DynamicUpdate, IsDsIntegrated
# Le client pointe-t-il un DNS qui fait autorité pour la zone ?
Get-DnsClientServerAddress
# Test manuel d'enregistrement
ipconfig /registerdns
Get-WinEvent -LogName "System" | Where-Object Id -in 5774,5775 | Select-Object -First 3
```
**Causes fréquentes :** zone en `None` au lieu de `Secure`, client qui pointe vers un DNS externe ou un secondaire, permissions insuffisantes sur la zone (héritage cassé), ou horloge désynchronisée (Kerberos refuse → l'update sécurisé échoue silencieusement).

