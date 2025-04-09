'use server';

import ky from 'ky';

import * as config from '@/app/config/const';
import * as apiConfig from '@/app/config/api';

interface GenerateAccountResponse {
  success: boolean;
  message: string;
  data?: apiConfig.GenerateAccountResponse;
}

export async function generateAccount(): Promise<GenerateAccountResponse> {
  let account: apiConfig.GenerateAccountResponse;
  try {
    const request = await ky.post(config.GenerateAccountURL, { prefixUrl: config.API_URL });
    account = await request.json<apiConfig.GenerateAccountResponse>();
  } catch (error) {
    console.error(`Error generating account: ${error}`);

    return {
      success: false,
      message: "Error generating account",
    }
  }

  return {
    success: true,
    message: "Account generated successfully",
    data: account,
  }
}