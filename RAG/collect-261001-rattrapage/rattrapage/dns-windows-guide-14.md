---
id: collect-261001-rattrapage/rattrapage/dns-windows-guide-14
title: "DNS sous Windows Server en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/dns_windows_guide.md
source_anchor: ""
source_lines: [1778, 1912]
sha256: ab6eb6372ce6474b38fa309317a4b1ca59669d9b0faec1755b5d7dc3c8d10cd7
---

# DNS sous Windows Server en entreprise — Guide technique ultra-complet

**Symptômes :** après changement d'IP d'un DC (ou migration de serveur), des clients tentent encore l'ancienne IP ; réplication en erreur.

**Checklist post-changement :**
```powershell
# 1. Les A du DC sont-ils à jour ? (domaine nu + nom du DC + _msdcs)
Get-DnsServerResourceRecord -ZoneName "contoso.local" -RRType A |
    Where-Object { $_.RecordData.IPv4Address -eq "10.0.0.99" }  # ancienne IP
# 2. Supprimer les A obsolètes (préciser -RecordData)
# 3. Forcer le ré-enregistrement
Invoke-Command -ComputerName "SRV-DC-03" -ScriptBlock { nltest /dsregdns; ipconfig /registerdns }
# 4. Vider les caches (serveurs + clients critiques)
Clear-DnsServerCache -ComputerName "SRV-DNS-01","SRV-DNS-02" -Force
# 5. Vérifier la réplication AD vers le DC
repadmin /showrepl SRV-DC-03
```
**Prévention :** avant toute migration, baissez les TTL à 300 s (§12), planifiez le scavenging, et gardez l'ancienne IP en alias le temps de la propagation (TTL × 2).

---

# Chapitre 13 — Sécurité

## 115. Surface d'attaque du DNS

Le DNS est une cible de choix : il voit **tout** le trafic de nommage, il est interrogé par **tout** le monde, et une réponse falsifiée redirige silencieusement les utilisateurs.

| Menace | Impact | Contre-mesure (section) |
|---|---|---|
| Cache poisoning | Redirection vers un faux serveur | Cache locking (§119), DNSSEC (§69), ports source aléatoires |
| Récursion ouverte | Amplification DDoS, exfiltration | Restreindre la récursion (§117) |
| AXFR sauvage | Cartographie complète du SI | Liste blanche de transferts (§50) |
| Détournement de mises à jour dynamiques | Enregistrements forgés | Updates sécurisés uniquement (§78) |
| Tunneling DNS (exfiltration via requêtes) | Fuite de données | Filtrage, analytique DNS (§97), pare-feu |
| Escalade via DnsAdmins | Contrôle du DC | Ne pas donner DnsAdmins (§120) |
| DGA / C2 de malware | Communication malware | Monitoring des QNAME (§97), DNS filtrant |

## 116. Cache poisoning et protections

L'attaque : un attaquant devine/forge une réponse à une requête récursive en cours et **empoisonne le cache** du résolveur pour la durée du TTL.

Défenses en profondeur (toutes activables sur Windows Server) :
1. **Ports source UDP aléatoires** (activés par défaut depuis 2008) : rend la devinette beaucoup plus dure.
2. **Cache locking** (§119) : une entrée en cache ne peut pas être écrasée avant un pourcentage de son TTL.
3. **0x20 encoding** (casse aléatoire) : le résolveur Windows mélange minuscules/majuscules dans la question et vérifie la réponse.
4. **DNSSEC** (§69) : la parade cryptographique définitive pour les zones signées.
5. Ne jamais exposer la récursion (§117).

```powershell
# Vérifier que le verrouillage du cache est actif (100 % par défaut)
Get-DnsServerCache | Select-Object LockingPercent, MaxTTL, MaxNegativeTTL
```

## 117. Restreindre la récursion

Un serveur **autoritaire public** (qui héberge vos zones Internet) ne doit **pas** faire de récursion pour le monde entier : c'est la porte ouverte à l'amplification DDoS (petite requête → grosse réponse, avec IP source usurpée).

```powershell
# Désactiver totalement la récursion (serveur autoritaire exposé)
Set-DnsServerRecursion -Enable $false

# Ou : récursion uniquement pour vos réseaux (2016+, recursion scopes)
Add-DnsServerRecursionScope -Name "Interne" -EnableRecursion $true
Add-DnsServerClientSubnet -Name "Subnet-LAN" -IPv4Subnet "10.0.0.0/8"
Add-DnsServerQueryResolutionPolicy -Name "Recursion-Interne" -Action ALLOW `
    -ClientSubnet "EQ,Subnet-LAN" -RecursionScope "Interne" -PassThru
