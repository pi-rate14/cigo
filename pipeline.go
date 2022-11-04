package main

import "time"

const NUM_PIPELINE_STAGES = 4

type executer interface {
	execute() (string, error)
}

func buildPipeline(proj string) []executer {
	pipeline := make([]executer, NUM_PIPELINE_STAGES)

	pipeline[0] = newStep(
		"go build",
		"go",
		"Go Build: SUCCESS",
		proj,
		[]string{"build", ".", "errors"},
	)

	pipeline[1] = newStep(
		"go test",
		"go",
		"Go Test: SUCCESS",
		proj,
		[]string{"test", "-v"},
	)

	pipeline[2] = newExceptionStep(
		"go fmt",
		"gofmt",
		"Gofmt: SUCCESS",
		proj,
		[]string{"-l", "."},
	)

	pipeline[3] = newTimeoutStep(
		"git push",
		"git",
		"Git Push: SUCCESS",
		proj,
		[]string{"push", "origin", "master"},
		30*time.Second,
	)

	return pipeline
}
