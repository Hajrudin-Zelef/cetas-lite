---
id: collect-261001-rattrapage/rattrapage/maintenance-windows-guide-13
title: "Maintenance et exploitation Windows en entreprise"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/maintenance_windows_guide.md
source_anchor: ""
source_lines: [2291, 2521]
sha256: b1890982829772843be19dfd45609637700a1dff80a8da558e08f9c46ea5c138
---

# Maintenance et exploitation Windows en entreprise

```powershell
# Le ping est-il OK + quel est le chemin ?
Test-NetConnection -ComputerName srv-fic-01

# Tester un PORT TCP précis (le vrai test : "l'application répond-elle ?")
Test-NetConnection -ComputerName srv-fic-01 -Port 445      # SMB
Test-NetConnection -ComputerName srv-ad-01  -Port 389      # LDAP
Test-NetConnection -ComputerName srv-web-01 -Port 443      # HTTPS
Test-NetConnection -ComputerName srv-rds-01 -Port 3389     # RDP

# Détail complet : IP résolue, latence, route
Test-NetConnection -ComputerName intranet.contoso.local -Port 443 -InformationLevel Detailed |
  Select-Object ComputerName, RemoteAddress, TcpTestSucceeded, PingSucceeded,
    @{n='LatenceMs';e={$_.PingReplyDetails.RoundtripTime}}, TraceRoute
```

> `TcpTestSucceeded : False` + ping OK = le service n'écoute pas, le pare-feu
> bloque, ou vous frappez à la mauvaise IP. C'est **le** test qui tranche entre
> « problème réseau » et « problème applicatif ».

---

## 80. netstat et Get-NetTCPConnection

Qui écoute quoi, qui est connecté à qui :

```cmd
:: Connexions actives + ports d'écoute, avec le PID (-o) et résolution (-b = exécutable)
netstat -ano
netstat -anob
```

```powershell
# Équivalent PowerShell, filtrable
Get-NetTCPConnection | Where-Object State -eq 'Listen' |
  Select-Object LocalAddress, LocalPort,
    @{n='PID';e={$_.OwningProcess}},
    @{n='Processus';e={(Get-Process -Id $_.OwningProcess -ErrorAction SilentlyContinue).ProcessName}} |
  Sort-Object LocalPort | Format-Table -AutoSize

# Qui occupe le port 8080 ?
Get-NetTCPConnection -LocalPort 8080 |
  Select-Object LocalAddress, LocalPort, State, OwningProcess,
    @{n='Processus';e={(Get-Process -Id $_.OwningProcess).ProcessName}}
```

Cas d'usage : « l'application ne démarre plus : port déjà utilisé » (trouvez le
squatteur), « connexions TIME_WAIT par milliers » (épuisement de ports éphémères
sur un serveur très sollicité → ajuster `MaxUserPort`/`TcpTimedWaitDelay` avec
parcimonie et mesure).

---

## 81. DNS côté client : nslookup et Resolve-DnsName

```cmd
:: Requête simple (utilise le DNS configuré)
nslookup intranet.contoso.local

:: Interroger un serveur précis (pour comparer : mon DNS vs un autre)
nslookup intranet.contoso.local 10.10.10.11
nslookup intranet.contoso.local 10.10.10.12

:: Types d'enregistrements (MX, SRV pour AD, TXT...)
nslookup -type=SRV _ldap._tcp.contoso.local
```

```powershell
# Version PowerShell (plus propre, scriptable)
Resolve-DnsName intranet.contoso.local
Resolve-DnsName intranet.contoso.local -Server 10.10.10.11
Resolve-DnsName -Name _ldap._tcp.contoso.local -Type SRV
```

Diagnostic AD : si les SRV `_ldap._tcp.<domaine>` ne se résolvent pas, **rien**
dans le domaine ne fonctionnera (jonction, GPO, logon). C'est le test n°1 d'un
poste « le domaine ne répond plus ».

---

## 82. DNS : vider le cache, vérifier les suffixes

```cmd
ipconfig /flushdns
```

```powershell
# Suffixes DNS configurés (indispensables pour résoudre les noms courts)
Get-DnsClientGlobalSetting | Select-Object -ExpandProperty SuffixSearchList

# Serveurs DNS par interface
Get-DnsClientServerAddress -AddressFamily IPv4 |
  Select-Object InterfaceAlias, ServerAddresses

# Corriger : imposer les DNS du domaine sur une interface (ex. Ethernet0)
Set-DnsClientServerAddress -InterfaceAlias 'Ethernet0' -ServerAddresses '10.10.10.11','10.10.10.12'
```

**Classique** : un poste avec un DNS externe (box, 8.8.8.8) en premier ne résout
plus les noms internes et perd le domaine par intermittence. Sur un domaine AD,
les clients doivent **toujours** pointer vers les DC/DNS internes en premier.

---

## 83. DHCP côté client : renouvellement et diagnostic

