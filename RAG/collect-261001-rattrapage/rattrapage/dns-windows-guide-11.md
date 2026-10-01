---
id: collect-261001-rattrapage/rattrapage/dns-windows-guide-11
title: "DNS sous Windows Server en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/dns_windows_guide.md
source_anchor: ""
source_lines: [1344, 1502]
sha256: 90100be83be8958b5e64625f10ab21132a027b73b7ae2bec4da7a8d81666a727
---

# DNS sous Windows Server en entreprise — Guide technique ultra-complet

| Sonde | Seuil d'alerte | Commande |
|---|---|---|
| Port 53 TCP/UDP répond | Down > 2 min | `Test-NetConnection -ComputerName 10.0.0.10 -Port 53` |
| Résolution d'un témoin interne | Échec | `Resolve-DnsName temoin.contoso.local -Server 10.0.0.10` |
| Résolution récursive (Internet) | Échec | `Resolve-DnsName www.example.com -Server 10.0.0.10` |
| Temps de réponse | > 500 ms répété | `Measure-Command { Resolve-DnsName … }` |
| Serial SOA cohérent entre serveurs | Divergence > 1 h | Script §87 |
| Scavenging : Event 2501 | Compte anormal | Collecte des événements DNS |

```powershell
# Sonde complète minimale (à planifier toutes les 5 min)
$dns = "10.0.0.10"
$tests = @(
    @{ Nom="Témoin interne";  OK = [bool](Resolve-DnsName "temoin.contoso.local" -Server $dns -ErrorAction SilentlyContinue) },
    @{ Nom="Récursion";       OK = [bool](Resolve-DnsName "www.example.com"     -Server $dns -ErrorAction SilentlyContinue) },
    @{ Nom="SRV AD";          OK = [bool](Resolve-DnsName "_ldap._tcp.dc._msdcs.contoso.local" -Type SRV -Server $dns -ErrorAction SilentlyContinue) }
)
$tests | ForEach-Object { if (-not $_.OK) { Write-Warning "ÉCHEC : $($_.Nom) sur $dns" } }
```

---

# Chapitre 12 — Dépannage

## 90. Méthodologie de dépannage DNS

Ne partez jamais dans tous les sens. Ordre systématique :

1. **Qualifier** : un poste ? un site ? tout le monde ? Un nom ? tous les noms ? Interne ? Internet ?
2. **Localiser la couche** : client (cache, suffixe, carte) → serveur (service, zone, transfert) → récursion (redirecteurs, pare-feu) → autoritaire distant.
3. **Reproduire** avec `Resolve-DnsName` en précisant `-Server` : comparez serveur par serveur.
4. **Éliminer le cache** : videz le cache client PUIS serveur avant de conclure (§95).
5. **Lire les logs** : événements DNS + journal analytique (§96-98).
6. **Corriger, vérifier, documenter** : notez la cause racine dans le ticket, pas juste « redémarré le service ».

```powershell
# Kit de premier niveau (à garder sous la main)
$Nom = "appli-metier.contoso.local"; $Srv = "10.0.0.10"
Resolve-DnsName -Name $Nom -Server $Srv -DnsOnly          # le serveur sait-il ?
Resolve-DnsName -Name $Nom -Server $Srv -CacheOnly        # vient-ce du cache ?
Resolve-DnsName -Name $Nom                                 # config client par défaut
```

## 91. nslookup : bases et pièges

`nslookup` existe toujours, mais il a des comportements trompeurs :

```
C:\> nslookup www.contoso.local
Serveur :   SRV-DNS-01.contoso.local
Address:    10.0.0.10
Nom :       www.contoso.local
Address:    10.0.1.50
```

**Pièges :**
- En mode interactif, `nslookup` ajoute parfois des suffixes de façon inattendue ; préférez `Resolve-DnsName` en script.
- `nslookup` n'affiche qu'**une** réponse A même s'il y en a plusieurs (round-robin masqué).
- Le timeout par défaut (2 s × 2 essais) est court : un « DNS request timed out » peut être un pare-feu qui *drop* (pas de rejet ICMP) plutôt qu'un DNS lent.
- Pour tester un transfert : `nslookup` → `ls contoso.local` — refusé si non autorisé (§50).

## 92. Resolve-DnsName : le successeur

L'applet moderne : pipeline PowerShell, types complets, DNSSEC, EDNS, cache.

```powershell
# Usages essentiels
Resolve-DnsName -Name "www.contoso.local" -Type ALL -Server "10.0.0.10"
Resolve-DnsName -Name "contoso.local" -Type MX -DnsOnly            # sans cache client
Resolve-DnsName -Name "contoso.local" -Type MX -CacheOnly         # QUE le cache
Resolve-DnsName -Name "www.contoso.local" -Server "10.0.0.10" -NoHostsFile  # ignore le hosts
Resolve-DnsName -Name "10.1.0.50"                                  # reverse auto (PTR)
Resolve-DnsName -Name "contoso.com" -Type SOA -DnssecOk           # avec bit DNSSEC
Resolve-DnsName -Name "www.example.com" -Trace                    # trace itérative complète

# Sortie exploitable en script
(Resolve-DnsName "www.contoso.local" -Type A -Server "10.0.0.10").IPAddress
```

