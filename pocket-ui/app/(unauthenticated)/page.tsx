"use client";

import React, { useState } from 'react';
import {
  Page,
  Navbar,
  List,
  ListInput,
  Button,
} from 'konsta/react';

export default function IndexPage() {
  return (
    <Page>
      <Navbar
        title="e-pocket"
      />

      <div className='max-w-xl mx-auto my-5 px-4'>
        <h3 className='text-white'>Please login first, before use any feature</h3>
        <h3 className='text-white'>If you dont have any account, <span className='text-blue-300 hover:underline cursor-pointer'>click this to make one</span></h3>

        <List strongIos insetIos>
          <ListInput label="Username" type="text" placeholder="Username" />

          <ListInput
            label="Password"
            type="password"
            placeholder="Your password"
          />
        </List>

        <Button>Login</Button>
      </div>

    </Page>
  );
}