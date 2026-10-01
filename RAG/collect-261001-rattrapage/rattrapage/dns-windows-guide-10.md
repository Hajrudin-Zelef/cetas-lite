---
id: collect-261001-rattrapage/rattrapage/dns-windows-guide-10
title: "DNS sous Windows Server en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-rattrapage/dns_windows_guide.md
source_anchor: ""
source_lines: [1219, 1343]
sha256: bf7cc917efeccbc052ec483b45732cf3ffc6b37ef4da80044520b7704cf07d03
---

# DNS sous Windows Server en entreprise — Guide technique ultra-complet

| Client | A | PTR |
|---|---|---|
| Windows (modèle A) | Le client | Le DHCP |
| Non-Windows / modèle B | Le DHCP (avec credentials) | Le DHCP |

```powershell
# Voir la config DNS d'une étendue DHCP (module DhcpServer)
Get-DhcpServerv4DnsSetting -ComputerName "SRV-DHCP-01"
```

## 80. Configurer les credentials DHCP

Sans compte dédié, le DHCP enregistre avec son **compte machine** : si plusieurs serveurs DHCP coexistent (ou après une réinstallation), les enregistrements deviennent orphelins (le nouveau serveur n'a pas le droit de les écraser). La bonne pratique : un **compte de service dédié**, mot de passe géré, membre du groupe **DnsUpdateProxy**.

```powershell
# 1. Créer le compte de service (adaptez selon votre nommage)
New-ADUser -Name "svc-dhcp-dns" -SamAccountName "svc-dhcp-dns" -AccountPassword (Read-Host -AsSecureString) `
    -Enabled $true -PasswordNeverExpires $true -Description "Mises à jour DNS pour DHCP"

# 2. L'ajouter au groupe DnsUpdateProxy du domaine
Add-ADGroupMember -Identity "DnsUpdateProxy" -Members "svc-dhcp-dns"

# 3. Déclarer les credentials sur chaque serveur DHCP (GUI : propriétés IPv4 → onglet Avancé → Informations d'identification)
Set-DhcpServerDnsCredential -Credential (Get-Credential "CONTOSO\svc-dhcp-dns") -ComputerName "SRV-DHCP-01"

# 4. Vérifier
Get-DhcpServerDnsCredential -ComputerName "SRV-DHCP-01"
```

⚠️ Le compte doit avoir un mot de passe **qui n'expire pas** (ou géré via gMSA — voir §149) sinon les mises à jour s'arrêtent le jour de l'expiration.

## 81. Le groupe DnsUpdateProxy : utilité et danger

**Utilité :** les membres de `DnsUpdateProxy` peuvent créer des enregistrements **sans en devenir propriétaires** : n'importe quel serveur DHCP du groupe peut ensuite les mettre à jour. Sans ça, le premier DHCP qui enregistre une imprimante verrouille l'enregistrement (ACL) et les autres ne peuvent plus le toucher.

**Danger :** les enregistrements créés par un membre de DnsUpdateProxy sont **non sécurisés** (pas d'ACL propriétaire). Conséquences :
- **Ne mettez JAMAIS un contrôleur de domaine dans DnsUpdateProxy** (erreur n°17, §139) : ses enregistrements deviendraient modifiables par tous.
- Limitez le groupe au strict nécessaire (comptes DHCP dédiés).
- Combinez avec le scavenging (§56) : les enregistrements non rafraîchis seront nettoyés.

## 82. Protéger les enregistrements existants

Quand vous basculez un DHCP vers le modèle B (ou changez de compte), les **anciens** enregistrements appartiennent à l'ancien propriétaire → le nouveau compte ne peut pas les mettre à jour → doublons et conflits.

Procédure de bascule propre :
1. Dans la console DNS, affichez la colonne **Horodatage**.
2. Pour les enregistrements critiques (serveurs en DHCP — à éviter, mais ça existe), passez-les en **statiques** : décochez « Supprimer cet enregistrement lorsqu'il devient périmé » — ou recréez-les à la main.
3. Option radicale mais propre sur une zone de test : `dnscmd /AgeAllRecords` puis laissez le scavenging nettoyer, **après** sauvegarde (§143).

## 83. Scavenging et DHCP : le duo gagnant

Le cycle vertueux : bail DHCP 8 jours → no-refresh 7 j + refresh 7 j → scavenging 7 j (§53). Le DHCP rafraîchit les enregistrements à chaque renouvellement de bail ; le scavenging supprime ceux des machines parties.

**Point de vigilance :** les équipements à bail **long ou réservation** (imprimantes, onduleurs, caméras) : si leur enregistrement est dynamique et que l'équipement ne rafraîchit pas (firmware basique), il sera scavengé. Pour ces équipements : **réservation DHCP + enregistrement statique** (ou modèle B avec credentials, et vérifiez le rafraîchissement).

## 84. Cas des clients non-Windows et statiques

- **Linux/macOS** : peuvent s'enregistrer via `nsupdate` + Kerberos (krb5) — en pratique, laissez le DHCP faire le travail (modèle B).
- **Équipements réseau** (imprimantes, onduleurs, switchs) : enregistrements **statiques** créés à la main, documentés dans l'inventaire. Pas de scavenging dessus (timestamp à 0).
- **Serveurs** : **IP fixe + A statique**, jamais de DHCP (sauf réservation + documentation). Un serveur dont l'IP change au gré du DHCP est une panne en puissance.

```powershell
# Créer un enregistrement statique (pas de timestamp → jamais scavengé)
Add-DnsServerResourceRecordA -Name "prt-compta-01" -ZoneName "contoso.local" -IPv4Address "10.1.5.30" -TimeToLive 01:00:00
# Vérifier qu'il est statique : Timestamp doit être vide
Get-DnsServerResourceRecord -ZoneName "contoso.local" -Name "prt-compta-01" | Select-Object HostName, Timestamp
```

---

# Chapitre 11 — Haute disponibilité

## 85. Plusieurs serveurs DNS : architecture type

Minimum viable en entreprise : **2 serveurs DNS** par site important, sur **2 hyperviseurs / 2 baies** distincts si virtualisés.

```
Site Paris-Siege :  SRV-DNS-01 (10.0.0.10)  +  SRV-DNS-02 (10.0.0.11)
Site Lyon-Usine  :  SRV-DNS-03 (10.2.0.10)  +  SRV-DNS-04 (10.2.0.11)
```

- Les clients ont **2 DNS** (DHCP options 006, ou GPO) : le client interroge le second si le premier ne répond pas (timeout ~1-2 s, variable selon l'OS).
- Chaque DC est DNS (zones AD-integrated) : la redondance DNS suit la redondance AD.
- ⚠️ Les clients n'équilibrent **pas** la charge entre leurs 2 DNS : ils utilisent le premier tant qu'il répond. Pour répartir, variez l'ordre via DHCP par site.

## 86. AD-integrated : la HA naturelle

Avec des zones AD-integrated, **chaque DC/DNS est un primaire** : pas de maître unique, pas de transfert à configurer, écriture possible partout. La réplication AD (15 s intra-site par défaut, compressée, chiffrée) propage les changements.

```powershell
# Vérifier que la zone est bien présente et saine sur chaque DC/DNS
$serveurs = "SRV-DNS-01","SRV-DNS-02","SRV-DNS-03"
foreach ($s in $serveurs) {
    Get-DnsServerZone -Name "contoso.local" -ComputerName $s |
        Select-Object @{n='Serveur';e={$s}}, ZoneName, ZoneType, IsDsIntegrated
}
```

**Limite assumée :** si AD est en panne (réplication cassée), le DNS multi-maîtres diverge. C'est le prix de l'intégration — d'où l'importance de superviser la réplication AD (`repadmin /showrepl`, `dcdiag /test:dns`).

## 87. Secondaires et transferts pour la HA

Quand un DNS **n'est pas** DC (DMZ, site sans AD, appliance), utilisez des **zones secondaires** (§25) : le primaire notifie (§51), le secondaire transfère.

```powershell
# Sur le secondaire de DMZ : vérifier la fraîcheur (serial identique au primaire)
$primaire   = (Resolve-DnsName "dmz.contoso.local" -Type SOA -Server "10.0.0.10").SerialNumber
$secondaire = (Resolve-DnsName "dmz.contoso.local" -Type SOA -Server "192.168.99.10").SerialNumber
if ($primaire -ne $secondaire) { Write-Warning "DIVERGENCE : primaire=$primaire secondaire=$secondaire" }
```

**Expire SOA** (§43) : sur un secondaire isolé longtemps, la zone expire → le serveur **cesse de répondre** pour cette zone. Réglez l'Expire à 2–4 semaines, pas moins.

## 88. Anycast DNS : principe

L'**anycast** : plusieurs serveurs DNS partagent la **même IP**, annoncée via BGP depuis plusieurs sites ; le routage envoie chaque client au plus proche. Si un site tombe, le trafic bascule automatiquement (convergence BGP, quelques secondes à minutes).

- En entreprise, c'est pertinent pour des **résolveurs centraux** multi-sites ou des zones publiques critiques.
- Windows Server ne fait pas de BGP natif : l'anycast se met en place avec des routeurs/pare-feu en frontal (ou un ADC). Mentionné ici pour que le terme ne soit pas un mystère en réunion d'architecture.
- Alternative simple sans BGP : **2 IP virtuelles** (une par site) + bascule manuelle documentée.

## 89. Supervision de la disponibilité DNS

Ce qu'il faut monitorer (Zabbix, SCOM, PRTG, ou script maison) :

