---
id: collect-261001-rattrapage/rattrapage/dns-windows-guide-5
title: "DNS sous Windows Server en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["incident"]
source: docs/RAG/collect-261001-rattrapage/dns_windows_guide.md
source_anchor: ""
source_lines: [528, 677]
sha256: f2a0598adfb50c6afcb7615ce325ea4c31622dad5e15d8991e3549320e7ee2bc
---

# DNS sous Windows Server en entreprise — Guide technique ultra-complet

Un enregistrement = **Nom** + **TTL** + **Classe** (toujours `IN` = Internet) + **Type** + **Données** (+ horodatage si aging actif).

```powershell
# Inspecter un enregistrement en détail
Get-DnsServerResourceRecord -ZoneName "contoso.local" -Name "www" -RRType A |
    Format-List HostName, RecordType, TimeToLive, Timestamp, RecordData
```

- `Timestamp` vide/`0` = enregistrement **statique** (créé à la main, jamais scavengé).
- `Timestamp` renseigné = **dynamique**, soumis au vieillissement (§52).

## 36. A et AAAA

Le cœur du DNS : nom → IPv4 / IPv6.

```powershell
Add-DnsServerResourceRecordA    -Name "www"       -ZoneName "contoso.local" -IPv4Address "10.0.1.50" -TimeToLive 01:00:00
Add-DnsServerResourceRecordAAAA -Name "www"       -ZoneName "contoso.local" -IPv6Address "2001:db8:10::50"
Add-DnsServerResourceRecordA    -Name "intranet"  -ZoneName "contoso.local" -IPv4Address "10.0.1.51","10.0.1.52" -TimeToLive 00:05:00
```

Le 3ᵉ exemple crée **deux A** sur le même nom : round-robin DNS (répartition naïve, sans contrôle de santé — à ne pas confondre avec un vrai load balancer).

**Bonnes pratiques :** un A par usage, nom explicite (`srv-fichiers-01`, pas `serveur2`) ; TTL court (5 min) si l'IP peut changer (bascule), long (1 h+) si stable.

## 37. CNAME : alias et pièges

Un CNAME fait pointer un nom vers un autre nom (pas directement vers une IP).

```powershell
Add-DnsServerResourceRecordCName -Name "intranet" -ZoneName "contoso.local" -HostNameAlias "srv-web-01.contoso.local." -TimeToLive 01:00:00
# Notez le point final : sans lui, le suffixe de zone est ajouté automatiquement
```

**Pièges classiques :**
- Un nom avec CNAME **ne peut avoir aucun autre enregistrement** (pas de MX, pas de TXT, pas de second CNAME). D'où l'interdiction du CNAME au **sommet de zone** (erreur n°5, §127) : l'apex porte déjà SOA et NS.
- Chaînes de CNAME trop longues = résolutions lentes et fragiles. Un seul niveau suffit.
- Pour un alias « nu » d'apex vers l'extérieur, utilisez un A (ou un redirecteur applicatif), pas un CNAME.

## 38. MX : messagerie

Désigne les serveurs de messagerie du domaine, avec **préférence** (plus petit = prioritaire).

```powershell
Add-DnsServerResourceRecordMX -Name "@" -ZoneName "contoso.local" -MailExchange "mx1.contoso.local." -Preference 10 -TimeToLive 01:00:00
Add-DnsServerResourceRecordMX -Name "@" -ZoneName "contoso.local" -MailExchange "mx2.contoso.local." -Preference 20
# "@" = le sommet de zone
```

Chaque `MailExchange` **doit** avoir un A/AAAA (jamais un CNAME, RFC 2181). Vérifiez le reverse (PTR) des MX : beaucoup d'anti-spam le contrôlent.

## 39. SRV : services

Localisation générique d'un service : `_service._proto.nom → cible:port` avec priorité/poids.

```powershell
# Exemple : annuaire LDAP redondant
Add-DnsServerResourceRecordSrv -Name "_ldap._tcp" -ZoneName "contoso.local" `
    -DomainName "srv-ad-01.contoso.local." -Port 389 -Priority 10 -Weight 100
Add-DnsServerResourceRecordSrv -Name "_ldap._tcp" -ZoneName "contoso.local" `
    -DomainName "srv-ad-02.contoso.local." -Port 389 -Priority 10 -Weight 100
Add-DnsServerResourceRecordSrv -Name "_ldap._tcp" -ZoneName "contoso.local" `
    -DomainName "srv-ad-03.contoso.local." -Port 389 -Priority 20 -Weight 0
