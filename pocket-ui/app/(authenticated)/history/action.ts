'use server';

import * as config from '@/app/config/const';
import * as apiConfig from '@/app/config/api';
import { auth } from "@/app/lib/auth";

export interface TransactionHistory {
  success: boolean;
  data?: apiConfig.TransactionHistoryResponse;
}

export async function getTransactionHistory(): Promise<TransactionHistory> {
  const session = await auth();
  if (session === null) {
    return { success: false }
  }

  let transactionHistoryResponse: apiConfig.TransactionHistoryResponse;
  try {
    const request = await fetch(`${config.API_URL}${config.ListTransactionHistoryURL}?userID=${session.user.id}`, {
      method: "GET",
    });

    transactionHistoryResponse = await request.json();
  } catch (error) {
    console.error(`Error fetching transaction history: ${error}`);
    return {
      success: false,
    }
  }

  return {
    success: true,
    data: transactionHistoryResponse,
  }
}