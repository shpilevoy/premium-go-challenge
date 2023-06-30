package main

import (
	"context"
)

type TaskInput string

type TaskResult map[string]interface{}


type Kafka interface {
	Produce(ctx context.Context, key TaskInput) error
	Consume() (context.Context, TaskInput, TaskResult, error)
}


func main() {

}
