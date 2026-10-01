---
id: collect-261001-rattrapage/rattrapage/maintenance-windows-guide-15
title: "Maintenance et exploitation Windows en entreprise"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["exploit", "arr"]
source: docs/RAG/collect-261001-rattrapage/maintenance_windows_guide.md
source_anchor: ""
source_lines: [2721, 2933]
sha256: 34ff517db2272c5929007608bdee26ef3999ec8987b88406628e33937d2bba89
---

# Maintenance et exploitation Windows en entreprise

**Diagnostic commenté :**

```cmd
:: Voir qui squatte l'IP : l'adresse MAC du conflit apparaît dans les événements
ipconfig /all
arp -a | findstr 10.10.20.45
```

```powershell
# Journal système : l'événement 4199 (Tcpip) signale le conflit avec la MAC adverse
Get-WinEvent -FilterHashtable @{ LogName='System'; Id=4199 } -MaxEvents 5 |
  Select-Object TimeCreated, Message | Format-Table -AutoSize -Wrap
```

**Résolution** : identifier la machine via la MAC (table du switch : `show mac
address-table`), puis : IP en dur oubliée sur un équipement (imprimante, caméra)
→ passer en DHCP + réservation ; étendue DHCP qui chevauche des IP fixes →
**exclure** la plage des IP fixes de l'étendue ; deux DHCP sauvages → supprimer
l'intrus (§86).

---

## 95. Cas pratique 11 : profil réseau Public au lieu de Domaine

**Symptôme** : le pare-feu se durcit tout seul, les partages et la découverte
réseau ne marchent plus, parfois les GPO ne s'appliquent pas.

**Diagnostic commenté :**

```powershell
# Quel profil le réseau a-t-il ?
Get-NetConnectionProfile | Select-Object Name, InterfaceAlias, NetworkCategory
# NetworkCategory = Public au lieu de DomainAuthenticated -> problème
```

**Causes** : le poste ne parvient pas à contacter un DC au moment de la
détection (DC injoignable, DNS faux §87, service NLA démarré avant le réseau).
**Résolution** : corriger la cause racine (DNS/DC), puis forcer la
re-détection :

```powershell
Restart-Service nlasvc -Force
# Vérifier 30 s après :
Get-NetConnectionProfile | Select-Object Name, NetworkCategory
```

> Ne « corrigez » pas en passant le profil en Privé à la main via le registre :
> vous masquez le symptôme (le poste ne voit toujours pas le domaine) et vous
> affaiblissez le pare-feu.

---

## 96. Cas pratique 12 : débit anormalement bas

**Symptôme** : copie de fichiers à 10 Mo/s sur un lien gigabit.

**Diagnostic commenté :**

```powershell
# 1. Négociation du lien
Get-NetAdapter | Select-Object Name, LinkSpeed, FullDuplex, MediaConnectionState
# LinkSpeed = 100 Mbps ou FullDuplex = False -> cause trouvée.

# 2. Erreurs d'interface (compteurs)
Get-Counter '\Network Interface(*)\Packets Received Errors',
            '\Network Interface(*)\Packets Outbound Errors' |
  Select-Object -ExpandProperty CounterSamples | Format-Table -AutoSize

# 3. Mesure réelle : copie d'un fichier témoin de 1 Go chronométrée
Measure-Command { Copy-Item \\srv-fic-01\partage\fichier-1go.bin C:\Temp\ }
```

**Causes** : câble abîmé (paire HS → 100 Mb/s), port switch forcé en 100/half,
pilote de carte réseau obsolète, **SMB signing** ou antivirus qui inspecte les
flux, Wi-Fi saturé. **Méthode** : mesurez **avant/après** chaque changement ;
un « ça semble mieux » n'est pas un diagnostic.

---

## 97. Cas pratique 13 : RDP impossible vers un serveur

