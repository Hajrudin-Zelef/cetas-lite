---
id: collect-261001-rattrapage/rattrapage/maintenance-windows-guide-20
title: "Maintenance et exploitation Windows en entreprise"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["incident"]
source: docs/RAG/collect-261001-rattrapage/maintenance_windows_guide.md
source_anchor: ""
source_lines: [3733, 3915]
sha256: 30a481f4c119a379614aa92730e6d7bca970505bae6447d220e09e79d47d720a
---

# Maintenance et exploitation Windows en entreprise

Modèle de procédure (1 page recto idéalement, 3 max) :

```text
TITRE : Redémarrage planifié du serveur SRV-FIC-01
OBJECTIF : appliquer les correctifs mensuels sans perte de données
PRÉREQUIS : fenêtre validée, sauvegarde OK, information utilisateurs
ÉTAPES :
  1. Vérifier qu'aucun job n'est en cours : ...
  2. Notifier : msg * /SERVER:srv-fic-01 "..."
  3. Installer les MAJ : ...
  4. Redémarrer, vérifier les services : Get-Service ...
  5. Contrôle post : partages accessibles, journal Système sans erreur
RETOUR ARRIÈRE : si échec > 30 min -> rollback §20 / restauration §124
TRACÉ : ticket n°, heure début/fin, opérateur
```

**Principes** : impératif, commandes copiables, critères de succès explicites,
procédure de retour arrière **systématique**, nom du rédacteur + date de dernière
relecture. Une procédure non relue depuis 2 ans est une procédure fausse.

---

## 130. Gestion des changements

Tout changement hors routine suit un mini-processus (proportionné à votre taille) :

| Étape | Contenu |
|---|---|
| **Demande** | Quoi, pourquoi, impact, risque |
| **Validation** | Chef de service (ou comité selon criticité) |
| **Planification** | Fenêtre (§4), communication, retour arrière |
| **Exécution** | Par la personne désignée, avec la procédure (§129) |
| **Vérification** | Tests post-changement, supervision |
| **Clôture** | Ticket complété, doc mise à jour, retour d'expérience si incident |

**Changements standard** (pré-approuvés) : patch mensuel selon anneaux (§11),
redémarrage planifié, ajout d'utilisateur — procédure connue, pas de validation
au cas par cas. **Changements normaux** : migration, nouveau rôle, modification
GPO sensible — validation requise. **Urgents** : incident de sécurité, panne P1 —
validation a posteriori sous 24 h, mais traçabilité immédiate.

---

## 131. Passation et continuité (angle chef de service)

Le service ne doit jamais dépendre d'une seule personne (« bus factor ») :

- [ ] **Coffre-fort des accès** : mots de passe admin, DSRM, certificats, clés —
      dans un coffre chiffré d'équipe, jamais dans la tête d'un seul.
- [ ] **Binômes** : chaque système critique a un référent **et** un suppléant
      formé (qui a déjà fait la manip, pas juste lu la doc).
- [ ] **Passation** : à chaque départ/arrivée, revue des DTI (§128), des tâches
      planifiées (§47) et des accès (révoquer le jour J).
- [ ] **Astreinte** : planning écrit, téléphone d'astreinte, procédure d'escalade
      (qui appeler à 3h du matin pour quoi), compensation.
- [ ] **Journal de bord** : un canal unique (ticket/OneNote/wiki) où l'équipe note
      interventions et observations — relu en revue hebdo (§7).

> Test annuel : « si je suis injoignable pendant 2 semaines, l'équipe tourne-t-elle ? »
> Si la réponse est non, la priorité n°1 n'est pas technique, elle est
> organisationnelle.

---

## 132. 20 erreurs classiques d'exploitation Windows

1. **Ne jamais redémarrer les serveurs** (« uptime 400 jours ») → patchs non
   finalisés, fuites mémoire, pannes surprises. Redémarrage mensuel minimum (§8).
2. **`Win32_Product` en inventaire** → reconfiguration MSI en cascade (§116).
3. **WSUS « tous produits, toutes langues »** → base obèse, console inutilisable (§13).
4. **Désactiver IPv6 en décochant la carte** → composants Windows cassés (§98).
5. **Réactiver SMBv1** pour un vieux NAS → faille de sécurité ; mettre à jour
   l'équipement (§90).
6. **Purger WinSxS à la main** → système instable ; utiliser DISM (§28-29).
7. **`chkdsk /r` lancé à 14h sur un volume de prod** → indisponibilité de plusieurs
   heures (§25).
8. **Mot de passe en dur dans un script** → utilisez gMSA (§46) ou un coffre.
9. **Tâche planifiée sous un compte utilisateur** → échec à la prochaine
   expiration de mot de passe (§45).
