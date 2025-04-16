'use client'

import { useEffect, useState } from 'react';
import {
  Block,
  Card,
  List,
  ListItem,
  Button
} from 'konsta/react';
import { FaPlus, FaMinus } from "react-icons/fa";

import { TransactionHistory, getTransactionHistory } from '@/app/(authenticated)/history/action'

export default function HistoryPage() {
  const [transactionHistory, setTransactionHistory] = useState<TransactionHistory>()

  useEffect(() => {
    const fetchTransactionHistory = async () => {
      const transactionHistory = await getTransactionHistory();
      if (!transactionHistory.success) {
        alert("failed to fetch transaction history");
        return
      }
      if (transactionHistory.data === undefined) {
        alert("failed to fetch transaction history");
        return
      }

      setTransactionHistory(transactionHistory);
    }

    fetchTransactionHistory();
  }, [])

  return (
    <>
      <Block strong className="space-y-4">
        <h2 className='font-bold text-lg'>My History</h2>
      </Block>
      {transactionHistory?.data?.transactionHistory.length != 0 && (
        <List>
          {transactionHistory?.data?.transactionHistory.map((transaction) => {
            let media = <FaPlus />
            if (transaction.transactionType === "OUTGOING") {
              media = <FaMinus />
            }

            return (
              <ListItem key={transaction.id} title={`${transaction.transactionType} Rp ${transaction.amount}`} media={media} />
            )
          })}
        </List>
      )}
    </>
  )
}