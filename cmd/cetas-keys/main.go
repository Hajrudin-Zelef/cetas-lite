// Commande cetas-keys — gestion des clés API des providers pour CETAS Lite.
//
// Adaptée du setup.py du CETAS original, sans dépendance à l'interface web :
// les clés sont saisies en aveugle, VALIDÉES par un appel réel au provider,
// stockées chiffrées dans le coffre (même vault.enc que l'application),
// puis exportées vers un .env où elles n'apparaissent JAMAIS en clair
// (entrées scellées AES-256-GCM).
//
// Usage :
//
//	cetas-keys [--home DIR] [--env PATH] <commande>
//
// Commandes : init | add [provider] | list | test [provider] | delete <provider>
// | export | passwd | status | sync [--api URL]
package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"golang.org/x/term"

	"cetas-lite/internal/keysetup"
	"cetas-lite/internal/securevault"
)

const (
	exitOK    = 0
	exitErr   = 1
	exitUsage = 2
)

var (
	flagHome string
	flagEnv  string
)

func main() {
	os.Exit(run(os.Args[1:]))
}

func run(args []string) int {
	// Flags globaux.
	rest := args[:]
	for len(rest) > 0 && strings.HasPrefix(rest[0], "--") {
		switch {
		case rest[0] == "--home" && len(rest) > 1:
			flagHome = rest[1]
			rest = rest[2:]
		case strings.HasPrefix(rest[0], "--home="):
			flagHome = strings.TrimPrefix(rest[0], "--home=")
			rest = rest[1:]
		case rest[0] == "--env" && len(rest) > 1:
			flagEnv = rest[1]
			rest = rest[2:]
		case strings.HasPrefix(rest[0], "--env="):
			flagEnv = strings.TrimPrefix(rest[0], "--env=")
			rest = rest[1:]
		default:
			return usage(fmt.Sprintf("option inconnue : %s", rest[0]))
		}
	}
	home, err := resolveHome()
	if err != nil {
		return fail(err)
	}
	envPath := flagEnv
	if envPath == "" {
		envPath = filepath.Join(home, ".env")
	}
	ctx := &cli{home: home, st: keysetup.New(home), envPath: envPath, in: bufio.NewReader(os.Stdin)}

	if len(rest) == 0 {
		// Sans argument : assistant interactif façon setup.py.
		if !ctx.checkSetupAuth() {
			return exitErr
		}
		return ctx.runWizard()
	}
	cmd, cmdArgs := rest[0], rest[1:]

	// L'aide ne nécessite pas le mot de passe de protection.
	switch cmd {
	case "help", "-h", "--help":
		return usage("")
	}

	// Porte d'entrée anti-vol (équivalent du "setup password" du setup.py) :
	// définit une fois, demandée ensuite à chaque lancement. Seule une
	// empreinte scrypt est stockée, jamais le mot de passe en clair.
	if !ctx.checkSetupAuth() {
		return exitErr
	}

	switch cmd {
	case "init":
		return ctx.cmdInit()
	case "add":
		return ctx.cmdAdd(optArg(cmdArgs, 0))
	case "list":
		return ctx.cmdList()
	case "test":
		return ctx.cmdTest(optArg(cmdArgs, 0))
	case "delete":
		if optArg(cmdArgs, 0) == "" {
			return usage("delete requiert un provider : cetas-keys delete <provider>")
		}
		return ctx.cmdDelete(cmdArgs[0])
	case "export":
		return ctx.cmdExport()
	case "passwd":
		return ctx.cmdPasswd()
	case "status":
		return ctx.cmdStatus()
	case "sync":
		return ctx.cmdSync(apiFlag(cmdArgs))
	case "help", "-h", "--help":
		return usage("")
	default:
		return usage(fmt.Sprintf("commande inconnue : %s", cmd))
	}
}

