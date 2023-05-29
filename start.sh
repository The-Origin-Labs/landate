#!/bin/bash
go run storage/main.go & P1=$!
go run api-gateway/main.go & P2=$!
go run document/main.go & P3=$!
wait $P1 $P2 $P3