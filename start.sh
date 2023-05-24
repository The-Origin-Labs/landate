#!/bin/bash
go run storage/main.go & P1=$!
go run api-gateway/main.go & P2=$!
wait $P1 $P2 