// checkSetupAuth — porte d'entrée anti-vol : à la première utilisation on
// définit le mot de passe de protection de l'outil, ensuite il est demandé
// à chaque lancement (3 tentatives). Seule une empreinte scrypt salée est
// stockée dans <home>/setup.auth — jamais le mot de passe en clair.
func (c *cli) checkSetupAuth() bool {
	if !keysetup.AuthExists(c.home) {
		fmt.Println()
		frame("Authentification")
		fmt.Println()
		// Anti-contournement : si un coffre existe déjà, la (re)définition
		// du mot de passe de protection exige le mot de passe maître.
		// Supprimer setup.auth ne suffit donc pas à passer outre.
		if c.st.Exists() {
			fmt.Println(red("  Un coffre existe déjà : le mot de passe maître est requis"))
			fmt.Println(red("  pour (re)définir la protection de l'outil."))
			masterOk := false
			for attempt := 1; attempt <= 3; attempt++ {
				mpw := c.readSecret(white("  Mot de passe maître du coffre : "))
				if _, err := c.st.Load(mpw); err == nil {
					masterOk = true
					break
				}
				errLine(fmt.Sprintf("Mot de passe incorrect (%d tentative(s) restante(s)).", 3-attempt))
			}
			if !masterOk {
				errLine("Accès refusé.")
				return false
			}
		} else {
			infoLine("Ce mot de passe protège cetas-keys lui-même (anti-vol).")
		}
		infoLine("Stocké sous forme d'empreinte chiffrée, jamais en clair.")
		pw, ok := c.readPasswordTwice(white("  Mot de passe de protection"))
		if !ok {
			return false
		}
		if err := keysetup.SetAuthPassword(c.home, pw); err != nil {
			errLine(fmt.Sprintf("Erreur : %v", err))
			return false
		}
		okLine("Protection activée.")
		return true
	}
	frame("Authentification")
	fmt.Println()
	for attempt := 1; attempt <= 3; attempt++ {
		pw := c.readSecret(white("  Mot de passe de l'outil : "))
		if keysetup.VerifyAuthPassword(c.home, pw) {
			okLine("Accès autorisé.\n")
			return true
		}
		errLine(fmt.Sprintf("Mot de passe incorrect (%d tentative(s) restante(s)).\n", 3-attempt))
	}
	errLine("Accès refusé.\n")
	return false
}

type cli struct {
	home    string
	st      *keysetup.Store
	envPath string
	in      *bufio.Reader
}

// ---------------------------------------------------------------------------
// commandes
// ---------------------------------------------------------------------------

func (c *cli) cmdInit() int {
	if c.st.Exists() {
		fmt.Printf("Un coffre existe déjà : %s\n", c.st.Path())
		fmt.Println("Utilisez 'passwd' pour changer le mot de passe, 'add' pour ajouter des clés.")
		return exitErr
	}
	fmt.Println("── Création du coffre ──")
	if !keysetup.PepperSet() {
		fmt.Println("Note : CETAS_PEPPER n'est pas défini (recommandé en production).")
	}
	suggest, _ := securevault.GeneratePassword(20)
	fmt.Printf("Suggestion de mot de passe fort : %s\n\n", suggest)
	pw, ok := c.readPasswordTwice("Nouveau mot de passe maître")
	if !ok {
		return exitErr
	}
	if err := c.st.Init(pw); err != nil {
		return fail(err)
	}
	fmt.Printf("Coffre créé : %s\n", c.st.Path())
	fmt.Println("Chiffrement : AES-256-GCM | Scrypt (N=2^16) | pepper CETAS_PEPPER")
	return exitOK
}

func (c *cli) cmdAdd(only string) int {
	pw, ok := c.requireVault()
	if !ok {
		return exitErr
	}
	for {
		id := only
		if id == "" {
			id = c.chooseProvider(pw)
			if id == "" {
				return exitOK // 0 / retour
			}
		}
		p, found := keysetup.ByID(id)
		if !found {
			fmt.Printf("Provider inconnu : %s\n", id)
			return exitUsage
		}
		c.configureProvider(pw, p)
		if only != "" {
			return exitOK
		}
	}
}

func (c *cli) cmdList() int {
	pw, ok := c.requireVault()
	if !ok {
		return exitErr
	}
	data, err := c.st.Load(pw)
	if err != nil {
		return fail(err)
	}
	keys := keysetup.APIKeys(data)
	if len(keys) == 0 {
		fmt.Println("Aucune clé API enregistrée.")
		return exitOK
	}
	ids := sortedKeys(keys)
	fmt.Println("Clés API enregistrées :")
	for _, id := range ids {
		label := id
		if p, found := keysetup.ByID(id); found {
			label = p.Label
		}
		fmt.Printf("  %-15s %s\n", label, keysetup.Mask(keys[id]))
	}
	return exitOK
}

