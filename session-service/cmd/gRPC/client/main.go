/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a client for Session service.
package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"os"
	"time"

	"google.golang.org/grpc"
	pb "github.com/develamit/go-grpc-gateway/session-service/api/proto/v1"
)

const (
	address     = "localhost:4101"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("did not connect: ", err)
	}
	defer conn.Close()
	c := pb.NewSessionClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SessionPost(ctx, &pb.SessionRequest{Name: name})
	if err != nil {
		log.Fatal("could not receive: ", err)
	}
	log.Info("Client received: ", r.GetMessage())
}
