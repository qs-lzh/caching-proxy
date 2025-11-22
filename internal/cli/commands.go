package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/qs-lzh/caching-proxy/internal/proxy"
)

var (
	port       int
	origin     string
	clearCache bool
)

func init() {
	rootCmd.Flags().IntVar(&port, "port", 3000, "Port to run the caching proxy server")
	rootCmd.Flags().StringVar(&origin, "origin", "", "Origin server URL")
	rootCmd.Flags().BoolVar(&clearCache, "clear-cache", false, "Clear cached responses")
}

var rootCmd = &cobra.Command{
	Use:   "caching-proxy",
	Short: "A caching proxy server",
	Long: `A caching proxy that forwards requests to an origin server and caches responses.
Use --port and --origin to run the proxy server.
Use --clear-cache to clear stored cached responses.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if clearCache {
			fmt.Println("Cache cleared.")
			return nil
		}
		if origin == "" {
			return fmt.Errorf("--origin is required unless --clear-cache is used")
		}
		fmt.Printf("Starting proxy on port %d, origin=%s\n", port, origin)

		if err := proxy.StartProxyServer(proxy.ProxyConfig{Port: port, Origin: origin}); err != nil {
			return fmt.Errorf("failed to start proxy server: %w\n", err)
		}

		return nil
	},
}

func CommandExecute() error {
	return rootCmd.Execute()
}