func (c *cli) cmdTest(only string) int {
	pw, ok := c.requireVault()
	if !ok {
		return exitErr
	}
	data, err := c.st.Load(pw)
	if err != nil {
		return fail(err)
	}
	keys := keysetup.APIKeys(data)
	if only != "" {
		p, found := keysetup.ByID(only)
		if !found {
			return usage(fmt.Sprintf("provider inconnu : %s", only))
		}
		k, has := keys[only]
		if !has {
			fmt.Printf("%s : aucune clé enregistrée.\n", p.Label)
			return exitErr
		}
		return c.testOne(p, k)
	}
	if len(keys) == 0 {
		fmt.Println("Aucune clé à tester.")
		return exitOK
	}
	failed := 0
	for _, id := range sortedKeys(keys) {
		p, found := keysetup.ByID(id)
		if !found {
			fmt.Printf("  %-15s ignoré (provider inconnu)\n", id)
			continue
		}
		if c.testOne(p, keys[id]) != exitOK {
			failed++
		}
	}
	if failed > 0 {
		fmt.Printf("%d clé(s) en échec.\n", failed)
		return exitErr
	}
	fmt.Println("Toutes les clés sont valides.")
	return exitOK
}

func (c *cli) cmdDelete(id string) int {
	p, found := keysetup.ByID(id)
	if !found {
		return usage(fmt.Sprintf("provider inconnu : %s", id))
	}
	pw, ok := c.requireVault()
	if !ok {
		return exitErr
	}
	if !c.confirm(fmt.Sprintf("Supprimer la clé %s ? [o/N] : ", p.Label)) {
		fmt.Println("Annulé.")
		return exitOK
	}
	deleted, err := c.st.DeleteAPIKey(pw, id)
	if err != nil {
		return fail(err)
	}
	if !deleted {
		fmt.Printf("Aucune clé %s enregistrée.\n", p.Label)
		return exitOK
	}
	if err := c.exportEnv(pw); err != nil {
		return fail(err)
	}
	fmt.Printf("Clé %s supprimée.\n", p.Label)
	return exitOK
}

func (c *cli) cmdExport() int {
	pw, ok := c.requireVault()
	if !ok {
		return exitErr
	}
	if err := c.exportEnv(pw); err != nil {
		return fail(err)
	}
	return exitOK
}

func (c *cli) cmdPasswd() int {
	pw, ok := c.requireVault()
	if !ok {
		return exitErr
	}
	if _, changed := c.changePassword(pw); !changed {
		return exitErr
	}
	return exitOK
}

func (c *cli) cmdStatus() int {
	fmt.Printf("Coffre : %s\n", c.st.Path())
	if !c.st.Exists() {
		fmt.Println("État  : inexistant (cetas-keys init pour le créer)")
	} else {
		fmt.Printf("État  : présent (%d octets)\n", c.st.Size())
	}
	fmt.Printf("Pepper : %s\n", yesNo(keysetup.PepperSet()))
	fmt.Printf(".env   : %s\n", c.envPath)
	if fi, err := os.Stat(c.envPath); err == nil {
		n := countEnvLines(c.envPath)
		fmt.Printf("État  : présent (%d octets, %d entrée(s) scellée(s), 0600=%v)\n",
			fi.Size(), n, fi.Mode().Perm() == 0o600)
	} else {
		fmt.Println("État  : absent")
	}
	return exitOK
}

// cmdSync pousse les clés validées du coffre vers l'application via son API,
// sans passer par l'interface web. Seuls les providers gérés par CETAS Lite
// (openrouter, deepseek, opencode, opencode-go) sont synchronisés.
func (c *cli) cmdSync(apiURL string) int {
	pw, ok := c.requireVault()
	if !ok {
		return exitErr
	}
	return c.syncToApp(pw, apiURL)
}

