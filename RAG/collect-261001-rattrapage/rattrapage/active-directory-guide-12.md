---
id: collect-261001-rattrapage/rattrapage/active-directory-guide-12
title: "Active Directory & GPO en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["agent", "agents", "arr", "attribution"]
source: docs/RAG/collect-261001-rattrapage/active_directory_guide.md
source_anchor: ""
source_lines: [1825, 1988]
sha256: 45df2d8529998f35c9ec4782de3b8194e4a051893b46bf38d7976e3c8ddb30ad
---

# Active Directory & GPO en entreprise — Guide technique ultra-complet

- Action **Mettre à jour** (pas Remplacer, qui supprime/recrée à chaque fois = lent).
- Un élément = un ciblage clair.
- Évitez les chemins en dur vers des serveurs uniques : utilisez des **noms DNS d'espace de noms DFS** (`\\ad.entreprise.fr\partages\commun`).

---

## 64. Ciblage au niveau élément (Item-Level Targeting)

Le **ciblage** applique un élément GPP uniquement si des conditions sont vraies : groupe, OU, site, adresse IP, version d'OS, jour/heure, état du registre, etc. Conditions combinables en ET/OU.

**Exemples concrets** :

| Besoin | Ciblage |
|---|---|
| Imprimante Lyon uniquement | « Le site AD est Site-Lyon » |
| Lecteur compta pour les comptables | « L'utilisateur est membre de GG_Compta » |
| Raccourci VPN sur les portables | « Le nom d'ordinateur commence par LT- » (via requête LDAP/WMI) |
| Paramètre hors heures ouvrées | Plage de dates/heures |

C'est ce qui permet **une seule GPO « Mappages »** pour toute l'entreprise avec 10 éléments ciblés, plutôt que 10 GPO. Lisible, maintenable.

---

## 65. GPO de sécurité : audit et droits utilisateur

**Audit** (`Configuration ordinateur > Stratégies > Paramètres Windows > Paramètres de sécurité > Stratégies d'audit`) : activez a minima (recommandation) :

| Stratégie | Succès | Échec |
|---|---|---|
| Ouverture/fermeture de session | ✅ | ✅ |
| Gestion des comptes | ✅ | ✅ |
| Accès aux objets (partages sensibles) | ✅ | ✅ |
| Modification de stratégie | ✅ | ✅ |
| Utilisation des privilèges | ❌ (trop verbeux) | ✅ |

Sur 2019+ : utilisez les **sous-catégories d'audit avancé** (`Configuration avancée de la stratégie d'audit`) pour un contrôle fin.

**Attribution des droits utilisateur** (`Stratégies locales > Attribution des droits utilisateur`) :

- « Ouvrir une session locale » : restreignez sur les serveurs (pas les utilisateurs lambda sur un serveur applicatif !).
- « Arrêter le système », « Sauvegarder fichiers et répertoires » : cercle restreint.
- « Ajouter des stations de travail au domaine » : par défaut 10 pour tout utilisateur authentifié — **limitez** via ce paramètre ou `ms-DS-MachineAccountQuota`.

```powershell
# Limiter la jonction de domaine à 0 pour les utilisateurs standard (seuls les admins/Helpdesk via délégation)
Set-ADDomain -Identity "ad.entreprise.fr" -Replace @{ "ms-DS-MachineAccountQuota" = 0 }
```

---

## 66. Déploiement de logiciels via GPO (MSI)

`Configuration ordinateur (ou utilisateur) > Stratégies > Installation de logiciel` : déploie des **MSI** (ou ZAP, obsolète).

- **Attribué (Assigned)** : installé au démarrage (ordinateur) ou à l'ouverture (utilisateur), **obligatoire**.
- **Publié (Published)** : disponible dans « Programmes et fonctionnalités » / ajout de programmes, **optionnel** (utilisateur uniquement).

```powershell
# 1. Placer le MSI sur un partage accessible en LECTURE par "Ordinateurs du domaine"
#    (ou "Utilisateurs authentifiés" pour un déploiement utilisateur)
# 2. Dans la GPO : Nouveau > Package > chemin UNC (JAMAIS de lettre de lecteur !)
#    \\srv-logiciels\deploy\agent_inventaire.msi
# 3. Méthode de déploiement : Attribué
```

**Bonnes pratiques** :

- Testez le MSI en silencieux d'abord : `msiexec /i package.msi /qn`.
- Utilisez les **transformations MST** pour pré-configurer (clés, options).
- **Mises à jour** : nouvelle GPO ou « mise à niveau » du package dans la même GPO.
- **Désinstallation** : « Désinstaller quand hors du périmètre » pour le nettoyage automatique.
- Pour du lourd (Office, etc.) : préférez SCCM/Intune/PDQ au déploiement GPO, limité aux petits agents/outils.
- Surveillez les **démarrages lents** : un MSI attribué ordinateur retarde le boot ; planifiez.

---

## 67. Scripts de démarrage / arrêt / ouverture / fermeture de session

| Script | Contexte | Déclenchement |
|---|---|---|
| Démarrage | **Ordinateur** (SYSTEM) | Boot, avant l'ouverture de session |
| Arrêt | Ordinateur (SYSTEM) | Extinction |
| Ouverture de session | **Utilisateur** | Logon |
| Fermeture de session | Utilisateur | Logoff |

Emplacement GPO : `Configuration ordinateur/utilisateur > Stratégies > Paramètres Windows > Scripts`.

```powershell
# Exemple de script de démarrage : installer l'agent d'inventaire si absent
# startup_inventaire.ps1 (déployé via la GPO, exécution SYSTEM)
if (-not (Test-Path "C:\Program Files\Inventaire\agent.exe")) {
    Start-Process msiexec.exe -ArgumentList '/i "\\srv-logiciels\deploy\agent.msi" /qn' -Wait
}
```

**PowerShell vs Batch** : préférez PowerShell ; réglez la **stratégie d'exécution** via GPO (`Activer l'exécution de scripts` = `AllSigned` ou `RemoteSigned` selon votre maturité de signature).

