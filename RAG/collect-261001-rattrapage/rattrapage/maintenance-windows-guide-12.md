---
id: collect-261001-rattrapage/rattrapage/maintenance-windows-guide-12
title: "Maintenance et exploitation Windows en entreprise"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["arr", "ethernet", "memory"]
source: docs/RAG/collect-261001-rattrapage/maintenance_windows_guide.md
source_anchor: ""
source_lines: [2079, 2290]
sha256: fb8d4b563c95c0469fca523c0a5f9865cdd81bd24e5e03b626c43aa55cdf739b
---

# Maintenance et exploitation Windows en entreprise

```text
!analyze -v        :: analyse automatique détaillée (le 90 % du travail)
k                 :: pile d'appels (stack)
lm                :: modules/pilotes chargés
!thread           :: thread fautif
```

> Les symboles se téléchargent à la première analyse (connexion Internet requise,
> quelques minutes). Sans symboles corrects, `!analyze` reste approximatif.

---

## 72. Analyser un minidump : !analyze -v

`!analyze -v` produit un rapport dont les lignes clés sont :

```text
IRQL_NOT_LESS_OR_EQUAL (a)
...
Probably caused by : driver-fic.sys ( driver_fic+1234 )
...
STACK_TEXT:
...
IMAGE_NAME:  driver-fic.sys
FAILURE_BUCKET_ID:  AV_driver-fic!Unknown_Function
```

Lecture :

| Champ | Sens |
|---|---|
| Code d'arrêt + nom | La famille d'erreur (§73) |
| `Probably caused by` | Le suspect n°1 (indication, pas preuve absolue) |
| `IMAGE_NAME` | Le module incriminé |
| `STACK_TEXT` | La pile d'appels : ce que faisait le noyau |
| `FAILURE_BUCKET_ID` | Signature : cherchez-la telle quelle sur le web |

**Conduite à tenir :**

1. Si le fautif est un **pilote tiers** (`*.sys` non-Microsoft) : mettez-le à jour,
   revenez à la version précédente, ou désactivez le matériel/logiciel associé.
2. Si c'est `ntoskrnl.exe` / `hal.dll` : c'est rarement le vrai coupable — cherchez
   plus bas dans la pile, suspectez **RAM** (test mémoire) ou **pilote** qui a
   corrompu le noyau.
3. BSOD variés et changeants à chaque crash = **matériel** (RAM, alimentation,
   surchauffe) dans 80 % des cas, pas un pilote.

Test mémoire intégré : `mdsched.exe` (redémarrage requis) ; en doute, MemTest86+
depuis une clé USB.

---

## 73. Codes d'arrêt courants : tableau de référence

| Code | Nom | Cause typique | Piste |
|---|---|---|---|
| `0x0A` | IRQL_NOT_LESS_OR_EQUAL | Pilote accédant à une adresse invalide | Pilote (MAJ/restauration) |
| `0x1A` | MEMORY_MANAGEMENT | RAM défectueuse ou pilote | Test mémoire, pilotes |
| `0x3B` | SYSTEM_SERVICE_EXCEPTION | Pilote/service système | Dump + MAJ pilote |
| `0x7E` | SYSTEM_THREAD_EXCEPTION_NOT_HANDLED | Pilote (souvent graphique/stockage) | Mode sans échec, restauration pilote |
| `0x50` | PAGE_FAULT_IN_NONPAGED_AREA | RAM ou pilote, antivirus | RAM, exclusion AV |
| `0x7A` | KERNEL_DATA_INPAGE_ERROR | Disque : page noyau illisible | SMART, chkdsk, câble |
| `0x77` | KERNEL_STACK_INPAGE_ERROR | Idem, pile du noyau | Disque |
| `0xF4` | CRITICAL_OBJECT_TERMINATION | Processus critique tué (souvent disque) | Disque/SMART |
| `0xEF` | CRITICAL_PROCESS_DIED | Processus système mort | Disque, corruption système (DISM/SFC) |
| `0x133` | DPC_WATCHDOG_VIOLATION | Pilote bloquant (SSD/NVMe, Wi-Fi) | Pilote stockage/réseau |
| `0x124` | WHEA_UNCORRECTABLE_ERROR | **Matériel** (CPU, RAM, PCIe) | Diag constructeur, températures |
| `0x109` | CRITICAL_STRUCTURE_CORRUPTION | Pilote vérolé / rootkit | Driver Verifier, analyse AV |
| `0xC000021A` | WINLOGON_FATAL_ERROR | Sous-système utilisateur mort | Corruption système, restauration |
| `0x21` | QUOTA_UNDERFLOW | Pilote (retour de quota incorrect) | Pilote |

---

## 74. Driver Verifier : traquer un pilote fautif

**Driver Verifier** (`verifier.exe`) soumet les pilotes à des contrôles stricts
pour faire planter **immédiatement et avec un fautif identifié** un pilote qui
corrompt silencieusement le système.

```cmd
:: Lancer l'assistant
verifier
```

Procédure :

1. « Créer des paramètres personnalisés » > cochez : Special Pool, Pool Tracking,
   Force IRQL Checking, I/O Verification, Deadlock Detection.