```

Ici : `srv-ad-01`/`02` se partagent la charge (même priorité, poids égaux), `srv-ad-03` n'est utilisé qu'en secours (priorité 20). C'est exactement le mécanisme qu'AD utilise (§6).

## 40. PTR : résolution inverse

Dans la zone inversée, le **nom** est le dernier octet (ou nibble IPv6), la donnée est le FQDN.

```powershell
# 10.1.0.50 → www.contoso.local (zone 1.10.in-addr.arpa)
Add-DnsServerResourceRecordPtr -Name "50" -ZoneName "1.10.in-addr.arpa" -PtrDomainName "www.contoso.local." -TimeToLive 01:00:00

# Vérification croisée (doit être cohérent dans les deux sens)
Resolve-DnsName -Name "10.1.0.50" -Type PTR
Resolve-DnsName -Name "www.contoso.local" -Type A
```

**Forward-confirmed reverse DNS** : certaines applications exigent que `PTR → nom → A → même IP`. Maintenez les deux zones synchronisées (le DHCP peut le faire, §79).

## 41. TXT : usages

Champ texte libre (255 caractères par chaîne, plusieurs chaînes possibles). Usages en entreprise :

| Usage | Exemple de contenu |
|---|---|
| SPF (anti-spoofing mail) | `v=spf1 mx ip4:10.0.5.25 -all` |
| Vérification de domaine (M365, certificats) | `MS=ms12345678` |
| DKIM/DMARC | `_dmarc.contoso.local` → `v=DMARC1; p=quarantine;` |
| Documentation | `owner=equipe-reseau; ticket=CHG-2026-0912` |

```powershell
Add-DnsServerResourceRecordTXT -Name "@" -ZoneName "contoso.local" -DescriptiveText "v=spf1 mx ip4:10.0.5.25 -all"
```

⚠️ Un TXT « documentation » est visible par **tous** : n'y mettez jamais de secret, mot de passe ou info sensible.

## 42. NS et glue records

Les NS délèguent l'autorité. Chaque zone a des NS ; les **glue records** sont les A/AAAA des NS eux-mêmes quand ils sont *dans* la zone déléguée (sinon, boucle impossible à résoudre).

```powershell
# Ajouter un NS autoritaire à la zone (avec son glue si interne)
Add-DnsServerResourceRecord -Name "@" -NS -ZoneName "contoso.local" -NameServer "SRV-DNS-02.contoso.local."
```

Vérifiez la cohérence : les NS de la zone = les NS chez le parent (délégation) = les serveurs qui répondent vraiment. Un NS « fantôme » (serveur décommissionné resté dans la liste) = timeouts intermittents (cas n°2, §101).

## 43. SOA : le début d'autorité

Un seul SOA par zone, au sommet. Champs :

| Champ | Rôle | Valeur conseillée |
|---|---|---|
| Serial | Version de la zone (déclenche IXFR/AXFR) | `AAAAMMJJNN` (ex : 2026092701) |
| Refresh | Fréquence de vérification du secondaire | 15 min – 2 h |
| Retry | Nouvel essai après échec | 5 – 15 min |
| Expire | Durée max d'utilisation d'une copie périmée | 1 – 4 semaines |
| Minimum (TTL négatif) | Durée de cache des réponses NXDOMAIN | 15 min – 1 h |

```powershell
# Lire le SOA
Get-DnsServerResourceRecord -ZoneName "contoso.local" -RRType SOA | Select-Object -ExpandProperty RecordData
# Modifier le serial (via clone + Set)
$old = Get-DnsServerResourceRecord -ZoneName "contoso.local" -RRType SOA
$new = $old.Clone()
$new.RecordData.SerialNumber = 2026092701
Set-DnsServerResourceRecord -NewInputObject $new -OldInputObject $old -ZoneName "contoso.local"
```

Incrémentez le serial à chaque modification manuelle d'une zone **standard** (les zones AD le font automatiquement).

## 44. CAA : autorité de certification

Un CAA déclare quelle autorité de certification a le droit d'émettre des certificats pour le domaine. Les CA publiques sont tenues de le vérifier.

```powershell
# Seule Let's Encrypt peut émettre pour contoso.local ; aucun wildcard
Add-DnsServerResourceRecord -CAA -Name "@" -ZoneName "contoso.local" -Flags 0 -Tag "issue" -Value "letsencrypt.org"
Add-DnsServerResourceRecord -CAA -Name "@" -ZoneName "contoso.local" -Flags 0 -Tag "issuewild" -Value ";"
```

Tags : `issue` (certificats simples), `issuewild` (wildcards), `iodef` (URL de rapport d'incident). En interne avec une CA d'entreprise, le CAA est moins critique mais reste une bonne hygiène pour les zones exposées.

## 45. Les enregistrements SRV d'AD : liste complète

Référence exhaustive de ce que Netlogon enregistre pour `contoso.local` (zone `_msdcs.contoso.local` + zone `contoso.local`) :

