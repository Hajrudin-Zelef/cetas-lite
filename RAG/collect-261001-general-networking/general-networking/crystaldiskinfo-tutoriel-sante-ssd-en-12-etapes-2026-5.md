---
id: collect-261001-general-networking/general-networking/crystaldiskinfo-tutoriel-sante-ssd-en-12-etapes-2026-5
title: "Calculer l'empreinte SHA-256 de l'installeur téléchargé"
domain: general-networking
role: reference
task: reference
actors: ["Microsoft", "Samsung"]
dates: []
keywords: ["attention", "gpu", "open source"]
source: docs/RAG/collect-261001-general-networking/crystaldiskinfo-tutoriel-sante-ssd-en-12-etapes-2026.md
source_anchor: ""
source_lines: [259, 336]
sha256: df501bd80fede66940cc7249a7e92d7ddeba1f7f9b87995d8696be156cbcb15e
---

# Calculer l'empreinte SHA-256 de l'installeur téléchargé

```
# watchdog-smart.ps1 -- analyse le rapport CrystalDiskInfo et alerte
$rapport = "C:\Outils\CrystalDiskInfo\DiskInfo.txt"
$contenu = Get-Content $rapport -Encoding UTF8
# Extraire les lignes "Etat de sante" de chaque disque
$etats = $contenu | Select-String -Pattern "Etat de sante|Health Status"
$alertes = @()
foreach ($ligne in $etats) {
    if ($ligne -match "Attention|Caution|Mauvais|Bad") {
        $alertes += $ligne.Line.Trim()
    }
}
if ($alertes.Count -gt 0) {
    $corps = "Disque(s) a risque sur $env:COMPUTERNAME :`n`n" + ($alertes -join "`n")
    Send-MailMessage -SmtpServer "smtp.votre-domaine.fr" -Port 587 -UseSsl `
        -From "[email protected]" -To "[email protected]" `
        -Subject "[SMART] Alerte disque sur $env:COMPUTERNAME" -Body $corps `
        -Credential (Import-Clixml "C:\Outils\smtp-cred.xml")
    Write-Host "ALERTE envoyee : $($alertes.Count) disque(s) concerne(s)."
} else {
    Write-Host "Tous les disques sont en bon etat."
}
```
Enchaînez ce watchdog juste après l’export nocturne (une seconde tâche planifiée à 03h10, par exemple). Vous obtenez alors une supervision de bout en bout, gratuite et sans serveur : CrystalDiskInfo lit le matériel, le batch exporte, PowerShell analyse et alerte. Pour sécuriser les identifiants SMTP, stockez-les chiffrés avec `Get-Credential | Export-Clixml "C:\Outils\smtp-cred.xml"` – le fichier n’est déchiffrable que par le compte qui l’a créé. Ce montage, éprouvé, se déploie identiquement sur dix ou mille postes.

## Erreurs fréquentes à éviter avec CrystalDiskInfo

Même un outil aussi simple réserve des pièges. Voici les cinq erreurs qui reviennent le plus souvent, et comment les contourner :

- **Confondre valeur brute et valeur normalisée.** La colonne normalisée (Actuel/Pire/Seuil) est trompeuse : « 100 » n’y veut pas dire « 100 % de santé ». Fiez-vous à la colonne*Valeurs brutes* , en clair, pour interpréter un attribut.
- **Paniquer devant un état « Attention » stable.** Quelques secteurs réalloués figés depuis des mois ne justifient pas un remplacement immédiat. C’est la*progression* du compteur qui compte, pas sa simple présence.
- **Télécharger depuis un site tiers.** Les portails de téléchargement injectent des logiciels indésirables. Passez toujours par la source officielle ou le Microsoft Store.
- **Diagnostiquer via un boîtier USB.** Beaucoup de ponts USB masquent le S.M.A.R.T. et affichent « Inconnu ». Pour un vrai diagnostic, branchez le disque en SATA ou NVMe direct.
- **Oublier .NET Framework 4.8 pour les e-mails.** Sans lui, la fonction Alert Mail reste grisée. Installez-le avant de configurer les alertes.

Une sixième erreur, plus insidieuse, mérite d’être citée : croire que CrystalDiskInfo *répare* un disque. Il ne fait que *lire* et *diagnostiquer*. Il n’effectue ni analyse de surface, ni correction d’erreurs, ni mise à jour de micrologiciel. Un état « Mauvais » ne se « corrige » pas dans le logiciel : il impose une sauvegarde et un remplacement.

## Dépannage : 8 problèmes courants et leurs solutions

