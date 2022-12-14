package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"
)

func run(proj string, out io.Writer) error {
	if proj == "" {
		return fmt.Errorf("project directory is required : %w", ErrValidatoin)
	}

	pipeline := buildPipeline(proj)

	sig := make(chan os.Signal, 1)

	errCh := make(chan error)
	done := make(chan struct{})

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for _, step := range pipeline {
			msg, err := step.execute()
			if err != nil {
				errCh <- err
			}

			_, err = fmt.Fprintln(out, msg)
			if err != nil {
				errCh <- err
			}
		}
		close(done)
	}()

	for {
		select {
		case rec := <-sig:
			signal.Stop(sig)
			return fmt.Errorf("%s: exiting: %w", rec, ErrSignal)
		case err := <-errCh:
			return err
		case <-done:
			return nil
		}
	}
}

func main() {
	proj := flag.String("p", "", "Project Directory")
	flag.Parse()

	if err := run(*proj, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
