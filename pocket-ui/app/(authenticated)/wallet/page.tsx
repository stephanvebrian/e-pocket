'use client'

import { useEffect, useState } from 'react';
import {
  Block,
  List, ListItem, Icon,
} from 'konsta/react';
import { MdAccountBalanceWallet } from 'react-icons/md';

import * as apiConfig from '@/app/config/api';
import { getAccounts } from './action'

export default function WalletPage() {
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
        <h2 className='font-bold text-lg'>My Wallet</h2>
      </Block>
      <div>
        {accounts.map((account) => {
          return (
            <List key={account.accountNumber} strong className='!my-2'>
              <ListItem
                link
                media={<Icon material={<MdAccountBalanceWallet />} className="text-5xl" />}
                header={`${account.accountName} - ${account.accountNumber}`}
                title={`Rp ${new Intl.NumberFormat('id-ID', { minimumFractionDigits: 2, maximumFractionDigits: 2, }).format(account.balance / 100)}`}
              // footer={"Tabungan"}
              />
            </List>
          )
        })}
      </div>
    </>
  );
}