```

**Test externe :** depuis Internet, `nslookup www.example.com <votre-IP-publique>` doit **échouer** (refusé) sur un autoritaire pur, **réussir** sur un résolveur interne (depuis le LAN uniquement).

## 118. Response Rate Limiting (RRL)

Le RRL limite le nombre de réponses **identiques** envoyées à un même client : il casse l'amplification DDoS (l'attaquant ne reçoit plus qu'un filet de réponses) sans pénaliser les clients légitimes.

```powershell
# Activer le RRL (Windows Server 2016+)
Set-DnsServerResponseRateLimiting -Enable $true -ResponsesPerSec 5 -WindowInSec 5 `
    -LeakRate 2 -MaxResponses 50 -PassThru

# Vérifier
Get-DnsServerResponseRateLimiting | Format-List Enable, ResponsesPerSec, WindowInSec
```

⚠️ Version : RRL = **2016 et +**. Réglages : commencez permissif (`ResponsesPerSec 10`), observez les faux positifs (un NAT qui concentre 500 postes derrière une IP peut déclencher le RRL — whitelistez vos NAT sortants si besoin).

## 119. Cache locking

Le verrouillage du cache empêche l'écrasement d'une entrée avant qu'un pourcentage de son TTL se soit écoulé. À 100 % (défaut), une entrée empoisonnée par erreur ne peut pas être « corrigée » par une réponse légitime avant expiration — mais une entrée légitime ne peut pas non plus être empoisonnée.

```powershell
# 100 % = protection maximale (défaut, recommandé)
Set-DnsServerCache -LockingPercent 100 -PassThru
# En dépannage uniquement : baisser temporairement pour forcer un rafraîchissement
Set-DnsServerCache -LockingPercent 0; Clear-DnsServerCache -Force; Set-DnsServerCache -LockingPercent 100
```

## 120. Le groupe DnsAdmins : risque d'escalade

**Fait peu connu, critique :** un membre du groupe **DnsAdmins** peut, sur un DC qui est aussi DNS, charger une DLL arbitraire comme plugin du service DNS (`ServerLevelPluginDll`) → exécution de code **SYSTEM** sur le DC → compromission du domaine. (Vecteur documenté depuis 2017, toujours d'actualité par design.)

**Règles :**
1. **Ne mettez personne dans DnsAdmins** en routine. Les admin DNS quotidiens n'en ont pas besoin.
2. Déléguez via des **ACL fines** : les opérateurs peuvent gérer les enregistrements d'une zone sans être DnsAdmins (clic droit zone → Propriétés → Sécurité).
3. Auditez le groupe régulièrement :
```powershell
Get-ADGroupMember -Identity "DnsAdmins" | Select-Object Name, SamAccountName
```
4. Si un DC est DNS (cas général), considérez DnsAdmins comme équivalent à « Domain Admins » dans votre modèle de menace.

## 121. Durcissement : liste de mesures

- [ ] Récursion désactivée ou restreinte sur tout serveur exposé (§117)
- [ ] Transferts de zone en liste blanche (§50)
- [ ] Mises à jour dynamiques en `Secure` sur toutes les zones AD (§78)
- [ ] Cache locking à 100 % (§119), RRL activé sur les exposés (§118)
- [ ] Groupe DnsAdmins vide ou quasi-vide, audité (§120)
- [ ] Journal **Audit** DNS activé en permanence (§96)
- [ ] `cache.dns` (root hints) à jour : `Get-DnsServerRootHint` vs liste officielle
- [ ] Pas de DNS externe dans la carte des DC (§20)
- [ ] Global Query Block List active (bloque `wpad`/`isatap` par défaut — anti-hijacking WPAD) :
```powershell
Get-DnsServerGlobalQueryBlockList | Select-Object -ExpandProperty ListName
```
- [ ] Mises à jour Windows à jour sur les DNS (le service DNS a eu des CVE critiques, ex : SIGRed CVE-2020-1350 — 17 ans de vulnérabilité RCE wormable)
- [ ] Comptes de service DHCP dédiés, mots de passe gérés (§80)

## 122. DNS et pare-feu : flux réseau

| Flux | Ports | Sens | Commentaire |
|---|---|---|---|
| Clients → DNS interne | UDP/TCP 53 | Entrant | TCP pour les grosses réponses/DNSSEC |
| DNS → redirecteurs / Internet | UDP/TCP 53 | Sortant | Restreignez aux redirecteurs si vous en utilisez |
| DNS ↔ DNS (transferts) | TCP 53 | Bilatéral | AXFR/IXFR = TCP uniquement |
| DC ↔ DC (AD-integrated) | 53 + ports AD (389, 636, 3268, 445, 135, RPC dyn.) | Bilatéral | La réplication de zone passe par AD |
| DHCP → DNS (updates) | UDP/TCP 53 | Sortant DHCP | Mises à jour dynamiques |
| Admin → DNS | RPC/WMI/WinRM | Entrant | Console DNS, PowerShell distant |

