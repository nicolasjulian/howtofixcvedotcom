# HowToFixCVE.com Monorepo

This monorepo contains the source code for HowToFixCVE.com, a project for searching and fixing Common Vulnerabilities and Exposures (CVEs). It is organized into two main components: `backend` and `frontend`, managed using Bazel build system.

**Notes : The Front end is very very terrible**
## Project Structure

The repository is structured as follows:

```
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

- `backend`: Contains the backend code for HowToFixCVE.com.
  - `configs`: Configuration files for the backend.
  - `controllers`: Controllers for handling API requests.
  - `utils`: Utility functions and libraries.
  - `main.go`: The entry point for the backend application.
- `frontend`: Contains the frontend code for HowToFixCVE.com.
  - `app.js`: The main JavaScript file for the frontend.
  - `public`: Static assets and HTML templates.
- `BUILD.bazel`: Bazel build configuration file.
- `WORKSPACE.bazel`: Bazel workspace configuration file.

## Getting Started

Follow these steps to set up and run the project locally:

### Prerequisites

- [Bazel](https://bazel.build/): You need to have Bazel installed on your system.

### Building and Running

1. Clone this repository:

   ```bash
   git clone https://github.com/yourusername/howtofixcvedotcom.git
   cd howtofixcvedotcom
   ```

2. Build the backend and frontend:

   ```bash
   bazel build //backend:main
   bazel build //frontend:app
   ```

3. Run the backend:

   ```bash
   bazel-bin/backend/main
   ```

4. Run the frontend:

   ```bash
   bazel-bin/frontend/app
   ```

Now, you should be able to access the HowToFixCVE.com web application locally.

## Contributing

We welcome contributions! If you'd like to contribute to this project, please follow our [contribution guidelines](CONTRIBUTING.md).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
