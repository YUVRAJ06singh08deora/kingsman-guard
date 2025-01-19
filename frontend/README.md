## Lock Management System
 
This system provides management of smart locks, enabling users to configure locks, set up temporary guest codes, and enable/disable lock automation.

## Project Setup

```sh
npm install
```

## .env Setup

```sh
VITE_BACKEND_URL=<Hosted Backend URL LINK>
VITE_COGNITO_AUTHORITY=<Cognito AUTHORITY LINK (Get It from CONITO setup Page on AWS CONSOLE)>
VITE_COGNITO_CLIENT_ID=<Client ID (Get It from CONITO setup Page on AWS CONSOLE) >
VITE_COGNITO_DOMAIN=<Cognito Domain (Get It from CONITO setup Page on AWS CONSOLE)>

```

### Compile and Hot-Reload for Development

```sh
npm run dev
```

### Type-Check, Compile and Minify for Production

```sh
npm run build
```

### Lint with ESLint

```sh
npm run lint
```

### Format with Prettier

```sh
npm run format
```