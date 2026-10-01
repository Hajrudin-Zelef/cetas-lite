---
id: collect-261001-rattrapage/rattrapage/dhcp-windows-guide-11
title: "Guide technique ultra-complet : DHCP sous Windows Server en entreprise"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["attribution"]
source: docs/RAG/collect-261001-rattrapage/dhcp_windows_guide.md
source_anchor: ""
source_lines: [1733, 1907]
sha256: 49127eee42f26526bb398c57c53cfed3c355256ffe17218c9f3fa0553f5a8411
---

# Si le conflit vient d'un bail DHCP attribué à tort : supprimer le bail
Remove-DhcpServerv4Lease -ScopeId 192.0.2.0 -IPAddress 192.0.2.87
```

**Prévention** : inventaire des IP statiques + exclusions systématiques (section 17). Activer la **détection de conflit** : le serveur ping l'adresse avant de l'attribuer (console → propriétés du serveur → onglet **Avancé** → *Nombre de tentatives de détection de conflit*, mettre 1 ou 2). Coût : léger délai à l'attribution.

```powershell
# Activer la détection de conflit (2 essais) — niveau serveur
Set-DhcpServerSetting -ComputerName "srv-dhcp-01.contoso.local" -ConflictDetectionAttempts 2
```

## 78. Cas n°2 — Étendue épuisée (plus d'adresses libres)

**Symptômes** : les nouveaux clients n'obtiennent pas d'IP (`ipconfig` → 169.254.x.x APIPA), les anciens fonctionnent. Événement **1020** dans le journal DHCP-Server.

**Diagnostic** :

```powershell
# Taux d'utilisation par étendue
Get-DhcpServerv4ScopeStatistics |
    Select-Object ScopeId,
        @{N="Utilisation_%";E={[math]::Round($_.PercentageInUse,1)}},
        Free, InUse, Pending |
    Sort-Object "Utilisation_%" -Descending | Format-Table -AutoSize

# Analyser : qui consomme ? (baux par âge, équipements inconnus)
Get-DhcpServerv4Lease -ScopeId 192.0.3.0 |
    Where-Object { $_.AddressState -eq "Active" } |
    Group-Object { $_.HostName -replace '\..*$','' } |
    Sort-Object Count -Descending | Select-Object -First 10 Name, Count
```

**Résolution (par ordre de préférence)** :

1. **Réduire la durée du bail** (ex. : Wi-Fi de 8 jours → 1 jour) : libère les adresses des visiteurs partis.
2. **Nettoyer** les baux expirés (section 28).
3. **Étendre la plage** (section 25) si le masque le permet.
4. **Recréer l'étendue** avec un masque plus large (/24 → /23) — heure creuse.
5. **Scinder** : créer un VLAN supplémentaire (solution d'architecture).

```powershell
# Exemple : Wi-Fi saturé, on passe le bail à 8 heures pour accélérer la rotation
Set-DhcpServerv4Scope -ScopeId 192.0.3.0 -LeaseDuration (New-TimeSpan -Hours 8)
```

**Prévention** : supervision du taux à 80 % (section 24), dimensionnement initial avec 30 % de marge (section 101).

## 79. Cas n°3 — Rogue DHCP : un serveur pirate distribue des adresses

**Symptômes** : des clients obtiennent une **passerelle ou des DNS bizarres** (ex. : DNS = 192.0.2.99 inconnu), pannes Internet intermittentes, parfois par vague (selon qui répond le plus vite au DISCOVER).

**Diagnostic** :

```powershell
# Sur un client affecté : qui est le serveur DHCP ?
ipconfig /all | Select-String "Serveur DHCP"

# Si ce n'est pas 192.0.2.10/11 → rogue confirmé. Trouver sa MAC :
arp -a | Select-String "192.0.2.99"

# Sur le réseau : capturer les OFFER (Wireshark, filtre bootp.option.dhcp == 2)
# L'IP source des OFFER suspects = le rogue
```

**Résolution immédiate** :

1. Identifier le port du switch via la MAC (`show mac address-table` côté switch).
2. **Shutdown** le port / débrancher l'équipement.
3. Sur les clients : `ipconfig /release` + `/renew` pour reprendre un bail légitime.

**Prévention (défense en profondeur, détails section 92-96)** :

- Autorisation AD (section 9) — bloque les serveurs Windows non autorisés.
- **DHCP snooping** sur les switchs — bloque les OFFER venant des ports non trusted.
- 802.1X — empêche les équipements non autorisés de parler au réseau.

> ⚠️ L'autorisation AD ne bloque que les serveurs **Windows**. Une box 4G, un routeur domestique ou un serveur Linux pirate **ne sont pas bloqués** par l'autorisation AD → le DHCP snooping est indispensable.

## 80. Cas n°4 — Un VLAN entier n'obtient plus d'IP (relais DHCP en cause)

**Symptômes** : tout un VLAN (ex. : Wi-Fi) n'a plus de DHCP, les autres VLANs fonctionnent.

**Checklist ordonnée** :

```powershell
# 1. L'étendue existe et est active ?
Get-DhcpServerv4Scope -ScopeId 192.0.3.0 | Select-Object Name, State