// syncToApp — cœur partagé entre la commande sync et l'assistant.
func (c *cli) syncToApp(pw, apiURL string) int {
	data, err := c.st.Load(pw)
	if err != nil {
		return fail(err)
	}
	keys := keysetup.APIKeys(data)
	var syncable []keysetup.Provider
	for _, id := range sortedKeys(keys) {
		if p, found := keysetup.ByID(id); found && p.AppID != "" {
			syncable = append(syncable, p)
		}
	}
	if len(syncable) == 0 {
		fmt.Println("Aucune clé synchronisable (openrouter, deepseek, opencode, opencode-go).")
		return exitOK
	}
	if !isLocalhost(apiURL) {
		fmt.Printf("Attention : l'API %s n'est pas locale — les clés y transiteraient en clair.\n", apiURL)
		if !c.confirm("Continuer quand même ? [o/N] : ") {
			fmt.Println("Annulé.")
			return exitOK
		}
	}
	fmt.Printf("── Synchronisation vers %s ──\n", apiURL)
	fmt.Print("Utilisateur de l'application : ")
	username := c.readLine()
	if username == "" {
		return fail(fmt.Errorf("utilisateur vide"))
	}
	appPw := c.readSecret("Mot de passe de l'application : ")
	defer wipe([]byte(appPw))

	token, err := apiLogin(apiURL, username, appPw)
	if err != nil {
		return fail(fmt.Errorf("connexion à l'application : %w", err))
	}
	failed := 0
	for _, p := range syncable {
		if err := apiPutProviderKey(apiURL, token, p.AppID, keys[p.ID]); err != nil {
			fmt.Printf("  %-15s échec : %v\n", p.Label, err)
			failed++
			continue
		}
		fmt.Printf("  %-15s synchronisé\n", p.Label)
	}
	if failed > 0 {
		fmt.Printf("%d synchronisation(s) en échec.\n", failed)
		return exitErr
	}
	fmt.Println("Clés synchronisées avec l'application (prises en compte immédiatement).")
	return exitOK
}

// ---------------------------------------------------------------------------
// flux interactifs
// ---------------------------------------------------------------------------

// chooseProvider affiche le menu des providers avec leur état, retourne "" si
// l'utilisateur choisit 0 (terminer).
func (c *cli) chooseProvider(pw string) string {
	data, err := c.st.Load(pw)
	if err != nil {
		errLine(fmt.Sprintf("Erreur : %v", err))
		return ""
	}
	keys := keysetup.APIKeys(data)
	fmt.Println()
	fmt.Println(cyanS("  ────────────────────────────────────────────────────────"))
	fmt.Println(white("  Providers disponibles :"))
	fmt.Println()
	for i, p := range keysetup.Providers {
		var status string
		if k, has := keys[p.ID]; has && k != "" {
			status = green("✓ " + keysetup.Mask(k))
		} else {
			status = gray("─ Non configuré")
		}
		fmt.Println(white(fmt.Sprintf("  %d. %-15s", i+1, p.Label)) + "  " + status)
	}
	fmt.Println()
	fmt.Println(yellow("  0. Retour"))
	fmt.Println()
	fmt.Print(white("  Choisir un provider à configurer (ou 0) : "))
	choice := c.readLine()
	if choice == "0" || choice == "" {
		return ""
	}
	for i, p := range keysetup.Providers {
		if fmt.Sprint(i+1) == choice {
			return p.ID
		}
	}
	errLine("Choix invalide.\n")
	return c.chooseProvider(pw)
}

// configureProvider — saisie aveugle + validation live + enregistrement.
func (c *cli) configureProvider(pw string, p keysetup.Provider) {
	data, err := c.st.Load(pw)
	if err != nil {
		errLine(fmt.Sprintf("Erreur : %v", err))
		return
	}
	existing := keysetup.APIKeys(data)[p.ID]
	fmt.Println()
	fmt.Println(cyanS(fmt.Sprintf("  ── Configuration de %s ──", p.Label)))
	if existing != "" {
		fmt.Println(gray(fmt.Sprintf("  Clé actuelle : %s", keysetup.Mask(existing))))
		if !c.confirm(yellow("  Modifier cette clé ? [O/n] : ")) {
			fmt.Println(gray(fmt.Sprintf("  %s conservé.\n", p.Label)))
			return
		}
	}
	for {
		// Trim explicite : l'interface web trimme les clés API avant stockage
		// (internal/web/providers.go) — le copier-coller ajoute souvent des
		// espaces parasites. Le mot de passe maître, lui, n'est jamais trimmé.
		key := strings.TrimSpace(c.readSecret(white(fmt.Sprintf("  Clé %s (%s) : ", p.Label, p.Hint))))
		defer wipe([]byte(key))
		if key == "" {
			fmt.Println(gray(fmt.Sprintf("  %s ignoré.\n", p.Label)))
			return
		}
		fmt.Printf("  Test de la clé %s ... ", p.Label)
		ok, reason := keysetup.TestKey(p, key)
		if ok {
			fmt.Println(green("OK"))
		} else {
			fmt.Println(red(fmt.Sprintf("Échec : %s", reason)))
			if !c.confirm(yellow("  Réessayer ? [O/n] : ")) {
				fmt.Println(gray(fmt.Sprintf("  %s ignoré.\n", p.Label)))
				return
			}
			continue
		}
		if err := c.st.SetAPIKey(pw, p.ID, key); err != nil {
			errLine(fmt.Sprintf("Erreur d'enregistrement : %v", err))
			return
		}
		if err := c.exportEnv(pw); err != nil {
			errLine(fmt.Sprintf("Clé enregistrée, mais export .env impossible : %v", err))
			return
		}
		okLine(fmt.Sprintf("Clé %s enregistrée et .env mis à jour.\n", p.Label))
		return
	}
}