2. « Sélectionner les pilotes dans une liste » > cochez les pilotes **tiers
   suspects** (jamais tous les pilotes Microsoft en production).
3. Redémarrez : le système devient **plus lent** (normal) et plantera vite si un
   pilote surveillé faute → le dump désignera clairement le coupable.
4. Après diagnostic : `verifier /reset` puis redémarrage pour **désactiver**.

> ⚠️ Driver Verifier en production = risque de BSOD en boucle si le pilote est
> vraiment fautif. Prévoyez le mode sans échec (`verifier /reset` depuis WinRE si
> besoin) et une fenêtre de maintenance. Ne le laissez jamais activé « au cas où ».

---

## 75. Dépannage réseau : méthode générale

Ne cliquez pas au hasard : suivez la pyramide, du bas vers le haut.

1. **Physique** : câble branché ? voyant du switch ? Wi-Fi connecté au bon SSID ?
2. **Local** : IP valide ? (`ipconfig`, §76) Passerelle joignable ? (`ping`, §77)
3. **Résolution** : le DNS répond-il ? (`nslookup`, §81)
4. **Routage** : où ça casse ? (`tracert`/`pathping`, §77-78)
5. **Transport** : le port est-il ouvert ? (`Test-NetConnection`, §79)
6. **Applicatif** : le service répond-il ? (logs, test applicatif)

**Règle d'or** : isolez le périmètre — « est-ce que ça marche depuis une autre
machine / un autre réseau ? » divise le problème par deux à chaque fois.
Et **documentez** : heure, symptômes exacts, messages d'erreur recopiés, commandes
lancées et résultats. Un ticket « ça marche pas » est inexploitable.

---

## 76. ipconfig : lecture et cas d'usage

```cmd
:: Configuration complète : IP, masque, passerelle, DNS, DHCP, bail
ipconfig /all

:: Renouveler le bail DHCP (après correction côté serveur)
ipconfig /release
ipconfig /renew

:: Vider / afficher le cache DNS client
ipconfig /flushdns
ipconfig /displaydns
```

Lecture d'un `ipconfig /all` sain (exemple) :

```text
Carte Ethernet Ethernet0 :
   Adresse IPv4 . . . . . . . . : 10.10.20.45
   Masque de sous-réseau . . . . : 255.255.255.0
   Passerelle par défaut . . . . : 10.10.20.1
   Serveur DHCP . . . . . . . . . : 10.10.20.10
   Serveurs DNS . . . . . . . . . : 10.10.10.11
                                     10.10.10.12
   Bail obtenu . . . . . . . . . : dimanche 27 septembre 2026 08:12:03
```

Anomalies à repérer d'un coup d'œil : IP en `169.254.x.x` (APIPA → pas de DHCP,
§86), passerelle absente, DNS en `8.8.8.8` sur un poste du domaine (doit pointer
vers les DC !), suffixe DNS manquant.

---

## 77. Test-Connection, ping, pathping

```powershell
# ping moderne : 4 échos, avec résolution et temps
Test-Connection -ComputerName 10.10.20.1 -Count 4

# ping continu avec horodatage (surveillance d'une coupure)
Test-Connection -ComputerName srv-fic-01 -Count 9999 |
  Select-Object @{n='Heure';e={Get-Date -Format 'HH:mm:ss'}},
    Address, ResponseTime, StatusCode

# Équivalent CMD classique
ping -t 10.10.20.1
pathping 10.10.30.5   # ping + traceroute avec stats par saut (lent : ~5 min)
```

Interprétation :

- **Délai dépassé** vers la passerelle = problème local (câble, switch, carte).
- **OK passerelle, KO au-delà** = routage / lien inter-sites / pare-feu.
- **Pertes intermittentes** = notez le % : < 1 % tolérable, > 5 % = à traiter
  (saturation, duplex, Wi-Fi).
- ⚠️ Un ping qui ne répond pas ne prouve pas une panne : ICMP est souvent filtré.
  Confirmez toujours avec un test **TCP** sur le port réel (§79).

---

## 78. Tracert et diagnostic de routage

```cmd
:: Traceroute classique (lent : résout chaque saut)
tracert 10.10.30.5

:: Sans résolution DNS (beaucoup plus rapide, à préférer en diagnostic)
tracert -d 10.10.30.5
```

```powershell
# Équivalent PowerShell (nécessite souvent une élévation / module)
Test-NetConnection -ComputerName 10.10.30.5 -TraceRoute
```

Lecture : le saut où les `*` commencent = le point de rupture (ou un équipement
qui ne répond pas à ICMP — vérifiez que les sauts suivants répondent avant de
conclure). Comparez avec un `tracert` depuis une machine qui fonctionne : une
**route asymétrique ou changée** après une modification réseau est un classique.

Table de routage locale (routes persistantes parasites après un VPN mal
désinstallé, par exemple) :

```cmd
route print -4
:: Supprimer une route persistante fautive
route delete 10.99.0.0
```

---

## 79. Test-NetConnection : le couteau suisse

