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

export interface TransferRequest {
  idempotencyKey: string;
  senderAccountNumber: string;
  receiverAccountNumber: string;
  amount: number;
}

export interface TransferResponse {
  success: boolean;
  data?: apiConfig.CreateTransferResponse;
}

export async function transfer(request: TransferRequest): Promise<TransferResponse> {
  const session = await auth();
  if (session === null) {
    return { success: false }
  }

  let transferRequest: apiConfig.CreateTransferRequest = {
    idempotencyKey: request.idempotencyKey,
    sender: {
      number: request.senderAccountNumber,
    },
    receiver: {
      number: request.receiverAccountNumber,
    },
    amount: request.amount,
    userID: session.user.id,
  }
  let transferResponse: apiConfig.CreateTransferResponse;
  try {
    const request = await fetch(`${config.API_URL}${config.CreateTransferURL}`, {
      method: "POST",
      body: JSON.stringify(transferRequest),
    })

    transferResponse = await request.json();
  } catch (error) {
    console.error(`Error creating transfer: ${error}`);
    return {
      success: false,
    }
  }

  return {
    success: true,
    data: transferResponse,
  }
}