---
id: collect-261001-rattrapage/rattrapage/dhcp-windows-guide-9
title: "Guide technique ultra-complet : DHCP sous Windows Server en entreprise"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-rattrapage/dhcp_windows_guide.md
source_anchor: ""
source_lines: [1387, 1555]
sha256: e482a1c0627bdaf35c19fa655cc61ed335750ee157906d5897dd6cf0ba99244f
---

# 1. Options de niveau serveur
$serverOptions = Get-DhcpServerv4OptionValue -ComputerName $primary
foreach ($opt in $serverOptions) {
    # Supprimer puis recréer côté secondaire (idempotent)
    Remove-DhcpServerv4OptionValue -ComputerName $secondary -OptionId $opt.OptionId -ErrorAction SilentlyContinue
}
# (recréation : à adapter selon tes options — exemple DNS + domaine)
Set-DhcpServerv4OptionValue -ComputerName $secondary `
    -DnsServer 192.0.2.53, 192.0.2.54 -DnsDomain "contoso.local" -OptionId 42 -Value 192.0.2.53

# 2. Filtres MAC
Set-DhcpServerv4FilterList -ComputerName $secondary `
    -Allow (Get-DhcpServerv4FilterList -ComputerName $primary).Allow `
    -Deny  (Get-DhcpServerv4FilterList -ComputerName $primary).Deny
Get-DhcpServerv4Filter -ComputerName $primary -List Allow | ForEach-Object {
    Add-DhcpServerv4Filter -ComputerName $secondary -List Allow `
        -MacAddress $_.MacAddress -Description $_.Description -ErrorAction SilentlyContinue
}

Write-Output "Synchronisation primaire -> secondaire terminee. Verifier avec :"
Write-Output "  Get-DhcpServerv4OptionValue -ComputerName $secondary"
```

## 64. DHCP sur plusieurs VLANs : l'agent de relais (relay agent)

**Le problème** : le DHCPDISCOVER est un **broadcast** (255.255.255.255). Les routeurs ne forwardent pas les broadcasts → sans relais, seul le VLAN du serveur DHCP obtient des adresses.

**La solution** : un **agent de relais** sur chaque VLAN/routeur qui convertit le broadcast en **unicast** vers le(s) serveur(s) DHCP, en ajoutant l'info du sous-réseau d'origine (option 82 / champ `giaddr`).

```text
Client VLAN 20 (192.0.3.0/24)
   │ DHCPDISCOVER (broadcast)
   ▼
┌──────────────────────────────┐
│ Routeur / Switch L3 VLAN 20  │
│  ip helper-address 192.0.2.10│  ← agent de relais
│  ip helper-address 192.0.2.11│  ← vers les 2 serveurs !
└──────────────┬───────────────┘
               │ unicast UDP 67, giaddr=192.0.3.1
               ▼
        Serveur DHCP (VLAN 10)
        → choisit l'étendue 192.0.3.0 grâce au giaddr
```

Le serveur utilise le champ **giaddr** (adresse de l'agent de relais) pour savoir **dans quelle étendue** piocher. D'où l'importance d'**une étendue par VLAN**.

## 65. Configurer l'agent de relais — côté routeur/switch (exemples constructeurs)

**Cisco IOS** :

```text
interface Vlan20
 description WiFi-Corporate
 ip address 192.0.3.1 255.255.255.0
 ip helper-address 192.0.2.10
 ip helper-address 192.0.2.11
```

**HP/Aruba (Comware)** :

```text
interface Vlan-interface20
 ip address 192.0.3.1 255.255.255.0
 dhcp select relay
 dhcp relay server-address 192.0.2.10
 dhcp relay server-address 192.0.2.11
```

> ⚠️ Toujours déclarer **les deux serveurs DHCP** dans le helper : si tu ne mets que le primaire et qu'il tombe, le VLAN n'a plus de DHCP malgré le basculement.

## 66. Configurer l'agent de relais — Windows Server (RRAS)

Si un serveur Windows fait office de routeur (RRAS), on peut y activer le relais DHCP :

```powershell
# Installer RRAS (routage uniquement, sans NAT/VPN si non nécessaire)
Install-WindowsFeature -Name Routing -IncludeManagementTools

# Via la console RRAS (rrasmgmt.msc) :
#   IPv4 -> Général -> clic droit -> Nouveau protocole de routage -> Agent de relais DHCP
#   Clic droit sur "Agent de relais DHCP" -> Propriétés -> ajouter 192.0.2.10 et 192.0.2.11
#   Clic droit -> Nouvelle interface -> choisir l'interface du VLAN client
```

> 💡 En pratique d'entreprise, le relais est sur les **switchs L3/routeurs**, pas sur Windows. RRAS = solution de secours ou labo.

## 67. Vérifier que le relais fonctionne

```powershell
# Côté serveur DHCP : les requêtes relayées apparaissent avec le giaddr
# Activer le log détaillé temporairement et filtrer :
Get-WinEvent -LogName "Microsoft-Windows-DHCP-Server/Operational" -MaxEvents 50 |
    Where-Object { $_.Message -like "*giaddr*" } |
    Select-Object TimeCreated, Message | Format-List