**Réflexe :** un résultat différent entre `-DnsOnly` et sans option = le cache client vous mentait. Un résultat différent entre deux `-Server` = divergence de serveurs.

## 93. dnscmd pour le diagnostic

```powershell
dnscmd SRV-DNS-01 /Info                    # config serveur complète
dnscmd SRV-DNS-01 /ZoneInfo contoso.local  # détail d'une zone (aging, transferts…)
dnscmd SRV-DNS-01 /ZonePrint contoso.local # dump complet de la zone (équiv. fichier)
dnscmd SRV-DNS-01 /EnumZones               # toutes les zones + types
dnscmd SRV-DNS-01 /Statistics              # compteurs (requêtes, récursion, transferts)
```

`/Statistics` est précieux pour objectiver « le DNS est lent » : `Recursive queries`, `Recursive query failures`, `UDP/TCP queries received`.

## 94. Test-DnsServer

L'applet de test fonctionnel intégrée : elle exécute une batterie de tests (résolution simple, récursive, transfert…) contre un serveur.

```powershell
# Batterie complète contre 10.0.0.10
Test-DnsServer -IPAddress "10.0.0.10" | Format-Table TestName, Result -AutoSize

# Test ciblé : le serveur fait-il autorité pour la zone ?
Test-DnsServer -IPAddress "10.0.0.10" -ZoneName "contoso.local" -Context "Authoritative"

# Le redirecteur répond-il ?
Test-DnsServer -IPAddress "1.1.1.1" -Context "Forwarder" -ComputerName "SRV-DNS-01"
```

Un `Result = Failure` sur le contexte `Recursion` + `Authoritative = Pass` = problème de redirecteurs/racine, pas de zone.

## 95. Vider le cache : client et serveur

Le cache est le suspect n°1 des « ça marche sur son poste mais pas le mien ».

```powershell
# Côté client
Clear-DnsClientCache
ipconfig /displaydns | Select-String "www.contoso.local" -Context 2,6  # inspecter avant de vider

# Côté serveur (vide le cache du service DNS : récursion + réponses négatives)
Clear-DnsServerCache -ComputerName "SRV-DNS-01" -Force

# Voir la taille du cache serveur
(Get-DnsServerStatistics -ComputerName "SRV-DNS-01").CacheStatistics
```

⚠️ Vider le cache serveur en pleine journée = pic de requêtes récursives (tout est à re-résoudre). À faire en heure creuse si le serveur est très sollicité. Le cache **négatif** (NXDOMAIN, TTL = Minimum du SOA, §43) est souvent le vrai coupable après la création d'un enregistrement « qui ne se résout pas ».

## 96. Les journaux DNS : lequel regarder

| Journal | Où | Contenu |
|---|---|---|
| Événements DNS Server | Observateur d'événements → Journaux Windows → DNS Server | Démarrage, transferts, scavenging (2501), erreurs AD (4013/4015) |
| Analytique DNS | Journaux des applications et services → Microsoft → Windows → DNS-Server → Analytical | **Chaque requête/réponse** (à activer, verbeux) |
| Audit DNS | … → DNS-Server → Audit | Modifications de zones (qui a créé/supprimé quoi) |
| Debug logging | Fichier `%SystemRoot%\System32\dns\dns.log` | Requêtes brutes avec flags (ancien mais précis) |
| Netlogon | `%SystemRoot%\debug\netlogon.log` | Enregistrements SRV des DC (5774/5775 en échec) |

Activez l'**Audit** en permanence (coût faible, précieux en enquête) ; l'**Analytique** uniquement pendant un dépannage ciblé (volume énorme).

```powershell
# Activer le journal analytique (désactivé par défaut)
wevtutil sl Microsoft-Windows-DNSServer/Analytical /e:true
# Le désactiver après le dépannage !
wevtutil sl Microsoft-Windows-DNSServer/Analytical /e:false
```

## 97. Journal analytique DNS

Une fois activé, chaque requête apparaît avec : horodatage, IP client, QNAME, QTYPE, réponse (RCODE), durée. C'est l'équivalent d'un Wireshark ciblé DNS.

```powershell
# Requêtes vers un nom suspect sur les 30 dernières minutes
$debut = (Get-Date).AddMinutes(-30)
Get-WinEvent -LogName "Microsoft-Windows-DNSServer/Analytical" |
    Where-Object { $_.TimeCreated -ge $debut -and $_.Message -match "appli-metier" } |
    Select-Object TimeCreated, @{n='Détail';e={$_.Message.Substring(0,120)}} | Format-Table -AutoSize
```