```cmd
ipconfig /release
ipconfig /renew
ipconfig /all   :: vérifier : serveur DHCP, bail, durée
```

```powershell
# Journal des événements DHCP client
Get-WinEvent -LogName 'Microsoft-Windows-DHCP Client Events/Admin' -MaxEvents 20 |
  Select-Object TimeCreated, Id, Message | Format-Table -AutoSize -Wrap
```

Si `ipconfig /renew` échoue (« impossible de contacter le serveur DHCP ») :

1. Le client voit-il le réseau ? (ping passerelle, §77)
2. Le serveur DHCP est-il joignable ? (`Test-NetConnection srv-dhcp -Port 67`
   ne marche pas en TCP — utilisez le journal serveur §40 et un test depuis un
   poste sain du même VLAN)
3. **Relais DHCP (IP helper)** configuré sur le routeur du VLAN ? (cause n°1
   après un changement de VLAN)
4. Étendue saturée ? (ID 1020 côté serveur, §40)
5. Filtrage MAC / classeur DHCP qui exclut le poste ?

---

## 84. Wi-Fi : diagnostic de base

```cmd
:: Réseaux visibles et signal
netsh wlan show networks mode=bssid

:: État de la connexion actuelle : SSID, signal, canal, débit
netsh wlan show interfaces

:: Rapport HTML complet (utile pour le support)
netsh wlan show wlanreport
:: Génère C:\ProgramData\Microsoft\Windows\WlanReport\wlan-report-latest.html
```

```powershell
# Profils Wi-Fi mémorisés
netsh wlan show profiles
# Oublier un profil corrompu (clé changée côté AP par ex.)
netsh wlan delete profile name="WiFi-Atelier"
```

Réflexes : signal < -70 dBm = couverture insuffisante ; débit négocié très
inférieur au standard = interférences ou pilote ; **pilote Wi-Fi à jour** avant
toute chose (c'est la cause n°1 des déconnexions inexpliquées sur Windows 11).

---

## 85. Cas pratique 1 : pas d'accès réseau du tout

**Symptôme** : le poste n'atteint rien, pas même la passerelle. Pas d'Internet,
pas d'intranet.

**Diagnostic commenté :**

```cmd
:: 1. La carte a-t-elle une IP ?
ipconfig
:: -> Si "Média déconnecté" : câble/switch. Si APIPA : voir cas n°2.
```

```powershell
# 2. La carte est-elle UP côté Windows ?
Get-NetAdapter | Select-Object Name, Status, LinkSpeed
# Status = Disconnected -> physique. Disabled -> réactiver :
Enable-NetAdapter -Name 'Ethernet0' -Confirm:$false
```

```powershell
# 3. Passerelle joignable ?
Test-Connection -ComputerName 10.10.20.1 -Count 2
# KO -> tester le câble sur une autre prise, vérifier le voyant du switch,
#      essayer un autre port du switch (port HS ou VLAN mal affecté).
```

**Résolution typique** : câble débranché/cassé (60 %), port du switch en erreur
(`err-disabled` après une boucle), carte désactivée dans Windows, ou pilote
corrompu (réinstaller le pilote). **Toujours** tester avec un câble et un port
connus bons avant d'incriminer la carte.

---

## 86. Cas pratique 2 : adresse APIPA 169.254.x.x

**Symptôme** : `ipconfig` montre `169.254.x.x`, masque `255.255.0.0`, pas de
passerelle. Le poste s'est auto-attribué une adresse = **aucun serveur DHCP
n'a répondu**.

**Diagnostic commenté :**

```cmd
ipconfig /all
:: Vérifier : "Serveur DHCP . . . : 255.255.255.255" ou absent -> aucun contact
```

1. **Un seul poste ?** → local : `ipconfig /release` + `/renew` ; si échec,
   câble/switch (cas n°1) ou pile réseau corrompue :
   ```cmd
   netsh winsock reset
   netsh int ip reset
   :: puis redémarrage
   ```
2. **Tout un VLAN ?** → serveur/relais : le service DHCP tourne-t-il ?
   l'étendue est-elle saturée (§40) ? le relais IP helper est-il toujours
   configuré sur le routeur (vérifiez après chaque changement de VLAN) ?
3. **Après un changement** (nouveau switch, nouveau VLAN) → dans 90 % des cas,
   le relais DHCP n'a pas été configuré sur le nouvel équipement.

**Piège** : un serveur DHCP « sauvage » (box, point d'accès mal configuré)
distribue de mauvaises IP : `ipconfig /all` montre alors un serveur DHCP
inattendu → traquez-le via son adresse MAC (`arp -a`).

---

## 87. Cas pratique 3 : DNS ne résout plus

**Symptôme** : le ping d'IP fonctionne (`ping 10.10.30.5` OK) mais les noms
échouent (`ping intranet` → « hôte introuvable »).

**Diagnostic commenté :**

