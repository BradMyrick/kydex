# KYDEX

KYDEX is a secure, cross-network decentralized exchange (DEX) that enables users to trade digital assets across multiple blockchain networks. Built with a focus on security and performance, KYDEX aims to provide a seamless and user-friendly trading experience.

## Features

- Cross-network trading
- Gas-efficient transactions
- Secure and transparent asset management
- Advanced cryptographic algorithms
- Decentralized order management

## Getting Started

### Prerequisites

- Go 1.16 or higher
- A supported database (e.g., PostgreSQL, MySQL)

### Installation

1. Clone the repository:
    
```bash
git clone https://github.com/BradMyrick/kydex.git
```

2. Change to the project directory:

```bash
cd kydex
```

3. Install dependencies:

```bash
go mod download
```

4. Configure the application:

Copy the sample configuration file and update it with your database credentials and other settings:

```bash
cp .env.sample .env
```

5. Build and run the application:

```bash
go build -o kydex ./cmd/kydex
./kydex
```