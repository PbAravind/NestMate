# Firebase Auth Setup Guide

This guide explains how to set up Firebase Authentication for the NestMate backend.

## Prerequisites

1. A Firebase project
2. Firebase Admin SDK service account credentials

## Step 1: Create Firebase Project

1. Go to [Firebase Console](https://console.firebase.google.com/)
2. Click "Create a project" or select an existing project
3. Enable Authentication in the Firebase console
4. Configure sign-in methods (Email/Password, Google, etc.)

## Step 2: Generate Service Account Key

1. In Firebase Console, go to Project Settings (gear icon)
2. Navigate to "Service accounts" tab
3. Click "Generate new private key"
4. Download the JSON file containing your credentials

## Step 3: Configure Environment Variables

From the downloaded JSON file, extract these values and set them as environment variables:

```bash
# From the JSON file
FIREBASE_PROJECT_ID=your-project-id
FIREBASE_PRIVATE_KEY="-----BEGIN PRIVATE KEY-----\nYour-Private-Key-Here\n-----END PRIVATE KEY-----\n"
FIREBASE_CLIENT_EMAIL=firebase-adminsdk-xxxxx@your-project.iam.gserviceaccount.com
```

**Important Notes:**
- The private key must include the `\n` characters for line breaks
- Wrap the private key in double quotes
- Keep the `-----BEGIN PRIVATE KEY-----` and `-----END PRIVATE KEY-----` markers

## Step 4: Client-Side Integration

The backend expects Firebase ID tokens from client applications. Clients should:

1. Initialize Firebase SDK with your project configuration
2. Authenticate users using Firebase Auth methods
3. Get the ID token: `await user.getIdToken()`
4. Send the ID token to backend endpoints in the Authorization header:
   ```
   Authorization: Bearer <firebase-id-token>
   ```

## API Endpoints

### Public Endpoints (No Authentication Required)

- `POST /api/v1/auth/register` - Register new user
- `POST /api/v1/auth/login` - Login with Firebase ID token

### Protected Endpoints (Require Authentication)

- `POST /api/v1/auth/logout` - Logout user
- `POST /api/v1/auth/refresh` - Refresh token
- `GET /api/v1/auth/profile` - Get user profile
- `PUT /api/v1/auth/profile` - Update user profile

## Authentication Flow

### Registration Flow
1. Client creates user with Firebase Auth
2. Client gets ID token from Firebase
3. Client sends registration request with user details
4. Backend creates user in Firebase Auth and local database
5. Backend returns custom token for API access

### Login Flow
1. Client authenticates with Firebase Auth
2. Client gets ID token from Firebase
3. Client sends ID token to `/api/v1/auth/login`
4. Backend verifies token and returns custom token for API access

### API Access Flow
1. Client includes Firebase ID token in Authorization header
2. Backend middleware verifies token with Firebase
3. Backend extracts user information and allows access

## Error Handling

The backend handles various authentication errors:

- `MISSING_AUTH_HEADER` - No Authorization header provided
- `INVALID_AUTH_FORMAT` - Invalid header format
- `INVALID_TOKEN` - Invalid or expired Firebase token
- `USER_NOT_FOUND` - User doesn't exist in database
- `REGISTRATION_FAILED` - User registration failed
- `LOGIN_FAILED` - Login process failed

## Security Considerations

1. **Token Validation**: All tokens are verified with Firebase
2. **User Isolation**: Users can only access their own data
3. **Token Expiration**: Tokens have limited lifetime
4. **Secure Storage**: Store credentials securely in environment variables
5. **HTTPS Only**: Use HTTPS in production for token transmission

## Testing

For testing without real Firebase credentials:
1. The service gracefully handles missing/invalid credentials
2. Mock implementations are available for unit testing
3. Integration tests can use Firebase Auth emulator

## Troubleshooting

### Common Issues

1. **"Failed to initialize Firebase app"**
   - Check that all environment variables are set correctly
   - Verify the private key format includes proper line breaks

2. **"Invalid token"**
   - Ensure client is sending Firebase ID token, not custom token
   - Check token hasn't expired
   - Verify project ID matches between client and server

3. **"User not found"**
   - User exists in Firebase but not in local database
   - Registration process may have failed partway through

### Debug Mode

Set `GIN_MODE=debug` to see detailed request/response logs.