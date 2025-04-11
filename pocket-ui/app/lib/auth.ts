import type {
  GetServerSidePropsContext,
  NextApiRequest,
  NextApiResponse,
} from "next"
import NextAuth from 'next-auth';
import Credentials from 'next-auth/providers/credentials';
import type { NextAuthOptions } from "next-auth"
import { getServerSession } from "next-auth"
import ky from 'ky';

import * as config from '@/app/config/const'
import * as apiConfig from '@/app/config/api';

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

        let validateUser: apiConfig.ValidateUserResponse;
        try {
          // there is a bug on ky package, hence switch to use native fetch instead
          // const validateUserRequest = await ky.post(config.ValidateUserURL, { prefixUrl: config.API_URL, json: credentials });
          // validateUser = await validateUserRequest.json<apiConfig.ValidateUserResponse>();
          const validateUserRequest = await fetch(`${config.API_URL}${config.ValidateUserURL}`, {
            body: JSON.stringify(credentials),
            method: "POST",
          })
          validateUser = await validateUserRequest.json();
        } catch (error) {
          console.error(`Error validating user: ${error}`);
          throw new Error("Error validating user");
        }

        if (validateUser.isValid === false) {
          throw new Error("Invalid username or password");
        }

        return { id: validateUser.userID, username: credentials.username };
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
