"use client";
import { signIn } from 'next-auth/react';
import React, { useState } from 'react';
import {
  Page,
  Navbar,
  List,
  ListInput,
  Button,
  Block,
} from 'konsta/react';

import { generateAccount } from './action';

export default function IndexPage() {
  const [username, setUsername] = useState<string>();
  const [password, setPassword] = useState<string>();
  const [showGenerateAccount, setShowGenerateAccount] = useState<boolean>(false);

  const handleGenerateAccount = async () => {
    const response = await generateAccount();

    if (!response.success) {
      // show the generated account
      alert("failed to generate an account");
      return
    }

    if (response.data === undefined) {
      alert("failed to generate an account");
      return
    }

    setUsername(response.data.username);
    setPassword(response.data.password);
    setShowGenerateAccount(true);
  }

  return (
    <Page>
      <Navbar
        title="e-pocket"
      />

      <div className='max-w-xl mx-auto my-5 px-4'>
        <h3 className='text-white'>Please login first, before use any feature</h3>
        <h3 className='text-white'>If you dont have an account, <span className='text-blue-300 hover:underline cursor-pointer' onClick={handleGenerateAccount}>click here to create one</span></h3>
        {showGenerateAccount && (
          <Block>
            <p>
              Username: {username}
            </p>
            <p>
              Password: {password}
            </p>
          </Block>
        )}

        <List strongIos insetIos>
          <ListInput label="Username" type="text" placeholder="Username" onChange={(e) => setUsername(e.target.value)} value={username || ""} />

          <ListInput
            label="Password"
            type="password"
            placeholder="Your password"
            onChange={(e) => setPassword(e.target.value)}
            value={password || ""}
          />
        </List>

        <Button onClick={() => signIn('username-creds', { username, password })}>Login</Button>
      </div>

    </Page>
  );
}

