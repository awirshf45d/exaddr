# Go Data Extractor

A lightweight CLI tool written in Go designed to extract structured data (IP addresses and Subdomains) from unstructured, "garbage" input files.

Stop manually grepping through messy log files, raw HTML, or mixed-content text buffers. Just dump the garbage into a file and let this tool extract exactly what you need.

## Use Case Scenario

**The Problem:**
You are doing reconnaissance on a target. You visit [crt.sh](https://crt.sh) to look for SSL certificates. You want the list of subdomains, but the page is full of HTML tags, table borders, and metadata. Copying the text gives you a mess.

**The Solution:**
1. Select All (Ctrl+A) and Copy (Ctrl+C) the raw text/HTML from the browser.
2. Paste it into a file (e.g., `raw_data.txt`).
3. Run this tool:
   ```bash
   exaddr -file raw_data.txt -d example.com
   ```
4. **Result:** A clean, sorted list of unique subdomains (e.g., `admin.example.com`, `vpn.example.com`) printed directly to your terminal.

## Features

*   **Subdomain Extraction:** extract subdomains belonging to specific root domains (supports comma-separated lists).
*   **IP Address Extraction:** regex-based extraction of IPv4 addresses.
*   **Garbage Filtering:** ignores HTML tags, random text, and special characters.
*   **Deduplication and Sorting**
*   **Flexible Output:** print results to the CLI (Standard Output) or save them directly to a file.

## Installation
```bash
go install github.com/awirshf45d/exAddr@latest
```

## Usage
```
exaddr -file raw_input.txt -d example.com -o cleaned-example.com-subs.txt
```
