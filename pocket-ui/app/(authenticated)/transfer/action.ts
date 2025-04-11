'use server';

import * as config from '@/app/config/const';
import * as apiConfig from '@/app/config/api';
import { auth } from "@/app/lib/auth";

export interface InquiryAccount {
  success: boolean;
  data?: apiConfig.InquiryAccountResponse;
}

export async function inquiry(accountNumber: string): Promise<InquiryAccount> {
  const session = await auth();
  if (session === null) {
    return { success: false }
  }

  let inquiryResponse: apiConfig.InquiryAccountResponse;
  try {
    const request = await fetch(`${config.API_URL}${config.InquiryURL}?accountNumber=${accountNumber}`, {
      method: "GET",
    });

    if (request.status == 404) {
      console.log(`Account not found: ${accountNumber}`);
      return {
        success: true,
      }
    }

    if (request.status >= 400) {
      console.error(`Error inquiring account: ${request.statusText}`);
      return {
        success: false,
      }
    }

    inquiryResponse = await request.json();
  } catch (error) {
    console.error(`Error inquiring account: ${error}`);

    return {
      success: false,
    }
  }

  return {
    success: true,
    data: inquiryResponse,
  }
}