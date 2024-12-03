#!/bin/bash
# Navigate to the project directory
cd /Users/arga.samosir/Documents/gateway_project/testing_tools

# Run tests and output results to a timestamped XML file
go test -v gateway_test.go > automation_result/results_$(date +%Y%m%d%H%M%S).xml
