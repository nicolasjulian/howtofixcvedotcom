# HowToFixCVE.com Monorepo

Welcome to the monorepo for HowToFixCVE.com, a platform designed for searching and addressing Common Vulnerabilities and Exposures (CVEs).

## Quick Links

- **Public API:** [https://api.howtofixcve.com](https://api.howtofixcve.com)
- **Frontend:** [https://howtofixcve.com](https://howtofixcve.com)
- **ArgoCD Manifests:** [https://github.com/nicolasjulian/argo-howtofixcvedotcom](https://github.com/nicolasjulian/argo-howtofixcvedotcom)

## API Endpoints

- **Latest CVEs:** `/latest`
- **Specific CVE Details:** `/info/{CVE-ID}`

Both endpoints return responses in JSON format.

## Tech Stack

- **Kubernetes:** For container orchestration.
- **ArgoCD:** Deployments are managed through ArgoCD, with manifests housed in a [separate repository](https://github.com/nicolasjulian/argo-howtofixcvedotcom).
- **Bazel:** Used to manage Golang dependencies in this monorepo.
- **Cloudflare:** Provides DNS, SSL, and Tunnels.

## Deployment

Whenever changes are made to the backend or frontend:

1. The corresponding workflow is triggered.
2. The new version is built and the image pushed to DockerHub.
3. The deploy stage updates image tags in the ArgoCD repository.
4. Finally, ArgoCD applies a blue-green deployment to roll out the changes.

## About the Project

> **Disclaimer:** The frontend design is terrible – this project is mostly for fun!

### Structure

```plaintext
.
├── backend/
│   ├── configs/
│   ├── controllers/
│   ├── utils/
│   ├── main.go
│   ├── go.mod
│   └── go.sum
├── frontend/
│   ├── app.js
│   ├── node_modules/
│   ├── package-lock.json
│   ├── package.json
│   └── public/
├── BUILD.bazel
├── README.md
├── WORKSPACE.bazel
```

- **backend**: Houses backend code, configs, controllers, utilities, and the main entry point.
- **frontend**: Contains the frontend's main JavaScript file and static assets.
- **Bazel Configuration**: `BUILD.bazel` and `WORKSPACE.bazel` contain Bazel-specific configurations.

## Getting Started

### Prerequisites

- **Bazel:** Ensure [Bazel](https://bazel.build/) is installed on your machine.

### Setup & Execution

1. Clone the repo:
   ```bash
   git clone https://github.com/yourusername/howtofixcvedotcom.git
   cd howtofixcvedotcom
   ```

2. Build both backend and frontend:
   ```bash
   bazel build backend/...
   ```

3. Execute the backend:
   ```bash
   export AUTH_USER=<your_auth_user>
   export AUTH_PASSWORD=<your_auth_password>
   export CVE_API="https://www.opencve.io"

   bazel run backend:backend_howtofixcvedotcom
   ```

4. Bazel & Gazelle Commands:
   ```bash
   bazel run //:gazelle
   bazel query "//backend:*"
   bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=deps.bzl%go_dependencies
   ```

### Access

Once executed, the HowToFixCVE.com web application should be accessible locally.

## Contributing

Interested in making a contribution? Please consult our [contribution guidelines](CONTRIBUTING.md) for more information.

## License

This project is licensed under the MIT License. Refer to the [LICENSE](LICENSE) file for details.
