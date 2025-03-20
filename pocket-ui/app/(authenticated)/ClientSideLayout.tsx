'use client'

import React, { useState, useEffect } from 'react';
import { usePathname, useRouter } from 'next/navigation';
import {
  Page,
  Navbar,
  Tabbar,
  TabbarLink,
  Block,
  Icon,
} from 'konsta/react';
import { MdLibraryBooks, MdAttachMoney, MdHistory } from 'react-icons/md';

interface ClientSideLayoutProps {
  children: React.ReactNode;
}

const tabBarMap = {
  "/wallet": {
    tabName: 'wallet-tab',
  },
  "/transfer": {
    tabName: 'transfer-tab',
  },
  "/history": {
    tabName: 'history-tab',
  },
} as const;

type tabBarPath = keyof typeof tabBarMap;

const showTabBarLabels = true
const showTabBarIcons = true

export default function ClientSideLayout(props: ClientSideLayoutProps) {
  const pathname = usePathname(); // Get the current URL path
  const router = useRouter();

  const [activeTab, setActiveTab] = useState<string | null>(null); // Initialize as null
  const [isLoading, setIsLoading] = useState(true); // Loading state

  // Update the active tab based on the URL path
  useEffect(() => {
    // TODO: consider to have 404 page instead
    if (!(pathname in tabBarMap)) {
      console.log(`invalid pathname: `, pathname);
      return;
    }

    // set active tab
    const tabBarData = tabBarMap[pathname as tabBarPath];
    setActiveTab(tabBarData.tabName);
    setIsLoading(false);
  }, [pathname]);

  return (
    <Page>
      <Navbar
        title="e-pocket"
      />

      <Tabbar
        labels={showTabBarLabels}
        icons={showTabBarIcons}
        className="left-0 bottom-0 fixed"
      >
        <TabbarLink
          active={activeTab === 'wallet-tab'}
          onClick={() => { setActiveTab('wallet-tab'); router.push("/wallet"); }}
          icon={
            <Icon
              material={<MdLibraryBooks className="w-6 h-6" />}
            />
          }
          label={showTabBarLabels && 'Wallet'}
        />
        <TabbarLink
          active={activeTab === 'transfer-tab'}
          onClick={() => { setActiveTab('transfer-tab'); router.push("/transfer"); }}
          icon={
            <Icon
              material={<MdAttachMoney className="w-6 h-6" />}
            />
          }
          label={showTabBarLabels && 'Transfer'}
        />
        <TabbarLink
          active={activeTab === 'history-tab'}
          onClick={() => { setActiveTab('history-tab'); router.push("/history"); }}
          icon={
            <Icon
              material={<MdHistory className="w-6 h-6" />}
            />
          }
          label={showTabBarLabels && 'History'}
        />
      </Tabbar>

      <div className='max-w-2xl mx-auto'>
        {isLoading && (
          <Block strong inset className="space-y-4">
            <p>
              <b>Loading...</b>
            </p>
            <p>
              <span>
                Please wait while we load the data...
              </span>
            </p>
          </Block>
        )}

        {!isLoading && props.children}
      </div>
    </Page>
  );
}