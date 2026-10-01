---
id: collect-261001-rattrapage/rattrapage/active-directory-guide-18
title: "Active Directory & GPO en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: ["2024-01-09", "2029-01-09"]
keywords: ["attention"]
source: docs/RAG/collect-261001-rattrapage/active_directory_guide.md
source_anchor: ""
source_lines: [2501, 2610]
sha256: db1c6d690703194a3d0db3e8f2525692c2b88812f36d6dbbd851b26b024b1d74
---

# Active Directory & GPO en entreprise — Guide technique ultra-complet

### Cas n°14 — Sauvegarder et restaurer une GPO critique
**Contexte** : modification risquée de la GPO « Proxy entreprise ».
**Objectif** : pouvoir revenir en arrière en 2 minutes.
```powershell
Backup-GPO -Name "USER - Tous - Proxy entreprise" -Path "E:\Sauvegardes\GPO" `
  -Comment "Avant changement proxy $(Get-Date -Format 'yyyy-MM-dd')"
# ... modification ...
# En cas de problème :
Restore-GPO -Name "USER - Tous - Proxy entreprise" -Path "E:\Sauvegardes\GPO\<GUID_sauvegarde>"
gpupdate /force
```
**Commentaire** : la sauvegarde GPO ne sauvegarde que la GPO, pas son lien ni son filtrage — documentez-les aussi (export `Get-GPOReport -ReportType Html`).

### Cas n°15 — Préparer la montée de niveau fonctionnel vers 2016/2025
**Contexte** : dernier DC 2012 R2 décommissionné, parc en 2019/2022.
**Objectif** : monter le niveau fonctionnel.
**Étapes** : 1) `Get-ADForest/Domain` : vérifier les niveaux. 2) `dcdiag /e`, `repadmin /replsummary` : santé parfaite. 3) Sauvegarde état système. 4) `Set-ADDomainMode` puis `Set-ADForestMode` (irréversible). 5) Activer la corbeille si ce n'est pas fait. 6) Documenter.
**Commentaire** : ne montez le niveau que quand **tous** les DC sont au bon OS et que vous êtes sûr de ne pas devoir réintroduire un vieux DC.

### Cas n°16 — Mettre en place le rapport de santé quotidien
**Contexte** : pas de supervision AD, on découvre les pannes par les utilisateurs.
**Objectif** : mail quotidien automatique.
**Solution** : script de la section 73 déposé sur le serveur d'administration, tâche planifiée quotidienne 7h avec un gMSA, envoi vers l'équipe exploitation. Ajouter une alerte immédiate (pas quotidienne) sur les ID critiques 2108/16650/2162 via le planificateur (« déclencher sur un événement »).
**Commentaire** : commencez simple (mail quotidien), affinez ensuite (filtrage du bruit, intégration au SIEM/Wazuh). Un rapport que personne ne lit ne sert à rien : désignez un lecteur responsable par roulement.

---

## 80. Pour aller plus loin

**Documentation officielle** :
- Microsoft Learn : « Active Directory Domain Services Overview » — la référence, à jour pour Server 2025.
- Référence des cmdlets : module `ActiveDirectory` et `GroupPolicy` (docs Microsoft).

**Sécurisation avancée** :
- **Tiering administratif** (modèle ESAE / « Red Forest » simplifié) : comptes admin par niveau (Tier 0 = DC), jamais de compte Tier 0 sur un poste utilisateur.
- **LAPS** (Windows Local Administrator Password Solution, intégré depuis 2023) : mots de passe admin locaux uniques et rotatifs.
- **Protected Users** : groupe imposant Kerberos durci (pas de NTLM, pas de délégation).
- **Authentication Policies & Silos** : restreindre où un compte peut s'authentifier (2012 R2+).
- Durcissement LDAP : exiger **LDAP signing** et **LDAPS** (avis Microsoft ADV190023).

**Outils complémentaires** :
- **PingCastle** : audit de sécurité AD gratuit, rapport actionnable — à lancer trimestriellement.
- **ADRecon / AD ACL Scanner** : inventaire et revue des ACL.
- **Specops / Lithnet** : protection contre les mots de passe compromis.
- **Wazuh** (votre guide existant) : centraliser les journaux AD et alerter sur les ID critiques de la section 74.

**Montée en compétence** :
- Labo : 2-3 VM (un DC 2022, un DC 2025, un client Windows 11) sous Hyper-V/Proxmox — rejouez chaque cas pratique de la section 79.
- Exercices : cassez volontairement (GPO filtrée, heure décalée, DC isolé) puis réparez avec ce guide.
- Veille : blog « Directory Services » Microsoft, notes de version de Windows Server.

**Sujets connexes à vos autres guides** : la supervision AD via **Zabbix/Prometheus** (compteurs NTDS), les sauvegardes avec **Veeam/Borg** (complément de wbadmin), l'automatisation avec **Ansible** (module `microsoft.ad`), le durcissement avec votre guide **Debian/Ubuntu** pour les serveurs membres Linux joints au domaine (SSSD/realmd).

---

## 81. Checklists de mise en production

### Nouveau domaine / nouvelle forêt
- [ ] Nom de domaine choisi définitivement (sous-domaine d'un nom public possédé).
- [ ] Niveau fonctionnel cible décidé.
- [ ] Premier DC promu, DNS intégré, GC actif.
- [ ] `redircmp` exécuté vers `OU=Postes`.
- [ ] Corbeille AD activée.
- [ ] Mots de passe DSRM au coffre.
- [ ] Arborescence OU créée et protégée.
- [ ] Groupes de base (AGDLP) créés.
- [ ] Default Domain Policy configurée (mots de passe, verrouillage, audit de base).
- [ ] PSO Admins créé.
- [ ] Sauvegarde état système planifiée + testée.
- [ ] Rapport de santé quotidien en place.
- [ ] Documentation (FSMO, topologie, schéma réseau) au wiki.

### Nouveau DC
- [ ] `dcdiag /e` propre avant.
- [ ] Promotion réussie, redémarrage OK.
- [ ] `repadmin /showrepl` : réplication entrante/sortante OK.
- [ ] DNS : enregistrements `_msdcs` présents.
- [ ] GC activé si prévu.
- [ ] SYSVOL partagé (`net share`).
- [ ] DHCP/DNS clients mis à jour.
- [ ] Ajouté à la supervision et aux sauvegardes.

### Nouvelle GPO
- [ ] Objectif unique, nommage conforme, commentaire renseigné.
- [ ] Testée sur OU pilote.
- [ ] `Backup-GPO` effectué.
- [ ] Filtrage de sécurité vérifié (lecture ordinateurs !).
- [ ] `gpresult /r` sur poste pilote : appliquée.
- [ ] Déploiement progressif, puis généralisation.
- [ ] Entrée au journal des changements.

---

## 82. Annexe : versions et niveaux fonctionnels

| Version | Niveau fonctionnel max | Points d'attention |
|---|---|---|
| Windows Server 2019 | 2016 (`WinThreshold`) | Fin du support mainstream : 09/01/2024 ; support étendu jusqu'au 09/01/2029. |
| Windows Server 2022 | 2016 (`WinThreshold`) | Pas de nouveau niveau fonctionnel ; améliorations sécurité (SMB AES-256, DNS-over-HTTPS). |
| Windows Server 2025 | 2025 (`Win2025`) ⚠️ | Nouveau niveau : base NTDS en pages 32K, NUMA, DLNA... Ne montez à `Win2025` que si tous les DC sont en 2022+ et que vous n'avez plus besoin de compatibilité antérieure. |

⚠️ **Avertissement version** : le niveau fonctionnel `Win2025` est irréversible et encore jeune au moment de la rédaction (2026) — validez en labo avant production, surtout si des applications tierces étendent le schéma.

**Matrice de compatibilité rapide** : un DC 2025 peut cohabiter avec des DC 2019/2022 en niveau 2016. La montée vers `Win2025` exige l'absence de DC 2019 et antérieurs.

**Fin de vie à surveiller** : planifiez le retrait des DC 2019 avant janvier 2029 (fin du support étendu). Inscrivez-le à votre feuille de route.

---

*Fin du guide. Gardez ce document à jour : chaque changement d'architecture (nouveau site, nouveau niveau fonctionnel, nouvelle PSO) mérite une mise à jour des sections concernées et de la documentation d'exploitation.*
