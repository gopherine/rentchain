#!/bin/bash

# Set up key variables
owner="cosmos1x7wjvlwz74fmrpxtvpc3cm05ddrv8yxujme6w9"
renter="cosmos13j6u9w2sy9fdnndcxzjqfrw6tzp4ryzqeg8kyv"
itemId="125"
price="10token"
startTime=$(date +%s)
duration="3600" # 1 hour in seconds
chainId="rentchain"

output=$(rentchaind query auth account cosmos1x7wjvlwz74fmrpxtvpc3cm05ddrv8yxujme6w9 --chain-id rentchain)

sequence=$(echo "$output" | awk '/sequence:/ {gsub(/"/, "", $2); print $2}')
sequence=$(echo "$sequence" | tr -d '[:space:]') # Remove any whitespace
sequence=$(echo "$sequence" | awk '{print int($0)}') # Parse as integer

# Creating rental
echo "Creating rental..."
createOutput=$(echo "y" | rentchaind tx rentchain create-rental \
  $itemId \
  $owner \
  $renter \
  $price \
  $startTime \
  $duration \
  --from $owner \
  --chain-id $chainId \
  --trace \
  --yes)
echo "Create Rental Output:"
echo "$createOutput"