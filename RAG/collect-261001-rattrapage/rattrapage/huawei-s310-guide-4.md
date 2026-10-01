---
id: collect-261001-rattrapage/rattrapage/huawei-s310-guide-4
title: "Guide ultra-complet — Huawei eKit S310"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["incident"]
source: docs/RAG/collect-261001-rattrapage/huawei_s310_guide.md
source_anchor: ""
source_lines: [458, 676]
sha256: a15a87cbd98e73417492bd7162306f2585baec5a9fdb8b33f76b9eb8ac9f06a8
---

# Guide ultra-complet — Huawei eKit S310

1. Branche la console, ouvre le terminal à 9 600 8N1.
2. Allume le switch, attends la fin du boot.
3. Appuie sur **Entrée** : le switch affiche une invite de login (ou directement
   l'invite `<>` selon la version).
4. **Mot de passe initial :** sur les versions récentes, le switch impose de
   **créer un mot de passe admin au premier démarrage** (politique de sécurité
   Huawei). L'identifiant et la procédure exacte sont **à vérifier sur le guide
   de démarrage rapide du modèle exact** (étiquette / documentation fournie).
5. ⚠️ Choisis immédiatement un mot de passe **fort et unique par site**
   (12+ caractères, documenté dans le coffre de l'entreprise — voir section 117).
   Un switch avec mot de passe par défaut sur le réseau = porte ouverte.

**Les 3 niveaux d'invite VRP (à connaître par cœur) :**

```
<HUAWEI>            ← vue utilisateur : commandes d'affichage (display ...)
[HUAWEI]            ← vue système (après "system-view") : configuration
[HUAWEI-GigabitEthernet0/0/1]  ← vue d'interface : config d'un port
```

🔧 Commandes de survie immédiates :

```
display version              # version logicielle et matérielle
display current-configuration # config active
display interface brief      # état de tous les ports
save                         # SAUVEGARDER la config (sinon perdue au reboot !)
```

---

## 20. Configuration de l'IP de management (VLANIF)

Sans IP de management, pas de web, pas de SSH, pas de SNMP. La méthode Huawei :
une interface virtuelle **VLANIF** rattachée au VLAN de management.

**Exemple : VLAN 99 = management, réseau 192.168.99.0/24, switch en .10 :**

```
system-view
vlan 99
 description VLAN_MANAGEMENT
quit
interface Vlanif99
 ip address 192.168.99.10 255.255.255.0
quit
ip route-static 0.0.0.0 0.0.0.0 192.168.99.1   # passerelle par défaut
quit
save
```

Puis, sur un port d'accès au VLAN 99 (ex. port 24 réservé à l'admin) :

```
system-view
interface GigabitEthernet0/0/24
 port link-type access
 port default vlan 99
 description PORT_ADMIN
quit
save
```

[web] Équivalent : *Configuration → VLAN → créer VLAN 99*, puis *Interface VLANIF*,
puis *Routage → route statique*.

✅ **Bonnes pratiques :**
- VLAN de management **dédié** (pas le VLAN 1, pas le VLAN utilisateurs).
- Documente l'IP de chaque switch dans un **tableau d'adressage** (section 119).
- La passerelle par défaut est indispensable pour le cloud eKit et le syslog distant.

⚠️ **Le VLAN 1** est le VLAN natif par défaut sur tous les ports : ne jamais y mettre
le management en production (tout équipement branché y aurait accès).

---

## 21. Accès web (HTTP/HTTPS) : activation et premier login

```
system-view
http server enable            # active le serveur web
http secure-server enable     # active HTTPS (recommandé)
quit
save
```

Puis dans un navigateur : `https://192.168.99.10` (accepte le certificat
auto-signé — normal sur un équipement neuf).

**Premier login :** identifiant/mot de passe = ceux définis à la première connexion
(section 19). La procédure exacte et les identifiants initiaux sont
**à vérifier sur la documentation du modèle exact**.

[web] Une fois connecté, l'assistant de configuration initiale guide pour :
hostname, IP de management, mot de passe admin, date/heure.

⚠️ **HTTPS partout :** désactive HTTP pur dès que HTTPS fonctionne
(`undo http server enable`). Un mot de passe admin en clair sur le réseau,
même en interne, c'est non (voir durcissement, section 117).

---

## 22. L'app Huawei eKit : onboarding cloud pas à pas

Le mode cloud permet de **déployer sans toucher au CLI** : le switch récupère sa
configuration depuis le cloud après simple branchement (plug-and-play).

**Prérequis :**
- Compte sur la plateforme **Huawei eKit** (créé par le revendeur ou l'admin).
- Le switch a accès à Internet (passerelle par défaut configurée, section 20,
  + DNS si besoin).
- Numéro de série du switch (photo de l'étiquette, section 13).

**Procédure type :**

1. Dans le portail eKit (ou l'app mobile), crée le **site** (ex. « Agence Nord »).
2. Ajoute le switch au site via son **numéro de série** (ou scan du QR code
   sur l'emballage/l'étiquette si disponible).
3. Branche le switch au réseau avec accès Internet : il s'enregistre
   automatiquement (« deployment through the registration query center »).
4. Applique un **template de configuration** (VLAN, SSID si AP eKit associés)
   depuis le portail.
5. Vérifie dans le portail : switch « en ligne », version logicielle, alertes.

✅ **Intérêt majeur multi-sites :** un technicien branche, l'admin configure à
distance. Zéro déplacement pour un remplacement standard.

⚠️ **Dépendance Internet :** sans accès Internet, le switch ne peut ni
s'enregistrer ni recevoir de template. Toujours prévoir le **plan B local**
(console + config de base sur clé USB / documentée) pour les sites sans
connectivité fiable.

---

## 23. Cloud vs local : bascule et coexistence

- Les switches S310 supportent **les deux modes** et la bascule de l'un à l'autre
  selon le besoin (valeur datasheet).
- En mode cloud, la configuration poussée par le portail **prime** ; une
  modification locale en CLI peut être écrasée à la prochaine synchronisation.
- 🔧 **Règle d'équipe :** on choisit **un** mode de référence par site et on s'y
  tient. Le mélange « un coup cloud, un coup CLI » est la première source de
  configurations fantômes.
- Pour un dépannage urgent en CLI sur un switch cloud-managed : fais-le, mais
  **reporte la modification dans le template cloud** juste après, sinon elle
  disparaîtra.

---

## 24. Date, heure et NTP : le détail qui change tout

Des logs sans heure fiable sont inutilisables pour le dépannage (et pour la
corrélation avec la vidéosurveillance ou les tickets).

```
system-view
clock timezone Paris add 1        # fuseau (heure d'hiver ; +2 en été si pas de DST auto)
ntp-service unicast-server 192.168.99.5   # serveur NTP interne (ou  pool public si autorisé)
quit
save
```

🔧 Vérifications :

```
display clock
display ntp-service status
```

⚠️ Si ton pare-feu bloque le NTP (UDP 123) vers Internet, utilise **ton propre
serveau NTP interne** (souvent le contrôleur de domaine ou le routeur). Un switch
à l'heure, c'est aussi des certificats HTTPS qui fonctionnent et des logs
exploitables.

---

## 25. Nommer le switch : hostname et descriptions

```
system-view
sysname SW-ACC-01-AGENCE-NORD
quit
save
```

**Convention de nommage suggérée :** `SW-<rôle>-<num>-<site>`
(`SW-ACC-01-NORD`, `SW-CORE-01-SIEGE`). Et surtout, **décris chaque port** :

```
interface GigabitEthernet0/0/7
 description AP_ETAGE1_COULOIR
```

✅ Un `display interface brief` avec des descriptions, c'est 30 minutes de
dépannage économisées à chaque incident. Rends la description obligatoire dans
tes procédures d'ajout d'équipement.

---

## 26. Checklist de première configuration (à cocher sur site)

- [ ] Console OK, mot de passe admin fort défini et **documenté au coffre**
- [ ] Hostname selon la convention (`SW-...`)
- [ ] VLAN management créé, VLANIF avec IP, passerelle par défaut
- [ ] HTTPS activé, HTTP désactivé (ou planifié)
- [ ] SSH activé (section 117), Telnet désactivé
- [ ] NTP configuré, `display clock` correct
- [ ] Enregistrement eKit cloud (si mode cloud choisi)
- [ ] Ports décrits au fur et à mesure du câblage
- [ ] `save` effectué (vérifié par `display saved-configuration` vs `display current-configuration`)
- [ ] Config sauvegardée **hors du switch** (fichier texte, section 99)
- [ ] Étiquette sur le switch : hostname + IP de management (au marqueur sur
      scotch d'électricien, ou étiqueteuse — le post-it ne tient pas 3 ans)

🔧 **Le test final :** débranche/rebranche électriquement le switch. Tout doit
revenir (IP, config, PoE). Si quelque chose manque, c'est qu'un `save` a été oublié.

---

