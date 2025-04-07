import type {
  GetServerSidePropsContext,
  NextApiRequest,
  NextApiResponse,
} from "next"
import NextAuth from 'next-auth';
import Credentials from 'next-auth/providers/credentials';
import type { NextAuthOptions } from "next-auth"
import { getServerSession } from "next-auth"

import * as config from '@/app/config/const'


export const authOptions: NextAuthOptions = {
  providers: [
    Credentials({
      id: "username-creds",
      name: 'Username Credentials',
      credentials: {
        username: { label: 'Username', type: 'text' },
        password: { label: 'Password', type: 'password' },
      },
      authorize: async (credentials) => {
        if (!credentials) {
          // TODO: define appropriate error
          return null;
        }

        // validate credentials must be filled
        if (!credentials.username || !credentials.password) {
          return null;
        }


        if (credentials.username === "username" && credentials.password === "password") {
          return { id: "1", username: "username" };
        }

        return null;
      },
    }),
  ],
  pages: {
    signIn: "/",
    signOut: "/logout",
  },
  callbacks: {
    async jwt({ token, user }) {
      // If the user object is available (during sign-in)
      if (user) {
        // token.username = user.username;
        token.user = {
          id: user.id || "",
          username: user.username,
        }
      }
      return token;
    },
    async session({ session, token }) {
      return {
        ...session,
        user: {
          ...session.user,
          id: token.user.id,
          username: token.user.username,
        },
      };
    },
  },
  secret: config.AUTH_SECRET,
}

export const handler = NextAuth(authOptions);

// Use it in server contexts
export function auth(
  ...args:
    | [GetServerSidePropsContext["req"], GetServerSidePropsContext["res"]]
    | [NextApiRequest, NextApiResponse]
    | []
) {
  return getServerSession(...args, authOptions)
}
