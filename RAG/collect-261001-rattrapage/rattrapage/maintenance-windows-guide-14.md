---
id: collect-261001-rattrapage/rattrapage/maintenance-windows-guide-14
title: "Maintenance et exploitation Windows en entreprise"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["arr", "full-duplex"]
source: docs/RAG/collect-261001-rattrapage/maintenance_windows_guide.md
source_anchor: ""
source_lines: [2522, 2720]
sha256: b7051eb984fd655ec1a840f6059865f7fce031becbd5f4a856ea206c15925f5e
---

# Maintenance et exploitation Windows en entreprise

```powershell
# 1. Quels DNS le poste utilise-t-il ?
Get-DnsClientServerAddress -AddressFamily IPv4 | Select-Object InterfaceAlias, ServerAddresses
# -> Si DNS externe/box en premier sur un poste du domaine : corriger (§82).

# 2. Le DNS configuré répond-il ?
Resolve-DnsName intranet.contoso.local -Server 10.10.10.11
# KO -> tester le second DNS ; si les deux KO depuis ce poste mais OK depuis
# un autre : pare-feu local ou réseau (port 53 UDP/TCP filtré).

# 3. Simple problème de cache ?
ipconfig /flushdns
```

**Résolutions typiques** : remettre les DNS internes en premier (§82) ; corriger
le suffixe DNS ; côté serveur : zone DNS corrompue ou service DNS arrêté
(journal DNS Server §39). **Test décisif** : `nslookup nom 10.10.10.11` vs
`nslookup nom 10.10.10.12` — si un seul répond, le problème est ce serveur-là,
pas « le DNS » en général.

---

## 88. Cas pratique 4 : un seul site inaccessible

**Symptôme** : tout marche sauf `https://appli-metier.contoso.local` (ou un site
Internet précis).

**Diagnostic commenté :**

```powershell
# 1. Résolution OK ?
Resolve-DnsName appli-metier.contoso.local
# 2. Port ouvert ?
Test-NetConnection -ComputerName appli-metier.contoso.local -Port 443
# 3. Où ça casse ?
tracert -d appli-metier.contoso.local
```

- **Résolution KO** → DNS (cas n°3), entrée manquante.
- **Port fermé** → service arrêté sur le serveur, pare-feu (Windows ou réseau),
  mauvaise IP (DNS menteur / fichier `hosts` — vérifiez
  `C:\Windows\System32\drivers\etc\hosts` !).
- **Tracert bloqué à un saut précis** → équipement intermédiaire (souvent après
  un changement de routage ou une règle de pare-feu).
- **Site Internet unique KO** → proxy (§99), filtrage, ou panne du site
  (vérifiez depuis un smartphone en 4G/5G pour trancher).

**Le fichier `hosts`** est un classique sous-estimé : une ligne oubliée par un
ancien admin redirige le nom vers une IP morte. Toujours le contrôler.

---

## 89. Cas pratique 5 : lenteurs réseau intermittentes

**Symptôme** : « c'est lent » par moments, sans panne franche.

**Diagnostic commenté :**

```powershell
# 1. Objectiver : ping long avec horodatage vers la ressource
Test-Connection -ComputerName srv-fic-01 -Count 600 |
  Where-Object { $_.ResponseTime -gt 100 -or $_.StatusCode -ne 0 } |
  Select-Object @{n='Heure';e={Get-Date -Format 'HH:mm:ss'}}, ResponseTime, StatusCode
# -> Des pics réguliers (ex. toutes les heures) = un batch, une sauvegarde, un antivirus.
```

2. **Heure des pics** : corrélez avec sauvegardes, antivirus (scan planifié),
   WSUS/WUfB (téléchargements), sauvegardes des VM sur l'hôte.
3. **Wi-Fi ?** → signal, interférences, roaming (§84).
4. **Duplex** : `Get-NetAdapter | Select Name, LinkSpeed, FullDuplex` — un lien
   en half-duplex ou 100 Mb/s explique tout (§58).
5. **Saturation** : compteurs `\Network Interface(*)\Bytes Total/sec` (§58) aux
   heures de pics.

**Résolution** : déplacer les batchs gourmands hors heures ouvrées, activer le
QoS si pertinent, passer le lien en 1 Gb/s full-duplex forcé si la négociation
échoue, ou ajouter de la bande passante — **après** avoir prouvé la saturation
par les compteurs, pas avant.

---

## 90. Cas pratique 6 : partage SMB inaccessible

**Symptôme** : `\\srv-fic-01\partage` → « chemin introuvable » ou « accès refusé ».

**Diagnostic commenté :**

