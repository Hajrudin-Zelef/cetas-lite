---
id: collect-261001-rattrapage/rattrapage/dns-windows-guide-6
title: "DNS sous Windows Server en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["incident"]
source: docs/RAG/collect-261001-rattrapage/dns_windows_guide.md
source_anchor: ""
source_lines: [678, 827]
sha256: e7a67fe464d4087f8cc558bbc1c20675ce2a81a42ff58b247c32a407091c0223
---

# DNS sous Windows Server en entreprise — Guide technique ultra-complet

```
# --- Zone _msdcs.contoso.local ---
_ldap._tcp.dc._msdcs               SRV 0 100 389  SRV-DC-01.contoso.local
_ldap._tcp.<Site>._sites.dc._msdcs SRV 0 100 389  SRV-DC-01.contoso.local   (par site)
_kerberos._tcp.dc._msdcs           SRV 0 100 88   SRV-DC-01.contoso.local
_kerberos._tcp.<Site>._sites.dc._msdcs
_kerberos._udp.dc._msdcs           SRV 0 100 88   SRV-DC-01.contoso.local
_kpasswd._tcp / _kpasswd._udp      SRV 0 100 464  SRV-DC-01.contoso.local
_gc._tcp                           SRV 0 100 3268 SRV-DC-01.contoso.local  (si GC)
_gc._tcp.<Site>._sites
_ldap._tcp.pdc._msdcs               SRV 0 100 389  <émulateur PDC>
_ldap._tcp.DomainDnsZones / ForestDnsZones  SRV 0 100 389 ... (DC hébergeant ces partitions)
<GUID>._msdcs                      CNAME SRV-DC-01.contoso.local           (réplication)

# --- Zone contoso.local ---
_ldap._tcp / _kerberos._tcp / _kpasswd._tcp / _gc._tcp   (SRV génériques, sans _msdcs)
contoso.local                      A     10.0.0.10 / 10.0.0.11 ...          (A de chaque DC)
```

```powershell
# Forcer le ré-enregistrement (après un incident)
nltest /dsregdns
# Sur tous les DC du domaine (à exécuter par DC ou via Invoke-Command)
Invoke-Command -ComputerName "SRV-DC-01","SRV-DC-02" -ScriptBlock { nltest /dsregdns }
```

## 46. Créer des enregistrements en PowerShell : synthèse