func (c *cli) testOne(p keysetup.Provider, key string) int {
	fmt.Printf("  Test de la clé %s ... ", p.Label)
	ok, reason := keysetup.TestKey(p, key)
	if ok {
		fmt.Println(green("OK"))
		return exitOK
	}
	fmt.Println(red(fmt.Sprintf("Échec : %s", reason)))
	return exitErr
}

// exportEnv régénère le .env scellé depuis le coffre.
func (c *cli) exportEnv(pw string) error {
	data, err := c.st.Load(pw)
	if err != nil {
		return err
	}
	proxyKey, err := c.st.ProxyKey(pw)
	if err != nil {
		return err
	}
	n, err := keysetup.WriteEnvFile(c.envPath, keysetup.APIKeys(data), proxyKey)
	if err != nil {
		return err
	}
	fmt.Printf(".env écrit : %d entrée(s) scellée(s) (aucune clé en clair).\n", n)
	return nil
}

// requireVault vérifie l'existence du coffre et demande le mot de passe maître.
func (c *cli) requireVault() (string, bool) {
	if !c.st.Exists() {
		fmt.Printf("Aucun coffre : %s\n", c.st.Path())
		fmt.Println("Créez-le d'abord avec : cetas-keys init")
		return "", false
	}
	pw := c.readSecret("Mot de passe maître : ")
	if _, err := c.st.Load(pw); err != nil {
		fmt.Printf("Erreur : %v\n", err)
		return "", false
	}
	return pw, true
}

// ---------------------------------------------------------------------------
// entrées / sorties
// ---------------------------------------------------------------------------

func (c *cli) readLine() string {
	line, _ := c.in.ReadString('\n')
	return strings.TrimSpace(line)
}

// readSecret lit une saisie sans écho (jamais affichée, jamais loggée).
// Renvoie la valeur BRUTE, sans trim : comme l'interface web, qui ne trimme
// pas le mot de passe maître (le trim des clés API est appliqué
// explicitement aux sites de saisie, cf. configureProvider).
func (c *cli) readSecret(prompt string) string {
	fmt.Print(prompt)
	raw, err := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println()
	if err != nil {
		return ""
	}
	return string(raw)
}

func (c *cli) readPasswordTwice(prompt string) (string, bool) {
	for {
		p1 := c.readSecret(prompt + " : ")
		defer wipe([]byte(p1))
		if p1 == "" {
			fmt.Println("Mot de passe vide.")
			return "", false
		}
		p2 := c.readSecret("Confirmer          : ")
		defer wipe([]byte(p2))
		if p1 != p2 {
			fmt.Println("Les mots de passe ne correspondent pas.")
			continue
		}
		if ok, msg := securevault.CheckPassword(p1); !ok {
			fmt.Printf("Mot de passe refusé : %s\n", msg)
			continue
		}
		return p1, true
	}
}

// confirm — [O/n] par défaut oui, [o/N] par défaut non selon le libellé.
func (c *cli) confirm(prompt string) bool {
	fmt.Print(prompt)
	ans := strings.ToLower(c.readLine())
	if strings.Contains(prompt, "[O/n]") {
		return ans != "n"
	}
	return ans == "o"
}

// wipe — effacement best-effort d'un secret en mémoire.
func wipe(b []byte) {
	for i := range b {
		b[i] = 0
	}
}

func yesNo(b bool) string {
	if b {
		return "oui"
	}
	return "non"
}

