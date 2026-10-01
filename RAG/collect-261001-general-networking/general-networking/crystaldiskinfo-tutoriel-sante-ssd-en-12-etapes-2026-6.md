---
id: collect-261001-general-networking/general-networking/crystaldiskinfo-tutoriel-sante-ssd-en-12-etapes-2026-6
title: "Calculer l'empreinte SHA-256 de l'installeur téléchargé"
domain: general-networking
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["attention", "mai", "open source"]
source: docs/RAG/collect-261001-general-networking/crystaldiskinfo-tutoriel-sante-ssd-en-12-etapes-2026.md
source_anchor: ""
source_lines: [337, 375]
sha256: 67fabeb2f5cb6ea82b30b9a35060521a9a907d31e8121bdb6bb641b12c1dfb32
---

# Calculer l'empreinte SHA-256 de l'installeur téléchargé

Le verdict : pour la lecture S.M.A.R.T. sous Windows, CrystalDiskInfo offre le meilleur rapport simplicité/profondeur, entièrement gratuit et sans marque imposée. Les administrateurs Linux ou multiplateformes préféreront `smartctl` (de la suite smartmontools), plus scriptable. Et pour une prédiction de durée de vie plus élaborée avec historique graphique, Hard Disk Sentinel reste une option, moyennant licence. Dans la majorité des cas – poste de travail, PME, homelab – CrystalDiskInfo couvre 100 % du besoin de diagnostic.

### À lire également sur Tech-Insider

## Foire aux questions (FAQ)

### CrystalDiskInfo est-il vraiment gratuit et sûr ?

Oui. CrystalDiskInfo est un logiciel open source et entièrement gratuit, développé par Crystal Dew World, dont le code est publié sur GitHub. Il est sûr à condition de le télécharger depuis la source officielle, SourceForge ou le Microsoft Store. Méfiez-vous uniquement des portails tiers qui empaquettent parfois l’installeur avec des logiciels indésirables.

### Quelle est la dernière version de CrystalDiskInfo en 2026 ?

La version stable la plus récente est CrystalDiskInfo 9.9.2, publiée le 25 juillet 2026 sous la forme d’un installeur de 6,2 Mo et d’une archive portable de 8,5 Mo (build daté du 29 juillet sur Uptodown, noté 4,5/5). Elle succède à la 9.9.1 du 23 mai 2026 – déjà à 159 438 téléchargements hebdomadaires sur SourceForge –, à la 9.9.0 de la mi-mai 2026 et à la 9.8.0 du 16 février 2026 (archive ZIP de 8,1 Mo) – qui elle-même prolongeait la lignée 9.7.x lancée à l’été 2025. L’édition Standard est recommandée pour la plupart des usages ; les éditions Shizuku, Kurei Kei et Aoi, nettement plus lourdes (jusqu’à 65-67 Mo pour l’édition Aoi), n’ajoutent qu’un habillage graphique.

### Que signifie l’état « Attention » (jaune) ?

Il signale qu’un attribut S.M.A.R.T. a franchi un seuil d’alerte – le plus souvent des secteurs réalloués ou en attente sur un disque dur, ou une usure élevée sur un SSD. Ce n’est pas une panne immédiate, mais un avertissement. Surveillez si le compteur progresse : une valeur stable se tolère, une valeur qui augmente impose de sauvegarder et de remplacer le disque.

### CrystalDiskInfo peut-il réparer un disque défectueux ?

Non. CrystalDiskInfo est un outil de diagnostic : il lit et interprète les données S.M.A.R.T., mais n’effectue aucune réparation, analyse de surface ou correction d’erreurs. Un état « Mauvais » impose une sauvegarde immédiate et le remplacement du disque, pas une tentative de réparation logicielle.

### Pourquoi mon SSD affiche-t-il « Inconnu » ou aucune donnée ?

Deux causes principales. Soit vous consultez le disque via un boîtier USB dont le pont ne relaie pas les commandes S.M.A.R.T. – branchez-le alors en SATA ou NVMe direct. Soit vous n’avez pas lancé le logiciel en administrateur, ce qui bloque l’accès matériel de bas niveau. Pour les SSD NVMe, vérifiez aussi que vous êtes sous Windows 10 ou ultérieur.

### Comment surveiller la santé d’un SSD sur le long terme ?

Activez le mode Résident et le lancement au démarrage pour une surveillance continue, réglez une actualisation automatique (10 minutes convient), et configurez des alertes par e-mail. Pour un parc, automatisez l’export via l’option en ligne de commande `/CopyExit` et une tâche planifiée, puis analysez les rapports avec un script PowerShell comme celui présenté dans ce tutoriel.

### Quelle température est dangereuse pour un disque ?

En dessous de 40 °C, tout va bien. Entre 45 et 55 °C, c’est acceptable sous charge. Au-delà de 60-70 °C, un SSD NVMe déclenche une limitation thermique qui réduit ses performances, et une exposition prolongée à ces températures accélère l’usure. Si vous constatez des pics récurrents, ajoutez un dissipateur M.2 ou améliorez la ventilation du boîtier.

### CrystalDiskInfo et CrystalDiskMark, est-ce la même chose ?

Non, ce sont deux outils distincts du même développeur. CrystalDiskInfo mesure la *santé* du disque (données S.M.A.R.T., température, usure). CrystalDiskMark mesure ses *performances* (vitesses de lecture et d’écriture). Ils sont complémentaires : l’un vous dit si le disque va bien, l’autre s’il est rapide.

*Article publié le 03 juin 2026. Les versions logicielles et données techniques citées reflètent les informations disponibles à cette date. Vérifiez toujours l’état de vos sauvegardes avant toute manipulation sur un disque signalé en « Attention » ou « Mauvais ».*