Tableau récapitulatif des applets (toutes acceptent `-TimeToLive` en `[TimeSpan]` et `-AgeRecord` pour activer le vieillissement sur l'enregistrement) :

| Besoin | Applet |
|---|---|
| A / AAAA | `Add-DnsServerResourceRecordA` / `AAAA` |
| CNAME | `Add-DnsServerResourceRecordCName` |
| MX | `Add-DnsServerResourceRecordMX` |
| SRV | `Add-DnsServerResourceRecordSrv` |
| PTR | `Add-DnsServerResourceRecordPtr` |
| TXT | `Add-DnsServerResourceRecordTXT` |
| NS / SOA / CAA / DS… | `Add-DnsServerResourceRecord -NS/-CAA…` (générique) |

```powershell
# Lister tous les enregistrements d'une zone, triés par type
Get-DnsServerResourceRecord -ZoneName "contoso.local" |
    Sort-Object RecordType, HostName |
    Format-Table HostName, RecordType,
        @{n='TTL';e={$_.TimeToLive.ToString()}},
        @{n='Données';e={$_.RecordData.ToString()}} -AutoSize

# Supprimer proprement (préciser la donnée quand plusieurs RR coexistent)
Remove-DnsServerResourceRecord -ZoneName "contoso.local" -Name "intranet" -RRType A -RecordData "10.0.1.51" -Force
```

---

# Chapitre 5 — Transferts de zone AXFR/IXFR

## 47. AXFR vs IXFR

| | AXFR (Full) | IXFR (Incremental) |
|---|---|---|
| Principe | Copie **complète** de la zone | Seulement les **différences** depuis un serial |
| Déclenchement | Premier transfert, serial trop ancien, demande explicite | Vérification périodique (Refresh SOA) |
| Protocole | TCP 53 | TCP 53 |
| Coût | Lourd (toute la zone) | Léger |

Le secondaire compare son serial SOA à celui du maître : s'il est inférieur, il demande un IXFR ; si le maître ne gère pas l'IXFR (ou si l'écart est trop grand), repli en AXFR. Windows Server gère les deux nativement.

## 48. Configurer le transfert de zone (GUI)

Sur la zone **primaire** : Propriétés → onglet **Transferts de zone** → cocher « Autoriser les transferts de zone » → choisir :
- **Vers tout serveur** : ⚠️ jamais en production (fuite d'information : un AXFR liste toute la zone).
- **Uniquement vers les serveurs listés dans l'onglet Serveurs de noms** : le bon choix courant.
- **Uniquement vers les serveurs suivants** : liste explicite d'IP (le plus strict).

Onglet **Serveurs de noms** : y déclarer chaque secondaire autorisé (avec son glue si besoin).

## 49. Configurer en PowerShell / dnscmd

```powershell
# Autoriser uniquement vers des IP explicites (secondaires 10.0.0.11 et 192.168.99.10)
Set-DnsServerZoneTransfer -Name "contoso.local" -SecondaryServers @("10.0.0.11","192.168.99.10") -PassThru

# Vérifier
Get-DnsServerZoneTransfer -Name "contoso.local" | Format-List *

# Notifier ces serveurs après chaque changement
Set-DnsServerPrimaryZone -Name "contoso.local" -NotifyServers @("10.0.0.11","192.168.99.10")

# Côté secondaire : déclarer le maître et forcer le transfert
Set-DnsServerSecondaryZone -Name "contoso.local" -MasterServers "10.0.0.10"
Sync-DnsServerZone -Name "contoso.local"
```

`dnscmd` équivalent : `dnscmd SRV-DNS-01 /ZoneResetSecondaries contoso.local /SecureList 10.0.0.11 192.168.99.10`.

## 50. Sécuriser les transferts

Checklist de sécurisation :
1. **Liste blanche d'IP** (§49), jamais « tout serveur ».
2. **Pare-feu** : TCP 53 ouvert uniquement vers les secondaires autorisés.
3. **TSIG** : si le secondaire est BIND/Linux, signez les transferts (clé partagée) — côté Windows, la sécurisation repose surtout sur la liste d'IP et le pare-feu.
4. **Zones AD-integrated** : pas de transfert nécessaire entre DC (réplication AD) ; n'activez les transferts que si un secondaire non-DC existe.
5. **Audit** : un AXFR inattendu dans les logs = signal d'alerte (reconnaissance).

```powershell
# Tester qu'un serveur NON autorisé est bien refusé (depuis ce dernier)
# nslookup interactif : ls contoso.local  → doit échouer (refusé)
```

## 51. Notify, SOA serial et cohérence

Le **NOTIFY** (RFC 1996) : quand le primaire change, il *pousse* une notification aux secondaires listés, qui viennent chercher l'IXFR sans attendre le Refresh SOA. Windows l'envoie automatiquement aux serveurs autorisés.

```powershell
# Voir / configurer les serveurs notifiés
Get-DnsServerZone -Name "contoso.local" | Select-Object -ExpandProperty NotifyServers
Set-DnsServerPrimaryZone -Name "contoso.local" -NotifyServers @("10.0.0.11","192.168.99.10")
```

**Cohérence :** après une modification, vérifiez que le serial a augmenté **partout** :
```powershell
"10.0.0.10","10.0.0.11" | ForEach-Object {
    [pscustomobject]@{ Serveur=$_; Serial=(Resolve-DnsName contoso.local -Type SOA -Server $_).SerialNumber }
}
```
Un serial qui ne bouge pas sur le secondaire = transfert en échec (pare-feu, liste blanche, maître injoignable) → cas pratique n°7 (§106).

---

# Chapitre 6 — Vieillissement (aging) et scavenging

## 52. Le problème des enregistrements périmés

Avec les mises à jour dynamiques (DHCP, Netlogon), les enregistrements A/PTR sont créés et rafraîchis automatiquement. Mais quand une machine est **retirée du parc** (mise au rebut, renommée, VM supprimée), son enregistrement reste : la zone se remplit de **noms fantômes** qui pointent vers des IP réattribuées → connexions vers la mauvaise machine, inventaires faux, alertes parasites.

Le **scavenging** est le ramasse-miettes automatique du DNS Windows : il supprime les enregistrements dynamiques périmés. Mal configuré, il supprime aussi des enregistrements **valides** (cas n°5, §104). Ce chapitre est donc à lire deux fois.

## 53. No-refresh et refresh intervals : le mécanisme

Chaque enregistrement dynamique porte un **horodatage** (timestamp). Deux intervalles protègent contre les réplications inutiles :

```
Timeline d'un enregistrement dynamique (ex : no-refresh 7 j, refresh 7 j)
─────────────────────────────────────────────────────────────────────────
J0 : création (timestamp = J0)
J0 ──► J7   : NO-REFRESH — les refresh clients sont IGNORÉS (pas de réplication AD)
J7 ──► J14  : REFRESH — un refresh client met à jour le timestamp (répliqué)
J14+         : périmé — éligible au scavenging (supprimé au prochain cycle)
```

