# **DNS Zone File Creator** 🌐

_DNS Zone File Creator_ is an application written in Go that queries DNS records for a given domain and writes the results to a DNS zone file. This project aims to simplify the process of getting DNS records and converting them into a well-structured zone file.

## **🚀 Getting Started**

### **Prerequisites**

- Go v1.16 or later. Installation instructions can be found [here](https://golang.org/dl/).

### **Installation**

1. Clone the repository:
   ```bash
   git clone https://github.com/enesconf/dnszonefile.git
   cd dnszonefile
   ```
2. Run the program:
   ```bash
   go run main.go --domain yourdomain.com
   ```

## **🔍 What Does It Do?**

_DNS Zone File Creator_ generates a zone file named `yourdomain.com.zone`. This file contains DNS records for the domain and its subdomains.

## **🛠️ Customizing Subdomains**

By default, _DNS Zone File Creator_ queries a list of common subdomains (`www`, `mail`, `ftp`). If you want to customize this list, modify the `subdomains` slice in `main.go`.
