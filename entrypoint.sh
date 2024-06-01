#!/bin/sh
ipfs init
# Set IPFS API HTTP Headers
ipfs config --json API.HTTPHeaders.Access-Control-Allow-Origin '["webui://-", "http://localhost:3000", "http://127.0.0.1:5001", "http://0.0.0.0:5001", "http://172.18.0.3:5001", "https://webui.ipfs.io"]'
ipfs config --json API.HTTPHeaders.Access-Control-Allow-Methods '["PUT", "POST"]'

# Start the IPFS daemon
exec ipfs daemon
