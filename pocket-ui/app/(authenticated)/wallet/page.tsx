'use client'

import {
  Block,
  List, ListItem, Icon,
} from 'konsta/react';

import { MdAccountBalanceWallet, MdSavings, MdPerson } from 'react-icons/md';

export default function WalletPage() {
  return (
    <>
      <Block strong className="space-y-4">
        <h2 className='font-bold text-lg'>My Wallet</h2>
      </Block>
      <div>
        <List strong className='!my-2'>
          <ListItem
            link
            media={<Icon material={<MdAccountBalanceWallet />} className="text-5xl" />}
            header={"Tabungan Utama"}
            title={"Rp30.522,52"}
          // footer={"Tabungan"}
          />
        </List>
        <List strong className='!my-2'>
          <ListItem
            link
            media={<Icon material={<MdAccountBalanceWallet />} className="text-5xl" />}
            header={"Tabungan Utama"}
            title={"Rp30.522,52"}
          // footer={"Tabungan"}
          />
        </List>
      </div>
    </>
  );
}
