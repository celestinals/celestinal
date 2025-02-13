<img src="assets/tkx.png" alt="tickex logo" width="120" height="120">

![Node.js](https://img.shields.io/badge/node-%3E%3D20.x-green.svg)
![PNPM](https://img.shields.io/badge/pnpm-%3E%3D9.x-yellow.svg)
![Turborepo](https://img.shields.io/badge/turbo-monorepo-orange.svg)

# 🌐 Tickex

Welcome to the **web monorepo** for Tickex! Built with:

- ⚛️ **Next.js 15**
- 🎨 **Tailwind CSS**
- 🧩 **ShadCN's Monorepo**
- ✅ **ESLint 9**

Enjoy a modern, scalable development workflow with shared resources and optimized builds.

---

## 📚 Table of Contents

- [🚀 Getting Started](#-getting-started)
- [🔧 Requirements](#-requirements)
- [📂 Setup Instructions](#-setup-instructions)
- [💻 Development](#-development)
- [🗂️ Folder Structure](#️-folder-structure)
- [🏗️ Monorepo Features](#️-monorepo-features)

---

## 🚀 Getting Started

Set up the web monorepo for Tickex and start building your applications with ease. Leverage shared components, utilities, and configurations to accelerate your workflow.

---

## 🔧 Requirements

Ensure the following tools are installed:

- **Node.js** (>= 20.x)
- **npm** (>= 10.x) or **yarn** (>= 1.x)
- **Docker engine** (optional for API integration)
- **pnpm** (>= 9.x) for efficient dependency management in the monorepo

---

## 📂 Setup Instructions

### 1. Clone the Repository

Clone the project repository to your local machine:

```bash
git clone https://github.com/tickexvn/tickex-website.git
cd web
```

### 2. Install Dependencies

Use `pnpm` to install dependencies across all workspaces:

```bash
pnpm install
```

### 3. Environment Variables

Set up environment variables by creating a `.env.local` file in the root directory and populating it as follows:

```bash
NEXT_PUBLIC_API_URL=https://api.tickex.com
NEXT_PUBLIC_APP_NAME=tickex
```

### 4. Run the Development Server

Start the development server for apps:

```bash
pnpm dev
```

Visit [http://localhost:3000](http://localhost:3000) in your browser.

---

## 💻 Development

### Monorepo Architecture

This project uses a monorepo structure for efficient code sharing across apps and packages. All reusable components, utilities, and configurations are centralized in the `packages/` directory.

### Code Quality

We enforce code quality and standards using **ESLint 9** and **Prettier**. To lint your code:

```bash
pnpm lint
```

### Styling

**Tailwind CSS** is used for utility-first styling. Modify or extend styles in the `tailwind.config.js` file located in `packages/config/`.

### Component Management

This project leverages **ShadCN's Monorepo** for shared components. Add new components using the command:

```bash
pnpm shadcn add [component-name]
```

Refer to the [ShadCN documentation](https://ui.shadcn.com/docs) for guidance.

### Building for Production

Build the application for production:

```bash
pnpm build
```

The output will be generated in the `.next/` directory for each app.

---

## 🗂️ Folder Structure

```plaintext
web/
├── apps/                  # Applications (e.g., web)
│   ├── web/               # Main web app
│   ├── components/        # Shared UI components for web
│   ├── hooks/             # Shared React hooks for web
│   ├── lib/               # Shared library code specific to the web app
│   └── app/               # Routing and pages structure
├── packages/              # Shared packages across apps
│   ├── components/        # Shared UI components
│   ├── config/            # Shared Tailwind and ESLint configurations
│   ├── hooks/             # Shared React hooks
│   ├── utils/             # Utility functions
│   └── types/             # Shared TypeScript types
├── pnpm-lock.yaml         # Dependency lock file for reproducible installs
├── .eslintrc.js           # ESLint configuration
├── turbo.json             # Turborepo configuration for monorepo tasks
├── package.json           # Root dependencies and scripts
├── pnpm-workspace.yaml    # PNPM workspace configuration
└── README.md              # Project documentation
```

---

## 🏗️ Monorepo Features

### Dependency Management

Using **PNPM** ensures efficient dependency management with:

- **Deduplication**: Dependencies are installed only once and reused across workspaces.
- **Strict Versioning**: Each app/package has isolated version control.

### Turbo-powered Builds

**Turborepo** accelerates builds and task execution by caching and optimizing tasks such as `build`, `lint`, and `test`.

### Shared Packages

The `packages/` directory centralizes reusable code, reducing duplication and enhancing maintainability.

### Incremental Development

Workspaces enable independent development and testing of apps/packages, allowing teams to iterate faster.

---

Made in 🇻🇳 🚀