**Symptôme** : le client RDP échoue (« l'ordinateur distant n'a pas répondu »,
erreur d'authentification, écran noir).

**Diagnostic commenté :**

```powershell
# 1. Port 3389 ouvert ?
Test-NetConnection -ComputerName srv-rds-01 -Port 3389
# 2. Le service est-il en écoute sur le serveur (via un autre accès : console, PSRemoting) ?
Invoke-Command -ComputerName srv-rds-01 { Get-Service TermService | Select-Object Status }
# 3. RDP autorisé dans le pare-feu ?
Invoke-Command -ComputerName srv-rds-01 {
    Get-NetFirewallRule -DisplayGroup 'Bureau à distance' | Select-Object DisplayName, Enabled
}
```

**Arbre de décision** : port fermé → service arrêté ou pare-feu ; port ouvert
mais échec d'authentification → compte verrouillé/expiré, heure désynchronisée
(Kerberos, §104), NLA ; connexion OK puis écran noir → session bloquée ou
problème graphique (tentez `/admin`) ; **après un patch** → vérifiez les
correctifs RDP récents (rollback §20 si régression avérée).

---

## 98. Cas pratique 14 : double pile IPv4/IPv6 capricieuse

**Symptôme** : résolution ou connexion aléatoire ; certaines applis préfèrent
IPv6 et échouent alors qu'IPv4 marche.

**Diagnostic commenté :**

```powershell
# Que résout le nom ? (souvent une adresse IPv6 link-local ou ULA en premier)
Resolve-DnsName srv-fic-01 | Select-Object Name, Type, IPAddress

# Forcer le test en IPv4 pour comparer
Test-NetConnection -ComputerName srv-fic-01 -Port 445   # utilise la résolution normale
ping -4 srv-fic-01
ping -6 srv-fic-01
```

**Résolutions** : si IPv6 n'est **pas** exploité dans votre infra (pas de DHCPv6,
pas de DNSv6), le plus propre est de **préférer IPv4** plutôt que de désactiver
IPv6 sauvagement (la désactivation totale casse des composants Windows : ne
décochez pas IPv6 sur la carte !). Ajustez les **stratégies de préfixe**
(`netsh interface ipv6 show prefixpolicies`) ou corrigez le DNS pour ne plus
publier d'enregistrements AAAA parasites.

---

## 99. Cas pratique 15 : proxy / PAC qui casse tout

**Symptôme** : Internet OK pour certains sites/applis, KO pour d'autres ;
erreurs de certificat ; applis qui « ne voient pas Internet » alors que le
navigateur oui.

**Diagnostic commenté :**

```powershell
# 1. Quel proxy est configuré ?
Get-ItemProperty 'HKCU:\Software\Microsoft\Windows\CurrentVersion\Internet Settings' |
  Select-Object ProxyEnable, ProxyServer, AutoConfigURL
# ProxyEnable=1 + ProxyServer vide/bizarre -> résidu d'un ancien proxy.

# 2. Le PAC est-il accessible et sain ?
Invoke-WebRequest -Uri 'http://proxy.contoso.local:8080/proxy.pac' -UseBasicParsing |
  Select-Object StatusCode

# 3. Test sans proxy (session de test)
netsh winhttp show proxy   # proxy au niveau machine (services !)
```

**Points clés** : le proxy **machine** (`netsh winhttp`) ≠ le proxy
**utilisateur** (navigateur) — les services et beaucoup d'applis utilisent le
premier ; un fichier PAC corrompu/injoignable = pannes intermittentes par site ;
après la suppression d'un proxy, pensez à nettoyer **les deux** niveaux + les
variables d'environnement `http_proxy` résiduelles.

---

## 100. AD/GPO côté client : gpupdate et gpresult

```cmd
:: Appliquer les stratégies (ordinateur + utilisateur)
gpupdate /force

:: Cibler : uniquement l'ordinateur (nécessite souvent un redémarrage)
gpupdate /target:computer /force
```

```powershell
# Rapport des stratégies appliquées (texte)
gpresult /r
# Rapport HTML complet (à ouvrir dans le navigateur) - idéal pour le support
gpresult /h C:\Temp\gpresult.html
# Stratégies pour un autre utilisateur/ordinateur (depuis un poste admin)
gpresult /s srv-poste-042 /user contoso\jdupont /h C:\Temp\gpo-042.html
```

**Lecture du rapport** : vérifiez « Dernière application de la stratégie de
groupe » (une date vieille de plusieurs jours = problème), la liste des GPO
appliquées **et refusées** (avec le motif : filtre de sécurité, WMI, GPO vide).
Une GPO « refusée » pour cause de filtrage de sécurité est le n°1 des « ma GPO ne
s'applique pas ».

---

## 101. RSOP et journaux du service de stratégie de groupe

- `rsop.msc` : jeu de stratégies résultant (visuel, équivalent de `gpresult`).
- Journal dédié : `Applications et services > Microsoft > Windows >
  GroupPolicy > Opérationnel` — chaque application de GPO y est tracée avec les
  erreurs par extension (scripts, registre, mappages…).

```powershell
# Erreurs GroupPolicy sur 7 jours
Get-WinEvent -LogName 'Microsoft-Windows-GroupPolicy/Operational' -MaxEvents 200 |
  Where-Object LevelDisplayName -eq 'Error' |
  Select-Object TimeCreated, Id, Message | Format-Table -AutoSize -Wrap
```

**IDs utiles** : 1030/1058 (accès SYSVOL/NETLOGON impossible → réseau/DNS),
1129 (échec de traitement au démarrage), 7016/7017 (extensions en timeout).
Si les erreurs 1030/1058 sont massives sur le parc : suspectez **DFSR/SYSVOL**
côté DC (§39), pas les postes.

---

## 102. nltest : diagnostiquer la relation au domaine

`nltest` interroge le canal sécurisé et les DC :

