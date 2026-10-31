// Package cmd contains the entrypoint and subcommands for the application.
package cmd

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"

	"github.com/t-richards/vulnsearch/internal/nvd"
	"github.com/t-richards/vulnsearch/internal/paths"
	"github.com/t-richards/vulnsearch/internal/serve"
	"github.com/t-richards/vulnsearch/internal/store"
)

const usageHelp = `vulnsearch: A fast, offline-capable replacement for other CVE search tools.

Usage: vulnsearch (subcommand) [years...]

Subcommands:
  serve      Start the web interface. [DEFAULT]
  fetch      Download CVE feeds from NVD; all years by default.
  load       Load downloaded CVE feeds into the database; all downloaded years by default.
  optimize   Optimize the database after loading.
  clobber    Delete the database and all downloaded feeds.
`

// errUsage is returned for unknown subcommands or invalid arguments.
var errUsage = errors.New("invalid usage")

// Execute determines the subcommand to run (if any) and executes it.
func Execute() {
	log.SetFlags(0)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	err := run(ctx, os.Args[1:])

	stop()

	if errors.Is(err, errUsage) {
		fmt.Fprint(os.Stderr, usageHelp)
		os.Exit(2)
	}

	if err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context, args []string) error {
	if len(args) == 0 {
		return serveCmd(ctx)
	}

	switch args[0] {
	case "serve":
		return serveCmd(ctx)
	case "fetch":
		return fetchCmd(ctx, args[1:])
	case "load":
		return loadCmd(ctx, args[1:])
	case "optimize":
		return optimizeCmd(ctx)
	case "clobber":
		return clobberCmd()
	case "help", "-h", "-help", "--help":
		fmt.Print(usageHelp)

		return nil
	default:
		return errUsage
	}
}

func openStore(ctx context.Context) (*store.Store, error) {
	path, err := paths.DBPath()
	if err != nil {
		return nil, err
	}

	return store.Open(ctx, path)
}

func serveCmd(ctx context.Context) error {
	st, err := openStore(ctx)
	if err != nil {
		return err
	}
	defer st.Close()

	app, err := serve.NewApplication(st)
	if err != nil {
		return err
	}

	addr := ":3000"
	if port, ok := os.LookupEnv("PORT"); ok {
		addr = ":" + port
	}

	log.Printf("Vulnsearch listening on http://localhost%s/", addr)

	return app.ListenAndServe(ctx, addr)
}

// parseYears parses year arguments, defaulting to all years.
func parseYears(args []string) ([]int, error) {
	if len(args) == 0 {
		return nvd.Years(), nil
	}

	years := make([]int, 0, len(args))

	for _, arg := range args {
		year, err := strconv.Atoi(arg)
		if err != nil || year < nvd.EarliestYear {
			return nil, fmt.Errorf("%w: invalid year %q", errUsage, arg)
		}

		years = append(years, year)
	}

	return years, nil
}

func fetchCmd(ctx context.Context, args []string) error {
	years, err := parseYears(args)
	if err != nil {
		return err
	}

	dir, err := paths.DataDir()
	if err != nil {
		return err
	}

	return nvd.Fetch(ctx, dir, years)
}

func loadCmd(ctx context.Context, args []string) error {
	years, err := parseYears(args)
	if err != nil {
		return err
	}

	dir, err := paths.DataDir()
	if err != nil {
		return err
	}

	st, err := openStore(ctx)
	if err != nil {
		return err
	}
	defer st.Close()

	for _, year := range years {
		path := filepath.Join(dir, nvd.ArchiveName(year))

		_, err := os.Stat(path)
		if errors.Is(err, os.ErrNotExist) && len(args) == 0 {
			continue // Only load what has been fetched.
		}

		log.Printf("Loading: %d", year)

		count, err := st.Import(ctx, nvd.OpenArchive(path))
		if err != nil {
			return fmt.Errorf("while loading %d: %w", year, err)
		}

		log.Printf("Loaded %d CVEs", count)
	}

	log.Printf("Rebuilding product index")

	return st.RebuildProducts(ctx)
}

func optimizeCmd(ctx context.Context) error {
	st, err := openStore(ctx)
	if err != nil {
		return err
	}
	defer st.Close()

	log.Printf("Optimizing database...")

	return st.Optimize(ctx)
}

const clobberMsg = `
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@    DANGER: IRREVOCABLE DATA LOSS!     @
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@

This action will delete the following directory and its contents:

    %s

Do you wish to continue? [y/N]: `

func clobberCmd() error {
	dir, err := paths.CacheDir()
	if err != nil {
		return err
	}

	fmt.Printf(clobberMsg, dir)

	answer, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	if !strings.EqualFold(strings.TrimSpace(answer), "y") {
		fmt.Println("Cancelled.")

		return nil
	}

	err = os.RemoveAll(dir)
	if err != nil {
		return fmt.Errorf("while removing %s: %w", dir, err)
	}

	return nil
}
