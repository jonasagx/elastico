package commands

import (
	"context"
	"errors"
	"fmt"

	"github.com/elastic/go-elasticsearch/v7"
)

var (
	commands map[string]Command

	ErrNoClient    = errors.New("no client provided")
	ErrWrongClient = errors.New("client time is not elasticsearch")
)

func init() {
	commands = make(map[string]Command)

	commands["ping"] = Ping
	commands["indices"] = Indices
}

func Get(cmdName string, args ...string) (Command, bool) {
	cmd, ok := commands[cmdName]
	return cmd, ok
}

type Command func(ctx context.Context, args ...string) error

func getClient(ctx context.Context) (*elasticsearch.Client, error) {
	rcli := ctx.Value("client")
	if rcli == nil {
		return nil, ErrNoClient
	}
	es, ok := rcli.(*elasticsearch.Client)
	if !ok {
		return nil, ErrWrongClient
	}

	return es, nil
}

func Ping(ctx context.Context, args ...string) error {
	es, err := getClient(ctx)
	if err != nil {
		return err
	}

	res, err := es.Info()
	if err != nil {
		return fmt.Errorf("failed to ping with: %w", err)
	}

	// info := map[string]interface{}{}
	// err = json.NewDecoder(res.Body).Decode(&info)
	// if err != nil {
	// 	return fmt.Errorf("failed processing ping with: %w", err)
	// }

	fmt.Printf("pinged cluster: %s\n", res)
	return nil
}

func Indices(ctx context.Context, args ...string) error {
	es, err := getClient(ctx)
	if err != nil {
		return err
	}

	res, err := es.Cat.Aliases()
	if err != nil {
		return fmt.Errorf("failed getting indices with: %w", err)
	}

	fmt.Printf("indices: %s\n", res)
	return nil
}
