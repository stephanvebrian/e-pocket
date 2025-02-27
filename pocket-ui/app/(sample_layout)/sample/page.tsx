"use client";

import {
  Page,
  Navbar,
  NavbarBackLink,
  Tabbar,
  TabbarLink,
  Block,
  Icon,
  List,
  ListItem,
  Toggle,
} from 'konsta/react';

export default function SampleIndex() {
  return (
    <Page>
      <Navbar
        title="e-pocket"
      />

      <Tabbar
        className="left-0 bottom-0 fixed"

      >
        <TabbarLink
          label={'Tab 1'}
        />
        <TabbarLink
          label={'Tab 2'}
        />
        <TabbarLink
          label={'Tab 3'}
        />
      </Tabbar>

      <div className='max-w-md mx-auto w-full'>
        <h2 className='text-white'>Test Page</h2>
      </div>

      <List strong inset>
        <ListItem
          title="Tabbar Labels"
        />
        <ListItem
          title="Tabbar Icons"
        />
      </List>
    </Page>
  );
}
