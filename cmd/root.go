package cmd

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"v3board/initialization"
	"v3board/router"

	"github.com/spf13/cobra"
)

var (
	debug   = false
	workDir string
	rootCmd = &cobra.Command{
		Use:   "v3board",
		Short: "new panel for personal",
		Run: func(cmd *cobra.Command, args []string) {
			if workDir == "" {
				log.Fatalln("please set v3board work dir by -w=/path/to/dir")
			}
			initialization.InitLog(debug)

			// TODO 动态端口
			srv := &http.Server{
				Addr:    ":8080",
				Handler: router.Router(),
			}

			go func() {
				if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					log.Println("server closed: ", err)
				}
			}()

			// Wait for interrupt signal to gracefully shutdown the server with
			// a timeout of 5 seconds.
			quit := make(chan os.Signal, 1)
			// kill (no param) default send syscall.SIGTERM
			// kill -2 is syscall.SIGINT
			// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
			signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
			<-quit
			log.Println("Shutting down server...")

			// The context is used to inform the server it has 5 seconds to finish
			// the request it is currently handling
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			if err := srv.Shutdown(ctx); err != nil {
				log.Fatal("Server forced to shutdown:", err)
			}

			log.Println("Server exited")
		},
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.PersistentFlags().StringVarP(&workDir, "workDir", "w", "/usr/local", "-w=/path/to/dir define v3board work dir, storage data and logs")
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "-d=true or --debug=true to enable debug mode")
}

func Execute() error {
	return rootCmd.Execute()
}
