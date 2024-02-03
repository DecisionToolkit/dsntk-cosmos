#!/usr/bin/env bash

echo "Calculating SLA:"
curl -s \
     -d "{ YearsAsCustomer: 1, NumberOfUnits: 1000 }" \
     -H "Content-Type: application/json" \
     -X POST http://0.0.0.0:22022/evaluate/io/dsntk/DecisionContract/SLA

echo ""
echo "Calculating fine:"

curl -s \
     -d "{ YearsAsCustomer: 1, NumberOfUnits: 1000, DefectiveUnits: 0.034 }" \
     -H "Content-Type: application/json" \
     -X POST http://0.0.0.0:22022/evaluate/io/dsntk/DecisionContract/Fine

echo ""