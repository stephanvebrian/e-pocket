'use client'

import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation'
import {
  Block,
  Card,
  List,
  ListInput,
  Button
} from 'konsta/react';
import { MdAccountBalanceWallet, MdSavings, MdPerson } from 'react-icons/md';
import { v6 as uuidv6 } from 'uuid';

import * as apiConfig from '@/app/config/api';
import { getAccounts } from '@/app/(authenticated)/wallet/action';
import { inquiry, transfer, randomAccount } from '@/app/(authenticated)/transfer/action';

export default function TransferPage() {
  const router = useRouter()

  const [accounts, setAccounts] = useState<apiConfig.AccountData[]>([]);

  const [sourceAccount, setSourceAccount] = useState<apiConfig.AccountData>()
  const [destinationNumber, setDestinationNumber] = useState<string>('')
  const [destinationAccount, setDestinationAccount] = useState<apiConfig.InquiryAccountResponse>()
  const [amount, setAmount] = useState<number>(0)
  const [enableTransfer, setEnableTransfer] = useState<boolean>(false)

  const fetchAccounts = async () => {
    const { success, data } = await getAccounts();
    if (!success || data === undefined) {
      alert("failed to fetch accounts");
      return
    }

    setAccounts(data.accounts);
    if (data.accounts.length > 0) {
      setSourceAccount(data.accounts[0]);
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

  const handleRandomAccount = async () => {
    const { success, data } = await randomAccount();
    if (!success) {
      alert("failed to generate random account");
      return
    }

    if (!data) {
      alert("failed to generate random account");
      return
    }

    setDestinationNumber(data.account.accountNumber);
  }

  const handleTransfer = async () => {
    if (sourceAccount === undefined) {
      alert("Please select source account");
      return;
    }

    if (destinationAccount === undefined) {
      alert("Please select destination account");
      return;
    }

    setEnableTransfer(false)

    const response = await transfer({
      idempotencyKey: uuidv6(),
      senderAccountNumber: sourceAccount.accountNumber,
      receiverAccountNumber: destinationAccount.accountNumber,
      amount: amount * 100,
    })

    if (!response.success || response.data === undefined) {
      alert("Transfer failed");
      return;
    }

    // redirect to history page
    router.push('/history')
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

  useEffect(() => {
    if (sourceAccount === undefined) {
      setEnableTransfer(false);
      return;
    }
    if (destinationNumber == '') {
      setEnableTransfer(false);
      return;
    }
    if (destinationAccount === undefined) {
      setEnableTransfer(false);
      return;
    }
    if (amount <= 0) {
      setEnableTransfer(false);
      return;
    }
    if (amount > sourceAccount.balance / 100) {
      setEnableTransfer(false);
      return;
    }

    setEnableTransfer(true);
  }, [sourceAccount, destinationNumber, destinationAccount, amount])

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
            setSourceAccount(accounts.find((account) => account.accountNumber === e.target.value))
          }}
          value={sourceAccount?.accountNumber || ''}
        >
          {accounts.map((account) => {
            const balanceStr = `Rp ${new Intl.NumberFormat('id-ID', { minimumFractionDigits: 2, maximumFractionDigits: 2, }).format(account.balance / 100)}`
            return (
              <option key={account.accountNumber} value={account.accountNumber}>{account.accountName} - {account.accountNumber} - {balanceStr}</option>
            )
          })}
        </ListInput>

        <div className='flex w-full'>
          <div className='w-4/5'>
            <ListInput label="Destination" type="text" placeholder="Destination Number" onChange={(e) => { setDestinationNumber(e.target.value) }} value={destinationNumber} />
          </div>
          <div className='flex items-center justify-center'>
            <Button onClick={handleRandomAccount}>Randomize</Button>
          </div>
        </div>
        {destinationAccount && (
          <Card>
            <p>Destination Account:</p>
            <p>{destinationAccount.accountName} - {destinationAccount.accountNumber}</p>
          </Card>
        )}

        <ListInput label="Amount" type="number" placeholder="Amount" onChange={(e) => { setAmount(e.target.valueAsNumber) }} value={amount || 0} />
      </List>

      <Button disabled={!enableTransfer} onClick={handleTransfer}>Transfer</Button>
    </>
  );
}
