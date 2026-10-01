---
id: collect-261001-general-networking/general-networking/crystaldiskinfo-tutoriel-sante-ssd-en-12-etapes-2026-3
title: "Calculer l'empreinte SHA-256 de l'installeur téléchargé"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["attention", "nand"]
source: docs/RAG/collect-261001-general-networking/crystaldiskinfo-tutoriel-sante-ssd-en-12-etapes-2026.md
source_anchor: ""
source_lines: [100, 160]
sha256: 7c903b0ae6cb567e663b9fd4af05e3b98d3a6d77abc8da512c7b46348533bf55
---

# Calculer l'empreinte SHA-256 de l'installeur téléchargé

| ID | Attribut | Ce qu’il révèle | Valeur brute idéale | 
|---|---|---|---|
| 05 | Secteurs réalloués | Secteurs défectueux remplacés par des secteurs de réserve | 0 (toute hausse est un signal d’alerte) | 
| C5 | Secteurs en attente (Current Pending) | Secteurs instables en cours de surveillance | 0 | 
| C6 | Secteurs non corrigibles (Uncorrectable) | Secteurs définitivement illisibles | 0 | 
| C4 | Événements de réallocation | Nombre de tentatives de réallocation | 0 | 
| 0A | Nouvelles tentatives de rotation (Spin Retry) | Difficulté du moteur à atteindre sa vitesse | 0 | 
| C2 | Température | Température du disque en °C | < 45 °C | 
| 09 | Heures de fonctionnement | Durée d’utilisation cumulée | Indicatif (usure globale) | 

Les attributs 05, C5 et C6 forment le **trio de la mort** pour un disque dur : dès que l’un d’eux quitte le zéro et progresse, planifiez le remplacement. À l’inverse, un compteur « Heures de fonctionnement » élevé n’est pas alarmant en soi – un disque peut tourner 60 000 heures sans un seul secteur réalloué. Pour un SSD (SATA ou NVMe), la grille change complètement, car il n’y a plus de pièces mobiles ; c’est l’endurance d’écriture qui prime :

| Attribut (SSD / NVMe) | Ce qu’il révèle | À surveiller | 
|---|---|---|
| Percentage Used (NVMe) | Usure estimée par le contrôleur (0 % = neuf, 100 % = fin de garantie d’endurance) | Approche de 100 % | 
| Total Host Writes / TBW | Volume total écrit depuis l’origine | Comparer au TBW garanti du constructeur | 
| Available Spare (NVMe) | Réserve de blocs de secours restante | Chute sous le « Spare Threshold » | 
| Wear Leveling Count (SSD SATA) | Cycles d’effacement moyens des cellules | Décroît vers son seuil | 
| Unsafe Shutdowns | Coupures d’alimentation brutales | Hausse anormale (alim/onduleur) | 
| Media and Data Integrity Errors | Erreurs d’intégrité des données | Toute valeur > 0 | 

## Étape 6 – Surveiller la température de vos disques

La température est un tueur silencieux de stockage, en particulier pour les SSD NVMe M.2 logés sous une carte graphique ou dans un ordinateur portable exigu. CrystalDiskInfo affiche la température instantanée dans son cartouche dédié, et l’attribut C2 (ou « Composite Temperature » en NVMe) en conserve l’historique haut/bas. Les plages à retenir : en dessous de 40 °C, tout va bien ; entre 45 et 55 °C, c’est acceptable sous charge ; au-delà de 60-70 °C, un SSD NVMe déclenche généralement une *limitation thermique* (throttling) qui réduit ses performances pour se protéger.

CrystalDiskInfo peut afficher la température en permanence dans la **zone de notification** de Windows (barre des tâches), sous forme d’icône numérique colorée. Activez cette option depuis `Fonction → Notification → Icône de température`. Vous obtenez alors une surveillance passive : un coup d’œil à la barre des tâches suffit à repérer un disque qui chauffe anormalement. Si vous constatez des pics récurrents, la solution passe souvent par un dissipateur M.2 (radiateur) ou une meilleure ventilation du boîtier – pas par le logiciel.

Petite subtilité d’affichage : certains disques rapportent deux capteurs de température (par exemple le contrôleur et la NAND sur un NVMe). CrystalDiskInfo affiche alors la valeur la plus élevée ou les deux selon le modèle. Ce n’est pas une anomalie ; c’est la richesse du télémétrage du disque qui transparaît.

