## Architecture

Frontend: React + Bootstrap
Backend: Go (Gin)
Database: MongoDB Atlas
Authentication: JWT + Refresh Tokens
Deployment: Vercel + Render


ARCHITECTURE DIAGRAM:

┌─────────────┐
│    User     │
└──────┬──────┘
       │
       ▼
┌──────────────────────────┐
│ React Frontend (Vercel)  │
├──────────────────────────┤
│ • Login                  │
│ • Register               │
│ • Browse Movies          │
│ • Recommendations        │
│ • Stream Movies          │
└──────────┬───────────────┘
           │ HTTPS / Axios
           ▼
┌──────────────────────────┐
│ Go Backend (Gin + Render)│
├──────────────────────────┤
│ • JWT Authentication     │
│ • Refresh Tokens         │
│ • Movie APIs             │
│ • Recommendation Engine  │
│ • User Management        │
└──────────┬───────────────┘
           │
           ▼
┌──────────────────────────┐
│ MongoDB Atlas            │
├──────────────────────────┤
│ • Users Collection       │
│ • Movies Collection      │
│ • Genres Collection      │
└──────────────────────────┘
