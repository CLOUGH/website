# Get Firebase Web App Configuration

## Step 1: Access Firebase Console

1. Go to [Firebase Console](https://console.firebase.google.com/)
2. Select your project: **my-website-a3970**

## Step 2: Add Web App (if not already done)

1. In the project overview, click the **Settings gear** → **Project settings**
2. Scroll down to **"Your apps"** section
3. If you don't see a web app (</> icon), click **"Add app"** → **Web**
4. Enter app nickname: "My Website"
5. **Check** "Also set up Firebase Hosting"
6. Click **"Register app"**

## Step 3: Get Configuration

You'll see a configuration object like this:

```javascript
// Your Firebase config
const firebaseConfig = {
  apiKey: "AIzaSyC...",
  authDomain: "my-website-a3970.firebaseapp.com",
  projectId: "my-website-a3970",
  storageBucket: "my-website-a3970.appspot.com",
  messagingSenderId: "123456789",
  appId: "1:123456789:web:abcdef123456"
};
```

## Step 4: Update Your .env File

Replace the dummy values in your `.env` file:

```bash
FIREBASE_API_KEY=AIzaSyC...                    # Copy from apiKey
FIREBASE_MESSAGING_SENDER_ID=123456789         # Copy from messagingSenderId  
FIREBASE_APP_ID=1:123456789:web:abcdef123456   # Copy from appId
```

## Step 5: Enable Authentication

1. In Firebase Console, go to **Authentication** → **Sign-in method**
2. Click on **Google** provider
3. Click **Enable**
4. Add your email to **Project settings** → **Users and permissions**
5. Set **Authorized domains** (add your domain when deploying)

## Step 6: Test the Application

```bash
# Start your server
go run main.go

# Visit: http://localhost:8080/admin/login
```

You should now see the Google Sign-in button working properly!

## Troubleshooting

If you still get errors:

1. **Check browser console** for any remaining Firebase errors
2. **Verify** all environment variables are set correctly
3. **Ensure** Authentication is enabled in Firebase Console
4. **Check** that your email has admin access (update the `IsUserAdmin` function if needed)