'use client'

import { useEffect, useState } from 'react';
import {
  Block,
  Card,
  List,
  ListInput
} from 'konsta/react';
import { MdAccountBalanceWallet, MdSavings, MdPerson } from 'react-icons/md';

import * as apiConfig from '@/app/config/api';
import { getAccounts } from '@/app/(authenticated)/wallet/action';
import { inquiry } from '@/app/(authenticated)/transfer/action';
import { resolve } from 'path';

export default function TransferPage() {
  const [accounts, setAccounts] = useState<apiConfig.AccountData[]>([]);

  const [sourceNumber, setSourceNumber] = useState<string>('')
  const [destinationNumber, setDestinationNumber] = useState<string>('')
  const [destinationAccount, setDestinationAccount] = useState<apiConfig.InquiryAccountResponse>()
  const [amount, setAmount] = useState<number>(0)

  const fetchAccounts = async () => {
    const { success, data } = await getAccounts();
    if (!success || data === undefined) {
      alert("failed to fetch accounts");
      return
    }

    setAccounts(data.accounts);
    if (data.accounts.length > 0) {
      setSourceNumber(data.accounts[0].accountNumber);
    }
  }

  const resolveDestination = async (accountNumber: string) => {
    const { success, data } = await inquiry(accountNumber);
    if (!success) {
      alert("failed to resolve destination account");
      return
    }

    if (!data) {
      setDestinationAccount(undefined);
      return;
    }

    setDestinationAccount(data);
  }

  useEffect(() => {
    fetchAccounts();
  }, [])

  useEffect(() => {
    if (!destinationNumber) return;

    // debounce
    const handler = setTimeout(() => {
      resolveDestination(destinationNumber);
    }, 500); // delay in ms for debounce

    return () => clearTimeout(handler); // clear previous timeout if input changes
  }, [destinationNumber]);

  return (
    <>
      <Block strong className="space-y-4">
        <h2 className='font-bold text-lg'>Transfer Money</h2>
      </Block>
      <List strongIos insetIos>
        <ListInput
          label="Sender"
          type="select"
          dropdown
          onChange={(e) => {
            setSourceNumber(e.target.value)
          }}
          value={sourceNumber}
        >
          {accounts.map((account) => {
            return (
              <option key={account.accountNumber} value={account.accountNumber}>{account.accountName} - {account.accountNumber}</option>
            )
          })}
        </ListInput>

        <ListInput label="Destination" type="text" placeholder="Destination Number" onChange={(e) => { setDestinationNumber(e.target.value) }} value={destinationNumber} />
        {destinationAccount && (
          <Card>
            <p>Destination Account:</p>
            <p>{destinationAccount.accountName} - {destinationAccount.accountNumber}</p>
          </Card>
        )}

        <ListInput label="Amount" type="number" placeholder="Amount" onChange={(e) => { setAmount(e.target.valueAsNumber) }} value={amount} />

      </List>
    </>
  );
}