```powershell
# 1. Le serveur répond-il sur SMB (445) ?
Test-NetConnection -ComputerName srv-fic-01 -Port 445
# KO -> serveur éteint, pare-feu, ou SMB désactivé.

# 2. Quelles sont les erreurs SMB côté client ?
Get-WinEvent -LogName 'Microsoft-Windows-SMBClient/Connectivity' -MaxEvents 10 |
  Select-Object TimeCreated, Id, Message | Format-Table -AutoSize -Wrap

# 3. Accès refusé -> droits NTFS + droits de partage (les DEUX doivent autoriser)
# Sur le serveur : vérifier le partage et les ACL
Get-SmbShare -Name 'partage' | Select-Object Name, Path, Description
Get-SmbShareAccess -Name 'partage'
```

**Distinguer** :

- « Chemin d'accès introuvable » (0x80070035) = réseau/nom/partage inexistant.
- « Accès refusé » (0x80070005) = authentification ou droits (vérifiez avec quel
  compte : `whoami`, session verrouillée avec d'anciens identifiants ?).
- **Après une mise à jour** : SMBv1 désactivé par défaut depuis longtemps — un
  vieux NAS/copieur qui ne parle que SMBv1 ne se connecte plus (ne **pas**
  réactiver SMBv1 par facilité : mettez à jour l'équipement ou isolez-le).

---

## 91. Cas pratique 7 : imprimante réseau hors ligne

**Symptôme** : l'imprimante apparaît « hors ligne » alors qu'elle est allumée.

**Diagnostic commenté :**

```powershell
# 1. L'imprimante répond-elle ?
Test-Connection -ComputerName 10.10.40.25 -Count 2
# 2. Le port d'impression est-il joignable (9100 = RAW, 515 = LPR) ?
Test-NetConnection -ComputerName 10.10.40.25 -Port 9100
# 3. État côté Windows
Get-Printer | Where-Object Name -like '*Atelier*' | Select-Object Name, PrinterStatus, PortName
```

**Causes fréquentes** : IP de l'imprimante changée (DHCP sans réservation →
**toujours** des réservations DHCP ou IP fixes pour les imprimantes), SNMP mal
configuré (Windows la déclare hors ligne si la communauté SNMP ne répond pas :
désactivez « État SNMP » sur le port ou configurez-le), spouleur planté côté
serveur d'impression (`Restart-Service Spooler`).

**Bon réflexe** : imprimez la page de configuration réseau depuis l'imprimante
elle-même — elle donne sa vraie IP, souvent différente de celle que Windows croit.

---

## 92. Cas pratique 8 : VPN qui ne se connecte plus

**Symptôme** : le client VPN échoue (erreur 800, 809, 13801…).

**Diagnostic commenté :**

1. **Internet OK sans VPN ?** (sinon : cas n°1-3 d'abord).
2. **Code d'erreur exact** : 809 = ports IPsec bloqués (UDP 500/4500) — classique
   derrière une box/hôtel ; 800 = serveur injoignable ; 13801 = certificat.
3. **Depuis un autre réseau** (partage de connexion smartphone) : si ça marche →
   le réseau local filtre (pare-feu de la box, Wi-Fi public).
4. Côté client : journal `Applications et services > Microsoft > Windows >
   NetworkProfile` + événements RAS :

```powershell
Get-WinEvent -LogName Application -MaxEvents 50 |
  Where-Object { $_.ProviderName -like '*Ras*' } |
  Select-Object TimeCreated, Id, Message | Format-Table -AutoSize -Wrap
```

**Résolutions** : ouvrir/relayer UDP 500 et 4500 + protocole ESP (50) côté
pare-feu ; réparer le profil VPN (supprimer/recréer) ; vérifier le certificat
client (expiré ? §120) ; mettre à jour le client VPN.

---

## 93. Cas pratique 9 : port fermé / application injoignable

**Symptôme** : `Test-NetConnection -Port XXXX` → `TcpTestSucceeded : False`.

**Diagnostic commenté :**

```powershell
# 1. Le service écoute-t-il VRAIMENT sur ce serveur ?
Invoke-Command -ComputerName srv-appli-01 {
    Get-NetTCPConnection -State Listen | Where-Object LocalPort -eq 8443
}
# Rien -> le service n'est pas démarré ou n'écoute pas sur ce port/interface.

# 2. Le pare-feu Windows laisse-t-il passer ?
Get-NetFirewallRule -DisplayName '*appli*' | Select-Object DisplayName, Enabled, Direction, Action
Get-NetFirewallPortFilter -Protocol TCP | Where-Object LocalPort -eq 8443 |
  Get-NetFirewallRule | Select-Object DisplayName, Enabled, Action

# 3. Test depuis le serveur lui-même (loopback) :
#    OK en local + KO à distance = pare-feu (Windows ou réseau).
```

**Ordre de vérification** : service démarré → écoute sur la bonne interface
(`0.0.0.0` vs `127.0.0.1` — une appli liée au loopback n'est joignable que
localement) → pare-feu Windows → pare-feu réseau inter-VLAN. **90 % des « ports
fermés » sont un service arrêté ou un pare-feu Windows**, pas le réseau.

---

## 94. Cas pratique 10 : conflit d'adresse IP

**Symptôme** : message « Conflit d'adresse IP », connectivité intermittente
(ça marche, ça coupe).

