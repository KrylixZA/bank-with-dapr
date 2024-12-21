# Features

## Account Management Service

- **Create Account**
- **Get Account Details**
- **Update Account**
- **Delete Account**
- **Dapr Building Blocks**:
  - State Management (Redis)
  - Actors (for representing individual accounts)

## Transaction Service

- **Create Transaction**
- **Get Transaction History**
- **Update Transaction**
- **Delete Transaction**
- **Dapr Building Blocks**:
  - State Management (Redis)
  - PubSub (for transaction events)

## User Management Service

- **Create User**
- **Get User Details**
- **Update User**
- **Delete User**
- **Dapr Building Blocks**:
  - State Management (Redis)

## Notification Service

- **Send Notifications (Email/SMS)**
- **Get Notification History**
- **Dapr Building Blocks**:
  - PubSub (for notification events)
  - Service Invocation (for sending notifications)

## Reporting Service

- **Generate Account Statements**
- **Generate Transaction Reports**
- **Dapr Building Blocks**:
  - Service Invocation (for generating reports)

## Authentication and Authorization Service

- **User Authentication**
- **Role-Based Access Control**
- **Dapr Building Blocks**:
  - State Management (Redis)
  - Service Invocation (for authentication)