# 2. Reste-t-il des adresses ?
Get-DhcpServerv4ScopeStatistics -ScopeId 192.0.3.0 | Select-Object Free, InUse

# 3. Le service DHCP tourne-t-il sur les 2 serveurs ?
Get-Service -ComputerName "srv-dhcp-01.contoso.local" -Name DHCPServer
Get-Service -ComputerName "srv-dhcp-02.contoso.local" -Name DHCPServer
```

Puis côté réseau :

4. `show ip interface brief` / `display ip interface` : l'interface VLAN est-elle up ?
5. L'`ip helper-address` est-il toujours configuré ? (une mise à jour de switch peut l'effacer)
6. Le pare-feu inter-VLAN laisse-t-il passer UDP 67/68 ?
7. Test depuis le routeur : `ping 192.0.2.10` (joignabilité du serveur DHCP).

**Cas vécu** : après un remplacement de switch, la config `ip helper-address` n'a pas été reprise → tout le VLAN voix sans DHCP. **Leçon** : la config réseau des VLANs (helpers inclus) doit être versionnée/sauvegardée comme le reste.

## 81. Cas n°5 — Un seul client n'obtient pas d'IP

**Symptômes** : un poste reste en 169.254.x.x, ses voisins fonctionnent.

**Diagnostic côté client** :

```powershell
# 1. Le service client DHCP tourne ?
Get-Service -Name Dhcp

# 2. Carte réseau : câble, pilote, état ?
Get-NetAdapter | Select-Object Name, Status, LinkSpeed, MacAddress

# 3. Forcer un renouvellement en regardant les erreurs
ipconfig /release
ipconfig /renew
# Noter le message d'erreur exact

# 4. La MAC est-elle bannie par un filtre Deny ?
Get-DhcpServerv4Filter -List Deny | Where-Object { $_.MacAddress -eq "AA-BB-CC-DD-EE-FF" }

# 5. Y a-t-il une réservation pointant vers une autre MAC ? (conflit de réservation)
Get-DhcpServerv4Reservation -ScopeId 192.0.2.0 | Where-Object { $_.IPAddress -eq "192.0.2.87" }
```

**Causes fréquentes** : câble défectueux, port switch en erreur (errdisable), filtre Deny, 802.1X qui échoue, client avec IP statique résiduelle sur une autre plage, pilote NIC corrompu.

## 82. Cas n°6 — Le PXE ne fonctionne pas (pas d'IP au boot réseau)

**Symptômes** : au boot PXE, le client affiche *"No DHCP offers"* ou timeout, alors que Windows obtient une IP normalement.

**Particularité** : le client PXE (firmware) est **plus strict** que Windows sur les timings et les options.

Checklist :

```powershell
# 1. Les options 066/067 sont-elles configurées ? (si WDS séparé du DHCP)
Get-DhcpServerv4OptionValue -ScopeId 192.0.2.0 -OptionId 66
Get-DhcpServerv4OptionValue -ScopeId 192.0.2.0 -OptionId 67

# 2. Les stratégies par architecture existent-elles ? (BIOS vs UEFI)
Get-DhcpServerv4Policy -ScopeId 192.0.2.0 | Select-Object Name, Enabled

# 3. WDS et DHCP sur le même serveur ? -> vérifier la config WDS (pas d'options 66/67 !)
# Console WDS -> Propriétés du serveur -> onglet DHCP :
#   [x] Ne pas écouter sur le port 67
#   [x] Configurer les options DHCP sur 060 (PXEClient)
```

Autres causes :

- **IP helper** : certains relais ne forwardent pas correctement les requêtes PXE (option 60). Sur Cisco : `ip helper-address` suffit normalement ; sinon ajouter `ip forward-protocol udp 4011`.
- **Délai** : le firmware PXE timeout vite ; si le serveur DHCP est lent (détection de conflit activée, section 77), le PXE échoue alors que Windows réussit. → Désactiver la détection de conflit ou augmenter le timeout PXE.
- **UEFI Secure Boot** : un fichier de boot non signé est refusé (pas un problème DHCP, mais symptôme similaire).

## 83. Cas n°7 — Le basculement ne réplique plus (State ≠ Normal)

**Symptômes** : les baux créés sur le primaire n'apparaissent pas sur le secondaire. `Get-DhcpServerv4Failover` → `State = CommunicationInterrupted` ou `PartnerDown`.

**Diagnostic** :

```powershell
# État détaillé
Get-DhcpServerv4Failover -ComputerName "srv-dhcp-01.contoso.local" |
    Format-List Name, Mode, State, ServerRole, PartnerServer

# Causes à vérifier dans l'ordre :
# 1. Connectivité TCP 647 entre les deux serveurs
Test-NetConnection -ComputerName "srv-dhcp-02.contoso.local" -Port 647

