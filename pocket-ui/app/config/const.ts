
export const AUTH_SECRET: string = process.env.AUTH_SECRET as string || "";
export const API_URL: string = process.env.API_URL as string || "";

export const GenerateAccountURL: string = "v1/account/generate";
export const ListAccountURL: string = "/v1/account";
export const InquiryURL: string = "/v1/account/inquiry";
export const GetRandomAccountURL: string = "/v1/account/random";

export const ValidateUserURL: string = "/v1/user/validate";

export const CreateTransferURL: string = "/v1/transfer";

export const ListTransactionHistoryURL: string = "/v1/transaction/history";
