---
id: collect-261001-rattrapage/rattrapage/huawei-s310-guide-12
title: "Guide ultra-complet — Huawei eKit S310"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["arr", "attention"]
source: docs/RAG/collect-261001-rattrapage/huawei_s310_guide.md
source_anchor: ""
source_lines: [2428, 2643]
sha256: 2abae9140ed01a0d38ce18ae62b7265a769e3a417e3f63cb41ff3eb0edd7946f
---

# Guide ultra-complet — Huawei eKit S310

- **Vue topologique** auto (via LLDP), état des équipements, alertes.
- **Inspection intelligente** : l'app propose des diagnostics guidés.
- **Montées de version** pilotées (Smart Upgrade, section 102).
- Alertes (port down, PoE en surcharge, température...) remontées en push.

⚠️ Le cloud **complète** mais ne **remplace** pas SNMP/syslog locaux : si
Internet tombe, tu perds la supervision cloud mais pas la locale. En site
critique : les deux.

---

## 85. Port mirroring : capturer le trafic pour Wireshark

Pour analyser un problème applicatif (paquets perdus, DHCP qui ne répond pas),
on **miroite** un port vers un port d'analyse où tourne Wireshark.

```
system-view
observe-port 1 interface GigabitEthernet0/0/24   # port d'analyse (PC + Wireshark)
quit
interface GigabitEthernet0/0/5                   # port à observer
 port-mirroring to observe-port 1 inbound
quit
save
```

> Syntaxe (`observe-port`, `port-mirroring`) **à vérifier sur la version
> logicielle du modèle exact**.

⚠️ **Règles :**
- Le port d'analyse ne sert **qu'à ça** pendant la capture (pas de production).
- Ne miroite pas un uplink 10G vers un port 1G : tu perds des paquets.
- **Retire le mirroring après** (`undo`) : ça consomme des ressources et peut
  fausser les stats.

---

## 86. VCT : le testeur de câble intégré (virtual cable test)

Un port qui monte/descend, un débit anormal : avant de changer le câble dans le
faux plafond, fais un **VCT** — le switch mesure la longueur et détecte les
défauts par paire.

```
test cable GigabitEthernet0/0/5
```

Résultat type : longueur par paire, état (OK / ouvert / court-circuit /
désadapté), distance du défaut.

⚠️ Le test **coupe brièvement** le lien : à faire en heure creuse ou sur un port
déjà en panne. Disponibilité **à vérifier sur la version du modèle exact**.

✅ **Usage pro :** VCT systématique quand tu brasses une nouvelle prise avant de
déclarer « la prise est morte ». 50 % des « prises mortes » sont des jarretières
ou des ports du switch mal configurés.

---

## 87. Statistiques de ports et RMON : les compteurs qui parlent

```
display interface GigabitEthernet0/0/5
```

Regarde : **input/output errors, CRC, collisions, giants, runts**. Quelques
erreurs au boot = normal. Des compteurs qui **augmentent en continu** = problème
(câble, duplex, carte réseau distante).

**Remise à zéro pour mesurer :**

```
reset counters interface GigabitEthernet0/0/5
# ... attends 10 minutes de trafic réel ...
display interface GigabitEthernet0/0/5
```

Le S310 supporte aussi **RMON** et la **collecte de statistiques de trafic**
(valeur datasheet) pour l'historique via SNMP.

🔧 **Checklist « port malade » :**
- CRC/errors qui montent → câble ou jarretière (VCT, section 86).
- ` Giants` → MTU / équipement qui envoie des trames surdimensionnées.
- Débit négocié à 100M au lieu de 1G → câble (paires manquantes) ou
  négociation forcée d'un côté.

---

## 88. Sauvegarde de la configuration : 3 méthodes

**Méthode 1 — Fichier local (le minimum) :**

```
save                          # config active → flash (vrpcfg.zip)
display saved-configuration   # vérifie que c'est bien sauvegardé
```