Voici les huit situations les plus signalées par les utilisateurs de CrystalDiskInfo, avec la marche à suivre pour chacune.

| Problème | Cause probable | Solution | 
|---|---|---|
| État « Inconnu » (gris) | Pont USB ou RAID masquant le S.M.A.R.T. | Brancher en SATA/NVMe direct ; tester un autre boîtier | 
| Le disque n’apparaît pas du tout | Manque de privilèges administrateur | Relancer en tant qu’administrateur (clic droit) | 
| Aucun SSD NVMe détecté | Windows antérieur à la version 10 | Mettre à jour vers Windows 10/11 | 
| Option Alert Mail grisée | .NET Framework 4.8 absent | Installer .NET Framework 4.8 ou supérieur | 
| E-mail de test qui échoue | Port/SSL ou mot de passe SMTP incorrect | Vérifier port 587/465 ; mot de passe d’application | 
| Température affichée aberrante | Capteur mal interprété (vieux disque/pont) | Recouper avec PowerShell ; ignorer si valeur figée | 
| Attributs vendeur illisibles | Noms d’attributs propriétaires | Consulter la fiche du constructeur ; c’est normal | 
| L’outil ne se lance pas au démarrage | Option « Démarrage » non cochée / bloquée par l’antivirus | Activer Fonction → Démarrage ; ajouter une exception | 

Cas particulier récurrent : après une mise à jour majeure de Windows, l’icône de température peut disparaître de la barre des tâches. Il suffit de rouvrir `Fonction → Notification → Icône de température` pour la restaurer. De même, si vous migrez d’un poste à l’autre, copiez le fichier `DiskInfo.ini` pour retrouver instantanément tous vos seuils et réglages.

## Astuces avancées pour les administrateurs et passionnés

Une fois les bases maîtrisées, plusieurs techniques élèvent votre pratique de la surveillance de stockage :

- **Croiser avec les outils constructeur.** Samsung Magician, WD Dashboard ou Crucial Storage Executive offrent des fonctions (mise à jour de micrologiciel, Over-Provisioning) absentes de CrystalDiskInfo. Utilisez-les en complément, pas en remplacement, du diagnostic S.M.A.R.T.
- **Cloner avant qu’il ne soit trop tard.** Dès qu’un disque passe « Attention » avec un compteur en progression, clonez-le immédiatement (Macrium Reflect, Clonezilla) tant qu’il est encore lisible. Attendre l’état « Mauvais », c’est parier sur la chance.
- **Documenter le TBW à l’achat.** Notez le*Total Host Writes* le jour de l’installation d’un SSD. Vous saurez ensuite précisément votre rythme d’écriture quotidien et pourrez extrapoler la durée de vie réelle.
- **Surveiller les « Unsafe Shutdowns ».** Une hausse rapide de ce compteur trahit un problème d’alimentation ou l’absence d’onduleur – un risque pour l’intégrité des données bien plus immédiat que l’usure.
- **Coupler à une supervision globale.** Sur un homelab, exportez les rapports vers un serveur central. Notre tutoriel Proxmox VE montre comment bâtir l’infrastructure de stockage qui héberge ce type de collecte.

Enfin, gardez à l’esprit la complémentarité des outils. CrystalDiskInfo excelle sur la santé des disques, mais pour le reste du matériel – cartes graphiques, processeurs, tensions – un utilitaire comme HWiNFO prend le relais, et pour nettoyer proprement des pilotes GPU récalcitrants, notre guide DDU reste la référence. La surveillance d’un PC fiable repose sur une boîte à outils, pas sur un logiciel unique.

## CrystalDiskInfo face aux alternatives

CrystalDiskInfo n’est pas seul sur ce créneau. Comprendre ses forces et ses limites face aux alternatives vous aide à choisir le bon outil selon le contexte.

| Outil | Type | Points forts | Limites | 
|---|---|---|---|
| CrystalDiskInfo 9.9.1 | Gratuit, open source | S.M.A.R.T. détaillé, alertes e-mail, résident, CLI | Disques uniquement ; pas de réparation | 
| HWiNFO | Gratuit | Surveillance de tout le système (CPU/GPU/disques) | Moins détaillé sur l’endurance SSD | 
| smartmontools (smartctl) | Gratuit, open source | Multiplateforme, scriptable, très complet | Ligne de commande, courbe d’apprentissage | 
| Outils constructeur (Magician, WD Dashboard…) | Gratuit | Mise à jour firmware, Over-Provisioning | Limités à une seule marque | 
| Hard Disk Sentinel | Payant (version gratuite bridée) | Prédiction de durée de vie, historique | Fonctions clés payantes | 