## Étape 7 – Régler les seuils d’alerte de santé et de température

Par défaut, CrystalDiskInfo se contente d’afficher l’état. Pour qu’il vous *alerte* activement, configurez les seuils. Ouvrez `Fonction → Paramètres avancés` : vous y trouverez la définition des seuils de santé et de température par disque. Vous pouvez par exemple décider qu’un disque doit passer en « Attention » dès le premier secteur réalloué, ou qu’une alarme se déclenche au-delà de 55 °C.

Ces seuils personnalisés sont particulièrement utiles pour un parc hétérogène. Sur un serveur de fichiers, vous serez plus strict (alerte à 50 °C, tolérance zéro sur les secteurs en attente) ; sur un vieux PC secondaire, vous pourrez lever le pied. La documentation complète de ces réglages figure dans le manuel des fonctionnalités avancées de Crystal Dew World. Une fois les seuils définis, associez-leur une réaction : une alarme sonore, une icône, ou – le plus utile – une notification par e-mail, que nous configurons à l’étape suivante.

## Étape 8 – Activer les alertes par e-mail (Alert Mail)

C’est la fonction qui transforme CrystalDiskInfo d’un outil de consultation en véritable sentinelle. La fonction **Alert Mail** envoie un courriel automatique dès qu’un seuil est franchi. Prérequis impératif : **.NET Framework 4.8 ou supérieur** doit être installé, faute de quoi l’option reste grisée. Rendez-vous dans `Fonction → Paramètres avancés → Notification par e-mail` et renseignez les paramètres SMTP de votre fournisseur.

Voici un exemple de configuration SMTP fonctionnelle, ici avec un serveur classique en TLS sur le port 587. Adaptez-la à votre fournisseur (Gmail exige un mot de passe d’application, OVH ou Infomaniak utilisent leurs propres hôtes SMTP) :

```
Serveur SMTP    : smtp.votre-domaine.fr
Port            : 587 (STARTTLS) ou 465 (SSL)
Authentification: activee
Utilisateur     : [email protected]
Mot de passe    : ******** (mot de passe d'application si Gmail/Outlook)
Expediteur      : [email protected]
Destinataire    : [email protected]
Objet           : [CrystalDiskInfo] Alerte disque sur POSTE-01
```
Après la saisie, utilisez impérativement le bouton **« Envoyer un e-mail de test »**. Il valide en direct la connexion SMTP et vous évite de découvrir, le jour d’une vraie panne, que l’authentification échouait silencieusement. Une fois le test reçu, votre poste enverra automatiquement un courriel à chaque passage en « Attention » ou « Mauvais ». Pour un administrateur gérant plusieurs sites en France ou en Europe, centraliser ces alertes sur une boîte dédiée (par exemple `stockage-alertes@`) crée un tableau de bord de fait, sans aucun serveur de supervision supplémentaire.

## Étape 9 – Activer le mode Résident et le lancement au démarrage

Une alerte e-mail ne sert à rien si le logiciel n’est pas ouvert. Le mode **Résident** résout ce problème : CrystalDiskInfo se réduit dans la zone de notification et continue de surveiller les disques en arrière-plan, à intervalle régulier. Activez-le via `Fonction → Résident`. Combinez-le avec `Fonction → Démarrage` pour que l’outil se lance automatiquement à l’ouverture de session Windows.

Réglez ensuite la fréquence de vérification via `Fonction → Paramètres avancés → Actualisation automatique`. Un intervalle de 10 minutes constitue un bon compromis pour un poste de travail ; sur un serveur, 30 à 60 minutes suffisent et allègent l’accès disque. Attention : une actualisation trop fréquente (toutes les minutes) réveille en permanence les disques durs mécaniques qui tentent de se mettre en veille, ce qui use inutilement le mécanisme. Sur SSD, l’impact est négligeable.

En version portable, ces trois réglages (Résident, Démarrage, Actualisation) s’inscrivent dans le fichier `DiskInfo.ini`. Voici à quoi ressemble une configuration type prête pour un déploiement automatisé – vous pouvez la distribuer telle quelle sur un parc :