**Méthode 2 — Export vers un serveur (l'automatique) :**

```
# Via TFTP/FTP/SFTP selon la version (à vérifier sur le modèle exact)
tftp 192.168.50.20 put vrpcfg.zip backup/SW-ACC-01_2026-09-27.zip
```

**Méthode 3 — Copier-coller (le rustique qui sauve) :**

```
display current-configuration
```
→ copie tout dans un fichier texte daté. En cas de remplacement urgent, tu
recolles la config sur le switch neuf en 10 minutes.

✅ **Politique :** sauvegarde **automatique hebdo** (script + TFTP/SFTP) +
sauvegarde **manuelle avant chaque changement**. Rétention : 3 mois minimum.
Teste la **restauration** une fois (section 89) — une sauvegarde jamais testée
n'est pas une sauvegarde.

---

## 89. Restauration de configuration : procédure

1. Transfère le fichier de sauvegarde sur le switch (TFTP/FTP, clé USB si
   supporté — à vérifier sur le modèle exact).
2. Deux approches :
   - **Remplacement du fichier de démarrage** puis reboot (le plus propre).
   - **Rejeu en CLI** : colle la config (méthode 3) — attention aux dépendances
     d'ordre (créer les VLAN avant de les affecter).
3. `save`, reboot, vérifie : `display current-configuration` vs fichier.
4. ✅ **Teste toujours la restauration sur un switch de maquette** avant d'en
   avoir besoin en urgence à 2h du matin.

⚠️ Restaurer une config d'un **autre modèle** (ex. 24P4S → 48P4S) : les numéros
de ports ne correspondent pas. Relis et adapte à la main, ne colle pas aveuglément.

---

## 90. Architecture des fichiers : vrpcfg, système, patches

```
dir flash:/
display startup
```

- **Fichier de configuration** : `vrpcfg.zip` (config de démarrage).
- **Fichier système** : l'image logicielle (`.cc` / nom selon version).
- `display startup` montre **quel** fichier système et **quelle** config seront
  utilisés au prochain boot — **vérifie toujours après une mise à jour**.

⚠️ Deux images système peuvent coexister en flash : c'est ce qui permet le
**rollback** (section 91). Ne supprime jamais l'ancienne image avant d'avoir
validé la nouvelle en production pendant au moins une semaine.

---

## 91. Mise à jour firmware : procédure complète (Smart Upgrade + manuel)

**Option A — Smart Upgrade via HOUP (recommandé si Internet OK) :**
La plateforme HOUP (Huawei Online Upgrade Platform) calcule le **chemin de mise
à jour standardisé**, pré-charge la nouvelle version, et propose une mise à jour
en un clic (valeur datasheet). Le Perpetual PoE maintient l'alimentation des
équipements pendant l'opération (section 59).

**Option B — Manuelle (classique, sans Internet) :**

1. **Sauvegarde** la config (section 88) — non négociable.
2. Télécharge l'image depuis le support Huawei (via ton revendeur / compte).
   **Vérifie la compatibilité modèle + version** (notes de release).
3. Transfère l'image sur le switch (FTP/SFTP/TFTP).
4. Déclare la nouvelle image comme image de démarrage :
   ```
   startup system-software <nom-image.cc>
   display startup          # VÉRIFIE : la bonne image est marquée "next startup"
   ```
   > La syntaxe exacte (`startup system-software`) est **à vérifier sur la
   > version logicielle du modèle exact**.
5. `save`, puis **reboot** en heure creuse : `reboot`.
6. Après reboot : `display version` → la nouvelle version est active.
7. **Batterie de tests** : ports up, VLAN, PoE, DHCP, un ping de chaque VLAN,
   accès web/SSH. Ne déclare pas « c'est bon » avant.
8. Garde l'ancienne image en flash pendant 1 à 2 semaines.

⚠️ **Fenêtre de maintenance :** même avec Perpetual PoE, le plan de commutation
est interrompu pendant le reboot. Jamais de mise à jour en pleine journée sans
validation écrite du responsable.

---

## 92. Rollback : revenir en arrière quand ça tourne mal

**Scénario :** la nouvelle version a un bug (fonction X qui ne marche plus).

1. En console/SSH : redéclare l'**ancienne image** comme image de démarrage
   (`startup system-software <ancienne-image>`).
2. `display startup` pour vérifier.
3. `reboot`.
4. Vérifie la version et refais la batterie de tests (section 91, étape 7).
5. Restaure la config sauvegardée **avant** la mise à jour si la nouvelle
   version l'a modifiée/migrée (ça arrive : compare les fichiers).

✅ **C'est pour ça qu'on garde l'ancienne image** (section 90). Un rollback sans
image de secours = réinstallation complète = heures d'arrêt.

---

## 93. Reset usine : procédure et précautions

**Quand :** switch à reconditionner, config irrécupérable, changement de site.

**Méthode CLI (à vérifier sur la version exacte) :**

```
reset saved-configuration
reboot
# Au reboot : [Y/N] → Y pour effacer
```