# Test client : depuis un poste du VLAN distant
# ipconfig /release puis ipconfig /renew
# Doit obtenir une IP de la bonne étendue (192.0.3.x pour le VLAN 20)
```

Si le client n'obtient rien : vérifier dans l'ordre (cas pratique n°4, section 80) :

1. Le helper est-il configuré sur la bonne interface VLAN ?
2. Le pare-feu laisse-t-il passer UDP 67/68 entre relais et serveurs ?
3. L'étendue du VLAN existe-t-elle et est-elle active ?
4. Y a-t-il des adresses libres ?

## 68. Bonnes pratiques — HA et multi-VLAN (récap)

1. ✅ 2 serveurs DHCP en Load Balance 50/50 (ou Hot Standby justifié).
2. ✅ Secret partagé robuste, stocké en coffre, rotation annuelle.
3. ✅ IP helper vers **les deux** serveurs sur chaque VLAN.
4. ✅ Une étendue par VLAN, même nommage (`VLAN20-WiFi-Corp`).
5. ✅ Surveillance de l'état du failover (`State = Normal`).
6. ✅ Script de synchro des paramètres non répliqués (section 63).
7. ❌ Jamais de helper vers un seul serveur.
8. ❌ Jamais deux étendues avec le même ScopeId sur des serveurs non en failover.

---

# Bloc F — DNS dynamique, sauvegarde, migration

## 69. DNS dynamique : pourquoi le DHCP doit mettre à jour le DNS

Sans DNS dynamique : un poste obtient `192.0.2.87` mais `poste-123.contoso.local` ne résout pas → impossible de joindre les machines par leur nom, la supervision échoue, les GPO ciblées par nom patinent.

Avec DNS dynamique : à chaque bail, le serveur DHCP **crée/met à jour** l'enregistrement A (et PTR) dans la zone AD.

```powershell
# Activer les mises à jour DNS dynamiques sur une étendue
Set-DhcpServerv4Scope -ScopeId 192.0.2.0 -DynamicDnsUpdateEnabled $true

# Vérifier
Get-DhcpServerv4Scope -ScopeId 192.0.2.0 | Select-Object Name, DynamicDnsUpdateEnabled
```

## 70. Configuration fine du DNS dynamique (onglet DNS de l'étendue)

Options disponibles (console → propriétés de l'étendue → onglet **DNS**, ou PowerShell) :

| Option | Recommandation |
|--------|----------------|
| Activer les mises à jour DNS dynamiques | ✅ Oui |
| Toujours mettre à jour (même si le client ne demande pas) | ✅ Oui (clients non-Windows, MFP, téléphones) |
| Supprimer l'enregistrement quand le bail expire/supprimé | ✅ Oui (évite les DNS fantômes) |
| Nom de protection (DHCID) | ✅ Oui (voir section 72) |

```powershell
# Configuration complète en PowerShell
Set-DhcpServerv4DnsSetting `
    -ComputerName "srv-dhcp-01.contoso.local" `
    -DynamicUpdates "Always" `
    -DeleteDnsRRonLeaseExpiry $true `
    -UpdateDnsRRForOlderClients $true `
    -DisableDynamicUpdateForDnsPTRRecord $false `
    -NameProtection $true

# Vérifier
Get-DhcpServerv4DnsSetting -ComputerName "srv-dhcp-01.contoso.local" | Format-List *
```

> **Lien métier copieurs** : les MFP ne mettent généralement **pas** à jour le DNS eux-mêmes (pas de client DNS dynamique). Avec *"Toujours mettre à jour"*, le serveur DHCP crée `mfp-compta-rdc.contoso.local` → tu peux joindre le copieur par son nom depuis les postes et la supervision. **Indispensable** pour un parc MFP propre.

## 71. Compte de service dédié pour les mises à jour DNS (sécurisation)

Par défaut, le DHCP met à jour le DNS avec le **compte machine** du serveur. Problème : si plusieurs serveurs DHCP (failover) mettent à jour les mêmes enregistrements, ou si un enregistrement est créé par un client, les permissions se mélangent → erreurs, enregistrements non supprimés.

**Bonne pratique** : créer un **compte de service dédié** (utilisateur AD simple, mot de passe robuste, *"le mot de passe n'expire jamais"* + documenté en coffre) :

