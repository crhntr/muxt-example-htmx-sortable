package main

import (
	"cmp"
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"

	"github.com/typelate/example-sortable/internal/database"
	"github.com/typelate/example-sortable/internal/domain"
	"github.com/typelate/example-sortable/internal/hypertext"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	db, err := database.Setup(ctx)
	if err != nil {
		panic(err)
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:       slog.LevelInfo,
		AddSource:   true,
		ReplaceAttr: nil,
	}))
	txMgr := database.NewTransactions(db)

	svc := domain.New(logger, txMgr)

	mux := http.NewServeMux()
	hypertext.Routes(mux, svc)
	srv := &http.Server{
		Handler: mux,
		Addr:    ":" + cmp.Or(os.Getenv("PORT"), "8080"),
	}

	go func() {
		<-ctx.Done()
		defer db.Close()
		if err = srv.Close(); err != nil {
			panic(err)
		}
	}()
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}
}
