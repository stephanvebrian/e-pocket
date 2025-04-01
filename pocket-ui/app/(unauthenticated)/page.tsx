"use client";
import { signIn } from 'next-auth/react';
import React, { useState } from 'react';
import {
  Page,
  Navbar,
  List,
  ListInput,
  Button,
} from 'konsta/react';

import { generateAccount } from './action';

export default function IndexPage() {
  const [username, setUsername] = useState<string>();
  const [password, setPassword] = useState<string>();

  const handleGenerateAccount = async () => {
    const res = await generateAccount();
    // alert(res.message); // TODO: add alert here

    if (res.success) {
      // show the generated account
    }
  }

  return (
    <Page>
      <Navbar
        title="e-pocket"
      />

      <div className='max-w-xl mx-auto my-5 px-4'>
        <h3 className='text-white'>Please login first, before use any feature</h3>
        <h3 className='text-white'>If you dont have an account, <span className='text-blue-300 hover:underline cursor-pointer' onClick={handleGenerateAccount}>click here to create one</span></h3>

        <List strongIos insetIos>
          <ListInput label="Username" type="text" placeholder="Username" onChange={(e) => setUsername(e.target.value)} />

          <ListInput
            label="Password"
            type="password"
            placeholder="Your password"
            onChange={(e) => setPassword(e.target.value)}
          />
        </List>

        <Button onClick={() => signIn('username-creds', { username, password })}>Login</Button>
      </div>

    </Page>
  );
}

