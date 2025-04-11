'use server';

import * as config from '@/app/config/const';
import * as apiConfig from '@/app/config/api';
import { auth } from "@/app/lib/auth";

export interface GetAccounts {
  success: boolean;
  data?: apiConfig.ListAccountResponse;
}

export async function getAccounts(): Promise<GetAccounts> {
  const session = await auth();
  if (session === null) {
    return { success: false }
  }

  let accountsResponse: apiConfig.ListAccountResponse;
  try {
    const request = await fetch(`${config.API_URL}${config.ListAccountURL}?userID=${session.user.id}`, {
      method: "GET",
    });

    if (request.status >= 400) {
      console.error(`Error fetching accounts: ${request.statusText}`);
      return {
        success: false,
      }
    }

    accountsResponse = await request.json();
  } catch (error) {
    console.error(`Error fetching accounts: ${error}`);

    return {
      success: false,
    }
  }

  return {
    success: true,
    data: accountsResponse,
  }
}