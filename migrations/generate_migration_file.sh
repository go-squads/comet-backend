#!/bin/bash

timestamp() {
  date +%s
}

operation=$1

current_time=$(timestamp)

touch ${current_time}_$operation.up.sql
touch ${current_time}_$operation.down.sql

echo "Generated file up and down for ${current_time}_$operation"
