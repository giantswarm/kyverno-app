package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/spf13/cobra"

	"github.com/giantswarm/kyverno-app/tools/values-sync/internal/chart"
	"github.com/giantswarm/kyverno-app/tools/values-sync/internal/config"
	"github.com/giantswarm/kyverno-app/tools/values-sync/internal/schema"
	"github.com/giantswarm/kyverno-app/tools/values-sync/internal/values"
)

type options struct {
	chartDir   string
	configPath string
	dryRun     bool
	addNew     bool
	output     string
}

// report is the JSON-serialisable sync report.
type report struct {
	ChartDir string       `json:"chartDir"`
	Results  []syncResult `json:"results"`
	Schema   schemaResult `json:"schema"`
}

type syncResult struct {
	Subchart string   `json:"subchart"`
	Removed  []string `json:"removed"`
	New      []string `json:"new"`
}

type schemaResult struct {
	Updated bool   `json:"updated"`
	Path    string `json:"path"`
	Error   string `json:"error,omitempty"`
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	opts := &options{}

	cmd := &cobra.Command{
		Use:   "values-sync",
		Short: "Sync values.yaml and schema after upstream vendir update",
		RunE: func(cmd *cobra.Command, args []string) error {
			return execute(opts)
		},
	}

	cmd.Flags().StringVar(&opts.chartDir, "chart-dir", "", "Path to the parent Helm chart directory (defaults to first helm/*/ match)")
	cmd.Flags().StringVar(&opts.configPath, "config", "", "Path to values-sync.yaml config (auto-detected if not set)")
	cmd.Flags().BoolVar(&opts.dryRun, "dry-run", false, "Print what would change without modifying files")
	cmd.Flags().BoolVar(&opts.addNew, "add-new", false, "Auto-add new upstream keys to values.yaml with upstream default value")
	cmd.Flags().StringVar(&opts.output, "output", "text", "Output format: text or json")

	return cmd.Execute()
}

func execute(opts *options) error {
	// Resolve chart directory.
	chartDir := opts.chartDir
	if chartDir == "" {
		detected, err := detectChartDir()
		if err != nil {
			return fmt.Errorf("detecting chart directory: %w", err)
		}
		chartDir = detected
		fmt.Fprintf(os.Stderr, "Auto-detected chart directory: %s\n", chartDir)
	}

	// Discover subchart names from Chart.yaml.
	deps, err := chart.LoadDependencies(chartDir)
	if err != nil {
		return err
	}
	if len(deps) == 0 {
		return fmt.Errorf("no dependencies found in %s/Chart.yaml", chartDir)
	}

	// Load config.
	configPath := opts.configPath
	if configPath == "" {
		configPath = detectConfigPath()
	}
	cfg, err := config.Load(configPath)
	if err != nil {
		return fmt.Errorf("loading config: %w", err)
	}
	if configPath != "" && len(cfg.Exclude) > 0 {
		fmt.Fprintf(os.Stderr, "Loaded config from %s (%d exclusions)\n", configPath, len(cfg.Exclude))
	}

	// Load our values.yaml.
	valuesPath := filepath.Join(chartDir, "values.yaml")
	doc, err := values.LoadValuesDoc(valuesPath)
	if err != nil {
		return err
	}

	// Sync each subchart.
	rep := report{ChartDir: chartDir}
	for _, dep := range deps {
		upstreamPath := filepath.Join(chartDir, "charts", dep, "values.yaml")
		if _, err := os.Stat(upstreamPath); os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "Warning: upstream values not found at %s, skipping %s\n", upstreamPath, dep)
			continue
		}

		res, err := values.SyncSubchart(doc, dep, upstreamPath, values.SyncOptions{
			DryRun:  opts.dryRun,
			AddNew:  opts.addNew,
			Exclude: cfg.Exclude,
		})
		if err != nil {
			return fmt.Errorf("syncing subchart %s: %w", dep, err)
		}

		sort.Strings(res.Removed)
		sort.Strings(res.New)

		rep.Results = append(rep.Results, syncResult{
			Subchart: res.Subchart,
			Removed:  res.Removed,
			New:      res.New,
		})
	}

	// Write updated values.yaml (unless dry-run).
	if !opts.dryRun {
		if err := values.WriteValues(valuesPath, doc); err != nil {
			return fmt.Errorf("writing updated values.yaml: %w", err)
		}
	}

	// Regenerate schema (unless dry-run).
	schemaPath := filepath.Join(chartDir, "values.schema.json")
	rep.Schema.Path = schemaPath
	if !opts.dryRun {
		if err := schema.Regenerate(valuesPath, schemaPath); err != nil {
			rep.Schema.Error = err.Error()
		} else {
			rep.Schema.Updated = true
		}
	}

	// Print report.
	switch opts.output {
	case "json":
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		if err := enc.Encode(rep); err != nil {
			return fmt.Errorf("encoding JSON report: %w", err)
		}
	default:
		printTextReport(rep, opts)
	}

	return nil
}

func detectConfigPath() string {
	candidates := []string{
		"tools/values-sync/values-sync.yaml",   // from repo root
		"values-sync.yaml",                     // from tools/values-sync dir
		"../tools/values-sync/values-sync.yaml",
	}
	for _, c := range candidates {
		if _, err := os.Stat(c); err == nil {
			return c
		}
	}
	return ""
}

func detectChartDir() (string, error) {
	matches, err := filepath.Glob("helm/*/")
	if err != nil || len(matches) == 0 {
		matches, err = filepath.Glob("../helm/*/")
		if err != nil || len(matches) == 0 {
			return "", fmt.Errorf("no helm/*/ directory found; use --chart-dir")
		}
	}
	return matches[0], nil
}

func printTextReport(rep report, opts *options) {
	fmt.Printf("SYNC REPORT: %s\n", rep.ChartDir)
	for _, r := range rep.Results {
		if len(r.Removed) == 0 && len(r.New) == 0 {
			fmt.Printf("  [%s] no changes\n", r.Subchart)
			continue
		}
		fmt.Printf("  [%s]\n", r.Subchart)
		if len(r.Removed) > 0 {
			action := "Removed from values.yaml"
			if opts.dryRun {
				action = "Would remove from values.yaml"
			}
			fmt.Printf("    %s:\n", action)
			for _, p := range r.Removed {
				fmt.Printf("      - %s\n", p)
			}
		}
		if len(r.New) > 0 {
			action := "Added new upstream keys"
			if opts.dryRun {
				action = "Would add new upstream keys"
			}
			fmt.Printf("    %s:\n", action)
			for _, p := range r.New {
				fmt.Printf("      - %s\n", p)
			}
		}
	}
	if opts.dryRun {
		fmt.Println("  Schema: dry-run, not regenerated")
	} else if rep.Schema.Updated {
		fmt.Printf("  Schema: Updated %s\n", rep.Schema.Path)
	} else if rep.Schema.Error != "" {
		fmt.Printf("  Schema: ERROR: %s\n", rep.Schema.Error)
	}
}
