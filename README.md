# Formatio: Simplifying Cloud DevOps for Developers ☁️

Formatio is a robust platform engineered to streamline the entire software deployment lifecycle, bringing the power of cloud-based virtual machines (Formachs) directly to developers. It integrates seamlessly with your existing Git workflows, offering automated CI/CD pipelines, intuitive network management, and comprehensive billing, all wrapped in a modern, user-friendly interface. Our goal is to eliminate DevOps complexities, allowing you to focus purely on building amazing applications.

## Table of Contents

*   [Installation](#installation)
*   [Usage](#usage)
*   [Features](#features)
*   [Technologies Used](#technologies-used)
*   [Author Info](#author-info)

## Installation

Getting Formatio up and running locally is straightforward. Follow these steps to set up both the backend API and the frontend application.

### ⚙️ Prerequisites

Before you start, ensure you have the following installed on your system:

*   [Go (1.22.x)](https://golang.org/dl/)
*   [Docker](https://www.docker.com/get-started)
*   [Docker Compose](https://docs.docker.com/compose/install/)
*   [Node.js (v18+)](https://nodejs.org/en/download/)
*   [Yarn](https://classic.yarnpkg.com/lang/en/docs/install/)

### 🚀 Clone the Repository

Start by cloning the project repository to your local machine:

```bash
git clone git@github.com:struckchure/formatio.git
cd formatio
```

### 📦 Backend Setup

1.  **Environment Variables**: Create a `.env` file by copying the sample provided:
    ```bash
    cp .env.sample .env
    ```
    Open `.env` and fill in the necessary values. Pay close attention to `PG_`, `APP_PORT`, `JWT_`, `GH_APP_`, `AUTH0_`, `RABBITMQ_URL`, `K8S_CLUSTER_CONFIG`, `INGRESS_ROOT_DOMAIN`, and `REDIS_URL`. For local development, many of these can use default values or be configured to point to your local Docker Compose services. `K8S_CLUSTER_CONFIG` can point to your local `.kube/config` file (already included in the repository).

2.  **Database Migrations**: Formatio uses Prisma for database management. Apply the migrations to set up your database schema:
    ```bash
    go run github.com/steebchen/prisma-client-go migrate deploy
    go run github.com/steebchen/prisma-client-go generate
    ```

3.  **Run Services**: Start the essential backend services (PostgreSQL, RabbitMQ, Redis) using Docker Compose:
    ```bash
    docker-compose up -d
    ```

4.  **Install Go Modules**: Ensure all Go dependencies are in place:
    ```bash
    go mod tidy
    ```

5.  **Start the Backend**: You can run the backend in development mode with hot-reloading using `air` (configured via `.air.toml`) or directly:
    ```bash
    # For hot-reloading (recommended for development)
    air

    # Or, for a standard run
    task dev
    ```

### 💻 Frontend Setup

1.  **Navigate to Client Directory**:
    ```bash
    cd client
    ```

2.  **Environment Variables**: Create a `.env` file for the frontend:
    ```bash
    cp .env.sample .env
    ```
    Update the `VITE_API_URL` and other relevant `VITE_` variables to match your backend setup.

3.  **Install Dependencies**:
    ```bash
    yarn
    ```

4.  **Generate API Clients**: The frontend uses `kubb` to generate API hooks and types from the backend's Swagger/OpenAPI specification.
    ```bash
    yarn kubb:generate
    ```

5.  **Start the Frontend**:
    ```bash
    yarn dev
    ```

Now, both your backend API and frontend application should be running locally!

## Usage

### 🌐 Accessing the Application

Once both the backend and frontend are running, you can access the Formatio web interface by navigating to `http://localhost:3000` in your web browser (or the port specified in `client/vite.config.ts`).

### 🔑 Authentication

*   **Local Registration/Login**: Use the `/register` and `/login` routes to create a new account or sign in with email and password.
*   **Social Login**: Formatio supports social authentication via Google and GitHub, powered by Auth0. Click the respective buttons on the login/registration page to connect your accounts.

### 🚀 Creating and Managing Formachs

1.  **Create a Formach**: Navigate to the "Formachs" section in the dashboard and click "Create Formach." You'll choose a name, a machine image (e.g., Alpine, Ubuntu), and a plan (CPU/memory allocation).
2.  **Monitor Status**: Once created, monitor your Formach's status (creating, running, shutting down) directly from the dashboard.
3.  **Configure Networking**: Under a Formach's "Networking" tab, you can set up port forwarding and generate public hostnames for your deployed applications.
4.  **Connect Repositories**: In the "Repository" tab, link your GitHub repositories to a Formach. This step is crucial for enabling continuous deployments.

### 🔄 Automated Deployments

1.  **GitHub Integration**: Ensure your GitHub account is authorized. If you haven't done so during social login, you can authorize it from the dashboard.
2.  **Define Actions**: Formatio looks for a `.formatio/action.yaml` file in your connected repository. This YAML file defines the build and deployment steps that Formatio will execute on your Formach.
3.  **Push to Deploy**: After connecting a repository and defining your action, every push to the configured branch (e.g., `main` or `master`) will automatically trigger a deployment to your Formach.
4.  **View Logs**: Track the progress and output of your deployments in real-time within the "Deployments" tab of your Formach's details page.

### 📈 API Documentation

The backend API is documented using Swagger. Once the backend is running, you can access the interactive API documentation at:

*   `http://localhost:<APP_PORT>/swagger/index.html` (replace `<APP_PORT>` with your configured backend port, typically `8000`).

This interface allows you to explore all available API endpoints, their expected parameters, and response structures.

## Features

*   **☁️ Cloud VM Provisioning (Formachs)**: Quickly provision and manage cloud-based virtual machines tailored to your application's needs.
*   **🔗 Integrated Git/GitHub CI/CD**: Seamlessly connect your GitHub repositories to automate build and deployment pipelines via webhooks.
*   **🚦 Automated Network Configuration**: Easily configure public access to your applications with automated ingress and port mapping.
*   **💸 Billing & Subscription Management**: Track usage and manage payment methods and invoices for your cloud resources.
*   **🔐 Flexible User Authentication**: Supports traditional email/password login and modern social authentication (Google, GitHub) via Auth0.
*   **🖥️ Real-time Deployment Logs**: Monitor your CI/CD pipeline execution with live streaming logs, providing immediate feedback on deployment status.
*   **🧩 Extensible Action Workflow**: Define custom build and deployment steps using intuitive YAML configurations (leveraging Formatio Storm).
*   **✨ Modern User Interface**: A clean, intuitive React-based frontend built with Vite, Tailwind CSS, and TanStack Query for an exceptional user experience.
*   **🧪 Robust Testing Framework**: Comprehensive unit and integration tests ensure reliability and maintainability of critical services.

## Technologies Used

| Category        | Technology          | Description                                         | Link                                                       |
| :-------------- | :------------------ | :-------------------------------------------------- | :--------------------------------------------------------- |
| **Backend**     | Go (1.22.x)         | Core programming language for high performance.     | [Go](https://golang.org/)                                  |
|                 | Fiber               | Express-inspired web framework for Go.              | [Fiber](https://gofiber.io/)                               |
|                 | Prisma Client Go    | Type-safe ORM for database interactions.            | [Prisma Client Go](https://www.prisma.io/client/go)        |
|                 | PostgreSQL          | Robust relational database.                         | [PostgreSQL](https://www.postgresql.org/)                  |
|                 | RabbitMQ            | Message broker for asynchronous tasks.              | [RabbitMQ](https://www.rabbitmq.com/)                      |
|                 | Redis               | In-memory data store for caching & session management. | [Redis](https://redis.io/)                                 |
|                 | Docker              | Containerization platform.                          | [Docker](https://www.docker.com/)                          |
|                 | Kubernetes client-go | Go client for Kubernetes API interaction.           | [client-go](https://github.com/kubernetes/client-go)       |
|                 | AWS SDK for Go      | Interacting with AWS services (e.g., Secrets Manager). | [AWS SDK for Go](https://aws.github.io/aws-sdk-go-v2/)     |
|                 | go-jwt/jwt          | JWT implementation for Go.                          | [go-jwt](https://github.com/golang-jwt/jwt)                |
|                 | Auth0 Go SDK        | Authentication and authorization provider.          | [Auth0](https://auth0.com/)                                |
|                 | Rodelar Go SDK      | Realtime events SDK for notifications.              | [Rodelar](https://github.com/Overal-X/rodelar-go-sdk)      |
|                 | Formatio Storm      | Internal DevOps orchestration library.              | [Formatio Storm](https://github.com/Overal-X/formatio.storm)|
|                 | Flutterwave SDK     | Payment processing.                                 | [Flutterwave](https://flutterwave.com/)                    |
| **Frontend**    | React               | JavaScript library for building user interfaces.    | [React](https://react.dev/)                                |
|                 | TypeScript          | Superset of JavaScript for type safety.             | [TypeScript](https://www.typescriptlang.org/)              |
|                 | Vite                | Fast frontend build tool.                           | [Vite](https://vitejs.dev/)                                |
|                 | TanStack Query      | Powerful asynchronous state management (React Query). | [TanStack Query](https://tanstack.com/query)               |
|                 | Zustand             | Small, fast, and scalable bear-necessities state-management. | [Zustand](https://zustand-bear.github.io/)                 |
|                 | Tailwind CSS        | Utility-first CSS framework.                        | [Tailwind CSS](https://tailwindcss.com/)                   |
|                 | Sass                | CSS pre-processor.                                  | [Sass](https://sass-lang.com/)                             |
|                 | Auth0 React SDK     | React integration for Auth0 authentication.         | [Auth0 React SDK](https://auth0.com/docs/libraries/auth0-react) |
|                 | Rodelar React SDK   | Realtime events SDK for notifications in React.     | [Rodelar](https://github.com/Overal-X/rodelar-ts-sdk)      |

## Author Info

👋 Hi there! I built this project to deepen my understanding of cloud infrastructure, distributed systems, and modern web development practices. If you have any questions or want to connect, feel free to reach out!

*   **Email**: [ameenmohammed2311@gmail.com](mailto:ameenmohammed2311@gmail.com)
*   **LinkedIn**: [Mohammed Al-Ameen](https://linkedin.com/in/struckchure)
*   **Twitter**: [@struckchure](https://twitter.com/struckchure)

---

[![Go Version](https://img.shields.io/badge/Go-1.22.x-00ADD8?style=flat-square&logo=go)](https://golang.org)
[![Docker](https://img.shields.io/badge/Docker-25.0.5-2496ED?style=flat-square&logo=docker)](https://www.docker.com/)
[![React](https://img.shields.io/badge/React-18.x-61DAFB?style=flat-square&logo=react)](https://react.dev/)
[![TypeScript](https://img.shields.io/badge/TypeScript-5.2-3178C6?style=flat-square&logo=typescript)](https://www.typescriptlang.org/)
[![Project Status](https://img.shields.io/badge/Status-Active-brightgreen?style=flat-square)](https://github.com/struckchure/formatio)

[![Readme was generated by Dokugen](https://img.shields.io/badge/Readme%20was%20generated%20by-Dokugen-brightgreen)](https://www.npmjs.com/package/dokugen)
