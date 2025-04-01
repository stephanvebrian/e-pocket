'use server';

interface GenerateAccountResponse {
  success: boolean;
  message: string;
}

export async function generateAccount(): Promise<GenerateAccountResponse> {
  console.log(`ServerAction: GENERATE ACCOUNT`)

  return {
    success: true,
    message: "Account generated successfully",
  }
}