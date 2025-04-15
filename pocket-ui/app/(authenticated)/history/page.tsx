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

export default function HistoryPage() {
  return (
    <>
      <Block strong className="space-y-4">
        <h2 className='font-bold text-lg'>My History</h2>
      </Block>
      <List>
        <ListItem title="Outgoing Rp 10.000,00" media={<FaPlus />} />
        <ListItem title="Incoming Rp 5.000,00" media={<FaMinus />} />
        <ListItem title="Outgoing Rp 1.000,00" media={<FaPlus />} />
      </List>
    </>
  )
}