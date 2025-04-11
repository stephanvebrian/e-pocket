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