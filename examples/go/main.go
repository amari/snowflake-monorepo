package main

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"

	snowflakev1 "github.com/amari/snowflake-monorepo/snowflake-go/pkg/proto/snowflake/v1"
)

func main() {
	conn, err := grpc.NewClient("localhost:8081", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	cli := snowflakev1.NewSnowflakeServiceClient(conn)

	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	resp, err := cli.NextSnowflake(ctx, &snowflakev1.NextSnowflakeRequest{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Created a Snowflake: %s\n", resp.GetSnowflake().GetStringValue())
}
