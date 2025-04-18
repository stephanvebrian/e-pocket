export type GenerateAccountResponse = {
  userID: string;
  username: string;
  password: string;
  accountNumber: string;
  name: string;
  balance: number;
  status: string;
}

export type ValidateUserResponse = {
  isValid: boolean;
  userID: string;
}

export type ListAccountResponse = {
  accounts: AccountData[];
}

export type AccountData = {
  accountNumber: string;
  accountName: string;
  balance: number;
  status: string;
}

export type InquiryAccountResponse = {
  accountNumber: string;
  accountName: string;
}

export type CreateTransferRequest = {
  idempotencyKey: string;
  sender: CreateTransferAccount;
  receiver: CreateTransferAccount;
  amount: number;
  userID: string;
}

export type CreateTransferAccount = {
  number: string;
}

export type CreateTransferResponse = {
  idempotencyKey: string;
  transactionID: string;
  status: string;
}
export type TransactionHistoryResponse = {
  transactionHistory: TransactionData[];
}

export type TransactionData = {
  id: string;
  userID: string;
  account: AccountData;
  transactionType: string;
  amount: number;
  endingBalance: number;
  status: string;
  createdAt: string;
  updatedAt: string;
}

export type RandomAccountResponse = {
  account: AccountData;
}