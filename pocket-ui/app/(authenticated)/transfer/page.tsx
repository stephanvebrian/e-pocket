'use client'

import { useEffect, useState } from 'react';
import {
  Block,
  BlockTitle,
  List,
  ListInput
} from 'konsta/react';
import { MdAccountBalanceWallet, MdSavings, MdPerson } from 'react-icons/md';

import * as apiConfig from '@/app/config/api';
import { getAccounts } from '@/app/(authenticated)/wallet/action';

export default function TransferPage() {
  const [accounts, setAccounts] = useState<apiConfig.AccountData[]>([]);

  const fetchAccounts = async () => {
    const { success, data } = await getAccounts();
    if (!success || data === undefined) {
      alert("failed to fetch accounts");
      return
    }

    setAccounts(data.accounts);
  }

  useEffect(() => {
    fetchAccounts();
  }, [])

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
        >
          {accounts.map((account) => {
            return (
              <option key={account.accountNumber} value={account.accountNumber}>{account.accountName} - {account.accountNumber}</option>
            )
          })}
        </ListInput>

        <ListInput label="Destination" type="text" placeholder="Destination Number" />

        <ListInput label="Amount" type="number" placeholder="Amount" />

      </List>
    </>
  );
}
