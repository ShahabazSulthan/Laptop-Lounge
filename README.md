

# Laptop Lounge

Laptop Lounge is an e-commerce website dedicated to laptops, featuring three main modules: User, Seller, and Admin. This platform facilitates a seamless shopping experience with a wide range of functionalities for each type of user.

## Table of Contents

1. [Features](#features)
2. [Installation](#installation)
3. [Usage](#usage)
4. [Modules](#modules)
5. [Testing & Mocking](#testing--mocking)
6. [CI/CD Pipeline](#cicd-pipeline)
7. [Docker & Kubernetes](#docker--kubernetes)
8. [Contributing](#contributing)
9. [License](#license)

## Features

### User
- **Signup and Login**: Includes Twilio OTP verification.
- **Forgot Password**: Reset password functionality.
- **Search Products**: Sort by A-Z, Z-A, High to Low, Low to High.
- **Filter Products**: Various filters available.
- **Cart Management**: Add and manage cart items.
- **Order Management**: Place, return, and cancel orders.
- **Invoice Generation**: Automated invoice creation.
- **Payment Options**: Razorpay, COD (Cash on Delivery), Wallet.
- **Wishlist**: Save products for future purchase.
- **Product Reviews**: Submit and view reviews.
- **Helpdesk**: Customer support.

### Seller
- **Signup and Login**: Requires admin verification using GST number.
- **Order Management**: Manage incoming orders.
- **Category Offers**: Apply offers to product categories.
- **Product Management**: Add, edit, and delete products.
- **Sales Report**: Generate sales reports in Excel and PDF formats.

### Admin
- **Login**: Secure admin access.
- **User & Seller Verification**: Verify, block, and unblock users and sellers.
- **Category & Brand Management**: Add, edit, and delete categories and brands.
- **Coupon Management**: Create and manage discount coupons.

## Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/ShahabazSulthan/Laptop-Lounge.git
   ```
2. Navigate to the project directory:
   ```sh
   cd laptop-lounge
   ```
3. Install dependencies:
   ```sh
   go mod download
   ```
4. Set up environment variables for database, Twilio, Razorpay, and other configurations.

5. Run the application:
   ```sh
   go run cmd/main.go
   ```

## Usage

### User Module
- Signup and login to the application.
- Browse and search for products.
- Add products to the cart and proceed to checkout.
- Choose payment methods (Razorpay, COD, Wallet).
- Track, return, and cancel orders.
- Generate invoices for orders.
- Utilize wishlist and product review features.
- Contact helpdesk for support.

### Seller Module
- Signup and login to the seller portal.
- Manage product listings (add, edit, delete).
- Handle orders and manage category-specific offers.
- Generate sales reports in both Excel and PDF formats.

### Admin Module
- Login to the admin dashboard.
- Verify, block, and unblock users and sellers.
- Manage product categories and brands.
- Create and manage discount coupons.

## Testing & Mocking

Laptop Lounge includes comprehensive testing and mocking to ensure reliability and maintainability:
- **Unit Tests**: Thorough tests for individual components.
- **Mocking**: Integrated with `gomock` to mock dependencies and isolate test scenarios.

To run tests:
```sh
go test ./...
```

## CI/CD Pipeline

The project is integrated with GitHub Actions for Continuous Integration and Continuous Deployment:
- **Automated Testing**: Ensures code quality with every commit.
- **Deployment**: The Docker image is automatically built and deployed to Kubernetes.

## Docker & Kubernetes

The project is fully containerized using Docker and deployed in a Kubernetes cluster:
- **Docker Image**: The Docker image can be found on [Docker Hub](https://hub.docker.com/repository/docker/shahabaz4573/laptop_lounge/general).
- **Kubernetes Deployment**: The application is deployed in a scalable Kubernetes environment.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any features, bug fixes, or enhancements.

