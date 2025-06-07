#!/bin/bash

echo "Generating SQLC code..."
sqlc -f ./config/sqlc.yaml generate
echo "SQLC code generation completed."