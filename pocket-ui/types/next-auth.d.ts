// types/next-auth.d.ts
import NextAuth, { type DefaultSession } from "next-auth"

declare module 'next-auth' {
  interface Session {
    user: {
      id: string;
      username: string;
    } & DefaultSession["user"]
  }

  interface User {
    id: string;
    username: string;
  }
}

import { JWT } from "next-auth/jwt"

declare module 'next-auth/jwt' {
  interface JWT {
    user: {
      id: string;
      username: string;
    }
  }
}