10. **Sauvegarde jamais testée** → découverte le jour du sinistre (§127).
11. **Snapshot de DC restauré sans procédure** → USN rollback, AD corrompu (§126).
12. **`/ResetBase` avant validation du patch** → rollback impossible (§29).
13. **DNS externe en premier sur un poste du domaine** → domaine intermittent (§82).
14. **Tous les DC patchés la même nuit** → aucun annuaire pendant l'opération (§11).
15. **Journal Sécurité trop petit** → écrasé en 3 jours, aucune enquête possible (§41).
16. **Driver Verifier oublié activé** → BSOD en production (§74).
17. **Pas de point de restauration avant un changement pilote** (§67).
18. **Ignorer les alertes** (« oui on sait, le disque est plein depuis 2 mois ») →
    l'alerte fatigue tue la supervision (§107).
19. **Documenter nulle part** → le savoir part avec la personne (§128, §131).
20. **Patcher sans fenêtre ni communication** → utilisateurs bloqués, confiance
    perdue (§4, §19).

---

## 133. Pense-bête des commandes PowerShell

```powershell
# --- Système & santé ---
sfc /scannow
DISM /Online /Cleanup-Image /RestoreHealth
Get-ComputerInfo | Select-Object WindowsProductName, OsBuildNumber
Get-HotFix | Sort-Object InstalledOn -Descending | Select-Object -First 5
systeminfo | Select-String 'Heure de démarrage'

# --- Disques ---
Get-PSDrive C
Get-Volume | Format-Table DriveLetter, FileSystemType, SizeRemaining, Size
Optimize-Volume -DriveLetter C -ReTrim -Verbose
Get-PhysicalDisk | Get-StorageReliabilityCounter | Format-Table DeviceId, Wear, TemperatureCelsius

# --- Journaux ---
Get-WinEvent -LogName System -MaxEvents 20
Get-WinEvent -FilterHashtable @{LogName='System'; Level=2; StartTime=(Get-Date).AddDays(-1)}
wevtutil qe System /c:10 /rd:true /f:text

# --- Réseau ---
Test-Connection -ComputerName 10.10.20.1 -Count 4
Test-NetConnection -ComputerName srv-fic-01 -Port 445
Resolve-DnsName intranet.contoso.local
Get-NetAdapter | Select-Object Name, Status, LinkSpeed
Get-NetTCPConnection -State Listen | Sort-Object LocalPort

# --- AD / GPO ---
gpupdate /force
gpresult /r ; gpresult /h C:\Temp\gpo.html
nltest /dsgetdc:contoso.local ; nltest /sc_query:contoso.local
Test-ComputerSecureChannel -Repair
w32tm /query /status ; w32tm /resync
repadmin /replsummary ; dcdiag /q

# --- Services & processus ---
Get-Service | Where-Object Status -eq 'Running'
Get-Process | Sort-Object CPU -Descending | Select-Object -First 10
resmon   # analyseur de ressources

# --- Tâches planifiées ---
Get-ScheduledTask | Get-ScheduledTaskInfo | Where-Object LastTaskResult -ne 0

# --- À distance ---
Enable-PSRemoting -Force
Invoke-Command -ComputerName srv-01, srv-02 { Get-Service Spooler }
Enter-PSSession -ComputerName srv-01

# --- Certificats ---
Get-ChildItem Cert:\LocalMachine\My | Select-Object Subject, NotAfter
```

---

## 134. Pense-bête des commandes CMD / classiques

```cmd
:: --- Démarrage / récupération ---
shutdown /r /o /f /t 0        :: redémarrer sur options avancées (WinRE)
reagentc /info                :: état de WinRE
bcdedit /enum                 :: magasin de démarrage
bootrec /fixmbr & bootrec /fixboot & bootrec /rebuildbcd
chkdsk C: /f

:: --- Réseau ---
ipconfig /all & ipconfig /flushdns & ipconfig /renew
ping -t 10.10.20.1 & tracert -d 10.10.30.5 & pathping 10.10.30.5
netstat -ano & arp -a & route print -4
nslookup intranet.contoso.local 10.10.10.11
netsh wlan show interfaces & netsh wlan show wlanreport
netsh winsock reset & netsh int ip reset   :: (puis reboot)

:: --- Système ---
sfc /scannow
DISM /Online /Cleanup-Image /RestoreHealth
cleanmgr /sageset:10 & cleanmgr /sagerun:10
vssadmin list shadows & vssadmin list writers
wbadmin get versions
wecutil qc   :: initier le collecteur d'événements
verifier      :: Driver Verifier (assistant)
mdsched       :: diagnostic mémoire (reboot)
```

---

## 135. Quiz : 10 questions