⚠️ **Délais** : les scripts synchrones retardent boot/logon. Mettez un **timeout**, rendez-les idempotents, et loguez dans un partage central pour le dépannage.

---

## 68. Dépannage GPO : gpresult, rsop.msc, journaux

```powershell
# Rapport des stratégies appliquées (utilisateur + ordinateur)
/gpresult /r
gpresult /r /scope:computer
gpresult /r /scope:user

# Rapport HTML complet (le plus utile à distance)
gpresult /h C:\Temp\gpresult.html /f

# Forcer l'actualisation
gpupdate /force
gpupdate /force /target:computer

# Forcer à distance (sans déranger l'utilisateur si /target:computer)
Invoke-GPUpdate -Computer "PC-ATELIER-042" -Target "Computer" -Force
```

**rsop.msc** : jeu de stratégies résultant, vue console. **GPMC > Modélisation de stratégie de groupe** : simule l'application pour un utilisateur/poste (idéal pour « que se passerait-il si... »).

**Journaux** : `Observateur d'événements > Journaux des applications et des services > Microsoft > Windows > GroupPolicy > Operational`. Événements clés : **1500-1503** (traitement), **1058/1030** (SYSVOL inaccessible — souvent DNS ou DFSR).

**Test de connectivité SYSVOL** : `\\ad.entreprise.fr\SYSVOL\ad.entreprise.fr\Policies` doit s'ouvrir depuis le poste.

---

## 69. GPO qui ne s'applique pas : méthodologie complète

Checklist ordonnée (dans 90 % des cas, le problème est dans les 5 premiers) :

1. [ ] **Liaison** : la GPO est-elle liée au bon niveau (domaine/OU/site) ? `Get-GPInheritance`.
2. [ ] **Objet au bon endroit** : l'utilisateur/le PC est-il vraiment dans l'OU ciblée ? (Erreur classique : objet dans `CN=Computers`.)
3. [ ] **Filtrage de sécurité** : le groupe a-t-il `Appliquer` + `Lecture` ? Les ordinateurs ont-ils `Lecture` ? (section 59, piège MS16-072.)
4. [ ] **Filtre WMI** : la requête retourne-t-elle vrai sur le poste ? Testez avec `Get-WmiObject`.
5. [ ] **Héritage bloqué / Enforced** : `Get-GPInheritance -Target ...` montre `GpoInheritanceBlocked`.
6. [ ] **Moitié désactivée** : la GPO a-t-elle la moitié (ordinateur/utilisateur) dont vous avez besoin d'activée ?
7. [ ] **Réplication SYSVOL** : le paramètre existe-t-il sur le SYSVOL du DC qui sert le client ? (`dfsrdiag`, comparez deux DC.)
8. [ ] **gpupdate / reboot** : les stratégies ordinateur nécessitent souvent un redémarrage ; `gpupdate /force` ne suffit pas toujours.
9. [ ] **Bouclage** : actif sur ce poste ? (section 61.)
10. [ ] **Délai de réplication AD** : GPO créée sur DC01, poste authentifié sur DC02 pas encore répliqué → attendez ou forcez `repadmin /syncall`.
11. [ ] **Journaux** : GroupPolicy Operational, événements 1058/1030/1500.

```powershell
# Script de diagnostic rapide à lancer sur le poste en cause
gpresult /r
nltest /dsgetsite
Test-Path "\\ad.entreprise.fr\SYSVOL\ad.entreprise.fr\Policies"
```

---

## 70. Bonnes pratiques de nommage et gestion du cycle de vie des GPO

**Nommage** : `[PÉRIMÈTRE] - [Objet] - [Action]` :

- `ORDI - Tous - Verrouillage session 10 min`
- `USER - Exploitation - Mappage lecteur S`
- `SECU - Serveurs - Audit avancé`
- `SRV-RDS - Bouclage bureau verrouillé`

**Cycle de vie** :