func sortedKeys(m map[string]string) []string {
	ids := make([]string, 0, len(m))
	for id := range m {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	return ids
}

func countEnvLines(path string) int {
	raw, err := os.ReadFile(path)
	if err != nil {
		return 0
	}
	n := 0
	for _, l := range strings.Split(string(raw), "\n") {
		if _, _, ok := keysetup.ParseEnvLine(l); ok {
			n++
		}
	}
	return n
}

func optArg(args []string, i int) string {
	if i < len(args) {
		return args[i]
	}
	return ""
}

func apiFlag(args []string) string {
	for i, a := range args {
		if a == "--api" && i+1 < len(args) {
			return args[i+1]
		}
		if strings.HasPrefix(a, "--api=") {
			return strings.TrimPrefix(a, "--api=")
		}
	}
	if v := strings.TrimSpace(os.Getenv("CETAS_LITE_ADDR")); v != "" {
		if !strings.Contains(v, "://") {
			v = "http://" + v
		}
		return v
	}
	return "http://127.0.0.1:8787"
}

func isLocalhost(url string) bool {
	u := strings.ToLower(url)
	return strings.Contains(u, "127.0.0.1") || strings.Contains(u, "localhost") || strings.Contains(u, "[::1]")
}

func resolveHome() (string, error) {
	if flagHome != "" {
		abs, err := filepath.Abs(flagHome)
		if err != nil {
			return "", err
		}
		return abs, nil
	}
	return keysetup.ResolveHome()
}

// ---------------------------------------------------------------------------
// API de l'application (sync)
// ---------------------------------------------------------------------------

var apiClient = &http.Client{Timeout: 15 * time.Second}

func apiLogin(apiURL, username, password string) (string, error) {
	body, _ := json.Marshal(map[string]string{"username": username, "password": password})
	req, err := http.NewRequest("POST", strings.TrimSuffix(apiURL, "/")+"/api/auth/login", bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := apiClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	raw, _ := io.ReadAll(io.LimitReader(resp.Body, 8192))
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("HTTP %d — %s", resp.StatusCode, strings.TrimSpace(string(raw)))
	}
	var out struct {
		Token string `json:"token"`
	}
	if err := json.Unmarshal(raw, &out); err != nil || out.Token == "" {
		return "", fmt.Errorf("réponse de connexion illisible")
	}
	return out.Token, nil
}

func apiPutProviderKey(apiURL, token, providerID, key string) error {
	body, _ := json.Marshal(map[string]string{"key": key})
	req, err := http.NewRequest("PUT", strings.TrimSuffix(apiURL, "/")+"/api/providers/"+providerID, bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := apiClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	raw, _ := io.ReadAll(io.LimitReader(resp.Body, 4096))
	if resp.StatusCode != 200 {
		return fmt.Errorf("HTTP %d — %s", resp.StatusCode, strings.TrimSpace(string(raw)))
	}
	return nil
}

// ---------------------------------------------------------------------------

func fail(err error) int {
	fmt.Fprintf(os.Stderr, "Erreur : %v\n", err)
	return exitErr
}

func usage(msg string) int {
	if msg != "" {
		fmt.Fprintf(os.Stderr, "Erreur : %s\n\n", msg)
	}
	fmt.Fprintln(os.Stderr, `cetas-keys — gestion des clés API (coffre chiffré + .env scellé)

Usage : cetas-keys [--home DIR] [--env PATH] [commande]

Sans commande : assistant interactif (comme setup.py) — il guide tout seul.

Commandes (pour scripts) :
  init              créer le coffre (mot de passe maître)
  add [provider]    ajouter des clés (saisie aveugle + validation live)
  list              lister les clés (masquées)
  test [provider]   re-valider une ou toutes les clés
  delete <provider> supprimer une clé
  export            régénérer le .env scellé depuis le coffre
  passwd            changer le mot de passe maître
  status            état du coffre et du .env
  sync [--api URL]  pousser les clés vers l'app (sans passer par l'UI)

Providers : nvidia, groq, openrouter, deepseek, freellmapi, anthropic, openai,
            grok, perplexity, google, mistral, qwen, kimi, glm,
            opencode, opencode-go

Le .env ne contient jamais de clé en clair : chaque entrée est scellée
(AES-256-GCM) avec la proxy_key du coffre — format <id>_key=<iv>:<ct>.`)
	if msg != "" {
		return exitUsage
	}
	return exitOK
}
